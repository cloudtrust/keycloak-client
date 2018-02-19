package client

import (
	"gopkg.in/h2non/gentleman.v2/plugins/body"
	"gopkg.in/h2non/gentleman.v2/plugins/url"
)

const (
	userPath      = "/auth/admin/realms/:realm/users"
	userCountPath = userPath + "/count"
	userIDPath    = userPath + "/:id"
)

func (c *Client) GetUsers(realm string) ([]UserRepresentation, error) {
	var resp = []UserRepresentation{}
	var err = c.get(&resp, url.Path(userPath), url.Param("realm", realm))
	return resp, err
}

func (c *Client) CreateUser(realm string, user UserRepresentation) error {
	return c.post(url.Path(userPath), url.Param("realm", realm), body.JSON(user))
}

func (c *Client) CountUsers(realm string) (int, error) {
	var resp = 0
	var err = c.get(&resp, url.Path(userCountPath), url.Param("realm", realm))
	return resp, err
}

func (c *Client) GetUser(realm, userID string) (UserRepresentation, error) {
	var resp = UserRepresentation{}
	var err = c.get(&resp, url.Path(userIDPath), url.Param("realm", realm), url.Param("id", userID))
	return resp, err
}

func (c *Client) UpdateUser(realm, userID string, user UserRepresentation) error {
	return c.put(url.Path(userIDPath), url.Param("realm", realm), url.Param("id", userID), body.JSON(user))
}

func (c *Client) DeleteUser(realm, userID string) error {
	return c.delete(url.Path(userIDPath), url.Param("realm", realm), url.Param("id", userID))
}
