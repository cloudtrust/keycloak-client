package api

import (
	"github.com/cloudtrust/keycloak-client"
	"gopkg.in/h2non/gentleman.v2/plugins/url"
)

const (
	statisticsPath        = "/auth/realms/master/api/admin/realms/:realm/statistics"
	statisticsUsers       = statisticsPath + "/users"
	statisticsCredentials = statisticsPath + "/credentials"
)

// GetStatisticsUsers returns statisctics on the total number of users and on their status
func (c *Client) GetStatisticsUsers(accessToken string, realmName string) (keycloak.StatisticsUsersRepresentation, error) {
	var resp = keycloak.StatisticsUsersRepresentation{}
	var err = c.get(accessToken, &resp, url.Path(statisticsUsers), url.Param("realm", realmName))
	return resp, err
}

// GetStatisticsAuthenticators returns statistics on the authenticators used by the users on a certain realm
func (c *Client) GetStatisticsAuthenticators(accessToken string, realmName string) (map[string]int64, error) {
	var resp = make(map[string]int64)
	var err = c.get(accessToken, &resp, url.Path(statisticsCredentials), url.Param("realm", realmName))
	return resp, err
}
