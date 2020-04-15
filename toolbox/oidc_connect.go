package toolbox

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	errorhandler "github.com/cloudtrust/common-service/errors"
	"github.com/cloudtrust/common-service/log"
	"github.com/cloudtrust/keycloak-client"
)

// OidcTokenProvider provides OIDC tokens
type OidcTokenProvider interface {
	ProvideToken(ctx context.Context) (string, error)
}

type oidcToken struct {
	AccessToken      string `json:"access_token,omitempty"`
	ExpiresIn        int64  `json:"expires_in,omitempty"`
	RefreshToken     string `json:"refresh_token,omitempty"`
	RefreshExpiresIn int64  `json:"refresh_expires_in,omitempty"`
	TokenType        string `json:"token_type,omitempty"`
	NotBeforePolicy  int    `json:"not-before-policy,omitempty"`
	SessionState     string `json:"session_state,omitempty"`
	Scope            string `json:"scope,omitempty"`
}

type oidcTokenProvider struct {
	timeout    time.Duration
	tokenURL   string
	reqBody    string
	logger     log.Logger
	oidcToken  oidcToken
	validUntil int64
}

const (
	// Max processing delay: let's assume that the user of OidcTokenProvider will have a maximum of 5 seconds to use the provided OIDC token
	maxProcessingDelay = int64(5)
)

// NewOidcTokenProvider creates an OidcTokenProvider
func NewOidcTokenProvider(config keycloak.Config, realm, username, password, clientID string, logger log.Logger) OidcTokenProvider {
	var urls = strings.Split(config.AddrTokenProvider, " ")
	var keycloakPublicURL = urls[0]

	var tokenURL = fmt.Sprintf("%s/auth/realms/%s/protocol/openid-connect/token", keycloakPublicURL, realm)
	// If needed, can add &client_secret={secret}
	var body = fmt.Sprintf("grant_type=password&client_id=%s&username=%s&password=%s",
		url.QueryEscape(clientID), url.QueryEscape(username), url.QueryEscape(password))

	return &oidcTokenProvider{
		timeout:  config.Timeout,
		tokenURL: tokenURL,
		reqBody:  body,
		logger:   logger,
	}
}

func (o *oidcTokenProvider) ProvideToken(ctx context.Context) (string, error) {
	if o.validUntil+maxProcessingDelay > time.Now().Unix() {
		return o.oidcToken.AccessToken, nil
	}

	var mimeType = "application/x-www-form-urlencoded"
	var httpClient = http.Client{
		Timeout: o.timeout,
	}
	var resp, err = httpClient.Post(o.tokenURL, mimeType, strings.NewReader(o.reqBody))
	if err != nil {
		o.logger.Warn(ctx, "msg", err.Error())
		return "", errorhandler.CreateInternalServerError("unexpected.httpResponse")
	}
	if err == nil && resp.StatusCode == http.StatusUnauthorized {
		o.logger.Warn(ctx, "msg", "Technical user credentials are invalid")
		return "", errorhandler.Error{
			Status:  http.StatusUnauthorized,
			Message: errorhandler.GetEmitter() + ".unauthorized",
		}
	}
	if resp.StatusCode >= 400 || resp.Body == http.NoBody || resp.Body == nil {
		o.logger.Warn(ctx, "msg", fmt.Sprintf("Unexpected behavior: unexpected http status (%d) or response has no body", resp.StatusCode))
		return "", errorhandler.CreateInternalServerError("unexpected.httpResponse")
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	err = json.Unmarshal(buf.Bytes(), &o.oidcToken)
	if err != nil {
		o.logger.Warn(ctx, "msg", fmt.Sprintf("Can't deserialize token. JSON: %s", buf.String()))
		return "", errorhandler.CreateInternalServerError("unexpected.oidcToken")
	}
	o.validUntil = time.Now().Unix() + o.oidcToken.ExpiresIn

	return o.oidcToken.AccessToken, nil
}
