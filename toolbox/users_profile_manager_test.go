package toolbox

import (
	"context"
	"errors"
	"testing"

	"github.com/cloudtrust/keycloak-client/v2"
	"github.com/cloudtrust/keycloak-client/v2/toolbox/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

const (
	token = "access-token"
	realm = "my-realm"
)

var (
	errAny = errors.New("any error")
)

func TestGetRealmUserProfile(t *testing.T) {
	var mockCtrl = gomock.NewController(t)
	defer mockCtrl.Finish()

	var mockProfileRetriever = mock.NewProfileRetriever(mockCtrl)
	var mockTokenProvider = mock.NewOidcTokenProvider(mockCtrl)
	var cache = NewUserProfileCache(mockProfileRetriever, mockTokenProvider)

	var ctx = context.TODO()
	var expectedResult = keycloak.UserProfileRepresentation{
		Attributes: []keycloak.ProfileAttrbRepresentation{},
		Groups:     []keycloak.ProfileGroupRepresentation{},
	}
	expectedResult.InitDynamicAttributes()

	t.Run("Token provider fails", func(t *testing.T) {
		mockTokenProvider.EXPECT().ProvideTokenForRealm(gomock.Any(), realm).Return("", errAny)
		var _, err = cache.GetRealmUserProfile(ctx, realm)
		assert.Equal(t, errAny, err)
	})

	// Now, token will always be provided without error
	mockTokenProvider.EXPECT().ProvideTokenForRealm(gomock.Any(), realm).Return(token, nil).AnyTimes()

	t.Run("Keycloak fails", func(t *testing.T) {
		mockProfileRetriever.EXPECT().GetUserProfile(token, realm).Return(keycloak.UserProfileRepresentation{}, errAny)
		var _, err = cache.GetRealmUserProfile(ctx, realm)
		assert.Equal(t, errAny, err)
	})

	t.Run("Success - result is not cached yet", func(t *testing.T) {
		mockProfileRetriever.EXPECT().GetUserProfile(token, realm).Return(expectedResult, nil)
		var res, err = cache.GetRealmUserProfile(ctx, realm)
		assert.Nil(t, err)
		assert.Equal(t, expectedResult, res)
	})

	t.Run("Success - result is already cached", func(t *testing.T) {
		var res, err = cache.GetRealmUserProfile(ctx, realm)
		assert.Nil(t, err)
		assert.Equal(t, expectedResult, res)
	})
}

func TestGetRealmUserProfileWithToken(t *testing.T) {
	var mockCtrl = gomock.NewController(t)
	defer mockCtrl.Finish()

	var token = "access-token"
	var realm = "my-realm"
	var errAny = errors.New("any error")
	var expectedResult = keycloak.UserProfileRepresentation{
		Attributes: []keycloak.ProfileAttrbRepresentation{},
		Groups:     []keycloak.ProfileGroupRepresentation{},
	}
	expectedResult.InitDynamicAttributes()

	var mockProfileRetriever = mock.NewProfileRetriever(mockCtrl)
	var cache = NewUserProfileCache(mockProfileRetriever, nil)

	t.Run("Provide access token, Keycloak fails", func(t *testing.T) {
		mockProfileRetriever.EXPECT().GetUserProfile(token, realm).Return(keycloak.UserProfileRepresentation{}, errAny)
		var _, err = cache.GetRealmUserProfileWithToken(token, realm)
		assert.Equal(t, errAny, err)
	})

	t.Run("Provide access token, success", func(t *testing.T) {
		mockProfileRetriever.EXPECT().GetUserProfile(token, realm).Return(expectedResult, nil)
		var res, err = cache.GetRealmUserProfileWithToken(token, realm)
		assert.Nil(t, err)
		assert.Equal(t, expectedResult, res)
	})

	t.Run("Provide access token, cached value", func(t *testing.T) {
		var res, err = cache.GetRealmUserProfileWithToken(token, realm)
		assert.Nil(t, err)
		assert.Equal(t, expectedResult, res)
	})
}
