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
	anyError            = errors.New("any error")
	defaultProfileError = errors.New("default profile error")
)

func defaultProfile(theRealm string) (keycloak.UserProfileRepresentation, error) {
	if theRealm == realm {
		// Default profile can be recognized with its 5 groups
		return keycloak.UserProfileRepresentation{
			Attributes: nil,
			Groups: []keycloak.ProfileGroupRepresentation{
				{}, {}, {}, {}, {},
			},
		}, nil
	}
	return keycloak.UserProfileRepresentation{}, defaultProfileError
}

func TestGetRealmUserProfile(t *testing.T) {
	var mockCtrl = gomock.NewController(t)
	defer mockCtrl.Finish()

	var mockProfileRetriever = mock.NewProfileRetriever(mockCtrl)
	var mockTokenProvider = mock.NewOidcTokenProvider(mockCtrl)
	var cache = NewUserProfileCache(mockProfileRetriever, mockTokenProvider, defaultProfile)

	var ctx = context.TODO()
	var expectedResult = keycloak.UserProfileRepresentation{
		Attributes: []keycloak.ProfileAttrbRepresentation{},
		Groups:     []keycloak.ProfileGroupRepresentation{},
	}

	t.Run("Token provider fails", func(t *testing.T) {
		mockTokenProvider.EXPECT().ProvideTokenForRealm(gomock.Any(), realm).Return("", anyError)
		var _, err = cache.GetRealmUserProfile(ctx, realm)
		assert.Equal(t, anyError, err)
	})

	// Now, token will always be provided without error
	mockTokenProvider.EXPECT().ProvideTokenForRealm(gomock.Any(), realm).Return(token, nil).AnyTimes()

	t.Run("GetRealm fails", func(t *testing.T) {
		mockProfileRetriever.EXPECT().GetRealm(token, realm).Return(keycloak.RealmRepresentation{}, anyError)
		var _, err = cache.GetRealmUserProfile(ctx, realm)
		assert.Equal(t, anyError, err)
	})
	t.Run("UserProfile disabled, can't get a default user profile", func(t *testing.T) {
		var anotherRealm = "another-realm"
		mockTokenProvider.EXPECT().ProvideTokenForRealm(gomock.Any(), anotherRealm).Return(token, nil).AnyTimes()
		mockProfileRetriever.EXPECT().GetRealm(token, anotherRealm).Return(keycloak.RealmRepresentation{}, nil)
		var _, err = cache.GetRealmUserProfile(ctx, anotherRealm)
		assert.Equal(t, defaultProfileError, err)
	})
	t.Run("UserProfile disabled, gets a default user profile", func(t *testing.T) {
		mockProfileRetriever.EXPECT().GetRealm(token, realm).Return(keycloak.RealmRepresentation{}, nil)
		var res, err = cache.GetRealmUserProfile(ctx, realm)
		assert.Nil(t, err)
		assert.Len(t, res.Attributes, 0)
		assert.Len(t, res.Groups, 5)
	})

	// Now, token will always be provided without error and user profile enabled
	mockProfileRetriever.EXPECT().GetRealm(token, realm).Return(keycloak.RealmRepresentation{
		Attributes: ptrMapStringStringPtr(map[string]*string{"userProfileEnabled": ptr("true")}),
	}, nil).AnyTimes()

	t.Run("Keycloak fails", func(t *testing.T) {
		mockProfileRetriever.EXPECT().GetUserProfile(token, realm).Return(keycloak.UserProfileRepresentation{}, anyError)
		var _, err = cache.GetRealmUserProfile(ctx, realm)
		assert.Equal(t, anyError, err)
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
	var anyError = errors.New("any error")
	var expectedResult = keycloak.UserProfileRepresentation{
		Attributes: []keycloak.ProfileAttrbRepresentation{},
		Groups:     []keycloak.ProfileGroupRepresentation{},
	}

	var mockProfileRetriever = mock.NewProfileRetriever(mockCtrl)
	var cache = NewUserProfileCache(mockProfileRetriever, nil, defaultProfile)

	// user profile always enabled (disabled cases tested in TestGetRealmUserProfile)
	mockProfileRetriever.EXPECT().GetRealm(token, realm).Return(keycloak.RealmRepresentation{
		Attributes: ptrMapStringStringPtr(map[string]*string{"userProfileEnabled": ptr("true")}),
	}, nil).AnyTimes()

	t.Run("Provide access token, Keycloak fails", func(t *testing.T) {
		mockProfileRetriever.EXPECT().GetUserProfile(token, realm).Return(keycloak.UserProfileRepresentation{}, anyError)
		var _, err = cache.GetRealmUserProfileWithToken(token, realm)
		assert.Equal(t, anyError, err)
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
