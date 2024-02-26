# WhatsappCommerceSettingsUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsCartEnabled** | Pointer to **bool** | When enabled, cart-related buttons appear in the conversation, catalog, and product details views. When the cart is disabled, customers can see products and their details, but all cart related buttons will not appear in any view. | [optional] 
**IsCatalogVisible** | Pointer to **bool** | When enabled, the catalog storefront icon and catalog-related buttons appear in conversation and business profile views. When the catalog is disabled, the storefront icon and catalog-related buttons will not appear in any views and the catalog preview with thumbnails will not appear in the business profile view. | [optional] 

## Methods

### NewWhatsappCommerceSettingsUpdateRequest

`func NewWhatsappCommerceSettingsUpdateRequest() *WhatsappCommerceSettingsUpdateRequest`

NewWhatsappCommerceSettingsUpdateRequest instantiates a new WhatsappCommerceSettingsUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappCommerceSettingsUpdateRequestWithDefaults

`func NewWhatsappCommerceSettingsUpdateRequestWithDefaults() *WhatsappCommerceSettingsUpdateRequest`

NewWhatsappCommerceSettingsUpdateRequestWithDefaults instantiates a new WhatsappCommerceSettingsUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsCartEnabled

`func (o *WhatsappCommerceSettingsUpdateRequest) GetIsCartEnabled() bool`

GetIsCartEnabled returns the IsCartEnabled field if non-nil, zero value otherwise.

### GetIsCartEnabledOk

`func (o *WhatsappCommerceSettingsUpdateRequest) GetIsCartEnabledOk() (*bool, bool)`

GetIsCartEnabledOk returns a tuple with the IsCartEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCartEnabled

`func (o *WhatsappCommerceSettingsUpdateRequest) SetIsCartEnabled(v bool)`

SetIsCartEnabled sets IsCartEnabled field to given value.

### HasIsCartEnabled

`func (o *WhatsappCommerceSettingsUpdateRequest) HasIsCartEnabled() bool`

HasIsCartEnabled returns a boolean if a field has been set.

### GetIsCatalogVisible

`func (o *WhatsappCommerceSettingsUpdateRequest) GetIsCatalogVisible() bool`

GetIsCatalogVisible returns the IsCatalogVisible field if non-nil, zero value otherwise.

### GetIsCatalogVisibleOk

`func (o *WhatsappCommerceSettingsUpdateRequest) GetIsCatalogVisibleOk() (*bool, bool)`

GetIsCatalogVisibleOk returns a tuple with the IsCatalogVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCatalogVisible

`func (o *WhatsappCommerceSettingsUpdateRequest) SetIsCatalogVisible(v bool)`

SetIsCatalogVisible sets IsCatalogVisible field to given value.

### HasIsCatalogVisible

`func (o *WhatsappCommerceSettingsUpdateRequest) HasIsCatalogVisible() bool`

HasIsCatalogVisible returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


