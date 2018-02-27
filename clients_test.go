package keycloak

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func initTest(t *testing.T) *Client {
	var config = Config{
		Addr:     "http://127.0.0.1",
		Username: "admin",
		Password: "admin",
		Timeout:  time.Second * 20,
	}
	var client *Client
	{
		var err error
		client, err = New(config)
		require.Nil(t, err, "could not create client")
	}
	return client
}

func TestCreateRealm(t *testing.T) {
	var client = initTest(t)
	var clients, err = client.GetClients("master")
	for i, c := range clients {
		fmt.Println(i, *(c.Id), *c.ClientId)
	}
	assert.Nil(t, err)
}

func TestGetClient(t *testing.T) {
	var client = initTest(t)
	var c, err = client.GetClient("318ab6db-c056-4d2f-b4f6-c0b585ee45b3", "master")
	fmt.Println(*(c.Id), *c.ClientId, c.Secret)
	assert.Nil(t, err)
}

func TestGetSecret(t *testing.T) {
	var client = initTest(t)
	var c, err = client.GetSecret("318ab6db-c056-4d2f-b4f6-c0b585ee45b3", "master")
	fmt.Println(*(c.Value))
	assert.Nil(t, err)
}
