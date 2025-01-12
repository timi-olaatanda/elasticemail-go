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

// MessageCategory the model 'MessageCategory'
type MessageCategory string

// List of MessageCategory
const (
	UNKNOWN_MESSAGE_CATEGORY                 MessageCategory = "Unknown"
	IGNORE_MESSAGE_CATEGORY                  MessageCategory = "Ignore"
	SPAM_MESSAGE_CATEGORY                    MessageCategory = "Spam"
	BLACK_LISTED_MESSAGE_CATEGORY            MessageCategory = "BlackListed"
	NO_MAILBOX_MESSAGE_CATEGORY              MessageCategory = "NoMailbox"
	GREY_LISTED_MESSAGE_CATEGORY             MessageCategory = "GreyListed"
	THROTTLED_MESSAGE_CATEGORY               MessageCategory = "Throttled"
	TIMEOUT_MESSAGE_CATEGORY                 MessageCategory = "Timeout"
	CONNECTION_PROBLEM_MESSAGE_CATEGORY      MessageCategory = "ConnectionProblem"
	SPF_PROBLEM_MESSAGE_CATEGORY             MessageCategory = "SPFProblem"
	ACCOUNT_PROBLEM_MESSAGE_CATEGORY         MessageCategory = "AccountProblem"
	DNS_PROBLEM_MESSAGE_CATEGORY             MessageCategory = "DNSProblem"
	NOT_DELIVERED_CANCELLED_MESSAGE_CATEGORY MessageCategory = "NotDeliveredCancelled"
	CODE_ERROR_MESSAGE_CATEGORY              MessageCategory = "CodeError"
	MANUAL_CANCEL_MESSAGE_CATEGORY           MessageCategory = "ManualCancel"
	CONNECTION_TERMINATED_MESSAGE_CATEGORY   MessageCategory = "ConnectionTerminated"
	NOT_DELIVERED_MESSAGE_CATEGORY           MessageCategory = "NotDelivered"
)

// All allowed values of MessageCategory enum
var AllowedMessageCategoryEnumValues = []MessageCategory{
	"Unknown",
	"Ignore",
	"Spam",
	"BlackListed",
	"NoMailbox",
	"GreyListed",
	"Throttled",
	"Timeout",
	"ConnectionProblem",
	"SPFProblem",
	"AccountProblem",
	"DNSProblem",
	"NotDeliveredCancelled",
	"CodeError",
	"ManualCancel",
	"ConnectionTerminated",
	"NotDelivered",
}

func (v *MessageCategory) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MessageCategory(value)
	for _, existing := range AllowedMessageCategoryEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MessageCategory", value)
}

// NewMessageCategoryFromValue returns a pointer to a valid MessageCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMessageCategoryFromValue(v string) (*MessageCategory, error) {
	ev := MessageCategory(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MessageCategory: valid values are %v", v, AllowedMessageCategoryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MessageCategory) IsValid() bool {
	for _, existing := range AllowedMessageCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MessageCategory value
func (v MessageCategory) Ptr() *MessageCategory {
	return &v
}

type NullableMessageCategory struct {
	value *MessageCategory
	isSet bool
}

func (v NullableMessageCategory) Get() *MessageCategory {
	return v.value
}

func (v *NullableMessageCategory) Set(val *MessageCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageCategory(val *MessageCategory) *NullableMessageCategory {
	return &NullableMessageCategory{value: val, isSet: true}
}

func (v NullableMessageCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
