package api

import (
	"github.com/cloudtrust/keycloak-client/v2"
	"gopkg.in/h2non/gentleman.v2/plugins/body"
	"gopkg.in/h2non/gentleman.v2/plugins/url"
)

const (
	realmRootPath               = "/auth/admin/realms"
	realmPath                   = realmRootPath + "/:realm"
	realmCredentialRegistrators = realmPath + "/credential-registrators"
	exportRealmPath             = "/auth/realms/:realm/export/realm"
)

// GetRealms get the top level represention of all the realms. Nested information like users are
// not included.
func (c *Client) GetRealms(accessToken string) ([]keycloak.RealmRepresentation, error) {
	var resp = []keycloak.RealmRepresentation{}
	var err = c.get(accessToken, &resp, url.Path(realmRootPath))
	return resp, err
}

// CreateRealm creates the realm from its RealmRepresentation.
func (c *Client) CreateRealm(accessToken string, realm keycloak.RealmRepresentation) (string, error) {
	return c.post(accessToken, nil, url.Path(realmRootPath), body.JSON(realm))
}

// GetRealm get the top level represention of the realm. Nested information like users are
// not included.
func (c *Client) GetRealm(accessToken string, realmName string) (keycloak.RealmRepresentation, error) {
	var resp = keycloak.RealmRepresentation{}
	var err = c.get(accessToken, &resp, url.Path(realmPath), url.Param("realm", realmName))
	return resp, err
}

// UpdateRealm update the top lovel information of the realm. Any user, role or client information
// from the realm representation will be ignored.
func (c *Client) UpdateRealm(accessToken string, realmName string, realm keycloak.RealmRepresentation) error {
	return c.put(accessToken, url.Path(realmPath), url.Param("realm", realmName), body.JSON(realm))
}

// DeleteRealm deletes the realm.
func (c *Client) DeleteRealm(accessToken string, realmName string) error {
	return c.delete(accessToken, url.Path(realmPath), url.Param("realm", realmName))
}

// ExportRealm recovers the full realm.
func (c *Client) ExportRealm(accessToken string, realmName string) (keycloak.RealmRepresentation, error) {
	var resp = keycloak.RealmRepresentation{}
	var err = c.get(accessToken, &resp, url.Path(exportRealmPath), url.Param("realm", realmName))
	return resp, err
}

// GetRealmCredentialRegistrators returns list of credentials types available for the realm
func (c *Client) GetRealmCredentialRegistrators(accessToken string, realmName string) ([]string, error) {
	var resp = []string{}
	var err = c.get(accessToken, &resp, url.Path(realmCredentialRegistrators), url.Param("realm", realmName), hdrAcceptJSON)
	return resp, err
}
