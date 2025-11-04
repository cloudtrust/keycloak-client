package toolbox

import (
	"context"
	"fmt"
	"net/url"
	"sync"
	"time"

	"github.com/cloudtrust/keycloak-client/v2"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/pkg/errors"
)

// OidcVerifierProvider is an interface for a provider of OidcVerifier instances
type OidcVerifierProvider interface {
	GetOidcVerifier(realm string) (OidcVerifier, error)
}

// OidcVerifier is an interface for OIDC token verifiers
type OidcVerifier interface {
	Verify(accessToken string) error
}

type verifierCache struct {
	internalURL    *url.URL
	externalURL    *url.URL
	verifiers      map[string]cachedVerifier
	verifiersMutex sync.RWMutex
}

type cachedVerifier struct {
	verifier  *oidc.IDTokenVerifier
	createdAt time.Time
	ctx       context.Context
}

// NewVerifierCache create an instance of OIDC verifier cache
func NewVerifierCache(internalURL *url.URL, externalURL *url.URL) OidcVerifierProvider {
	return &verifierCache{
		internalURL:    internalURL,
		externalURL:    externalURL,
		verifiers:      make(map[string]cachedVerifier),
		verifiersMutex: sync.RWMutex{},
	}
}

func (vc *verifierCache) GetOidcVerifier(realm string) (OidcVerifier, error) {
	vc.verifiersMutex.RLock()
	v, ok := vc.verifiers[realm]
	vc.verifiersMutex.RUnlock()
	if ok {
		return &v, nil
	}

	ctx := ContextWithForwarded(context.Background(), vc.internalURL, vc.externalURL)
	var oidcProvider *oidc.Provider
	{
		var err error
		var issuer = fmt.Sprintf("%s://%s/auth/realms/%s", vc.externalURL.Scheme, vc.externalURL.Host, realm)
		oidcProvider, err = oidc.NewProvider(ctx, issuer)
		if err != nil {
			return nil, errors.Wrap(err, keycloak.MsgErrCannotCreate+"."+keycloak.OIDCProvider)
		}
	}

	ov := oidcProvider.Verifier(&oidc.Config{SkipClientIDCheck: true})
	res := cachedVerifier{
		createdAt: time.Now(),
		verifier:  ov,
		ctx:       ctx,
	}
	vc.verifiersMutex.Lock()
	vc.verifiers[realm] = res
	vc.verifiersMutex.Unlock()

	return &res, nil
}

func (cv *cachedVerifier) Verify(accessToken string) error {
	_, err := cv.verifier.Verify(cv.ctx, accessToken)
	return err
}
