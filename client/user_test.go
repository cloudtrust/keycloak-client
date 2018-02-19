package client

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetUsers(t *testing.T) {
	var client = initTest(t)
	var users []UserRepresentation
	{
		var err error
		users, err = client.GetUsers("__internal")
		require.Nil(t, err, "could not get users")
	}
	for _, i := range users {
		fmt.Println(*i.Username)
	}
}

func TestCreateUser(t *testing.T) {
	var client = initTest(t)
	var realm = "__internal"
	var user = UserRepresentation{
		Username: str("johanr"),
	}
	var err = client.CreateUser(realm, user)
	assert.Nil(t, err)
}

func TestCountUsers(t *testing.T) {
	var client = initTest(t)
	var realm = "__internal"

	var count, err = client.CountUsers(realm)
	assert.Nil(t, err)
	assert.NotZero(t, count)
}
func TestGetUser(t *testing.T) {
	var client = initTest(t)
	var user UserRepresentation
	{
		var err error
		user, err = client.GetUser("__internal", "078f735b-ac07-4b39-88cb-88647c4ff47c")
		require.Nil(t, err, "could not get users")
	}
	fmt.Println(*user.Username)
}

func TestUpdateUser(t *testing.T) {
	var client = initTest(t)

	var user = UserRepresentation{
		Email: str("john.doe@elca.ch"),
	}
	var err = client.UpdateUser("__internal", "078f735b-ac07-4b39-88cb-88647c4ff47c", user)
	assert.Nil(t, err)
}
func TestDeleteUser(t *testing.T) {
	var client = initTest(t)

	var err = client.DeleteUser("__internal", "078f735b-ac07-4b39-88cb-88647c4ff47c")
	assert.Nil(t, err)
}
