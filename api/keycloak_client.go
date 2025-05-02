package api

import (
	"encoding/json"
	"regexp"
	"strings"

	"fmt"
	"net/http"
	"net/url"

	commonhttp "github.com/cloudtrust/common-service/v2/errors"
	"github.com/cloudtrust/keycloak-client/v2"
	"github.com/cloudtrust/keycloak-client/v2/toolbox"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugin"
	"gopkg.in/h2non/gentleman.v2/plugins/query"
	"gopkg.in/h2non/gentleman.v2/plugins/timeout"
)

// Client is the keycloak client.
type Client struct {
	apiURL        *url.URL
	httpClient    *gentleman.Client
	account       *AccountClient
	issuerManager toolbox.IssuerManager
}

// AccountClient structure
type AccountClient struct {
	client *Client
}

const (
	msgErrUnkownHTTPContentType = "unkownHTTPContentType"
)

/*********************************************************/
/*** START: copy of internal_response.go in httpclient ***/

// Internal Response is the only way to let ReadContent be testable
type internalResponse struct {
	gentlemanResponse *gentleman.Response
	bytes             []byte
}

func buildInternalResponse(resp *gentleman.Response) *internalResponse {
	return &internalResponse{
		gentlemanResponse: resp,
		bytes:             nil,
	}
}

func (ir *internalResponse) StatusCode() int {
	return ir.gentlemanResponse.StatusCode
}

func (ir *internalResponse) GetHeader(name string) string {
	return ir.gentlemanResponse.Header.Get(name)
}

func (ir *internalResponse) Bytes() []byte {
	if ir.bytes == nil {
		ir.bytes = ir.gentlemanResponse.Bytes()
	}
	return ir.bytes
}

func (ir *internalResponse) JSON(data any) error {
	return json.Unmarshal(ir.Bytes(), data)
}

func (ir *internalResponse) String() string {
	return string(ir.Bytes())
}

/*** END: copy of internal_response.go in httpclient ***/
/*******************************************************/

// New returns a keycloak client.
func New(config keycloak.Config) (*Client, error) {
	var issuerMgr toolbox.IssuerManager
	{
		var err error
		issuerMgr, err = toolbox.NewIssuerManager(config)
		if err != nil {
			return nil, errors.Wrap(err, keycloak.MsgErrCannotParse+"."+keycloak.TokenProviderURL)
		}
	}

	var uAPI *url.URL
	{
		var err error
		uAPI, err = url.Parse(config.AddrAPI)
		if err != nil {
			return nil, errors.Wrap(err, keycloak.MsgErrCannotParse+"."+keycloak.APIURL)
		}
	}

	var httpClient = gentleman.New()
	{
		httpClient = httpClient.URL(uAPI.String())
		httpClient = httpClient.Use(timeout.Request(config.Timeout))
	}

	var client = &Client{
		apiURL:        uAPI,
		httpClient:    httpClient,
		issuerManager: issuerMgr,
	}

	client.account = &AccountClient{
		client: client,
	}

	return client, nil
}

// VerifyToken verifies a token. It returns an error it is malformed, expired,...
func (c *Client) VerifyToken(issuer string, realmName string, accessToken string) error {
	oidcVerifierProvider, err := c.issuerManager.GetOidcVerifierProvider(issuer)
	if err != nil {
		return err
	}

	verifier, err := oidcVerifierProvider.GetOidcVerifier(realmName)
	if err != nil {
		return err
	}
	return verifier.Verify(accessToken)
}

// AccountClient gets the associated AccountClient
func (c *Client) AccountClient() *AccountClient {
	return c.account
}

func (c *Client) checkError(resp *internalResponse) error {
	switch {
	case resp.StatusCode() == http.StatusUnauthorized:
		return keycloak.ClientDetailedError{HTTPStatus: http.StatusUnauthorized, Message: string(resp.Bytes())}
	case resp.StatusCode() >= 400:
		return treatErrorStatus(resp)
	case resp.StatusCode() >= 200:
		return nil
	default:
		return fmt.Errorf("%s.%v", keycloak.MsgErrUnknownResponseStatusCode, resp.StatusCode())
	}
}

func (c *Client) readContent(resp *internalResponse, data any) (retError error) {
	defer func() {
		if err := recover(); err != nil {
			retError = fmt.Errorf("Unexpected panic. Ensure data is declared with the expected type: %v", err)
		}
	}()
	var hdr = resp.GetHeader("Content-Type")
	switch strings.Split(hdr, ";")[0] {
	case "application/json":
		retError = resp.JSON(data)
	case "text/plain":
		*(data.(*string)) = resp.String()
		retError = nil
	case "text/html":
		*(data.(*string)) = resp.String()
		retError = nil
	case "application/octet-stream", "application/zip", "application/pdf", "text/xml":
		*(data.(*[]byte)) = resp.Bytes()
		retError = nil
	default:
		if len(resp.Bytes()) == 0 {
			retError = nil
		} else {
			retError = fmt.Errorf("%s.%v", msgErrUnkownHTTPContentType, hdr)
		}
	}
	return retError
}

// get is a HTTP get method.
func (c *Client) get(accessToken string, data any, plugins ...plugin.Plugin) error {
	var err error
	var req = c.httpClient.Get()
	req = applyPlugins(req, plugins...)
	req, err = setAuthorisationAndHostHeaders(req, accessToken)

	if err != nil {
		return err
	}

	var gresp *gentleman.Response
	{
		var err error
		gresp, err = req.Do()
		if err != nil {
			return errors.Wrap(err, keycloak.MsgErrCannotObtain+"."+keycloak.Response)
		}

		var resp = buildInternalResponse(gresp)
		err = c.checkError(resp)
		if err != nil {
			return err
		}
		return c.readContent(resp, data)
	}
}

func (c *Client) post(accessToken string, data any, plugins ...plugin.Plugin) (string, error) {
	var err error
	var req = c.httpClient.Post()
	req = applyPlugins(req, plugins...)
	req, err = setAuthorisationAndHostHeaders(req, accessToken)

	if err != nil {
		return "", err
	}

	var gresp *gentleman.Response
	{
		var err error
		gresp, err = req.Do()
		if err != nil {
			return "", errors.Wrap(err, keycloak.MsgErrCannotObtain+"."+keycloak.Response)
		}
		var resp = buildInternalResponse(gresp)

		err = c.checkError(resp)
		if err != nil {
			return "", err
		}
		return resp.GetHeader("Location"), c.readContent(resp, data)
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
			return errors.Wrap(err, keycloak.MsgErrCannotObtain+"."+keycloak.Response)
		}

		return c.checkError(buildInternalResponse(resp))
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
			return errors.Wrap(err, keycloak.MsgErrCannotObtain+"."+keycloak.Response)
		}

		return c.checkError(buildInternalResponse(resp))
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
			return "", errors.Wrap(err, keycloak.MsgErrCannotParse+"."+keycloak.TokenProviderURL)
		}
	}

	return u.Host, nil
}

func extractIssuerFromToken(tokenStr string) (string, error) {
	token, _, err := jwt.NewParser().ParseUnverified(tokenStr, jwt.MapClaims{})
	if err != nil {
		return "", errors.Wrap(err, keycloak.MsgErrCannotParse+"."+keycloak.TokenMsg)
	}

	issuer, err := token.Claims.GetIssuer()
	if err != nil {
		return "", errors.Wrap(err, keycloak.MsgErrCannotGetIssuer+"."+keycloak.TokenMsg)
	}

	return issuer, nil
}

// createQueryPlugins create query parameters with the key values paramKV.
func createQueryPlugins(paramKV ...string) []plugin.Plugin {
	var plugins = []plugin.Plugin{}
	for i := 0; i < len(paramKV); i += 2 {
		var k = paramKV[i]
		var v = paramKV[i+1]
		plugins = append(plugins, query.Add(k, v))
	}
	return plugins
}

func treatErrorStatus(resp *internalResponse) error {
	var response map[string]any
	err := json.Unmarshal(resp.Bytes(), &response)
	if message, ok := response["errorMessage"]; ok && err == nil {
		return whitelistErrors(resp.StatusCode(), message.(string))
	}
	return keycloak.HTTPError{
		HTTPStatus: resp.StatusCode(),
		Message:    string(resp.Bytes()),
	}
}

func whitelistErrors(statusCode int, message string) error {
	// whitelist errors from Keycloak
	reg := regexp.MustCompile("invalidPassword[a-zA-Z]*Message")
	errorMessages := map[string]string{
		"User exists with same username or email": keycloak.MsgErrExistingValue + "." + keycloak.UserOrEmail,
		"usernameExistsMessage":                   keycloak.MsgErrExistingValue + "." + keycloak.UserOrEmail,
		"emailExistsMessage":                      keycloak.MsgErrExistingValue + "." + keycloak.UserOrEmail,
		"User exists with same username":          keycloak.MsgErrExistingValue + "." + keycloak.Username,
		"User exists with same email":             keycloak.MsgErrExistingValue + "." + keycloak.Email,
		"readOnlyUsernameMessage":                 keycloak.MsgErrReadOnly + "." + keycloak.Username,
	}

	switch {
	//POST account/credentials/password with error message related to invalid value for the password
	// of the format invalidPassword{a-zA-Z}*Message, e.g. invalidPasswordMinDigitsMessage
	case reg.MatchString(message):
		return commonhttp.Error{
			Status:  statusCode,
			Message: "keycloak." + message,
		}
	// update account in back-office or self-service
	case errorMessages[message] != "":
		return commonhttp.Error{
			Status:  statusCode,
			Message: "keycloak." + errorMessages[message],
		}
	default:
		return keycloak.ClientDetailedError{
			HTTPStatus: statusCode,
			Message:    message,
		}
	}
}

// Token is JWT token.
// We need to define our own structure as the library define aud as a string but it can also be a string array.
// To fix this issue, we remove aud as we do not use it here.
type Token struct {
	hdr            *header
	Issuer         string `json:"iss,omitempty"`
	Subject        string `json:"sub,omitempty"`
	ExpirationTime int64  `json:"exp,omitempty"`
	NotBefore      int64  `json:"nbf,omitempty"`
	IssuedAt       int64  `json:"iat,omitempty"`
	ID             string `json:"jti,omitempty"`
	Username       string `json:"preferred_username,omitempty"`
}

type header struct {
	Algorithm   string `json:"alg,omitempty"`
	KeyID       string `json:"kid,omitempty"`
	Type        string `json:"typ,omitempty"`
	ContentType string `json:"cty,omitempty"`
}
