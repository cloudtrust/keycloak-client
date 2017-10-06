package client

import (
	"fmt"
	"net/url"
	"github.com/pkg/errors"
	"time"
	"net/http"
	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugin"
	"gopkg.in/h2non/gentleman.v2/plugins/timeout"
	//"gopkg.in/h2non/gentleman.v2/plugins/multipart"
)


type Client interface {
	GetRealms() ([]RealmRepresentation, error)
	GetUsers(realm string) ([]UserRepresentation, error)
}

type HttpConfig struct {
	Addr string
	Username string
	Password string
	Timeout time.Duration
}

type client struct {
	username string
	password string
	accessToken string
	httpClient *gentleman.Client
}


func NewHttpClient(config HttpConfig) (Client, error) {
	var u *url.URL
	{
		var err error
		u, err = url.Parse(config.Addr)
		if err != nil {
			return nil, errors.Wrap(err, "Parse failed")
		}
	}

	if u.Scheme != "http" {
		var m string = fmt.Sprintf("Unsupported protocol %s. Your address must start with http://", u.Scheme)
		return nil, errors.New(m)
	}

	var httpClient *gentleman.Client = gentleman.New()
	{
		httpClient = httpClient.URL(u.String())
		httpClient = httpClient.Use(timeout.Request(config.Timeout))
	}

	return &client{
		username: config.Username,
		password: config.Password,
		httpClient: httpClient,
	}, nil
}

func (c *client) getToken() error {
	var req *gentleman.Request
	{
		var authPath string = "/auth/realms/master/protocol/openid-connect/token"
		//var formData multipart.FormData = multipart.FormData{
		//	Data: map[string]multipart.Values{
		//		"username": multipart.Values{c.username},
		//		"password": multipart.Values{c.password},
		//		"grant_type": multipart.Values{"password"},
		//		"client_id": multipart.Values{"admin-cli"},
		//	},
		//}
		req = c.httpClient.Post()
		req = req.SetHeader("Content-Type", "application/x-www-form-urlencoded")
		req = req.Path(authPath)
		req = req.Type("urlencoded")
		req = req.BodyString(fmt.Sprintf("username=%s&password=%s&grant_type=password&client_id=admin-cli",c.username,c.password))
	}

	var resp *gentleman.Response
	{
		var err error
		resp, err = req.Do()
		if err != nil {
			return errors.Wrap(err, "Failed to get the token response")
		}
	}
	defer resp.Close()

	var unmarshalledBody map[string]interface{}
	{
		var err error
		err = resp.JSON(&unmarshalledBody)
		if err != nil {
			return errors.Wrap(err, "Failed to unmarshal response json")
		}
	}

	var accessToken interface{}
	{
		var ok bool
		accessToken, ok = unmarshalledBody["access_token"]
		if !ok {
			return errors.New("No access token in reponse body")
		}
	}

	c.accessToken = accessToken.(string)

	return nil
}

func (c *client) do(path string, plugins ...plugin.Plugin) (*gentleman.Response, error) {
	var req *gentleman.Request = c.httpClient.Get()
	{
		req = req.Path(path)
		req = req.SetHeader("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
		for _, p := range plugins {
			req = req.Use(p)
		}
	}
	var resp *gentleman.Response
	{
		var err error
		resp, err = req.Do()
		if err != nil {
			return nil, errors.Wrap(err, "Failed to get the response")
		}
		switch resp.StatusCode {
		case http.StatusUnauthorized:
			var err = c.getToken()
			//This induces a potential infinite loop, where a new token gets requested and the
			//process gets delayed so much it expires before the recursion.
			//It is decided that should this happen, the machine would be considered to be in terrible shape
			//and the loop wouldn't be the biggest problem.
			if err != nil {
				return nil, errors.Wrap(err, "Failed to get token")
			}
			return c.do(path)
		default:
			return resp, nil
		}
	}
}