package keycloak

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	oidc "github.com/coreos/go-oidc"
	"github.com/pkg/errors"
	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugin"
	"gopkg.in/h2non/gentleman.v2/plugins/query"
	"gopkg.in/h2non/gentleman.v2/plugins/timeout"
)

// Config is the keycloak client http config.
type Config struct {
	Addr     string
	Username string
	Password string
	Timeout  time.Duration
}

// Client is the keycloak client.
type Client struct {
	username     string
	password     string
	accessToken  string
	oidcProvider *oidc.Provider
	httpClient   *gentleman.Client
}

// New returns a keycloak client.
func New(config Config) (*Client, error) {
	var u *url.URL
	{
		var err error
		u, err = url.Parse(config.Addr)
		if err != nil {
			return nil, errors.Wrap(err, "could not parse URL")
		}
	}

	// if u.Scheme != "http" {
	// 	return nil, fmt.Errorf("protocol not supported, your address must start with http://, not %v", u.Scheme)
	// }

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
			return nil, errors.Wrap(err, "could not create oidc provider")
		}
	}

	return &Client{
		username:     config.Username,
		password:     config.Password,
		oidcProvider: oidcProvider,
		httpClient:   httpClient,
	}, nil
}

// getToken get a token from keycloak.
func (c *Client) getToken() error {
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
			return errors.Wrap(err, "could not get token")
		}
	}
	defer resp.Close()

	var unmarshalledBody map[string]interface{}
	{
		var err error
		err = resp.JSON(&unmarshalledBody)
		if err != nil {
			return errors.Wrap(err, "could not unmarshal response")
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

// verifyToken token verify a token. It returns an error it is malformed, expired,...
func (c *Client) verifyToken() error {
	var v = c.oidcProvider.Verifier(&oidc.Config{SkipClientIDCheck: true})

	var err error
	_, err = v.Verify(context.Background(), c.accessToken)
	return err
}

// get is a HTTP get method.
func (c *Client) get(data interface{}, plugins ...plugin.Plugin) error {
	var req = c.httpClient.Get()
	req = applyPlugins(req, c.accessToken, plugins...)

	var resp *gentleman.Response
	{
		var err error
		resp, err = req.Do()
		if err != nil {
			return errors.Wrap(err, "could not get response")
		}

		switch {
		case resp.StatusCode == http.StatusUnauthorized:
			// If the token is not valid (expired, ...) ask a new one.
			if err = c.verifyToken(); err != nil {
				var err = c.getToken()
				if err != nil {
					return errors.Wrap(err, "could not get token: %v")
				}
			}
			return c.get(data, plugins...)
		case resp.StatusCode >= 400:
			return fmt.Errorf("invalid status code: '%v': %v", resp.RawResponse.Status, string(resp.Bytes()))
		case resp.StatusCode >= 200:
			return json.Unmarshal(resp.Bytes(), data)
		default:
			return fmt.Errorf("unknown response status code: %v", resp.StatusCode)
		}
	}
}

func (c *Client) post(plugins ...plugin.Plugin) error {
	var req = c.httpClient.Post()
	req = applyPlugins(req, c.accessToken, plugins...)

	var resp *gentleman.Response
	{
		var err error
		resp, err = req.Do()
		if err != nil {
			return errors.Wrap(err, "could not get response")
		}

		switch {
		case resp.StatusCode == http.StatusUnauthorized:
			// If the token is not valid (expired, ...) ask a new one.
			if err = c.verifyToken(); err != nil {
				var err = c.getToken()
				if err != nil {
					return errors.Wrap(err, "could not get token")
				}
			}
			return c.post(plugins...)
		case resp.StatusCode >= 400:
			return fmt.Errorf("invalid status code: '%v': %v", resp.RawResponse.Status, string(resp.Bytes()))
		case resp.StatusCode >= 200:
			return nil
		default:
			return fmt.Errorf("unknown response status code: %v", resp.StatusCode)
		}
	}
}

func (c *Client) delete(plugins ...plugin.Plugin) error {
	var req = c.httpClient.Delete()
	req = applyPlugins(req, c.accessToken, plugins...)

	var resp *gentleman.Response
	{
		var err error
		resp, err = req.Do()
		if err != nil {
			return errors.Wrap(err, "could not get response")
		}

		switch {
		case resp.StatusCode == http.StatusUnauthorized:
			// If the token is not valid (expired, ...) ask a new one.
			if err = c.verifyToken(); err != nil {
				var err = c.getToken()
				if err != nil {
					return errors.Wrap(err, "could not get token")
				}
			}
			return c.delete(plugins...)
		case resp.StatusCode >= 400:
			return fmt.Errorf("invalid status code: '%v': %v", resp.RawResponse.Status, string(resp.Bytes()))
		case resp.StatusCode >= 200:
			return nil
		default:
			return fmt.Errorf("unknown response status code: %v", resp.StatusCode)
		}
	}
}

func (c *Client) put(plugins ...plugin.Plugin) error {
	var req = c.httpClient.Put()
	req = applyPlugins(req, c.accessToken, plugins...)

	var resp *gentleman.Response
	{
		var err error
		resp, err = req.Do()
		if err != nil {
			return errors.Wrap(err, "could not get response")
		}

		switch {
		case resp.StatusCode == http.StatusUnauthorized:
			// If the token is not valid (expired, ...) ask a new one.
			if err = c.verifyToken(); err != nil {
				var err = c.getToken()
				if err != nil {
					return errors.Wrap(err, "could not get token: %v")
				}
			}
			return c.put(plugins...)
		case resp.StatusCode >= 400:
			return fmt.Errorf("invalid status code: '%v': %v", resp.RawResponse.Status, string(resp.Bytes()))
		case resp.StatusCode >= 200:
			return nil
		default:
			return fmt.Errorf("unknown response status code: %v", resp.StatusCode)
		}
	}
}

// applyPlugins apply all the plugins to the request req.
func applyPlugins(req *gentleman.Request, accessToken string, plugins ...plugin.Plugin) *gentleman.Request {
	var r = req.SetHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	for _, p := range plugins {
		r = r.Use(p)
	}
	return r
}

// createQueryPlugins create query parameters with the key values paramKV.
func createQueryPlugins(paramKV ...string) []plugin.Plugin {
	var plugins = []plugin.Plugin{}
	for i := 0; i < len(paramKV); i += 2 {
		var k = paramKV[i]
		var v = paramKV[i+1]
		plugins = append(plugins, query.Set(k, v))
	}
	return plugins
}

func str(s string) *string {
	return &s
}
