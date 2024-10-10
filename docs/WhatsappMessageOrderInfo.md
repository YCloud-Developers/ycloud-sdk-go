# WhatsappMessageOrderInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**WhatsappMessageOrderStatusEnum**](WhatsappMessageOrderStatusEnum.md) |  | [optional] 
**Type** | Pointer to **string** | Only supported value is &#x60;quick_pay&#x60;. When this field is passed in we hide the \&quot;Review and Pay\&quot; button and only show the \&quot;Pay Now\&quot; button in the order details bubble. | [optional] 
**CatalogId** | Pointer to **string** | Unique identifier of the Facebook catalog being used by the business. If you do not provide this field, you must provide the following fields inside the items object: &#x60;country_of_origin&#x60;, &#x60;importer_name&#x60;, and &#x60;importer_address&#x60;. | [optional] 
**Items** | Pointer to [**[]WhatsappMessageOrderItem**](WhatsappMessageOrderItem.md) | Array of items in the order. | [optional] 
**Subtotal** | Pointer to [**WhatsappMessageOrderAmount**](WhatsappMessageOrderAmount.md) |  | [optional] 
**Tax** | Pointer to [**WhatsappMessageOrderAmount**](WhatsappMessageOrderAmount.md) |  | [optional] 
**Shipping** | Pointer to [**WhatsappMessageOrderAmount**](WhatsappMessageOrderAmount.md) |  | [optional] 
**Discount** | Pointer to [**WhatsappMessageOrderAmount**](WhatsappMessageOrderAmount.md) |  | [optional] 
**Expiration** | Pointer to [**WhatsappMessageOrderExpiration**](WhatsappMessageOrderExpiration.md) |  | [optional] 
**Description** | Pointer to **string** | **Optional.** Text for sharing status related information. Could be useful while sending cancellation. Max character limit is 120 characters. | [optional] 

## Methods

### NewWhatsappMessageOrderInfo

`func NewWhatsappMessageOrderInfo() *WhatsappMessageOrderInfo`

NewWhatsappMessageOrderInfo instantiates a new WhatsappMessageOrderInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappMessageOrderInfoWithDefaults

`func NewWhatsappMessageOrderInfoWithDefaults() *WhatsappMessageOrderInfo`

NewWhatsappMessageOrderInfoWithDefaults instantiates a new WhatsappMessageOrderInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *WhatsappMessageOrderInfo) GetStatus() WhatsappMessageOrderStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WhatsappMessageOrderInfo) GetStatusOk() (*WhatsappMessageOrderStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WhatsappMessageOrderInfo) SetStatus(v WhatsappMessageOrderStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WhatsappMessageOrderInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *WhatsappMessageOrderInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WhatsappMessageOrderInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WhatsappMessageOrderInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WhatsappMessageOrderInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCatalogId

`func (o *WhatsappMessageOrderInfo) GetCatalogId() string`

GetCatalogId returns the CatalogId field if non-nil, zero value otherwise.

### GetCatalogIdOk

`func (o *WhatsappMessageOrderInfo) GetCatalogIdOk() (*string, bool)`

GetCatalogIdOk returns a tuple with the CatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogId

`func (o *WhatsappMessageOrderInfo) SetCatalogId(v string)`

SetCatalogId sets CatalogId field to given value.

### HasCatalogId

`func (o *WhatsappMessageOrderInfo) HasCatalogId() bool`

HasCatalogId returns a boolean if a field has been set.

### GetItems

`func (o *WhatsappMessageOrderInfo) GetItems() []WhatsappMessageOrderItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *WhatsappMessageOrderInfo) GetItemsOk() (*[]WhatsappMessageOrderItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *WhatsappMessageOrderInfo) SetItems(v []WhatsappMessageOrderItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *WhatsappMessageOrderInfo) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetSubtotal

`func (o *WhatsappMessageOrderInfo) GetSubtotal() WhatsappMessageOrderAmount`

GetSubtotal returns the Subtotal field if non-nil, zero value otherwise.

### GetSubtotalOk

`func (o *WhatsappMessageOrderInfo) GetSubtotalOk() (*WhatsappMessageOrderAmount, bool)`

GetSubtotalOk returns a tuple with the Subtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotal

`func (o *WhatsappMessageOrderInfo) SetSubtotal(v WhatsappMessageOrderAmount)`

SetSubtotal sets Subtotal field to given value.

### HasSubtotal

`func (o *WhatsappMessageOrderInfo) HasSubtotal() bool`

HasSubtotal returns a boolean if a field has been set.

### GetTax

`func (o *WhatsappMessageOrderInfo) GetTax() WhatsappMessageOrderAmount`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *WhatsappMessageOrderInfo) GetTaxOk() (*WhatsappMessageOrderAmount, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *WhatsappMessageOrderInfo) SetTax(v WhatsappMessageOrderAmount)`

SetTax sets Tax field to given value.

### HasTax

`func (o *WhatsappMessageOrderInfo) HasTax() bool`

HasTax returns a boolean if a field has been set.

### GetShipping

`func (o *WhatsappMessageOrderInfo) GetShipping() WhatsappMessageOrderAmount`

GetShipping returns the Shipping field if non-nil, zero value otherwise.

### GetShippingOk

`func (o *WhatsappMessageOrderInfo) GetShippingOk() (*WhatsappMessageOrderAmount, bool)`

GetShippingOk returns a tuple with the Shipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipping

`func (o *WhatsappMessageOrderInfo) SetShipping(v WhatsappMessageOrderAmount)`

SetShipping sets Shipping field to given value.

### HasShipping

`func (o *WhatsappMessageOrderInfo) HasShipping() bool`

HasShipping returns a boolean if a field has been set.

### GetDiscount

`func (o *WhatsappMessageOrderInfo) GetDiscount() WhatsappMessageOrderAmount`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *WhatsappMessageOrderInfo) GetDiscountOk() (*WhatsappMessageOrderAmount, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *WhatsappMessageOrderInfo) SetDiscount(v WhatsappMessageOrderAmount)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *WhatsappMessageOrderInfo) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetExpiration

`func (o *WhatsappMessageOrderInfo) GetExpiration() WhatsappMessageOrderExpiration`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *WhatsappMessageOrderInfo) GetExpirationOk() (*WhatsappMessageOrderExpiration, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *WhatsappMessageOrderInfo) SetExpiration(v WhatsappMessageOrderExpiration)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *WhatsappMessageOrderInfo) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetDescription

`func (o *WhatsappMessageOrderInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WhatsappMessageOrderInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WhatsappMessageOrderInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WhatsappMessageOrderInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


