package keycloak

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayground(t *testing.T) {
	var client = initTest(t)
	var clients, err = client.GetClients("master")

	for _, c := range clients {
		printStruct(c)
	}

	assert.Nil(t, err)
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

func printStruct(data interface{}) {
	var s, err = json.Marshal(data)
	if err != nil {
		fmt.Println("could not marshal json")
		return
	}
	fmt.Println(string(s))
	fmt.Println()
}
