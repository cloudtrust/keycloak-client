package client

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func initTest(t *testing.T) *client {
	var config = HttpConfig{
		Addr:     "http://172.19.0.3:8080",
		Username: "admin",
		Password: "admin",
		Timeout:  time.Second * 20,
	}
	var client *client
	{
		var err error
		client, err = New(config)
		require.Nil(t, err, "could not create client")
	}
	return client
}

func TestGetToken(t *testing.T) {
	var client = initTest(t)
	var err = client.getToken()
	require.Nil(t, err, "could not get token")
	fmt.Println(client.accessToken)
}
