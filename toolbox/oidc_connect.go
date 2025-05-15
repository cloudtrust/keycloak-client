package toolbox

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	errorhandler "github.com/cloudtrust/common-service/v2/errors"
	"github.com/cloudtrust/keycloak-client/v2"
	"golang.org/x/oauth2"
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
	username          string
	password          string
	defaultKey        string
	logger            Logger
}

type oidcTokenInfo struct {
	forwarded    string
	oauth2Config *oauth2.Config
	tokenSource  oauth2.TokenSource
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
			oauth2Config: &oauth2.Config{
				ClientID: clientID,
				Endpoint: oauth2.Endpoint{TokenURL: fmt.Sprintf("%s/auth/realms/%s/protocol/openid-connect/token", config.AddrInternalAPI, realm)},
			},
			tokenSource: nil,
		}
	})

	return &oidcTokenProvider{
		timeout:           config.Timeout,
		perRealmTokenInfo: perRealmTokenInfo,
		username:          username,
		password:          password,
		defaultKey:        config.URIProvider.GetDefaultKey(),
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
	// First time we are requesting a token
	if oti.tokenSource == nil {
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
	return token.AccessToken, nil
}
