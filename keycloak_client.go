package keycloak

import (
	"context"
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

	jwt "github.com/gbrlsnchs/jwt"
)

// Config is the keycloak client http config.
type Config struct {
	AddrTokenProvider string
	AddrAPI           string
	Timeout           time.Duration
}

// Client is the keycloak client.
type Client struct {
	tokenProviderUrl *url.URL
	apiUrl           *url.URL
	httpClient       *gentleman.Client
}

// New returns a keycloak client.
func New(config Config) (*Client, error) {
	var uToken *url.URL
	{
		var err error
		uToken, err = url.Parse(config.AddrTokenProvider)
		if err != nil {
			return nil, errors.Wrap(err, "could not parse Token Provider URL")
		}
	}

	var uAPI *url.URL
	{
		var err error
		uAPI, err = url.Parse(config.AddrAPI)
		if err != nil {
			return nil, errors.Wrap(err, "could not parse API URL")
		}
	}

	var httpClient = gentleman.New()
	{
		httpClient = httpClient.URL(uAPI.String())
		httpClient = httpClient.Use(timeout.Request(config.Timeout))
	}

	return &Client{
		tokenProviderUrl: uToken,
		apiUrl:           uAPI,
		httpClient:       httpClient,
	}, nil
}

// getToken returns a valid token from keycloak.
func (c *Client) GetToken(realm string, username string, password string) (string, error) {
	var req *gentleman.Request
	{
		var authPath = fmt.Sprintf("/auth/realms/%s/protocol/openid-connect/token", realm)
		req = c.httpClient.Post()
		req = req.SetHeader("Content-Type", "application/x-www-form-urlencoded")
		req = req.Path(authPath)
		req = req.Type("urlencoded")
		req = req.BodyString(fmt.Sprintf("username=%s&password=%s&grant_type=password&client_id=admin-cli", username, password))
	}

	var resp *gentleman.Response
	{
		var err error
		resp, err = req.Do()
		if err != nil {
			return "", errors.Wrap(err, "could not get token")
		}
	}
	defer resp.Close()

	var unmarshalledBody map[string]interface{}
	{
		var err error
		err = resp.JSON(&unmarshalledBody)
		if err != nil {
			return "", errors.Wrap(err, "could not unmarshal response")
		}
	}

	var accessToken interface{}
	{
		var ok bool
		accessToken, ok = unmarshalledBody["access_token"]
		if !ok {
			return "", fmt.Errorf("could not find access token in response body")
		}
	}

	fmt.Printf("%s", accessToken.(string))
	fmt.Println()

	return accessToken.(string), nil
}

// verifyToken token verify a token. It returns an error it is malformed, expired,...
func (c *Client) VerifyToken(realmName string, accessToken string) error {
	var oidcProvider *oidc.Provider
	{
		var err error
		var issuer = fmt.Sprintf("%s/auth/realms/%s", c.tokenProviderUrl.String(), realmName)
		oidcProvider, err = oidc.NewProvider(context.Background(), issuer)
		if err != nil {
			return errors.Wrap(err, "could not create oidc provider")
		}
	}

	var v = oidcProvider.Verifier(&oidc.Config{SkipClientIDCheck: true})

	var err error
	_, err = v.Verify(context.Background(), accessToken)
	return err
}

// get is a HTTP get method.
func (c *Client) get(accessToken string, data interface{}, plugins ...plugin.Plugin) error {
	var err error
	var req = c.httpClient.Get()
	req = applyPlugins(req, plugins...)
	req, err = setAuthorisationAndHostHeaders(req, accessToken)

	if err != nil {
		return err
	}

	var resp *gentleman.Response
	{
		var err error
		resp, err = req.Do()
		if err != nil {
			return errors.Wrap(err, "could not get response")
		}

		switch {
		case resp.StatusCode == http.StatusUnauthorized:
			return fmt.Errorf("unauthorized request: '%v': %v", resp.RawResponse.Status, string(resp.Bytes()))
		case resp.StatusCode >= 400:
			return fmt.Errorf("invalid status code: '%v': %v", resp.RawResponse.Status, string(resp.Bytes()))
		case resp.StatusCode >= 200:
			switch resp.Header.Get("Content-Type") {
			case "application/json":
				return resp.JSON(data)
			case "application/octet-stream":
				data = resp.Bytes()
				return nil
			default:
				return fmt.Errorf("unkown http content-type: %v", resp.Header.Get("Content-Type"))
			}
		default:
			return fmt.Errorf("unknown response status code: %v", resp.StatusCode)
		}
	}
}

func (c *Client) post(accessToken string, data interface{}, plugins ...plugin.Plugin) (string, error) {
	var err error
	var req = c.httpClient.Post()
	req = applyPlugins(req, plugins...)
	req, err = setAuthorisationAndHostHeaders(req, accessToken)

	if err != nil {
		return "", err
	}

	var resp *gentleman.Response
	{
		var err error
		resp, err = req.Do()
		if err != nil {
			return "", errors.Wrap(err, "could not get response")
		}

		switch {
		case resp.StatusCode == http.StatusUnauthorized:
			return "", fmt.Errorf("unauthorized request: '%v': %v", resp.RawResponse.Status, string(resp.Bytes()))
		case resp.StatusCode >= 400:
			return "", fmt.Errorf("invalid status code: '%v': %v", resp.RawResponse.Status, string(resp.Bytes()))
		case resp.StatusCode >= 200:
			var location = resp.Header.Get("Location")

			switch resp.Header.Get("Content-Type") {
			case "application/json":
				return location, resp.JSON(data)
			case "application/octet-stream":
				data = resp.Bytes()
				return location, nil
			default:
				return location, nil
			}
		default:
			return "", fmt.Errorf("unknown response status code: %v", resp.StatusCode)
		}
	}
}

func (c *Client) delete(accessToken string, plugins ...plugin.Plugin) error {
	var err error
	var req = c.httpClient.Delete()
	req = applyPlugins(req, plugins...)
	req, err = setAuthorisationAndHostHeaders(req, accessToken)

	if err != nil {
		return err
	}

	var resp *gentleman.Response
	{
		var err error
		resp, err = req.Do()
		if err != nil {
			return errors.Wrap(err, "could not get response")
		}

		switch {
		case resp.StatusCode == http.StatusUnauthorized:
			return fmt.Errorf("unauthorized request: '%v': %v", resp.RawResponse.Status, string(resp.Bytes()))
		case resp.StatusCode >= 400:
			return fmt.Errorf("invalid status code: '%v': %v", resp.RawResponse.Status, string(resp.Bytes()))
		case resp.StatusCode >= 200:
			return nil
		default:
			return fmt.Errorf("unknown response status code: %v", resp.StatusCode)
		}
	}
}

func (c *Client) put(accessToken string, plugins ...plugin.Plugin) error {
	var err error
	var req = c.httpClient.Put()
	req = applyPlugins(req, plugins...)
	req, err = setAuthorisationAndHostHeaders(req, accessToken)

	if err != nil {
		return err
	}

	var resp *gentleman.Response
	{
		var err error
		resp, err = req.Do()
		if err != nil {
			return errors.Wrap(err, "could not get response")
		}

		switch {
		case resp.StatusCode == http.StatusUnauthorized:
			return fmt.Errorf("unauthorized request: '%v': %v", resp.RawResponse.Status, string(resp.Bytes()))
		case resp.StatusCode >= 400:
			return fmt.Errorf("invalid status code: '%v': %v", resp.RawResponse.Status, string(resp.Bytes()))
		case resp.StatusCode >= 200:
			return nil
		default:
			return fmt.Errorf("unknown response status code: %v", resp.StatusCode)
		}
	}
}

func setAuthorisationAndHostHeaders(req *gentleman.Request, accessToken string) (*gentleman.Request, error) {
	host, err := extractHostFromToken(accessToken)

	if err != nil {
		return req, err
	}

	var r = req.SetHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	r = r.SetHeader("X-Forwarded-Proto", "https")

	r.Context.Request.Host = host

	return r, nil
}

// applyPlugins apply all the plugins to the request req.
func applyPlugins(req *gentleman.Request, plugins ...plugin.Plugin) *gentleman.Request {
	var r = req
	for _, p := range plugins {
		r = r.Use(p)
	}
	return r
}

func extractHostFromToken(token string) (string, error) {
	issuer, err := extractIssuerFromToken(token)

	if err != nil {
		return "", err
	}

	var u *url.URL
	{
		var err error
		u, err = url.Parse(issuer)
		if err != nil {
			return "", errors.Wrap(err, "could not parse Token issuer URL")
		}
	}

	return u.Host, nil
}

func extractIssuerFromToken(token string) (string, error) {
	payload, _, err := jwt.Parse(token)

	if err != nil {
		return "", fmt.Errorf("Error while parsing token")
	}

	var jot jwt.JWT

	if err = jwt.Unmarshal(payload, &jot); err != nil {
		return "", fmt.Errorf("Error while unmarshalling token")
	}

	return jot.Issuer, nil
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
