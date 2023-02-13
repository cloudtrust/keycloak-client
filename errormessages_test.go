package keycloak

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTTPError(t *testing.T) {
	var err = HTTPError{HTTPStatus: 400, Message: "error message"}
	assert.Equal(t, "400:error message", err.Error())
}

func TestClientDetailedError(t *testing.T) {
	var err = ClientDetailedError{HTTPStatus: 400, Message: "error message"}
	assert.Equal(t, "400:error message", err.Error())
	assert.Equal(t, 400, err.Status())
	assert.Equal(t, "error message", err.ErrorMessage())
}
