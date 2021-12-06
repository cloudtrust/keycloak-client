package toolbox

import (
	"context"
	"fmt"
	"net/url"
	"sync"
	"time"

	"github.com/cloudtrust/keycloak-client/v2"
	oidc "github.com/coreos/go-oidc"
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
	tokenURL       *url.URL
	verifiers      map[string]cachedVerifier
	verifiersMutex sync.RWMutex
}

type cachedVerifier struct {
	verifier  *oidc.IDTokenVerifier
	createdAt time.Time
}

// NewVerifierCache create an instance of OIDC verifier cache
func NewVerifierCache(tokenURL *url.URL) OidcVerifierProvider {
	return &verifierCache{
		tokenURL:       tokenURL,
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
	var oidcProvider *oidc.Provider
	{
		var err error
		var issuer = fmt.Sprintf("%s/auth/realms/%s", vc.tokenURL.String(), realm)
		oidcProvider, err = oidc.NewProvider(context.Background(), issuer)
		if err != nil {
			return nil, errors.Wrap(err, keycloak.MsgErrCannotCreate+"."+keycloak.OIDCProvider)
		}
	}

	ov := oidcProvider.Verifier(&oidc.Config{SkipClientIDCheck: true})
	res := cachedVerifier{
		createdAt: time.Now(),
		verifier:  ov,
	}
	vc.verifiersMutex.Lock()
	vc.verifiers[realm] = res
	vc.verifiersMutex.Unlock()

	return &res, nil
}

func (cv *cachedVerifier) Verify(accessToken string) error {
	_, err := cv.verifier.Verify(context.Background(), accessToken)
	return err
}
