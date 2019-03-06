package keycloak

import (
	"fmt"

	"gopkg.in/h2non/gentleman.v2/plugins/url"
)

const (
	clientsPath  = "/auth/admin/realms/:realm/clients"
	clientIDPath = clientsPath + "/:id"
	clientSecret = clientsPath + "/client-secret"
)

// GetClients returns a list of clients belonging to the realm.
// Parameters: clientId (filter by clientId),
// viewableOnly (filter clients that cannot be viewed in full by admin, default="false")
func (c *Client) GetClients(realmName string, paramKV ...string) ([]ClientRepresentation, error) {
	if len(paramKV)%2 != 0 {
		return nil, fmt.Errorf("the number of key/val parameters should be even")
	}

	var resp = []ClientRepresentation{}
	var plugins = append(createQueryPlugins(paramKV...), url.Path(clientsPath), url.Param("realm", realmName))
	var err = c.get(&resp, plugins...)
	return resp, err
}

// GetClient get the representation of the client. idClient is the id of client (not client-id).
func (c *Client) GetClient(realmName, idClient string) (ClientRepresentation, error) {
	var resp = ClientRepresentation{}
	var err = c.get(&resp, url.Path(clientIDPath), url.Param("realm", realmName), url.Param("id", idClient))
	return resp, err
}

// GetSecret get the client secret. idClient is the id of client (not client-id).
func (c *Client) GetSecret(realmName, idClient string) (CredentialRepresentation, error) {
	var resp = CredentialRepresentation{}
	var err = c.get(&resp, url.Path(clientSecret), url.Param("realm", realmName), url.Param("id", idClient))
	return resp, err
}
