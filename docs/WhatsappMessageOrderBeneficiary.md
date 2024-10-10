# WhatsappMessageOrderBeneficiary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the individual or business receiving the physical goods. Cannot exceed 200 characters. | 
**AddressLine1** | **string** | Shipping address (Door/Tower Number, Street Name etc.). Cannot exceed 100 characters. | 
**AddressLine2** | Pointer to **string** | Shipping address (Landmark, Area, etc.). Cannot exceed 100 characters. | [optional] 
**City** | **string** | Name of the city. | 
**State** | **string** | Name of the state. | 
**Country** | **string** | Name of the country. Currently the only supported value is &#x60;India&#x60;. | 
**PostalCode** | **string** | 6-digit zipcode of shipping address. | 

## Methods

### NewWhatsappMessageOrderBeneficiary

`func NewWhatsappMessageOrderBeneficiary(name string, addressLine1 string, city string, state string, country string, postalCode string, ) *WhatsappMessageOrderBeneficiary`

NewWhatsappMessageOrderBeneficiary instantiates a new WhatsappMessageOrderBeneficiary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappMessageOrderBeneficiaryWithDefaults

`func NewWhatsappMessageOrderBeneficiaryWithDefaults() *WhatsappMessageOrderBeneficiary`

NewWhatsappMessageOrderBeneficiaryWithDefaults instantiates a new WhatsappMessageOrderBeneficiary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WhatsappMessageOrderBeneficiary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WhatsappMessageOrderBeneficiary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WhatsappMessageOrderBeneficiary) SetName(v string)`

SetName sets Name field to given value.


### GetAddressLine1

`func (o *WhatsappMessageOrderBeneficiary) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *WhatsappMessageOrderBeneficiary) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *WhatsappMessageOrderBeneficiary) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.


### GetAddressLine2

`func (o *WhatsappMessageOrderBeneficiary) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *WhatsappMessageOrderBeneficiary) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *WhatsappMessageOrderBeneficiary) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *WhatsappMessageOrderBeneficiary) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetCity

`func (o *WhatsappMessageOrderBeneficiary) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *WhatsappMessageOrderBeneficiary) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *WhatsappMessageOrderBeneficiary) SetCity(v string)`

SetCity sets City field to given value.


### GetState

`func (o *WhatsappMessageOrderBeneficiary) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *WhatsappMessageOrderBeneficiary) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *WhatsappMessageOrderBeneficiary) SetState(v string)`

SetState sets State field to given value.


### GetCountry

`func (o *WhatsappMessageOrderBeneficiary) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *WhatsappMessageOrderBeneficiary) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *WhatsappMessageOrderBeneficiary) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetPostalCode

`func (o *WhatsappMessageOrderBeneficiary) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *WhatsappMessageOrderBeneficiary) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *WhatsappMessageOrderBeneficiary) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


