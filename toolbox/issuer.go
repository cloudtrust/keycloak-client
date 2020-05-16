package toolbox

import (
	"context"
	"errors"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/cloudtrust/keycloak-client"
)

// IssuerManager provides URL according to a given context
type IssuerManager interface {
	GetIssuer(ctx context.Context) (OidcVerifierProvider, error)
}

type issuerManager struct {
	domainToIssuer         map[string]OidcVerifierProvider
	keyContextIssuerDomain interface{}
}

func getProtocolAndDomain(URL string) string {
	var r = regexp.MustCompile(`^\w+:\/\/[^\/]+`)
	var match = r.FindStringSubmatch(URL)
	if match != nil {
		return strings.ToLower(match[0])
	}
	// Best effort: if not found return the whole input string
	return URL
}

// NewIssuerManager creates a new URLProvider
func NewIssuerManager(config keycloak.Config, keyContextIssuerDomain interface{}) (IssuerManager, error) {
	URLs := config.AddrTokenProvider
	// Use default values when clients are not initializing these values
	cacheTTL := config.CacheTTL
	if cacheTTL == 0 {
		cacheTTL = 15 * time.Minute
	}
	errTolerance := config.ErrorTolerance
	if errTolerance == 0 {
		errTolerance = time.Minute
	}

	var domainToIssuer = make(map[string]OidcVerifierProvider)

	for _, value := range strings.Split(URLs, " ") {
		uToken, err := url.Parse(value)
		if err != nil {
			return nil, err
		}
		issuer := NewVerifierCache(uToken, cacheTTL, errTolerance)
		domainToIssuer[getProtocolAndDomain(value)] = issuer
	}
	return &issuerManager{
		domainToIssuer:         domainToIssuer,
		keyContextIssuerDomain: keyContextIssuerDomain,
	}, nil
}

func (im *issuerManager) GetIssuer(ctx context.Context) (OidcVerifierProvider, error) {
	if rawValue := ctx.Value(im.keyContextIssuerDomain); rawValue != nil {
		// The issuer domain has been found in the context
		issuerDomain := getProtocolAndDomain(rawValue.(string))
		if issuer, ok := im.domainToIssuer[issuerDomain]; ok {
			return issuer, nil
		}
	}
	return nil, errors.New("Unknown issuer")
}
