package toolbox

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
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

var lastForwarded string

func toBase64(value []byte) string {
	return strings.ReplaceAll(base64.StdEncoding.EncodeToString(value), "=", "")
}

func toJSONBase64(value any) string {
	res, _ := json.Marshal(value)
	return toBase64(res)
}

func generateJWT(issuer string) string {
	var header = toJSONBase64(map[string]any{"alg": "RS256", "typ": "JWT"})
	var payload = toJSONBase64(map[string]any{"sub": "1234567890", "name": "John Doe", "iat": time.Now().Unix(), "iss": issuer})
	var sign = toBase64([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})
	return fmt.Sprintf("%s.%s.%s", header, payload, sign)
}

func createCerts() any {
	return map[string]any{
		"keys": []map[string]any{
			{
				"kid": "kid1",
				"kty": "RSA",
				"alg": "RS256",
				"use": "sig",
				"x5c": []string{
					"MIICnTCCAYUCBgFvrUPBDTANBgkqhkiG9w0BAQsFADASMRAwDgYDVQQDDAd0cnVzdGlkMB4XDTIwMDExNjA3Mjk1NloXDTMwMDExNjA3MzEzNlowEjEQMA4GA1UEAwwHdHJ1c3RpZDCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBALZnC3z8Z5YoJRDwii4fILVnHJluqbevvLQEuyCMfgGDPwLZtDy6X2ksg0/IMRxf+2k1qkiTlQodEdOm0Ypl9AB78wUj9Anh0mkWOSv2Xv2Qqq4TXhygyPcHi0DaP+slCjeyBNdXWm4CEVV81ylWy1wqPM8JTBCkOirGsC8IaVguov+41p/uj8oCrg7X0hdlVNfIGABiuYagoG6JfAL/jL2pykY8UOb3BPtTHGeS9lZCdOf6sPIonhc9BynU2cigfjoimhFgbJqBXMlnl+OB8YyT/cp/+Jrms3Tm6UjhdoC50EKTIK0Wis6jFf+S9dwlhGlgaa9poNI2SvY3W9HTqSMCAwEAATANBgkqhkiG9w0BAQsFAAOCAQEAYnM8axU9w9lQtt9lkEitTn9yhy9cCjTWM0utllutq1y7rVyLhuu+P//6SXUv1fuZQKAxFr3rteX55dBPvwCoT+lKl0SoLXgZv9DD1WUfiRApYNXfBsw3mluaYFwZceIaTBhHu0b6blTpl9wJndZx69TcLFsMKcbP/CifwFfutG6S0SToQHSoLi6rGKEVFKL6UIKFF35k0AG7Qg9ZwAc61GxKoCXm8U0JUdMc8Sq6yiFuqHQiKbuRW7KQVD36/whVInVEPRPLdQIu2A4TT5+dG/Vwz+2egS4ItsjKGN5B74ljK3jvH817HLxmBUdnWXM86LbHj7C6U/sQK56TSIUfyw==",
				},
				"x5t":      "k98EZetZbspu32ZCWZkgLIkxAN0",
				"x5t#S256": "jI9wl5oGxT9t5Zp1drnWO99Q4M57K6iKS-XKrX9gMIg",
				"n":        "tmcLfPxnliglEPCKLh8gtWccmW6pt6-8tAS7IIx-AYM_Atm0PLpfaSyDT8gxHF_7aTWqSJOVCh0R06bRimX0AHvzBSP0CeHSaRY5K_Ze_ZCqrhNeHKDI9weLQNo_6yUKN7IE11dabgIRVXzXKVbLXCo8zwlMEKQ6KsawLwhpWC6i_7jWn-6PygKuDtfSF2VU18gYAGK5hqCgbol8Av-MvanKRjxQ5vcE-1McZ5L2VkJ05_qw8iieFz0HKdTZyKB-OiKaEWBsmoFcyWeX44HxjJP9yn_4muazdObpSOF2gLnQQpMgrRaKzqMV_5L13CWEaWBpr2mg0jZK9jdb0dOpIw",
				"e":        "AQAB",
			},
		},
	}
}

func createProtocolOIDC(baseURL string, realm string) any {
	var issuer = baseURL + "/auth/realms/" + realm
	var protocolOIDC = issuer + "/protocol/openid-connect"
	var res = map[string]any{
		"issuer":                 issuer,
		"authorization_endpoint": protocolOIDC + "/auth",
		"token_endpoint":         protocolOIDC + "/token",
		"introspection_endpoint": protocolOIDC + "/introspect",
		"userinfo_endpoint":      protocolOIDC + "/userinfo",
		"end_session_endpoint":   protocolOIDC + "/logout",
		"jwks_uri":               protocolOIDC + "/certs",
	}
	return res
}

func decodeRequest(_ context.Context, req *http.Request) (any, error) {
	res := map[string]string{"realm": mux.Vars(req)["realm"], "host": req.Host, "scheme": req.URL.Scheme, "uri": req.RequestURI}
	if forwarded := req.Header.Get("Forwarded"); forwarded != "" {
		lastForwarded = forwarded
	} else {
		lastForwarded = ""
	}
	// Following values are not necessary for the test but can be useful for debugging
	for k, v := range req.Header {
		res[k] = v[0]
	}
	return res, nil
}

func encodeReply(_ context.Context, w http.ResponseWriter, rep any) error {
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

func endpoint(_ context.Context, request any) (response any, err error) {
	var query = request.(map[string]string)
	if !strings.Contains(query["realm"], "realm") {
		return nil, errors.New("invalid endpoint: " + query["uri"])
	}
	if strings.HasSuffix(query["uri"], "/certs") {
		return createCerts(), nil
	}
	var host string
	if forwarded, ok := query["Forwarded"]; ok {
		var tmpHost = strings.Split(strings.Split(forwarded, "=")[1], ";")[0]
		var tmpProto = strings.Split(forwarded, "=")[2]
		host = fmt.Sprintf("%s://%s", tmpProto, tmpHost)
	} else {
		host = fmt.Sprintf("%s://%s", query["scheme"], query["host"])
	}
	return createProtocolOIDC(host, query["realm"]), nil
}

func TestGetOidcVerifier(t *testing.T) {
	verifierHandler := http_transport.NewServer(endpoint, decodeRequest, encodeReply, http_transport.ServerErrorEncoder(errorHandler))

	r := mux.NewRouter()
	r.Handle("/auth/realms/{realm}/.well-known/openid-configuration", verifierHandler)
	r.Handle("/auth/realms/{realm}/protocol/openid-connect/certs", verifierHandler)

	ts := httptest.NewServer(r)
	defer ts.Close()

	internalURL, _ := url.Parse(ts.URL)
	defaultExternalURL, _ := url.Parse("https://idp-dev.trustid.ch")
	whiteLabelURL, _ := url.Parse("https://white.label.url/auth/realms/myrealm")

	{
		// First test with a verifier which hardly expires
		defaultVerifier := NewVerifierCache(internalURL, defaultExternalURL)
		whiteLabelVerifier := NewVerifierCache(internalURL, whiteLabelURL)

		jwtDefault := generateJWT("https://idp-dev.trustid.ch/auth/realms/trustid")
		jwtWhiteLabel := generateJWT("https://white.label.url/auth/realms/trustid")

		t.Run("Unknown realm: can't get verifier", func(t *testing.T) {
			_, err := defaultVerifier.GetOidcVerifier("unknown")
			assert.NotNil(t, err)
			assert.Contains(t, lastForwarded, "idp-dev")
		})

		var v1 OidcVerifier
		t.Run("Ask for a verifier for realm1", func(t *testing.T) {
			var e error
			v1, e = defaultVerifier.GetOidcVerifier("realm1")
			assert.Nil(t, e)
			assert.Contains(t, lastForwarded, "idp-dev")

			e = v1.Verify(jwtDefault)
			assert.NotNil(t, e)
			assert.Contains(t, e.Error(), "failed to verify signature")
		})

		t.Run("Ask for the same realm", func(t *testing.T) {
			v2, _ := defaultVerifier.GetOidcVerifier("realm1")
			assert.Equal(t, v1, v2)
			assert.Contains(t, lastForwarded, "idp-dev")
		})

		var v3 OidcVerifier
		t.Run("Ask for a different verifier", func(t *testing.T) {
			v3, _ = defaultVerifier.GetOidcVerifier("realm2")
			assert.NotEqual(t, v1, v3)
			assert.Contains(t, lastForwarded, "idp-dev")
		})

		t.Run("Ask for a white labeled verifier", func(t *testing.T) {
			v4, err := whiteLabelVerifier.GetOidcVerifier("realm3")
			assert.Nil(t, err)
			assert.NotNil(t, v4)
			assert.NotEqual(t, v1, v4)
			assert.NotEqual(t, v3, v4)
			assert.Contains(t, lastForwarded, "white.label.url")

			err = v4.Verify(jwtWhiteLabel)
			assert.NotNil(t, err)
			assert.Contains(t, err.Error(), "failed to verify signature")
		})
	}
}
