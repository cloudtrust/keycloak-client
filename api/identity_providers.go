package api

import (
	"github.com/cloudtrust/keycloak-client/v2"
	"gopkg.in/h2non/gentleman.v2/plugins/body"
	"gopkg.in/h2non/gentleman.v2/plugins/url"
)

const (
	idpsPath       = "/auth/admin/realms/:realm/identity-provider/instances"
	idpAliasPath   = idpsPath + "/:alias"
	idpMappersPath = idpAliasPath + "/mappers"
)

// GetIdps gets the list of identity providers
func (c *Client) GetIdps(accessToken string, realmName string) ([]keycloak.IdentityProviderRepresentation, error) {
	var resp = []keycloak.IdentityProviderRepresentation{}
	var err = c.get(accessToken, &resp, url.Path(idpsPath), url.Param("realm", realmName))
	return resp, err
}

// UpdateIdp updates the identity provider. idpAlias is the alias of identity provider.
func (c *Client) UpdateIdp(accessToken string, realmName, idpAlias string, idpRep keycloak.IdentityProviderRepresentation) error {
	return c.put(accessToken, url.Path(idpAliasPath), url.Param("realm", realmName), url.Param("alias", idpAlias), body.JSON(idpRep))
}

// GetIdp gets an identity provider matching the given alias
func (c *Client) GetIdp(accessToken string, realmName string, idpAlias string) (keycloak.IdentityProviderRepresentation, error) {
	var resp = keycloak.IdentityProviderRepresentation{}
	var err = c.get(accessToken, &resp, url.Path(idpAliasPath), url.Param("realm", realmName), url.Param("alias", idpAlias))
	return resp, err
}

// GetIdpMappers gets the mappers of the specified identity provider
func (c *Client) GetIdpMappers(accessToken string, realmName string, idpAlias string) ([]keycloak.IdentityProviderMapperRepresentation, error) {
	var resp = []keycloak.IdentityProviderMapperRepresentation{}
	var err = c.get(accessToken, &resp, url.Path(idpMappersPath), url.Param("realm", realmName), url.Param("alias", idpAlias))
	return resp, err
}
