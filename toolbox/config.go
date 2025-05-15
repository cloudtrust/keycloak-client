package toolbox

import (
	"errors"
	"time"

	"github.com/cloudtrust/keycloak-client/v2"
)

// InternalConfig struct
type InternalConfig struct {
	InternalURI    string            `mapstructure:"internal-uri"`
	RealmPublicURI map[string]string `mapstructure:"realm-public-uri-map"`
	DefaultKey     *string           `mapstructure:"default-key"`
	Timeout        time.Duration     `mapstructure:"timeout"`
}

// ConfigurationProvider interface
type ConfigurationProvider func(target any) error

// NewConfig returns a Keycloak configuration
func NewConfig(confUnmarshal ConfigurationProvider) (keycloak.Config, error) {
	var config InternalConfig
	var err = confUnmarshal(&config)
	if err != nil {
		return keycloak.Config{}, errors.New("Cannot get keycloak configuration")
	}
	var defaultKey = "default"
	if config.DefaultKey != nil {
		defaultKey = *config.DefaultKey
	}
	var uriProvider keycloak.KeycloakURIProvider
	uriProvider, err = NewKeycloakURIProvider(config.RealmPublicURI, defaultKey)
	if err != nil {
		return keycloak.Config{}, err
	}
	return keycloak.Config{
		URIProvider:     uriProvider,
		AddrInternalAPI: config.InternalURI,
		Timeout:         config.Timeout,
	}, err
}
