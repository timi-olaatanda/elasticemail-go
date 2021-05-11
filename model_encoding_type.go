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

// EncodingType Encoding type for the email headers
type EncodingType string

// List of EncodingType
const (
	USER_PROVIDED EncodingType = "UserProvided"
	NONE EncodingType = "None"
	RAW7BIT EncodingType = "Raw7bit"
	RAW8BIT EncodingType = "Raw8bit"
	QUOTED_PRINTABLE EncodingType = "QuotedPrintable"
	BASE64 EncodingType = "Base64"
	UUE EncodingType = "Uue"
)

func (v *EncodingType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EncodingType(value)
	for _, existing := range []EncodingType{ "UserProvided", "None", "Raw7bit", "Raw8bit", "QuotedPrintable", "Base64", "Uue",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EncodingType", value)
}

// Ptr returns reference to EncodingType value
func (v EncodingType) Ptr() *EncodingType {
	return &v
}

type NullableEncodingType struct {
	value *EncodingType
	isSet bool
}

func (v NullableEncodingType) Get() *EncodingType {
	return v.value
}

func (v *NullableEncodingType) Set(val *EncodingType) {
	v.value = val
	v.isSet = true
}

func (v NullableEncodingType) IsSet() bool {
	return v.isSet
}

func (v *NullableEncodingType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEncodingType(val *EncodingType) *NullableEncodingType {
	return &NullableEncodingType{value: val, isSet: true}
}

func (v NullableEncodingType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEncodingType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

