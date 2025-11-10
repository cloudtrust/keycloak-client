package toolbox

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/cloudtrust/keycloak-client/v2"
)

// Errors definition
var (
	ErrConfigKeyNotFound   = errors.New("component config key not found")
	ErrInvalidConfigFormat = errors.New("invalid config encoding")
)

// ComponentTool interface
type ComponentTool interface {
	FindComponent(components []keycloak.ComponentRepresentation) *keycloak.ComponentRepresentation
	InitializeComponent(realmName string, idpID string, initial any) (keycloak.ComponentRepresentation, error)
	GetComponentEntry(comp *keycloak.ComponentRepresentation, key string, out any) error
	UpdateComponentEntry(comp *keycloak.ComponentRepresentation, key string, value any) error
	DeleteComponentEntry(comp *keycloak.ComponentRepresentation, key string) (bool, error)

	GetProviderType() string
}

// ComponentConfig struct
type ComponentConfig struct {
	ProviderType string `mapstructure:"provider-type"`
	ProviderID   string `mapstructure:"provider-id"`
	ConfigName   string `mapstructure:"config-name"`
}

// GenericComponentTool struct
type GenericComponentTool struct {
	ProviderType string
	ProviderID   string
	ConfigName   string
}

// NewComponentTool creates
func NewComponentTool(config ComponentConfig) ComponentTool {
	return &GenericComponentTool{
		ProviderType: config.ProviderType,
		ProviderID:   config.ProviderID,
		ConfigName:   config.ConfigName,
	}
}

// ComponentConfigKeyValue struct
type ComponentConfigKeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// GetProviderType returns the provider type
func (ct *GenericComponentTool) GetProviderType() string {
	return ct.ProviderType
}

// FindComponent searches a component
func (ct *GenericComponentTool) FindComponent(components []keycloak.ComponentRepresentation) *keycloak.ComponentRepresentation {
	for i := range components {
		c := &components[i]
		if c.ProviderID != nil && *c.ProviderID == ct.ProviderID {
			return c
		}
	}

	return nil
}

// InitializeComponent initializes a component
func (ct *GenericComponentTool) InitializeComponent(parentID string, idpID string, initial any) (keycloak.ComponentRepresentation, error) {
	valueJSON, err := json.Marshal(initial)
	if err != nil {
		return keycloak.ComponentRepresentation{}, err
	}

	arr := []ComponentConfigKeyValue{{
		Key:   idpID,
		Value: string(valueJSON),
	}}

	arrJSON, err := json.Marshal(arr)
	if err != nil {
		return keycloak.ComponentRepresentation{}, err
	}

	config := map[string][]string{
		ct.ConfigName: {string(arrJSON)},
	}

	return keycloak.ComponentRepresentation{
		ProviderID:   &ct.ProviderID,
		ProviderType: &ct.ProviderType,
		ParentID:     &parentID,
		Config:       config,
	}, nil
}

func (ct *GenericComponentTool) loadConfigArray(comp *keycloak.ComponentRepresentation) ([]ComponentConfigKeyValue, error) {
	if comp == nil || comp.Config == nil {
		return nil, errors.New("no config found")
	}

	entries := comp.Config[ct.ConfigName]
	if len(entries) == 0 {
		return nil, ErrInvalidConfigFormat
	}

	var arr []ComponentConfigKeyValue
	if err := json.Unmarshal([]byte(entries[0]), &arr); err != nil {
		return nil, ErrInvalidConfigFormat
	}

	return arr, nil
}

func (ct *GenericComponentTool) saveConfigArray(comp *keycloak.ComponentRepresentation, arr []ComponentConfigKeyValue) error {
	data, err := json.Marshal(arr)
	if err != nil {
		return err
	}
	comp.Config[ct.ConfigName] = []string{string(data)}

	return nil
}

func (ct *GenericComponentTool) getComponentConfigKeyValue(comp *keycloak.ComponentRepresentation, key string) (ComponentConfigKeyValue, error) {
	arr, err := ct.loadConfigArray(comp)
	if err != nil {
		return ComponentConfigKeyValue{}, err
	}
	for _, kv := range arr {
		if kv.Key == key {
			return kv, nil
		}
	}

	return ComponentConfigKeyValue{}, ErrConfigKeyNotFound
}

func (ct *GenericComponentTool) updateComponentConfigKeyValue(comp *keycloak.ComponentRepresentation, cckv ComponentConfigKeyValue) error {
	arr, err := ct.loadConfigArray(comp)
	if errors.Is(err, ErrInvalidConfigFormat) || errors.Is(err, ErrConfigKeyNotFound) {
		arr = nil
	} else if err != nil {
		return err
	}

	// Update or append
	found := false
	for i, kv := range arr {
		if kv.Key == cckv.Key {
			arr[i] = cckv
			found = true
			break
		}
	}
	if !found {
		arr = append(arr, cckv)
	}

	return ct.saveConfigArray(comp, arr)
}

// GetComponentEntry decodes a stored JSON value into the provided struct.
func (ct *GenericComponentTool) GetComponentEntry(comp *keycloak.ComponentRepresentation, key string, out any) error {
	cckv, err := ct.getComponentConfigKeyValue(comp, key)
	if err != nil {
		return err
	}
	if err := json.Unmarshal([]byte(cckv.Value), out); err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidConfigFormat, err)
	}

	return nil
}

// UpdateComponentEntry encodes a struct and stores it under the key.
func (ct *GenericComponentTool) UpdateComponentEntry(comp *keycloak.ComponentRepresentation, key string, value any) error {
	valueJSON, err := json.Marshal(value)
	if err != nil {
		return err
	}

	kv := ComponentConfigKeyValue{
		Key:   key,
		Value: string(valueJSON),
	}

	return ct.updateComponentConfigKeyValue(comp, kv)
}

// DeleteComponentEntry deletes a component entry
func (ct *GenericComponentTool) DeleteComponentEntry(comp *keycloak.ComponentRepresentation, key string) (bool, error) {
	arr, err := ct.loadConfigArray(comp)
	if err != nil {
		return false, err
	}

	newArr := make([]ComponentConfigKeyValue, 0, len(arr))
	deleted := false
	for _, kv := range arr {
		if kv.Key == key {
			deleted = true
			continue
		}
		newArr = append(newArr, kv)
	}
	if !deleted {
		return false, nil
	}

	if err = ct.saveConfigArray(comp, newArr); err != nil {
		return false, err
	}

	return true, nil
}
