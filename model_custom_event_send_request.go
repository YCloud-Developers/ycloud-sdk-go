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

// CustomEventSendRequest Contains the properties of the custom event data to be sent.
type CustomEventSendRequest struct {
	// Name of the event. One of the custom event names you previously defined.
	EventName string `json:"eventName"`
	// The time at which the event occurred, formatted in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339). e.g., `2022-06-01T12:00:00.000Z`, if not provided, the current time will be used.
	OccurTime *time.Time `json:"occurTime,omitempty"`
	// ID of the object that the event is associated with. For events defined with `objectType` as `CONTACT`, the `objectId` should be a `contact` ID. Alternatively, you can use the `contactPhoneNumber` field to specify the contact.
	ObjectId *string `json:"objectId,omitempty"`
	// The phone number of the contact for events defined with `objectType` as `CONTACT`.
	ContactPhoneNumber *string `json:"contactPhoneNumber,omitempty"`
	// The properties of the custom event.
	Properties map[string]map[string]interface{} `json:"properties,omitempty"`
}

// NewCustomEventSendRequest instantiates a new CustomEventSendRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomEventSendRequest(eventName string) *CustomEventSendRequest {
	this := CustomEventSendRequest{}
	this.EventName = eventName
	return &this
}

// NewCustomEventSendRequestWithDefaults instantiates a new CustomEventSendRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomEventSendRequestWithDefaults() *CustomEventSendRequest {
	this := CustomEventSendRequest{}
	return &this
}

// GetEventName returns the EventName field value
func (o *CustomEventSendRequest) GetEventName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventName
}

// GetEventNameOk returns a tuple with the EventName field value
// and a boolean to check if the value has been set.
func (o *CustomEventSendRequest) GetEventNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventName, true
}

// SetEventName sets field value
func (o *CustomEventSendRequest) SetEventName(v string) {
	o.EventName = v
}

// GetOccurTime returns the OccurTime field value if set, zero value otherwise.
func (o *CustomEventSendRequest) GetOccurTime() time.Time {
	if o == nil || o.OccurTime == nil {
		var ret time.Time
		return ret
	}
	return *o.OccurTime
}

// GetOccurTimeOk returns a tuple with the OccurTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEventSendRequest) GetOccurTimeOk() (*time.Time, bool) {
	if o == nil || o.OccurTime == nil {
		return nil, false
	}
	return o.OccurTime, true
}

// HasOccurTime returns a boolean if a field has been set.
func (o *CustomEventSendRequest) HasOccurTime() bool {
	if o != nil && o.OccurTime != nil {
		return true
	}

	return false
}

// SetOccurTime gets a reference to the given time.Time and assigns it to the OccurTime field.
func (o *CustomEventSendRequest) SetOccurTime(v time.Time) {
	o.OccurTime = &v
}

// GetObjectId returns the ObjectId field value if set, zero value otherwise.
func (o *CustomEventSendRequest) GetObjectId() string {
	if o == nil || o.ObjectId == nil {
		var ret string
		return ret
	}
	return *o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEventSendRequest) GetObjectIdOk() (*string, bool) {
	if o == nil || o.ObjectId == nil {
		return nil, false
	}
	return o.ObjectId, true
}

// HasObjectId returns a boolean if a field has been set.
func (o *CustomEventSendRequest) HasObjectId() bool {
	if o != nil && o.ObjectId != nil {
		return true
	}

	return false
}

// SetObjectId gets a reference to the given string and assigns it to the ObjectId field.
func (o *CustomEventSendRequest) SetObjectId(v string) {
	o.ObjectId = &v
}

// GetContactPhoneNumber returns the ContactPhoneNumber field value if set, zero value otherwise.
func (o *CustomEventSendRequest) GetContactPhoneNumber() string {
	if o == nil || o.ContactPhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.ContactPhoneNumber
}

// GetContactPhoneNumberOk returns a tuple with the ContactPhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEventSendRequest) GetContactPhoneNumberOk() (*string, bool) {
	if o == nil || o.ContactPhoneNumber == nil {
		return nil, false
	}
	return o.ContactPhoneNumber, true
}

// HasContactPhoneNumber returns a boolean if a field has been set.
func (o *CustomEventSendRequest) HasContactPhoneNumber() bool {
	if o != nil && o.ContactPhoneNumber != nil {
		return true
	}

	return false
}

// SetContactPhoneNumber gets a reference to the given string and assigns it to the ContactPhoneNumber field.
func (o *CustomEventSendRequest) SetContactPhoneNumber(v string) {
	o.ContactPhoneNumber = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *CustomEventSendRequest) GetProperties() map[string]map[string]interface{} {
	if o == nil || o.Properties == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEventSendRequest) GetPropertiesOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *CustomEventSendRequest) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]map[string]interface{} and assigns it to the Properties field.
func (o *CustomEventSendRequest) SetProperties(v map[string]map[string]interface{}) {
	o.Properties = v
}

func (o CustomEventSendRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["eventName"] = o.EventName
	}
	if o.OccurTime != nil {
		toSerialize["occurTime"] = o.OccurTime
	}
	if o.ObjectId != nil {
		toSerialize["objectId"] = o.ObjectId
	}
	if o.ContactPhoneNumber != nil {
		toSerialize["contactPhoneNumber"] = o.ContactPhoneNumber
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	return json.Marshal(toSerialize)
}

type NullableCustomEventSendRequest struct {
	value *CustomEventSendRequest
	isSet bool
}

func (v NullableCustomEventSendRequest) Get() *CustomEventSendRequest {
	return v.value
}

func (v *NullableCustomEventSendRequest) Set(val *CustomEventSendRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomEventSendRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomEventSendRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomEventSendRequest(val *CustomEventSendRequest) *NullableCustomEventSendRequest {
	return &NullableCustomEventSendRequest{value: val, isSet: true}
}

func (v NullableCustomEventSendRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomEventSendRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


