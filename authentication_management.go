package keycloak

import (
	"gopkg.in/h2non/gentleman.v2/plugins/body"
	"gopkg.in/h2non/gentleman.v2/plugins/url"
)

const (
	authenticationManagementPath = "/auth/admin/realms/:realm/authentication"
)

// GetAuthenticatorProviders returns a list of authenticator providers.
func (c *Client) GetAuthenticatorProviders(realmName string) ([]map[string]interface{}, error) {
	var resp = []map[string]interface{}{}
	var err = c.get(&resp, url.Path(authenticationManagementPath+"/authenticator-providers"), url.Param("realm", realmName))
	return resp, err
}

// GetClientAuthenticatorProviders returns a list of client authenticator providers.
func (c *Client) GetClientAuthenticatorProviders(realmName string) ([]map[string]interface{}, error) {
	var resp = []map[string]interface{}{}
	var err = c.get(&resp, url.Path(authenticationManagementPath+"/client-authenticator-providers"), url.Param("realm", realmName))
	return resp, err
}

// GetAuthenticatorProviderConfig returns the authenticator provider’s configuration description.
func (c *Client) GetAuthenticatorProviderConfig(realmName, providerID string) (AuthenticatorConfigInfoRepresentation, error) {
	var resp = AuthenticatorConfigInfoRepresentation{}
	var err = c.get(&resp, url.Path(authenticationManagementPath+"/config-description/:providerID"), url.Param("realm", realmName), url.Param("providerID", providerID))
	return resp, err
}

// GetAuthenticatorConfig returns the authenticator configuration.
func (c *Client) GetAuthenticatorConfig(realmName, configID string) (AuthenticatorConfigRepresentation, error) {
	var resp = AuthenticatorConfigRepresentation{}
	var err = c.get(&resp, url.Path(authenticationManagementPath+"/config/:id"), url.Param("realm", realmName), url.Param("id", configID))
	return resp, err
}

// UpdateAuthenticatorConfig updates the authenticator configuration.
func (c *Client) UpdateAuthenticatorConfig(realmName, configID string, config AuthenticatorConfigRepresentation) error {
	return c.put(url.Path(authenticationManagementPath+"/config/:id"), url.Param("realm", realmName), url.Param("id", configID), body.JSON(config))
}

// DeleteAuthenticatorConfig deletes the authenticator configuration.
func (c *Client) DeleteAuthenticatorConfig(realmName, configID string) error {
	return c.delete(url.Path(authenticationManagementPath+"/config/:id"), url.Param("realm", realmName), url.Param("id", configID))
}

// CreateAuthenticationExecution add new authentication execution
func (c *Client) CreateAuthenticationExecution(realmName string, authExec AuthenticationExecutionRepresentation) error {
	return c.post(nil, url.Path(authenticationManagementPath+"/executions"), url.Param("realm", realmName), body.JSON(authExec))
}

// DeleteAuthenticationExecution deletes the execution.
func (c *Client) DeleteAuthenticationExecution(realmName, executionID string) error {
	return c.delete(url.Path(authenticationManagementPath+"/executions/:id"), url.Param("realm", realmName), url.Param("id", executionID))
}

// UpdateAuthenticationExecution update execution with new configuration.
func (c *Client) UpdateAuthenticationExecution(realmName, executionID string, authConfig AuthenticatorConfigRepresentation) error {
	return c.post(nil, url.Path(authenticationManagementPath+"/executions/:id/config"), url.Param("realm", realmName), url.Param("id", executionID), body.JSON(authConfig))
}

// LowerExecutionPriority lowers the execution’s priority.
func (c *Client) LowerExecutionPriority(realmName, executionID string) error {
	return c.post(nil, url.Path(authenticationManagementPath+"/executions/:id/lower-priority"), url.Param("realm", realmName), url.Param("id", executionID))
}

// RaiseExecutionPriority raise the execution’s priority.
func (c *Client) RaiseExecutionPriority(realmName, executionID string) error {
	return c.post(nil, url.Path(authenticationManagementPath+"/executions/:id/raise-priority"), url.Param("realm", realmName), url.Param("id", executionID))
}

// CreateAuthenticationFlow creates a new authentication flow.
func (c *Client) CreateAuthenticationFlow(realmName string, authFlow AuthenticationFlowRepresentation) error {
	return c.post(nil, url.Path(authenticationManagementPath+"/flows"), url.Param("realm", realmName), body.JSON(authFlow))
}

// GetAuthenticationFlows returns a list of authentication flows.
func (c *Client) GetAuthenticationFlows(realmName string) ([]AuthenticationFlowRepresentation, error) {
	var resp = []AuthenticationFlowRepresentation{}
	var err = c.get(&resp, url.Path(authenticationManagementPath+"/flows"), url.Param("realm", realmName))
	return resp, err
}

// CopyExistingAuthenticationFlow copy the existing authentication flow under a new name.
// 'flowAlias' is the name of the existing authentication flow,
// 'newName' is the new name of the authentication flow.
func (c *Client) CopyExistingAuthenticationFlow(realmName, flowAlias, newName string) error {
	var m = map[string]string{"newName": newName}
	return c.post(nil, url.Path(authenticationManagementPath+"/flows/:flowAlias/copy"), url.Param("realm", realmName), url.Param("flowAlias", flowAlias), body.JSON(m))
}

// GetAuthenticationExecutionForFlow returns the authentication executions for a flow.
func (c *Client) GetAuthenticationExecutionForFlow(realmName, flowAlias string) (AuthenticationExecutionInfoRepresentation, error) {
	var resp = AuthenticationExecutionInfoRepresentation{}
	var err = c.get(&resp, url.Path(authenticationManagementPath+"/flows/:flowAlias/executions"), url.Param("realm", realmName), url.Param("flowAlias", flowAlias))
	return resp, err
}

// UpdateAuthenticationExecutionForFlow updates the authentication executions of a flow.
func (c *Client) UpdateAuthenticationExecutionForFlow(realmName, flowAlias string, authExecInfo AuthenticationExecutionInfoRepresentation) error {
	return c.put(url.Path(authenticationManagementPath+"/flows/:flowAlias/executions"), url.Param("realm", realmName), url.Param("flowAlias", flowAlias), body.JSON(authExecInfo))
}

// CreateAuthenticationExecutionForFlow add a new authentication execution to a flow.
// 'flowAlias' is the alias of the parent flow.
func (c *Client) CreateAuthenticationExecutionForFlow(realmName, flowAlias, provider string) error {
	var m = map[string]string{"provider": provider}
	return c.post(url.Path(authenticationManagementPath+"/flows/:flowAlias/executions/execution"), url.Param("realm", realmName), url.Param("flowAlias", flowAlias), body.JSON(m))
}

// CreateFlowWithExecutionForExistingFlow add a new flow with a new execution to an existing flow.
// 'flowAlias' is the alias of the parent authentication flow.
func (c *Client) CreateFlowWithExecutionForExistingFlow(realmName, flowAlias, alias, flowType, provider, description string) error {
	var m = map[string]string{"alias": alias, "type": flowType, "provider": provider, "description": description}
	return c.post(url.Path(authenticationManagementPath+"/flows/:flowAlias/executions/flow"), url.Param("realm", realmName), url.Param("flowAlias", flowAlias), body.JSON(m))
}

// GetAuthenticationFlow gets the authentication flow for id.
func (c *Client) GetAuthenticationFlow(realmName, flowID string) (AuthenticationFlowRepresentation, error) {
	var resp = AuthenticationFlowRepresentation{}
	var err = c.get(&resp, url.Path(authenticationManagementPath+"/flows/:id"), url.Param("realm", realmName), url.Param("id", flowID))
	return resp, err
}

// DeleteAuthenticationFlow deletes an authentication flow.
func (c *Client) DeleteAuthenticationFlow(realmName, flowID string) error {
	return c.delete(url.Path(authenticationManagementPath+"/flows/:id"), url.Param("realm", realmName), url.Param("id", flowID))
}

// GetFormActionProviders returns a list of form action providers.
func (c *Client) GetFormActionProviders(realmName string) ([]map[string]interface{}, error) {
	var resp = []map[string]interface{}{}
	var err = c.get(&resp, url.Path(authenticationManagementPath+"/form-action-providers"), url.Param("realm", realmName))
	return resp, err
}

// GetFormProviders returns a list of form providers.
func (c *Client) GetFormProviders(realmName string) ([]map[string]interface{}, error) {
	var resp = []map[string]interface{}{}
	var err = c.get(&resp, url.Path(authenticationManagementPath+"/form-providers"), url.Param("realm", realmName))
	return resp, err
}

// GetConfigDescriptionForClients returns the configuration descriptions for all clients.
func (c *Client) GetConfigDescriptionForClients(realmName string) (map[string]interface{}, error) {
	var resp = map[string]interface{}{}
	var err = c.get(&resp, url.Path(authenticationManagementPath+"/per-client-config-description"), url.Param("realm", realmName))
	return resp, err
}

// RegisterRequiredAction register a new required action.
func (c *Client) RegisterRequiredAction(realmName, providerID, name string) error {
	var m = map[string]string{"providerId": providerID, "name": name}
	return c.post(url.Path(authenticationManagementPath+"/register-required-action"), url.Param("realm", realmName), body.JSON(m))
}

// GetRequiredActions returns a list of required actions.
func (c *Client) GetRequiredActions(realmName string) ([]RequiredActionProviderRepresentation, error) {
	var resp = []RequiredActionProviderRepresentation{}
	var err = c.get(&resp, url.Path(authenticationManagementPath+"/required-actions"), url.Param("realm", realmName))
	return resp, err
}

// GetRequiredAction returns the required action for the alias.
func (c *Client) GetRequiredAction(realmName, actionAlias string) (RequiredActionProviderRepresentation, error) {
	var resp = RequiredActionProviderRepresentation{}
	var err = c.get(&resp, url.Path(authenticationManagementPath+"/required-actions/:alias"), url.Param("realm", realmName), url.Param("alias", actionAlias))
	return resp, err
}

// UpdateRequiredAction updates the required action.
func (c *Client) UpdateRequiredAction(realmName, actionAlias string, action RequiredActionProviderRepresentation) error {
	return c.put(url.Path(authenticationManagementPath+"/required-actions/:alias"), url.Param("realm", realmName), url.Param("alias", actionAlias), body.JSON(action))
}

// DeleteRequiredAction deletes the required action.
func (c *Client) DeleteRequiredAction(realmName, actionAlias string) error {
	return c.delete(url.Path(authenticationManagementPath+"/required-actions/:alias"), url.Param("realm", realmName), url.Param("alias", actionAlias))
}

// GetUnregisteredRequiredActions returns a list of unregistered required actions.
func (c *Client) GetUnregisteredRequiredActions(realmName string) ([]map[string]interface{}, error) {
	var resp = []map[string]interface{}{}
	var err = c.get(&resp, url.Path(authenticationManagementPath+"/unregistered-required-actions"), url.Param("realm", realmName))
	return resp, err
}
