package keycloak

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetUsers(t *testing.T) {
	var c = initTest(t)
	var users []UserRepresentation
	{
		var err error
		users, err = c.GetUsers("master")
		require.Nil(t, err, "could not get users")
	}
	for _, i := range users {
		fmt.Println(i.Credentials)
		assert.NotZero(t, *i.Username)
	}
}
