package toolbox

import (
	"context"

	"github.com/cloudtrust/keycloak-client/v2"
)

// ProfileRetriever interface
type ProfileRetriever interface {
	GetRealm(accessToken string, realmName string) (keycloak.RealmRepresentation, error)
	GetUserProfile(accessToken string, realmName string) (keycloak.UserProfileRepresentation, error)
}

// UserProfileCache struct
type UserProfileCache struct {
	tokenProvider  OidcTokenProvider
	retriever      ProfileRetriever
	cachedProfiles map[string]keycloak.UserProfileRepresentation
}

// NewUserProfileCache creates a UserProfileCache instance
func NewUserProfileCache(retriever ProfileRetriever, tokenProvider OidcTokenProvider) *UserProfileCache {
	return &UserProfileCache{
		tokenProvider:  tokenProvider,
		retriever:      retriever,
		cachedProfiles: map[string]keycloak.UserProfileRepresentation{},
	}
}

// GetRealmUserProfile gets the realm users profile using the token provider provided when creating the UserProfileCache instance
func (upc *UserProfileCache) GetRealmUserProfile(ctx context.Context, realmName string) (keycloak.UserProfileRepresentation, error) {
	return upc.getRealmUserProfile(func() (string, error) {
		return upc.tokenProvider.ProvideTokenForRealm(ctx, realmName)
	}, realmName)
}

// GetRealmUserProfileWithToken gets the realm users profile using the provided access token
func (upc *UserProfileCache) GetRealmUserProfileWithToken(accessToken string, realmName string) (keycloak.UserProfileRepresentation, error) {
	var f = func() (string, error) {
		return accessToken, nil
	}
	return upc.getRealmUserProfile(f, realmName)
}

func (upc *UserProfileCache) getRealmUserProfile(provideToken func() (string, error), realmName string) (keycloak.UserProfileRepresentation, error) {
	// Already loaded profile?
	if profile, ok := upc.getProfile(realmName); ok {
		return profile, nil
	}
	// Get access token
	var accessToken, err = provideToken()
	if err != nil {
		return keycloak.UserProfileRepresentation{}, err
	}
	// Retrieve the profile
	var profile keycloak.UserProfileRepresentation
	profile, err = upc.retriever.GetUserProfile(accessToken, realmName)
	if err != nil {
		return keycloak.UserProfileRepresentation{}, err
	}

	profile.InitDynamicAttributes()

	// Write profile in cache then return it
	upc.setProfile(realmName, profile)
	return profile, nil
}

func (upc *UserProfileCache) getProfile(realm string) (keycloak.UserProfileRepresentation, bool) {
	res, ok := upc.cachedProfiles[realm]
	return res, ok
}

func (upc *UserProfileCache) setProfile(realm string, profile keycloak.UserProfileRepresentation) {
	upc.cachedProfiles[realm] = profile
}
