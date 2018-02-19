package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	oidc "github.com/coreos/go-oidc"
	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugin"
	"gopkg.in/h2non/gentleman.v2/plugins/timeout"
)

type HttpConfig struct {
	Addr     string
	Username string
	Password string
	Timeout  time.Duration
}

type client struct {
	username     string
	password     string
	accessToken  string
	oidcProvider *oidc.Provider
	httpClient   *gentleman.Client
}

func New(config HttpConfig) (*client, error) {
	var u *url.URL
	{
		var err error
		u, err = url.Parse(config.Addr)
		if err != nil {
			return nil, fmt.Errorf("could not parse URL: %v", err)
		}
	}

	if u.Scheme != "http" {
		return nil, fmt.Errorf("protocol not supported, your address must start with http://, not %v", u.Scheme)
	}

	var httpClient = gentleman.New()
	{
		httpClient = httpClient.URL(u.String())
		httpClient = httpClient.Use(timeout.Request(config.Timeout))
	}

	var oidcProvider *oidc.Provider
	{
		var err error
		var issuer = fmt.Sprintf("%s/auth/realms/master", u.String())
		oidcProvider, err = oidc.NewProvider(context.Background(), issuer)
		if err != nil {
			return nil, fmt.Errorf("could not create oidc provider: %v", err)
		}
	}

	return &client{
		username:     config.Username,
		password:     config.Password,
		oidcProvider: oidcProvider,
		httpClient:   httpClient,
	}, nil
}

func (c *client) getToken() error {
	var req *gentleman.Request
	{
		var authPath = "/auth/realms/master/protocol/openid-connect/token"
		req = c.httpClient.Post()
		req = req.SetHeader("Content-Type", "application/x-www-form-urlencoded")
		req = req.Path(authPath)
		req = req.Type("urlencoded")
		req = req.BodyString(fmt.Sprintf("username=%s&password=%s&grant_type=password&client_id=admin-cli", c.username, c.password))
	}

	var resp *gentleman.Response
	{
		var err error
		resp, err = req.Do()
		if err != nil {
			return fmt.Errorf("could not get token: %v", err)
		}
	}
	defer resp.Close()

	var unmarshalledBody map[string]interface{}
	{
		var err error
		err = resp.JSON(&unmarshalledBody)
		if err != nil {
			return fmt.Errorf("could not unmarshal response: %v", err)
		}
	}

	var accessToken interface{}
	{
		var ok bool
		accessToken, ok = unmarshalledBody["access_token"]
		if !ok {
			return fmt.Errorf("could not find access token in response body")
		}
	}

	c.accessToken = accessToken.(string)
	return nil
}

func (c *client) verifyToken() error {
	var v = c.oidcProvider.Verifier(&oidc.Config{SkipClientIDCheck: true})

	var err error
	_, err = v.Verify(context.Background(), c.accessToken)
	return err
}

func (c *client) get(data interface{}, plugins ...plugin.Plugin) error {
	var req = c.httpClient.Get()
	req = applyPlugins(req, c.accessToken, plugins...)

	var resp *gentleman.Response
	{
		var err error
		resp, err = req.Do()
		if err != nil {
			return fmt.Errorf("could not get response: %v", err)
		}

		switch {
		case resp.StatusCode == http.StatusUnauthorized:
			// If the token is not valid (expired, ...) ask a new one.
			if err = c.verifyToken(); err != nil {
				var err = c.getToken()
				if err != nil {
					return fmt.Errorf("could not get token: %v", err)
				}
			}
			return c.get(data, plugins...)
		case resp.StatusCode >= 400:
			return handleError(resp)
		case resp.StatusCode >= 200:
			return json.Unmarshal(resp.Bytes(), data)
		default:
			return fmt.Errorf("unknown response status code: %v", resp.StatusCode)
		}
	}
}

func (c *client) post(plugins ...plugin.Plugin) error {
	var req = c.httpClient.Post()
	req = applyPlugins(req, c.accessToken, plugins...)

	var resp *gentleman.Response
	{
		var err error
		resp, err = req.Do()
		if err != nil {
			return fmt.Errorf("could not get response: %v", err)
		}

		switch {
		case resp.StatusCode == http.StatusUnauthorized:
			// If the token is not valid (expired, ...) ask a new one.
			if err = c.verifyToken(); err != nil {
				var err = c.getToken()
				if err != nil {
					return fmt.Errorf("could not get token: %v", err)
				}
			}
			return c.post(plugins...)
		case resp.StatusCode >= 400:
			return handleError(resp)
		case resp.StatusCode >= 200:
			return nil
		default:
			return fmt.Errorf("unknown response status code: %v", resp.StatusCode)
		}
	}
}

func (c *client) delete(plugins ...plugin.Plugin) error {
	var req = c.httpClient.Delete()
	req = applyPlugins(req, c.accessToken, plugins...)

	var resp *gentleman.Response
	{
		var err error
		resp, err = req.Do()
		if err != nil {
			return fmt.Errorf("could not get response: %v", err)
		}

		switch {
		case resp.StatusCode == http.StatusUnauthorized:
			// If the token is not valid (expired, ...) ask a new one.
			if err = c.verifyToken(); err != nil {
				var err = c.getToken()
				if err != nil {
					return fmt.Errorf("could not get token: %v", err)
				}
			}
			return c.delete(plugins...)
		case resp.StatusCode >= 400:
			return handleError(resp)
		case resp.StatusCode >= 200:
			return nil
		default:
			return fmt.Errorf("unknown response status code: %v", resp.StatusCode)
		}
	}
}

func (c *client) put(plugins ...plugin.Plugin) error {
	var req = c.httpClient.Put()
	req = applyPlugins(req, c.accessToken, plugins...)

	var resp *gentleman.Response
	{
		var err error
		resp, err = req.Do()
		if err != nil {
			return fmt.Errorf("could not get response: %v", err)
		}

		switch {
		case resp.StatusCode == http.StatusUnauthorized:
			// If the token is not valid (expired, ...) ask a new one.
			if err = c.verifyToken(); err != nil {
				var err = c.getToken()
				if err != nil {
					return fmt.Errorf("could not get token: %v", err)
				}
			}
			return c.put(plugins...)
		case resp.StatusCode >= 400:
			return handleError(resp)
		case resp.StatusCode >= 200:
			return nil
		default:
			return fmt.Errorf("unknown response status code: %v", resp.StatusCode)
		}
	}
}

func applyPlugins(req *gentleman.Request, accessToken string, plugins ...plugin.Plugin) *gentleman.Request {
	var r = req.SetHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	for _, p := range plugins {
		r = r.Use(p)
	}
	return r
}

func handleError(resp *gentleman.Response) error {
	var m = map[string]string{}
	var err = json.Unmarshal(resp.Bytes(), &m)
	if err != nil {
		return fmt.Errorf("invalid status code: %v; could not unmarshal response: %v", resp.StatusCode, err)
	}
	return fmt.Errorf("error message: %v", m)
}

func str(s string) *string {
	return &s
}
