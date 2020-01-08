package keycloak

import "gopkg.in/h2non/gentleman.v2/plugins/url"
import "gopkg.in/h2non/gentleman.v2/plugins/query"

const (
	recoveryCodePath = "/auth/realms/:realm/recovery-code"
)

type RecoveryCodeRepresentation struct {
	Code *string `json:"code,omitempty"`
}

// CreateRecoveryCode creates a new recovery code authenticator and returns the code.
func (c *Client) CreateRecoveryCode(accessToken string, realmName string, userID string) (RecoveryCodeRepresentation, error) {
	var resp = RecoveryCodeRepresentation{}

	_, err := c.post(accessToken, &resp, query.Add("userId", userID), url.Path(recoveryCodePath), url.Param("realm", realmName))
	return resp, err
}
