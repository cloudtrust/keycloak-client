package keycloak

import (
	"gopkg.in/h2non/gentleman.v2/plugins/body"
	"gopkg.in/h2non/gentleman.v2/plugins/url"
)

const (
	groupsPath                 = "/auth/admin/realms/:realm/groups"
	groupByIDPath              = groupsPath + "/:id"
	groupClientRoleMappingPath = groupByIDPath + "/role-mappings/clients/:clientId"
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

// DeleteGroup delete a specific group’s representation
func (c *Client) DeleteGroup(accessToken string, realmName string, groupID string) error {
	return c.delete(accessToken, url.Path(groupByIDPath), url.Param("realm", realmName), url.Param("id", groupID))
}

// AssignClientRole assign client roles to a specific group
func (c *Client) AssignClientRole(accessToken string, realmName string, groupID string, clientID string, role []RoleRepresentation) error {
	_, err := c.post(accessToken, url.Path(groupClientRoleMappingPath), url.Param("realm", realmName), url.Param("id", groupID), url.Param("clientId", clientID))
	return err
}
