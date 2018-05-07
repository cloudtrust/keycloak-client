package keycloak

import (
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
)

func TestExportRealms(t *testing.T) {

	var client, err = New(Config{
		Addr:     "http://keycloak:80",
		Username: "admin",
		Password: "admin",
		Timeout:  10 * time.Second,
	})
	require.Nil(t, err, "Err wasnt nil!", err)
	{
		var realm, err = client.ExportRealm("master")
		require.Nil(t, err, "Err isnt nil!", err)
		spew.Dump(realm)
	}
}
