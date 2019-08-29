package keycloak

import (
	"gopkg.in/h2non/gentleman.v2/plugins/body"
	"gopkg.in/h2non/gentleman.v2/plugins/url"
)

const (
	clientRoleMappingPath = "/auth/admin/realms/:realm/users/:id/role-mappings/clients/:client"
	realmRoleMappingPath  = "/auth/admin/realms/:realm/users/:id/role-mappings/realm"
)

// AddClientRolesToUserRoleMapping add client-level roles to the user role mapping.
func (c *Client) AddClientRolesToUserRoleMapping(accessToken string, realmName, userID, clientID string, roles []RoleRepresentation) error {
	_, err := c.post(accessToken, nil, url.Path(clientRoleMappingPath), url.Param("realm", realmName), url.Param("id", userID), url.Param("client", clientID), body.JSON(roles))
	return err
}

// GetClientRoleMappings gets client-level role mappings for the user, and the app.
func (c *Client) GetClientRoleMappings(accessToken string, realmName, userID, clientID string) ([]RoleRepresentation, error) {
	var resp = []RoleRepresentation{}
	var err = c.get(accessToken, &resp, url.Path(clientRoleMappingPath), url.Param("realm", realmName), url.Param("id", userID), url.Param("client", clientID))
	return resp, err
}

// DeleteClientRolesFromUserRoleMapping deletes client-level roles from user role mapping.
func (c *Client) DeleteClientRolesFromUserRoleMapping(accessToken string, realmName, userID, clientID string) error {
	return c.delete(accessToken, url.Path(clientRoleMappingPath), url.Param("realm", realmName), url.Param("id", userID), url.Param("client", clientID))
}

// GetRealmLevelRoleMappings gets realm level role mappings
func (c *Client) GetRealmLevelRoleMappings(accessToken string, realmName, userID string) ([]RoleRepresentation, error) {
	var resp = []RoleRepresentation{}
	var err = c.get(accessToken, &resp, url.Path(realmRoleMappingPath), url.Param("realm", realmName), url.Param("id", userID))
	return resp, err
}
