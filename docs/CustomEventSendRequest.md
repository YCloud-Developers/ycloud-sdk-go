# CustomEventSendRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventName** | **string** | Name of the event. One of the custom event names you previously defined. | 
**OccurTime** | Pointer to **time.Time** | The time at which the event occurred, formatted in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339). e.g., &#x60;2022-06-01T12:00:00.000Z&#x60;, if not provided, the current time will be used. | [optional] 
**ObjectId** | Pointer to **string** | ID of the object that the event is associated with. For events defined with &#x60;objectType&#x60; as &#x60;CONTACT&#x60;, the &#x60;objectId&#x60; should be a &#x60;contact&#x60; ID. Alternatively, you can use the &#x60;contactPhoneNumber&#x60; field to specify the contact. | [optional] 
**ContactPhoneNumber** | Pointer to **string** | The phone number of the contact for events defined with &#x60;objectType&#x60; as &#x60;CONTACT&#x60;. | [optional] 
**Properties** | Pointer to **map[string]map[string]interface{}** | The properties of the custom event. | [optional] 

## Methods

### NewCustomEventSendRequest

`func NewCustomEventSendRequest(eventName string, ) *CustomEventSendRequest`

NewCustomEventSendRequest instantiates a new CustomEventSendRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomEventSendRequestWithDefaults

`func NewCustomEventSendRequestWithDefaults() *CustomEventSendRequest`

NewCustomEventSendRequestWithDefaults instantiates a new CustomEventSendRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventName

`func (o *CustomEventSendRequest) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *CustomEventSendRequest) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *CustomEventSendRequest) SetEventName(v string)`

SetEventName sets EventName field to given value.


### GetOccurTime

`func (o *CustomEventSendRequest) GetOccurTime() time.Time`

GetOccurTime returns the OccurTime field if non-nil, zero value otherwise.

### GetOccurTimeOk

`func (o *CustomEventSendRequest) GetOccurTimeOk() (*time.Time, bool)`

GetOccurTimeOk returns a tuple with the OccurTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurTime

`func (o *CustomEventSendRequest) SetOccurTime(v time.Time)`

SetOccurTime sets OccurTime field to given value.

### HasOccurTime

`func (o *CustomEventSendRequest) HasOccurTime() bool`

HasOccurTime returns a boolean if a field has been set.

### GetObjectId

`func (o *CustomEventSendRequest) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *CustomEventSendRequest) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *CustomEventSendRequest) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *CustomEventSendRequest) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetContactPhoneNumber

`func (o *CustomEventSendRequest) GetContactPhoneNumber() string`

GetContactPhoneNumber returns the ContactPhoneNumber field if non-nil, zero value otherwise.

### GetContactPhoneNumberOk

`func (o *CustomEventSendRequest) GetContactPhoneNumberOk() (*string, bool)`

GetContactPhoneNumberOk returns a tuple with the ContactPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPhoneNumber

`func (o *CustomEventSendRequest) SetContactPhoneNumber(v string)`

SetContactPhoneNumber sets ContactPhoneNumber field to given value.

### HasContactPhoneNumber

`func (o *CustomEventSendRequest) HasContactPhoneNumber() bool`

HasContactPhoneNumber returns a boolean if a field has been set.

### GetProperties

`func (o *CustomEventSendRequest) GetProperties() map[string]map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CustomEventSendRequest) GetPropertiesOk() (*map[string]map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CustomEventSendRequest) SetProperties(v map[string]map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CustomEventSendRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


