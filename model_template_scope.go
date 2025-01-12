/*
Elastic Email REST API

This API is based on the REST API architecture, allowing the user to easily manage their data with this resource-based approach.    Every API call is established on which specific request type (GET, POST, PUT, DELETE) will be used.    The API has a limit of 20 concurrent connections and a hard timeout of 600 seconds per request.    To start using this API, you will need your Access Token (available <a target=\"_blank\" href=\"https://elasticemail.com/account#/settings/new/manage-api\">here</a>). Remember to keep it safe. Required access levels are listed in the given request’s description.    This is the documentation for REST API. If you’d like to read our legacy documentation regarding Web API v2 click <a target=\"_blank\" href=\"https://api.elasticemail.com/public/help\">here</a>.    Downloadable library clients can be found in our Github repository <a target=\"_blank\" href=\"https://github.com/ElasticEmail?tab=repositories&q=%22rest+api%22+in%3Areadme\">here</a>

API version: 4.0.0
Contact: support@elasticemail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ElasticEmail

import (
	"encoding/json"
	"fmt"
)

// TemplateScope Visibility of a template
type TemplateScope string

// List of TemplateScope
const (
	PERSONAL_TEMPLATE_SCOPE TemplateScope = "Personal"
	GLOBAL_TEMPLATE_SCOPE   TemplateScope = "Global"
)

// All allowed values of TemplateScope enum
var AllowedTemplateScopeEnumValues = []TemplateScope{
	"Personal",
	"Global",
}

func (v *TemplateScope) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TemplateScope(value)
	for _, existing := range AllowedTemplateScopeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TemplateScope", value)
}

// NewTemplateScopeFromValue returns a pointer to a valid TemplateScope
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTemplateScopeFromValue(v string) (*TemplateScope, error) {
	ev := TemplateScope(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TemplateScope: valid values are %v", v, AllowedTemplateScopeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TemplateScope) IsValid() bool {
	for _, existing := range AllowedTemplateScopeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TemplateScope value
func (v TemplateScope) Ptr() *TemplateScope {
	return &v
}

type NullableTemplateScope struct {
	value *TemplateScope
	isSet bool
}

func (v NullableTemplateScope) Get() *TemplateScope {
	return v.value
}

func (v *NullableTemplateScope) Set(val *TemplateScope) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateScope) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateScope(val *TemplateScope) *NullableTemplateScope {
	return &NullableTemplateScope{value: val, isSet: true}
}

func (v NullableTemplateScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
