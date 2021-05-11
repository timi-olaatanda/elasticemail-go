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
)

// InboundRoute struct for InboundRoute
type InboundRoute struct {
	PublicId *string `json:"PublicId,omitempty"`
	// Name of this route
	Name *string `json:"Name,omitempty"`
	// Type of the filter
	FilterType *InboundRouteFilterType `json:"FilterType,omitempty"`
	// Filter of the inbound data
	Filter *string `json:"Filter,omitempty"`
	// Type of action to take
	ActionType *InboundRouteActionType `json:"ActionType,omitempty"`
	// URL address or Email to notify about the inbound
	ActionParameter *string `json:"ActionParameter,omitempty"`
	// Place of this route in your routes queue's order
	SortOrder *int32 `json:"SortOrder,omitempty"`
}

// NewInboundRoute instantiates a new InboundRoute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInboundRoute() *InboundRoute {
	this := InboundRoute{}
	return &this
}

// NewInboundRouteWithDefaults instantiates a new InboundRoute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInboundRouteWithDefaults() *InboundRoute {
	this := InboundRoute{}
	return &this
}

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *InboundRoute) GetPublicId() string {
	if o == nil || o.PublicId == nil {
		var ret string
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundRoute) GetPublicIdOk() (*string, bool) {
	if o == nil || o.PublicId == nil {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *InboundRoute) HasPublicId() bool {
	if o != nil && o.PublicId != nil {
		return true
	}

	return false
}

// SetPublicId gets a reference to the given string and assigns it to the PublicId field.
func (o *InboundRoute) SetPublicId(v string) {
	o.PublicId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InboundRoute) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundRoute) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InboundRoute) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InboundRoute) SetName(v string) {
	o.Name = &v
}

// GetFilterType returns the FilterType field value if set, zero value otherwise.
func (o *InboundRoute) GetFilterType() InboundRouteFilterType {
	if o == nil || o.FilterType == nil {
		var ret InboundRouteFilterType
		return ret
	}
	return *o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundRoute) GetFilterTypeOk() (*InboundRouteFilterType, bool) {
	if o == nil || o.FilterType == nil {
		return nil, false
	}
	return o.FilterType, true
}

// HasFilterType returns a boolean if a field has been set.
func (o *InboundRoute) HasFilterType() bool {
	if o != nil && o.FilterType != nil {
		return true
	}

	return false
}

// SetFilterType gets a reference to the given InboundRouteFilterType and assigns it to the FilterType field.
func (o *InboundRoute) SetFilterType(v InboundRouteFilterType) {
	o.FilterType = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *InboundRoute) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundRoute) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *InboundRoute) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *InboundRoute) SetFilter(v string) {
	o.Filter = &v
}

// GetActionType returns the ActionType field value if set, zero value otherwise.
func (o *InboundRoute) GetActionType() InboundRouteActionType {
	if o == nil || o.ActionType == nil {
		var ret InboundRouteActionType
		return ret
	}
	return *o.ActionType
}

// GetActionTypeOk returns a tuple with the ActionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundRoute) GetActionTypeOk() (*InboundRouteActionType, bool) {
	if o == nil || o.ActionType == nil {
		return nil, false
	}
	return o.ActionType, true
}

// HasActionType returns a boolean if a field has been set.
func (o *InboundRoute) HasActionType() bool {
	if o != nil && o.ActionType != nil {
		return true
	}

	return false
}

// SetActionType gets a reference to the given InboundRouteActionType and assigns it to the ActionType field.
func (o *InboundRoute) SetActionType(v InboundRouteActionType) {
	o.ActionType = &v
}

// GetActionParameter returns the ActionParameter field value if set, zero value otherwise.
func (o *InboundRoute) GetActionParameter() string {
	if o == nil || o.ActionParameter == nil {
		var ret string
		return ret
	}
	return *o.ActionParameter
}

// GetActionParameterOk returns a tuple with the ActionParameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundRoute) GetActionParameterOk() (*string, bool) {
	if o == nil || o.ActionParameter == nil {
		return nil, false
	}
	return o.ActionParameter, true
}

// HasActionParameter returns a boolean if a field has been set.
func (o *InboundRoute) HasActionParameter() bool {
	if o != nil && o.ActionParameter != nil {
		return true
	}

	return false
}

// SetActionParameter gets a reference to the given string and assigns it to the ActionParameter field.
func (o *InboundRoute) SetActionParameter(v string) {
	o.ActionParameter = &v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise.
func (o *InboundRoute) GetSortOrder() int32 {
	if o == nil || o.SortOrder == nil {
		var ret int32
		return ret
	}
	return *o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundRoute) GetSortOrderOk() (*int32, bool) {
	if o == nil || o.SortOrder == nil {
		return nil, false
	}
	return o.SortOrder, true
}

// HasSortOrder returns a boolean if a field has been set.
func (o *InboundRoute) HasSortOrder() bool {
	if o != nil && o.SortOrder != nil {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given int32 and assigns it to the SortOrder field.
func (o *InboundRoute) SetSortOrder(v int32) {
	o.SortOrder = &v
}

func (o InboundRoute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PublicId != nil {
		toSerialize["PublicId"] = o.PublicId
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.FilterType != nil {
		toSerialize["FilterType"] = o.FilterType
	}
	if o.Filter != nil {
		toSerialize["Filter"] = o.Filter
	}
	if o.ActionType != nil {
		toSerialize["ActionType"] = o.ActionType
	}
	if o.ActionParameter != nil {
		toSerialize["ActionParameter"] = o.ActionParameter
	}
	if o.SortOrder != nil {
		toSerialize["SortOrder"] = o.SortOrder
	}
	return json.Marshal(toSerialize)
}

type NullableInboundRoute struct {
	value *InboundRoute
	isSet bool
}

func (v NullableInboundRoute) Get() *InboundRoute {
	return v.value
}

func (v *NullableInboundRoute) Set(val *InboundRoute) {
	v.value = val
	v.isSet = true
}

func (v NullableInboundRoute) IsSet() bool {
	return v.isSet
}

func (v *NullableInboundRoute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInboundRoute(val *InboundRoute) *NullableInboundRoute {
	return &NullableInboundRoute{value: val, isSet: true}
}

func (v NullableInboundRoute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInboundRoute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


