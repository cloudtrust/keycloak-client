package api

import (
	"github.com/cloudtrust/keycloak-client"
	"gopkg.in/h2non/gentleman.v2/plugins/body"
	"gopkg.in/h2non/gentleman.v2/plugins/url"
)

const (
	clientRoleMappingPath = "/auth/admin/realms/:realm/users/:id/role-mappings/clients/:client"
	realmRoleMappingPath  = "/auth/admin/realms/:realm/users/:id/role-mappings/realm"
)

// AddClientRolesToUserRoleMapping add client-level roles to the user role mapping.
func (c *Client) AddClientRolesToUserRoleMapping(accessToken string, realmName, userID, clientID string, roles []keycloak.RoleRepresentation) error {
	_, err := c.post(accessToken, nil, url.Path(clientRoleMappingPath), url.Param("realm", realmName), url.Param("id", userID), url.Param("client", clientID), body.JSON(roles))
	return err
}

// GetClientRoleMappings gets client-level role mappings for the user, and the app.
func (c *Client) GetClientRoleMappings(accessToken string, realmName, userID, clientID string) ([]keycloak.RoleRepresentation, error) {
	var resp = []keycloak.RoleRepresentation{}
	var err = c.get(accessToken, &resp, url.Path(clientRoleMappingPath), url.Param("realm", realmName), url.Param("id", userID), url.Param("client", clientID))
	return resp, err
}

// DeleteClientRolesFromUserRoleMapping deletes client-level roles from user role mapping.
func (c *Client) DeleteClientRolesFromUserRoleMapping(accessToken string, realmName, userID, clientID string) error {
	return c.delete(accessToken, url.Path(clientRoleMappingPath), url.Param("realm", realmName), url.Param("id", userID), url.Param("client", clientID))
}

// GetRealmLevelRoleMappings gets realm level role mappings
func (c *Client) GetRealmLevelRoleMappings(accessToken string, realmName, userID string) ([]keycloak.RoleRepresentation, error) {
	var resp = []keycloak.RoleRepresentation{}
	var err = c.get(accessToken, &resp, url.Path(realmRoleMappingPath), url.Param("realm", realmName), url.Param("id", userID))
	return resp, err
}

// AddRealmLevelRoleMappings gets realm level role mappings
func (c *Client) AddRealmLevelRoleMappings(accessToken string, realmName, userID string, roles []keycloak.RoleRepresentation) error {
	var _, err = c.post(accessToken, nil, url.Path(realmRoleMappingPath), url.Param("realm", realmName), url.Param("id", userID), body.JSON(roles))
	return err
}

// DeleteRealmLevelRoleMappings gets realm level role mappings
func (c *Client) DeleteRealmLevelRoleMappings(accessToken string, realmName, userID string, roles []keycloak.RoleRepresentation) error {
	return c.delete(accessToken, url.Path(realmRoleMappingPath), url.Param("realm", realmName), url.Param("id", userID), body.JSON(roles))
}
