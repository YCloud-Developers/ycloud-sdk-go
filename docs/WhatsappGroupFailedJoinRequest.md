# WhatsappGroupFailedJoinRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JoinRequestId** | Pointer to **string** | The join request ID. | [optional] 
**Errors** | Pointer to **[]map[string]interface{}** | Errors returned by WhatsApp for this join request. | [optional] 

## Methods

### NewWhatsappGroupFailedJoinRequest

`func NewWhatsappGroupFailedJoinRequest() *WhatsappGroupFailedJoinRequest`

NewWhatsappGroupFailedJoinRequest instantiates a new WhatsappGroupFailedJoinRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappGroupFailedJoinRequestWithDefaults

`func NewWhatsappGroupFailedJoinRequestWithDefaults() *WhatsappGroupFailedJoinRequest`

NewWhatsappGroupFailedJoinRequestWithDefaults instantiates a new WhatsappGroupFailedJoinRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJoinRequestId

`func (o *WhatsappGroupFailedJoinRequest) GetJoinRequestId() string`

GetJoinRequestId returns the JoinRequestId field if non-nil, zero value otherwise.

### GetJoinRequestIdOk

`func (o *WhatsappGroupFailedJoinRequest) GetJoinRequestIdOk() (*string, bool)`

GetJoinRequestIdOk returns a tuple with the JoinRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinRequestId

`func (o *WhatsappGroupFailedJoinRequest) SetJoinRequestId(v string)`

SetJoinRequestId sets JoinRequestId field to given value.

### HasJoinRequestId

`func (o *WhatsappGroupFailedJoinRequest) HasJoinRequestId() bool`

HasJoinRequestId returns a boolean if a field has been set.

### GetErrors

`func (o *WhatsappGroupFailedJoinRequest) GetErrors() []map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *WhatsappGroupFailedJoinRequest) GetErrorsOk() (*[]map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *WhatsappGroupFailedJoinRequest) SetErrors(v []map[string]interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *WhatsappGroupFailedJoinRequest) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


