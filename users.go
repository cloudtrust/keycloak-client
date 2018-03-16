package keycloak

import (
	"fmt"

	"gopkg.in/h2non/gentleman.v2/plugins/body"
	"gopkg.in/h2non/gentleman.v2/plugins/url"
)

const (
	userPath      = "/auth/admin/realms/:realm/users"
	userCountPath = userPath + "/count"
	userIDPath    = userPath + "/:id"
)

// GetUsers returns a list of users, filtered according to the query parameters.
// Parameters: email, first (paging offset, int), firstName, lastName, username,
// max (maximum result size, default = 100),
// search (string contained in username, firstname, lastname or email)
func (c *Client) GetUsers(realmName string, paramKV ...string) ([]UserRepresentation, error) {
	if len(paramKV)%2 != 0 {
		return nil, fmt.Errorf("the number of key/val parameters should be even")
	}

	var resp = []UserRepresentation{}
	var plugins = append(createQueryPlugins(paramKV...), url.Path(userPath), url.Param("realm", realmName))
	var err = c.get(&resp, plugins...)
	return resp, err
}

// CreateUser creates the user from its UserRepresentation. The username must be unique.
func (c *Client) CreateUser(realm string, user UserRepresentation) error {
	return c.post(url.Path(userPath), url.Param("realm", realm), body.JSON(user))
}

// CountUsers returns the number of users in the realm.
func (c *Client) CountUsers(realmName string) (int, error) {
	var resp = 0
	var err = c.get(&resp, url.Path(userCountPath), url.Param("realm", realmName))
	return resp, err
}

// GetUser get the represention of the user.
func (c *Client) GetUser(realmName, userID string) (UserRepresentation, error) {
	var resp = UserRepresentation{}
	var err = c.get(&resp, url.Path(userIDPath), url.Param("realm", realmName), url.Param("id", userID))
	return resp, err
}

// UpdateUser updates the user.
func (c *Client) UpdateUser(realmName, userID string, user UserRepresentation) error {
	return c.put(url.Path(userIDPath), url.Param("realm", realmName), url.Param("id", userID), body.JSON(user))
}

// DeleteUser deletes the user.
func (c *Client) DeleteUser(realmName, userID string) error {
	return c.delete(url.Path(userIDPath), url.Param("realm", realmName), url.Param("id", userID))
}
