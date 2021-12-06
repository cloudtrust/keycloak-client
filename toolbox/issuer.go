package toolbox

import (
	"errors"
	"net/url"
	"regexp"
	"strings"

	"github.com/cloudtrust/keycloak-client"
)

// IssuerManager provides URL according to a given context
type IssuerManager interface {
	GetOidcVerifierProvider(issuer string) (OidcVerifierProvider, error)
}

type issuerManager struct {
	domainToVerifier map[string]OidcVerifierProvider
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
func NewIssuerManager(config keycloak.Config) (IssuerManager, error) {
	if err := ImportLegacyAddrTokenProvider(&config); err != nil {
		return nil, err
	}

	var domainToVerifier = make(map[string]OidcVerifierProvider)
	for _, value := range config.URIProvider.GetAllBaseURIs() {
		uToken, err := url.Parse(value)
		if err != nil {
			return nil, err
		}
		verifier := NewVerifierCache(uToken)
		domainToVerifier[getProtocolAndDomain(value)] = verifier
	}
	return &issuerManager{
		domainToVerifier: domainToVerifier,
	}, nil
}

func (im *issuerManager) GetOidcVerifierProvider(issuer string) (OidcVerifierProvider, error) {
	issuerDomain := getProtocolAndDomain(issuer)
	if verifier, ok := im.domainToVerifier[issuerDomain]; ok {
		return verifier, nil
	}
	return nil, errors.New("unknownIssuer")
}
