package keycloak

import "gopkg.in/h2non/gentleman.v2/plugins/url"

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
	var err = c.get(&resp, url.Path(authenticationManagementPath+"client-authenticator-providers"), url.Param("realm", realmName))
	return resp, err
}

// GetAuthenticatorProviderConfig returns the authenticator provider’s configuration description.
func (c *Client) GetAuthenticatorProviderConfig(realmName, providerID string) (AuthenticatorConfigInfoRepresentation, error) {
	var resp = AuthenticatorConfigInfoRepresentation{}
	var err = c.get(&resp, url.Path(authenticationManagementPath+"config-description/:providerID"), url.Param("realm", realmName), url.Param("providerID", providerID))
	return resp, err
}

// GetAuthenticatorConfig returns the authenticator configuration.
func (c *Client) GetAuthenticatorConfig(realmName, configID string) (AuthenticatorConfigRepresentation, error) {
	var resp = AuthenticatorConfigRepresentation{}
	var err = c.get(&resp, url.Path(authenticationManagementPath+"config/:id"), url.Param("realm", realmName), url.Param("id", configID))
	return resp, err
}

// Update authenticator configuration
// PUT /{realm}/authentication/config/{id}

// Delete authenticator configuration
// DELETE /{realm}/authentication/config/{id}

// Add new authentication execution
// POST /{realm}/authentication/executions

// Delete execution
// DELETE /{realm}/authentication/executions/{executionId}

// Update execution with new configuration
// POST /{realm}/authentication/executions/{executionId}/config

// Lower execution’s priority
// POST /{realm}/authentication/executions/{executionId}/lower-priority

// Raise execution’s priority
// POST /{realm}/authentication/executions/{executionId}/raise-priority

// Create a new authentication flow
// POST /{realm}/authentication/flows

// Get authentication flows Returns a list of authentication flows.
// GET /{realm}/authentication/flows

// Copy existing authentication flow under a new name The new name is given as 'newName' attribute of the passed JSON object
// POST /{realm}/authentication/flows/{flowAlias}/copy

// Get authentication executions for a flow
// GET /{realm}/authentication/flows/{flowAlias}/executions

// Update authentication executions of a flow
// PUT /{realm}/authentication/flows/{flowAlias}/executions

// Add new authentication execution to a flow
// POST /{realm}/authentication/flows/{flowAlias}/executions/execution

// Add new flow with new execution to existing flow
// POST /{realm}/authentication/flows/{flowAlias}/executions/flow

// Get authentication flow for id
// GET /{realm}/authentication/flows/{id}

// Delete an authentication flow
// DELETE /{realm}/authentication/flows/{id}

// Get form action providers Returns a list of form action providers.
// GET /{realm}/authentication/form-action-providers

// Get form providers Returns a list of form providers.
// GET /{realm}/authentication/form-providers

// Get configuration descriptions for all clients
// GET /{realm}/authentication/per-client-config-description

// Register a new required actions
// POST /{realm}/authentication/register-required-action

// Get required actions Returns a list of required actions.
// GET /{realm}/authentication/required-actions

// Get required action for alias
// GET /{realm}/authentication/required-actions/{alias}

// Update required action
// PUT /{realm}/authentication/required-actions/{alias}

// Delete required action
// DELETE /{realm}/authentication/required-actions/{alias}

// Get unregistered required actions Returns a list of unregistered required actions.
// GET /{realm}/authentication/unregistered-required-actions
