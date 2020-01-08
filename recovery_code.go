package keycloak

import "gopkg.in/h2non/gentleman.v2/plugins/url"

const (
	recoveryCodePath = "/auth/realms/:realm/recovery-code"
)

type RecoveryCodeRepresentation struct {
	Code *string `json:"code,omitempty"`
}

// CreateRecoveryCode creates a new recovery code authenticator and returns the code.
func (c *Client) CreateRecoveryCode(accessToken string, realmName string, userID string) (RecoveryCodeRepresentation, error) {
	var resp = RecoveryCodeRepresentation{}

	var plugins = append(createQueryPlugins("userId", userID), url.Path(recoveryCodePath), url.Param("realm", realmName))
	_, err := c.post(accessToken, &resp, plugins...)
	return resp, err
}
