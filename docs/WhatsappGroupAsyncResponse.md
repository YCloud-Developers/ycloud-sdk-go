# WhatsappGroupAsyncResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | Pointer to **string** | The request ID for tracking the asynchronous operation in webhooks. | [optional] 
**Status** | Pointer to **string** | The asynchronous request status. | [optional] 

## Methods

### NewWhatsappGroupAsyncResponse

`func NewWhatsappGroupAsyncResponse() *WhatsappGroupAsyncResponse`

NewWhatsappGroupAsyncResponse instantiates a new WhatsappGroupAsyncResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappGroupAsyncResponseWithDefaults

`func NewWhatsappGroupAsyncResponseWithDefaults() *WhatsappGroupAsyncResponse`

NewWhatsappGroupAsyncResponseWithDefaults instantiates a new WhatsappGroupAsyncResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *WhatsappGroupAsyncResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *WhatsappGroupAsyncResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *WhatsappGroupAsyncResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *WhatsappGroupAsyncResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetStatus

`func (o *WhatsappGroupAsyncResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WhatsappGroupAsyncResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WhatsappGroupAsyncResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WhatsappGroupAsyncResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


