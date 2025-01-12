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
	"time"
)

// SubAccountInfo Detailed information about SubAccount.
type SubAccountInfo struct {
	// Public key for limited access to your Account such as contact/add so you can use it safely on public websites.
	PublicAccountID *string `json:"PublicAccountID,omitempty"`
	// Proper email address.
	Email *string `json:"Email,omitempty"`
	Settings *SubaccountSettingsInfo `json:"Settings,omitempty"`
	// Date of last activity on Account
	LastActivity *time.Time `json:"LastActivity,omitempty"`
	// Amount of email credits
	EmailCredits *int32 `json:"EmailCredits,omitempty"`
	// Amount of emails sent from this Account
	TotalEmailsSent *int64 `json:"TotalEmailsSent,omitempty"`
	// Numeric reputation
	Reputation *float64 `json:"Reputation,omitempty"`
	Status *AccountStatusEnum `json:"Status,omitempty"`
	// How many contacts this SubAccount has stored
	ContactsCount *int32 `json:"ContactsCount,omitempty"`
}

// NewSubAccountInfo instantiates a new SubAccountInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubAccountInfo() *SubAccountInfo {
	this := SubAccountInfo{}
	var status AccountStatusEnum = DISABLED
	this.Status = &status
	return &this
}

// NewSubAccountInfoWithDefaults instantiates a new SubAccountInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubAccountInfoWithDefaults() *SubAccountInfo {
	this := SubAccountInfo{}
	var status AccountStatusEnum = DISABLED
	this.Status = &status
	return &this
}

// GetPublicAccountID returns the PublicAccountID field value if set, zero value otherwise.
func (o *SubAccountInfo) GetPublicAccountID() string {
	if o == nil || o.PublicAccountID == nil {
		var ret string
		return ret
	}
	return *o.PublicAccountID
}

// GetPublicAccountIDOk returns a tuple with the PublicAccountID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAccountInfo) GetPublicAccountIDOk() (*string, bool) {
	if o == nil || o.PublicAccountID == nil {
		return nil, false
	}
	return o.PublicAccountID, true
}

// HasPublicAccountID returns a boolean if a field has been set.
func (o *SubAccountInfo) HasPublicAccountID() bool {
	if o != nil && o.PublicAccountID != nil {
		return true
	}

	return false
}

// SetPublicAccountID gets a reference to the given string and assigns it to the PublicAccountID field.
func (o *SubAccountInfo) SetPublicAccountID(v string) {
	o.PublicAccountID = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *SubAccountInfo) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAccountInfo) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *SubAccountInfo) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *SubAccountInfo) SetEmail(v string) {
	o.Email = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *SubAccountInfo) GetSettings() SubaccountSettingsInfo {
	if o == nil || o.Settings == nil {
		var ret SubaccountSettingsInfo
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAccountInfo) GetSettingsOk() (*SubaccountSettingsInfo, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *SubAccountInfo) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given SubaccountSettingsInfo and assigns it to the Settings field.
func (o *SubAccountInfo) SetSettings(v SubaccountSettingsInfo) {
	o.Settings = &v
}

// GetLastActivity returns the LastActivity field value if set, zero value otherwise.
func (o *SubAccountInfo) GetLastActivity() time.Time {
	if o == nil || o.LastActivity == nil {
		var ret time.Time
		return ret
	}
	return *o.LastActivity
}

// GetLastActivityOk returns a tuple with the LastActivity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAccountInfo) GetLastActivityOk() (*time.Time, bool) {
	if o == nil || o.LastActivity == nil {
		return nil, false
	}
	return o.LastActivity, true
}

// HasLastActivity returns a boolean if a field has been set.
func (o *SubAccountInfo) HasLastActivity() bool {
	if o != nil && o.LastActivity != nil {
		return true
	}

	return false
}

// SetLastActivity gets a reference to the given time.Time and assigns it to the LastActivity field.
func (o *SubAccountInfo) SetLastActivity(v time.Time) {
	o.LastActivity = &v
}

// GetEmailCredits returns the EmailCredits field value if set, zero value otherwise.
func (o *SubAccountInfo) GetEmailCredits() int32 {
	if o == nil || o.EmailCredits == nil {
		var ret int32
		return ret
	}
	return *o.EmailCredits
}

// GetEmailCreditsOk returns a tuple with the EmailCredits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAccountInfo) GetEmailCreditsOk() (*int32, bool) {
	if o == nil || o.EmailCredits == nil {
		return nil, false
	}
	return o.EmailCredits, true
}

// HasEmailCredits returns a boolean if a field has been set.
func (o *SubAccountInfo) HasEmailCredits() bool {
	if o != nil && o.EmailCredits != nil {
		return true
	}

	return false
}

// SetEmailCredits gets a reference to the given int32 and assigns it to the EmailCredits field.
func (o *SubAccountInfo) SetEmailCredits(v int32) {
	o.EmailCredits = &v
}

// GetTotalEmailsSent returns the TotalEmailsSent field value if set, zero value otherwise.
func (o *SubAccountInfo) GetTotalEmailsSent() int64 {
	if o == nil || o.TotalEmailsSent == nil {
		var ret int64
		return ret
	}
	return *o.TotalEmailsSent
}

// GetTotalEmailsSentOk returns a tuple with the TotalEmailsSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAccountInfo) GetTotalEmailsSentOk() (*int64, bool) {
	if o == nil || o.TotalEmailsSent == nil {
		return nil, false
	}
	return o.TotalEmailsSent, true
}

// HasTotalEmailsSent returns a boolean if a field has been set.
func (o *SubAccountInfo) HasTotalEmailsSent() bool {
	if o != nil && o.TotalEmailsSent != nil {
		return true
	}

	return false
}

// SetTotalEmailsSent gets a reference to the given int64 and assigns it to the TotalEmailsSent field.
func (o *SubAccountInfo) SetTotalEmailsSent(v int64) {
	o.TotalEmailsSent = &v
}

// GetReputation returns the Reputation field value if set, zero value otherwise.
func (o *SubAccountInfo) GetReputation() float64 {
	if o == nil || o.Reputation == nil {
		var ret float64
		return ret
	}
	return *o.Reputation
}

// GetReputationOk returns a tuple with the Reputation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAccountInfo) GetReputationOk() (*float64, bool) {
	if o == nil || o.Reputation == nil {
		return nil, false
	}
	return o.Reputation, true
}

// HasReputation returns a boolean if a field has been set.
func (o *SubAccountInfo) HasReputation() bool {
	if o != nil && o.Reputation != nil {
		return true
	}

	return false
}

// SetReputation gets a reference to the given float64 and assigns it to the Reputation field.
func (o *SubAccountInfo) SetReputation(v float64) {
	o.Reputation = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SubAccountInfo) GetStatus() AccountStatusEnum {
	if o == nil || o.Status == nil {
		var ret AccountStatusEnum
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAccountInfo) GetStatusOk() (*AccountStatusEnum, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SubAccountInfo) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given AccountStatusEnum and assigns it to the Status field.
func (o *SubAccountInfo) SetStatus(v AccountStatusEnum) {
	o.Status = &v
}

// GetContactsCount returns the ContactsCount field value if set, zero value otherwise.
func (o *SubAccountInfo) GetContactsCount() int32 {
	if o == nil || o.ContactsCount == nil {
		var ret int32
		return ret
	}
	return *o.ContactsCount
}

// GetContactsCountOk returns a tuple with the ContactsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAccountInfo) GetContactsCountOk() (*int32, bool) {
	if o == nil || o.ContactsCount == nil {
		return nil, false
	}
	return o.ContactsCount, true
}

// HasContactsCount returns a boolean if a field has been set.
func (o *SubAccountInfo) HasContactsCount() bool {
	if o != nil && o.ContactsCount != nil {
		return true
	}

	return false
}

// SetContactsCount gets a reference to the given int32 and assigns it to the ContactsCount field.
func (o *SubAccountInfo) SetContactsCount(v int32) {
	o.ContactsCount = &v
}

func (o SubAccountInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PublicAccountID != nil {
		toSerialize["PublicAccountID"] = o.PublicAccountID
	}
	if o.Email != nil {
		toSerialize["Email"] = o.Email
	}
	if o.Settings != nil {
		toSerialize["Settings"] = o.Settings
	}
	if o.LastActivity != nil {
		toSerialize["LastActivity"] = o.LastActivity
	}
	if o.EmailCredits != nil {
		toSerialize["EmailCredits"] = o.EmailCredits
	}
	if o.TotalEmailsSent != nil {
		toSerialize["TotalEmailsSent"] = o.TotalEmailsSent
	}
	if o.Reputation != nil {
		toSerialize["Reputation"] = o.Reputation
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.ContactsCount != nil {
		toSerialize["ContactsCount"] = o.ContactsCount
	}
	return json.Marshal(toSerialize)
}

type NullableSubAccountInfo struct {
	value *SubAccountInfo
	isSet bool
}

func (v NullableSubAccountInfo) Get() *SubAccountInfo {
	return v.value
}

func (v *NullableSubAccountInfo) Set(val *SubAccountInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSubAccountInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSubAccountInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubAccountInfo(val *SubAccountInfo) *NullableSubAccountInfo {
	return &NullableSubAccountInfo{value: val, isSet: true}
}

func (v NullableSubAccountInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubAccountInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


