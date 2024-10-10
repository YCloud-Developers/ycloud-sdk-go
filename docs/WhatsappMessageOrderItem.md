# WhatsappMessageOrderItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RetailerId** | Pointer to **string** | Content ID for an item in the order from your catalog. | [optional] 
**Name** | **string** | The item&#39;s name to be displayed to the user. Cannot exceed 60 characters. | 
**Image** | Pointer to [**WhatsappMessageMedia**](WhatsappMessageMedia.md) |  | [optional] 
**Amount** | [**WhatsappMessageOrderAmount**](WhatsappMessageOrderAmount.md) |  | 
**SaleAmount** | Pointer to [**WhatsappMessageOrderAmount**](WhatsappMessageOrderAmount.md) |  | [optional] 
**Quantity** | **int32** | The number of items in the order. | 
**CountryOfOrigin** | Pointer to **string** | Required if &#x60;catalog_id&#x60; is not present. The country of origin of the product. | [optional] 
**ImporterName** | Pointer to **string** | Required if &#x60;catalog_id&#x60; is not present. Name of the importer company. | [optional] 
**ImporterAddress** | Pointer to **string** | Required if &#x60;catalog_id&#x60; is not present. Address of importer company. | [optional] 

## Methods

### NewWhatsappMessageOrderItem

`func NewWhatsappMessageOrderItem(name string, amount WhatsappMessageOrderAmount, quantity int32, ) *WhatsappMessageOrderItem`

NewWhatsappMessageOrderItem instantiates a new WhatsappMessageOrderItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappMessageOrderItemWithDefaults

`func NewWhatsappMessageOrderItemWithDefaults() *WhatsappMessageOrderItem`

NewWhatsappMessageOrderItemWithDefaults instantiates a new WhatsappMessageOrderItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetailerId

`func (o *WhatsappMessageOrderItem) GetRetailerId() string`

GetRetailerId returns the RetailerId field if non-nil, zero value otherwise.

### GetRetailerIdOk

`func (o *WhatsappMessageOrderItem) GetRetailerIdOk() (*string, bool)`

GetRetailerIdOk returns a tuple with the RetailerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetailerId

`func (o *WhatsappMessageOrderItem) SetRetailerId(v string)`

SetRetailerId sets RetailerId field to given value.

### HasRetailerId

`func (o *WhatsappMessageOrderItem) HasRetailerId() bool`

HasRetailerId returns a boolean if a field has been set.

### GetName

`func (o *WhatsappMessageOrderItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WhatsappMessageOrderItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WhatsappMessageOrderItem) SetName(v string)`

SetName sets Name field to given value.


### GetImage

`func (o *WhatsappMessageOrderItem) GetImage() WhatsappMessageMedia`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *WhatsappMessageOrderItem) GetImageOk() (*WhatsappMessageMedia, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *WhatsappMessageOrderItem) SetImage(v WhatsappMessageMedia)`

SetImage sets Image field to given value.

### HasImage

`func (o *WhatsappMessageOrderItem) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetAmount

`func (o *WhatsappMessageOrderItem) GetAmount() WhatsappMessageOrderAmount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *WhatsappMessageOrderItem) GetAmountOk() (*WhatsappMessageOrderAmount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *WhatsappMessageOrderItem) SetAmount(v WhatsappMessageOrderAmount)`

SetAmount sets Amount field to given value.


### GetSaleAmount

`func (o *WhatsappMessageOrderItem) GetSaleAmount() WhatsappMessageOrderAmount`

GetSaleAmount returns the SaleAmount field if non-nil, zero value otherwise.

### GetSaleAmountOk

`func (o *WhatsappMessageOrderItem) GetSaleAmountOk() (*WhatsappMessageOrderAmount, bool)`

GetSaleAmountOk returns a tuple with the SaleAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleAmount

`func (o *WhatsappMessageOrderItem) SetSaleAmount(v WhatsappMessageOrderAmount)`

SetSaleAmount sets SaleAmount field to given value.

### HasSaleAmount

`func (o *WhatsappMessageOrderItem) HasSaleAmount() bool`

HasSaleAmount returns a boolean if a field has been set.

### GetQuantity

`func (o *WhatsappMessageOrderItem) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *WhatsappMessageOrderItem) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *WhatsappMessageOrderItem) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetCountryOfOrigin

`func (o *WhatsappMessageOrderItem) GetCountryOfOrigin() string`

GetCountryOfOrigin returns the CountryOfOrigin field if non-nil, zero value otherwise.

### GetCountryOfOriginOk

`func (o *WhatsappMessageOrderItem) GetCountryOfOriginOk() (*string, bool)`

GetCountryOfOriginOk returns a tuple with the CountryOfOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryOfOrigin

`func (o *WhatsappMessageOrderItem) SetCountryOfOrigin(v string)`

SetCountryOfOrigin sets CountryOfOrigin field to given value.

### HasCountryOfOrigin

`func (o *WhatsappMessageOrderItem) HasCountryOfOrigin() bool`

HasCountryOfOrigin returns a boolean if a field has been set.

### GetImporterName

`func (o *WhatsappMessageOrderItem) GetImporterName() string`

GetImporterName returns the ImporterName field if non-nil, zero value otherwise.

### GetImporterNameOk

`func (o *WhatsappMessageOrderItem) GetImporterNameOk() (*string, bool)`

GetImporterNameOk returns a tuple with the ImporterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImporterName

`func (o *WhatsappMessageOrderItem) SetImporterName(v string)`

SetImporterName sets ImporterName field to given value.

### HasImporterName

`func (o *WhatsappMessageOrderItem) HasImporterName() bool`

HasImporterName returns a boolean if a field has been set.

### GetImporterAddress

`func (o *WhatsappMessageOrderItem) GetImporterAddress() string`

GetImporterAddress returns the ImporterAddress field if non-nil, zero value otherwise.

### GetImporterAddressOk

`func (o *WhatsappMessageOrderItem) GetImporterAddressOk() (*string, bool)`

GetImporterAddressOk returns a tuple with the ImporterAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImporterAddress

`func (o *WhatsappMessageOrderItem) SetImporterAddress(v string)`

SetImporterAddress sets ImporterAddress field to given value.

### HasImporterAddress

`func (o *WhatsappMessageOrderItem) HasImporterAddress() bool`

HasImporterAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


