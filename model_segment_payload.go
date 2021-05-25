/*
 * Elastic Email REST API
 *
 * This API is based on the REST API architecture, allowing the user to easily manage their data with this resource-based approach.    Every API call is established on which specific request type (GET, POST, PUT, DELETE) will be used.    To start using this API, you will need your Access Token (available <a target=\"_blank\" href=\"https://elasticemail.com/account#/settings/new/manage-api\">here</a>). Remember to keep it safe. Required access levels are listed in the given request’s description.    This is the documentation for REST API. If you’d like to read our legacy documentation regarding Web API v2 click <a target=\"_blank\" href=\"https://api.elasticemail.com/public/help\">here</a>.    Downloadable library clients can be found in our Github repository <a target=\"_blank\" href=\"https://github.com/ElasticEmail?tab=repositories&q=%22rest+api%22+in%3Areadme\">here</a>
 *
 * API version: 4.0.0
 * Contact: support@elasticemail.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ElasticEmail

import (
	"encoding/json"
)

// SegmentPayload struct for SegmentPayload
type SegmentPayload struct {
	// Segment name
	Name *string `json:"Name,omitempty"`
	// SQL-like rule to determine which Contacts belong to this Segment. Help for building a segment rule can be found here: https://help.elasticemail.com/en/articles/5162182-segment-rules
	Rule *string `json:"Rule,omitempty"`
}

// NewSegmentPayload instantiates a new SegmentPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSegmentPayload() *SegmentPayload {
	this := SegmentPayload{}
	return &this
}

// NewSegmentPayloadWithDefaults instantiates a new SegmentPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSegmentPayloadWithDefaults() *SegmentPayload {
	this := SegmentPayload{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SegmentPayload) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentPayload) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SegmentPayload) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SegmentPayload) SetName(v string) {
	o.Name = &v
}

// GetRule returns the Rule field value if set, zero value otherwise.
func (o *SegmentPayload) GetRule() string {
	if o == nil || o.Rule == nil {
		var ret string
		return ret
	}
	return *o.Rule
}

// GetRuleOk returns a tuple with the Rule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentPayload) GetRuleOk() (*string, bool) {
	if o == nil || o.Rule == nil {
		return nil, false
	}
	return o.Rule, true
}

// HasRule returns a boolean if a field has been set.
func (o *SegmentPayload) HasRule() bool {
	if o != nil && o.Rule != nil {
		return true
	}

	return false
}

// SetRule gets a reference to the given string and assigns it to the Rule field.
func (o *SegmentPayload) SetRule(v string) {
	o.Rule = &v
}

func (o SegmentPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Rule != nil {
		toSerialize["Rule"] = o.Rule
	}
	return json.Marshal(toSerialize)
}

type NullableSegmentPayload struct {
	value *SegmentPayload
	isSet bool
}

func (v NullableSegmentPayload) Get() *SegmentPayload {
	return v.value
}

func (v *NullableSegmentPayload) Set(val *SegmentPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableSegmentPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableSegmentPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSegmentPayload(val *SegmentPayload) *NullableSegmentPayload {
	return &NullableSegmentPayload{value: val, isSet: true}
}

func (v NullableSegmentPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSegmentPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

