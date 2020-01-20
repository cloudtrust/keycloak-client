package keycloak

import (
	"gopkg.in/h2non/gentleman.v2/plugins/body"
	"gopkg.in/h2non/gentleman.v2/plugins/url"
)

const (
	groupsPath                          = "/auth/admin/realms/:realm/groups"
	groupByIDPath                       = groupsPath + "/:id"
	groupClientRoleMappingPath          = groupByIDPath + "/role-mappings/clients/:clientId"
	availableGroupClientRoleMappingPath = groupClientRoleMappingPath + "/available"
)

// GetGroups gets all groups for the realm
func (c *Client) GetGroups(accessToken string, realmName string) ([]GroupRepresentation, error) {
	var resp = []GroupRepresentation{}
	var err = c.get(accessToken, &resp, url.Path(groupsPath), url.Param("realm", realmName))
	return resp, err
}

// GetGroup gets a specific group’s representation
func (c *Client) GetGroup(accessToken string, realmName string, groupID string) (GroupRepresentation, error) {
	var resp = GroupRepresentation{}
	var err = c.get(accessToken, &resp, url.Path(groupByIDPath), url.Param("realm", realmName), url.Param("id", groupID))
	return resp, err
}

// CreateGroup creates the group from its GroupRepresentation. The group name must be unique.
func (c *Client) CreateGroup(accessToken string, reqRealmName string, group GroupRepresentation) (string, error) {
	return c.post(accessToken, nil, url.Path(groupsPath), url.Param("realm", reqRealmName), body.JSON(group))
}

// DeleteGroup deletes a specific group’s representation
func (c *Client) DeleteGroup(accessToken string, realmName string, groupID string) error {
	return c.delete(accessToken, url.Path(groupByIDPath), url.Param("realm", realmName), url.Param("id", groupID))
}

// AssignClientRole assigns client roles to a specific group
func (c *Client) AssignClientRole(accessToken string, realmName string, groupID string, clientID string, roles []RoleRepresentation) error {
	_, err := c.post(accessToken, nil, url.Path(groupClientRoleMappingPath), url.Param("realm", realmName), url.Param("id", groupID), url.Param("clientId", clientID), body.JSON(roles))
	return err
}

// RemoveClientRole deletes client roles from a specific group
func (c *Client) RemoveClientRole(accessToken string, realmName string, groupID string, clientID string, roles []RoleRepresentation) error {
	return c.delete(accessToken, url.Path(groupClientRoleMappingPath), url.Param("realm", realmName), url.Param("id", groupID), url.Param("clientId", clientID), body.JSON(roles))
}

// GetGroupClientRoles gets client roles assigned to a specific group
func (c *Client) GetGroupClientRoles(accessToken string, realmName string, groupID string, clientID string) ([]RoleRepresentation, error) {
	var roles = []RoleRepresentation{}
	var err = c.get(accessToken, &roles, url.Path(groupClientRoleMappingPath), url.Param("realm", realmName), url.Param("id", groupID), url.Param("clientId", clientID))
	return roles, err
}

// GetAvailableGroupClientRoles gets client roles available in a specific group
func (c *Client) GetAvailableGroupClientRoles(accessToken string, realmName string, groupID string, clientID string) ([]RoleRepresentation, error) {
	var roles = []RoleRepresentation{}
	var err = c.get(accessToken, &roles, url.Path(availableGroupClientRoleMappingPath), url.Param("realm", realmName), url.Param("id", groupID), url.Param("clientId", clientID))
	return roles, err
}
