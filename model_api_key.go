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
	"time"
)

// ApiKey ApiKey info
type ApiKey struct {
	// Access level or permission to be assigned to this ApiKey.
	AccessLevel *[]AccessLevel `json:"AccessLevel,omitempty"`
	// Name of the ApiKey.
	Name *string `json:"Name,omitempty"`
	// Date this ApiKey was created.
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// Date this ApiKey was last used.
	LastUse NullableTime `json:"LastUse,omitempty"`
	// Date this ApiKey expires.
	Expires NullableTime `json:"Expires,omitempty"`
	// Which IPs can use this ApiKey
	RestrictAccessToIPRange *[]string `json:"RestrictAccessToIPRange,omitempty"`
}

// NewApiKey instantiates a new ApiKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiKey() *ApiKey {
	this := ApiKey{}
	return &this
}

// NewApiKeyWithDefaults instantiates a new ApiKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiKeyWithDefaults() *ApiKey {
	this := ApiKey{}
	return &this
}

// GetAccessLevel returns the AccessLevel field value if set, zero value otherwise.
func (o *ApiKey) GetAccessLevel() []AccessLevel {
	if o == nil || o.AccessLevel == nil {
		var ret []AccessLevel
		return ret
	}
	return *o.AccessLevel
}

// GetAccessLevelOk returns a tuple with the AccessLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKey) GetAccessLevelOk() (*[]AccessLevel, bool) {
	if o == nil || o.AccessLevel == nil {
		return nil, false
	}
	return o.AccessLevel, true
}

// HasAccessLevel returns a boolean if a field has been set.
func (o *ApiKey) HasAccessLevel() bool {
	if o != nil && o.AccessLevel != nil {
		return true
	}

	return false
}

// SetAccessLevel gets a reference to the given []AccessLevel and assigns it to the AccessLevel field.
func (o *ApiKey) SetAccessLevel(v []AccessLevel) {
	o.AccessLevel = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiKey) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKey) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiKey) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiKey) SetName(v string) {
	o.Name = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *ApiKey) GetDateCreated() time.Time {
	if o == nil || o.DateCreated == nil {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKey) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || o.DateCreated == nil {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *ApiKey) HasDateCreated() bool {
	if o != nil && o.DateCreated != nil {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *ApiKey) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetLastUse returns the LastUse field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiKey) GetLastUse() time.Time {
	if o == nil || o.LastUse.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUse.Get()
}

// GetLastUseOk returns a tuple with the LastUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiKey) GetLastUseOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastUse.Get(), o.LastUse.IsSet()
}

// HasLastUse returns a boolean if a field has been set.
func (o *ApiKey) HasLastUse() bool {
	if o != nil && o.LastUse.IsSet() {
		return true
	}

	return false
}

// SetLastUse gets a reference to the given NullableTime and assigns it to the LastUse field.
func (o *ApiKey) SetLastUse(v time.Time) {
	o.LastUse.Set(&v)
}
// SetLastUseNil sets the value for LastUse to be an explicit nil
func (o *ApiKey) SetLastUseNil() {
	o.LastUse.Set(nil)
}

// UnsetLastUse ensures that no value is present for LastUse, not even an explicit nil
func (o *ApiKey) UnsetLastUse() {
	o.LastUse.Unset()
}

// GetExpires returns the Expires field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiKey) GetExpires() time.Time {
	if o == nil || o.Expires.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Expires.Get()
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiKey) GetExpiresOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Expires.Get(), o.Expires.IsSet()
}

// HasExpires returns a boolean if a field has been set.
func (o *ApiKey) HasExpires() bool {
	if o != nil && o.Expires.IsSet() {
		return true
	}

	return false
}

// SetExpires gets a reference to the given NullableTime and assigns it to the Expires field.
func (o *ApiKey) SetExpires(v time.Time) {
	o.Expires.Set(&v)
}
// SetExpiresNil sets the value for Expires to be an explicit nil
func (o *ApiKey) SetExpiresNil() {
	o.Expires.Set(nil)
}

// UnsetExpires ensures that no value is present for Expires, not even an explicit nil
func (o *ApiKey) UnsetExpires() {
	o.Expires.Unset()
}

// GetRestrictAccessToIPRange returns the RestrictAccessToIPRange field value if set, zero value otherwise.
func (o *ApiKey) GetRestrictAccessToIPRange() []string {
	if o == nil || o.RestrictAccessToIPRange == nil {
		var ret []string
		return ret
	}
	return *o.RestrictAccessToIPRange
}

// GetRestrictAccessToIPRangeOk returns a tuple with the RestrictAccessToIPRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKey) GetRestrictAccessToIPRangeOk() (*[]string, bool) {
	if o == nil || o.RestrictAccessToIPRange == nil {
		return nil, false
	}
	return o.RestrictAccessToIPRange, true
}

// HasRestrictAccessToIPRange returns a boolean if a field has been set.
func (o *ApiKey) HasRestrictAccessToIPRange() bool {
	if o != nil && o.RestrictAccessToIPRange != nil {
		return true
	}

	return false
}

// SetRestrictAccessToIPRange gets a reference to the given []string and assigns it to the RestrictAccessToIPRange field.
func (o *ApiKey) SetRestrictAccessToIPRange(v []string) {
	o.RestrictAccessToIPRange = &v
}

func (o ApiKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessLevel != nil {
		toSerialize["AccessLevel"] = o.AccessLevel
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.DateCreated != nil {
		toSerialize["DateCreated"] = o.DateCreated
	}
	if o.LastUse.IsSet() {
		toSerialize["LastUse"] = o.LastUse.Get()
	}
	if o.Expires.IsSet() {
		toSerialize["Expires"] = o.Expires.Get()
	}
	if o.RestrictAccessToIPRange != nil {
		toSerialize["RestrictAccessToIPRange"] = o.RestrictAccessToIPRange
	}
	return json.Marshal(toSerialize)
}

type NullableApiKey struct {
	value *ApiKey
	isSet bool
}

func (v NullableApiKey) Get() *ApiKey {
	return v.value
}

func (v *NullableApiKey) Set(val *ApiKey) {
	v.value = val
	v.isSet = true
}

func (v NullableApiKey) IsSet() bool {
	return v.isSet
}

func (v *NullableApiKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiKey(val *ApiKey) *NullableApiKey {
	return &NullableApiKey{value: val, isSet: true}
}

func (v NullableApiKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

