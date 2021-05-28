package toolbox

import (
	"testing"

	"github.com/cloudtrust/keycloak-client"
	"github.com/stretchr/testify/assert"
)

func TestGetProtocolAndDomain(t *testing.T) {
	var invalidURL = "not a valid URL"
	assert.Equal(t, invalidURL, getProtocolAndDomain(invalidURL))
	assert.Equal(t, "https://elca.ch", getProtocolAndDomain("https://ELCA.CH/PATH/TO/TARGET"))
}

func TestNewIssuerManager(t *testing.T) {
	t.Run("Empty config", func(t *testing.T) {
		_, err := NewIssuerManager(keycloak.Config{AddrTokenProvider: nil})
		assert.NotNil(t, err)
	})
	t.Run("Invalid URL", func(t *testing.T) {
		_, err := NewIssuerManager(keycloak.Config{AddrTokenProvider: []string{":"}})
		assert.NotNil(t, err)
	})

	defaultPath := "http://default.domain.com:5555"
	myDomainPath := "http://my.domain.com/path/to/somewhere"
	otherDomainPath := "http://other.domain.com:2120/"
	allDomains := []string{defaultPath, myDomainPath, otherDomainPath}

	prov, err := NewIssuerManager(keycloak.Config{AddrTokenProvider: allDomains})
	assert.Nil(t, err)
	assert.NotNil(t, prov)

	// No issuer provided with context
	issuerNoContext, _ := prov.GetOidcVerifierProvider("")
	// Unrecognized issuer provided in context
	issuerDefault, _ := prov.GetOidcVerifierProvider("http://unknown.issuer.com/one/path")
	// Case insensitive
	issuerMyDomain, _ := prov.GetOidcVerifierProvider("http://MY.DOMAIN.COM/issuer")
	// Other domain
	issuerOtherDomain, _ := prov.GetOidcVerifierProvider("http://other.domain.com:2120/any/thing/here")

	assert.Equal(t, issuerNoContext, issuerDefault)
	assert.NotEqual(t, issuerNoContext, issuerMyDomain)
	assert.NotEqual(t, issuerNoContext, issuerOtherDomain)
	assert.NotEqual(t, issuerMyDomain, issuerOtherDomain)
}
