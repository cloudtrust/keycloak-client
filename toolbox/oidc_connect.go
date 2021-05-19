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
	"github.com/cloudtrust/keycloak-client"
)

// OidcTokenProvider provides OIDC tokens
type OidcTokenProvider interface {
	ProvideToken(ctx context.Context) (string, error)
	ProvideTokenForRealm(ctx context.Context, realm string) (string, error)
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
	timeout           time.Duration
	perRealmTokenInfo map[string]*oidcTokenInfo
	reqBody           string
	defaultKey        string
	logger            Logger
}

type oidcTokenInfo struct {
	url        string
	oidcToken  oidcToken
	validUntil int64
}

const (
	// Max processing delay: let's assume that the user of OidcTokenProvider will have a maximum of 5 seconds to use the provided OIDC token
	maxProcessingDelay = int64(5)
)

// NewOidcTokenProvider creates an OidcTokenProvider
func NewOidcTokenProvider(config keycloak.Config, realm, username, password, clientID string, logger Logger) OidcTokenProvider {
	var key = "__default__"
	var targets = map[string]string{key: config.AddrTokenProvider[0]}
	return NewOidcTokenProviderMap(targets, key, config.Timeout, realm, username, password, clientID, logger)
}

// NewOidcTokenProviderMap creates an OidcTokenProvider with possible multiple targets
func NewOidcTokenProviderMap(targets map[string]string, defaultKey string, timeout time.Duration, realm, username, password, clientID string, logger Logger) OidcTokenProvider {
	var perRealmTokenInfo = make(map[string]*oidcTokenInfo)
	for targetRealm, keycloakURL := range targets {
		perRealmTokenInfo[strings.ToLower(targetRealm)] = &oidcTokenInfo{
			url: fmt.Sprintf("%s/auth/realms/%s/protocol/openid-connect/token", keycloakURL, realm),
		}
	}

	// If needed, can add &client_secret={secret}
	var body = fmt.Sprintf("grant_type=password&client_id=%s&username=%s&password=%s",
		url.QueryEscape(clientID), url.QueryEscape(username), url.QueryEscape(password))

	return &oidcTokenProvider{
		timeout:           timeout,
		perRealmTokenInfo: perRealmTokenInfo,
		reqBody:           body,
		defaultKey:        defaultKey,
		logger:            logger,
	}
}

func (o *oidcTokenProvider) ProvideToken(ctx context.Context) (string, error) {
	return o.ProvideTokenForRealm(ctx, o.defaultKey)
}

func (o *oidcTokenProvider) ProvideTokenForRealm(ctx context.Context, realm string) (string, error) {
	var oti *oidcTokenInfo
	var ok bool
	if oti, ok = o.perRealmTokenInfo[strings.ToLower(realm)]; !ok {
		if realm == o.defaultKey {
			return "", errorhandler.CreateInternalServerError("unknownRealm")
		}
		return o.ProvideTokenForRealm(ctx, o.defaultKey)
	}
	if time.Now().Unix()+maxProcessingDelay < oti.validUntil {
		return oti.oidcToken.AccessToken, nil
	}

	var mimeType = "application/x-www-form-urlencoded"
	var httpClient = http.Client{
		Timeout: o.timeout,
	}
	var resp, err = httpClient.Post(oti.url, mimeType, strings.NewReader(o.reqBody))
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
	_, _ = buf.ReadFrom(resp.Body)

	err = json.Unmarshal(buf.Bytes(), &oti.oidcToken)
	if err != nil {
		o.logger.Warn(ctx, "msg", fmt.Sprintf("Can't deserialize token. JSON: %s", buf.String()))
		return "", errorhandler.CreateInternalServerError("unexpected.oidcToken")
	}
	oti.validUntil = time.Now().Unix() + oti.oidcToken.ExpiresIn

	return oti.oidcToken.AccessToken, nil
}
