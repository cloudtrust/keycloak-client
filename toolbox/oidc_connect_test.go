package toolbox

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	"github.com/cloudtrust/keycloak-client"
	"github.com/cloudtrust/keycloak-client/toolbox/mock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

type TestResponse struct {
	StatusCode   int
	NoBody       bool
	ResponseBody string
}

func (t *TestResponse) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(t.StatusCode)
	if !t.NoBody {
		w.Write([]byte(t.ResponseBody))
	}
	time.Sleep(20 * time.Millisecond)
}

func TestCreateToken(t *testing.T) {
	var mockCtrl = gomock.NewController(t)
	defer mockCtrl.Finish()

	var mockLogger = mock.NewLogger(mockCtrl)
	mockLogger.EXPECT().Warn(gomock.Any(), gomock.Any()).AnyTimes()

	var oidcToken1 = oidcToken{
		AccessToken: `eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJZRTUtNUpBb2NOcG5zeEpEaGRyYVdyWlVCZTBMR2xfanNILUtnb1EwWi1FIn0.eyJqdGkiOiIxYjJkZDY2NS01ZGE1LTRiMzAtODY0MS0wNWQ4ZTk0NTQ2ZWQiLCJleHAiOjE1NzkyMDQ0ODYsIm5iZiI6MCwiaWF0IjoxNTc5MTY4NDg2LCJpc3MiOiJodHRwOi8vMTI3LjAuMC4xOjgwODAvYXV0aC9yZWFsbXMvbWFzdGVyIiwiYXVkIjpbInBhc3NmbG93LXJlYWxtIiwibXlfcmVhbG0tcmVhbG0iLCJDbG91ZHRydXN0LXJlYWxtIiwibWFzdGVyLXJlYWxtIiwiYWNjb3VudCJdLCJzdWIiOiI3OTU5MjhjMy03N2Y1LTRmMjQtOTI0NC02NzBkMGJmMDJhMmQiLCJ0eXAiOiJCZWFyZXIiLCJhenAiOiJhZG1pbi1jbGkiLCJhdXRoX3RpbWUiOjAsInNlc3Npb25fc3RhdGUiOiI5ZTQ5NDA1MS1kZGQ1LTRhODctYTczZC1hOWU5YjMwYmFlZGEiLCJhY3IiOiIxIiwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbImNyZWF0ZS1yZWFsbSIsIm9mZmxpbmVfYWNjZXNzIiwiYWRtaW4iLCJ1bWFfYXV0aG9yaXphdGlvbiJdfSwicmVzb3VyY2VfYWNjZXNzIjp7InBhc3NmbG93LXJlYWxtIjp7InJvbGVzIjpbInZpZXctcmVhbG0iLCJ2aWV3LWlkZW50aXR5LXByb3ZpZGVycyIsIm1hbmFnZS1pZGVudGl0eS1wcm92aWRlcnMiLCJpbXBlcnNvbmF0aW9uIiwiY3JlYXRlLWNsaWVudCIsIm1hbmFnZS11c2VycyIsInF1ZXJ5LXJlYWxtcyIsInZpZXctYXV0aG9yaXphdGlvbiIsInF1ZXJ5LWNsaWVudHMiLCJxdWVyeS11c2VycyIsIm1hbmFnZS1ldmVudHMiLCJtYW5hZ2UtcmVhbG0iLCJ2aWV3LWV2ZW50cyIsInZpZXctdXNlcnMiLCJ2aWV3LWNsaWVudHMiLCJtYW5hZ2UtYXV0aG9yaXphdGlvbiIsIm1hbmFnZS1jbGllbnRzIiwicXVlcnktZ3JvdXBzIl19LCJteV9yZWFsbS1yZWFsbSI6eyJyb2xlcyI6WyJ2aWV3LXJlYWxtIiwidmlldy1pZGVudGl0eS1wcm92aWRlcnMiLCJtYW5hZ2UtaWRlbnRpdHktcHJvdmlkZXJzIiwiaW1wZXJzb25hdGlvbiIsImNyZWF0ZS1jbGllbnQiLCJtYW5hZ2UtdXNlcnMiLCJxdWVyeS1yZWFsbXMiLCJ2aWV3LWF1dGhvcml6YXRpb24iLCJxdWVyeS1jbGllbnRzIiwicXVlcnktdXNlcnMiLCJtYW5hZ2UtZXZlbnRzIiwibWFuYWdlLXJlYWxtIiwidmlldy1ldmVudHMiLCJ2aWV3LXVzZXJzIiwidmlldy1jbGllbnRzIiwibWFuYWdlLWF1dGhvcml6YXRpb24iLCJtYW5hZ2UtY2xpZW50cyIsInF1ZXJ5LWdyb3VwcyJdfSwiQ2xvdWR0cnVzdC1yZWFsbSI6eyJyb2xlcyI6WyJ2aWV3LWlkZW50aXR5LXByb3ZpZGVycyIsInZpZXctcmVhbG0iLCJtYW5hZ2UtaWRlbnRpdHktcHJvdmlkZXJzIiwiaW1wZXJzb25hdGlvbiIsImNyZWF0ZS1jbGllbnQiLCJtYW5hZ2UtdXNlcnMiLCJxdWVyeS1yZWFsbXMiLCJ2aWV3LWF1dGhvcml6YXRpb24iLCJxdWVyeS1jbGllbnRzIiwicXVlcnktdXNlcnMiLCJtYW5hZ2UtZXZlbnRzIiwibWFuYWdlLXJlYWxtIiwidmlldy1ldmVudHMiLCJ2aWV3LXVzZXJzIiwidmlldy1jbGllbnRzIiwibWFuYWdlLWF1dGhvcml6YXRpb24iLCJtYW5hZ2UtY2xpZW50cyIsInF1ZXJ5LWdyb3VwcyJdfSwibWFzdGVyLXJlYWxtIjp7InJvbGVzIjpbInZpZXctcmVhbG0iLCJ2aWV3LWlkZW50aXR5LXByb3ZpZGVycyIsIm1hbmFnZS1pZGVudGl0eS1wcm92aWRlcnMiLCJpbXBlcnNvbmF0aW9uIiwiY3JlYXRlLWNsaWVudCIsIm1hbmFnZS11c2VycyIsInF1ZXJ5LXJlYWxtcyIsInZpZXctYXV0aG9yaXphdGlvbiIsInF1ZXJ5LWNsaWVudHMiLCJxdWVyeS11c2VycyIsIm1hbmFnZS1ldmVudHMiLCJtYW5hZ2UtcmVhbG0iLCJ2aWV3LWV2ZW50cyIsInZpZXctdXNlcnMiLCJ2aWV3LWNsaWVudHMiLCJtYW5hZ2UtYXV0aG9yaXphdGlvbiIsIm1hbmFnZS1jbGllbnRzIiwicXVlcnktZ3JvdXBzIl19LCJhY2NvdW50Ijp7InJvbGVzIjpbIm1hbmFnZS1hY2NvdW50IiwibWFuYWdlLWFjY291bnQtbGlua3MiLCJ2aWV3LXByb2ZpbGUiXX19LCJzY29wZSI6ImVtYWlsIHByb2ZpbGUgZ3JvdXBzIiwiZW1haWxfdmVyaWZpZWQiOmZhbHNlLCJncm91cHMiOlsiL3RvZV9hZG1pbmlzdHJhdG9yIl0sInByZWZlcnJlZF91c2VybmFtZSI6ImFkbWluIn0.GrPUDdQUID0S38ZSwVqarAuSDXrl5cJju3uq32y6bNGPK9a8jcHBP5FEfMcZ3vieQPtWFeySMycwTEcH6x6lc-9bj1w5veL4yyTA1zk_ERPshfiobk0u94vuljnoz-PW7JvBLOy47Bk5cP9pHPbJMPY0kOFHTpXZHd6KwfcE_X8gizLw4rDhIpK1NEtABQVzUNvP9fDZOm2I1PHJbl0odRE7EFu9Xh5ya8DaUQ2RKUb0E5csnA3DYlFdEhtMV1MAKRzqplDzj8zLQ8f8fflzC9_g4vmnDUEKSBxq1f1qKmzm1-XUuqRYTNWHfOtRR9rXrEzn-6fymFcRHIVGW7kgzg`,
		ExpiresIn:   3600,
	}
	var oidcToken2 = oidcToken{
		AccessToken: `eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJZRTUtNUpBb2NOcG5zeEpEaGRyYVdyWlVCZTBMR2xfanNILUtnb1EwWi1FIn0.eyJqdGkiOiAiMWIyZGQ2NjUtNWRhNS00YjMwLTg2NDEtMDVkOGU5NDU0NmVkIiwiZXhwIjogMTU3OTIwNDQ4NiwibmJmIjogMCwiaWF0IjogMTU3OTE2ODQ4NiwiaXNzIjogImh0dHA6Ly9zZWNvbmQtaWRwL2F1dGgvcmVhbG1zL21hc3RlciIsImF1ZCI6IFsiQ2xvdWR0cnVzdC1yZWFsbSIsImFjY291bnQiXSwic3ViIjogIjc5NTkyOGMzLTc3ZjUtNGYyNC05MjQ0LTY3MGQwYmYwMmEyZCIsInR5cCI6ICJCZWFyZXIiLCJhenAiOiAiYWRtaW4tY2xpIiwiYXV0aF90aW1lIjogMCwic2Vzc2lvbl9zdGF0ZSI6ICI5ZTQ5NDA1MS1kZGQ1LTRhODctYTczZC1hOWU5YjMwYmFlZGEiLCJhY3IiOiAiMSIsInJlYWxtX2FjY2VzcyI6IHsicm9sZXMiOiBbImFkbWluIl19LCJyZXNvdXJjZV9hY2Nlc3MiOiB7IkNsb3VkdHJ1c3QtcmVhbG0iOiB7InJvbGVzIjogWyJ2aWV3LWlkZW50aXR5LXByb3ZpZGVycyIsInZpZXctcmVhbG0iLCJ2aWV3LXVzZXJzIiwicXVlcnktZ3JvdXBzIl19LCJhY2NvdW50IjogeyJyb2xlcyI6IFsibWFuYWdlLWFjY291bnQiLCJtYW5hZ2UtYWNjb3VudC1saW5rcyIsInZpZXctcHJvZmlsZSJdfX0sInNjb3BlIjogImVtYWlsIHByb2ZpbGUgZ3JvdXBzIiwiZW1haWxfdmVyaWZpZWQiOiBmYWxzZSwiZ3JvdXBzIjogWyIvdG9lX2FkbWluaXN0cmF0b3IiXSwicHJlZmVycmVkX3VzZXJuYW1lIjogImFkbWluIn0=.GrPUDdQUID0S38ZSwVqarAuSDXrl5cJju3uq32y6bNGPK9a8jcHBP5FEfMcZ3vieQPtWFeySMycwTEcH6x6lc-9bj1w5veL4yyTA1zk_ERPshfiobk0u94vuljnoz-PW7JvBLOy47Bk5cP9pHPbJMPY0kOFHTpXZHd6KwfcE_X8gizLw4rDhIpK1NEtABQVzUNvP9fDZOm2I1PHJbl0odRE7EFu9Xh5ya8DaUQ2RKUb0E5csnA3DYlFdEhtMV1MAKRzqplDzj8zLQ8f8fflzC9_g4vmnDUEKSBxq1f1qKmzm1-XUuqRYTNWHfOtRR9rXrEzn-6fymFcRHIVGW7kgzg`,
		ExpiresIn:   3600,
	}

	var validJSON, _ = json.Marshal(oidcToken1)
	var validJSON2, _ = json.Marshal(oidcToken2)

	r := mux.NewRouter()
	r.Handle("/auth/realms/nobody/protocol/openid-connect/token", &TestResponse{StatusCode: http.StatusOK, NoBody: true})
	r.Handle("/auth/realms/invalid/protocol/openid-connect/token", &TestResponse{StatusCode: http.StatusUnauthorized})
	r.Handle("/auth/realms/bad-json/protocol/openid-connect/token", &TestResponse{StatusCode: http.StatusOK, ResponseBody: `{"truncated-`})
	r.Handle("/auth/realms/valid/protocol/openid-connect/token", &TestResponse{StatusCode: http.StatusOK, ResponseBody: string(validJSON)})
	r.Handle("/second/auth/realms/valid/protocol/openid-connect/token", &TestResponse{StatusCode: http.StatusOK, ResponseBody: string(validJSON2)})

	ts := httptest.NewServer(r)
	defer ts.Close()

	var uriProvider, _ = NewKeycloakURIProviderFromArray([]string{ts.URL})
	var invalidURIProvider, _ = NewKeycloakURIProviderFromArray([]string{ts.URL + "0"})
	var ctx = context.TODO()

	t.Run("No body in HTTP response", func(t *testing.T) {
		var p = NewOidcTokenProvider(keycloak.Config{URIProvider: uriProvider}, "nobody", "user", "passwd", "clientID", mockLogger)
		var _, err = p.ProvideToken(ctx)
		assert.NotNil(t, err)
	})

	t.Run("Invalid credentials", func(t *testing.T) {
		var p = NewOidcTokenProvider(keycloak.Config{URIProvider: uriProvider}, "invalid", "user", "passwd", "clientID", mockLogger)
		var _, err = p.ProvideToken(ctx)
		assert.NotNil(t, err)
	})

	t.Run("Invalid JSON", func(t *testing.T) {
		var p = NewOidcTokenProvider(keycloak.Config{URIProvider: uriProvider}, "bad-json", "user", "passwd", "clientID", mockLogger)
		var _, err = p.ProvideToken(ctx)
		assert.NotNil(t, err)
	})

	t.Run("No HTTP response", func(t *testing.T) {
		var p = NewOidcTokenProvider(keycloak.Config{URIProvider: invalidURIProvider}, "bad-json", "user", "passwd", "clientID", mockLogger)
		var _, err = p.ProvideToken(ctx)
		assert.NotNil(t, err)
	})

	t.Run("Valid credentials", func(t *testing.T) {
		var p = NewOidcTokenProvider(keycloak.Config{URIProvider: uriProvider}, "valid", "user", "passwd", "clientID", mockLogger)

		var timeStart = time.Now()

		// First call
		var token, err = p.ProvideToken(ctx)
		assert.Nil(t, err)
		assert.NotEqual(t, "", token)

		var timeAfterFirstCall = time.Now()

		// Second call
		token, err = p.ProvideToken(ctx)
		assert.Nil(t, err)
		assert.NotEqual(t, "", token)

		var timeAfterSecondCall = time.Now()

		var withHTTPDuration = int64(20 * time.Millisecond)
		var withoutHTTPDuration = int64(5 * time.Millisecond)
		var duration1 = timeAfterFirstCall.Sub(timeStart).Nanoseconds()
		var duration2 = timeAfterSecondCall.Sub(timeAfterFirstCall).Nanoseconds()
		var msg = fmt.Sprintf("Durations: no valid token loaded yet:%d (expected > %d), token not expired:%d (expected < %d)", duration1, withHTTPDuration, duration2, withoutHTTPDuration)
		assert.True(t, duration1 > withHTTPDuration, msg)
		assert.True(t, duration2 < withoutHTTPDuration, msg)
	})

	t.Run("Multiple issuers", func(t *testing.T) {
		var anotherIssuer = "second"
		var targets = map[string]string{"*": ts.URL, anotherIssuer: ts.URL + "/second"}
		var kup, _ = NewKeycloakURIProvider(targets, "*")
		var cfg = keycloak.Config{
			URIProvider: kup,
			Timeout:     time.Second,
		}
		var p = NewOidcTokenProvider(cfg, "valid", "user", "passwd", "clientID", mockLogger)

		var token1, err = p.ProvideTokenForRealm(ctx, "any")
		assert.Nil(t, err)
		assert.NotEqual(t, "", token1)

		var token2 string
		token2, err = p.ProvideTokenForRealm(ctx, "another")
		assert.Nil(t, err)
		assert.Equal(t, token1, token2)

		token2, err = p.ProvideTokenForRealm(ctx, anotherIssuer)
		assert.Nil(t, err)
		assert.NotEqual(t, token1, token2)
	})
}
