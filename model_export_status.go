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

// ExportStatus Current status of the export.
type ExportStatus string

// List of ExportStatus
const (
	ERROR ExportStatus = "Error"
	LOADING ExportStatus = "Loading"
	READY ExportStatus = "Ready"
	EXPIRED ExportStatus = "Expired"
)

func (v *ExportStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ExportStatus(value)
	for _, existing := range []ExportStatus{ "Error", "Loading", "Ready", "Expired",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ExportStatus", value)
}

// Ptr returns reference to ExportStatus value
func (v ExportStatus) Ptr() *ExportStatus {
	return &v
}

type NullableExportStatus struct {
	value *ExportStatus
	isSet bool
}

func (v NullableExportStatus) Get() *ExportStatus {
	return v.value
}

func (v *NullableExportStatus) Set(val *ExportStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableExportStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableExportStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportStatus(val *ExportStatus) *NullableExportStatus {
	return &NullableExportStatus{value: val, isSet: true}
}

func (v NullableExportStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

