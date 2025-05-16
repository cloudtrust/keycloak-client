package api

import (
	"github.com/cloudtrust/keycloak-client/v2"
	"gopkg.in/h2non/gentleman.v2/plugins/body"
	"gopkg.in/h2non/gentleman.v2/plugins/query"
	"gopkg.in/h2non/gentleman.v2/plugins/url"
)

const (
	// API Keycloak out-of-the-box
	kcRolePath       = "/auth/admin/realms/:realm/roles"
	kcRoleByIDPath   = "/auth/admin/realms/:realm/roles-by-id/:id"
	kcClientRolePath = "/auth/admin/realms/:realm/clients/:id/roles"
)

// GetClientRoles gets all roles for the realm or client
func (c *Client) GetClientRoles(accessToken string, realmName, idClient string) ([]keycloak.RoleRepresentation, error) {
	var resp = []keycloak.RoleRepresentation{}
	var err = c.forRealm(realmName).
		get(accessToken, &resp, url.Path(kcClientRolePath), url.Param("realm", realmName), url.Param("id", idClient))
	return resp, err
}

// CreateClientRole creates a new role for the realm or client
func (c *Client) CreateClientRole(accessToken string, realmName, clientID string, role keycloak.RoleRepresentation) (string, error) {
	return c.forRealm(realmName).
		post(accessToken, nil, url.Path(kcClientRolePath), url.Param("realm", realmName), url.Param("id", clientID), body.JSON(role))
}

// GetRoles gets all roles for the realm or client
func (c *Client) GetRoles(accessToken string, realmName string) ([]keycloak.RoleRepresentation, error) {
	var resp = []keycloak.RoleRepresentation{}
	var err = c.forRealm(realmName).
		get(accessToken, &resp, url.Path(kcRolePath), url.Param("realm", realmName))
	return resp, err
}

// GetRolesWithAttributes gets all roles for the realm or client with their attributes
func (c *Client) GetRolesWithAttributes(accessToken string, realmName string) ([]keycloak.RoleRepresentation, error) {
	var resp = []keycloak.RoleRepresentation{}
	var err = c.forRealm(realmName).
		get(accessToken, &resp, url.Path(kcRolePath), url.Param("realm", realmName), query.Add("briefRepresentation", "false"))
	return resp, err
}

// GetRole gets a specific role’s representation
func (c *Client) GetRole(accessToken string, realmName string, roleID string) (keycloak.RoleRepresentation, error) {
	var resp = keycloak.RoleRepresentation{}
	var err = c.forRealm(realmName).
		get(accessToken, &resp, url.Path(kcRoleByIDPath), url.Param("realm", realmName), url.Param("id", roleID))
	return resp, err
}

// CreateRole creates a new role in a realm
func (c *Client) CreateRole(accessToken string, realmName string, role keycloak.RoleRepresentation) (string, error) {
	return c.forRealm(realmName).
		post(accessToken, nil, url.Path(kcRolePath), url.Param("realm", realmName), body.JSON(role))
}

// UpdateRole updates a role in a realm
func (c *Client) UpdateRole(accessToken string, realmName string, roleID string, role keycloak.RoleRepresentation) error {
	return c.forRealm(realmName).
		put(accessToken, url.Path(kcRoleByIDPath), url.Param("realm", realmName), url.Param("id", roleID), body.JSON(role))
}

// DeleteRole deletes a role in a realm
func (c *Client) DeleteRole(accessToken string, realmName string, roleID string) error {
	return c.forRealm(realmName).
		delete(accessToken, url.Path(kcRoleByIDPath), url.Param("realm", realmName), url.Param("id", roleID))
}
