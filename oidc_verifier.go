package keycloak

import (
	"context"
	"fmt"
	"net/url"
	"time"

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
	duration       time.Duration
	errorTolerance time.Duration
	tokenURL       *url.URL
	verifiers      map[string]cachedVerifier
}

type cachedVerifier struct {
	verifier            *oidc.IDTokenVerifier
	createdAt           time.Time
	expireAt            time.Time
	invalidateOnErrorAt time.Time
}

// NewVerifierCache create an instance of OIDC verifier cache
func NewVerifierCache(tokenURL *url.URL, timeToLive time.Duration, errorTolerance time.Duration) OidcVerifierProvider {
	return &verifierCache{
		duration:       timeToLive,
		errorTolerance: errorTolerance,
		tokenURL:       tokenURL,
		verifiers:      make(map[string]cachedVerifier),
	}
}

func (vc *verifierCache) GetOidcVerifier(realm string) (OidcVerifier, error) {
	v, ok := vc.verifiers[realm]
	if ok && v.isValid() {
		return &v, nil
	}
	var oidcProvider *oidc.Provider
	{
		var err error
		var issuer = fmt.Sprintf("%s/auth/realms/%s", vc.tokenURL.String(), realm)
		oidcProvider, err = oidc.NewProvider(context.Background(), issuer)
		if err != nil {
			return nil, errors.Wrap(err, MsgErrCannotCreate+"."+OIDCProvider)
		}
	}

	ov := oidcProvider.Verifier(&oidc.Config{SkipClientIDCheck: true})
	res := cachedVerifier{
		createdAt:           time.Now(),
		expireAt:            time.Now().Add(vc.duration),
		invalidateOnErrorAt: time.Now().Add(vc.errorTolerance),
		verifier:            ov,
	}
	vc.verifiers[realm] = res

	return &res, nil
}

func (cv *cachedVerifier) isValid() bool {
	return time.Now().Before(cv.expireAt)
}

func (cv *cachedVerifier) Verify(accessToken string) error {
	_, err := cv.verifier.Verify(context.Background(), accessToken)
	if err != nil && time.Now().After(cv.invalidateOnErrorAt) {
		// An error occured and current time is after invalidateOnErrorAt
		// Let's make this verifier expire
		cv.expireAt = cv.createdAt
	}
	return err
}
