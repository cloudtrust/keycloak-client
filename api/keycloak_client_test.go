package api

import (
	"testing"

	"github.com/cloudtrust/keycloak-client/v2/toolbox"
	"github.com/stretchr/testify/assert"
)

func ptr(value string) *string {
	return &value
}

func TestForRealm(t *testing.T) {
	var kcConfig, err = toolbox.NewConfig(func(target any) error {
		var config = target.(*toolbox.InternalConfig)
		config.InternalURI = "http://cloudtrust:8080"
		config.DefaultKey = ptr("default")
		config.RealmPublicURI = map[string]string{
			"default": "https://my.domain.test",
			"other":   "https://my.other.domain.test",
		}
		return nil
	})
	assert.Nil(t, err)
	var c *Client
	c, err = New(kcConfig)
	assert.Nil(t, err)

	t.Run("Empty", func(t *testing.T) {
		assert.False(t, c.isIssuedByDefaultMaster(""))
	})
	t.Run("Invalid issuer", func(t *testing.T) {
		var jwtInvIssuer = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaXNzIjoxNTksImlhdCI6MTUxNjIzOTAyMn0.KMUFsIDTnFmyG3nMiGM6H9FNFUROf3wh7SmqJp-QV30"
		assert.False(t, c.isIssuedByDefaultMaster(jwtInvIssuer))
		assert.Equal(t, c.perRealmClients["other"], c.forRealm(jwtInvIssuer, "other"))
		assert.Equal(t, c.perRealmClients["default"], c.forRealm(jwtInvIssuer, "master"))
	})
	t.Run("Master token with other issuer domain", func(t *testing.T) {
		var jwtMaster = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaXNzIjoiaHR0cHM6Ly9teS5vdGhlci5kb21haW4udGVzdC9hdXRoL3JlYWxtcy9tYXN0ZXIiLCJhZG1pbiI6dHJ1ZSwiaWF0IjoxNTE2MjM5MDIyfQ.KMUFsIDTnFmyG3nMiGM6H9FNFUROf3wh7SmqJp-QV30"
		assert.False(t, c.isIssuedByDefaultMaster(jwtMaster))
		assert.Equal(t, c.perRealmClients["other"], c.forRealm(jwtMaster, "other"))
		assert.Equal(t, c.perRealmClients["default"], c.forRealm(jwtMaster, "master"))
	})
	t.Run("Master token with expected issuer domain", func(t *testing.T) {
		var jwtMaster = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaXNzIjoiaHR0cHM6Ly9teS5kb21haW4udGVzdC9hdXRoL3JlYWxtcy9tYXN0ZXIiLCJhZG1pbiI6dHJ1ZSwiaWF0IjoxNTE2MjM5MDIyfQ.KMUFsIDTnFmyG3nMiGM6H9FNFUROf3wh7SmqJp-QV30"
		assert.True(t, c.isIssuedByDefaultMaster(jwtMaster))
		assert.Equal(t, c.perRealmClients["default"], c.forRealm(jwtMaster, "other"))
		assert.Equal(t, c.perRealmClients["default"], c.forRealm(jwtMaster, "master"))
	})
	t.Run("Token from other realm", func(t *testing.T) {
		var jwtOther = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaXNzIjoiaHR0cHM6Ly9teS5kb21haW4udGVzdC9yZWFsbXMva2V5Y2xvYWsiLCJpYXQiOjE1MTYyMzkwMjJ9.KMUFsIDTnFmyG3nMiGM6H9FNFUROf3wh7SmqJp-QV30"
		assert.False(t, c.isIssuedByDefaultMaster(jwtOther))
		assert.Equal(t, c.perRealmClients["other"], c.forRealm(jwtOther, "other"))
		assert.Equal(t, c.perRealmClients["default"], c.forRealm(jwtOther, "master"))
	})
}
