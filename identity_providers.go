package keycloak

import (
	"errors"

	"gopkg.in/h2non/gentleman.v2/plugins/url"
)

const (
	idpsPath       = "/auth/admin/realms/:realm/identity-provider/instances"
	idpAliasPath   = idpsPath + "/:alias"
	idpMappersPath = idpAliasPath + "/mappers"
)

func (c *Client) GetIdps(accessToken string, realmName string, paramKV ...string) ([]IdentityProviderRepresentation, error) {
	if len(paramKV)%2 != 0 {
		return nil, errors.New(MsgErrInvalidParam + "." + EvenParams)
	}

	var resp = []IdentityProviderRepresentation{}
	var plugins = append(createQueryPlugins(paramKV...), url.Path(idpsPath), url.Param("realm", realmName))
	var err = c.get(accessToken, &resp, plugins...)
	return resp, err
}

func (c *Client) GetIdp(accessToken string, realmName string, idpAlias string, paramKV ...string) (IdentityProviderRepresentation, error) {
	var resp = IdentityProviderRepresentation{}

	if len(paramKV)%2 != 0 {
		return resp, errors.New(MsgErrInvalidParam + "." + EvenParams)
	}

	var plugins = append(createQueryPlugins(paramKV...), url.Path(idpAliasPath), url.Param("realm", realmName), url.Param("alias", idpAlias))
	var err = c.get(accessToken, &resp, plugins...)
	return resp, err
}

func (c *Client) GetIdpMappers(accessToken string, realmName string, idpAlias string, paramKV ...string) ([]IdentityProviderMapperRepresentation, error) {
	var resp = []IdentityProviderMapperRepresentation{}

	if len(paramKV)%2 != 0 {
		return resp, errors.New(MsgErrInvalidParam + "." + EvenParams)
	}

	var plugins = append(createQueryPlugins(paramKV...), url.Path(idpMappersPath), url.Param("realm", realmName), url.Param("alias", idpAlias))
	var err = c.get(accessToken, &resp, plugins...)
	return resp, err
}
