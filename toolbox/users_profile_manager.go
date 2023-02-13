package toolbox

import (
	"context"
	"encoding/json"

	"github.com/cloudtrust/keycloak-client/v2"
)

// ProfileRetriever interface
type ProfileRetriever interface {
	GetRealm(accessToken string, realmName string) (keycloak.RealmRepresentation, error)
	GetUserProfile(accessToken string, realmName string) (keycloak.UserProfileRepresentation, error)
}

// DefaultProfileProviderFunc function type
type DefaultProfileProviderFunc func(realmName string) (keycloak.UserProfileRepresentation, error)

// UserProfileCache struct
type UserProfileCache struct {
	tokenProvider      OidcTokenProvider
	retriever          ProfileRetriever
	cachedProfiles     map[string][]byte
	defProfileProvider DefaultProfileProviderFunc
}

func NewUserProfileCache(retriever ProfileRetriever, tokenProvider OidcTokenProvider, defProfileProvider DefaultProfileProviderFunc) *UserProfileCache {
	return &UserProfileCache{
		tokenProvider:      tokenProvider,
		retriever:          retriever,
		defProfileProvider: defProfileProvider,
		cachedProfiles:     map[string][]byte{},
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
	// Is user profile enabled for this realm?
	var realmRep keycloak.RealmRepresentation
	realmRep, err = upc.retriever.GetRealm(accessToken, realmName)
	if err != nil {
		return keycloak.UserProfileRepresentation{}, err
	}
	if !realmRep.IsUserProfileEnabled() {
		return upc.defProfileProvider(realmName)
	}
	// Retrieve the profile
	var profile keycloak.UserProfileRepresentation
	profile, err = upc.retriever.GetUserProfile(accessToken, realmName)
	if err != nil {
		return keycloak.UserProfileRepresentation{}, err
	}
	// Write profile in cache then return it
	upc.setProfile(realmName, profile)
	return profile, nil
}

func (upc *UserProfileCache) getProfile(realm string) (keycloak.UserProfileRepresentation, bool) {
	var res keycloak.UserProfileRepresentation
	if bytes, ok := upc.cachedProfiles[realm]; ok {
		json.Unmarshal(bytes, &res)
		return res, true
	}
	return res, false
}

func (upc *UserProfileCache) setProfile(realm string, profile keycloak.UserProfileRepresentation) {
	var bytes, _ = json.Marshal(profile)
	upc.cachedProfiles[realm] = bytes
}
