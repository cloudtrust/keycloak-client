package api

import (
	"strconv"

	"github.com/cloudtrust/keycloak-client/v2"
	"gopkg.in/h2non/gentleman.v2/plugins/body"
	"gopkg.in/h2non/gentleman.v2/plugins/url"
)

const (
	resetPasswordPath    = userIDPath + "/reset-password"
	credentialsPath      = userIDPath + "/credentials"
	credentialsTypesPath = realmPath + "/credentialTypes"
	credentialIDPath     = credentialsPath + "/:credentialID"
	labelPath            = credentialIDPath + "/label"
	moveFirstPath        = credentialIDPath + "/moveToFirst"
	moveAfterPath        = credentialIDPath + "/moveAfter/:previousCredentialID"
	logoutPath           = userIDPath + "/logout"
	// Paper card API
	papercardPath               = "/auth/realms/:realm/papercard"
	resetFailuresPath           = papercardPath + "/users/:userId/credentials/:credentialId/resetFailures"
	sendPaperCardsRemindersPath = papercardPath + "/expiryReminders"
)

// ResetPassword resets password of the user.
func (c *Client) ResetPassword(accessToken string, realmName, userID string, cred keycloak.CredentialRepresentation) error {
	return c.put(accessToken, url.Path(resetPasswordPath), url.Param("realm", realmName), url.Param("id", userID), body.JSON(cred))
}

// LogoutAllSessions of the user.
func (c *Client) LogoutAllSessions(accessToken string, realmName, userID string) error {
	var _, err = c.post(accessToken, nil, url.Path(logoutPath), url.Param("realm", realmName), url.Param("id", userID))
	return err
}

// GetCredentials returns the list of credentials of the user
func (c *Client) GetCredentials(accessToken string, realmName string, userID string) ([]keycloak.CredentialRepresentation, error) {
	var resp = []keycloak.CredentialRepresentation{}

	var err = c.get(accessToken, &resp, url.Path(credentialsPath), url.Param("realm", realmName), url.Param("id", userID))
	return resp, err
}

// GetCredentialTypes returns list of credentials types available for the realm
func (c *Client) GetCredentialTypes(accessToken string, realmName string) ([]string, error) {
	var resp = []string{}
	var err = c.get(accessToken, &resp, url.Path(credentialsTypesPath), url.Param("realm", realmName))
	return resp, err
}

// UpdateLabelCredential updates the label of credential
func (c *Client) UpdateLabelCredential(accessToken string, realmName string, userID string, credentialID string, label string) error {
	return c.put(accessToken, url.Path(labelPath), url.Param("realm", realmName), url.Param("id", userID), url.Param("credentialID", credentialID), body.String(label), hdrAcceptJSON, hdrContentTypeTextPlain)
}

// DeleteCredential deletes the credential
func (c *Client) DeleteCredential(accessToken string, realmName string, userID string, credentialID string) error {
	return c.delete(accessToken, url.Path(credentialIDPath), url.Param("realm", realmName), url.Param("id", userID), url.Param("credentialID", credentialID))
}

// MoveToFirst moves the credential at the top of the list
func (c *Client) MoveToFirst(accessToken string, realmName string, userID string, credentialID string) error {
	_, err := c.post(accessToken, url.Path(moveFirstPath), url.Param("realm", realmName), url.Param("id", userID), url.Param("credentialID", credentialID))
	return err
}

// MoveAfter moves the credential after the specified one into the list
func (c *Client) MoveAfter(accessToken string, realmName string, userID string, credentialID string, previousCredentialID string) error {
	_, err := c.post(accessToken, url.Path(moveAfterPath), url.Param("realm", realmName), url.Param("id", userID), url.Param("credentialID", credentialID), url.Param("previousCredentialID", previousCredentialID))
	return err
}

// UpdatePassword updates the user's password
// Parameters: realm, currentPassword, newPassword, confirmPassword
func (c *Client) UpdatePassword(accessToken, realm, currentPassword, newPassword, confirmPassword string) (string, error) {
	var m = map[string]string{"currentPassword": currentPassword, "newPassword": newPassword, "confirmation": confirmPassword}
	return c.post(accessToken, nil, url.Path(accountPasswordPath), url.Param("realm", realm), body.JSON(m))
}

// ResetPapercardFailures reset failures information in a paper card credential
func (c *Client) ResetPapercardFailures(accessToken, realmName, userID, credentialID string) error {
	return c.put(accessToken, url.Path(resetFailuresPath), url.Param("realm", realmName), url.Param("userId", userID), url.Param("credentialId", credentialID))
}

// RemindersResponse struct
type RemindersResponse struct {
	Partial bool `json:"partial"`
}

// SendPaperCardsReminders sends reminders to users of paper cards which will soon be expired
func (c *Client) SendPaperCardsReminders(accessToken, realmName string, firstReminderDays, nextReminderDays, maxCount int) (bool, error) {
	var paramKV = []string{
		"firstReminderDays", strconv.Itoa(firstReminderDays),
		"nextReminderDays", strconv.Itoa(nextReminderDays),
		"maxCount", strconv.Itoa(maxCount),
	}

	var resp RemindersResponse
	var plugins = append(createQueryPlugins(paramKV...), url.Path(sendPaperCardsRemindersPath), url.Param("realm", realmName))
	var _, err = c.post(accessToken, &resp, plugins...)
	if err != nil {
		return false, err
	}
	return resp.Partial, err
}
