package keycloak

import (
	"time"
)

type KeycloakURIProvider interface {
	GetDefaultKey() string
	GetAllBaseURIs() []string
	GetBaseURI(realmName string) string
	ForEachTokenURI(callback func(realm, tokenURI string))
}

// Config is the keycloak client http config.
type Config struct {
	// AddrTokenProvider is deprecated. Please prefer using URIProvider
	AddrTokenProvider []string
	URIProvider       KeycloakURIProvider
	AddrAPI           string
	Timeout           time.Duration
}
