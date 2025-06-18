package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsIssuedByMaster(t *testing.T) {
	var c = Client{}

	t.Run("Empty", func(t *testing.T) {
		assert.False(t, c.isIssuedByDefaultMaster(""))
	})
	t.Run("Invalid issuer", func(t *testing.T) {
		var jwtInvIssuer = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaXNzIjoxNTksImlhdCI6MTUxNjIzOTAyMn0.KMUFsIDTnFmyG3nMiGM6H9FNFUROf3wh7SmqJp-QV30"
		assert.False(t, c.isIssuedByDefaultMaster(jwtInvIssuer))
	})
	t.Run("Master token with unexpected issuer domain", func(t *testing.T) {
		var jwtMaster = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaXNzIjoiaHR0cHM6Ly9teS5kb21haW4udGVzdC9hdXRoL3JlYWxtcy9tYXN0ZXIiLCJhZG1pbiI6dHJ1ZSwiaWF0IjoxNTE2MjM5MDIyfQ.KMUFsIDTnFmyG3nMiGM6H9FNFUROf3wh7SmqJp-QV30"
		c.baseURL = "https://my.unexpected.test"
		assert.False(t, c.isIssuedByDefaultMaster(jwtMaster))
	})
	t.Run("Master token with expected issuer domain", func(t *testing.T) {
		var jwtMaster = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaXNzIjoiaHR0cHM6Ly9teS5kb21haW4udGVzdC9hdXRoL3JlYWxtcy9tYXN0ZXIiLCJhZG1pbiI6dHJ1ZSwiaWF0IjoxNTE2MjM5MDIyfQ.KMUFsIDTnFmyG3nMiGM6H9FNFUROf3wh7SmqJp-QV30"
		c.baseURL = "https://my.domain.test"
		assert.True(t, c.isIssuedByDefaultMaster(jwtMaster))
	})
	t.Run("Token from other realm", func(t *testing.T) {
		var jwtOther = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaXNzIjoiaHR0cHM6Ly9teS5kb21haW4udGVzdC9yZWFsbXMva2V5Y2xvYWsiLCJpYXQiOjE1MTYyMzkwMjJ9.KMUFsIDTnFmyG3nMiGM6H9FNFUROf3wh7SmqJp-QV30"
		assert.False(t, c.isIssuedByDefaultMaster(jwtOther))
	})
}
