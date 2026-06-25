# WhatsappBusinessUsernameDeleteResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** | Whether the delete request was accepted. | [optional] 
**Id** | Pointer to **string** | Phone number ID. | [optional] 
**WabaId** | Pointer to **string** | WhatsApp Business Account ID. | [optional] 
**PhoneNumber** | Pointer to **string** | Phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format. | [optional] 
**BusinessUsernameStatus** | Pointer to [**WhatsappBusinessUsernameStatus**](WhatsappBusinessUsernameStatus.md) |  | [optional] 
**BusinessUsernameUpdatedAt** | Pointer to **time.Time** | The time when the Business Username state was last updated. | [optional] 

## Methods

### NewWhatsappBusinessUsernameDeleteResult

`func NewWhatsappBusinessUsernameDeleteResult() *WhatsappBusinessUsernameDeleteResult`

NewWhatsappBusinessUsernameDeleteResult instantiates a new WhatsappBusinessUsernameDeleteResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappBusinessUsernameDeleteResultWithDefaults

`func NewWhatsappBusinessUsernameDeleteResultWithDefaults() *WhatsappBusinessUsernameDeleteResult`

NewWhatsappBusinessUsernameDeleteResultWithDefaults instantiates a new WhatsappBusinessUsernameDeleteResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *WhatsappBusinessUsernameDeleteResult) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *WhatsappBusinessUsernameDeleteResult) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *WhatsappBusinessUsernameDeleteResult) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *WhatsappBusinessUsernameDeleteResult) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetId

`func (o *WhatsappBusinessUsernameDeleteResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WhatsappBusinessUsernameDeleteResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WhatsappBusinessUsernameDeleteResult) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WhatsappBusinessUsernameDeleteResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetWabaId

`func (o *WhatsappBusinessUsernameDeleteResult) GetWabaId() string`

GetWabaId returns the WabaId field if non-nil, zero value otherwise.

### GetWabaIdOk

`func (o *WhatsappBusinessUsernameDeleteResult) GetWabaIdOk() (*string, bool)`

GetWabaIdOk returns a tuple with the WabaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWabaId

`func (o *WhatsappBusinessUsernameDeleteResult) SetWabaId(v string)`

SetWabaId sets WabaId field to given value.

### HasWabaId

`func (o *WhatsappBusinessUsernameDeleteResult) HasWabaId() bool`

HasWabaId returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *WhatsappBusinessUsernameDeleteResult) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *WhatsappBusinessUsernameDeleteResult) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *WhatsappBusinessUsernameDeleteResult) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *WhatsappBusinessUsernameDeleteResult) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetBusinessUsernameStatus

`func (o *WhatsappBusinessUsernameDeleteResult) GetBusinessUsernameStatus() WhatsappBusinessUsernameStatus`

GetBusinessUsernameStatus returns the BusinessUsernameStatus field if non-nil, zero value otherwise.

### GetBusinessUsernameStatusOk

`func (o *WhatsappBusinessUsernameDeleteResult) GetBusinessUsernameStatusOk() (*WhatsappBusinessUsernameStatus, bool)`

GetBusinessUsernameStatusOk returns a tuple with the BusinessUsernameStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUsernameStatus

`func (o *WhatsappBusinessUsernameDeleteResult) SetBusinessUsernameStatus(v WhatsappBusinessUsernameStatus)`

SetBusinessUsernameStatus sets BusinessUsernameStatus field to given value.

### HasBusinessUsernameStatus

`func (o *WhatsappBusinessUsernameDeleteResult) HasBusinessUsernameStatus() bool`

HasBusinessUsernameStatus returns a boolean if a field has been set.

### GetBusinessUsernameUpdatedAt

`func (o *WhatsappBusinessUsernameDeleteResult) GetBusinessUsernameUpdatedAt() time.Time`

GetBusinessUsernameUpdatedAt returns the BusinessUsernameUpdatedAt field if non-nil, zero value otherwise.

### GetBusinessUsernameUpdatedAtOk

`func (o *WhatsappBusinessUsernameDeleteResult) GetBusinessUsernameUpdatedAtOk() (*time.Time, bool)`

GetBusinessUsernameUpdatedAtOk returns a tuple with the BusinessUsernameUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUsernameUpdatedAt

`func (o *WhatsappBusinessUsernameDeleteResult) SetBusinessUsernameUpdatedAt(v time.Time)`

SetBusinessUsernameUpdatedAt sets BusinessUsernameUpdatedAt field to given value.

### HasBusinessUsernameUpdatedAt

`func (o *WhatsappBusinessUsernameDeleteResult) HasBusinessUsernameUpdatedAt() bool`

HasBusinessUsernameUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


