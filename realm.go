package keycloak

import (
	"gopkg.in/h2non/gentleman.v2/plugins/body"
	"gopkg.in/h2non/gentleman.v2/plugins/url"
)

const (
	realmRootPath = "/auth/admin/realms"
	realmPath     = realmRootPath + "/:realm"
)

// GetRealms get the top level represention of all the realms. Nested information like users are
// not included.
func (c *Client) GetRealms() ([]RealmRepresentation, error) {
	var resp = []RealmRepresentation{}
	var err = c.get(&resp, url.Path(realmRootPath))
	return resp, err
}

// CreateRealm creates the realm from its RealmRepresentation.
func (c *Client) CreateRealm(realm RealmRepresentation) error {
	return c.post(url.Path(realmRootPath), body.JSON(realm))
}

// GetRealm get the top level represention of the realm. Nested information like users are
// not included.
func (c *Client) GetRealm(realmName string) (RealmRepresentation, error) {
	var resp = RealmRepresentation{}
	var err = c.get(&resp, url.Path(realmPath), url.Param("realm", realmName))
	return resp, err
}

// UpdateRealm update the top lovel information of the realm. Any user, role or client information
// from the realm representation will be ignored.
func (c *Client) UpdateRealm(realmName string, realm RealmRepresentation) error {
	return c.put(url.Path(realmPath), url.Param("realm", realmName), body.JSON(realm))
}

// DeleteRealm deletes the realm.
func (c *Client) DeleteRealm(realmName string) error {
	return c.delete(url.Path(realmPath), url.Param("realm", realmName))
}
