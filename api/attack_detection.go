package api

import (
	"gopkg.in/h2non/gentleman.v2/plugins/url"
)

const (
	// API Keycloak out-of-the-box
	kcAttackDetectionPath   = "/auth/admin/realms/:realm/attack-detection/brute-force/users"
	kcAttackDetectionIDPath = kcAttackDetectionPath + "/:id"
)

// ClearAllLoginFailures clears any user login failures for all users. This can release temporary disabled users.
func (c *Client) ClearAllLoginFailures(accessToken string, realmName string) error {
	return c.forRealm(accessToken, realmName).delete(accessToken, url.Path(kcAttackDetectionPath), url.Param("realm", realmName))
}

// GetAttackDetectionStatus gets the status of a username in brute force detection.
func (c *Client) GetAttackDetectionStatus(accessToken string, realmName, userID string) (map[string]any, error) {
	var resp = map[string]any{}
	var err = c.forRealm(accessToken, realmName).get(accessToken, &resp, url.Path(kcAttackDetectionIDPath), url.Param("realm", realmName), url.Param("id", userID))
	return resp, err
}

// ClearUserLoginFailures clear any user login failures for the user. This can release temporary disabled user.
func (c *Client) ClearUserLoginFailures(accessToken string, realmName, userID string) error {
	return c.forRealm(accessToken, realmName).delete(accessToken, url.Path(kcAttackDetectionIDPath), url.Param("realm", realmName), url.Param("id", userID))
}
