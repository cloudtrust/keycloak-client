package api

import (
	"github.com/cloudtrust/keycloak-client/v2"
	"gopkg.in/h2non/gentleman.v2/plugins/body"
	"gopkg.in/h2non/gentleman.v2/plugins/url"
)

const (
	// API Keycloak out-of-the-box
	kcComponentsPath  = "/auth/admin/realms/:realm/components"
	kcComponentIDPath = kcComponentsPath + "/:id"
)

// GetComponents gets the list of components in a realm.
func (c *Client) GetComponents(accessToken string, realmName string, paramKV ...string) ([]keycloak.ComponentRepresentation, error) {
	resp := []keycloak.ComponentRepresentation{}

	plugins := append(c.createQueryPlugins(paramKV...), url.Path(kcComponentsPath), url.Param("realm", realmName))
	err := c.forRealm(accessToken, realmName).
		get(accessToken, &resp, plugins...)

	return resp, err
}

// CreateComponent creates a new component.
func (c *Client) CreateComponent(accessToken string, realmName string, compRep keycloak.ComponentRepresentation) error {
	_, err := c.forRealm(accessToken, realmName).
		post(accessToken, nil, url.Path(kcComponentsPath), url.Param("realm", realmName), body.JSON(compRep))
	return err
}

// UpdateComponent updates the component.
func (c *Client) UpdateComponent(accessToken string, realmName, componentID string, componentRep keycloak.ComponentRepresentation) error {
	return c.forRealm(accessToken, realmName).
		put(accessToken, url.Path(kcComponentIDPath), url.Param("realm", realmName), url.Param("id", componentID), body.JSON(componentRep))
}
