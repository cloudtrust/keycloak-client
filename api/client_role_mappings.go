package api

import (
	"github.com/cloudtrust/keycloak-client/v2"
	"gopkg.in/h2non/gentleman.v2/plugins/body"
	"gopkg.in/h2non/gentleman.v2/plugins/url"
)

const (
	// API Keycloak out-of-the-box
	kcClientRoleMappingPath = "/auth/admin/realms/:realm/users/:id/role-mappings/clients/:client"
	kcRealmRoleMappingPath  = "/auth/admin/realms/:realm/users/:id/role-mappings/realm"
)

// AddClientRolesToUserRoleMapping add client-level roles to the user role mapping.
func (c *Client) AddClientRolesToUserRoleMapping(accessToken string, realmName, userID, clientID string, roles []keycloak.RoleRepresentation) error {
	_, err := c.forRealm(accessToken, realmName).
		post(accessToken, nil, url.Path(kcClientRoleMappingPath), url.Param("realm", realmName), url.Param("id", userID), url.Param("client", clientID), body.JSON(roles))
	return err
}

// GetClientRoleMappings gets client-level role mappings for the user, and the app.
func (c *Client) GetClientRoleMappings(accessToken string, realmName, userID, clientID string) ([]keycloak.RoleRepresentation, error) {
	var resp = []keycloak.RoleRepresentation{}
	var err = c.forRealm(accessToken, realmName).
		get(accessToken, &resp, url.Path(kcClientRoleMappingPath), url.Param("realm", realmName), url.Param("id", userID), url.Param("client", clientID))
	return resp, err
}

// DeleteClientRolesFromUserRoleMapping deletes client-level roles from user role mapping.
func (c *Client) DeleteClientRolesFromUserRoleMapping(accessToken string, realmName, userID, clientID string, roles []keycloak.RoleRepresentation) error {
	return c.forRealm(accessToken, realmName).
		delete(accessToken, url.Path(kcClientRoleMappingPath), url.Param("realm", realmName), url.Param("id", userID), url.Param("client", clientID), body.JSON(roles))
}

// GetRealmLevelRoleMappings gets realm level role mappings
func (c *Client) GetRealmLevelRoleMappings(accessToken string, realmName, userID string) ([]keycloak.RoleRepresentation, error) {
	var resp = []keycloak.RoleRepresentation{}
	var err = c.forRealm(accessToken, realmName).
		get(accessToken, &resp, url.Path(kcRealmRoleMappingPath), url.Param("realm", realmName), url.Param("id", userID))
	return resp, err
}

// AddRealmLevelRoleMappings gets realm level role mappings
func (c *Client) AddRealmLevelRoleMappings(accessToken string, realmName, userID string, roles []keycloak.RoleRepresentation) error {
	var _, err = c.forRealm(accessToken, realmName).
		post(accessToken, nil, url.Path(kcRealmRoleMappingPath), url.Param("realm", realmName), url.Param("id", userID), body.JSON(roles))
	return err
}

// DeleteRealmLevelRoleMappings gets realm level role mappings
func (c *Client) DeleteRealmLevelRoleMappings(accessToken string, realmName, userID string, roles []keycloak.RoleRepresentation) error {
	return c.forRealm(accessToken, realmName).
		delete(accessToken, url.Path(kcRealmRoleMappingPath), url.Param("realm", realmName), url.Param("id", userID), body.JSON(roles))
}
