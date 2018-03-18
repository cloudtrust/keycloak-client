package keycloak

import (
	"gopkg.in/h2non/gentleman.v2/plugins/body"
	"gopkg.in/h2non/gentleman.v2/plugins/url"
)

const (
	clientRoleMappingPath = "/auth/admin/realms/:realm/groups/:id/role-mappings/clients/:client"
)

// CreateClientsRoleMapping add client-level roles to the user role mapping.
func (c *Client) CreateClientsRoleMapping(realmName, groupID, clientID string, roles []RoleRepresentation) error {
	return c.post(url.Path(clientRoleMappingPath), url.Param("realm", realmName), url.Param("id", groupID), url.Param("client", clientID), body.JSON(roles))
}

// GetClientsRoleMapping gets client-level role mappings for the user, and the app.
func (c *Client) GetClientsRoleMapping(realmName, groupID, clientID string) ([]RoleRepresentation, error) {
	var resp = []RoleRepresentation{}
	var err = c.get(&resp, url.Path(clientRoleMappingPath), url.Param("realm", realmName), url.Param("id", groupID), url.Param("client", clientID))
	return resp, err
}

// DeleteClientsRoleMapping deletes client-level roles from user role mapping.
func (c *Client) DeleteClientsRoleMapping(realmName, groupID, clientID string) error {
	return c.delete(url.Path(clientRoleMappingPath), url.Param("realm", realmName), url.Param("id", groupID), url.Param("client", clientID))
}
