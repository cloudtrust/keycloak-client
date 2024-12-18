package keycloak

import (
	"fmt"
)

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
	MsgErrCannotGetIssuer           = "cannotGetIssuer"

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
	return fmt.Sprintf("%d:%s", e.HTTPStatus, e.Message)
}

// ClientDetailedError struct
type ClientDetailedError struct {
	HTTPStatus int
	Message    string
}

// Error implements error
func (e ClientDetailedError) Error() string {
	return fmt.Sprintf("%d:%s", e.HTTPStatus, e.Message)
}

// Status implements common-service/errors/DetailedError
func (e ClientDetailedError) Status() int {
	return e.HTTPStatus
}

// ErrorMessage implements common-service/errors/DetailedError
func (e ClientDetailedError) ErrorMessage() string {
	return e.Message
}
