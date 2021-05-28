package toolbox

import (
	"testing"

	"github.com/cloudtrust/keycloak-client"
	"github.com/stretchr/testify/assert"
)

func TestNewKeycloakURIProvider(t *testing.T) {
	t.Run("FromArray-Empty slice", func(t *testing.T) {
		var _, err = NewKeycloakURIProviderFromArray(nil)
		assert.NotNil(t, err)
	})
	t.Run("FromArray-Two items slice", func(t *testing.T) {
		var uriProvider, err = NewKeycloakURIProviderFromArray([]string{"http://localhost:8080", "http://127.0.0.1:8080"})
		assert.Nil(t, err)
		assert.NotNil(t, uriProvider)
		assert.Len(t, uriProvider.(*kcURIProvider).entries, 2)
	})
	t.Run("FromMap-Default key not exists", func(t *testing.T) {
		var _, err = NewKeycloakURIProvider(map[string]string{"one": "http://localhost:8080", "two": "http://127.0.0.1:8080"}, "other-key")
		assert.NotNil(t, err)
	})
	t.Run("FromMap-Default key exists", func(t *testing.T) {
		var one = "http://localhost:8080"
		var two = "http://127.0.0.1:8080"
		var uriProvider, err = NewKeycloakURIProvider(map[string]string{"one": one, "two": two}, "two")
		assert.Nil(t, err)

		assert.Equal(t, one, uriProvider.GetBaseURI("one"))
		assert.Equal(t, two, uriProvider.GetBaseURI("two"))
		assert.Equal(t, two, uriProvider.GetBaseURI("other"))

		var allBaseURIs = uriProvider.GetAllBaseURIs()
		assert.Equal(t, two, allBaseURIs[0])
		assert.Equal(t, one, allBaseURIs[1])
	})
}

func TestImportLegacyAddrTokenProvider(t *testing.T) {
	var cfg = keycloak.Config{AddrTokenProvider: []string{"http://localhost:8080", "http://127.0.0.1:8080"}}
	assert.Nil(t, cfg.URIProvider)

	var err = ImportLegacyAddrTokenProvider(&cfg)
	assert.Nil(t, err)
	assert.NotNil(t, cfg.URIProvider)

	var uriProvider = cfg.URIProvider.(*kcURIProvider)
	assert.Equal(t, "default", uriProvider.defaultKey)
	assert.Equal(t, cfg.AddrTokenProvider[0], uriProvider.entries["default"])
	assert.Equal(t, cfg.AddrTokenProvider[1], uriProvider.entries["entry-1"])
}
