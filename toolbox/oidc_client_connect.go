package toolbox

import (
	"context"
	"fmt"
	"strings"

	errorhandler "github.com/cloudtrust/common-service/v2/errors"
	"github.com/cloudtrust/keycloak-client/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

// OAuth2Config struct
type OAuth2Config struct {
	Realm        *string `mapstructure:"realm"`
	Username     *string `mapstructure:"username"`
	Password     *string `mapstructure:"password"`
	ClientID     *string `mapstructure:"client-id"`
	ClientSecret *string `mapstructure:"client-secret"`
}

// IsClientConfig checks if the config is a client config or a username/password one
func (oac *OAuth2Config) IsClientConfig() bool {
	return oac != nil && oac.Realm != nil && oac.ClientID != nil && oac.ClientSecret != nil
}

type oauth2TokenProvider struct {
	perRealmTokenInfo map[string]*oauth2TokenInfo
	defaultKey        string
	logger            Logger
}

type oauth2TokenInfo struct {
	tokenSource oauth2.TokenSource
	token       *oauth2.Token
}

// NewOAuth2TokenProvider creates an OidcTokenProvider
func NewOAuth2TokenProvider(kcConfig keycloak.Config, oauth2Config OAuth2Config, logger Logger) OidcTokenProvider {
	if !oauth2Config.IsClientConfig() {
		return NewOidcTokenProvider(kcConfig, *oauth2Config.Realm, *oauth2Config.Username, *oauth2Config.Password, *oauth2Config.ClientID, logger)
	}
	var perRealmTokenInfo = make(map[string]*oauth2TokenInfo)
	_ = ImportLegacyAddrTokenProvider(&kcConfig)
	kcConfig.URIProvider.ForEachTokenURI(func(targetRealm, tokenURI string) {
		var cfg = clientcredentials.Config{
			ClientID:     *oauth2Config.ClientID,
			ClientSecret: *oauth2Config.ClientSecret,
			TokenURL:     fmt.Sprintf(tokenURI, *oauth2Config.Realm),
		}
		perRealmTokenInfo[targetRealm] = &oauth2TokenInfo{
			tokenSource: cfg.TokenSource(context.Background()),
		}
	})

	return &oauth2TokenProvider{
		perRealmTokenInfo: perRealmTokenInfo,
		defaultKey:        kcConfig.URIProvider.GetDefaultKey(),
		logger:            logger,
	}
}

func (o *oauth2TokenProvider) ProvideToken(ctx context.Context) (string, error) {
	return o.ProvideTokenForRealm(ctx, o.defaultKey)
}

func (o *oauth2TokenProvider) ProvideTokenForRealm(ctx context.Context, realm string) (string, error) {
	var oti *oauth2TokenInfo
	var ok bool
	if oti, ok = o.perRealmTokenInfo[strings.ToLower(realm)]; !ok {
		if realm == o.defaultKey {
			return "", errorhandler.CreateInternalServerError("unknownRealm")
		}
		return o.ProvideTokenForRealm(ctx, o.defaultKey)
	}
	if oti.token == nil || !oti.token.Valid() {
		var err error
		oti.token, err = oti.tokenSource.Token()
		if err != nil {
			return "", err
		}
	}
	return oti.token.AccessToken, nil
}
