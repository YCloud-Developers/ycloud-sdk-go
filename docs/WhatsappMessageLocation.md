# WhatsappMessageLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Latitude** | **float64** | Latitude of the location. | 
**Longitude** | **float64** | Longitude of the location. | 
**Name** | Pointer to **string** | Name of the location. | [optional] 
**Address** | Pointer to **string** | Address of the location. Only displayed if &#x60;name&#x60; is present. | [optional] 

## Methods

### NewWhatsappMessageLocation

`func NewWhatsappMessageLocation(latitude float64, longitude float64, ) *WhatsappMessageLocation`

NewWhatsappMessageLocation instantiates a new WhatsappMessageLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappMessageLocationWithDefaults

`func NewWhatsappMessageLocationWithDefaults() *WhatsappMessageLocation`

NewWhatsappMessageLocationWithDefaults instantiates a new WhatsappMessageLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLatitude

`func (o *WhatsappMessageLocation) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *WhatsappMessageLocation) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *WhatsappMessageLocation) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.


### GetLongitude

`func (o *WhatsappMessageLocation) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *WhatsappMessageLocation) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *WhatsappMessageLocation) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.


### GetName

`func (o *WhatsappMessageLocation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WhatsappMessageLocation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WhatsappMessageLocation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WhatsappMessageLocation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddress

`func (o *WhatsappMessageLocation) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *WhatsappMessageLocation) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *WhatsappMessageLocation) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *WhatsappMessageLocation) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
