package keycloak

import (
	"flag"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var (
	hostPort = flag.String("hostport", "10.244.18.2:80", "keycloak host:port")
	username = flag.String("username", "admin", "keycloak user name")
	password = flag.String("password", "admin", "keycloak password")
	to       = flag.Int("timeout", 20, "timeout in seconds")
)

func TestMain(m *testing.M) {
	flag.Parse()
	result := m.Run()
	os.Exit(result)
}

func initTest(t *testing.T) *Client {
	var config = Config{
		Addr:     fmt.Sprintf("http://%s", *hostPort),
		Username: *username,
		Password: *password,
		Timeout:  time.Duration(*to) * time.Second,
	}
	var client *Client
	{
		var err error
		client, err = New(config)
		require.Nil(t, err, "could not create client")
	}
	return client
}
