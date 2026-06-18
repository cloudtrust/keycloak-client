package api

import "gopkg.in/h2non/gentleman.v2/plugins/url"

const (
	ctIDPTeamMemberPath = "/auth/realms/:realmReq/api/idp/realms/:realm/team-members/:userId"
)

// DeleteExtIDPTeamMemberUser deletes a team member user
func (c *Client) DeleteExtIDPTeamMemberUser(accessToken string, realmReqName string, realmName string, userID string) error {
	return c.forRealm(accessToken, realmName).
		delete(accessToken, url.Path(ctIDPTeamMemberPath),
			url.Param("realmReq", realmReqName),
			url.Param("realm", realmName),
			url.Param("userId", userID))
}
