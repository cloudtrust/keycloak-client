package api

import (
	"errors"

	"github.com/cloudtrust/keycloak-client/v2"
	"gopkg.in/h2non/gentleman.v2/plugins/body"
	"gopkg.in/h2non/gentleman.v2/plugins/query"
	"gopkg.in/h2non/gentleman.v2/plugins/url"
)

const (
	// API Keycloak out-of-the-box
	kcUserPath           = "/auth/admin/realms/:realm/users"
	kcUserCountPath      = kcUserPath + "/count"
	kcUserIDPath         = kcUserPath + "/:id"
	kcUserGroupsPath     = kcUserIDPath + "/groups"
	kcUserGroupIDPath    = kcUserGroupsPath + "/:groupId"
	kcUserFederationPath = kcUserIDPath + "/federated-identity"
	kcShadowUser         = kcUserFederationPath + "/:provider"
	kcProfilePath        = kcUserPath + "/profile"

	// API keycloak-rest-api-extensions admin
	ctAdminRootPath                       = "/auth/realms/:realmReq/api/admin"
	ctAdminExtensionAPIPath               = ctAdminRootPath + "/realms/:realm"
	ctUsersAdminExtensionAPIPath          = ctAdminExtensionAPIPath + "/users"
	ctSendEmailAdminExtensionAPIPath      = ctAdminExtensionAPIPath + "/send-email"
	ctSendEmailUsersAdminExtensionAPIPath = ctUsersAdminExtensionAPIPath + "/:userId/send-email"
	ctGetUserPath                         = ctUsersAdminExtensionAPIPath + "/:id"
	ctExecuteActionsEmailPath             = ctUsersAdminExtensionAPIPath + "/:id/execute-actions-email"
	ctExpiredToUAcceptancePath            = ctAdminRootPath + "/expired-tou-acceptance"
	ctGetSupportInfoPath                  = ctAdminRootPath + "/support-infos"

	// API keycloak-sms
	ctSmsAPI              = "/auth/realms/:realm/smsApi"
	ctSendSmsCode         = ctSmsAPI + "/sendNewCode"
	ctSendSmsConsentCode  = ctSmsAPI + "/users/:userId/consent"
	ctCheckSmsConsentCode = ctSendSmsConsentCode + "/:consent"
	ctSendSMSPath         = ctSmsAPI + "/sendSms"

	// API keycloak-custom-flows Onboarding
	ctSendReminderEmailPath = "/auth/realms/:realm/ctcustom/sendReminderEmail"

	// API keycloak-custom-flows TrustID auth token
	ctGenerateTrustIDAuthToken = "/auth/realms/:realmReq/ctcustom/realms/:realm/users/:userId/generate-trustid-auth-token"
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

	var plugins = append(c.createQueryPlugins(paramKV...), url.Path(ctUsersAdminExtensionAPIPath), url.Param("realmReq", reqRealmName), url.Param("realm", targetRealmName))
	var err = c.forRealm(accessToken, reqRealmName).
		get(accessToken, &resp, plugins...)
	return resp, err
}

// CreateUser creates the user from its UserRepresentation. The username must be unique.
func (c *Client) CreateUser(accessToken string, reqRealmName, targetRealmName string, user keycloak.UserRepresentation, paramKV ...string) (string, error) {
	if len(paramKV)%2 != 0 {
		return "", errors.New(keycloak.MsgErrInvalidParam + "." + keycloak.EvenParams)
	}
	var plugins = append(c.createQueryPlugins(paramKV...), url.Path(ctUsersAdminExtensionAPIPath), url.Param("realmReq", reqRealmName), url.Param("realm", targetRealmName), body.JSON(user))
	return c.forRealm(accessToken, reqRealmName).
		post(accessToken, nil, plugins...)
}

// CountUsers returns the number of users in the realm.
func (c *Client) CountUsers(accessToken string, realmName string) (int, error) {
	var resp = 0
	var err = c.forRealm(accessToken, realmName).
		get(accessToken, &resp, url.Path(kcUserCountPath), url.Param("realm", realmName))
	return resp, err
}

// GetUser gets the represention of the user.
func (c *Client) GetUser(accessToken string, realmName, userID string) (keycloak.UserRepresentation, error) {
	var resp = keycloak.UserRepresentation{}
	var err = c.forRealm(accessToken, realmName).
		get(accessToken, &resp, url.Path(ctGetUserPath), url.Param("realmReq", realmName), url.Param("realm", realmName), url.Param("id", userID))
	return resp, err
}

// GetGroupsOfUser gets the groups of the user.
func (c *Client) GetGroupsOfUser(accessToken string, realmName, userID string) ([]keycloak.GroupRepresentation, error) {
	var resp = []keycloak.GroupRepresentation{}
	var err = c.forRealm(accessToken, realmName).
		get(accessToken, &resp, url.Path(kcUserGroupsPath), url.Param("realm", realmName), url.Param("id", userID))
	return resp, err
}

// AddGroupToUser adds a group to the groups of the user.
func (c *Client) AddGroupToUser(accessToken string, realmName, userID, groupID string) error {
	return c.forRealm(accessToken, realmName).
		put(accessToken, url.Path(kcUserGroupIDPath), url.Param("realm", realmName), url.Param("id", userID), url.Param("groupId", groupID))
}

// DeleteGroupFromUser adds a group to the groups of the user.
func (c *Client) DeleteGroupFromUser(accessToken string, realmName, userID, groupID string) error {
	return c.forRealm(accessToken, realmName).
		delete(accessToken, url.Path(kcUserGroupIDPath), url.Param("realm", realmName), url.Param("id", userID), url.Param("groupId", groupID))
}

// UpdateUser updates the user.
func (c *Client) UpdateUser(accessToken string, realmName, userID string, user keycloak.UserRepresentation) error {
	return c.forRealm(accessToken, realmName).
		put(accessToken, url.Path(kcUserIDPath), url.Param("realm", realmName), url.Param("id", userID), body.JSON(user))
}

// DeleteUser deletes the user.
func (c *Client) DeleteUser(accessToken string, realmName, userID string) error {
	return c.forRealm(accessToken, realmName).
		delete(accessToken, url.Path(kcUserIDPath), url.Param("realm", realmName), url.Param("id", userID))
}

// ExecuteActionsEmail sends an update account email to the user. An email contains a link the user can click to perform a set of required actions.
func (c *Client) ExecuteActionsEmail(accessToken string, reqRealmName string, targetRealmName string, userID string, actions []string, paramKV ...string) error {
	if len(paramKV)%2 != 0 {
		return errors.New(keycloak.MsgErrInvalidParam + "." + keycloak.EvenParams)
	}

	var plugins = append(c.createQueryPlugins(paramKV...), url.Path(ctExecuteActionsEmailPath), url.Param("realmReq", reqRealmName),
		url.Param("realm", targetRealmName), url.Param("id", userID), body.JSON(actions))

	return c.forRealm(accessToken, reqRealmName).
		put(accessToken, plugins...)
}

// SendSmsCode sends a SMS code and return it
func (c *Client) SendSmsCode(accessToken string, realmName string, userID string) (keycloak.SmsCodeRepresentation, error) {
	var paramKV []string
	paramKV = append(paramKV, "userid", userID)
	var plugins = append(c.createQueryPlugins(paramKV...), url.Path(ctSendSmsCode), url.Param("realm", realmName))
	var resp = keycloak.SmsCodeRepresentation{}

	_, err := c.forRealm(accessToken, realmName).
		post(accessToken, &resp, plugins...)

	return resp, err
}

// SendReminderEmail sends a reminder email to a user
func (c *Client) SendReminderEmail(accessToken string, realmName string, userID string, paramKV ...string) error {
	if len(paramKV)%2 != 0 {
		return errors.New(keycloak.MsgErrInvalidParam + "." + keycloak.EvenParams)
	}
	var newParamKV = append(paramKV, "userid", userID)

	var plugins = append(c.createQueryPlugins(newParamKV...), url.Path(ctSendReminderEmailPath), url.Param("realm", realmName))

	_, err := c.forRealm(accessToken, realmName).
		post(accessToken, nil, plugins...)
	return err
}

// GetFederatedIdentities gets the federated identities of a user in the given realm
func (c *Client) GetFederatedIdentities(accessToken string, realmName string, userID string) ([]keycloak.FederatedIdentityRepresentation, error) {
	var res []keycloak.FederatedIdentityRepresentation
	var err = c.forRealm(accessToken, realmName).
		get(accessToken, &res, url.Path(kcUserFederationPath), url.Param("realm", realmName), url.Param("id", userID))
	return res, err
}

// LinkShadowUser links shadow user to a realm in the context of brokering
func (c *Client) LinkShadowUser(accessToken string, reqRealmName string, userID string, provider string, fedIDKC keycloak.FederatedIdentityRepresentation) error {
	_, err := c.forRealm(accessToken, reqRealmName).post(accessToken, nil, url.Path(kcShadowUser), url.Param("realm", reqRealmName), url.Param("id", userID), url.Param("provider", provider), body.JSON(fedIDKC))
	return err
}

// UnlinkShadowUser unlinks shadow user to a realm in the context of brokering
func (c *Client) UnlinkShadowUser(accessToken string, reqRealmName string, userID string, provider string) error {
	return c.forRealm(accessToken, reqRealmName).delete(accessToken, url.Path(kcShadowUser), url.Param("realm", reqRealmName), url.Param("id", userID), url.Param("provider", provider))
}

// SendEmail sends an email to a user
func (c *Client) SendEmail(accessToken string, reqRealmName string, realmName string, emailRep keycloak.EmailRepresentation) error {
	_, err := c.forRealm(accessToken, reqRealmName).
		post(accessToken, nil, url.Path(ctSendEmailAdminExtensionAPIPath), url.Param("realmReq", reqRealmName), url.Param("realm", realmName), body.JSON(emailRep))
	return err
}

// SendEmailToUser sends an email to the user specified by the UserID
func (c *Client) SendEmailToUser(accessToken string, reqRealmName string, realmName string, userID string, emailRep keycloak.EmailRepresentation) error {
	_, err := c.forRealm(accessToken, reqRealmName).
		post(accessToken, nil, url.Path(ctSendEmailUsersAdminExtensionAPIPath), url.Param("realmReq", reqRealmName), url.Param("realm", realmName), url.Param("userId", userID), body.JSON(emailRep))
	return err
}

// SendSMS sends an SMS to a user
func (c *Client) SendSMS(accessToken string, realmName string, smsRep keycloak.SMSRepresentation) error {
	_, err := c.forRealm(accessToken, realmName).
		post(accessToken, nil, url.Path(ctSendSMSPath), url.Param("realm", realmName), body.JSON(smsRep))
	return err
}

// CheckConsentCodeSMS checks a consent code previously sent by SMS to a user
func (c *Client) CheckConsentCodeSMS(accessToken string, realmName string, userID string, consentCode string) error {
	return c.forRealm(accessToken, realmName).
		get(accessToken, nil, url.Path(ctCheckSmsConsentCode), url.Param("realm", realmName), url.Param("userId", userID), url.Param("consent", consentCode))
}

// SendConsentCodeSMS sends an SMS to a user with a consent code
func (c *Client) SendConsentCodeSMS(accessToken string, realmName string, userID string) error {
	_, err := c.forRealm(accessToken, realmName).
		post(accessToken, nil, url.Path(ctSendSmsConsentCode), url.Param("realm", realmName), url.Param("userId", userID))
	return err
}

// GetExpiredTermsOfUseAcceptance gets the list of users created for a
// long time (configured in Keycloak) who declined the terms of use
func (c *Client) GetExpiredTermsOfUseAcceptance(accessToken string) ([]keycloak.DeletableUserRepresentation, error) {
	var deletableUsers []keycloak.DeletableUserRepresentation
	var realmName = "master"
	err := c.forRealm(accessToken, realmName).
		get(accessToken, &deletableUsers, url.Path(ctExpiredToUAcceptancePath), url.Param("realmReq", realmName))
	return deletableUsers, err
}

// GetSupportInfo gets the list of accounts matching a given email address
func (c *Client) GetSupportInfo(accessToken string, email string) ([]keycloak.EmailInfoRepresentation, error) {
	var emailInfos []keycloak.EmailInfoRepresentation
	var realmName = "master"
	err := c.forRealm(accessToken, realmName).
		get(accessToken, &emailInfos, url.Path(ctGetSupportInfoPath), url.Param("realmReq", realmName), query.Add("email", email))
	return emailInfos, err
}

// GenerateTrustIDAuthToken generates a TrustID auth token
func (c *Client) GenerateTrustIDAuthToken(accessToken string, reqRealmName string, realmName string, userID string) (string, error) {
	var token keycloak.TrustIDAuthTokenRepresentation
	err := c.forRealm(accessToken, reqRealmName).
		get(accessToken, &token, url.Path(ctGenerateTrustIDAuthToken), url.Param("realmReq", reqRealmName), url.Param("realm", realmName), url.Param("userId", userID))
	return *token.Token, err
}

// GetUserProfile gets the configuration of attribute management
func (c *Client) GetUserProfile(accessToken string, realmName string) (keycloak.UserProfileRepresentation, error) {
	var profile keycloak.UserProfileRepresentation
	err := c.forRealm(accessToken, realmName).
		get(accessToken, &profile, url.Path(kcProfilePath), url.Param("realm", realmName))
	return profile, err
}
