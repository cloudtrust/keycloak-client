package api

import (
	"fmt"

	"github.com/cloudtrust/keycloak-client/v2"
	"gopkg.in/h2non/gentleman.v2/plugins/query"
	"gopkg.in/h2non/gentleman.v2/plugins/url"
)

const (
	statisticsPath        = "/auth/realms/master/api/admin/realms/:realm/statistics"
	statisticsUsers       = statisticsPath + "/users"
	statisticsCredentials = statisticsPath + "/credentials"
	statisticsOnboarding  = statisticsPath + "/onboarding"
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

// GetStatisticsOnboarding returns statistics on the onboarding of users of a client using the social realm in a specified range of date
func (c *Client) GetStatisticsOnboarding(accessToken string, socialRealmName, clientRealmName string, dateFrom, dateTo int64) (map[string]int64, error) {
	var resp = make(map[string]int64)
	var err = c.get(accessToken, &resp, url.Path(statisticsOnboarding), url.Param("realm", clientRealmName), query.Add("socialRealmName", socialRealmName),
		query.Add("dateFrom", fmt.Sprintf("%d", dateFrom)), query.Add("dateTo", fmt.Sprintf("%d", dateTo)))
	return resp, err
}
