package api

import (
	"github.com/cloudtrust/keycloak-client/v2"
	"gopkg.in/h2non/gentleman.v2/plugins/query"
	"gopkg.in/h2non/gentleman.v2/plugins/url"
)

const (
	// API keycloak-custom-flows
	ctRecoveryCodePath   = "/auth/realms/:realm/ctcustom/recovery-code"
	ctActivationCodePath = "/auth/realms/:realm/ctcustom/activation-code"
)

// CreateRecoveryCode creates a new recovery code authenticator and returns the code.
func (c *Client) CreateRecoveryCode(accessToken string, realmName string, userID string) (keycloak.RecoveryCodeRepresentation, error) {
	var resp = keycloak.RecoveryCodeRepresentation{}

	_, err := c.forRealm(realmName).
		post(accessToken, &resp, query.Add("userId", userID), url.Path(ctRecoveryCodePath), url.Param("realm", realmName))
	return resp, err
}

// CreateActivationCode creates a new activation code authenticator and returns the code.
func (c *Client) CreateActivationCode(accessToken string, realmName string, userID string) (keycloak.ActivationCodeRepresentation, error) {
	var resp = keycloak.ActivationCodeRepresentation{}

	_, err := c.forRealm(realmName).
		post(accessToken, &resp, query.Add("userId", userID), url.Path(ctActivationCodePath), url.Param("realm", realmName))
	return resp, err
}
