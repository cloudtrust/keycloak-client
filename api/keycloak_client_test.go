package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const accessTokenValid = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyLCJpc3MiOiJodHRwczovL3NhbXBsZS5jb20vIn0.xLlV0CYqKDIPI-_IEABEcjRnKVNklivaw9WRmR8SXto"

func TestExtractIssuerFromToken(t *testing.T) {
	t.Run("Can't parse JWT", func(t *testing.T) {
		var _, err = extractIssuerFromToken("AAABBBCCC")
		assert.NotNil(t, err)
	})
	t.Run("Can't unmarshal token", func(t *testing.T) {
		var _, err = extractIssuerFromToken("AAA.BBB.CCC")
		assert.NotNil(t, err)
	})
	t.Run("Valid token", func(t *testing.T) {
		var issuer, err = extractIssuerFromToken(accessTokenValid)
		assert.Nil(t, err)
		assert.Equal(t, "https://sample.com/", issuer)
	})
}
