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
	"time"
)

// Contact Represents a contact.
type Contact struct {
	// Unique ID for the object.
	Id string `json:"id"`
	// Contact's nickname.
	Nickname *string `json:"nickname,omitempty"`
	// Two-letter country abbreviation. See [ISO 3166-1 alpha-2 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	CountryCode *string `json:"countryCode,omitempty"`
	// Full country name.
	CountryName *string `json:"countryName,omitempty"`
	// Unique Phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format.
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	// The contact's email address. If present, the email address must be unique.
	Email *string `json:"email,omitempty"`
	// The time at which the contact last sent a message to your business, formatted in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339). e.g., `2022-06-01T12:00:00.000Z`.
	LastSeen *time.Time `json:"lastSeen,omitempty"`
	// The business phone number that the contact last sent a message to.
	LastMessageToPhoneNumber *string `json:"lastMessageToPhoneNumber,omitempty"`
	// Contact's tags.
	Tags []string `json:"tags,omitempty"`
	// The time at which the contact was created, formatted in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339). e.g., `2022-06-01T12:00:00.000Z`.
	CreateTime *time.Time `json:"createTime,omitempty"`
	// Contact's custom attributes.
	CustomAttributes []ContactCustomAttribute `json:"customAttributes,omitempty"`
	// The email address of the contact's owner.
	OwnerEmail *string `json:"ownerEmail,omitempty"`
}

// NewContact instantiates a new Contact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContact(id string) *Contact {
	this := Contact{}
	this.Id = id
	return &this
}

// NewContactWithDefaults instantiates a new Contact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactWithDefaults() *Contact {
	this := Contact{}
	return &this
}

// GetId returns the Id field value
func (o *Contact) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Contact) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Contact) SetId(v string) {
	o.Id = v
}

// GetNickname returns the Nickname field value if set, zero value otherwise.
func (o *Contact) GetNickname() string {
	if o == nil || o.Nickname == nil {
		var ret string
		return ret
	}
	return *o.Nickname
}

// GetNicknameOk returns a tuple with the Nickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contact) GetNicknameOk() (*string, bool) {
	if o == nil || o.Nickname == nil {
		return nil, false
	}
	return o.Nickname, true
}

// HasNickname returns a boolean if a field has been set.
func (o *Contact) HasNickname() bool {
	if o != nil && o.Nickname != nil {
		return true
	}

	return false
}

// SetNickname gets a reference to the given string and assigns it to the Nickname field.
func (o *Contact) SetNickname(v string) {
	o.Nickname = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *Contact) GetCountryCode() string {
	if o == nil || o.CountryCode == nil {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contact) GetCountryCodeOk() (*string, bool) {
	if o == nil || o.CountryCode == nil {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *Contact) HasCountryCode() bool {
	if o != nil && o.CountryCode != nil {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *Contact) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetCountryName returns the CountryName field value if set, zero value otherwise.
func (o *Contact) GetCountryName() string {
	if o == nil || o.CountryName == nil {
		var ret string
		return ret
	}
	return *o.CountryName
}

// GetCountryNameOk returns a tuple with the CountryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contact) GetCountryNameOk() (*string, bool) {
	if o == nil || o.CountryName == nil {
		return nil, false
	}
	return o.CountryName, true
}

// HasCountryName returns a boolean if a field has been set.
func (o *Contact) HasCountryName() bool {
	if o != nil && o.CountryName != nil {
		return true
	}

	return false
}

// SetCountryName gets a reference to the given string and assigns it to the CountryName field.
func (o *Contact) SetCountryName(v string) {
	o.CountryName = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *Contact) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contact) GetPhoneNumberOk() (*string, bool) {
	if o == nil || o.PhoneNumber == nil {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *Contact) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber != nil {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *Contact) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *Contact) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contact) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *Contact) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *Contact) SetEmail(v string) {
	o.Email = &v
}

// GetLastSeen returns the LastSeen field value if set, zero value otherwise.
func (o *Contact) GetLastSeen() time.Time {
	if o == nil || o.LastSeen == nil {
		var ret time.Time
		return ret
	}
	return *o.LastSeen
}

// GetLastSeenOk returns a tuple with the LastSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contact) GetLastSeenOk() (*time.Time, bool) {
	if o == nil || o.LastSeen == nil {
		return nil, false
	}
	return o.LastSeen, true
}

// HasLastSeen returns a boolean if a field has been set.
func (o *Contact) HasLastSeen() bool {
	if o != nil && o.LastSeen != nil {
		return true
	}

	return false
}

// SetLastSeen gets a reference to the given time.Time and assigns it to the LastSeen field.
func (o *Contact) SetLastSeen(v time.Time) {
	o.LastSeen = &v
}

// GetLastMessageToPhoneNumber returns the LastMessageToPhoneNumber field value if set, zero value otherwise.
func (o *Contact) GetLastMessageToPhoneNumber() string {
	if o == nil || o.LastMessageToPhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.LastMessageToPhoneNumber
}

// GetLastMessageToPhoneNumberOk returns a tuple with the LastMessageToPhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contact) GetLastMessageToPhoneNumberOk() (*string, bool) {
	if o == nil || o.LastMessageToPhoneNumber == nil {
		return nil, false
	}
	return o.LastMessageToPhoneNumber, true
}

// HasLastMessageToPhoneNumber returns a boolean if a field has been set.
func (o *Contact) HasLastMessageToPhoneNumber() bool {
	if o != nil && o.LastMessageToPhoneNumber != nil {
		return true
	}

	return false
}

// SetLastMessageToPhoneNumber gets a reference to the given string and assigns it to the LastMessageToPhoneNumber field.
func (o *Contact) SetLastMessageToPhoneNumber(v string) {
	o.LastMessageToPhoneNumber = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Contact) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contact) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Contact) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *Contact) SetTags(v []string) {
	o.Tags = v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *Contact) GetCreateTime() time.Time {
	if o == nil || o.CreateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contact) GetCreateTimeOk() (*time.Time, bool) {
	if o == nil || o.CreateTime == nil {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *Contact) HasCreateTime() bool {
	if o != nil && o.CreateTime != nil {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given time.Time and assigns it to the CreateTime field.
func (o *Contact) SetCreateTime(v time.Time) {
	o.CreateTime = &v
}

// GetCustomAttributes returns the CustomAttributes field value if set, zero value otherwise.
func (o *Contact) GetCustomAttributes() []ContactCustomAttribute {
	if o == nil || o.CustomAttributes == nil {
		var ret []ContactCustomAttribute
		return ret
	}
	return o.CustomAttributes
}

// GetCustomAttributesOk returns a tuple with the CustomAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contact) GetCustomAttributesOk() ([]ContactCustomAttribute, bool) {
	if o == nil || o.CustomAttributes == nil {
		return nil, false
	}
	return o.CustomAttributes, true
}

// HasCustomAttributes returns a boolean if a field has been set.
func (o *Contact) HasCustomAttributes() bool {
	if o != nil && o.CustomAttributes != nil {
		return true
	}

	return false
}

// SetCustomAttributes gets a reference to the given []ContactCustomAttribute and assigns it to the CustomAttributes field.
func (o *Contact) SetCustomAttributes(v []ContactCustomAttribute) {
	o.CustomAttributes = v
}

// GetOwnerEmail returns the OwnerEmail field value if set, zero value otherwise.
func (o *Contact) GetOwnerEmail() string {
	if o == nil || o.OwnerEmail == nil {
		var ret string
		return ret
	}
	return *o.OwnerEmail
}

// GetOwnerEmailOk returns a tuple with the OwnerEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contact) GetOwnerEmailOk() (*string, bool) {
	if o == nil || o.OwnerEmail == nil {
		return nil, false
	}
	return o.OwnerEmail, true
}

// HasOwnerEmail returns a boolean if a field has been set.
func (o *Contact) HasOwnerEmail() bool {
	if o != nil && o.OwnerEmail != nil {
		return true
	}

	return false
}

// SetOwnerEmail gets a reference to the given string and assigns it to the OwnerEmail field.
func (o *Contact) SetOwnerEmail(v string) {
	o.OwnerEmail = &v
}

func (o Contact) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Nickname != nil {
		toSerialize["nickname"] = o.Nickname
	}
	if o.CountryCode != nil {
		toSerialize["countryCode"] = o.CountryCode
	}
	if o.CountryName != nil {
		toSerialize["countryName"] = o.CountryName
	}
	if o.PhoneNumber != nil {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.LastSeen != nil {
		toSerialize["lastSeen"] = o.LastSeen
	}
	if o.LastMessageToPhoneNumber != nil {
		toSerialize["lastMessageToPhoneNumber"] = o.LastMessageToPhoneNumber
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.CreateTime != nil {
		toSerialize["createTime"] = o.CreateTime
	}
	if o.CustomAttributes != nil {
		toSerialize["customAttributes"] = o.CustomAttributes
	}
	if o.OwnerEmail != nil {
		toSerialize["ownerEmail"] = o.OwnerEmail
	}
	return json.Marshal(toSerialize)
}

type NullableContact struct {
	value *Contact
	isSet bool
}

func (v NullableContact) Get() *Contact {
	return v.value
}

func (v *NullableContact) Set(val *Contact) {
	v.value = val
	v.isSet = true
}

func (v NullableContact) IsSet() bool {
	return v.isSet
}

func (v *NullableContact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContact(val *Contact) *NullableContact {
	return &NullableContact{value: val, isSet: true}
}

func (v NullableContact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


