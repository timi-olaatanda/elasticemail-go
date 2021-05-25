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

// ContactActivity struct for ContactActivity
type ContactActivity struct {
	// Total emails sent.
	TotalSent *int32 `json:"TotalSent,omitempty"`
	// Total emails opened.
	TotalOpened *int32 `json:"TotalOpened,omitempty"`
	// Total emails clicked
	TotalClicked *int32 `json:"TotalClicked,omitempty"`
	// Total emails failed.
	TotalFailed *int32 `json:"TotalFailed,omitempty"`
	// Last date when an email was sent to this contact
	LastSent NullableTime `json:"LastSent,omitempty"`
	// Date this contact last opened an email
	LastOpened NullableTime `json:"LastOpened,omitempty"`
	// Date this contact last clicked an email
	LastClicked NullableTime `json:"LastClicked,omitempty"`
	// Last date when an email sent to this contact bounced
	LastFailed NullableTime `json:"LastFailed,omitempty"`
	// IP from which this contact opened or clicked their email last time
	LastIP *string `json:"LastIP,omitempty"`
	// Last RFC Error code if any occurred
	ErrorCode NullableInt32 `json:"ErrorCode,omitempty"`
	// Last RFC error message if any occurred
	FriendlyErrorMessage *string `json:"FriendlyErrorMessage,omitempty"`
}

// NewContactActivity instantiates a new ContactActivity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactActivity() *ContactActivity {
	this := ContactActivity{}
	return &this
}

// NewContactActivityWithDefaults instantiates a new ContactActivity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactActivityWithDefaults() *ContactActivity {
	this := ContactActivity{}
	return &this
}

// GetTotalSent returns the TotalSent field value if set, zero value otherwise.
func (o *ContactActivity) GetTotalSent() int32 {
	if o == nil || o.TotalSent == nil {
		var ret int32
		return ret
	}
	return *o.TotalSent
}

// GetTotalSentOk returns a tuple with the TotalSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactActivity) GetTotalSentOk() (*int32, bool) {
	if o == nil || o.TotalSent == nil {
		return nil, false
	}
	return o.TotalSent, true
}

// HasTotalSent returns a boolean if a field has been set.
func (o *ContactActivity) HasTotalSent() bool {
	if o != nil && o.TotalSent != nil {
		return true
	}

	return false
}

// SetTotalSent gets a reference to the given int32 and assigns it to the TotalSent field.
func (o *ContactActivity) SetTotalSent(v int32) {
	o.TotalSent = &v
}

// GetTotalOpened returns the TotalOpened field value if set, zero value otherwise.
func (o *ContactActivity) GetTotalOpened() int32 {
	if o == nil || o.TotalOpened == nil {
		var ret int32
		return ret
	}
	return *o.TotalOpened
}

// GetTotalOpenedOk returns a tuple with the TotalOpened field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactActivity) GetTotalOpenedOk() (*int32, bool) {
	if o == nil || o.TotalOpened == nil {
		return nil, false
	}
	return o.TotalOpened, true
}

// HasTotalOpened returns a boolean if a field has been set.
func (o *ContactActivity) HasTotalOpened() bool {
	if o != nil && o.TotalOpened != nil {
		return true
	}

	return false
}

// SetTotalOpened gets a reference to the given int32 and assigns it to the TotalOpened field.
func (o *ContactActivity) SetTotalOpened(v int32) {
	o.TotalOpened = &v
}

// GetTotalClicked returns the TotalClicked field value if set, zero value otherwise.
func (o *ContactActivity) GetTotalClicked() int32 {
	if o == nil || o.TotalClicked == nil {
		var ret int32
		return ret
	}
	return *o.TotalClicked
}

// GetTotalClickedOk returns a tuple with the TotalClicked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactActivity) GetTotalClickedOk() (*int32, bool) {
	if o == nil || o.TotalClicked == nil {
		return nil, false
	}
	return o.TotalClicked, true
}

// HasTotalClicked returns a boolean if a field has been set.
func (o *ContactActivity) HasTotalClicked() bool {
	if o != nil && o.TotalClicked != nil {
		return true
	}

	return false
}

// SetTotalClicked gets a reference to the given int32 and assigns it to the TotalClicked field.
func (o *ContactActivity) SetTotalClicked(v int32) {
	o.TotalClicked = &v
}

// GetTotalFailed returns the TotalFailed field value if set, zero value otherwise.
func (o *ContactActivity) GetTotalFailed() int32 {
	if o == nil || o.TotalFailed == nil {
		var ret int32
		return ret
	}
	return *o.TotalFailed
}

// GetTotalFailedOk returns a tuple with the TotalFailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactActivity) GetTotalFailedOk() (*int32, bool) {
	if o == nil || o.TotalFailed == nil {
		return nil, false
	}
	return o.TotalFailed, true
}

// HasTotalFailed returns a boolean if a field has been set.
func (o *ContactActivity) HasTotalFailed() bool {
	if o != nil && o.TotalFailed != nil {
		return true
	}

	return false
}

// SetTotalFailed gets a reference to the given int32 and assigns it to the TotalFailed field.
func (o *ContactActivity) SetTotalFailed(v int32) {
	o.TotalFailed = &v
}

// GetLastSent returns the LastSent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContactActivity) GetLastSent() time.Time {
	if o == nil || o.LastSent.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastSent.Get()
}

// GetLastSentOk returns a tuple with the LastSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContactActivity) GetLastSentOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastSent.Get(), o.LastSent.IsSet()
}

// HasLastSent returns a boolean if a field has been set.
func (o *ContactActivity) HasLastSent() bool {
	if o != nil && o.LastSent.IsSet() {
		return true
	}

	return false
}

// SetLastSent gets a reference to the given NullableTime and assigns it to the LastSent field.
func (o *ContactActivity) SetLastSent(v time.Time) {
	o.LastSent.Set(&v)
}
// SetLastSentNil sets the value for LastSent to be an explicit nil
func (o *ContactActivity) SetLastSentNil() {
	o.LastSent.Set(nil)
}

// UnsetLastSent ensures that no value is present for LastSent, not even an explicit nil
func (o *ContactActivity) UnsetLastSent() {
	o.LastSent.Unset()
}

// GetLastOpened returns the LastOpened field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContactActivity) GetLastOpened() time.Time {
	if o == nil || o.LastOpened.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastOpened.Get()
}

// GetLastOpenedOk returns a tuple with the LastOpened field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContactActivity) GetLastOpenedOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastOpened.Get(), o.LastOpened.IsSet()
}

// HasLastOpened returns a boolean if a field has been set.
func (o *ContactActivity) HasLastOpened() bool {
	if o != nil && o.LastOpened.IsSet() {
		return true
	}

	return false
}

// SetLastOpened gets a reference to the given NullableTime and assigns it to the LastOpened field.
func (o *ContactActivity) SetLastOpened(v time.Time) {
	o.LastOpened.Set(&v)
}
// SetLastOpenedNil sets the value for LastOpened to be an explicit nil
func (o *ContactActivity) SetLastOpenedNil() {
	o.LastOpened.Set(nil)
}

// UnsetLastOpened ensures that no value is present for LastOpened, not even an explicit nil
func (o *ContactActivity) UnsetLastOpened() {
	o.LastOpened.Unset()
}

// GetLastClicked returns the LastClicked field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContactActivity) GetLastClicked() time.Time {
	if o == nil || o.LastClicked.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastClicked.Get()
}

// GetLastClickedOk returns a tuple with the LastClicked field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContactActivity) GetLastClickedOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastClicked.Get(), o.LastClicked.IsSet()
}

// HasLastClicked returns a boolean if a field has been set.
func (o *ContactActivity) HasLastClicked() bool {
	if o != nil && o.LastClicked.IsSet() {
		return true
	}

	return false
}

// SetLastClicked gets a reference to the given NullableTime and assigns it to the LastClicked field.
func (o *ContactActivity) SetLastClicked(v time.Time) {
	o.LastClicked.Set(&v)
}
// SetLastClickedNil sets the value for LastClicked to be an explicit nil
func (o *ContactActivity) SetLastClickedNil() {
	o.LastClicked.Set(nil)
}

// UnsetLastClicked ensures that no value is present for LastClicked, not even an explicit nil
func (o *ContactActivity) UnsetLastClicked() {
	o.LastClicked.Unset()
}

// GetLastFailed returns the LastFailed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContactActivity) GetLastFailed() time.Time {
	if o == nil || o.LastFailed.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastFailed.Get()
}

// GetLastFailedOk returns a tuple with the LastFailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContactActivity) GetLastFailedOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastFailed.Get(), o.LastFailed.IsSet()
}

// HasLastFailed returns a boolean if a field has been set.
func (o *ContactActivity) HasLastFailed() bool {
	if o != nil && o.LastFailed.IsSet() {
		return true
	}

	return false
}

// SetLastFailed gets a reference to the given NullableTime and assigns it to the LastFailed field.
func (o *ContactActivity) SetLastFailed(v time.Time) {
	o.LastFailed.Set(&v)
}
// SetLastFailedNil sets the value for LastFailed to be an explicit nil
func (o *ContactActivity) SetLastFailedNil() {
	o.LastFailed.Set(nil)
}

// UnsetLastFailed ensures that no value is present for LastFailed, not even an explicit nil
func (o *ContactActivity) UnsetLastFailed() {
	o.LastFailed.Unset()
}

// GetLastIP returns the LastIP field value if set, zero value otherwise.
func (o *ContactActivity) GetLastIP() string {
	if o == nil || o.LastIP == nil {
		var ret string
		return ret
	}
	return *o.LastIP
}

// GetLastIPOk returns a tuple with the LastIP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactActivity) GetLastIPOk() (*string, bool) {
	if o == nil || o.LastIP == nil {
		return nil, false
	}
	return o.LastIP, true
}

// HasLastIP returns a boolean if a field has been set.
func (o *ContactActivity) HasLastIP() bool {
	if o != nil && o.LastIP != nil {
		return true
	}

	return false
}

// SetLastIP gets a reference to the given string and assigns it to the LastIP field.
func (o *ContactActivity) SetLastIP(v string) {
	o.LastIP = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContactActivity) GetErrorCode() int32 {
	if o == nil || o.ErrorCode.Get() == nil {
		var ret int32
		return ret
	}
	return *o.ErrorCode.Get()
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContactActivity) GetErrorCodeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ErrorCode.Get(), o.ErrorCode.IsSet()
}

// HasErrorCode returns a boolean if a field has been set.
func (o *ContactActivity) HasErrorCode() bool {
	if o != nil && o.ErrorCode.IsSet() {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given NullableInt32 and assigns it to the ErrorCode field.
func (o *ContactActivity) SetErrorCode(v int32) {
	o.ErrorCode.Set(&v)
}
// SetErrorCodeNil sets the value for ErrorCode to be an explicit nil
func (o *ContactActivity) SetErrorCodeNil() {
	o.ErrorCode.Set(nil)
}

// UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
func (o *ContactActivity) UnsetErrorCode() {
	o.ErrorCode.Unset()
}

// GetFriendlyErrorMessage returns the FriendlyErrorMessage field value if set, zero value otherwise.
func (o *ContactActivity) GetFriendlyErrorMessage() string {
	if o == nil || o.FriendlyErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.FriendlyErrorMessage
}

// GetFriendlyErrorMessageOk returns a tuple with the FriendlyErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactActivity) GetFriendlyErrorMessageOk() (*string, bool) {
	if o == nil || o.FriendlyErrorMessage == nil {
		return nil, false
	}
	return o.FriendlyErrorMessage, true
}

// HasFriendlyErrorMessage returns a boolean if a field has been set.
func (o *ContactActivity) HasFriendlyErrorMessage() bool {
	if o != nil && o.FriendlyErrorMessage != nil {
		return true
	}

	return false
}

// SetFriendlyErrorMessage gets a reference to the given string and assigns it to the FriendlyErrorMessage field.
func (o *ContactActivity) SetFriendlyErrorMessage(v string) {
	o.FriendlyErrorMessage = &v
}

func (o ContactActivity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TotalSent != nil {
		toSerialize["TotalSent"] = o.TotalSent
	}
	if o.TotalOpened != nil {
		toSerialize["TotalOpened"] = o.TotalOpened
	}
	if o.TotalClicked != nil {
		toSerialize["TotalClicked"] = o.TotalClicked
	}
	if o.TotalFailed != nil {
		toSerialize["TotalFailed"] = o.TotalFailed
	}
	if o.LastSent.IsSet() {
		toSerialize["LastSent"] = o.LastSent.Get()
	}
	if o.LastOpened.IsSet() {
		toSerialize["LastOpened"] = o.LastOpened.Get()
	}
	if o.LastClicked.IsSet() {
		toSerialize["LastClicked"] = o.LastClicked.Get()
	}
	if o.LastFailed.IsSet() {
		toSerialize["LastFailed"] = o.LastFailed.Get()
	}
	if o.LastIP != nil {
		toSerialize["LastIP"] = o.LastIP
	}
	if o.ErrorCode.IsSet() {
		toSerialize["ErrorCode"] = o.ErrorCode.Get()
	}
	if o.FriendlyErrorMessage != nil {
		toSerialize["FriendlyErrorMessage"] = o.FriendlyErrorMessage
	}
	return json.Marshal(toSerialize)
}

type NullableContactActivity struct {
	value *ContactActivity
	isSet bool
}

func (v NullableContactActivity) Get() *ContactActivity {
	return v.value
}

func (v *NullableContactActivity) Set(val *ContactActivity) {
	v.value = val
	v.isSet = true
}

func (v NullableContactActivity) IsSet() bool {
	return v.isSet
}

func (v *NullableContactActivity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactActivity(val *ContactActivity) *NullableContactActivity {
	return &NullableContactActivity{value: val, isSet: true}
}

func (v NullableContactActivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactActivity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

