package client

import (
	"fmt"
	"gopkg.in/h2non/gentleman.v2"
	"github.com/pkg/errors"
	"encoding/json"
)

func (c *client)GetUsers(realm string) ([]UserRepresentation, error) {
	var getUsers_Path string = fmt.Sprintf("/auth/admin/realms/%s/users", realm)
	var resp *gentleman.Response
	{
		var err error
		resp, err = c.do(getUsers_Path)
		if err != nil {
			return nil, errors.Wrap(err, "Get Realms failed.")
		}
	}
	var result []UserRepresentation
	{
		var err error
		err = json.Unmarshal(resp.Bytes(), &result)
		if err != nil {
			return nil, errors.Wrap(err, "Get Users failed to unmarshal response.")
		}
	}
	return result, nil
}