# WhatsappInboundMessageOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CatalogId** | Pointer to **string** | The catalog ID. | [optional] 
**ProductItems** | Pointer to [**[]WhatsappInboundMessageOrderProductItem**](WhatsappInboundMessageOrderProductItem.md) |  | [optional] 
**Text** | Pointer to **string** | Text message sent along with the order. | [optional] 

## Methods

### NewWhatsappInboundMessageOrder

`func NewWhatsappInboundMessageOrder() *WhatsappInboundMessageOrder`

NewWhatsappInboundMessageOrder instantiates a new WhatsappInboundMessageOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappInboundMessageOrderWithDefaults

`func NewWhatsappInboundMessageOrderWithDefaults() *WhatsappInboundMessageOrder`

NewWhatsappInboundMessageOrderWithDefaults instantiates a new WhatsappInboundMessageOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalogId

`func (o *WhatsappInboundMessageOrder) GetCatalogId() string`

GetCatalogId returns the CatalogId field if non-nil, zero value otherwise.

### GetCatalogIdOk

`func (o *WhatsappInboundMessageOrder) GetCatalogIdOk() (*string, bool)`

GetCatalogIdOk returns a tuple with the CatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogId

`func (o *WhatsappInboundMessageOrder) SetCatalogId(v string)`

SetCatalogId sets CatalogId field to given value.

### HasCatalogId

`func (o *WhatsappInboundMessageOrder) HasCatalogId() bool`

HasCatalogId returns a boolean if a field has been set.

### GetProductItems

`func (o *WhatsappInboundMessageOrder) GetProductItems() []WhatsappInboundMessageOrderProductItem`

GetProductItems returns the ProductItems field if non-nil, zero value otherwise.

### GetProductItemsOk

`func (o *WhatsappInboundMessageOrder) GetProductItemsOk() (*[]WhatsappInboundMessageOrderProductItem, bool)`

GetProductItemsOk returns a tuple with the ProductItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductItems

`func (o *WhatsappInboundMessageOrder) SetProductItems(v []WhatsappInboundMessageOrderProductItem)`

SetProductItems sets ProductItems field to given value.

### HasProductItems

`func (o *WhatsappInboundMessageOrder) HasProductItems() bool`

HasProductItems returns a boolean if a field has been set.

### GetText

`func (o *WhatsappInboundMessageOrder) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *WhatsappInboundMessageOrder) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *WhatsappInboundMessageOrder) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *WhatsappInboundMessageOrder) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


