package client

import (
	"gopkg.in/h2non/gentleman.v2/plugins/body"
	"gopkg.in/h2non/gentleman.v2/plugins/url"
)

const (
	realmRootPath = "/auth/admin/realms"
	realmPath     = realmRootPath + "/:realm"
)

func (c *client) GetRealms() ([]RealmRepresentation, error) {
	var resp = []RealmRepresentation{}
	var err = c.get(&resp, url.Path(realmRootPath))
	return resp, err
}

func (c *client) CreateRealm(realm RealmRepresentation) error {
	return c.post(url.Path(realmRootPath), body.JSON(realm))
}

func (c *client) GetRealm(realm string) (RealmRepresentation, error) {
	var resp = RealmRepresentation{}
	var err = c.get(&resp, url.Path(realmPath), url.Param("realm", realm))
	return resp, err
}

func (c *client) UpdateRealm(realmName string, realm RealmRepresentation) error {
	return c.put(url.Path(realmPath), url.Param("realm", realmName), body.JSON(realm))
}

func (c *client) DeleteRealm(realm string) error {
	return c.delete(url.Path(realmPath), url.Param("realm", realm))
}
