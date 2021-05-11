/*
 * Elastic Email REST API
 *
 * This API is based on the REST API architecture, allowing the user to easily manage their data with this resource-based approach.    Every API call is established on which specific request type (GET, POST, PUT, DELETE) will be used.    To start using this API, you will need your Access Token (available <a href=\"https://elasticemail.com/account#/settings/new/manage-api\">here</a>). Remember to keep it safe. Required access levels are listed in the given request’s description.    This is the documentation for REST API. If you’d like to read our legacy documentation regarding Web API v2 click <a href=\"https://api.elasticemail.com/public/help\">here</a>.
 *
 * API version: 4.0.0
 * Contact: support@elasticemail.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ElasticEmail

import (
	"encoding/json"
	"fmt"
)

// BodyContentType Type of body part
type BodyContentType string

// List of BodyContentType
const (
	HTML BodyContentType = "HTML"
	PLAIN_TEXT BodyContentType = "PlainText"
	AMP BodyContentType = "AMP"
	CSS BodyContentType = "CSS"
)

func (v *BodyContentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BodyContentType(value)
	for _, existing := range []BodyContentType{ "HTML", "PlainText", "AMP", "CSS",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BodyContentType", value)
}

// Ptr returns reference to BodyContentType value
func (v BodyContentType) Ptr() *BodyContentType {
	return &v
}

type NullableBodyContentType struct {
	value *BodyContentType
	isSet bool
}

func (v NullableBodyContentType) Get() *BodyContentType {
	return v.value
}

func (v *NullableBodyContentType) Set(val *BodyContentType) {
	v.value = val
	v.isSet = true
}

func (v NullableBodyContentType) IsSet() bool {
	return v.isSet
}

func (v *NullableBodyContentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBodyContentType(val *BodyContentType) *NullableBodyContentType {
	return &NullableBodyContentType{value: val, isSet: true}
}

func (v NullableBodyContentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBodyContentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

