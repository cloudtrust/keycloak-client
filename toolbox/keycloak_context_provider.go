package toolbox

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/cloudtrust/keycloak-client/v2"
)

type kcContextEntry struct {
	host    string
	baseURI string
}

type kcContextProvider struct {
	defaultKey string
	entries    map[string]kcContextEntry
}

// NewKeycloakURIProviderFromArray creates a Keycloak URI provider
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

// NewKeycloakURIProvider creates a Keycloak URI provider
func NewKeycloakURIProvider(entries map[string]string, defaultKey string) (keycloak.KeycloakURIProvider, error) {
	if len(entries) == 0 {
		return nil, errors.New("entries should not be empty")
	}
	entries = toLowerKeys(entries)
	defaultKey = strings.ToLower(defaultKey)
	if _, ok := entries[defaultKey]; !ok {
		return nil, errors.New("defaultKey is not an entry of the provided entries")
	}
	var kcEntries, err = toKcEntries(entries)
	if err != nil {
		return nil, err
	}

	return &kcContextProvider{
		defaultKey: defaultKey,
		entries:    kcEntries,
	}, nil
}

func toLowerKeys(entries map[string]string) map[string]string {
	var res = make(map[string]string)
	for key, value := range entries {
		res[strings.ToLower(key)] = value
	}
	return res
}

func toKcEntries(entries map[string]string) (map[string]kcContextEntry, error) {
	var res = map[string]kcContextEntry{}
	for key, baseURI := range entries {
		var host, err = extractHostFromURL(baseURI)
		if err != nil {
			return nil, err
		}
		res[key] = kcContextEntry{
			host:    host,
			baseURI: baseURI,
		}
	}
	return res, nil
}

func extractHostFromURL(anURL string) (string, error) {
	var u *url.URL
	{
		var err error
		u, err = url.Parse(anURL)
		if err != nil || u.Host == "" {
			return "", errors.New("Can't parse URL " + anURL)
		}
	}

	return u.Host, nil
}

func (kup *kcContextProvider) GetDefaultKey() string {
	return kup.defaultKey
}

func (kup *kcContextProvider) GetAllBaseURIs() []string {
	var res = []string{kup.entries[kup.defaultKey].baseURI}
	for _, value := range kup.entries {
		if value.baseURI != res[0] {
			res = append(res, value.baseURI)
		}
	}
	return res
}

func (kup *kcContextProvider) GetBaseURI(realmName string) string {
	if value, ok := kup.entries[strings.ToLower(realmName)]; ok {
		return value.baseURI
	}
	return kup.entries[kup.defaultKey].baseURI
}

func (kup *kcContextProvider) ForEachContextURI(callback func(realm, host, baseURI string)) {
	for realm, entry := range kup.entries {
		callback(realm, entry.host, entry.baseURI)
	}
}
