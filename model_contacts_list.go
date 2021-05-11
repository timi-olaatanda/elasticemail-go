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
	"time"
)

// ContactsList List of Lists, with detailed data about its contents.
type ContactsList struct {
	// Name of your list.
	ListName *string `json:"ListName,omitempty"`
	// ID code of list. Please note that this is different from the listid field.
	PublicListID NullableString `json:"PublicListID,omitempty"`
	// Date of creation in YYYY-MM-DDThh:ii:ss format
	DateAdded *time.Time `json:"DateAdded,omitempty"`
	// True: Allow unsubscribing from this list. Otherwise, false
	AllowUnsubscribe *bool `json:"AllowUnsubscribe,omitempty"`
}

// NewContactsList instantiates a new ContactsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactsList() *ContactsList {
	this := ContactsList{}
	return &this
}

// NewContactsListWithDefaults instantiates a new ContactsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactsListWithDefaults() *ContactsList {
	this := ContactsList{}
	return &this
}

// GetListName returns the ListName field value if set, zero value otherwise.
func (o *ContactsList) GetListName() string {
	if o == nil || o.ListName == nil {
		var ret string
		return ret
	}
	return *o.ListName
}

// GetListNameOk returns a tuple with the ListName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactsList) GetListNameOk() (*string, bool) {
	if o == nil || o.ListName == nil {
		return nil, false
	}
	return o.ListName, true
}

// HasListName returns a boolean if a field has been set.
func (o *ContactsList) HasListName() bool {
	if o != nil && o.ListName != nil {
		return true
	}

	return false
}

// SetListName gets a reference to the given string and assigns it to the ListName field.
func (o *ContactsList) SetListName(v string) {
	o.ListName = &v
}

// GetPublicListID returns the PublicListID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContactsList) GetPublicListID() string {
	if o == nil || o.PublicListID.Get() == nil {
		var ret string
		return ret
	}
	return *o.PublicListID.Get()
}

// GetPublicListIDOk returns a tuple with the PublicListID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContactsList) GetPublicListIDOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PublicListID.Get(), o.PublicListID.IsSet()
}

// HasPublicListID returns a boolean if a field has been set.
func (o *ContactsList) HasPublicListID() bool {
	if o != nil && o.PublicListID.IsSet() {
		return true
	}

	return false
}

// SetPublicListID gets a reference to the given NullableString and assigns it to the PublicListID field.
func (o *ContactsList) SetPublicListID(v string) {
	o.PublicListID.Set(&v)
}
// SetPublicListIDNil sets the value for PublicListID to be an explicit nil
func (o *ContactsList) SetPublicListIDNil() {
	o.PublicListID.Set(nil)
}

// UnsetPublicListID ensures that no value is present for PublicListID, not even an explicit nil
func (o *ContactsList) UnsetPublicListID() {
	o.PublicListID.Unset()
}

// GetDateAdded returns the DateAdded field value if set, zero value otherwise.
func (o *ContactsList) GetDateAdded() time.Time {
	if o == nil || o.DateAdded == nil {
		var ret time.Time
		return ret
	}
	return *o.DateAdded
}

// GetDateAddedOk returns a tuple with the DateAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactsList) GetDateAddedOk() (*time.Time, bool) {
	if o == nil || o.DateAdded == nil {
		return nil, false
	}
	return o.DateAdded, true
}

// HasDateAdded returns a boolean if a field has been set.
func (o *ContactsList) HasDateAdded() bool {
	if o != nil && o.DateAdded != nil {
		return true
	}

	return false
}

// SetDateAdded gets a reference to the given time.Time and assigns it to the DateAdded field.
func (o *ContactsList) SetDateAdded(v time.Time) {
	o.DateAdded = &v
}

// GetAllowUnsubscribe returns the AllowUnsubscribe field value if set, zero value otherwise.
func (o *ContactsList) GetAllowUnsubscribe() bool {
	if o == nil || o.AllowUnsubscribe == nil {
		var ret bool
		return ret
	}
	return *o.AllowUnsubscribe
}

// GetAllowUnsubscribeOk returns a tuple with the AllowUnsubscribe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactsList) GetAllowUnsubscribeOk() (*bool, bool) {
	if o == nil || o.AllowUnsubscribe == nil {
		return nil, false
	}
	return o.AllowUnsubscribe, true
}

// HasAllowUnsubscribe returns a boolean if a field has been set.
func (o *ContactsList) HasAllowUnsubscribe() bool {
	if o != nil && o.AllowUnsubscribe != nil {
		return true
	}

	return false
}

// SetAllowUnsubscribe gets a reference to the given bool and assigns it to the AllowUnsubscribe field.
func (o *ContactsList) SetAllowUnsubscribe(v bool) {
	o.AllowUnsubscribe = &v
}

func (o ContactsList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ListName != nil {
		toSerialize["ListName"] = o.ListName
	}
	if o.PublicListID.IsSet() {
		toSerialize["PublicListID"] = o.PublicListID.Get()
	}
	if o.DateAdded != nil {
		toSerialize["DateAdded"] = o.DateAdded
	}
	if o.AllowUnsubscribe != nil {
		toSerialize["AllowUnsubscribe"] = o.AllowUnsubscribe
	}
	return json.Marshal(toSerialize)
}

type NullableContactsList struct {
	value *ContactsList
	isSet bool
}

func (v NullableContactsList) Get() *ContactsList {
	return v.value
}

func (v *NullableContactsList) Set(val *ContactsList) {
	v.value = val
	v.isSet = true
}

func (v NullableContactsList) IsSet() bool {
	return v.isSet
}

func (v *NullableContactsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactsList(val *ContactsList) *NullableContactsList {
	return &NullableContactsList{value: val, isSet: true}
}

func (v NullableContactsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


