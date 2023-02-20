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

// WhatsappApiError The original error object returned by WhatsApp. See [Handling Errors](https://developers.facebook.com/docs/graph-api/guides/error-handling), [Cloud API Error Codes](https://developers.facebook.com/docs/whatsapp/cloud-api/support/error-codes).
type WhatsappApiError struct {
	// A human-readable description of the error.
	Message string `json:"message"`
	// An error code.
	Code string `json:"code"`
	// Error type.
	Type *string `json:"type,omitempty"`
	// Additional code about the error.
	ErrorSubcode *string `json:"error_subcode,omitempty"`
	// The message to display to the user. The language of the message is based on the locale of the API request.
	ErrorUserMsg *string `json:"error_user_msg,omitempty"`
	// The title of the dialog, if shown. The language of the message is based on the locale of the API request.
	ErrorUserTitle *string `json:"error_user_title,omitempty"`
	// Internal support identifier. When reporting a bug related to a Graph API call, include the fbtrace_id to help us find log data for debugging.
	FbtraceId *string `json:"fbtrace_id,omitempty"`
	// Additional data about the error. A string or map. - For template APIs, this field is a string describing the reason for the error.   - For message APIs, this field is a map with property `details` describing the reason for the error.
	ErrorData map[string]interface{} `json:"error_data,omitempty"`
}

// NewWhatsappApiError instantiates a new WhatsappApiError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsappApiError(message string, code string) *WhatsappApiError {
	this := WhatsappApiError{}
	this.Message = message
	this.Code = code
	return &this
}

// NewWhatsappApiErrorWithDefaults instantiates a new WhatsappApiError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsappApiErrorWithDefaults() *WhatsappApiError {
	this := WhatsappApiError{}
	return &this
}

// GetMessage returns the Message field value
func (o *WhatsappApiError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *WhatsappApiError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *WhatsappApiError) SetMessage(v string) {
	o.Message = v
}

// GetCode returns the Code field value
func (o *WhatsappApiError) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *WhatsappApiError) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *WhatsappApiError) SetCode(v string) {
	o.Code = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WhatsappApiError) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappApiError) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WhatsappApiError) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WhatsappApiError) SetType(v string) {
	o.Type = &v
}

// GetErrorSubcode returns the ErrorSubcode field value if set, zero value otherwise.
func (o *WhatsappApiError) GetErrorSubcode() string {
	if o == nil || o.ErrorSubcode == nil {
		var ret string
		return ret
	}
	return *o.ErrorSubcode
}

// GetErrorSubcodeOk returns a tuple with the ErrorSubcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappApiError) GetErrorSubcodeOk() (*string, bool) {
	if o == nil || o.ErrorSubcode == nil {
		return nil, false
	}
	return o.ErrorSubcode, true
}

// HasErrorSubcode returns a boolean if a field has been set.
func (o *WhatsappApiError) HasErrorSubcode() bool {
	if o != nil && o.ErrorSubcode != nil {
		return true
	}

	return false
}

// SetErrorSubcode gets a reference to the given string and assigns it to the ErrorSubcode field.
func (o *WhatsappApiError) SetErrorSubcode(v string) {
	o.ErrorSubcode = &v
}

// GetErrorUserMsg returns the ErrorUserMsg field value if set, zero value otherwise.
func (o *WhatsappApiError) GetErrorUserMsg() string {
	if o == nil || o.ErrorUserMsg == nil {
		var ret string
		return ret
	}
	return *o.ErrorUserMsg
}

// GetErrorUserMsgOk returns a tuple with the ErrorUserMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappApiError) GetErrorUserMsgOk() (*string, bool) {
	if o == nil || o.ErrorUserMsg == nil {
		return nil, false
	}
	return o.ErrorUserMsg, true
}

// HasErrorUserMsg returns a boolean if a field has been set.
func (o *WhatsappApiError) HasErrorUserMsg() bool {
	if o != nil && o.ErrorUserMsg != nil {
		return true
	}

	return false
}

// SetErrorUserMsg gets a reference to the given string and assigns it to the ErrorUserMsg field.
func (o *WhatsappApiError) SetErrorUserMsg(v string) {
	o.ErrorUserMsg = &v
}

// GetErrorUserTitle returns the ErrorUserTitle field value if set, zero value otherwise.
func (o *WhatsappApiError) GetErrorUserTitle() string {
	if o == nil || o.ErrorUserTitle == nil {
		var ret string
		return ret
	}
	return *o.ErrorUserTitle
}

// GetErrorUserTitleOk returns a tuple with the ErrorUserTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappApiError) GetErrorUserTitleOk() (*string, bool) {
	if o == nil || o.ErrorUserTitle == nil {
		return nil, false
	}
	return o.ErrorUserTitle, true
}

// HasErrorUserTitle returns a boolean if a field has been set.
func (o *WhatsappApiError) HasErrorUserTitle() bool {
	if o != nil && o.ErrorUserTitle != nil {
		return true
	}

	return false
}

// SetErrorUserTitle gets a reference to the given string and assigns it to the ErrorUserTitle field.
func (o *WhatsappApiError) SetErrorUserTitle(v string) {
	o.ErrorUserTitle = &v
}

// GetFbtraceId returns the FbtraceId field value if set, zero value otherwise.
func (o *WhatsappApiError) GetFbtraceId() string {
	if o == nil || o.FbtraceId == nil {
		var ret string
		return ret
	}
	return *o.FbtraceId
}

// GetFbtraceIdOk returns a tuple with the FbtraceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappApiError) GetFbtraceIdOk() (*string, bool) {
	if o == nil || o.FbtraceId == nil {
		return nil, false
	}
	return o.FbtraceId, true
}

// HasFbtraceId returns a boolean if a field has been set.
func (o *WhatsappApiError) HasFbtraceId() bool {
	if o != nil && o.FbtraceId != nil {
		return true
	}

	return false
}

// SetFbtraceId gets a reference to the given string and assigns it to the FbtraceId field.
func (o *WhatsappApiError) SetFbtraceId(v string) {
	o.FbtraceId = &v
}

// GetErrorData returns the ErrorData field value if set, zero value otherwise.
func (o *WhatsappApiError) GetErrorData() map[string]interface{} {
	if o == nil || o.ErrorData == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ErrorData
}

// GetErrorDataOk returns a tuple with the ErrorData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappApiError) GetErrorDataOk() (map[string]interface{}, bool) {
	if o == nil || o.ErrorData == nil {
		return nil, false
	}
	return o.ErrorData, true
}

// HasErrorData returns a boolean if a field has been set.
func (o *WhatsappApiError) HasErrorData() bool {
	if o != nil && o.ErrorData != nil {
		return true
	}

	return false
}

// SetErrorData gets a reference to the given map[string]interface{} and assigns it to the ErrorData field.
func (o *WhatsappApiError) SetErrorData(v map[string]interface{}) {
	o.ErrorData = v
}

func (o WhatsappApiError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["code"] = o.Code
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.ErrorSubcode != nil {
		toSerialize["error_subcode"] = o.ErrorSubcode
	}
	if o.ErrorUserMsg != nil {
		toSerialize["error_user_msg"] = o.ErrorUserMsg
	}
	if o.ErrorUserTitle != nil {
		toSerialize["error_user_title"] = o.ErrorUserTitle
	}
	if o.FbtraceId != nil {
		toSerialize["fbtrace_id"] = o.FbtraceId
	}
	if o.ErrorData != nil {
		toSerialize["error_data"] = o.ErrorData
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsappApiError struct {
	value *WhatsappApiError
	isSet bool
}

func (v NullableWhatsappApiError) Get() *WhatsappApiError {
	return v.value
}

func (v *NullableWhatsappApiError) Set(val *WhatsappApiError) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsappApiError) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsappApiError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsappApiError(val *WhatsappApiError) *NullableWhatsappApiError {
	return &NullableWhatsappApiError{value: val, isSet: true}
}

func (v NullableWhatsappApiError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsappApiError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

