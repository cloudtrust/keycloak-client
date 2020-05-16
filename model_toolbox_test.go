package keycloak

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	supportedDateLayouts = []string{"02.01.2006", "2006/01/02"}
)

func TestUserRepresentationAttributes(t *testing.T) {
	var userRep UserRepresentation

	var (
		keyMissing  = AttributeKey("missing")
		keyDate     = AttributeKey("date")
		keyGender   = AttributeKey("gender")
		keyMultiple = AttributeKey("multiple")
	)

	// Attributes are empty
	assert.Nil(t, userRep.GetAttribute(keyMissing), "Search with no attribute")
	assert.Nil(t, userRep.GetAttributeString(keyMissing), "Search with no attribute")
	assert.Nil(t, userRep.GetAttributeDate(keyDate, supportedDateLayouts), "Search with no attribute")

	// Sets some attributes
	userRep.SetAttributeString(keyDate, "2021/12/31")
	userRep.SetAttributeString(keyGender, "M")

	// Search for a missing attribute
	assert.Nil(t, userRep.GetAttribute(keyMissing), "Missing attribute")

	t.Run("Check that a Set/Get cycle gives the correct values", func(t *testing.T) {
		userRep.SetAttribute(AttributeKey("multiple"), []string{"3", "7", "21"})
		assert.Equal(t, []string{"3", "7", "21"}, userRep.GetAttribute(keyMultiple), "Gets a multiple-value attribute")
		assert.Equal(t, []string{"2021/12/31"}, userRep.GetAttribute(keyDate), "Gets an array")
		assert.Equal(t, "2021/12/31", *userRep.GetAttributeString(keyDate), "Gets a single attribute")
		assert.Equal(t, "31.12.2021", *userRep.GetAttributeDate(keyDate, supportedDateLayouts), "Gets a birthdate")
	})

	t.Run("Attributes and dates", func(t *testing.T) {
		// Set a date in a different format than the one which will be used to store the information
		userRep.SetAttributeDate(keyDate, "2021/12/31", supportedDateLayouts)
		assert.Equal(t, "31.12.2021", *userRep.GetAttributeString(keyDate), "Gets a single attribute")
		assert.Equal(t, "31.12.2021", *userRep.GetAttributeDate(keyDate, supportedDateLayouts), "Gets a birthdate")

		// Do not override a parameter with a nil value
		userRep.Attributes.SetDateWhenNotNil(keyDate, nil, supportedDateLayouts)
		assert.Equal(t, "31.12.2021", *userRep.GetAttributeString(keyDate))
		var otherDate = "30.11.2022"
		userRep.Attributes.SetDateWhenNotNil(keyDate, &otherDate, supportedDateLayouts)
		assert.Equal(t, otherDate, *userRep.GetAttributeString(keyDate))
	})

	t.Run("Int tests", func(t *testing.T) {
		var keyInt = AttributeKey("numeric")
		var res, err = userRep.Attributes.GetInt(keyInt)
		var value = 5
		assert.Nil(t, err)
		assert.Nil(t, res)
		// Set int value
		userRep.Attributes.SetIntWhenNotNil(keyInt, &value)
		res, err = userRep.Attributes.GetInt(keyInt)
		assert.Nil(t, err)
		assert.Equal(t, value, *res)
		// Set when not nil : won't have any effect with nil
		userRep.Attributes.SetIntWhenNotNil(keyInt, nil)
		res, err = userRep.Attributes.GetInt(keyInt)
		assert.Nil(t, err)
		assert.Equal(t, value, *res)
		// Update to 10
		value = 10
		userRep.Attributes.SetIntWhenNotNil(keyInt, &value)
		res, err = userRep.Attributes.GetInt(keyInt)
		assert.Nil(t, err)
		assert.Equal(t, value, *res)
	})

	t.Run("Boolean tests", func(t *testing.T) {
		var keyBool = AttributeKey("boolean")
		var res, err = userRep.Attributes.GetBool(keyBool)
		var value = false
		assert.Nil(t, err)
		assert.Nil(t, res)
		// Set boolean value
		userRep.Attributes.SetBoolWhenNotNil(keyBool, &value)
		res, err = userRep.Attributes.GetBool(keyBool)
		assert.Nil(t, err)
		assert.False(t, *res)
		// Set when not nil : won't have any effect with nil
		userRep.Attributes.SetBoolWhenNotNil(keyBool, nil)
		res, err = userRep.Attributes.GetBool(keyBool)
		assert.Nil(t, err)
		assert.False(t, *res)
		// Update to true
		value = true
		userRep.Attributes.SetBoolWhenNotNil(keyBool, &value)
		res, err = userRep.Attributes.GetBool(keyBool)
		assert.Nil(t, err)
		assert.True(t, *res)
	})
}

func TestMergeAttributes(t *testing.T) {
	var (
		currentAttributes = make(Attributes)
		newAttributes     = make(Attributes)
		keyOne            = AttributeKey("one")
		keyTwo            = AttributeKey("two")
		keyThree          = AttributeKey("three")
	)

	currentAttributes.SetString(keyOne, "abc")
	currentAttributes.SetString(keyThree, "zyx")

	currentAttributes.Merge(nil)
	assert.Len(t, currentAttributes, 2)

	newAttributes.SetString(keyTwo, "def")
	newAttributes.SetString(keyThree, "ghi")

	currentAttributes.Merge(&newAttributes)
	assert.Len(t, currentAttributes, 3)
	assert.Equal(t, "abc", *currentAttributes.GetString(keyOne))
	assert.Equal(t, "def", *currentAttributes.GetString(keyTwo))
	assert.Equal(t, "ghi", *currentAttributes.GetString(keyThree))
}
