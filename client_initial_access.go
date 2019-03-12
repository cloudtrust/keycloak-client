package keycloak

import (
	"gopkg.in/h2non/gentleman.v2/plugins/body"
	"gopkg.in/h2non/gentleman.v2/plugins/url"
)

const (
	clientInitialAccessPath = "/auth/admin/realms/:realm/clients-initial-access"
)

// CreateClientInitialAccess creates a new initial access token.
func (c *Client) CreateClientInitialAccess(accessToken string, realmName string, access ClientInitialAccessCreatePresentation) (ClientInitialAccessPresentation, error) {
	var resp = ClientInitialAccessPresentation{}
	_, err := c.post(accessToken, &resp, nil, url.Path(clientInitialAccessPath), url.Param("realm", realmName), body.JSON(access))
	return resp, err
}

// GetClientInitialAccess returns a list of clients initial access.
func (c *Client) GetClientInitialAccess(accessToken string, realmName string) ([]ClientInitialAccessPresentation, error) {
	var resp = []ClientInitialAccessPresentation{}
	var err = c.get(accessToken, &resp, url.Path(clientInitialAccessPath), url.Param("realm", realmName))
	return resp, err
}

// DeleteClientInitialAccess deletes the client initial access.
func (c *Client) DeleteClientInitialAccess(accessToken string, realmName, accessID string) error {
	return c.delete(accessToken, url.Path(clientInitialAccessPath+"/:id"), url.Param("realm", realmName), url.Param("id", accessID))
}
