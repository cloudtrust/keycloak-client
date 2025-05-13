package api

import (
	"github.com/cloudtrust/keycloak-client/v2"
	"gopkg.in/h2non/gentleman.v2/plugin"
	"gopkg.in/h2non/gentleman.v2/plugins/body"
	"gopkg.in/h2non/gentleman.v2/plugins/headers"
	"gopkg.in/h2non/gentleman.v2/plugins/query"
	"gopkg.in/h2non/gentleman.v2/plugins/url"
)

const (
	// API keycloak-rest-extensions account
	ctAccountExtensionAPIPath            = "/auth/realms/master/api/account/realms/:realm"
	ctAccountExecuteActionsEmail         = ctAccountExtensionAPIPath + "/execute-actions-email"
	ctAccountSendEmail                   = ctAccountExtensionAPIPath + "/send-email"
	ctAccountCredentialsPath             = ctAccountExtensionAPIPath + "/credentials"
	ctAccountPasswordPath                = ctAccountCredentialsPath + "/password"
	ctAccountCredentialsRegistratorsPath = ctAccountCredentialsPath + "/registrators"
	ctAccountCredentialIDPath            = ctAccountCredentialsPath + "/:credentialID"
	ctAccountCredentialLabelPath         = ctAccountCredentialIDPath + "/label"
	ctAccountMoveFirstPath               = ctAccountCredentialIDPath + "/moveToFirst"
	ctAccountMoveAfterPath               = ctAccountCredentialIDPath + "/moveAfter/:previousCredentialID"
)

var (
	hdrAcceptJSON           = headers.Set("Accept", "application/json")
	hdrContentTypeTextPlain = headers.Set("Content-Type", "text/plain")
)

// GetCredentials returns the list of credentials of the user
func (c *AccountClient) GetCredentials(accessToken string, realmName string) ([]keycloak.CredentialRepresentation, error) {
	var resp = []keycloak.CredentialRepresentation{}
	var err = c.client.get(accessToken, &resp, url.Path(ctAccountCredentialsPath), url.Param("realm", realmName), hdrAcceptJSON)
	return resp, err
}

// GetCredentialRegistrators returns list of credentials types available for the user
func (c *AccountClient) GetCredentialRegistrators(accessToken string, realmName string) ([]string, error) {
	var resp = []string{}
	var err = c.client.get(accessToken, &resp, url.Path(ctAccountCredentialsRegistratorsPath), url.Param("realm", realmName), hdrAcceptJSON)
	return resp, err
}

// UpdateLabelCredential updates the label of credential
func (c *AccountClient) UpdateLabelCredential(accessToken string, realmName string, credentialID string, label string) error {
	return c.client.put(accessToken, url.Path(ctAccountCredentialLabelPath), url.Param("realm", realmName), url.Param("credentialID", credentialID), body.String(label), hdrAcceptJSON, hdrContentTypeTextPlain)
}

// DeleteCredential deletes the credential
func (c *AccountClient) DeleteCredential(accessToken string, realmName string, credentialID string) error {
	return c.client.delete(accessToken, url.Path(ctAccountCredentialIDPath), url.Param("realm", realmName), url.Param("credentialID", credentialID), hdrAcceptJSON)
}

// MoveToFirst moves the credential at the top of the list
func (c *AccountClient) MoveToFirst(accessToken string, realmName string, credentialID string) error {
	_, err := c.client.post(accessToken, nil, url.Path(ctAccountMoveFirstPath), url.Param("realm", realmName), url.Param("credentialID", credentialID), hdrAcceptJSON)
	return err
}

// MoveAfter moves the credential after the specified one into the list
func (c *AccountClient) MoveAfter(accessToken string, realmName string, credentialID string, previousCredentialID string) error {
	_, err := c.client.post(accessToken, nil, url.Path(ctAccountMoveAfterPath), url.Param("realm", realmName), url.Param("credentialID", credentialID), url.Param("previousCredentialID", previousCredentialID), hdrAcceptJSON)
	return err
}

// UpdatePassword updates the user's password
// Parameters: realm, currentPassword, newPassword, confirmPassword
func (c *AccountClient) UpdatePassword(accessToken, realm, currentPassword, newPassword, confirmPassword string) (string, error) {
	var m = map[string]string{"currentPassword": currentPassword, "newPassword": newPassword, "confirmation": confirmPassword}
	return c.client.post(accessToken, nil, url.Path(ctAccountPasswordPath), url.Param("realm", realm), body.JSON(m))
}

// GetAccount provides the user's information
func (c *AccountClient) GetAccount(accessToken string, realm string) (keycloak.UserRepresentation, error) {
	var resp = keycloak.UserRepresentation{}
	var err = c.client.get(accessToken, &resp, url.Path(ctAccountExtensionAPIPath), url.Param("realm", realm), hdrAcceptJSON)
	return resp, err
}

// UpdateAccount updates the user's information
func (c *AccountClient) UpdateAccount(accessToken string, realm string, user keycloak.UserRepresentation) error {
	_, err := c.client.post(accessToken, nil, url.Path(ctAccountExtensionAPIPath), url.Param("realm", realm), body.JSON(user))
	return err
}

// DeleteAccount deletes current user
func (c *AccountClient) DeleteAccount(accessToken string, realmName string) error {
	return c.client.delete(accessToken, url.Path(ctAccountExtensionAPIPath), url.Param("realm", realmName), hdrAcceptJSON)
}

// ExecuteActionsEmail sends an email with required actions to the user
func (c *AccountClient) ExecuteActionsEmail(accessToken string, realmName string, actions []string) error {
	return c.client.put(accessToken, url.Path(ctAccountExecuteActionsEmail), url.Param("realm", realmName), body.JSON(actions))
}

// SendEmail sends an email
func (c *AccountClient) SendEmail(accessToken, realmName, template, subject string, recipient *string, attributes map[string]string) error {
	var plugins []plugin.Plugin
	plugins = append(plugins, url.Path(ctAccountSendEmail), url.Param("realm", realmName))
	plugins = append(plugins, query.Add("template", template), query.Add("subject", subject))
	if recipient != nil && len(*recipient) >= 0 {
		plugins = append(plugins, query.Add("recipient", *recipient))
	}
	plugins = append(plugins, body.JSON(attributes))
	_, err := c.client.post(accessToken, nil, plugins...)
	return err
}
