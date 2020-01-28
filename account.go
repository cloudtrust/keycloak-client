package keycloak

import (
	"gopkg.in/h2non/gentleman.v2/plugins/body"
	"gopkg.in/h2non/gentleman.v2/plugins/headers"
	"gopkg.in/h2non/gentleman.v2/plugins/url"
)

const (
	accountPath                        = "/auth/realms/:realm/account"
	accountExtensionAPIPath            = "/auth/realms/master/api/account/realms/:realm"
	accountPasswordPath                = accountExtensionAPIPath + "/credentials/password"
	accountCredentialsPath             = accountExtensionAPIPath + "/credentials"
	accountCredentialsRegistratorsPath = accountCredentialsPath + "/registrators"
	accountCredentialIDPath            = accountCredentialsPath + "/:credentialID"
	accountCredentialLabelPath         = accountCredentialIDPath + "/label"
	accountMoveFirstPath               = accountCredentialIDPath + "/moveToFirst"
	accountMoveAfterPath               = accountCredentialIDPath + "/moveAfter/:previousCredentialID"
)

// GetCredentials returns the list of credentials of the user
func (c *AccountClient) GetCredentials(accessToken string, realmName string) ([]CredentialRepresentation, error) {
	var resp = []CredentialRepresentation{}
	var err = c.client.get(accessToken, &resp, url.Path(accountCredentialsPath), url.Param("realm", realmName), headers.Set("Accept", "application/json"))
	return resp, err
}

// GetCredentialRegistrators returns list of credentials types available for the user
func (c *AccountClient) GetCredentialRegistrators(accessToken string, realmName string) ([]string, error) {
	var resp = []string{}
	var err = c.client.get(accessToken, &resp, url.Path(accountCredentialsRegistratorsPath), url.Param("realm", realmName), headers.Set("Accept", "application/json"))
	return resp, err
}

// UpdateLabelCredential updates the label of credential
func (c *AccountClient) UpdateLabelCredential(accessToken string, realmName string, credentialID string, label string) error {
	return c.client.put(accessToken, url.Path(accountCredentialLabelPath), url.Param("realm", realmName), url.Param("credentialID", credentialID), body.String(label), headers.Set("Accept", "application/json"), headers.Set("Content-Type", "text/plain"))
}

// DeleteCredential deletes the credential
func (c *AccountClient) DeleteCredential(accessToken string, realmName string, credentialID string) error {
	return c.client.delete(accessToken, url.Path(accountCredentialIDPath), url.Param("realm", realmName), url.Param("credentialID", credentialID), headers.Set("Accept", "application/json"))
}

// MoveToFirst moves the credential at the top of the list
func (c *AccountClient) MoveToFirst(accessToken string, realmName string, credentialID string) error {
	_, err := c.client.post(accessToken, nil, url.Path(accountMoveFirstPath), url.Param("realm", realmName), url.Param("credentialID", credentialID), headers.Set("Accept", "application/json"))
	return err
}

// MoveAfter moves the credential after the specified one into the list
func (c *AccountClient) MoveAfter(accessToken string, realmName string, credentialID string, previousCredentialID string) error {
	_, err := c.client.post(accessToken, nil, url.Path(accountMoveAfterPath), url.Param("realm", realmName), url.Param("credentialID", credentialID), url.Param("previousCredentialID", previousCredentialID), headers.Set("Accept", "application/json"))
	return err
}

// UpdatePassword updates the user's password
// Parameters: realm, currentPassword, newPassword, confirmPassword
func (c *AccountClient) UpdatePassword(accessToken, realm, currentPassword, newPassword, confirmPassword string) (string, error) {
	var m = map[string]string{"currentPassword": currentPassword, "newPassword": newPassword, "confirmation": confirmPassword}
	return c.client.post(accessToken, nil, url.Path(accountPasswordPath), url.Param("realm", realm), body.JSON(m))
}

// GetAccount provides the user's information
func (c *AccountClient) GetAccount(accessToken string, realm string) (UserRepresentation, error) {
	var resp = UserRepresentation{}
	var err = c.client.get(accessToken, &resp, url.Path(accountExtensionAPIPath), url.Param("realm", realm), headers.Set("Accept", "application/json"))
	return resp, err
}

// UpdateAccount updates the user's information
func (c *AccountClient) UpdateAccount(accessToken string, realm string, user UserRepresentation) error {
	_, err := c.client.post(accessToken, nil, url.Path(accountExtensionAPIPath), url.Param("realm", realm), body.JSON(user))
	return err
}

// DeleteAccount delete current user
func (c *AccountClient) DeleteAccount(accessToken string, realmName string) error {
	return c.client.delete(accessToken, url.Path(accountExtensionAPIPath), url.Param("realm", realmName), headers.Set("Accept", "application/json"))
}
