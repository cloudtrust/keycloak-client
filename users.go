package keycloak

import (
	"fmt"

	"gopkg.in/h2non/gentleman.v2/plugins/body"
	"gopkg.in/h2non/gentleman.v2/plugins/url"
)

const (
	userPath                     = "/auth/admin/realms/:realm/users"
	userCountPath                = userPath + "/count"
	userIDPath                   = userPath + "/:id"
	userGroupsPath               = userIDPath + "/groups"
	resetPasswordPath            = userIDPath + "/reset-password"
	sendVerifyEmailPath          = userIDPath + "/send-verify-email"
	executeActionsEmailPath      = userIDPath + "/execute-actions-email"
	getCredentialsForUserPath    = "/auth/realms/:realmReq/api/admin/realms/:realm/users/:id/credentials"
	deleteCredentialsForUserPath = getCredentialsForUserPath + "/:credid"
)

// GetUsers returns a list of users, filtered according to the query parameters.
// Parameters: email, first (paging offset, int), firstName, lastName, username,
// max (maximum result size, default = 100),
// search (string contained in username, firstname, lastname or email)
func (c *Client) GetUsers(accessToken string, realmName string, paramKV ...string) ([]UserRepresentation, error) {
	if len(paramKV)%2 != 0 {
		return nil, fmt.Errorf("the number of key/val parameters should be even")
	}

	var resp = []UserRepresentation{}
	var plugins = append(createQueryPlugins(paramKV...), url.Path(userPath), url.Param("realm", realmName))
	var err = c.get(accessToken, &resp, plugins...)
	return resp, err
}

// CreateUser creates the user from its UserRepresentation. The username must be unique.
func (c *Client) CreateUser(accessToken string, realmName string, user UserRepresentation) (string, error) {
	return c.post(accessToken, nil, url.Path(userPath), url.Param("realm", realmName), body.JSON(user))
}

// CountUsers returns the number of users in the realm.
func (c *Client) CountUsers(accessToken string, realmName string) (int, error) {
	var resp = 0
	var err = c.get(accessToken, &resp, url.Path(userCountPath), url.Param("realm", realmName))
	return resp, err
}

// GetUser get the represention of the user.
func (c *Client) GetUser(accessToken string, realmName, userID string) (UserRepresentation, error) {
	var resp = UserRepresentation{}
	var err = c.get(accessToken, &resp, url.Path(userIDPath), url.Param("realm", realmName), url.Param("id", userID))
	return resp, err
}

// GetGroupsOfUser get the groups of the user.
func (c *Client) GetGroupsOfUser(accessToken string, realmName, userID string) ([]GroupRepresentation, error) {
	var resp = []GroupRepresentation{}
	var err = c.get(accessToken, &resp, url.Path(userGroupsPath), url.Param("realm", realmName), url.Param("id", userID))
	return resp, err
}

// UpdateUser updates the user.
func (c *Client) UpdateUser(accessToken string, realmName, userID string, user UserRepresentation) error {
	return c.put(accessToken, url.Path(userIDPath), url.Param("realm", realmName), url.Param("id", userID), body.JSON(user))
}

// DeleteUser deletes the user.
func (c *Client) DeleteUser(accessToken string, realmName, userID string) error {
	return c.delete(accessToken, url.Path(userIDPath), url.Param("realm", realmName), url.Param("id", userID))
}

// ResetPassword resets password of the user.
func (c *Client) ResetPassword(accessToken string, realmName, userID string, cred CredentialRepresentation) error {
	return c.put(accessToken, url.Path(resetPasswordPath), url.Param("realm", realmName), url.Param("id", userID), body.JSON(cred))
}

// SendVerifyEmail sends an email-verification email to the user An email contains a link the user can click to verify their email address.
func (c *Client) SendVerifyEmail(accessToken string, realmName string, userID string, paramKV ...string) error {
	if len(paramKV)%2 != 0 {
		return fmt.Errorf("the number of key/val parameters should be even")
	}

	var plugins = append(createQueryPlugins(paramKV...), url.Path(sendVerifyEmailPath), url.Param("realm", realmName), url.Param("id", userID))

	return c.put(accessToken, plugins...)
}

// ExecuteActionsEmail sends an update account email to the user. An email contains a link the user can click to perform a set of required actions.
func (c *Client) ExecuteActionsEmail(accessToken string, realmName string, userID string, actions []string, paramKV ...string) error {
	if len(paramKV)%2 != 0 {
		return fmt.Errorf("the number of key/val parameters should be even")
	}

	var plugins = append(createQueryPlugins(paramKV...), url.Path(executeActionsEmailPath), url.Param("realm", realmName), url.Param("id", userID), body.JSON(actions))

	return c.put(accessToken, plugins...)
}

// GetCredentialsForUser gets the credential list for a user
func (c *Client) GetCredentialsForUser(accessToken string, realmReq, realmName string, userID string) ([]CredentialRepresentation, error) {
	var resp = []CredentialRepresentation{}
	var err = c.get(accessToken, &resp, url.Path(getCredentialsForUserPath), url.Param("realmReq", realmReq), url.Param("realm", realmName), url.Param("id", userID))
	return resp, err
}

// DeleteCredentialsForUser remove credentials for a user
func (c *Client) DeleteCredentialsForUser(accessToken string, realmReq, realmName string, userID string, credentialID string) error {
	return c.delete(accessToken, url.Path(deleteCredentialsForUserPath), url.Param("realmReq", realmReq), url.Param("realm", realmName), url.Param("id", userID), url.Param("credid", userID))
}
