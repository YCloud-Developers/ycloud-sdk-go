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

// WhatsappBusinessAccount Represents a specific [WhatsApp Business Account (WABA)](https://www.facebook.com/business/help/1499554293524119).
type WhatsappBusinessAccount struct {
	// ID of the WhatApp Business Account.
	Id *string `json:"id,omitempty"`
	// User-friendly name to differentiate WhatsApp Business Accounts.
	Name *string `json:"name,omitempty"`
	// The currency in which the payment transactions for the WhatsApp Business Account will be processed.
	Currency *string `json:"currency,omitempty"`
	// Namespace string for the message templates that belong to the WhatsApp Business Account.
	MessageTemplateNamespace *string `json:"messageTemplateNamespace,omitempty"`
	AccountReviewStatus *WhatsappBusinessAccountReviewStatus `json:"accountReviewStatus,omitempty"`
	BusinessVerificationStatus *MetaBusinessAccountVerificationStatus `json:"businessVerificationStatus,omitempty"`
	// Country of the WhatsApp Business Account's owning Meta Business account.
	Country *string `json:"country,omitempty"`
	// Ownership type of the WhatsApp Business Account.
	OwnershipType *string `json:"ownershipType,omitempty"`
	// Primary funding ID for the WhatsApp Business Account paid service.
	PrimaryFundingId *string `json:"primaryFundingId,omitempty"`
	// The purchase order number supplied by the business for payment management purposes.
	PurchaseOrderNumber *string `json:"purchaseOrderNumber,omitempty"`
	// The timezone of the WhatsApp Business Account.
	TimezoneId *string `json:"timezoneId,omitempty"`
	Decision *WhatsappReviewDecision `json:"decision,omitempty"`
	UpdateEvent *WhatsappBusinessAccountUpdateEventEnum `json:"updateEvent,omitempty"`
	BanState *WhatsappBusinessAccountBanState `json:"banState,omitempty"`
	// The date when the WABA is banned.
	BanDate *string `json:"banDate,omitempty"`
	// Used to report violations imposed on the WABA. See also [WhatsApp Business Platform Policy Violations](https://developers.facebook.com/docs/whatsapp/overview/policy-enforcement/violations).
	ViolationType *string `json:"violationType,omitempty"`
	// Used to report restrictions imposed on the WABA, when that WABA violates [WhatsApp Business Platform policies](https://developers.facebook.com/docs/whatsapp/overview/policy-enforcement).
	Restrictions []WhatsappBusinessAccountRestrictionInfo `json:"restrictions,omitempty"`
}

// NewWhatsappBusinessAccount instantiates a new WhatsappBusinessAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsappBusinessAccount() *WhatsappBusinessAccount {
	this := WhatsappBusinessAccount{}
	return &this
}

// NewWhatsappBusinessAccountWithDefaults instantiates a new WhatsappBusinessAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsappBusinessAccountWithDefaults() *WhatsappBusinessAccount {
	this := WhatsappBusinessAccount{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WhatsappBusinessAccount) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappBusinessAccount) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WhatsappBusinessAccount) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WhatsappBusinessAccount) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WhatsappBusinessAccount) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappBusinessAccount) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WhatsappBusinessAccount) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WhatsappBusinessAccount) SetName(v string) {
	o.Name = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *WhatsappBusinessAccount) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappBusinessAccount) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *WhatsappBusinessAccount) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *WhatsappBusinessAccount) SetCurrency(v string) {
	o.Currency = &v
}

// GetMessageTemplateNamespace returns the MessageTemplateNamespace field value if set, zero value otherwise.
func (o *WhatsappBusinessAccount) GetMessageTemplateNamespace() string {
	if o == nil || o.MessageTemplateNamespace == nil {
		var ret string
		return ret
	}
	return *o.MessageTemplateNamespace
}

// GetMessageTemplateNamespaceOk returns a tuple with the MessageTemplateNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappBusinessAccount) GetMessageTemplateNamespaceOk() (*string, bool) {
	if o == nil || o.MessageTemplateNamespace == nil {
		return nil, false
	}
	return o.MessageTemplateNamespace, true
}

// HasMessageTemplateNamespace returns a boolean if a field has been set.
func (o *WhatsappBusinessAccount) HasMessageTemplateNamespace() bool {
	if o != nil && o.MessageTemplateNamespace != nil {
		return true
	}

	return false
}

// SetMessageTemplateNamespace gets a reference to the given string and assigns it to the MessageTemplateNamespace field.
func (o *WhatsappBusinessAccount) SetMessageTemplateNamespace(v string) {
	o.MessageTemplateNamespace = &v
}

// GetAccountReviewStatus returns the AccountReviewStatus field value if set, zero value otherwise.
func (o *WhatsappBusinessAccount) GetAccountReviewStatus() WhatsappBusinessAccountReviewStatus {
	if o == nil || o.AccountReviewStatus == nil {
		var ret WhatsappBusinessAccountReviewStatus
		return ret
	}
	return *o.AccountReviewStatus
}

// GetAccountReviewStatusOk returns a tuple with the AccountReviewStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappBusinessAccount) GetAccountReviewStatusOk() (*WhatsappBusinessAccountReviewStatus, bool) {
	if o == nil || o.AccountReviewStatus == nil {
		return nil, false
	}
	return o.AccountReviewStatus, true
}

// HasAccountReviewStatus returns a boolean if a field has been set.
func (o *WhatsappBusinessAccount) HasAccountReviewStatus() bool {
	if o != nil && o.AccountReviewStatus != nil {
		return true
	}

	return false
}

// SetAccountReviewStatus gets a reference to the given WhatsappBusinessAccountReviewStatus and assigns it to the AccountReviewStatus field.
func (o *WhatsappBusinessAccount) SetAccountReviewStatus(v WhatsappBusinessAccountReviewStatus) {
	o.AccountReviewStatus = &v
}

// GetBusinessVerificationStatus returns the BusinessVerificationStatus field value if set, zero value otherwise.
func (o *WhatsappBusinessAccount) GetBusinessVerificationStatus() MetaBusinessAccountVerificationStatus {
	if o == nil || o.BusinessVerificationStatus == nil {
		var ret MetaBusinessAccountVerificationStatus
		return ret
	}
	return *o.BusinessVerificationStatus
}

// GetBusinessVerificationStatusOk returns a tuple with the BusinessVerificationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappBusinessAccount) GetBusinessVerificationStatusOk() (*MetaBusinessAccountVerificationStatus, bool) {
	if o == nil || o.BusinessVerificationStatus == nil {
		return nil, false
	}
	return o.BusinessVerificationStatus, true
}

// HasBusinessVerificationStatus returns a boolean if a field has been set.
func (o *WhatsappBusinessAccount) HasBusinessVerificationStatus() bool {
	if o != nil && o.BusinessVerificationStatus != nil {
		return true
	}

	return false
}

// SetBusinessVerificationStatus gets a reference to the given MetaBusinessAccountVerificationStatus and assigns it to the BusinessVerificationStatus field.
func (o *WhatsappBusinessAccount) SetBusinessVerificationStatus(v MetaBusinessAccountVerificationStatus) {
	o.BusinessVerificationStatus = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *WhatsappBusinessAccount) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappBusinessAccount) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *WhatsappBusinessAccount) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *WhatsappBusinessAccount) SetCountry(v string) {
	o.Country = &v
}

// GetOwnershipType returns the OwnershipType field value if set, zero value otherwise.
func (o *WhatsappBusinessAccount) GetOwnershipType() string {
	if o == nil || o.OwnershipType == nil {
		var ret string
		return ret
	}
	return *o.OwnershipType
}

// GetOwnershipTypeOk returns a tuple with the OwnershipType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappBusinessAccount) GetOwnershipTypeOk() (*string, bool) {
	if o == nil || o.OwnershipType == nil {
		return nil, false
	}
	return o.OwnershipType, true
}

// HasOwnershipType returns a boolean if a field has been set.
func (o *WhatsappBusinessAccount) HasOwnershipType() bool {
	if o != nil && o.OwnershipType != nil {
		return true
	}

	return false
}

// SetOwnershipType gets a reference to the given string and assigns it to the OwnershipType field.
func (o *WhatsappBusinessAccount) SetOwnershipType(v string) {
	o.OwnershipType = &v
}

// GetPrimaryFundingId returns the PrimaryFundingId field value if set, zero value otherwise.
func (o *WhatsappBusinessAccount) GetPrimaryFundingId() string {
	if o == nil || o.PrimaryFundingId == nil {
		var ret string
		return ret
	}
	return *o.PrimaryFundingId
}

// GetPrimaryFundingIdOk returns a tuple with the PrimaryFundingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappBusinessAccount) GetPrimaryFundingIdOk() (*string, bool) {
	if o == nil || o.PrimaryFundingId == nil {
		return nil, false
	}
	return o.PrimaryFundingId, true
}

// HasPrimaryFundingId returns a boolean if a field has been set.
func (o *WhatsappBusinessAccount) HasPrimaryFundingId() bool {
	if o != nil && o.PrimaryFundingId != nil {
		return true
	}

	return false
}

// SetPrimaryFundingId gets a reference to the given string and assigns it to the PrimaryFundingId field.
func (o *WhatsappBusinessAccount) SetPrimaryFundingId(v string) {
	o.PrimaryFundingId = &v
}

// GetPurchaseOrderNumber returns the PurchaseOrderNumber field value if set, zero value otherwise.
func (o *WhatsappBusinessAccount) GetPurchaseOrderNumber() string {
	if o == nil || o.PurchaseOrderNumber == nil {
		var ret string
		return ret
	}
	return *o.PurchaseOrderNumber
}

// GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappBusinessAccount) GetPurchaseOrderNumberOk() (*string, bool) {
	if o == nil || o.PurchaseOrderNumber == nil {
		return nil, false
	}
	return o.PurchaseOrderNumber, true
}

// HasPurchaseOrderNumber returns a boolean if a field has been set.
func (o *WhatsappBusinessAccount) HasPurchaseOrderNumber() bool {
	if o != nil && o.PurchaseOrderNumber != nil {
		return true
	}

	return false
}

// SetPurchaseOrderNumber gets a reference to the given string and assigns it to the PurchaseOrderNumber field.
func (o *WhatsappBusinessAccount) SetPurchaseOrderNumber(v string) {
	o.PurchaseOrderNumber = &v
}

// GetTimezoneId returns the TimezoneId field value if set, zero value otherwise.
func (o *WhatsappBusinessAccount) GetTimezoneId() string {
	if o == nil || o.TimezoneId == nil {
		var ret string
		return ret
	}
	return *o.TimezoneId
}

// GetTimezoneIdOk returns a tuple with the TimezoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappBusinessAccount) GetTimezoneIdOk() (*string, bool) {
	if o == nil || o.TimezoneId == nil {
		return nil, false
	}
	return o.TimezoneId, true
}

// HasTimezoneId returns a boolean if a field has been set.
func (o *WhatsappBusinessAccount) HasTimezoneId() bool {
	if o != nil && o.TimezoneId != nil {
		return true
	}

	return false
}

// SetTimezoneId gets a reference to the given string and assigns it to the TimezoneId field.
func (o *WhatsappBusinessAccount) SetTimezoneId(v string) {
	o.TimezoneId = &v
}

// GetDecision returns the Decision field value if set, zero value otherwise.
func (o *WhatsappBusinessAccount) GetDecision() WhatsappReviewDecision {
	if o == nil || o.Decision == nil {
		var ret WhatsappReviewDecision
		return ret
	}
	return *o.Decision
}

// GetDecisionOk returns a tuple with the Decision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappBusinessAccount) GetDecisionOk() (*WhatsappReviewDecision, bool) {
	if o == nil || o.Decision == nil {
		return nil, false
	}
	return o.Decision, true
}

// HasDecision returns a boolean if a field has been set.
func (o *WhatsappBusinessAccount) HasDecision() bool {
	if o != nil && o.Decision != nil {
		return true
	}

	return false
}

// SetDecision gets a reference to the given WhatsappReviewDecision and assigns it to the Decision field.
func (o *WhatsappBusinessAccount) SetDecision(v WhatsappReviewDecision) {
	o.Decision = &v
}

// GetUpdateEvent returns the UpdateEvent field value if set, zero value otherwise.
func (o *WhatsappBusinessAccount) GetUpdateEvent() WhatsappBusinessAccountUpdateEventEnum {
	if o == nil || o.UpdateEvent == nil {
		var ret WhatsappBusinessAccountUpdateEventEnum
		return ret
	}
	return *o.UpdateEvent
}

// GetUpdateEventOk returns a tuple with the UpdateEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappBusinessAccount) GetUpdateEventOk() (*WhatsappBusinessAccountUpdateEventEnum, bool) {
	if o == nil || o.UpdateEvent == nil {
		return nil, false
	}
	return o.UpdateEvent, true
}

// HasUpdateEvent returns a boolean if a field has been set.
func (o *WhatsappBusinessAccount) HasUpdateEvent() bool {
	if o != nil && o.UpdateEvent != nil {
		return true
	}

	return false
}

// SetUpdateEvent gets a reference to the given WhatsappBusinessAccountUpdateEventEnum and assigns it to the UpdateEvent field.
func (o *WhatsappBusinessAccount) SetUpdateEvent(v WhatsappBusinessAccountUpdateEventEnum) {
	o.UpdateEvent = &v
}

// GetBanState returns the BanState field value if set, zero value otherwise.
func (o *WhatsappBusinessAccount) GetBanState() WhatsappBusinessAccountBanState {
	if o == nil || o.BanState == nil {
		var ret WhatsappBusinessAccountBanState
		return ret
	}
	return *o.BanState
}

// GetBanStateOk returns a tuple with the BanState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappBusinessAccount) GetBanStateOk() (*WhatsappBusinessAccountBanState, bool) {
	if o == nil || o.BanState == nil {
		return nil, false
	}
	return o.BanState, true
}

// HasBanState returns a boolean if a field has been set.
func (o *WhatsappBusinessAccount) HasBanState() bool {
	if o != nil && o.BanState != nil {
		return true
	}

	return false
}

// SetBanState gets a reference to the given WhatsappBusinessAccountBanState and assigns it to the BanState field.
func (o *WhatsappBusinessAccount) SetBanState(v WhatsappBusinessAccountBanState) {
	o.BanState = &v
}

// GetBanDate returns the BanDate field value if set, zero value otherwise.
func (o *WhatsappBusinessAccount) GetBanDate() string {
	if o == nil || o.BanDate == nil {
		var ret string
		return ret
	}
	return *o.BanDate
}

// GetBanDateOk returns a tuple with the BanDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappBusinessAccount) GetBanDateOk() (*string, bool) {
	if o == nil || o.BanDate == nil {
		return nil, false
	}
	return o.BanDate, true
}

// HasBanDate returns a boolean if a field has been set.
func (o *WhatsappBusinessAccount) HasBanDate() bool {
	if o != nil && o.BanDate != nil {
		return true
	}

	return false
}

// SetBanDate gets a reference to the given string and assigns it to the BanDate field.
func (o *WhatsappBusinessAccount) SetBanDate(v string) {
	o.BanDate = &v
}

// GetViolationType returns the ViolationType field value if set, zero value otherwise.
func (o *WhatsappBusinessAccount) GetViolationType() string {
	if o == nil || o.ViolationType == nil {
		var ret string
		return ret
	}
	return *o.ViolationType
}

// GetViolationTypeOk returns a tuple with the ViolationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappBusinessAccount) GetViolationTypeOk() (*string, bool) {
	if o == nil || o.ViolationType == nil {
		return nil, false
	}
	return o.ViolationType, true
}

// HasViolationType returns a boolean if a field has been set.
func (o *WhatsappBusinessAccount) HasViolationType() bool {
	if o != nil && o.ViolationType != nil {
		return true
	}

	return false
}

// SetViolationType gets a reference to the given string and assigns it to the ViolationType field.
func (o *WhatsappBusinessAccount) SetViolationType(v string) {
	o.ViolationType = &v
}

// GetRestrictions returns the Restrictions field value if set, zero value otherwise.
func (o *WhatsappBusinessAccount) GetRestrictions() []WhatsappBusinessAccountRestrictionInfo {
	if o == nil || o.Restrictions == nil {
		var ret []WhatsappBusinessAccountRestrictionInfo
		return ret
	}
	return o.Restrictions
}

// GetRestrictionsOk returns a tuple with the Restrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappBusinessAccount) GetRestrictionsOk() ([]WhatsappBusinessAccountRestrictionInfo, bool) {
	if o == nil || o.Restrictions == nil {
		return nil, false
	}
	return o.Restrictions, true
}

// HasRestrictions returns a boolean if a field has been set.
func (o *WhatsappBusinessAccount) HasRestrictions() bool {
	if o != nil && o.Restrictions != nil {
		return true
	}

	return false
}

// SetRestrictions gets a reference to the given []WhatsappBusinessAccountRestrictionInfo and assigns it to the Restrictions field.
func (o *WhatsappBusinessAccount) SetRestrictions(v []WhatsappBusinessAccountRestrictionInfo) {
	o.Restrictions = v
}

func (o WhatsappBusinessAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.MessageTemplateNamespace != nil {
		toSerialize["messageTemplateNamespace"] = o.MessageTemplateNamespace
	}
	if o.AccountReviewStatus != nil {
		toSerialize["accountReviewStatus"] = o.AccountReviewStatus
	}
	if o.BusinessVerificationStatus != nil {
		toSerialize["businessVerificationStatus"] = o.BusinessVerificationStatus
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	if o.OwnershipType != nil {
		toSerialize["ownershipType"] = o.OwnershipType
	}
	if o.PrimaryFundingId != nil {
		toSerialize["primaryFundingId"] = o.PrimaryFundingId
	}
	if o.PurchaseOrderNumber != nil {
		toSerialize["purchaseOrderNumber"] = o.PurchaseOrderNumber
	}
	if o.TimezoneId != nil {
		toSerialize["timezoneId"] = o.TimezoneId
	}
	if o.Decision != nil {
		toSerialize["decision"] = o.Decision
	}
	if o.UpdateEvent != nil {
		toSerialize["updateEvent"] = o.UpdateEvent
	}
	if o.BanState != nil {
		toSerialize["banState"] = o.BanState
	}
	if o.BanDate != nil {
		toSerialize["banDate"] = o.BanDate
	}
	if o.ViolationType != nil {
		toSerialize["violationType"] = o.ViolationType
	}
	if o.Restrictions != nil {
		toSerialize["restrictions"] = o.Restrictions
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsappBusinessAccount struct {
	value *WhatsappBusinessAccount
	isSet bool
}

func (v NullableWhatsappBusinessAccount) Get() *WhatsappBusinessAccount {
	return v.value
}

func (v *NullableWhatsappBusinessAccount) Set(val *WhatsappBusinessAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsappBusinessAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsappBusinessAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsappBusinessAccount(val *WhatsappBusinessAccount) *NullableWhatsappBusinessAccount {
	return &NullableWhatsappBusinessAccount{value: val, isSet: true}
}

func (v NullableWhatsappBusinessAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsappBusinessAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


