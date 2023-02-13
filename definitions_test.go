package keycloak

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsUserProfileEnabled(t *testing.T) {
	var realm = RealmRepresentation{Attributes: nil}
	var bFalse = "FALSE"
	var bTrue = "TRue"

	t.Run("No attributes", func(t *testing.T) {
		assert.False(t, realm.IsUserProfileEnabled())
	})
	t.Run("Empty attributes", func(t *testing.T) {
		realm.Attributes = &map[string]*string{}
		assert.False(t, realm.IsUserProfileEnabled())
	})
	t.Run("User profile attribute is false", func(t *testing.T) {
		(*realm.Attributes)["userProfileEnabled"] = &bFalse
		assert.False(t, realm.IsUserProfileEnabled())
	})
	t.Run("User profile attribute is true", func(t *testing.T) {
		(*realm.Attributes)["userProfileEnabled"] = &bTrue
		assert.True(t, realm.IsUserProfileEnabled())
	})
}
