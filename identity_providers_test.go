package keycloak

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	kcHost     = "localhost:8080"
	master     = "master"
	realm      = "identity-broker-realm"
	idpAlias   = "keycloak-oidc"
	mapperName = "test-mapper"
	username   = "admin"
	password   = "admin"
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

func TestGetTheIdp(t *testing.T) {
	client, err := makeDefaultClient()
	assert.Nil(t, err)

	token, err := client.GetToken(master, username, password)
	fmt.Printf(token)
	assert.Nil(t, err)

	idp, err := client.GetTheIdp(token, realm, idpAlias)
	assert.Nil(t, err)
	assert.Equal(t, idpAlias, *(idp.Alias))

}

func TestGetIdps(t *testing.T) {
	client, err := makeDefaultClient()
	assert.Nil(t, err)

	token, err := client.GetToken(master, username, password)
	fmt.Printf(token)
	assert.Nil(t, err)

	idps, err := client.GetIdps(token, realm)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(idps))
	assert.Equal(t, idpAlias, *(idps[0].Alias))
}

func TestGetIdpMappers(t *testing.T) {
	client, err := makeDefaultClient()
	assert.Nil(t, err)

	token, err := client.GetToken(master, username, password)
	fmt.Printf(token)
	assert.Nil(t, err)

	mappers, err := client.GetIdpMappers(token, realm, idpAlias)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(mappers))
	assert.Equal(t, mapperName, *(mappers[0].Name))
}
