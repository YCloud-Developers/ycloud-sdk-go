# WhatsappGroupJoinRequestActionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApprovedJoinRequests** | Pointer to **[]string** | Approved join request IDs. | [optional] 
**RejectedJoinRequests** | Pointer to **[]string** | Rejected join request IDs. | [optional] 
**FailedJoinRequests** | Pointer to [**[]WhatsappGroupFailedJoinRequest**](WhatsappGroupFailedJoinRequest.md) | Join requests that failed to be processed. | [optional] 
**Errors** | Pointer to **[]map[string]interface{}** | Errors returned by WhatsApp. | [optional] 

## Methods

### NewWhatsappGroupJoinRequestActionResponse

`func NewWhatsappGroupJoinRequestActionResponse() *WhatsappGroupJoinRequestActionResponse`

NewWhatsappGroupJoinRequestActionResponse instantiates a new WhatsappGroupJoinRequestActionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappGroupJoinRequestActionResponseWithDefaults

`func NewWhatsappGroupJoinRequestActionResponseWithDefaults() *WhatsappGroupJoinRequestActionResponse`

NewWhatsappGroupJoinRequestActionResponseWithDefaults instantiates a new WhatsappGroupJoinRequestActionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApprovedJoinRequests

`func (o *WhatsappGroupJoinRequestActionResponse) GetApprovedJoinRequests() []string`

GetApprovedJoinRequests returns the ApprovedJoinRequests field if non-nil, zero value otherwise.

### GetApprovedJoinRequestsOk

`func (o *WhatsappGroupJoinRequestActionResponse) GetApprovedJoinRequestsOk() (*[]string, bool)`

GetApprovedJoinRequestsOk returns a tuple with the ApprovedJoinRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedJoinRequests

`func (o *WhatsappGroupJoinRequestActionResponse) SetApprovedJoinRequests(v []string)`

SetApprovedJoinRequests sets ApprovedJoinRequests field to given value.

### HasApprovedJoinRequests

`func (o *WhatsappGroupJoinRequestActionResponse) HasApprovedJoinRequests() bool`

HasApprovedJoinRequests returns a boolean if a field has been set.

### GetRejectedJoinRequests

`func (o *WhatsappGroupJoinRequestActionResponse) GetRejectedJoinRequests() []string`

GetRejectedJoinRequests returns the RejectedJoinRequests field if non-nil, zero value otherwise.

### GetRejectedJoinRequestsOk

`func (o *WhatsappGroupJoinRequestActionResponse) GetRejectedJoinRequestsOk() (*[]string, bool)`

GetRejectedJoinRequestsOk returns a tuple with the RejectedJoinRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedJoinRequests

`func (o *WhatsappGroupJoinRequestActionResponse) SetRejectedJoinRequests(v []string)`

SetRejectedJoinRequests sets RejectedJoinRequests field to given value.

### HasRejectedJoinRequests

`func (o *WhatsappGroupJoinRequestActionResponse) HasRejectedJoinRequests() bool`

HasRejectedJoinRequests returns a boolean if a field has been set.

### GetFailedJoinRequests

`func (o *WhatsappGroupJoinRequestActionResponse) GetFailedJoinRequests() []WhatsappGroupFailedJoinRequest`

GetFailedJoinRequests returns the FailedJoinRequests field if non-nil, zero value otherwise.

### GetFailedJoinRequestsOk

`func (o *WhatsappGroupJoinRequestActionResponse) GetFailedJoinRequestsOk() (*[]WhatsappGroupFailedJoinRequest, bool)`

GetFailedJoinRequestsOk returns a tuple with the FailedJoinRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedJoinRequests

`func (o *WhatsappGroupJoinRequestActionResponse) SetFailedJoinRequests(v []WhatsappGroupFailedJoinRequest)`

SetFailedJoinRequests sets FailedJoinRequests field to given value.

### HasFailedJoinRequests

`func (o *WhatsappGroupJoinRequestActionResponse) HasFailedJoinRequests() bool`

HasFailedJoinRequests returns a boolean if a field has been set.

### GetErrors

`func (o *WhatsappGroupJoinRequestActionResponse) GetErrors() []map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *WhatsappGroupJoinRequestActionResponse) GetErrorsOk() (*[]map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *WhatsappGroupJoinRequestActionResponse) SetErrors(v []map[string]interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *WhatsappGroupJoinRequestActionResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


