package api

import (
	"errors"

	"github.com/cloudtrust/keycloak-client"
	"gopkg.in/h2non/gentleman.v2/plugins/url"
)

const (
	clientsPath       = "/auth/admin/realms/:realm/clients"
	clientIDPath      = clientsPath + "/:id"
	clientSecret      = clientsPath + "/client-secret"
	clientMappersPath = clientIDPath + "/evaluate-scopes/protocol-mappers"
)

// GetClients returns a list of clients belonging to the realm.
// Parameters: clientId (filter by clientId),
// viewableOnly (filter clients that cannot be viewed in full by admin, default="false")
func (c *Client) GetClients(accessToken string, realmName string, paramKV ...string) ([]keycloak.ClientRepresentation, error) {
	if len(paramKV)%2 != 0 {
		return nil, errors.New(keycloak.MsgErrInvalidParam + "." + keycloak.EvenParams)
	}

	var resp = []keycloak.ClientRepresentation{}
	var plugins = append(createQueryPlugins(paramKV...), url.Path(clientsPath), url.Param("realm", realmName))
	var err = c.get(accessToken, &resp, plugins...)
	return resp, err
}

// GetClient get the representation of the client. idClient is the id of client (not client-id).
func (c *Client) GetClient(accessToken string, realmName, idClient string) (keycloak.ClientRepresentation, error) {
	var resp = keycloak.ClientRepresentation{}
	var err = c.get(accessToken, &resp, url.Path(clientIDPath), url.Param("realm", realmName), url.Param("id", idClient))
	return resp, err
}

// GetClientMappers gets mappers of the client specified by id
func (c *Client) GetClientMappers(accessToke string, realmName, idClient string) ([]keycloak.ClientMapperRepresentation, error) {
	var resp = []keycloak.ClientMapperRepresentation{}
	var err = c.get(accessToke, &resp, url.Path(clientMappersPath), url.Param("realm", realmName), url.Param("id", idClient))
	return resp, err
}

// GetSecret get the client secret. idClient is the id of client (not client-id).
func (c *Client) GetSecret(accessToken string, realmName, idClient string) (keycloak.CredentialRepresentation, error) {
	var resp = keycloak.CredentialRepresentation{}
	var err = c.get(accessToken, &resp, url.Path(clientSecret), url.Param("realm", realmName), url.Param("id", idClient))
	return resp, err
}
