# WhatsappBusinessAccountPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]WhatsappBusinessAccount**](WhatsappBusinessAccount.md) | An array containing WhatsApp Business Account objects. | [optional] 
**Offset** | **int32** | The position of the item this page starts from, zero-based. e.g., the 11th item is at offset 10. | 
**Limit** | **int32** | A limit on the number of items to be returned, between 1 and 100, defaults to 10. | 
**Length** | **int32** | The actual number of items in the page. | 
**Total** | Pointer to **int32** | The total number of items. This field is returned only when the request parameter &#x60;includeTotal&#x60; is set to &#x60;true&#x60;. | [optional] 

## Methods

### NewWhatsappBusinessAccountPage

`func NewWhatsappBusinessAccountPage(offset int32, limit int32, length int32, ) *WhatsappBusinessAccountPage`

NewWhatsappBusinessAccountPage instantiates a new WhatsappBusinessAccountPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappBusinessAccountPageWithDefaults

`func NewWhatsappBusinessAccountPageWithDefaults() *WhatsappBusinessAccountPage`

NewWhatsappBusinessAccountPageWithDefaults instantiates a new WhatsappBusinessAccountPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *WhatsappBusinessAccountPage) GetItems() []WhatsappBusinessAccount`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *WhatsappBusinessAccountPage) GetItemsOk() (*[]WhatsappBusinessAccount, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *WhatsappBusinessAccountPage) SetItems(v []WhatsappBusinessAccount)`

SetItems sets Items field to given value.

### HasItems

`func (o *WhatsappBusinessAccountPage) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetOffset

`func (o *WhatsappBusinessAccountPage) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *WhatsappBusinessAccountPage) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *WhatsappBusinessAccountPage) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetLimit

`func (o *WhatsappBusinessAccountPage) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *WhatsappBusinessAccountPage) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *WhatsappBusinessAccountPage) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetLength

`func (o *WhatsappBusinessAccountPage) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *WhatsappBusinessAccountPage) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *WhatsappBusinessAccountPage) SetLength(v int32)`

SetLength sets Length field to given value.


### GetTotal

`func (o *WhatsappBusinessAccountPage) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *WhatsappBusinessAccountPage) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *WhatsappBusinessAccountPage) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *WhatsappBusinessAccountPage) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


