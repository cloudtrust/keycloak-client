package keycloak

import (
	"gopkg.in/h2non/gentleman.v2/plugins/body"
	"gopkg.in/h2non/gentleman.v2/plugins/url"
)

const (
	rolePath       = "/auth/admin/realms/:realm/roles"
	roleByIDPath   = "/auth/admin/realms/:realm/roles-by-id/:id"
	clientRolePath = "/auth/admin/realms/:realm/clients/:id/roles"
)

// GetClientRoles gets all roles for the realm or client
func (c *Client) GetClientRoles(accessToken string, realmName, idClient string) ([]RoleRepresentation, error) {
	var resp = []RoleRepresentation{}
	var err = c.get(accessToken, &resp, url.Path(clientRolePath), url.Param("realm", realmName), url.Param("id", idClient))
	return resp, err
}

// CreateClientRole creates a new role for the realm or client
func (c *Client) CreateClientRole(accessToken string, realmName, clientID string, role RoleRepresentation) (string, error) {
	return c.post(accessToken, nil, url.Path(clientRolePath), url.Param("realm", realmName), url.Param("id", clientID), body.JSON(role))
}

// GetRoles gets all roles for the realm or client
func (c *Client) GetRoles(accessToken string, realmName string) ([]RoleRepresentation, error) {
	var resp = []RoleRepresentation{}
	var err = c.get(accessToken, &resp, url.Path(rolePath), url.Param("realm", realmName))
	return resp, err
}

// GetRole gets a specific role’s representation
func (c *Client) GetRole(accessToken string, realmName string, roleID string) (RoleRepresentation, error) {
	var resp = RoleRepresentation{}
	var err = c.get(accessToken, &resp, url.Path(roleByIDPath), url.Param("realm", realmName), url.Param("id", roleID))
	return resp, err
}
