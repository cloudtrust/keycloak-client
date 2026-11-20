package api

import (
	"github.com/cloudtrust/keycloak-client/v2"
	"gopkg.in/h2non/gentleman.v2/plugins/url"
)

const (
	// API Cloudtrust
	ctExternalSystemAPIPath   = "/auth/realms/master/extservice"
	ctCircuitBreakerPath      = ctExternalSystemAPIPath + "/realms/:realm/circuit-breaker"
	ctCircuitBreakerOpenPath  = ctCircuitBreakerPath + "/open"
	ctCircuitBreakerClosePath = ctCircuitBreakerPath + "/open"
)

// GetCircuitBreakerStatus gets the circuit breaker status for the given realm
func (c *Client) GetCircuitBreakerStatus(accessToken string, realmName string) (keycloak.CircuitBreakerRepresentation, error) {
	var resp keycloak.CircuitBreakerRepresentation
	var err = c.forRealm(accessToken, realmName).
		get(accessToken, &resp, url.Path(ctCircuitBreakerPath), url.Param("realm", realmName))
	return resp, err
}

// OpenCircuitBreaker opens the circuit breaker of the given realm.
func (c *Client) OpenCircuitBreaker(accessToken string, realmName string) error {
	return c.updateCircuitBreakerStatus(accessToken, realmName, ctCircuitBreakerOpenPath)
}

// CloseCircuitBreaker closes the circuit breaker of the given realm.
func (c *Client) CloseCircuitBreaker(accessToken string, realmName string) error {
	return c.updateCircuitBreakerStatus(accessToken, realmName, ctCircuitBreakerClosePath)
}

func (c *Client) updateCircuitBreakerStatus(accessToken string, realmName string, requestPath string) error {
	_, err := c.forRealm(accessToken, realmName).
		post(accessToken, nil, url.Path(requestPath), url.Param("realm", realmName))
	return err
}
