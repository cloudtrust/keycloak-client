package client

import (
	"testing"
	"time"
	"github.com/stretchr/testify/require"
	"fmt"
)

func initTest(t *testing.T) Client {
	var config HttpConfig = HttpConfig{
		Addr: "http://172.17.0.2:8080",
		Username: "admin",
		Password: "admin",
		Timeout: time.Second * 5,
	}
	var client Client
	{
		var err error
		client, err = NewHttpClient(config)
		require.Nil(t, err, "Failed to create client")
	}
	return client
}

func TestClient_getToken(t *testing.T) {
	var genClient Client = initTest(t)
	var httpClient *client = genClient.(*client)
	var err error = httpClient.getToken()
	require.Nil(t, err, "Failed to get token")
	fmt.Println(httpClient.accessToken)
}

func TestClient_GetRealms(t *testing.T) {
	var client Client = initTest(t)
	var realms []RealmRepresentation
	{
		var err error
		realms, err = client.GetRealms()
		require.Nil(t, err, "Failed to get realms")
	}
	fmt.Println(realms)
}

func TestClient_GetUsers(t *testing.T) {
	var client Client = initTest(t)
	var users []UserRepresentation
	{
		var err error
		users, err = client.GetUsers("master")
		require.Nil(t, err, "Failed to get users")
	}
	fmt.Println(users[0])
}

