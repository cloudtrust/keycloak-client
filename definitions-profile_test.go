package keycloak

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAttribute(t *testing.T) {
	var attrb = ProfileAttrbRepresentation{}
	var key = "key"

	t.Run("Annotations is nil", func(t *testing.T) {
		assert.False(t, attrb.IsAnnotationFalse(key))
		assert.False(t, attrb.IsAnnotationTrue(key))
	})
	t.Run("Annotations is empty", func(t *testing.T) {
		attrb.Annotations = map[string]string{}
		assert.False(t, attrb.IsAnnotationFalse(key))
		assert.False(t, attrb.IsAnnotationTrue(key))
	})
	t.Run("Annotations is empty", func(t *testing.T) {
		attrb.Annotations[key] = "false"
		assert.True(t, attrb.IsAnnotationFalse(key))
		assert.False(t, attrb.IsAnnotationTrue(key))
	})
}
