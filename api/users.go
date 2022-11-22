package api

import (
	"errors"

	"github.com/cloudtrust/keycloak-client/v2"
	"gopkg.in/h2non/gentleman.v2/plugins/body"
	"gopkg.in/h2non/gentleman.v2/plugins/query"
	"gopkg.in/h2non/gentleman.v2/plugins/url"
)

const (
	userPath                            = "/auth/admin/realms/:realm/users"
	adminRootPath                       = "/auth/realms/:realmReq/api/admin"
	adminExtensionAPIPath               = adminRootPath + "/realms/:realm"
	usersAdminExtensionAPIPath          = adminExtensionAPIPath + "/users"
	sendEmailAdminExtensionAPIPath      = adminExtensionAPIPath + "/send-email"
	sendEmailUsersAdminExtensionAPIPath = usersAdminExtensionAPIPath + "/:userId/send-email"
	userCountPath                       = userPath + "/count"
	userIDPath                          = userPath + "/:id"
	userGroupsPath                      = userIDPath + "/groups"
	userGroupIDPath                     = userGroupsPath + "/:groupId"
	executeActionsEmailPath             = usersAdminExtensionAPIPath + "/:id/execute-actions-email"
	sendReminderEmailPath               = "/auth/realms/:realm/onboarding/sendReminderEmail"
	smsAPI                              = "/auth/realms/:realm/smsApi"
	sendSmsCode                         = smsAPI + "/sendNewCode"
	sendSmsConsentCode                  = smsAPI + "/users/:userId/consent"
	checkSmsConsentCode                 = sendSmsConsentCode + "/:consent"
	sendSMSPath                         = smsAPI + "/sendSms"
	userFederationPath                  = userIDPath + "/federated-identity"
	shadowUser                          = userFederationPath + "/:provider"
	expiredToUAcceptancePath            = adminRootPath + "/expired-tou-acceptance"
	getSupportInfoPath                  = adminRootPath + "/support-infos"
	generateTrustIDAuthToken            = "/auth/realms/:realmReq/trustid-auth-token/realms/:realm/users/:userId/generate"
	profilePath                         = userPath + "/profile"
)

// GetUsers returns a list of users, filtered according to the query parameters.
// Parameters: email, first (paging offset, int), firstName, lastName, username,
// max (maximum result size, default = 100),
// search (string contained in username, firstname, lastname or email. by default,
// value is searched with a like -%value%- but you can introduce your own % symbol
// or you can use =value to search an exact value)
func (c *Client) GetUsers(accessToken string, reqRealmName, targetRealmName string, paramKV ...string) (keycloak.UsersPageRepresentation, error) {
	var resp keycloak.UsersPageRepresentation
	if len(paramKV)%2 != 0 {
		return resp, errors.New(keycloak.MsgErrInvalidParam + "." + keycloak.EvenParams)
	}

	var plugins = append(createQueryPlugins(paramKV...), url.Path(usersAdminExtensionAPIPath), url.Param("realmReq", reqRealmName), url.Param("realm", targetRealmName))
	var err = c.get(accessToken, &resp, plugins...)
	return resp, err
}

// CreateUser creates the user from its UserRepresentation. The username must be unique.
func (c *Client) CreateUser(accessToken string, reqRealmName, targetRealmName string, user keycloak.UserRepresentation) (string, error) {
	return c.post(accessToken, nil, url.Path(usersAdminExtensionAPIPath), url.Param("realmReq", reqRealmName), url.Param("realm", targetRealmName), body.JSON(user))
}

// CountUsers returns the number of users in the realm.
func (c *Client) CountUsers(accessToken string, realmName string) (int, error) {
	var resp = 0
	var err = c.get(accessToken, &resp, url.Path(userCountPath), url.Param("realm", realmName))
	return resp, err
}

// GetUser gets the represention of the user.
func (c *Client) GetUser(accessToken string, realmName, userID string) (keycloak.UserRepresentation, error) {
	var resp = keycloak.UserRepresentation{}
	var err = c.get(accessToken, &resp, url.Path(userIDPath), url.Param("realm", realmName), url.Param("id", userID))
	return resp, err
}

// GetGroupsOfUser gets the groups of the user.
func (c *Client) GetGroupsOfUser(accessToken string, realmName, userID string) ([]keycloak.GroupRepresentation, error) {
	var resp = []keycloak.GroupRepresentation{}
	var err = c.get(accessToken, &resp, url.Path(userGroupsPath), url.Param("realm", realmName), url.Param("id", userID))
	return resp, err
}

// AddGroupToUser adds a group to the groups of the user.
func (c *Client) AddGroupToUser(accessToken string, realmName, userID, groupID string) error {
	return c.put(accessToken, url.Path(userGroupIDPath), url.Param("realm", realmName), url.Param("id", userID), url.Param("groupId", groupID))
}

// DeleteGroupFromUser adds a group to the groups of the user.
func (c *Client) DeleteGroupFromUser(accessToken string, realmName, userID, groupID string) error {
	return c.delete(accessToken, url.Path(userGroupIDPath), url.Param("realm", realmName), url.Param("id", userID), url.Param("groupId", groupID))
}

// UpdateUser updates the user.
func (c *Client) UpdateUser(accessToken string, realmName, userID string, user keycloak.UserRepresentation) error {
	return c.put(accessToken, url.Path(userIDPath), url.Param("realm", realmName), url.Param("id", userID), body.JSON(user))
}

// DeleteUser deletes the user.
func (c *Client) DeleteUser(accessToken string, realmName, userID string) error {
	return c.delete(accessToken, url.Path(userIDPath), url.Param("realm", realmName), url.Param("id", userID))
}

// ExecuteActionsEmail sends an update account email to the user. An email contains a link the user can click to perform a set of required actions.
func (c *Client) ExecuteActionsEmail(accessToken string, reqRealmName string, targetRealmName string, userID string, actions []string, paramKV ...string) error {
	if len(paramKV)%2 != 0 {
		return errors.New(keycloak.MsgErrInvalidParam + "." + keycloak.EvenParams)
	}

	var plugins = append(createQueryPlugins(paramKV...), url.Path(executeActionsEmailPath), url.Param("realmReq", reqRealmName),
		url.Param("realm", targetRealmName), url.Param("id", userID), body.JSON(actions))

	return c.put(accessToken, plugins...)
}

// SendSmsCode sends a SMS code and return it
func (c *Client) SendSmsCode(accessToken string, realmName string, userID string) (keycloak.SmsCodeRepresentation, error) {
	var paramKV []string
	paramKV = append(paramKV, "userid", userID)
	var plugins = append(createQueryPlugins(paramKV...), url.Path(sendSmsCode), url.Param("realm", realmName))
	var resp = keycloak.SmsCodeRepresentation{}

	_, err := c.post(accessToken, &resp, plugins...)

	return resp, err
}

// SendReminderEmail sends a reminder email to a user
func (c *Client) SendReminderEmail(accessToken string, realmName string, userID string, paramKV ...string) error {
	if len(paramKV)%2 != 0 {
		return errors.New(keycloak.MsgErrInvalidParam + "." + keycloak.EvenParams)
	}
	var newParamKV = append(paramKV, "userid", userID)

	var plugins = append(createQueryPlugins(newParamKV...), url.Path(sendReminderEmailPath), url.Param("realm", realmName))

	_, err := c.post(accessToken, nil, plugins...)
	return err
}

// GetFederatedIdentities gets the federated identities of a user in the given realm
func (c *Client) GetFederatedIdentities(accessToken string, realmName string, userID string) ([]keycloak.FederatedIdentityRepresentation, error) {
	var res []keycloak.FederatedIdentityRepresentation
	var err = c.get(accessToken, &res, url.Path(userFederationPath), url.Param("realm", realmName), url.Param("id", userID))
	return res, err
}

// LinkShadowUser links shadow user to a realm in the context of brokering
func (c *Client) LinkShadowUser(accessToken string, reqRealmName string, userID string, provider string, fedIDKC keycloak.FederatedIdentityRepresentation) error {
	_, err := c.post(accessToken, nil, url.Path(shadowUser), url.Param("realm", reqRealmName), url.Param("id", userID), url.Param("provider", provider), body.JSON(fedIDKC))
	return err
}

// SendEmail sends an email to a user
func (c *Client) SendEmail(accessToken string, reqRealmName string, realmName string, emailRep keycloak.EmailRepresentation) error {
	_, err := c.post(accessToken, nil, url.Path(sendEmailAdminExtensionAPIPath), url.Param("realmReq", reqRealmName), url.Param("realm", realmName), body.JSON(emailRep))
	return err
}

// SendEmailToUser sends an email to the user specified by the UserID
func (c *Client) SendEmailToUser(accessToken string, reqRealmName string, realmName string, userID string, emailRep keycloak.EmailRepresentation) error {
	_, err := c.post(accessToken, nil, url.Path(sendEmailUsersAdminExtensionAPIPath), url.Param("realmReq", reqRealmName), url.Param("realm", realmName), url.Param("userId", userID), body.JSON(emailRep))
	return err
}

// SendSMS sends an SMS to a user
func (c *Client) SendSMS(accessToken string, realmName string, smsRep keycloak.SMSRepresentation) error {
	_, err := c.post(accessToken, nil, url.Path(sendSMSPath), url.Param("realm", realmName), body.JSON(smsRep))
	return err
}

// CheckConsentCodeSMS checks a consent code previously sent by SMS to a user
func (c *Client) CheckConsentCodeSMS(accessToken string, realmName string, userID string, consentCode string) error {
	return c.get(accessToken, nil, url.Path(checkSmsConsentCode), url.Param("realm", realmName), url.Param("userId", userID), url.Param("consent", consentCode))
}

// SendConsentCodeSMS sends an SMS to a user with a consent code
func (c *Client) SendConsentCodeSMS(accessToken string, realmName string, userID string) error {
	_, err := c.post(accessToken, nil, url.Path(sendSmsConsentCode), url.Param("realm", realmName), url.Param("userId", userID))
	return err
}

// GetExpiredTermsOfUseAcceptance gets the list of users created for a
// long time (configured in Keycloak) who declined the terms of use
func (c *Client) GetExpiredTermsOfUseAcceptance(accessToken string) ([]keycloak.DeletableUserRepresentation, error) {
	var deletableUsers []keycloak.DeletableUserRepresentation
	err := c.get(accessToken, &deletableUsers, url.Path(expiredToUAcceptancePath), url.Param("realmReq", "master"))
	return deletableUsers, err
}

// GetSupportInfo gets the list of accounts matching a given email address
func (c *Client) GetSupportInfo(accessToken string, email string) ([]keycloak.EmailInfoRepresentation, error) {
	var emailInfos []keycloak.EmailInfoRepresentation
	err := c.get(accessToken, &emailInfos, url.Path(getSupportInfoPath), url.Param("realmReq", "master"), query.Add("email", email))
	return emailInfos, err
}

// GenerateTrustIDAuthToken generates a TrustID auth token
func (c *Client) GenerateTrustIDAuthToken(accessToken string, reqRealmName string, realmName string, userID string) (string, error) {
	var token keycloak.TrustIDAuthTokenRepresentation
	err := c.get(accessToken, &token, url.Path(generateTrustIDAuthToken), url.Param("realmReq", reqRealmName), url.Param("realm", realmName), url.Param("userId", userID))
	return *token.Token, err
}

// GetUserProfile gets the configuration of attribute management
func (c *Client) GetUserProfile(accessToken string, realmName string) (keycloak.UserProfileRepresentation, error) {
	var profile keycloak.UserProfileRepresentation
	err := c.get(accessToken, &profile, url.Path(profilePath), url.Param("realm", realmName))
	return profile, err
}
