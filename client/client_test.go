package client

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func initTest(t *testing.T) *Client {
	var config = HttpConfig{
		Addr:     "http://172.19.0.3:8080",
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

func TestGetToken(t *testing.T) {
	var client = initTest(t)
	var err = client.getToken()
	require.Nil(t, err, "could not get token")
	assert.NotZero(t, client.accessToken)
}
