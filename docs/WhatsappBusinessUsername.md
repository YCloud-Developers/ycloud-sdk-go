# WhatsappBusinessUsername

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Phone number ID. | [optional] 
**WabaId** | Pointer to **string** | WhatsApp Business Account ID. | [optional] 
**PhoneNumber** | Pointer to **string** | Phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format. | [optional] 
**DisplayPhoneNumber** | Pointer to **string** | Display phone number. | [optional] 
**BusinessUsername** | Pointer to **string** | Active Business Username. The value is a plain username without &#x60;@&#x60;. | [optional] 
**BusinessUsernameStatus** | Pointer to [**WhatsappBusinessUsernameStatus**](WhatsappBusinessUsernameStatus.md) |  | [optional] 
**RequestedBusinessUsername** | Pointer to **string** | Last requested Business Username that is still under review. This value can coexist with an active &#x60;businessUsername&#x60; while the new request is pending. | [optional] 
**BusinessUsernameUpdatedAt** | Pointer to **time.Time** | The time when the Business Username state was last updated. | [optional] 

## Methods

### NewWhatsappBusinessUsername

`func NewWhatsappBusinessUsername() *WhatsappBusinessUsername`

NewWhatsappBusinessUsername instantiates a new WhatsappBusinessUsername object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappBusinessUsernameWithDefaults

`func NewWhatsappBusinessUsernameWithDefaults() *WhatsappBusinessUsername`

NewWhatsappBusinessUsernameWithDefaults instantiates a new WhatsappBusinessUsername object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WhatsappBusinessUsername) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WhatsappBusinessUsername) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WhatsappBusinessUsername) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WhatsappBusinessUsername) HasId() bool`

HasId returns a boolean if a field has been set.

### GetWabaId

`func (o *WhatsappBusinessUsername) GetWabaId() string`

GetWabaId returns the WabaId field if non-nil, zero value otherwise.

### GetWabaIdOk

`func (o *WhatsappBusinessUsername) GetWabaIdOk() (*string, bool)`

GetWabaIdOk returns a tuple with the WabaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWabaId

`func (o *WhatsappBusinessUsername) SetWabaId(v string)`

SetWabaId sets WabaId field to given value.

### HasWabaId

`func (o *WhatsappBusinessUsername) HasWabaId() bool`

HasWabaId returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *WhatsappBusinessUsername) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *WhatsappBusinessUsername) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *WhatsappBusinessUsername) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *WhatsappBusinessUsername) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetDisplayPhoneNumber

`func (o *WhatsappBusinessUsername) GetDisplayPhoneNumber() string`

GetDisplayPhoneNumber returns the DisplayPhoneNumber field if non-nil, zero value otherwise.

### GetDisplayPhoneNumberOk

`func (o *WhatsappBusinessUsername) GetDisplayPhoneNumberOk() (*string, bool)`

GetDisplayPhoneNumberOk returns a tuple with the DisplayPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayPhoneNumber

`func (o *WhatsappBusinessUsername) SetDisplayPhoneNumber(v string)`

SetDisplayPhoneNumber sets DisplayPhoneNumber field to given value.

### HasDisplayPhoneNumber

`func (o *WhatsappBusinessUsername) HasDisplayPhoneNumber() bool`

HasDisplayPhoneNumber returns a boolean if a field has been set.

### GetBusinessUsername

`func (o *WhatsappBusinessUsername) GetBusinessUsername() string`

GetBusinessUsername returns the BusinessUsername field if non-nil, zero value otherwise.

### GetBusinessUsernameOk

`func (o *WhatsappBusinessUsername) GetBusinessUsernameOk() (*string, bool)`

GetBusinessUsernameOk returns a tuple with the BusinessUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUsername

`func (o *WhatsappBusinessUsername) SetBusinessUsername(v string)`

SetBusinessUsername sets BusinessUsername field to given value.

### HasBusinessUsername

`func (o *WhatsappBusinessUsername) HasBusinessUsername() bool`

HasBusinessUsername returns a boolean if a field has been set.

### GetBusinessUsernameStatus

`func (o *WhatsappBusinessUsername) GetBusinessUsernameStatus() WhatsappBusinessUsernameStatus`

GetBusinessUsernameStatus returns the BusinessUsernameStatus field if non-nil, zero value otherwise.

### GetBusinessUsernameStatusOk

`func (o *WhatsappBusinessUsername) GetBusinessUsernameStatusOk() (*WhatsappBusinessUsernameStatus, bool)`

GetBusinessUsernameStatusOk returns a tuple with the BusinessUsernameStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUsernameStatus

`func (o *WhatsappBusinessUsername) SetBusinessUsernameStatus(v WhatsappBusinessUsernameStatus)`

SetBusinessUsernameStatus sets BusinessUsernameStatus field to given value.

### HasBusinessUsernameStatus

`func (o *WhatsappBusinessUsername) HasBusinessUsernameStatus() bool`

HasBusinessUsernameStatus returns a boolean if a field has been set.

### GetRequestedBusinessUsername

`func (o *WhatsappBusinessUsername) GetRequestedBusinessUsername() string`

GetRequestedBusinessUsername returns the RequestedBusinessUsername field if non-nil, zero value otherwise.

### GetRequestedBusinessUsernameOk

`func (o *WhatsappBusinessUsername) GetRequestedBusinessUsernameOk() (*string, bool)`

GetRequestedBusinessUsernameOk returns a tuple with the RequestedBusinessUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedBusinessUsername

`func (o *WhatsappBusinessUsername) SetRequestedBusinessUsername(v string)`

SetRequestedBusinessUsername sets RequestedBusinessUsername field to given value.

### HasRequestedBusinessUsername

`func (o *WhatsappBusinessUsername) HasRequestedBusinessUsername() bool`

HasRequestedBusinessUsername returns a boolean if a field has been set.

### GetBusinessUsernameUpdatedAt

`func (o *WhatsappBusinessUsername) GetBusinessUsernameUpdatedAt() time.Time`

GetBusinessUsernameUpdatedAt returns the BusinessUsernameUpdatedAt field if non-nil, zero value otherwise.

### GetBusinessUsernameUpdatedAtOk

`func (o *WhatsappBusinessUsername) GetBusinessUsernameUpdatedAtOk() (*time.Time, bool)`

GetBusinessUsernameUpdatedAtOk returns a tuple with the BusinessUsernameUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUsernameUpdatedAt

`func (o *WhatsappBusinessUsername) SetBusinessUsernameUpdatedAt(v time.Time)`

SetBusinessUsernameUpdatedAt sets BusinessUsernameUpdatedAt field to given value.

### HasBusinessUsernameUpdatedAt

`func (o *WhatsappBusinessUsername) HasBusinessUsernameUpdatedAt() bool`

HasBusinessUsernameUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


