package toolbox

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	http_transport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func decodeRequest(_ context.Context, req *http.Request) (interface{}, error) {
	res := map[string]string{"realm": mux.Vars(req)["realm"], "host": req.Host}
	return res, nil
}

func encodeReply(_ context.Context, w http.ResponseWriter, rep interface{}) error {
	if rep == nil {
		w.WriteHeader(404)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)

	var json, err = json.Marshal(rep)
	if err == nil {
		w.Write(json)
	}
	return nil
}

func errorHandler(_ context.Context, _ error, w http.ResponseWriter) {
	w.WriteHeader(500)
}

func endpoint(_ context.Context, request interface{}) (response interface{}, err error) {
	var query = request.(map[string]string)
	if !strings.Contains(query["realm"], "realm") {
		return nil, nil
	}
	return map[string]string{
		"issuer":                 "http://" + query["host"] + "/auth/realms/" + query["realm"],
		"authorization_endpoint": "",
		"token_endpoint":         "",
		"jwks_uri":               "",
		"userinfo_endpoint":      "",
	}, nil
}

func TestGetOidcVerifier(t *testing.T) {
	verifierHandler := http_transport.NewServer(endpoint, decodeRequest, encodeReply, http_transport.ServerErrorEncoder(errorHandler))

	r := mux.NewRouter()
	r.Handle("/auth/realms/{realm}/.well-known/openid-configuration", verifierHandler)

	ts := httptest.NewServer(r)
	defer ts.Close()

	url, _ := url.Parse(ts.URL)

	{
		// First test with a verifier which hardly expires
		verifier := NewVerifierCache(url, time.Minute, 10*time.Minute)

		{
			// Unknown realm: can't get verifier
			_, err := verifier.GetOidcVerifier("unknown")
			assert.NotNil(t, err)
		}

		v1, e := verifier.GetOidcVerifier("realm1")
		assert.Nil(t, e)
		{
			// Ask for the same realm before its verifier expires
			v2, _ := verifier.GetOidcVerifier("realm1")
			assert.Equal(t, v1, v2)
		}
		{
			// Ask for a different verifier
			v3, _ := verifier.GetOidcVerifier("realm2")
			assert.NotEqual(t, v1, v3)
		}

		time.Sleep(100 * time.Millisecond)
		assert.NotNil(t, v1.Verify("abcdef"))
	}

	{
		// Now, test with a verifier which quickly expires on error
		verifier := NewVerifierCache(url, time.Minute, time.Millisecond)
		v1, _ := verifier.GetOidcVerifier("realm1")
		time.Sleep(100 * time.Millisecond)
		{
			// Ask for the same realm before its verifier expires
			v2, _ := verifier.GetOidcVerifier("realm1")
			assert.Equal(t, v1, v2)
		}
		{
			// Verify an invalid token
			assert.NotNil(t, v1.Verify("abcdef"))
			// Ask for the same realm before its verifier expires but after an error occured
			v2, _ := verifier.GetOidcVerifier("realm1")
			assert.NotEqual(t, v1, v2)
		}
	}
}
