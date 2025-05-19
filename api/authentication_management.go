package api

import (
	"github.com/cloudtrust/keycloak-client/v2"
	"gopkg.in/h2non/gentleman.v2/plugins/body"
	"gopkg.in/h2non/gentleman.v2/plugins/url"
)

const (
	// API Keycloak out-of-the-box
	kcAuthenticationManagementPath      = "/auth/admin/realms/:realm/authentication"
	kcAuthenticationConfigPath          = kcAuthenticationManagementPath + "/config/:id"
	kcAuthenticationRequiredActionsPath = kcAuthenticationManagementPath + "/required-actions"
	kcAuthenticationRequiredActionPath  = kcAuthenticationRequiredActionsPath + "/:alias"
)

// GetAuthenticatorProviders returns a list of authenticator providers.
func (c *Client) GetAuthenticatorProviders(accessToken string, realmName string) ([]map[string]any, error) {
	var resp = []map[string]any{}
	var err = c.forRealm(realmName).
		get(accessToken, &resp, url.Path(kcAuthenticationManagementPath+"/authenticator-providers"), url.Param("realm", realmName))
	return resp, err
}

// GetClientAuthenticatorProviders returns a list of client authenticator providers.
func (c *Client) GetClientAuthenticatorProviders(accessToken string, realmName string) ([]map[string]any, error) {
	var resp = []map[string]any{}
	var err = c.forRealm(realmName).
		get(accessToken, &resp, url.Path(kcAuthenticationManagementPath+"/client-authenticator-providers"), url.Param("realm", realmName))
	return resp, err
}

// GetAuthenticatorProviderConfig returns the authenticator provider’s configuration description.
func (c *Client) GetAuthenticatorProviderConfig(accessToken string, realmName, providerID string) (keycloak.AuthenticatorConfigInfoRepresentation, error) {
	var resp = keycloak.AuthenticatorConfigInfoRepresentation{}
	var err = c.forRealm(realmName).
		get(accessToken, &resp, url.Path(kcAuthenticationManagementPath+"/config-description/:providerID"), url.Param("realm", realmName), url.Param("providerID", providerID))
	return resp, err
}

// GetAuthenticatorConfig returns the authenticator configuration.
func (c *Client) GetAuthenticatorConfig(accessToken string, realmName, configID string) (keycloak.AuthenticatorConfigRepresentation, error) {
	var resp = keycloak.AuthenticatorConfigRepresentation{}
	var err = c.forRealm(realmName).
		get(accessToken, &resp, url.Path(kcAuthenticationConfigPath), url.Param("realm", realmName), url.Param("id", configID))
	return resp, err
}

// UpdateAuthenticatorConfig updates the authenticator configuration.
func (c *Client) UpdateAuthenticatorConfig(accessToken string, realmName, configID string, config keycloak.AuthenticatorConfigRepresentation) error {
	return c.forRealm(realmName).
		put(accessToken, url.Path(kcAuthenticationConfigPath), url.Param("realm", realmName), url.Param("id", configID), body.JSON(config))
}

// DeleteAuthenticatorConfig deletes the authenticator configuration.
func (c *Client) DeleteAuthenticatorConfig(accessToken string, realmName, configID string) error {
	return c.forRealm(realmName).
		delete(accessToken, url.Path(kcAuthenticationConfigPath), url.Param("realm", realmName), url.Param("id", configID))
}

// CreateAuthenticationExecution add new authentication execution
func (c *Client) CreateAuthenticationExecution(accessToken string, realmName string, authExec keycloak.AuthenticationExecutionRepresentation) (string, error) {
	return c.forRealm(realmName).
		post(accessToken, nil, url.Path(kcAuthenticationManagementPath+"/executions"), url.Param("realm", realmName), body.JSON(authExec))
}

// DeleteAuthenticationExecution deletes the execution.
func (c *Client) DeleteAuthenticationExecution(accessToken string, realmName, executionID string) error {
	return c.forRealm(realmName).
		delete(accessToken, url.Path(kcAuthenticationManagementPath+"/executions/:id"), url.Param("realm", realmName), url.Param("id", executionID))
}

// UpdateAuthenticationExecution update execution with new configuration.
func (c *Client) UpdateAuthenticationExecution(accessToken string, realmName, executionID string, authConfig keycloak.AuthenticatorConfigRepresentation) error {
	_, err := c.forRealm(realmName).
		post(accessToken, nil, url.Path(kcAuthenticationManagementPath+"/executions/:id/config"), url.Param("realm", realmName), url.Param("id", executionID), body.JSON(authConfig))
	return err
}

// LowerExecutionPriority lowers the execution’s priority.
func (c *Client) LowerExecutionPriority(accessToken string, realmName, executionID string) error {
	_, err := c.forRealm(realmName).
		post(accessToken, nil, url.Path(kcAuthenticationManagementPath+"/executions/:id/lower-priority"), url.Param("realm", realmName), url.Param("id", executionID))
	return err
}

// RaiseExecutionPriority raise the execution’s priority.
func (c *Client) RaiseExecutionPriority(accessToken string, realmName, executionID string) error {
	_, err := c.forRealm(realmName).
		post(accessToken, nil, url.Path(kcAuthenticationManagementPath+"/executions/:id/raise-priority"), url.Param("realm", realmName), url.Param("id", executionID))
	return err
}

// CreateAuthenticationFlow creates a new authentication flow.
func (c *Client) CreateAuthenticationFlow(accessToken string, realmName string, authFlow keycloak.AuthenticationFlowRepresentation) error {
	_, err := c.forRealm(realmName).
		post(accessToken, nil, url.Path(kcAuthenticationManagementPath+"/flows"), url.Param("realm", realmName), body.JSON(authFlow))
	return err
}

// GetAuthenticationFlows returns a list of authentication flows.
func (c *Client) GetAuthenticationFlows(accessToken string, realmName string) ([]keycloak.AuthenticationFlowRepresentation, error) {
	var resp = []keycloak.AuthenticationFlowRepresentation{}
	var err = c.forRealm(realmName).
		get(accessToken, &resp, url.Path(kcAuthenticationManagementPath+"/flows"), url.Param("realm", realmName))
	return resp, err
}

// CopyExistingAuthenticationFlow copy the existing authentication flow under a new name.
// 'flowAlias' is the name of the existing authentication flow,
// 'newName' is the new name of the authentication flow.
func (c *Client) CopyExistingAuthenticationFlow(accessToken string, realmName, flowAlias, newName string) error {
	var m = map[string]string{"newName": newName}
	_, err := c.forRealm(realmName).
		post(accessToken, nil, url.Path(kcAuthenticationManagementPath+"/flows/:flowAlias/copy"), url.Param("realm", realmName), url.Param("flowAlias", flowAlias), body.JSON(m))
	return err
}

// GetAuthenticationExecutionForFlow returns the authentication executions for a flow.
func (c *Client) GetAuthenticationExecutionForFlow(accessToken string, realmName, flowAlias string) (keycloak.AuthenticationExecutionInfoRepresentation, error) {
	var resp = keycloak.AuthenticationExecutionInfoRepresentation{}
	var err = c.forRealm(realmName).get(accessToken, &resp, url.Path(kcAuthenticationManagementPath+"/flows/:flowAlias/executions"), url.Param("realm", realmName), url.Param("flowAlias", flowAlias))
	return resp, err
}

// UpdateAuthenticationExecutionForFlow updates the authentication executions of a flow.
func (c *Client) UpdateAuthenticationExecutionForFlow(accessToken string, realmName, flowAlias string, authExecInfo keycloak.AuthenticationExecutionInfoRepresentation) error {
	return c.forRealm(realmName).
		put(accessToken, url.Path(kcAuthenticationManagementPath+"/flows/:flowAlias/executions"), url.Param("realm", realmName), url.Param("flowAlias", flowAlias), body.JSON(authExecInfo))
}

// CreateAuthenticationExecutionForFlow add a new authentication execution to a flow.
// 'flowAlias' is the alias of the parent flow.
func (c *Client) CreateAuthenticationExecutionForFlow(accessToken string, realmName, flowAlias, provider string) (string, error) {
	var m = map[string]string{"provider": provider}
	return c.forRealm(realmName).
		post(accessToken, nil, url.Path(kcAuthenticationManagementPath+"/flows/:flowAlias/executions/execution"), url.Param("realm", realmName), url.Param("flowAlias", flowAlias), body.JSON(m))
}

// CreateFlowWithExecutionForExistingFlow add a new flow with a new execution to an existing flow.
// 'flowAlias' is the alias of the parent authentication flow.
func (c *Client) CreateFlowWithExecutionForExistingFlow(accessToken string, realmName, flowAlias, alias, flowType, provider, description string) (string, error) {
	var m = map[string]string{"alias": alias, "type": flowType, "provider": provider, "description": description}
	return c.forRealm(realmName).
		post(accessToken, nil, url.Path(kcAuthenticationManagementPath+"/flows/:flowAlias/executions/flow"), url.Param("realm", realmName), url.Param("flowAlias", flowAlias), body.JSON(m))
}

// GetAuthenticationFlow gets the authentication flow for id.
func (c *Client) GetAuthenticationFlow(accessToken string, realmName, flowID string) (keycloak.AuthenticationFlowRepresentation, error) {
	var resp = keycloak.AuthenticationFlowRepresentation{}
	var err = c.forRealm(realmName).
		get(accessToken, &resp, url.Path(kcAuthenticationManagementPath+"/flows/:id"), url.Param("realm", realmName), url.Param("id", flowID))
	return resp, err
}

// DeleteAuthenticationFlow deletes an authentication flow.
func (c *Client) DeleteAuthenticationFlow(accessToken string, realmName, flowID string) error {
	return c.forRealm(realmName).
		delete(accessToken, url.Path(kcAuthenticationManagementPath+"/flows/:id"), url.Param("realm", realmName), url.Param("id", flowID))
}

// GetFormActionProviders returns a list of form action providers.
func (c *Client) GetFormActionProviders(accessToken string, realmName string) ([]map[string]any, error) {
	var resp = []map[string]any{}
	var err = c.forRealm(realmName).
		get(accessToken, &resp, url.Path(kcAuthenticationManagementPath+"/form-action-providers"), url.Param("realm", realmName))
	return resp, err
}

// GetFormProviders returns a list of form providers.
func (c *Client) GetFormProviders(accessToken string, realmName string) ([]map[string]any, error) {
	var resp = []map[string]any{}
	var err = c.forRealm(realmName).
		get(accessToken, &resp, url.Path(kcAuthenticationManagementPath+"/form-providers"), url.Param("realm", realmName))
	return resp, err
}

// GetConfigDescriptionForClients returns the configuration descriptions for all clients.
func (c *Client) GetConfigDescriptionForClients(accessToken string, realmName string) (map[string]any, error) {
	var resp = map[string]any{}
	var err = c.forRealm(realmName).
		get(accessToken, &resp, url.Path(kcAuthenticationManagementPath+"/per-client-config-description"), url.Param("realm", realmName))
	return resp, err
}

// RegisterRequiredAction register a new required action.
func (c *Client) RegisterRequiredAction(accessToken string, realmName, providerID, name string) error {
	var m = map[string]string{"providerId": providerID, "name": name}
	_, err := c.forRealm(realmName).
		post(accessToken, nil, url.Path(kcAuthenticationManagementPath+"/register-required-action"), url.Param("realm", realmName), body.JSON(m))
	return err
}

// GetRequiredActions returns a list of required actions.
func (c *Client) GetRequiredActions(accessToken string, realmName string) ([]keycloak.RequiredActionProviderRepresentation, error) {
	var resp = []keycloak.RequiredActionProviderRepresentation{}
	var err = c.forRealm(realmName).
		get(accessToken, &resp, url.Path(kcAuthenticationRequiredActionsPath), url.Param("realm", realmName))
	return resp, err
}

// GetRequiredAction returns the required action for the alias.
func (c *Client) GetRequiredAction(accessToken string, realmName, actionAlias string) (keycloak.RequiredActionProviderRepresentation, error) {
	var resp = keycloak.RequiredActionProviderRepresentation{}
	var err = c.forRealm(realmName).
		get(accessToken, &resp, url.Path(kcAuthenticationRequiredActionPath), url.Param("realm", realmName), url.Param("alias", actionAlias))
	return resp, err
}

// UpdateRequiredAction updates the required action.
func (c *Client) UpdateRequiredAction(accessToken string, realmName, actionAlias string, action keycloak.RequiredActionProviderRepresentation) error {
	return c.forRealm(realmName).
		put(accessToken, url.Path(kcAuthenticationRequiredActionPath), url.Param("realm", realmName), url.Param("alias", actionAlias), body.JSON(action))
}

// DeleteRequiredAction deletes the required action.
func (c *Client) DeleteRequiredAction(accessToken string, realmName, actionAlias string) error {
	return c.forRealm(realmName).
		delete(accessToken, url.Path(kcAuthenticationRequiredActionPath), url.Param("realm", realmName), url.Param("alias", actionAlias))
}

// GetUnregisteredRequiredActions returns a list of unregistered required actions.
func (c *Client) GetUnregisteredRequiredActions(accessToken string, realmName string) ([]map[string]any, error) {
	var resp = []map[string]any{}
	var err = c.forRealm(realmName).
		get(accessToken, &resp, url.Path(kcAuthenticationManagementPath+"/unregistered-required-actions"), url.Param("realm", realmName))
	return resp, err
}
