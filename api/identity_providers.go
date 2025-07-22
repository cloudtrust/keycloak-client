package api

import (
	"github.com/cloudtrust/keycloak-client/v2"
	"gopkg.in/h2non/gentleman.v2/plugins/body"
	"gopkg.in/h2non/gentleman.v2/plugins/url"
)

const (
	// API Keycloak out-of-the-box
	kcIdpsPath       = "/auth/admin/realms/:realm/identity-provider/instances"
	kcIdpAliasPath   = kcIdpsPath + "/:alias"
	kcIdpMappersPath = kcIdpAliasPath + "/mappers"
)

// GetIdps gets the list of identity providers
func (c *Client) GetIdps(accessToken string, realmName string) ([]keycloak.IdentityProviderRepresentation, error) {
	var resp = []keycloak.IdentityProviderRepresentation{}
	var err = c.forRealm(accessToken, realmName).
		get(accessToken, &resp, url.Path(kcIdpsPath), url.Param("realm", realmName))
	return resp, err
}

// CreateIdp creates a new identity provider
func (c *Client) CreateIdp(accessToken string, realmName string, idpRep keycloak.IdentityProviderRepresentation) error {
	_, err := c.forRealm(accessToken, realmName).
		post(accessToken, nil, url.Path(kcIdpsPath), url.Param("realm", realmName), body.JSON(idpRep))
	return err
}

// UpdateIdp updates the identity provider. idpAlias is the alias of identity provider.
func (c *Client) UpdateIdp(accessToken string, realmName, idpAlias string, idpRep keycloak.IdentityProviderRepresentation) error {
	return c.forRealm(accessToken, realmName).
		put(accessToken, url.Path(kcIdpAliasPath), url.Param("realm", realmName), url.Param("alias", idpAlias), body.JSON(idpRep))
}

// GetIdp gets an identity provider matching the given alias
func (c *Client) GetIdp(accessToken string, realmName string, idpAlias string) (keycloak.IdentityProviderRepresentation, error) {
	var resp = keycloak.IdentityProviderRepresentation{}
	var err = c.forRealm(accessToken, realmName).
		get(accessToken, &resp, url.Path(kcIdpAliasPath), url.Param("realm", realmName), url.Param("alias", idpAlias))
	return resp, err
}

// DeleteIdp deletes an identity provider matching the given alias
func (c *Client) DeleteIdp(accessToken string, realmName string, idpAlias string) error {
	return c.forRealm(accessToken, realmName).
		delete(accessToken, url.Path(kcIdpAliasPath), url.Param("realm", realmName), url.Param("alias", idpAlias))
}

// GetIdpMappers gets the mappers of the specified identity provider
func (c *Client) GetIdpMappers(accessToken string, realmName string, idpAlias string) ([]keycloak.IdentityProviderMapperRepresentation, error) {
	var resp = []keycloak.IdentityProviderMapperRepresentation{}
	var err = c.forRealm(accessToken, realmName).
		get(accessToken, &resp, url.Path(kcIdpMappersPath), url.Param("realm", realmName), url.Param("alias", idpAlias))
	return resp, err
}
