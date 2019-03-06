package keycloak

import "gopkg.in/h2non/gentleman.v2/plugins/url"

const (
	clientRegistrationPolicyPath = "/auth/admin/realms/:realm/client-registration-policy/providers"
)

// GetClientRegistrationPolicy is the base path to retrieve providers with the configProperties properly filled.
func (c *Client) GetClientRegistrationPolicy(accessToken string, realmName, configID string) ([]ComponentTypeRepresentation, error) {
	var resp = []ComponentTypeRepresentation{}
	var err = c.get(accessToken, &resp, url.Path(clientRegistrationPolicyPath), url.Param("realm", realmName))
	return resp, err
}
