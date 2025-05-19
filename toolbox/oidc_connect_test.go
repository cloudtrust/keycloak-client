package toolbox

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"go.uber.org/mock/gomock"

	"github.com/cloudtrust/keycloak-client/v2"
	"github.com/cloudtrust/keycloak-client/v2/toolbox/mock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

type TestResponse struct {
	StatusCode   int
	NoBody       bool
	ResponseBody string
}

func (t *TestResponse) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(t.StatusCode)
	if !t.NoBody {
		w.Write([]byte(t.ResponseBody))
	}
	time.Sleep(20 * time.Millisecond)
}

func createTechnicalUser(realm string, user string, password string, clientID string) OAuth2Config {
	return OAuth2Config{
		Realm:    &realm,
		Username: &user,
		Password: &password,
		ClientID: &clientID,
	}
}

func createServiceAccount(realm string, clientID string, clientSecret string) OAuth2Config {
	return OAuth2Config{
		Realm:        &realm,
		ClientID:     &clientID,
		ClientSecret: &clientSecret,
	}
}

func TestCreateToken(t *testing.T) {
	var mockCtrl = gomock.NewController(t)
	defer mockCtrl.Finish()

	var mockLogger = mock.NewLogger(mockCtrl)
	mockLogger.EXPECT().Warn(gomock.Any(), gomock.Any()).AnyTimes()

	r := mux.NewRouter()
	r.Handle("/auth/realms/nobody/protocol/openid-connect/token", &TestResponse{StatusCode: http.StatusOK, NoBody: true})
	r.Handle("/auth/realms/invalid/protocol/openid-connect/token", &TestResponse{StatusCode: http.StatusUnauthorized})
	r.Handle("/auth/realms/bad-json/protocol/openid-connect/token", &TestResponse{StatusCode: http.StatusOK, ResponseBody: `{"truncated-`})

	ts := httptest.NewServer(r)
	defer ts.Close()

	var uriProvider, _ = NewKeycloakURIProviderFromArray([]string{ts.URL})
	var invalidURIProvider, _ = NewKeycloakURIProviderFromArray([]string{ts.URL + "0"})
	var ctx = context.TODO()

	var runFailingTest = func(t *testing.T, uriProvider keycloak.KeycloakURIProvider, realm string) {
		t.Run("Technical user", func(t *testing.T) {
			var creds = createTechnicalUser(realm, "user", "passwd", "clientID")
			var p = NewOAuth2TokenProvider(keycloak.Config{URIProvider: uriProvider}, creds, mockLogger)
			var _, err = p.ProvideToken(ctx)
			assert.NotNil(t, err)
		})
		t.Run("Service account", func(t *testing.T) {
			var creds = createServiceAccount(realm, "clientID", "client-secret")
			var p = NewOAuth2TokenProvider(keycloak.Config{URIProvider: uriProvider}, creds, mockLogger)
			var _, err = p.ProvideToken(ctx)
			assert.NotNil(t, err)
		})
	}

	t.Run("No body in HTTP response", func(t *testing.T) {
		runFailingTest(t, uriProvider, "nobody")
	})
	t.Run("Invalid credentials", func(t *testing.T) {
		runFailingTest(t, uriProvider, "invalid")
	})
	t.Run("Invalid JSON", func(t *testing.T) {
		runFailingTest(t, uriProvider, "bad-json")
	})
	t.Run("No HTTP response", func(t *testing.T) {
		runFailingTest(t, invalidURIProvider, "bad-json")
	})
}
