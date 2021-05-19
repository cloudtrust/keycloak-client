package keycloak

import (
	"time"
)

// Config is the keycloak client http config.
type Config struct {
	AddrTokenProvider []string
	AddrAPI           string
	Timeout           time.Duration
	CacheTTL          time.Duration
	ErrorTolerance    time.Duration
}
