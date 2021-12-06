package toolbox

import (
	"errors"
	"fmt"
	"strings"

	"github.com/cloudtrust/keycloak-client/v2"
)

type kcURIProvider struct {
	defaultKey string
	entries    map[string]string
}

func NewKeycloakURIProviderFromArray(uris []string) (keycloak.KeycloakURIProvider, error) {
	var entries = make(map[string]string)
	var defaultKey = "default"
	for idx, value := range uris {
		if idx == 0 {
			entries[defaultKey] = value
		} else {
			entries[fmt.Sprintf("entry-%d", idx)] = value
		}
	}
	return NewKeycloakURIProvider(entries, defaultKey)
}

func NewKeycloakURIProvider(entries map[string]string, defaultKey string) (keycloak.KeycloakURIProvider, error) {
	if len(entries) == 0 {
		return nil, errors.New("entries should not be empty")
	}
	entries = toLowerKeys(entries)
	defaultKey = strings.ToLower(defaultKey)
	if _, ok := entries[defaultKey]; !ok {
		return nil, errors.New("defaultKey is not an entry of the provided entries")
	}

	return &kcURIProvider{
		defaultKey: defaultKey,
		entries:    entries,
	}, nil
}

func toLowerKeys(entries map[string]string) map[string]string {
	var res = make(map[string]string)
	for key, value := range entries {
		res[strings.ToLower(key)] = value
	}
	return res
}

func (kup *kcURIProvider) GetDefaultKey() string {
	return kup.defaultKey
}

func (kup *kcURIProvider) GetAllBaseURIs() []string {
	var res = []string{kup.entries[kup.defaultKey]}
	for _, value := range kup.entries {
		if value != res[0] {
			res = append(res, value)
		}
	}
	return res
}

func (kup *kcURIProvider) GetBaseURI(realmName string) string {
	if value, ok := kup.entries[strings.ToLower(realmName)]; ok {
		return value
	}
	return kup.entries[kup.defaultKey]
}

func (kup *kcURIProvider) ForEachTokenURI(callback func(realm, tokenURI string)) {
	for realm, baseURI := range kup.entries {
		callback(realm, baseURI+"/auth/realms/%s/protocol/openid-connect/token")
	}
}

func ImportLegacyAddrTokenProvider(c *keycloak.Config) error {
	if c.URIProvider != nil {
		return nil
	}
	var err error
	c.URIProvider, err = NewKeycloakURIProviderFromArray(c.AddrTokenProvider)
	return err
}
