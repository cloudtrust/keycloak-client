package client

import (
	"testing"
	"time"
	"github.com/stretchr/testify/require"
	"fmt"
	"crypto/cipher"
	"crypto/aes"
)

type setivable interface {
	cipher.BlockMode
	SetIV([]byte)
}


func Test_Poop(t *testing.T)  {
	var k []byte = make([]byte, 16)
	var iv []byte = k
	b, e := aes.NewCipher(k)
	if e != nil {
		fmt.Println(e)
		return
	}
	var c = cipher.NewCBCEncrypter(b, iv)
	switch cp := c.(type) {
	case setivable:
		fmt.Println("Haha!")
		cp.SetIV(k)
	case cipher.BlockMode:
		fmt.Println("Hoho!")
	}
}

func initTest(t *testing.T) Client {
	var config HttpConfig = HttpConfig{
		Addr: "http://127.0.0.1:8080",
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
	var realms []map[string]interface{}
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

