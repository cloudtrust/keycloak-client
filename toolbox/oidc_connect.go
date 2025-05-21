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

	errorhandler "github.com/cloudtrust/common-service/v2/errors"
	"github.com/cloudtrust/keycloak-client/v2"
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
	reqBody           string // Added for fix CLOUDTRUST-6415
	//username          string // Commented for fix CLOUDTRUST-6415
	//password          string // Commented for fix CLOUDTRUST-6415
	defaultKey string
	logger     Logger
}

type oidcTokenInfo struct {
	url        string    // Added for fix CLOUDTRUST-6415
	oidcToken  oidcToken // Added for fix CLOUDTRUST-6415
	validUntil int64     // Added for fix CLOUDTRUST-6415
	forwarded  string
	//oauth2Config *oauth2.Config // Commented for fix CLOUDTRUST-6415
	//tokenSource  oauth2.TokenSource // Commented for fix CLOUDTRUST-6415
}

const (
	// Max processing delay: let's assume that the user of OidcTokenProvider will have a maximum of 5 seconds to use the provided OIDC token
	maxProcessingDelay = int64(5)
)

// NewOidcTokenProvider creates an OidcTokenProvider
func NewOidcTokenProvider(config keycloak.Config, realm, username, password, clientID string, logger Logger) OidcTokenProvider {
	var perRealmTokenInfo = make(map[string]*oidcTokenInfo)
	config.URIProvider.ForEachContextURI(func(targetRealm, host, _ string) {
		perRealmTokenInfo[targetRealm] = &oidcTokenInfo{
			forwarded: fmt.Sprintf("host=%s;proto=https", host),
			/* Commented for fix CLOUDTRUST-6415
			oauth2Config: &oauth2.Config{
				ClientID: clientID,
				Endpoint: oauth2.Endpoint{TokenURL: fmt.Sprintf("%s/auth/realms/%s/protocol/openid-connect/token", config.AddrInternalAPI, realm)},
			},
			tokenSource: nil,
			*/
			// Added for fix CLOUDTRUST-6415
			url: fmt.Sprintf("%s/auth/realms/%s/protocol/openid-connect/token", config.AddrInternalAPI, realm),
		}
	})

	// If needed, can add &client_secret={secret}
	var body = fmt.Sprintf("grant_type=password&client_id=%s&username=%s&password=%s",
		url.QueryEscape(clientID), url.QueryEscape(username), url.QueryEscape(password))

	return &oidcTokenProvider{
		timeout:           config.Timeout,
		perRealmTokenInfo: perRealmTokenInfo,
		reqBody:           body, // Added for fix CLOUDTRUST-6415
		//username:          username, // Commented for fix CLOUDTRUST-6415
		//password:          password, // Commented for fix CLOUDTRUST-6415
		defaultKey: config.URIProvider.GetDefaultKey(),
		logger:     logger,
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
	// Added for fix CLOUDTRUST-6415
	if time.Now().Unix()+maxProcessingDelay < oti.validUntil {
		return oti.oidcToken.AccessToken, nil
	}

	var mimeType = "application/x-www-form-urlencoded"
	var httpClient = http.Client{
		Timeout: o.timeout,
	}
	//var req *http.Request
	var req, err = http.NewRequest("POST", oti.url, strings.NewReader(o.reqBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", mimeType)
	req.Header.Set("Forwarded", oti.forwarded)
	var resp *http.Response
	resp, err = httpClient.Do(req)
	if err != nil {
		o.logger.Warn(ctx, "msg", err.Error())
		return "", errorhandler.CreateInternalServerError("unexpected.httpResponse")
	}
	if resp.StatusCode == http.StatusUnauthorized {
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
	// First time we are requesting a token
	// Commented for fix CLOUDTRUST-6415
	/*if oti.tokenSource == nil {
		client := &http.Client{
			Transport: &customTransport{
				base:          http.DefaultTransport,
				forwardedHost: oti.forwarded,
			},
			Timeout: o.timeout,
		}
		var tokenCtx = context.WithValue(context.Background(), oauth2.HTTPClient, client)
		var token, err = oti.oauth2Config.PasswordCredentialsToken(tokenCtx, o.username, o.password)
		if err != nil {
			return "", err
		}
		oti.tokenSource = oti.oauth2Config.TokenSource(ctx, token)
		return token.AccessToken, nil
	}
	// A token has already been requested... Get it (automatically refresh it if needed)
	var token, err = oti.tokenSource.Token()
	if err != nil {
		return "", err
	}
	return token.AccessToken, nil*/
}
