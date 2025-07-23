package keycloak

import (
	"strconv"
	"time"

	"github.com/cloudtrust/common-service/v2/fields"
)

// Get a given attribute
func (a Attributes) Get(key AttributeKey) []string {
	return a[key]
}

// Set a given attribute
func (a Attributes) Set(key AttributeKey, value []string) {
	a[key] = value
}

// Remove a given attribute
func (a Attributes) Remove(key AttributeKey) {
	delete(a, key)
}

// GetString gets the first value of a given attribute
func (a Attributes) GetString(key AttributeKey) *string {
	var attrbs = a[key]
	if len(attrbs) == 1 && attrbs[0] == "" {
		// User profile sometimes sets a default empty value
		return nil
	}
	if len(attrbs) > 0 {
		return &attrbs[0]
	}
	return nil
}

// SetString sets the value of a given attribute
func (a Attributes) SetString(key AttributeKey, value string) {
	a.Set(key, []string{value})
}

// GetInt gets the first value of a given attribute
func (a Attributes) GetInt(key AttributeKey) (*int, error) {
	var attrb = a.GetString(key)
	if attrb != nil {
		var res64, err = strconv.ParseInt(*attrb, 0, 0)
		var res = int(res64)
		return &res, err
	}
	return nil, nil
}

// SetInt sets the value of a given attribute
func (a Attributes) SetInt(key AttributeKey, value int) {
	a.Set(key, []string{strconv.FormatInt(int64(value), 10)})
}

// GetBool gets the first value of a given attribute
func (a Attributes) GetBool(key AttributeKey) (*bool, error) {
	var attrb = a.GetString(key)
	if attrb != nil {
		var res, err = strconv.ParseBool(*attrb)
		return &res, err
	}
	return nil, nil
}

// SetBool sets the value of a given attribute
func (a Attributes) SetBool(key AttributeKey, value bool) {
	a.Set(key, []string{strconv.FormatBool(value)})
}

// GetDate returns an attribute which contains a date value
func (a Attributes) GetDate(key AttributeKey, dateLayouts []string) *string {
	var attrb = a.GetString(key)
	var formatted = a.reformatDate(attrb, dateLayouts)
	if formatted != nil {
		a[key] = []string{*formatted}
		return formatted
	}
	return attrb
}

// SetDate sets a date
func (a Attributes) SetDate(key AttributeKey, value string, dateLayouts []string) {
	var formatted = a.reformatDate(&value, dateLayouts)
	if formatted != nil {
		value = *formatted
	}
	a.Set(key, []string{value})
}

// GetTime returns an attribute which contains a date value
func (a Attributes) GetTime(key AttributeKey, dateLayouts []string) (*time.Time, error) {
	return a.parseDate(a.GetString(key), dateLayouts)
}

// SetTime sets a date
func (a Attributes) SetTime(key AttributeKey, value time.Time, dateLayout string) {
	a.Set(key, []string{value.Format(dateLayout)})
}

// SetStringWhenNotNil sets an attribute value if it is not nil
func (a Attributes) SetStringWhenNotNil(key AttributeKey, value *string) {
	if value != nil {
		a.Set(key, []string{*value})
	}
}

// SetIntWhenNotNil sets an attribute value if it is not nil
func (a Attributes) SetIntWhenNotNil(key AttributeKey, value *int) {
	if value != nil {
		a.Set(key, []string{strconv.FormatInt(int64(*value), 10)})
	}
}

// SetBoolWhenNotNil sets an attribute value if it is not nil
func (a Attributes) SetBoolWhenNotNil(key AttributeKey, value *bool) {
	if value != nil {
		a.Set(key, []string{strconv.FormatBool(*value)})
	}
}

// SetDateWhenNotNil sets a date attribute if it is not nil
func (a Attributes) SetDateWhenNotNil(key AttributeKey, value *string, dateLayouts []string) {
	if value != nil {
		a.SetDate(key, *value, dateLayouts)
	}
}

// SetTimeWhenNotNil sets a date attribute if it is not nil
func (a Attributes) SetTimeWhenNotNil(key AttributeKey, value *time.Time, dateLayout string) {
	if value != nil {
		a.SetTime(key, *value, dateLayout)
	}
}

func (a Attributes) parseDate(value *string, dateLayouts []string) (*time.Time, error) {
	if value == nil || len(dateLayouts) == 0 {
		return nil, nil
	}
	var date, firstErr = time.Parse(dateLayouts[0], *value)
	if firstErr == nil {
		return &date, nil
	}

	// Date does not have the expected layout. Try to convert it from supported layouts
	var err error
	for _, layout := range dateLayouts[1:] {
		date, err = time.Parse(layout, *value)
		if err == nil {
			return &date, nil
		}
	}

	return nil, firstErr
}

func (a Attributes) reformatDate(value *string, dateLayouts []string) *string {
	var date, err = a.parseDate(value, dateLayouts)
	if err != nil || date == nil {
		return nil
	}
	var res = date.Format(dateLayouts[0])
	return &res
}

// Merge current attributes with others (Values from others replace those with the same key in current attributes)
func (a Attributes) Merge(others *Attributes) {
	if others != nil {
		for key, attribute := range *others {
			a[key] = attribute
		}
	}
}

func (a Attributes) SetDynamicAttributes(dynamicAttributes map[string]any, profile UserProfileRepresentation) {
	for k, v := range dynamicAttributes {
		if _, ok := profile.dynamicAttributeKeys[k]; ok {
			strValue, ok := v.(string)
			if ok {
				// Hypothesis : all dynamic attributes are string
				a.SetString(AttributeKey(k), strValue)
			}
		}
	}
}

// RemoveAttribute removes an attribute given its key
func (u *UserRepresentation) RemoveAttribute(key AttributeKey) {
	if u.Attributes != nil {
		u.Attributes.Remove(key)
	}
}

// GetFieldValues returns a field value
func (u *UserRepresentation) GetFieldValues(field fields.Field) []string {
	switch field {
	case fields.Email:
		return toSlice(u.Email)
	case fields.FirstName:
		return toSlice(u.FirstName)
	case fields.LastName:
		return toSlice(u.LastName)
	}
	return u.GetAttribute(AttributeKey(field.AttributeName()))
}

func toSlice(value *string) []string {
	if value == nil {
		return nil
	}
	return []string{*value}
}

// SetFieldValues sets a field value
func (u *UserRepresentation) SetFieldValues(field fields.Field, values []string) {
	switch field {
	case fields.Email:
		u.Email = toSingleString(values)
	case fields.FirstName:
		u.FirstName = toSingleString(values)
	case fields.LastName:
		u.LastName = toSingleString(values)
	default:
		if len(values) == 0 {
			u.RemoveAttribute(AttributeKey(field.AttributeName()))
		} else {
			u.SetAttribute(AttributeKey(field.AttributeName()), values)
		}
	}
}

func toSingleString(values []string) *string {
	if len(values) == 1 {
		return &values[0]
	}
	return nil
}

// GetAttribute returns an attribute given its key
func (u *UserRepresentation) GetAttribute(key AttributeKey) []string {
	if u.Attributes != nil {
		return u.Attributes.Get(key)
	}
	return nil
}

// SetAttribute sets an attribute
func (u *UserRepresentation) SetAttribute(key AttributeKey, value []string) {
	if u.Attributes == nil {
		var attrbs = make(Attributes)
		u.Attributes = &attrbs
	}
	u.Attributes.Set(key, value)
}

// GetAttributeString returns the first value of an attribute given its key
func (u *UserRepresentation) GetAttributeString(key AttributeKey) *string {
	if u.Attributes != nil {
		return u.Attributes.GetString(key)
	}
	return nil
}

// SetAttributeString sets an attribute with a single value
func (u *UserRepresentation) SetAttributeString(key AttributeKey, value string) {
	u.SetAttribute(key, []string{value})
}

// GetAttributeBool returns the first value of an attribute given its key
func (u *UserRepresentation) GetAttributeBool(key AttributeKey) (*bool, error) {
	if u.Attributes != nil {
		return u.Attributes.GetBool(key)
	}
	return nil, nil
}

func (u *UserRepresentation) GetDynamicAttributes(profile UserProfileRepresentation) map[string]any {
	dynamicAttributes := map[string]any{}
	if u.Attributes != nil {
		attributes := map[AttributeKey][]string(*u.Attributes)
		for attributeName := range profile.dynamicAttributeKeys {
			attributeKey := AttributeKey(attributeName)
			if _, ok := attributes[attributeKey]; ok && len(attributes[attributeKey]) > 0 {
				// Hypothesis : all dynamic attributes can be returned as string
				dynamicAttributes[string(attributeKey)] = attributes[attributeKey][0]
			}
		}
	}
	return dynamicAttributes
}

// SetAttributeBool sets an attribute with a single value
func (u *UserRepresentation) SetAttributeBool(key AttributeKey, value bool) {
	u.SetAttribute(key, []string{strconv.FormatBool(value)})
}

// GetAttributeInt returns the first value of an attribute given its key
func (u *UserRepresentation) GetAttributeInt(key AttributeKey) (*int, error) {
	if u.Attributes != nil {
		return u.Attributes.GetInt(key)
	}
	return nil, nil
}

// SetAttributeInt sets an attribute with a single value
func (u *UserRepresentation) SetAttributeInt(key AttributeKey, value int) {
	u.SetAttribute(key, []string{strconv.FormatInt(int64(value), 10)})
}

// GetAttributeDate returns an attribute which contains a date value
func (u *UserRepresentation) GetAttributeDate(key AttributeKey, dateLayouts []string) *string {
	if u.Attributes != nil {
		return u.Attributes.GetDate(key, dateLayouts)
	}
	return nil
}

// SetAttributeDate sets a date attribute
func (u *UserRepresentation) SetAttributeDate(key AttributeKey, date string, dateLayouts []string) {
	if u.Attributes == nil {
		var attrbs = make(Attributes)
		u.Attributes = &attrbs
	}
	u.Attributes.SetDate(key, date, dateLayouts)
}

// GetAttributeTime returns an attribute which contains a date value
func (u *UserRepresentation) GetAttributeTime(key AttributeKey, dateLayouts []string) (*time.Time, error) {
	if u.Attributes != nil {
		return u.Attributes.GetTime(key, dateLayouts)
	}
	return nil, nil
}

// SetAttributeTime sets a date attribute
func (u *UserRepresentation) SetAttributeTime(key AttributeKey, date time.Time, dateLayout string) {
	u.SetAttributeString(key, date.Format(dateLayout))
}
