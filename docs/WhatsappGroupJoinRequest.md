# WhatsappGroupJoinRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JoinRequestId** | Pointer to **string** | The join request ID. | [optional] 
**WaId** | Pointer to **string** | WhatsApp user ID. | [optional] 
**UserId** | Pointer to **string** | Business-scoped user ID. | [optional] 
**ParentUserId** | Pointer to **string** | Parent business-scoped user ID. | [optional] 
**Username** | Pointer to **string** | WhatsApp username. | [optional] 
**CreationTimestamp** | Pointer to **int64** | Unix timestamp indicating when the join request was created. | [optional] 

## Methods

### NewWhatsappGroupJoinRequest

`func NewWhatsappGroupJoinRequest() *WhatsappGroupJoinRequest`

NewWhatsappGroupJoinRequest instantiates a new WhatsappGroupJoinRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappGroupJoinRequestWithDefaults

`func NewWhatsappGroupJoinRequestWithDefaults() *WhatsappGroupJoinRequest`

NewWhatsappGroupJoinRequestWithDefaults instantiates a new WhatsappGroupJoinRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJoinRequestId

`func (o *WhatsappGroupJoinRequest) GetJoinRequestId() string`

GetJoinRequestId returns the JoinRequestId field if non-nil, zero value otherwise.

### GetJoinRequestIdOk

`func (o *WhatsappGroupJoinRequest) GetJoinRequestIdOk() (*string, bool)`

GetJoinRequestIdOk returns a tuple with the JoinRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinRequestId

`func (o *WhatsappGroupJoinRequest) SetJoinRequestId(v string)`

SetJoinRequestId sets JoinRequestId field to given value.

### HasJoinRequestId

`func (o *WhatsappGroupJoinRequest) HasJoinRequestId() bool`

HasJoinRequestId returns a boolean if a field has been set.

### GetWaId

`func (o *WhatsappGroupJoinRequest) GetWaId() string`

GetWaId returns the WaId field if non-nil, zero value otherwise.

### GetWaIdOk

`func (o *WhatsappGroupJoinRequest) GetWaIdOk() (*string, bool)`

GetWaIdOk returns a tuple with the WaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaId

`func (o *WhatsappGroupJoinRequest) SetWaId(v string)`

SetWaId sets WaId field to given value.

### HasWaId

`func (o *WhatsappGroupJoinRequest) HasWaId() bool`

HasWaId returns a boolean if a field has been set.

### GetUserId

`func (o *WhatsappGroupJoinRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *WhatsappGroupJoinRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *WhatsappGroupJoinRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *WhatsappGroupJoinRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetParentUserId

`func (o *WhatsappGroupJoinRequest) GetParentUserId() string`

GetParentUserId returns the ParentUserId field if non-nil, zero value otherwise.

### GetParentUserIdOk

`func (o *WhatsappGroupJoinRequest) GetParentUserIdOk() (*string, bool)`

GetParentUserIdOk returns a tuple with the ParentUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentUserId

`func (o *WhatsappGroupJoinRequest) SetParentUserId(v string)`

SetParentUserId sets ParentUserId field to given value.

### HasParentUserId

`func (o *WhatsappGroupJoinRequest) HasParentUserId() bool`

HasParentUserId returns a boolean if a field has been set.

### GetUsername

`func (o *WhatsappGroupJoinRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *WhatsappGroupJoinRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *WhatsappGroupJoinRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *WhatsappGroupJoinRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetCreationTimestamp

`func (o *WhatsappGroupJoinRequest) GetCreationTimestamp() int64`

GetCreationTimestamp returns the CreationTimestamp field if non-nil, zero value otherwise.

### GetCreationTimestampOk

`func (o *WhatsappGroupJoinRequest) GetCreationTimestampOk() (*int64, bool)`

GetCreationTimestampOk returns a tuple with the CreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimestamp

`func (o *WhatsappGroupJoinRequest) SetCreationTimestamp(v int64)`

SetCreationTimestamp sets CreationTimestamp field to given value.

### HasCreationTimestamp

`func (o *WhatsappGroupJoinRequest) HasCreationTimestamp() bool`

HasCreationTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


