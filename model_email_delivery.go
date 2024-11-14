/*
YCloud API

The [YCloud](https://ycloud.com) API is organized around [REST](https://en.wikipedia.org/wiki/Representational_state_transfer). Our API is designed to have predictable, resource-oriented URLs, return [JSON](https://www.json.org) responses, and use standard HTTP response codes and verbs.

API version: v2
Contact: service@ycloud.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ycloud

import (
	"encoding/json"
)

// EmailDelivery Represents an email delivery report.
type EmailDelivery struct {
	// Unique ID for the related email you've previously sent.
	EmailId string `json:"emailId"`
	// A recipient's email address.
	RecipientAddress string `json:"recipientAddress"`
	// Delivery status of the email to the specific recipient address. - `sending`: The messaging request is accepted by our system. - `failed`: The message failed to be sent from our system. - `sent`: The message has been sent from our system. - `delivered`: Our system has received a delivery receipt indicating that message is delivered. - `undelivered`: Our system has received a delivery receipt indicating that message is not delivered.
	Status *string `json:"status,omitempty"`
	// Error code when the email is undeliverable.
	ErrorCode *string `json:"errorCode,omitempty"`
	// Error message when the email is undeliverable.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// The `externalId` you set when you sent the email.
	ExternalId *string `json:"externalId,omitempty"`
	// This can be either empty or one of `email`, or `verify`. Defaults to `email`. - `email`: Indicates that the message is sent via the **Email** product. - `verify`: Indicates that the message is sent via the **Verify** product.
	BizType *string `json:"bizType,omitempty"`
	// The verification ID. Included only when `bizType` is `verify`.
	VerificationId *string `json:"verificationId,omitempty"`
}

// NewEmailDelivery instantiates a new EmailDelivery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailDelivery(emailId string, recipientAddress string) *EmailDelivery {
	this := EmailDelivery{}
	this.EmailId = emailId
	this.RecipientAddress = recipientAddress
	return &this
}

// NewEmailDeliveryWithDefaults instantiates a new EmailDelivery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailDeliveryWithDefaults() *EmailDelivery {
	this := EmailDelivery{}
	return &this
}

// GetEmailId returns the EmailId field value
func (o *EmailDelivery) GetEmailId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmailId
}

// GetEmailIdOk returns a tuple with the EmailId field value
// and a boolean to check if the value has been set.
func (o *EmailDelivery) GetEmailIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailId, true
}

// SetEmailId sets field value
func (o *EmailDelivery) SetEmailId(v string) {
	o.EmailId = v
}

// GetRecipientAddress returns the RecipientAddress field value
func (o *EmailDelivery) GetRecipientAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecipientAddress
}

// GetRecipientAddressOk returns a tuple with the RecipientAddress field value
// and a boolean to check if the value has been set.
func (o *EmailDelivery) GetRecipientAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecipientAddress, true
}

// SetRecipientAddress sets field value
func (o *EmailDelivery) SetRecipientAddress(v string) {
	o.RecipientAddress = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EmailDelivery) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailDelivery) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EmailDelivery) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *EmailDelivery) SetStatus(v string) {
	o.Status = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *EmailDelivery) GetErrorCode() string {
	if o == nil || o.ErrorCode == nil {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailDelivery) GetErrorCodeOk() (*string, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *EmailDelivery) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *EmailDelivery) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *EmailDelivery) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailDelivery) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *EmailDelivery) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *EmailDelivery) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *EmailDelivery) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailDelivery) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *EmailDelivery) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *EmailDelivery) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetBizType returns the BizType field value if set, zero value otherwise.
func (o *EmailDelivery) GetBizType() string {
	if o == nil || o.BizType == nil {
		var ret string
		return ret
	}
	return *o.BizType
}

// GetBizTypeOk returns a tuple with the BizType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailDelivery) GetBizTypeOk() (*string, bool) {
	if o == nil || o.BizType == nil {
		return nil, false
	}
	return o.BizType, true
}

// HasBizType returns a boolean if a field has been set.
func (o *EmailDelivery) HasBizType() bool {
	if o != nil && o.BizType != nil {
		return true
	}

	return false
}

// SetBizType gets a reference to the given string and assigns it to the BizType field.
func (o *EmailDelivery) SetBizType(v string) {
	o.BizType = &v
}

// GetVerificationId returns the VerificationId field value if set, zero value otherwise.
func (o *EmailDelivery) GetVerificationId() string {
	if o == nil || o.VerificationId == nil {
		var ret string
		return ret
	}
	return *o.VerificationId
}

// GetVerificationIdOk returns a tuple with the VerificationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailDelivery) GetVerificationIdOk() (*string, bool) {
	if o == nil || o.VerificationId == nil {
		return nil, false
	}
	return o.VerificationId, true
}

// HasVerificationId returns a boolean if a field has been set.
func (o *EmailDelivery) HasVerificationId() bool {
	if o != nil && o.VerificationId != nil {
		return true
	}

	return false
}

// SetVerificationId gets a reference to the given string and assigns it to the VerificationId field.
func (o *EmailDelivery) SetVerificationId(v string) {
	o.VerificationId = &v
}

func (o EmailDelivery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["emailId"] = o.EmailId
	}
	if true {
		toSerialize["recipientAddress"] = o.RecipientAddress
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.ErrorCode != nil {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if o.ErrorMessage != nil {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if o.ExternalId != nil {
		toSerialize["externalId"] = o.ExternalId
	}
	if o.BizType != nil {
		toSerialize["bizType"] = o.BizType
	}
	if o.VerificationId != nil {
		toSerialize["verificationId"] = o.VerificationId
	}
	return json.Marshal(toSerialize)
}

type NullableEmailDelivery struct {
	value *EmailDelivery
	isSet bool
}

func (v NullableEmailDelivery) Get() *EmailDelivery {
	return v.value
}

func (v *NullableEmailDelivery) Set(val *EmailDelivery) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailDelivery) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailDelivery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailDelivery(val *EmailDelivery) *NullableEmailDelivery {
	return &NullableEmailDelivery{value: val, isSet: true}
}

func (v NullableEmailDelivery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailDelivery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
