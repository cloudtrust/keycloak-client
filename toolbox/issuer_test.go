package toolbox

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudtrust/keycloak-client"
	"github.com/stretchr/testify/assert"
)

type contextKey int

const (
	keyContextIssuerDomain contextKey = iota
)

func TestGetProtocolAndDomain(t *testing.T) {
	var invalidURL = "not a valid URL"
	assert.Equal(t, invalidURL, getProtocolAndDomain(invalidURL))
	assert.Equal(t, "https://elca.ch", getProtocolAndDomain("https://ELCA.CH/PATH/TO/TARGET"))
}

func TestNewIssuerManager(t *testing.T) {
	t.Run("Invalid URL", func(t *testing.T) {
		_, err := NewIssuerManager(keycloak.Config{AddrTokenProvider: ":"}, keyContextIssuerDomain)
		assert.NotNil(t, err)
	})

	defaultPath := "http://default.domain.com:5555"
	myDomainPath := "http://my.domain.com/path/to/somewhere"
	otherDomainPath := "http://other.domain.com:2120/"
	allDomains := fmt.Sprintf("%s %s %s", defaultPath, myDomainPath, otherDomainPath)

	prov, err := NewIssuerManager(keycloak.Config{AddrTokenProvider: allDomains}, keyContextIssuerDomain)
	assert.Nil(t, err)
	assert.NotNil(t, prov)

	// No issuer provided with context
	issuerNoContext, _ := prov.GetIssuer(context.Background())
	// Unrecognized issuer provided in context
	issuerDefault, _ := prov.GetIssuer(context.WithValue(context.Background(), keyContextIssuerDomain, "http://unknown.issuer.com/one/path"))
	// Case insensitive
	issuerMyDomain, _ := prov.GetIssuer(context.WithValue(context.Background(), keyContextIssuerDomain, "http://MY.DOMAIN.COM/issuer"))
	// Other domain
	issuerOtherDomain, _ := prov.GetIssuer(context.WithValue(context.Background(), keyContextIssuerDomain, "http://other.domain.com:2120/any/thing/here"))

	assert.Equal(t, issuerNoContext, issuerDefault)
	assert.NotEqual(t, issuerNoContext, issuerMyDomain)
	assert.NotEqual(t, issuerNoContext, issuerOtherDomain)
	assert.NotEqual(t, issuerMyDomain, issuerOtherDomain)
}
