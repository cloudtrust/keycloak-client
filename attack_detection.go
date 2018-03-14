package keycloak

import (
	"gopkg.in/h2non/gentleman.v2/plugins/url"
)

const (
	attackDetectionPath   = "/auth/admin/realms/:realm/attack-detection/brute-force/users"
	attackDetectionIDPath = attackDetectionPath + "/:id"
)

// ClearAllLoginFailures clears any user login failures for all users. This can release temporary disabled users.
func (c *Client) ClearAllLoginFailures(realmName string) error {
	return c.delete(url.Path(attackDetectionPath), url.Param("realm", realmName))
}

// GetAttackDetectionStatus gets the status of a username in brute force detection.
func (c *Client) GetAttackDetectionStatus(realmName, userID string) (map[string]interface{}, error) {
	var resp = map[string]interface{}{}
	var err = c.get(&resp, url.Path(attackDetectionIDPath), url.Param("realm", realmName), url.Param("id", userID))
	return resp, err
}

// ClearUserLoginFailures clear any user login failures for the user. This can release temporary disabled user.
func (c *Client) ClearUserLoginFailures(realmName, userID string) error {
	return c.delete(url.Path(attackDetectionIDPath), url.Param("realm", realmName), url.Param("id", userID))
}
