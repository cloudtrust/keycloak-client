package keycloak

import (
	"time"
)

// KeycloakURIProvider interface
type KeycloakURIProvider interface {
	GetDefaultKey() string
	GetAllBaseURIs() []string
	GetBaseURI(realmName string) string
	ForEachContextURI(callback func(realm, host, baseURI string))
}

// Config is the keycloak client http config.
type Config struct {
	// AddrTokenProvider is deprecated. Please prefer using URIProvider
	AddrTokenProvider []string
	URIProvider       KeycloakURIProvider
	AddrAPI           string
	Timeout           time.Duration
}
