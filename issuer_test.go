package keycloak

import (
	"context"
	"fmt"
	"testing"

	cs "github.com/cloudtrust/common-service"
	"github.com/stretchr/testify/assert"
)

func TestGetProtocolAndDomain(t *testing.T) {
	var invalidURL = "not a valid URL"
	assert.Equal(t, invalidURL, getProtocolAndDomain(invalidURL))
	assert.Equal(t, "https://elca.ch", getProtocolAndDomain("https://ELCA.CH/PATH/TO/TARGET"))
}

func TestNewIssuerManager(t *testing.T) {
	{
		_, err := NewIssuerManager(Config{AddrTokenProvider: ":"})
		assert.NotNil(t, err)
	}

	defaultPath := "http://default.domain.com:5555"
	myDomainPath := "http://my.domain.com/path/to/somewhere"
	otherDomainPath := "http://other.domain.com:2120/"
	allDomains := fmt.Sprintf("%s %s %s", defaultPath, myDomainPath, otherDomainPath)

	prov, err := NewIssuerManager(Config{AddrTokenProvider: allDomains})
	assert.Nil(t, err)
	assert.NotNil(t, prov)

	// No issuer provided with context
	issuerNoContext, _ := prov.GetIssuer(context.Background())
	// Unrecognized issuer provided in context
	issuerDefault, _ := prov.GetIssuer(context.WithValue(context.Background(), cs.CtContextIssuerDomain, "http://unknown.issuer.com/one/path"))
	// Case insensitive
	issuerMyDomain, _ := prov.GetIssuer(context.WithValue(context.Background(), cs.CtContextIssuerDomain, "http://MY.DOMAIN.COM/issuer"))
	// Other domain
	issuerOtherDomain, _ := prov.GetIssuer(context.WithValue(context.Background(), cs.CtContextIssuerDomain, "http://other.domain.com:2120/any/thing/here"))

	assert.Equal(t, issuerNoContext, issuerDefault)
	assert.NotEqual(t, issuerNoContext, issuerMyDomain)
	assert.NotEqual(t, issuerNoContext, issuerOtherDomain)
	assert.NotEqual(t, issuerMyDomain, issuerOtherDomain)
}
