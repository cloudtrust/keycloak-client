package toolbox

import (
	"fmt"
	"testing"

	"github.com/cloudtrust/keycloak-client/v2"
	"github.com/stretchr/testify/assert"
)

const (
	realmName = "test-community"

	idpID1     = "EXTIDP-12345678-abcd-efgh-ijkl-012345678901"
	idpRanges1 = "192.168.1.0/24"

	idpID2     = "EXTIDP-12345678-abcd-efgh-ijkl-012345678902"
	idpRanges2 = "192.168.2.0/24"

	compID           = "5b3f0a5d-a59d-4aff-8932-aa70f2806f04"
	compProviderType = "org.keycloak.services.ui.extend.UiTabProvider"
	compProviderID   = "Home-realm discovery settings"
	compConfigName   = "hrdSettings"
)

type hrdSettingModel struct {
	IPRangesList string `json:"ipRangesList"`
}

func ptr(value string) *string {
	return &value
}

func createComponentTool() ComponentTool {
	return &GenericComponentTool{
		ProviderType: compProviderType,
		ProviderID:   compProviderID,
		ConfigName:   compConfigName,
	}
}

func testComponent() keycloak.ComponentRepresentation {
	config := map[string][]string{
		compConfigName: {
			fmt.Sprintf(
				`[{"key":"%s","value":"{\"ipRangesList\":\"%s\"}"},{"key":"%s","value":"{\"ipRangesList\":\"%s\"}"}]`,
				idpID1, idpRanges1, idpID2, idpRanges2,
			),
		},
	}
	return keycloak.ComponentRepresentation{
		Config:       config,
		ID:           ptr(compID),
		ParentID:     ptr(realmName),
		ProviderID:   ptr(compProviderID),
		ProviderType: ptr(compProviderType),
	}
}

func TestFindComponent(t *testing.T) {
	tool := createComponentTool()
	comp := testComponent()

	t.Run("No component", func(t *testing.T) {
		res := tool.FindComponent([]keycloak.ComponentRepresentation{})
		assert.Nil(t, res)
	})

	t.Run("No component with this ID", func(t *testing.T) {
		stsComp := comp
		stsComp.ProviderID = ptr("STS settings")

		res := tool.FindComponent([]keycloak.ComponentRepresentation{stsComp})
		assert.Nil(t, res)
	})

	t.Run("Component is present", func(t *testing.T) {
		res := tool.FindComponent([]keycloak.ComponentRepresentation{comp})
		assert.NotNil(t, res)
		assert.Equal(t, comp, *res)
	})
}

func TestInitializeComponent(t *testing.T) {
	tool := createComponentTool()

	t.Run("Marshal of value failed", func(t *testing.T) {
		var initial complex128

		_, err := tool.InitializeComponent(realmName, idpID1, initial)
		assert.NotNil(t, err)
	})

	t.Run("success", func(t *testing.T) {
		initial := hrdSettingModel{
			IPRangesList: idpRanges1,
		}

		expectedConfig := map[string][]string{
			compConfigName: {
				fmt.Sprintf(
					`[{"key":"%s","value":"{\"ipRangesList\":\"%s\"}"}]`,
					idpID1, idpRanges1,
				),
			}}

		res, err := tool.InitializeComponent(realmName, idpID1, initial)
		assert.Nil(t, err)
		assert.Equal(t, compProviderID, *res.ProviderID)
		assert.Equal(t, compProviderType, *res.ProviderType)
		assert.Equal(t, realmName, *res.ParentID)
		assert.Equal(t, expectedConfig, res.Config)
	})
}

func TestGetComponentEntry(t *testing.T) {
	tool := createComponentTool()

	t.Run("Key not found", func(t *testing.T) {
		comp := testComponent()

		var out hrdSettingModel
		err := tool.GetComponentEntry(&comp, "non-existing-idp", &out)
		assert.ErrorIs(t, err, ErrConfigKeyNotFound)
	})

	t.Run("Invalid JSON format", func(t *testing.T) {
		comp := testComponent()

		// Corrupt config
		brokenComp := comp
		brokenComp.Config[compConfigName] = []string{`[{"key":"abc","value":"not-a-json"}]`}

		var out hrdSettingModel
		err := tool.GetComponentEntry(&brokenComp, "abc", &out)
		assert.ErrorIs(t, err, ErrInvalidConfigFormat)
	})

	t.Run("Success", func(t *testing.T) {
		comp := testComponent()

		var out hrdSettingModel
		err := tool.GetComponentEntry(&comp, idpID1, &out)
		assert.Nil(t, err)
		assert.Equal(t, idpRanges1, out.IPRangesList)
	})
}

func TestUpdateComponentEntry(t *testing.T) {
	tool := createComponentTool()

	t.Run("Add new entry", func(t *testing.T) {
		comp := testComponent()

		newModel := hrdSettingModel{IPRangesList: "10.0.0.0/8"}
		err := tool.UpdateComponentEntry(&comp, "new-idp", newModel)
		assert.Nil(t, err)

		var out hrdSettingModel
		err = tool.GetComponentEntry(&comp, "new-idp", &out)
		assert.Nil(t, err)
		assert.Equal(t, "10.0.0.0/8", out.IPRangesList)
	})

	t.Run("Update existing entry", func(t *testing.T) {
		comp := testComponent()

		updatedModel := hrdSettingModel{IPRangesList: "172.16.0.0/12"}
		err := tool.UpdateComponentEntry(&comp, idpID1, updatedModel)
		assert.Nil(t, err)

		var out hrdSettingModel
		err = tool.GetComponentEntry(&comp, idpID1, &out)
		assert.Nil(t, err)
		assert.Equal(t, "172.16.0.0/12", out.IPRangesList)
	})
}

func TestDeleteComponentEntry(t *testing.T) {
	tool := createComponentTool()

	t.Run("Delete existing entry", func(t *testing.T) {
		comp := testComponent()

		ok, err := tool.DeleteComponentEntry(&comp, idpID1)
		assert.Nil(t, err)
		assert.True(t, ok)

		var out hrdSettingModel
		err = tool.GetComponentEntry(&comp, idpID1, &out)
		assert.ErrorIs(t, err, ErrConfigKeyNotFound)
	})

	t.Run("Delete non-existing entry", func(t *testing.T) {
		comp := testComponent()

		ok, err := tool.DeleteComponentEntry(&comp, "does-not-exist")
		assert.Nil(t, err)
		assert.False(t, ok)
	})
}
