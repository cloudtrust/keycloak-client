package api

import (
	"github.com/cloudtrust/keycloak-client/v2"
	"gopkg.in/h2non/gentleman.v2/plugins/body"
	"gopkg.in/h2non/gentleman.v2/plugins/url"
)

const (
	// API Keycloak out-of-the-box
	kcGroupsPath                          = "/auth/admin/realms/:realm/groups"
	kcGroupByIDPath                       = kcGroupsPath + "/:id"
	kcGroupClientRoleMappingPath          = kcGroupByIDPath + "/role-mappings/clients/:clientId"
	kcAvailableGroupClientRoleMappingPath = kcGroupClientRoleMappingPath + "/available"
)

// GetGroups gets all groups for the realm
func (c *Client) GetGroups(accessToken string, realmName string) ([]keycloak.GroupRepresentation, error) {
	var resp = []keycloak.GroupRepresentation{}
	var err = c.forRealm(realmName).
		get(accessToken, &resp, url.Path(kcGroupsPath), url.Param("realm", realmName))
	return resp, err
}

// GetGroup gets a specific group’s representation
func (c *Client) GetGroup(accessToken string, realmName string, groupID string) (keycloak.GroupRepresentation, error) {
	var resp = keycloak.GroupRepresentation{}
	var err = c.forRealm(realmName).
		get(accessToken, &resp, url.Path(kcGroupByIDPath), url.Param("realm", realmName), url.Param("id", groupID))
	return resp, err
}

// CreateGroup creates the group from its GroupRepresentation. The group name must be unique.
func (c *Client) CreateGroup(accessToken string, reqRealmName string, group keycloak.GroupRepresentation) (string, error) {
	return c.forRealm(reqRealmName).
		post(accessToken, nil, url.Path(kcGroupsPath), url.Param("realm", reqRealmName), body.JSON(group))
}

// DeleteGroup deletes a specific group’s representation
func (c *Client) DeleteGroup(accessToken string, realmName string, groupID string) error {
	return c.forRealm(realmName).
		delete(accessToken, url.Path(kcGroupByIDPath), url.Param("realm", realmName), url.Param("id", groupID))
}

// AssignClientRole assigns client roles to a specific group
func (c *Client) AssignClientRole(accessToken string, realmName string, groupID string, clientID string, roles []keycloak.RoleRepresentation) error {
	_, err := c.forRealm(realmName).
		post(accessToken, nil, url.Path(kcGroupClientRoleMappingPath), url.Param("realm", realmName), url.Param("id", groupID), url.Param("clientId", clientID), body.JSON(roles))
	return err
}

// RemoveClientRole deletes client roles from a specific group
func (c *Client) RemoveClientRole(accessToken string, realmName string, groupID string, clientID string, roles []keycloak.RoleRepresentation) error {
	return c.forRealm(realmName).
		delete(accessToken, url.Path(kcGroupClientRoleMappingPath), url.Param("realm", realmName), url.Param("id", groupID), url.Param("clientId", clientID), body.JSON(roles))
}

// GetGroupClientRoles gets client roles assigned to a specific group
func (c *Client) GetGroupClientRoles(accessToken string, realmName string, groupID string, clientID string) ([]keycloak.RoleRepresentation, error) {
	var roles = []keycloak.RoleRepresentation{}
	var err = c.forRealm(realmName).
		get(accessToken, &roles, url.Path(kcGroupClientRoleMappingPath), url.Param("realm", realmName), url.Param("id", groupID), url.Param("clientId", clientID))
	return roles, err
}

// GetAvailableGroupClientRoles gets client roles available in a specific group
func (c *Client) GetAvailableGroupClientRoles(accessToken string, realmName string, groupID string, clientID string) ([]keycloak.RoleRepresentation, error) {
	var roles = []keycloak.RoleRepresentation{}
	var err = c.forRealm(realmName).
		get(accessToken, &roles, url.Path(kcAvailableGroupClientRoleMappingPath), url.Param("realm", realmName), url.Param("id", groupID), url.Param("clientId", clientID))
	return roles, err
}
