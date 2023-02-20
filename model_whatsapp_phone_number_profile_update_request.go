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

// WhatsappPhoneNumberProfileUpdateRequest WhatsApp Phone Number Business Profile. Customers can view your business profile by clicking your business's name or number in a conversation thread. See also [Business Profiles](https://developers.facebook.com/docs/whatsapp/cloud-api/reference/business-profiles).
type WhatsappPhoneNumberProfileUpdateRequest struct {
	// The business's **About** text. This text appears in the business's profile, beneath its profile image, phone number, and contact buttons. - String cannot be empty. - Strings must be between 1 and 139 characters. - Rendered emojis are supported however their unicode values are not. Emoji unicode values must be Java- or JavaScript-escape encoded. - Hyperlinks can be included but will not render as clickable links. - Markdown is not supported.
	About *string `json:"about,omitempty"`
	// Address of the business. Character limit 256.
	Address *string `json:"address,omitempty"`
	// Description of the business. Character limit 512.
	Description *string `json:"description,omitempty"`
	// The contact email address (in valid email format) of the business. Character limit 128.
	Email *string `json:"email,omitempty"`
	// URL of the profile picture that was uploaded to Meta.
	ProfilePictureUrl *string `json:"profilePictureUrl,omitempty"`
	Vertical *WhatsappPhoneNumberProfileVertical `json:"vertical,omitempty"`
	// The URLs associated with the business. For instance, a website, Facebook Page, or Instagram. You must include the http:// or https:// portion of the URL. There is a maximum of 2 websites with a maximum of 256 characters each.
	Websites []string `json:"websites,omitempty"`
}

// NewWhatsappPhoneNumberProfileUpdateRequest instantiates a new WhatsappPhoneNumberProfileUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsappPhoneNumberProfileUpdateRequest() *WhatsappPhoneNumberProfileUpdateRequest {
	this := WhatsappPhoneNumberProfileUpdateRequest{}
	return &this
}

// NewWhatsappPhoneNumberProfileUpdateRequestWithDefaults instantiates a new WhatsappPhoneNumberProfileUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsappPhoneNumberProfileUpdateRequestWithDefaults() *WhatsappPhoneNumberProfileUpdateRequest {
	this := WhatsappPhoneNumberProfileUpdateRequest{}
	return &this
}

// GetAbout returns the About field value if set, zero value otherwise.
func (o *WhatsappPhoneNumberProfileUpdateRequest) GetAbout() string {
	if o == nil || o.About == nil {
		var ret string
		return ret
	}
	return *o.About
}

// GetAboutOk returns a tuple with the About field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappPhoneNumberProfileUpdateRequest) GetAboutOk() (*string, bool) {
	if o == nil || o.About == nil {
		return nil, false
	}
	return o.About, true
}

// HasAbout returns a boolean if a field has been set.
func (o *WhatsappPhoneNumberProfileUpdateRequest) HasAbout() bool {
	if o != nil && o.About != nil {
		return true
	}

	return false
}

// SetAbout gets a reference to the given string and assigns it to the About field.
func (o *WhatsappPhoneNumberProfileUpdateRequest) SetAbout(v string) {
	o.About = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *WhatsappPhoneNumberProfileUpdateRequest) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappPhoneNumberProfileUpdateRequest) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *WhatsappPhoneNumberProfileUpdateRequest) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *WhatsappPhoneNumberProfileUpdateRequest) SetAddress(v string) {
	o.Address = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WhatsappPhoneNumberProfileUpdateRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappPhoneNumberProfileUpdateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WhatsappPhoneNumberProfileUpdateRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WhatsappPhoneNumberProfileUpdateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *WhatsappPhoneNumberProfileUpdateRequest) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappPhoneNumberProfileUpdateRequest) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *WhatsappPhoneNumberProfileUpdateRequest) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *WhatsappPhoneNumberProfileUpdateRequest) SetEmail(v string) {
	o.Email = &v
}

// GetProfilePictureUrl returns the ProfilePictureUrl field value if set, zero value otherwise.
func (o *WhatsappPhoneNumberProfileUpdateRequest) GetProfilePictureUrl() string {
	if o == nil || o.ProfilePictureUrl == nil {
		var ret string
		return ret
	}
	return *o.ProfilePictureUrl
}

// GetProfilePictureUrlOk returns a tuple with the ProfilePictureUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappPhoneNumberProfileUpdateRequest) GetProfilePictureUrlOk() (*string, bool) {
	if o == nil || o.ProfilePictureUrl == nil {
		return nil, false
	}
	return o.ProfilePictureUrl, true
}

// HasProfilePictureUrl returns a boolean if a field has been set.
func (o *WhatsappPhoneNumberProfileUpdateRequest) HasProfilePictureUrl() bool {
	if o != nil && o.ProfilePictureUrl != nil {
		return true
	}

	return false
}

// SetProfilePictureUrl gets a reference to the given string and assigns it to the ProfilePictureUrl field.
func (o *WhatsappPhoneNumberProfileUpdateRequest) SetProfilePictureUrl(v string) {
	o.ProfilePictureUrl = &v
}

// GetVertical returns the Vertical field value if set, zero value otherwise.
func (o *WhatsappPhoneNumberProfileUpdateRequest) GetVertical() WhatsappPhoneNumberProfileVertical {
	if o == nil || o.Vertical == nil {
		var ret WhatsappPhoneNumberProfileVertical
		return ret
	}
	return *o.Vertical
}

// GetVerticalOk returns a tuple with the Vertical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappPhoneNumberProfileUpdateRequest) GetVerticalOk() (*WhatsappPhoneNumberProfileVertical, bool) {
	if o == nil || o.Vertical == nil {
		return nil, false
	}
	return o.Vertical, true
}

// HasVertical returns a boolean if a field has been set.
func (o *WhatsappPhoneNumberProfileUpdateRequest) HasVertical() bool {
	if o != nil && o.Vertical != nil {
		return true
	}

	return false
}

// SetVertical gets a reference to the given WhatsappPhoneNumberProfileVertical and assigns it to the Vertical field.
func (o *WhatsappPhoneNumberProfileUpdateRequest) SetVertical(v WhatsappPhoneNumberProfileVertical) {
	o.Vertical = &v
}

// GetWebsites returns the Websites field value if set, zero value otherwise.
func (o *WhatsappPhoneNumberProfileUpdateRequest) GetWebsites() []string {
	if o == nil || o.Websites == nil {
		var ret []string
		return ret
	}
	return o.Websites
}

// GetWebsitesOk returns a tuple with the Websites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappPhoneNumberProfileUpdateRequest) GetWebsitesOk() ([]string, bool) {
	if o == nil || o.Websites == nil {
		return nil, false
	}
	return o.Websites, true
}

// HasWebsites returns a boolean if a field has been set.
func (o *WhatsappPhoneNumberProfileUpdateRequest) HasWebsites() bool {
	if o != nil && o.Websites != nil {
		return true
	}

	return false
}

// SetWebsites gets a reference to the given []string and assigns it to the Websites field.
func (o *WhatsappPhoneNumberProfileUpdateRequest) SetWebsites(v []string) {
	o.Websites = v
}

func (o WhatsappPhoneNumberProfileUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.About != nil {
		toSerialize["about"] = o.About
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.ProfilePictureUrl != nil {
		toSerialize["profilePictureUrl"] = o.ProfilePictureUrl
	}
	if o.Vertical != nil {
		toSerialize["vertical"] = o.Vertical
	}
	if o.Websites != nil {
		toSerialize["websites"] = o.Websites
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsappPhoneNumberProfileUpdateRequest struct {
	value *WhatsappPhoneNumberProfileUpdateRequest
	isSet bool
}

func (v NullableWhatsappPhoneNumberProfileUpdateRequest) Get() *WhatsappPhoneNumberProfileUpdateRequest {
	return v.value
}

func (v *NullableWhatsappPhoneNumberProfileUpdateRequest) Set(val *WhatsappPhoneNumberProfileUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsappPhoneNumberProfileUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsappPhoneNumberProfileUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsappPhoneNumberProfileUpdateRequest(val *WhatsappPhoneNumberProfileUpdateRequest) *NullableWhatsappPhoneNumberProfileUpdateRequest {
	return &NullableWhatsappPhoneNumberProfileUpdateRequest{value: val, isSet: true}
}

func (v NullableWhatsappPhoneNumberProfileUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsappPhoneNumberProfileUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


