package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetRealms(t *testing.T) {
	var client = initTest(t)
	var realms []RealmRepresentation
	{
		var err error
		realms, err = client.GetRealms()
		require.Nil(t, err, "could not get realms")
		assert.NotNil(t, realms)
	}
}

func TestCreateRealm(t *testing.T) {
	var client = initTest(t)
	var realm = RealmRepresentation{
		Realm: str("__internal"),
	}
	var err = client.CreateRealm(realm)
	assert.Nil(t, err)
}

func TestGetRealm(t *testing.T) {
	var client = initTest(t)

	var realm, err = client.GetRealm("__internal")
	assert.Nil(t, err)
	assert.NotNil(t, realm)
}

func TestUpdateRealm(t *testing.T) {
	var client = initTest(t)

	var realm = RealmRepresentation{
		DisplayName: str("Test realm"),
	}
	var err = client.UpdateRealm("__internal", realm)
	assert.Nil(t, err)
}
func TestDeleteRealm(t *testing.T) {
	var client = initTest(t)

	var err = client.DeleteRealm("__internal")
	assert.Nil(t, err)
}
