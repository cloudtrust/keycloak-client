package keycloak

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	kcHost          = "localhost:8080"
	masterRealm     = "master"
	testRealm       = "identity-broker-realm"
	idpAlias        = "keycloak-oidc"
	mapperName      = "test-mapper"
	adminUsername   = "admin"
	adminPassword   = "admin"
	supportUsername = "support_user"
	supportPassword = "support"
)

func makeDefaultClient() (*Client, error) {
	timeout, _ := time.ParseDuration("10s")
	config := Config{
		AddrTokenProvider: kcHost,
		AddrAPI:           kcHost,
		Timeout:           timeout,
	}

	return New(config)
}

func TestGetIdp(t *testing.T) {
	client, err := makeDefaultClient()
	assert.Nil(t, err)

	//token, err := client.GetToken(masterRealm, adminUsername, adminPassword)
	token, err := client.GetToken(testRealm, supportUsername, supportPassword)
	if err != nil {
		t.Skip("Skipping test, no Keycloak configured to test properly.")
	}
	assert.Nil(t, err)

	idp, err := client.GetIdp(token, testRealm, idpAlias)
	assert.Nil(t, err)
	assert.Equal(t, idpAlias, *(idp.Alias))

}

func TestGetIdps(t *testing.T) {
	client, err := makeDefaultClient()
	assert.Nil(t, err)

	//token, err := client.GetToken(masterRealm, adminUsername, adminPassword)
	token, err := client.GetToken(testRealm, supportUsername, supportPassword)
	if err != nil {
		t.Skip("Skipping test, no Keycloak configured to test properly.")
	}
	assert.Nil(t, err)

	idps, err := client.GetIdps(token, testRealm)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(idps))
	assert.Equal(t, idpAlias, *(idps[0].Alias))
}

func TestGetIdpMappers(t *testing.T) {
	client, err := makeDefaultClient()
	assert.Nil(t, err)

	//token, err := client.GetToken(masterRealm, adminUsername, adminPassword)
	token, err := client.GetToken(testRealm, supportUsername, supportPassword)
	if err != nil {
		t.Skip("Skipping test, no Keycloak configured to test properly.")
	}
	assert.Nil(t, err)

	mappers, err := client.GetIdpMappers(token, testRealm, idpAlias)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(mappers))
	assert.Equal(t, mapperName, *(mappers[0].Name))
}
