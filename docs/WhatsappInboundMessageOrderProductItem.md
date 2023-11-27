# WhatsappInboundMessageOrderProductItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductRetailerId** | Pointer to **string** | The product SKU identifier. | [optional] 
**Quantity** | Pointer to **int32** | Number of item. | [optional] 
**ItemPrice** | Pointer to **float64** | Unitary price of item. | [optional] 
**Currency** | Pointer to **string** | Price currency. [ISO 4217 currency code](https://en.wikipedia.org/wiki/ISO_4217). | [optional] 

## Methods

### NewWhatsappInboundMessageOrderProductItem

`func NewWhatsappInboundMessageOrderProductItem() *WhatsappInboundMessageOrderProductItem`

NewWhatsappInboundMessageOrderProductItem instantiates a new WhatsappInboundMessageOrderProductItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappInboundMessageOrderProductItemWithDefaults

`func NewWhatsappInboundMessageOrderProductItemWithDefaults() *WhatsappInboundMessageOrderProductItem`

NewWhatsappInboundMessageOrderProductItemWithDefaults instantiates a new WhatsappInboundMessageOrderProductItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductRetailerId

`func (o *WhatsappInboundMessageOrderProductItem) GetProductRetailerId() string`

GetProductRetailerId returns the ProductRetailerId field if non-nil, zero value otherwise.

### GetProductRetailerIdOk

`func (o *WhatsappInboundMessageOrderProductItem) GetProductRetailerIdOk() (*string, bool)`

GetProductRetailerIdOk returns a tuple with the ProductRetailerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductRetailerId

`func (o *WhatsappInboundMessageOrderProductItem) SetProductRetailerId(v string)`

SetProductRetailerId sets ProductRetailerId field to given value.

### HasProductRetailerId

`func (o *WhatsappInboundMessageOrderProductItem) HasProductRetailerId() bool`

HasProductRetailerId returns a boolean if a field has been set.

### GetQuantity

`func (o *WhatsappInboundMessageOrderProductItem) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *WhatsappInboundMessageOrderProductItem) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *WhatsappInboundMessageOrderProductItem) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *WhatsappInboundMessageOrderProductItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetItemPrice

`func (o *WhatsappInboundMessageOrderProductItem) GetItemPrice() float64`

GetItemPrice returns the ItemPrice field if non-nil, zero value otherwise.

### GetItemPriceOk

`func (o *WhatsappInboundMessageOrderProductItem) GetItemPriceOk() (*float64, bool)`

GetItemPriceOk returns a tuple with the ItemPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemPrice

`func (o *WhatsappInboundMessageOrderProductItem) SetItemPrice(v float64)`

SetItemPrice sets ItemPrice field to given value.

### HasItemPrice

`func (o *WhatsappInboundMessageOrderProductItem) HasItemPrice() bool`

HasItemPrice returns a boolean if a field has been set.

### GetCurrency

`func (o *WhatsappInboundMessageOrderProductItem) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *WhatsappInboundMessageOrderProductItem) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *WhatsappInboundMessageOrderProductItem) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *WhatsappInboundMessageOrderProductItem) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


