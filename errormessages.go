package keycloak

import "strconv"

// Constants for error management
const (
	MsgErrMissingParam              = "missingParameter"
	MsgErrInvalidParam              = "invalidParameter"
	MsgErrCannotObtain              = "cannotObtain"
	MsgErrCannotMarshal             = "cannotMarshal"
	MsgErrCannotUnmarshal           = "cannotUnmarshal"
	MsgErrCannotParse               = "cannotParse"
	MsgErrCannotCreate              = "cannotCreate"
	MsgErrUnkownHTTPContentType     = "unkownHTTPContentType"
	MsgErrUnknownResponseStatusCode = "unknownResponseStatusCode"
	MsgErrExistingValue             = "existing"
	MsgErrReadOnly                  = "readOnlyValue"

	EvenParams       = "key/valParametersShouldBeEven"
	TokenProviderURL = "tokenProviderURL"
	APIURL           = "APIURL"
	TokenMsg         = "token"
	Response         = "response"
	AccessToken      = "accessToken"
	OIDCProvider     = "OIDCProvider"
	UserOrEmail      = "UsernameOrEmail"
	Username         = "username"
	Email            = "email"
)

// HTTPError is returned when an error occured while contacting the keycloak instance.
type HTTPError struct {
	HTTPStatus int
	Message    string
}

func (e HTTPError) Error() string {
	return strconv.Itoa(e.HTTPStatus) + ":" + e.Message
}
