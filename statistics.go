package keycloak

import (
	"gopkg.in/h2non/gentleman.v2/plugins/url"
)

const (
	statisticsPath        = "/auth/realms/master/api/admin/realms/:realm/statistics"
	statisticsUsers       = statisticsPath + "/users"
	statisticsCredentials = statisticsPath + "/credentials"
)

// StatisticsUsersRepresentation elements returned by GetStatisticsUsers
type StatisticsUsersRepresentation struct {
	Total    int64 `json:"total,omitempty"`
	Disabled int64 `json:"disabled,omitempty"`
	Inactive int64 `json:"inactive,omitempty"`
}

// GetStatisticsUsers returns statisctics on the total number of users and on their status
func (c *Client) GetStatisticsUsers(accessToken string, realmName string) (StatisticsUsersRepresentation, error) {
	var resp = StatisticsUsersRepresentation{}
	var err = c.get(accessToken, &resp, url.Path(statisticsUsers), url.Param("realm", realmName))
	return resp, err
}

// GetStatisticsAuthenticators returns statistics on the authenticators used by the users on a certain realm
func (c *Client) GetStatisticsAuthenticators(accessToken string, realmName string) (map[string]int64, error) {
	var resp = make(map[string]int64)
	var err = c.get(accessToken, &resp, url.Path(statisticsCredentials), url.Param("realm", realmName))
	return resp, err
}
