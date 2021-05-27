package keycloak

import (
	"time"
)

type KeycloakURIProvider interface {
	GetDefaultKey() string
	GetAllBaseURIs() []string
	ForEachTokenURI(callback func(realm, tokenURI string))
}

// Config is the keycloak client http config.
type Config struct {
	// AddrTokenProvider is deprecated. Please prefer using URIProvider
	AddrTokenProvider []string
	URIProvider       KeycloakURIProvider
	AddrAPI           string
	Timeout           time.Duration
	CacheTTL          time.Duration
	ErrorTolerance    time.Duration
}
