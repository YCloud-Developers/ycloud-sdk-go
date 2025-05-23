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

// WhatsappPhoneNumber See [WhatsApp Business Phone Number](https://developers.facebook.com/docs/whatsapp/cloud-api/phone-numbers)
type WhatsappPhoneNumber struct {
	// Phone number ID.
	Id *string `json:"id,omitempty"`
	// Phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format.
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	// Display phone number.
	DisplayPhoneNumber *string `json:"displayPhoneNumber,omitempty"`
	// WhatsApp Business Account ID.
	WabaId *string `json:"wabaId,omitempty"`
	QualityRating *WhatsappPhoneNumberQualityRating `json:"qualityRating,omitempty"`
	// Messaging limits determine the maximum number of business-initiated conversations each phone number can start in a rolling 24-hour period. See also [Messaging Limits](https://developers.facebook.com/docs/whatsapp/messaging-limits). - `TIER_NOT_SET`: Unknown limit. - `TIER_50`: 50 business-initiated conversations in a rolling 24-hour period. - `TIER_250`: 250 business-initiated conversations in a rolling 24-hour period. - `TIER_1K`: 1K business-initiated conversations with unique customers in a rolling 24-hour period. - `TIER_10K`: 10K business-initiated conversations with unique customers in a rolling 24-hour period. - `TIER_100K`: 100K business-initiated conversations with unique customers in a rolling 24-hour period. - `TIER_UNLIMITED`: An unlimited number of business-initiated conversations in a rolling 24-hour period.
	MessagingLimit *string `json:"messagingLimit,omitempty"`
	// Verified name.
	VerifiedName *string `json:"verifiedName,omitempty"`
	CodeVerificationStatus *WhatsappPhoneNumberCodeVerificationStatus `json:"codeVerificationStatus,omitempty"`
	// Whether this phone number is an official business account or not. An official business account has a green checkmark badge in its profile and chat thread headers. See [Official Business Account](https://developers.facebook.com/docs/whatsapp/overview/business-accounts#official-business-account) for more information.
	IsOfficialBusinessAccount *bool `json:"isOfficialBusinessAccount,omitempty"`
	Status *WhatsappPhoneNumberStatus `json:"status,omitempty"`
	NameStatus *WhatsappPhoneNumberNameStatus `json:"nameStatus,omitempty"`
	NewNameStatus *WhatsappPhoneNumberNameStatus `json:"newNameStatus,omitempty"`
	Decision *WhatsappReviewDecision `json:"decision,omitempty"`
	// Last requested verified name.
	RequestedVerifiedName *string `json:"requestedVerifiedName,omitempty"`
	// Rejection reason.
	RejectionReason *string `json:"rejectionReason,omitempty"`
	QualityUpdateEvent *WhatsappPhoneNumberQualityUpdateEventEnum `json:"qualityUpdateEvent,omitempty"`
}

// NewWhatsappPhoneNumber instantiates a new WhatsappPhoneNumber object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsappPhoneNumber() *WhatsappPhoneNumber {
	this := WhatsappPhoneNumber{}
	return &this
}

// NewWhatsappPhoneNumberWithDefaults instantiates a new WhatsappPhoneNumber object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsappPhoneNumberWithDefaults() *WhatsappPhoneNumber {
	this := WhatsappPhoneNumber{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WhatsappPhoneNumber) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappPhoneNumber) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WhatsappPhoneNumber) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WhatsappPhoneNumber) SetId(v string) {
	o.Id = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *WhatsappPhoneNumber) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappPhoneNumber) GetPhoneNumberOk() (*string, bool) {
	if o == nil || o.PhoneNumber == nil {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *WhatsappPhoneNumber) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber != nil {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *WhatsappPhoneNumber) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetDisplayPhoneNumber returns the DisplayPhoneNumber field value if set, zero value otherwise.
func (o *WhatsappPhoneNumber) GetDisplayPhoneNumber() string {
	if o == nil || o.DisplayPhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.DisplayPhoneNumber
}

// GetDisplayPhoneNumberOk returns a tuple with the DisplayPhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappPhoneNumber) GetDisplayPhoneNumberOk() (*string, bool) {
	if o == nil || o.DisplayPhoneNumber == nil {
		return nil, false
	}
	return o.DisplayPhoneNumber, true
}

// HasDisplayPhoneNumber returns a boolean if a field has been set.
func (o *WhatsappPhoneNumber) HasDisplayPhoneNumber() bool {
	if o != nil && o.DisplayPhoneNumber != nil {
		return true
	}

	return false
}

// SetDisplayPhoneNumber gets a reference to the given string and assigns it to the DisplayPhoneNumber field.
func (o *WhatsappPhoneNumber) SetDisplayPhoneNumber(v string) {
	o.DisplayPhoneNumber = &v
}

// GetWabaId returns the WabaId field value if set, zero value otherwise.
func (o *WhatsappPhoneNumber) GetWabaId() string {
	if o == nil || o.WabaId == nil {
		var ret string
		return ret
	}
	return *o.WabaId
}

// GetWabaIdOk returns a tuple with the WabaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappPhoneNumber) GetWabaIdOk() (*string, bool) {
	if o == nil || o.WabaId == nil {
		return nil, false
	}
	return o.WabaId, true
}

// HasWabaId returns a boolean if a field has been set.
func (o *WhatsappPhoneNumber) HasWabaId() bool {
	if o != nil && o.WabaId != nil {
		return true
	}

	return false
}

// SetWabaId gets a reference to the given string and assigns it to the WabaId field.
func (o *WhatsappPhoneNumber) SetWabaId(v string) {
	o.WabaId = &v
}

// GetQualityRating returns the QualityRating field value if set, zero value otherwise.
func (o *WhatsappPhoneNumber) GetQualityRating() WhatsappPhoneNumberQualityRating {
	if o == nil || o.QualityRating == nil {
		var ret WhatsappPhoneNumberQualityRating
		return ret
	}
	return *o.QualityRating
}

// GetQualityRatingOk returns a tuple with the QualityRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappPhoneNumber) GetQualityRatingOk() (*WhatsappPhoneNumberQualityRating, bool) {
	if o == nil || o.QualityRating == nil {
		return nil, false
	}
	return o.QualityRating, true
}

// HasQualityRating returns a boolean if a field has been set.
func (o *WhatsappPhoneNumber) HasQualityRating() bool {
	if o != nil && o.QualityRating != nil {
		return true
	}

	return false
}

// SetQualityRating gets a reference to the given WhatsappPhoneNumberQualityRating and assigns it to the QualityRating field.
func (o *WhatsappPhoneNumber) SetQualityRating(v WhatsappPhoneNumberQualityRating) {
	o.QualityRating = &v
}

// GetMessagingLimit returns the MessagingLimit field value if set, zero value otherwise.
func (o *WhatsappPhoneNumber) GetMessagingLimit() string {
	if o == nil || o.MessagingLimit == nil {
		var ret string
		return ret
	}
	return *o.MessagingLimit
}

// GetMessagingLimitOk returns a tuple with the MessagingLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappPhoneNumber) GetMessagingLimitOk() (*string, bool) {
	if o == nil || o.MessagingLimit == nil {
		return nil, false
	}
	return o.MessagingLimit, true
}

// HasMessagingLimit returns a boolean if a field has been set.
func (o *WhatsappPhoneNumber) HasMessagingLimit() bool {
	if o != nil && o.MessagingLimit != nil {
		return true
	}

	return false
}

// SetMessagingLimit gets a reference to the given string and assigns it to the MessagingLimit field.
func (o *WhatsappPhoneNumber) SetMessagingLimit(v string) {
	o.MessagingLimit = &v
}

// GetVerifiedName returns the VerifiedName field value if set, zero value otherwise.
func (o *WhatsappPhoneNumber) GetVerifiedName() string {
	if o == nil || o.VerifiedName == nil {
		var ret string
		return ret
	}
	return *o.VerifiedName
}

// GetVerifiedNameOk returns a tuple with the VerifiedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappPhoneNumber) GetVerifiedNameOk() (*string, bool) {
	if o == nil || o.VerifiedName == nil {
		return nil, false
	}
	return o.VerifiedName, true
}

// HasVerifiedName returns a boolean if a field has been set.
func (o *WhatsappPhoneNumber) HasVerifiedName() bool {
	if o != nil && o.VerifiedName != nil {
		return true
	}

	return false
}

// SetVerifiedName gets a reference to the given string and assigns it to the VerifiedName field.
func (o *WhatsappPhoneNumber) SetVerifiedName(v string) {
	o.VerifiedName = &v
}

// GetCodeVerificationStatus returns the CodeVerificationStatus field value if set, zero value otherwise.
func (o *WhatsappPhoneNumber) GetCodeVerificationStatus() WhatsappPhoneNumberCodeVerificationStatus {
	if o == nil || o.CodeVerificationStatus == nil {
		var ret WhatsappPhoneNumberCodeVerificationStatus
		return ret
	}
	return *o.CodeVerificationStatus
}

// GetCodeVerificationStatusOk returns a tuple with the CodeVerificationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappPhoneNumber) GetCodeVerificationStatusOk() (*WhatsappPhoneNumberCodeVerificationStatus, bool) {
	if o == nil || o.CodeVerificationStatus == nil {
		return nil, false
	}
	return o.CodeVerificationStatus, true
}

// HasCodeVerificationStatus returns a boolean if a field has been set.
func (o *WhatsappPhoneNumber) HasCodeVerificationStatus() bool {
	if o != nil && o.CodeVerificationStatus != nil {
		return true
	}

	return false
}

// SetCodeVerificationStatus gets a reference to the given WhatsappPhoneNumberCodeVerificationStatus and assigns it to the CodeVerificationStatus field.
func (o *WhatsappPhoneNumber) SetCodeVerificationStatus(v WhatsappPhoneNumberCodeVerificationStatus) {
	o.CodeVerificationStatus = &v
}

// GetIsOfficialBusinessAccount returns the IsOfficialBusinessAccount field value if set, zero value otherwise.
func (o *WhatsappPhoneNumber) GetIsOfficialBusinessAccount() bool {
	if o == nil || o.IsOfficialBusinessAccount == nil {
		var ret bool
		return ret
	}
	return *o.IsOfficialBusinessAccount
}

// GetIsOfficialBusinessAccountOk returns a tuple with the IsOfficialBusinessAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappPhoneNumber) GetIsOfficialBusinessAccountOk() (*bool, bool) {
	if o == nil || o.IsOfficialBusinessAccount == nil {
		return nil, false
	}
	return o.IsOfficialBusinessAccount, true
}

// HasIsOfficialBusinessAccount returns a boolean if a field has been set.
func (o *WhatsappPhoneNumber) HasIsOfficialBusinessAccount() bool {
	if o != nil && o.IsOfficialBusinessAccount != nil {
		return true
	}

	return false
}

// SetIsOfficialBusinessAccount gets a reference to the given bool and assigns it to the IsOfficialBusinessAccount field.
func (o *WhatsappPhoneNumber) SetIsOfficialBusinessAccount(v bool) {
	o.IsOfficialBusinessAccount = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WhatsappPhoneNumber) GetStatus() WhatsappPhoneNumberStatus {
	if o == nil || o.Status == nil {
		var ret WhatsappPhoneNumberStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappPhoneNumber) GetStatusOk() (*WhatsappPhoneNumberStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WhatsappPhoneNumber) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given WhatsappPhoneNumberStatus and assigns it to the Status field.
func (o *WhatsappPhoneNumber) SetStatus(v WhatsappPhoneNumberStatus) {
	o.Status = &v
}

// GetNameStatus returns the NameStatus field value if set, zero value otherwise.
func (o *WhatsappPhoneNumber) GetNameStatus() WhatsappPhoneNumberNameStatus {
	if o == nil || o.NameStatus == nil {
		var ret WhatsappPhoneNumberNameStatus
		return ret
	}
	return *o.NameStatus
}

// GetNameStatusOk returns a tuple with the NameStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappPhoneNumber) GetNameStatusOk() (*WhatsappPhoneNumberNameStatus, bool) {
	if o == nil || o.NameStatus == nil {
		return nil, false
	}
	return o.NameStatus, true
}

// HasNameStatus returns a boolean if a field has been set.
func (o *WhatsappPhoneNumber) HasNameStatus() bool {
	if o != nil && o.NameStatus != nil {
		return true
	}

	return false
}

// SetNameStatus gets a reference to the given WhatsappPhoneNumberNameStatus and assigns it to the NameStatus field.
func (o *WhatsappPhoneNumber) SetNameStatus(v WhatsappPhoneNumberNameStatus) {
	o.NameStatus = &v
}

// GetNewNameStatus returns the NewNameStatus field value if set, zero value otherwise.
func (o *WhatsappPhoneNumber) GetNewNameStatus() WhatsappPhoneNumberNameStatus {
	if o == nil || o.NewNameStatus == nil {
		var ret WhatsappPhoneNumberNameStatus
		return ret
	}
	return *o.NewNameStatus
}

// GetNewNameStatusOk returns a tuple with the NewNameStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappPhoneNumber) GetNewNameStatusOk() (*WhatsappPhoneNumberNameStatus, bool) {
	if o == nil || o.NewNameStatus == nil {
		return nil, false
	}
	return o.NewNameStatus, true
}

// HasNewNameStatus returns a boolean if a field has been set.
func (o *WhatsappPhoneNumber) HasNewNameStatus() bool {
	if o != nil && o.NewNameStatus != nil {
		return true
	}

	return false
}

// SetNewNameStatus gets a reference to the given WhatsappPhoneNumberNameStatus and assigns it to the NewNameStatus field.
func (o *WhatsappPhoneNumber) SetNewNameStatus(v WhatsappPhoneNumberNameStatus) {
	o.NewNameStatus = &v
}

// GetDecision returns the Decision field value if set, zero value otherwise.
func (o *WhatsappPhoneNumber) GetDecision() WhatsappReviewDecision {
	if o == nil || o.Decision == nil {
		var ret WhatsappReviewDecision
		return ret
	}
	return *o.Decision
}

// GetDecisionOk returns a tuple with the Decision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappPhoneNumber) GetDecisionOk() (*WhatsappReviewDecision, bool) {
	if o == nil || o.Decision == nil {
		return nil, false
	}
	return o.Decision, true
}

// HasDecision returns a boolean if a field has been set.
func (o *WhatsappPhoneNumber) HasDecision() bool {
	if o != nil && o.Decision != nil {
		return true
	}

	return false
}

// SetDecision gets a reference to the given WhatsappReviewDecision and assigns it to the Decision field.
func (o *WhatsappPhoneNumber) SetDecision(v WhatsappReviewDecision) {
	o.Decision = &v
}

// GetRequestedVerifiedName returns the RequestedVerifiedName field value if set, zero value otherwise.
func (o *WhatsappPhoneNumber) GetRequestedVerifiedName() string {
	if o == nil || o.RequestedVerifiedName == nil {
		var ret string
		return ret
	}
	return *o.RequestedVerifiedName
}

// GetRequestedVerifiedNameOk returns a tuple with the RequestedVerifiedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappPhoneNumber) GetRequestedVerifiedNameOk() (*string, bool) {
	if o == nil || o.RequestedVerifiedName == nil {
		return nil, false
	}
	return o.RequestedVerifiedName, true
}

// HasRequestedVerifiedName returns a boolean if a field has been set.
func (o *WhatsappPhoneNumber) HasRequestedVerifiedName() bool {
	if o != nil && o.RequestedVerifiedName != nil {
		return true
	}

	return false
}

// SetRequestedVerifiedName gets a reference to the given string and assigns it to the RequestedVerifiedName field.
func (o *WhatsappPhoneNumber) SetRequestedVerifiedName(v string) {
	o.RequestedVerifiedName = &v
}

// GetRejectionReason returns the RejectionReason field value if set, zero value otherwise.
func (o *WhatsappPhoneNumber) GetRejectionReason() string {
	if o == nil || o.RejectionReason == nil {
		var ret string
		return ret
	}
	return *o.RejectionReason
}

// GetRejectionReasonOk returns a tuple with the RejectionReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappPhoneNumber) GetRejectionReasonOk() (*string, bool) {
	if o == nil || o.RejectionReason == nil {
		return nil, false
	}
	return o.RejectionReason, true
}

// HasRejectionReason returns a boolean if a field has been set.
func (o *WhatsappPhoneNumber) HasRejectionReason() bool {
	if o != nil && o.RejectionReason != nil {
		return true
	}

	return false
}

// SetRejectionReason gets a reference to the given string and assigns it to the RejectionReason field.
func (o *WhatsappPhoneNumber) SetRejectionReason(v string) {
	o.RejectionReason = &v
}

// GetQualityUpdateEvent returns the QualityUpdateEvent field value if set, zero value otherwise.
func (o *WhatsappPhoneNumber) GetQualityUpdateEvent() WhatsappPhoneNumberQualityUpdateEventEnum {
	if o == nil || o.QualityUpdateEvent == nil {
		var ret WhatsappPhoneNumberQualityUpdateEventEnum
		return ret
	}
	return *o.QualityUpdateEvent
}

// GetQualityUpdateEventOk returns a tuple with the QualityUpdateEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappPhoneNumber) GetQualityUpdateEventOk() (*WhatsappPhoneNumberQualityUpdateEventEnum, bool) {
	if o == nil || o.QualityUpdateEvent == nil {
		return nil, false
	}
	return o.QualityUpdateEvent, true
}

// HasQualityUpdateEvent returns a boolean if a field has been set.
func (o *WhatsappPhoneNumber) HasQualityUpdateEvent() bool {
	if o != nil && o.QualityUpdateEvent != nil {
		return true
	}

	return false
}

// SetQualityUpdateEvent gets a reference to the given WhatsappPhoneNumberQualityUpdateEventEnum and assigns it to the QualityUpdateEvent field.
func (o *WhatsappPhoneNumber) SetQualityUpdateEvent(v WhatsappPhoneNumberQualityUpdateEventEnum) {
	o.QualityUpdateEvent = &v
}

func (o WhatsappPhoneNumber) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.PhoneNumber != nil {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	if o.DisplayPhoneNumber != nil {
		toSerialize["displayPhoneNumber"] = o.DisplayPhoneNumber
	}
	if o.WabaId != nil {
		toSerialize["wabaId"] = o.WabaId
	}
	if o.QualityRating != nil {
		toSerialize["qualityRating"] = o.QualityRating
	}
	if o.MessagingLimit != nil {
		toSerialize["messagingLimit"] = o.MessagingLimit
	}
	if o.VerifiedName != nil {
		toSerialize["verifiedName"] = o.VerifiedName
	}
	if o.CodeVerificationStatus != nil {
		toSerialize["codeVerificationStatus"] = o.CodeVerificationStatus
	}
	if o.IsOfficialBusinessAccount != nil {
		toSerialize["isOfficialBusinessAccount"] = o.IsOfficialBusinessAccount
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.NameStatus != nil {
		toSerialize["nameStatus"] = o.NameStatus
	}
	if o.NewNameStatus != nil {
		toSerialize["newNameStatus"] = o.NewNameStatus
	}
	if o.Decision != nil {
		toSerialize["decision"] = o.Decision
	}
	if o.RequestedVerifiedName != nil {
		toSerialize["requestedVerifiedName"] = o.RequestedVerifiedName
	}
	if o.RejectionReason != nil {
		toSerialize["rejectionReason"] = o.RejectionReason
	}
	if o.QualityUpdateEvent != nil {
		toSerialize["qualityUpdateEvent"] = o.QualityUpdateEvent
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsappPhoneNumber struct {
	value *WhatsappPhoneNumber
	isSet bool
}

func (v NullableWhatsappPhoneNumber) Get() *WhatsappPhoneNumber {
	return v.value
}

func (v *NullableWhatsappPhoneNumber) Set(val *WhatsappPhoneNumber) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsappPhoneNumber) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsappPhoneNumber) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsappPhoneNumber(val *WhatsappPhoneNumber) *NullableWhatsappPhoneNumber {
	return &NullableWhatsappPhoneNumber{value: val, isSet: true}
}

func (v NullableWhatsappPhoneNumber) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsappPhoneNumber) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


