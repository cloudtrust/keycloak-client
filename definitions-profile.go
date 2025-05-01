package keycloak

import "strings"

// UserProfileRepresentation struct
type UserProfileRepresentation struct {
	Attributes []ProfileAttrbRepresentation `json:"attributes"`
	Groups     []ProfileGroupRepresentation `json:"groups"`
}

// ProfileAttrbRepresentation struct
type ProfileAttrbRepresentation struct {
	Name        *string                                `json:"name,omitempty"`
	DisplayName *string                                `json:"displayName,omitempty"`
	Group       *string                                `json:"group,omitempty"`
	Required    *ProfileAttrbRequiredRepresentation    `json:"required,omitempty"`
	Permissions *ProfileAttrbPermissionsRepresentation `json:"permissions,omitempty"`
	Validations ProfileAttrbValidationRepresentation   `json:"validations,omitempty"`
	Selector    *ProfileAttrbSelectorRepresentation    `json:"selector,omitempty"`
	Annotations map[string]string                      `json:"annotations,omitempty"`
}

// ProfileAttrbValidationRepresentation struct
// Known keys:
// - email: empty
// - length: min, max (int/string), trim-disabled (boolean as string)
// - integer: min, max (integer as string)
// - double: min, max (double as string)
// - options: options (array of allowed values)
// - pattern: pattern (regex as string), error-message (string)
// - local-date: empty
// - uri: empty
// - username-prohibited-characters: error-message (string)
// - person-name-prohibited-characters: error-message (string)
type ProfileAttrbValidationRepresentation map[string]ProfileAttrValidatorRepresentation

// ProfileAttrValidatorRepresentation type
type ProfileAttrValidatorRepresentation map[string]any

// ProfileAttrbRequiredRepresentation struct
type ProfileAttrbRequiredRepresentation struct {
	Roles  []string `json:"roles,omitempty"`
	Scopes []string `json:"scopes,omitempty"`
}

// ProfileAttrbPermissionsRepresentation struct
type ProfileAttrbPermissionsRepresentation struct {
	View []string `json:"view,omitempty"`
	Edit []string `json:"edit,omitempty"`
}

// ProfileAttrbSelectorRepresentation struct
type ProfileAttrbSelectorRepresentation struct {
	Scopes []string `json:"scopes,omitempty"`
}

// ProfileGroupRepresentation struct
type ProfileGroupRepresentation struct {
	Name               *string           `json:"name,omitempty"`
	DisplayHeader      *string           `displayHeader:"name,omitempty"`
	DisplayDescription *string           `displayDescription:"name,omitempty"`
	Annotations        map[string]string `annotations:"name,omitempty"`
}

// IsAnnotationTrue checks if an annotation is true
func (attrb *ProfileAttrbRepresentation) IsAnnotationTrue(key string) bool {
	return attrb.AnnotationEqualsIgnoreCase(key, "true")
}

// IsAnnotationFalse checks if an annotation is false
func (attrb *ProfileAttrbRepresentation) IsAnnotationFalse(key string) bool {
	return attrb.AnnotationEqualsIgnoreCase(key, "false")
}

// AnnotationEqualsIgnoreCase checks if an annotation
func (attrb *ProfileAttrbRepresentation) AnnotationEqualsIgnoreCase(key string, value string) bool {
	return attrb.AnnotationMatches(key, func(attrbValue string) bool {
		return strings.EqualFold(value, attrbValue)
	})
}

// AnnotationMatches checks if an annotation
func (attrb *ProfileAttrbRepresentation) AnnotationMatches(key string, matcher func(value string) bool) bool {
	if attrb.Annotations != nil {
		if attrbValue, ok := attrb.Annotations[key]; ok {
			return matcher(attrbValue)
		}
	}
	return false
}
