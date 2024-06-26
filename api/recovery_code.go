package api

import (
	"github.com/cloudtrust/keycloak-client/v2"
	"gopkg.in/h2non/gentleman.v2/plugins/query"
	"gopkg.in/h2non/gentleman.v2/plugins/url"
)

const (
	recoveryCodePath   = "/auth/realms/:realm/recovery-code"
	activationCodePath = "/auth/realms/:realm/activation-code"
)

// CreateRecoveryCode creates a new recovery code authenticator and returns the code.
func (c *Client) CreateRecoveryCode(accessToken string, realmName string, userID string) (keycloak.RecoveryCodeRepresentation, error) {
	var resp = keycloak.RecoveryCodeRepresentation{}

	_, err := c.post(accessToken, &resp, query.Add("userId", userID), url.Path(recoveryCodePath), url.Param("realm", realmName))
	return resp, err
}

// CreateActivationCode creates a new activation code authenticator and returns the code.
func (c *Client) CreateActivationCode(accessToken string, realmName string, userID string) (keycloak.ActivationCodeRepresentation, error) {
	var resp = keycloak.ActivationCodeRepresentation{}

	_, err := c.post(accessToken, &resp, query.Add("userId", userID), url.Path(activationCodePath), url.Param("realm", realmName))
	return resp, err
}
