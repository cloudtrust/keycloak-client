package keycloak

import (
	"gopkg.in/h2non/gentleman.v2/plugins/url"
)

const (
	idpsPath       = "/auth/admin/realms/:realm/identity-provider/instances"
	idpAliasPath   = idpsPath + "/:alias"
	idpMappersPath = idpAliasPath + "/mappers"
)

func (c *Client) GetIdps(accessToken string, realmName string) ([]IdentityProviderRepresentation, error) {
	var resp = []IdentityProviderRepresentation{}
	var err = c.get(accessToken, &resp, url.Path(idpsPath), url.Param("realm", realmName))
	return resp, err
}

func (c *Client) GetIdp(accessToken string, realmName string, idpAlias string) (IdentityProviderRepresentation, error) {
	var resp = IdentityProviderRepresentation{}
	var err = c.get(accessToken, &resp, url.Path(idpAliasPath), url.Param("realm", realmName), url.Param("alias", idpAlias))
	return resp, err
}

func (c *Client) GetIdpMappers(accessToken string, realmName string, idpAlias string) ([]IdentityProviderMapperRepresentation, error) {
	var resp = []IdentityProviderMapperRepresentation{}
	var err = c.get(accessToken, &resp, url.Path(idpMappersPath), url.Param("realm", realmName), url.Param("alias", idpAlias))
	return resp, err
}
