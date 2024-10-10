# WhatsappMessageOrderStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferenceId** | Pointer to **string** | Unique identifier for the order provided by the business. | [optional] 
**Order** | Pointer to [**WhatsappMessageOrderInfo**](WhatsappMessageOrderInfo.md) |  | [optional] 

## Methods

### NewWhatsappMessageOrderStatus

`func NewWhatsappMessageOrderStatus() *WhatsappMessageOrderStatus`

NewWhatsappMessageOrderStatus instantiates a new WhatsappMessageOrderStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappMessageOrderStatusWithDefaults

`func NewWhatsappMessageOrderStatusWithDefaults() *WhatsappMessageOrderStatus`

NewWhatsappMessageOrderStatusWithDefaults instantiates a new WhatsappMessageOrderStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferenceId

`func (o *WhatsappMessageOrderStatus) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *WhatsappMessageOrderStatus) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *WhatsappMessageOrderStatus) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *WhatsappMessageOrderStatus) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetOrder

`func (o *WhatsappMessageOrderStatus) GetOrder() WhatsappMessageOrderInfo`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *WhatsappMessageOrderStatus) GetOrderOk() (*WhatsappMessageOrderInfo, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *WhatsappMessageOrderStatus) SetOrder(v WhatsappMessageOrderInfo)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *WhatsappMessageOrderStatus) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


