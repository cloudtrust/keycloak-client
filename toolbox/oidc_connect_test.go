package toolbox

import (
	"context"
	"errors"
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

func TestNewOAuth2Config(t *testing.T) {
	var mockCtrl = gomock.NewController(t)
	defer mockCtrl.Finish()

	var (
		realm           = "the-realm"
		username        = "the-username"
		password        = "the-password"
		clientID        = "the-client-id"
		clientSecretKey = "key-client-secret"
		usernameKey     = "key-username"
		passwordKey     = "key-password"
		anError         = errors.New("an error")

		strLoader = func(key string) string {
			return key
		}
	)

	t.Run("Failure", func(t *testing.T) {
		loaderFailure := func(cfg *OAuth2Config) error {
			return anError
		}

		_, err := NewOAuth2Config(loaderFailure, strLoader, clientSecretKey, usernameKey, passwordKey)
		assert.NotNil(t, err)
	})
	t.Run("Success with ClientID", func(t *testing.T) {
		loaderSuccess := func(cfg *OAuth2Config) error {
			*cfg = OAuth2Config{
				Realm:    &realm,
				ClientID: &clientID,
				Username: &username,
				Password: &password,
			}
			return nil
		}

		res, err := NewOAuth2Config(loaderSuccess, strLoader, clientSecretKey, usernameKey, passwordKey)
		assert.Nil(t, err)
		assert.Equal(t, realm, *res.Realm)
		assert.Equal(t, usernameKey, *res.Username) // overriden by env variable
		assert.Equal(t, passwordKey, *res.Password) // overriden by env variable
		assert.Equal(t, clientID, *res.ClientID)
		assert.Nil(t, res.ClientSecret)
	})
	t.Run("Success with ClientSecret", func(t *testing.T) {
		loaderSuccess := func(cfg *OAuth2Config) error {
			*cfg = OAuth2Config{
				Realm:    &realm,
				ClientID: &clientID,
			}
			return nil
		}

		res, err := NewOAuth2Config(loaderSuccess, strLoader, clientSecretKey, usernameKey, passwordKey)
		assert.Nil(t, err)
		assert.Equal(t, realm, *res.Realm)
		assert.Nil(t, res.Username)
		assert.Nil(t, res.Password)
		assert.Equal(t, clientID, *res.ClientID)
		assert.Equal(t, clientSecretKey, *res.ClientSecret)
	})
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
