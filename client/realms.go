package client


import (
	"github.com/pkg/errors"
	gentleman "gopkg.in/h2non/gentleman.v2"
	"encoding/json"
)

func (c *client) GetRealms() ([]map[string]interface{}, error) {
	var getRealms_Path string = "/auth/admin/realms"
	var resp *gentleman.Response
	{
		var err error
		resp, err = c.do(getRealms_Path)
		if err != nil {
			return nil, errors.Wrap(err, "Get Realms failed.")
		}
	}
	var result []map[string]interface{}
	{
		var err error
		err = json.Unmarshal(resp.Bytes(), &result)
		if err != nil {
			return nil, errors.Wrap(err, "Get Realms failed to unmarshal response.")
		}
	}
	return result, nil
}
