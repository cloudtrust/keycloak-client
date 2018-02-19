package client

import (
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
		assert.NotZero(t, *i.Username)
	}
}

func TestCreateUser(t *testing.T) {
	var client = initTest(t)
	var realm = "__internal"
	var user = UserRepresentation{
		Username: str("john"),
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
		user, err = client.GetUser("__internal", "eb8b75ea-305d-40f6-87e5-ac8e16979c40")
		require.Nil(t, err, "could not get users")
		assert.NotZero(t, *user.Username)
	}
}

func TestUpdateUser(t *testing.T) {
	var client = initTest(t)

	var user = UserRepresentation{
		Email: str("john.doe@elca.ch"),
	}
	var err = client.UpdateUser("__internal", "eb8b75ea-305d-40f6-87e5-ac8e16979c40", user)
	assert.Nil(t, err)
}
func TestDeleteUser(t *testing.T) {
	var client = initTest(t)

	var err = client.DeleteUser("__internal", "eb8b75ea-305d-40f6-87e5-ac8e16979c40")
	assert.Nil(t, err)
}
