# WhatsappPhoneNumberPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]WhatsappPhoneNumber**](WhatsappPhoneNumber.md) | An array containing WhatsApp phone number objects. | [optional] 
**Offset** | **int32** | The position of the item this page starts from, zero-based. e.g., the 11th item is at offset 10. | 
**Limit** | **int32** | A limit on the number of items to be returned, between 1 and 100, defaults to 10. | 
**Length** | **int32** | The actual number of items in the page. | 
**Total** | Pointer to **int32** | The total number of items. This field is returned only when the request parameter &#x60;includeTotal&#x60; is set to &#x60;true&#x60;. | [optional] 

## Methods

### NewWhatsappPhoneNumberPage

`func NewWhatsappPhoneNumberPage(offset int32, limit int32, length int32, ) *WhatsappPhoneNumberPage`

NewWhatsappPhoneNumberPage instantiates a new WhatsappPhoneNumberPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappPhoneNumberPageWithDefaults

`func NewWhatsappPhoneNumberPageWithDefaults() *WhatsappPhoneNumberPage`

NewWhatsappPhoneNumberPageWithDefaults instantiates a new WhatsappPhoneNumberPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *WhatsappPhoneNumberPage) GetItems() []WhatsappPhoneNumber`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *WhatsappPhoneNumberPage) GetItemsOk() (*[]WhatsappPhoneNumber, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *WhatsappPhoneNumberPage) SetItems(v []WhatsappPhoneNumber)`

SetItems sets Items field to given value.

### HasItems

`func (o *WhatsappPhoneNumberPage) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetOffset

`func (o *WhatsappPhoneNumberPage) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *WhatsappPhoneNumberPage) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *WhatsappPhoneNumberPage) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetLimit

`func (o *WhatsappPhoneNumberPage) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *WhatsappPhoneNumberPage) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *WhatsappPhoneNumberPage) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetLength

`func (o *WhatsappPhoneNumberPage) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *WhatsappPhoneNumberPage) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *WhatsappPhoneNumberPage) SetLength(v int32)`

SetLength sets Length field to given value.


### GetTotal

`func (o *WhatsappPhoneNumberPage) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *WhatsappPhoneNumberPage) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *WhatsappPhoneNumberPage) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *WhatsappPhoneNumberPage) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


