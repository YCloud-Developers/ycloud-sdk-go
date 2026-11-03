# WhatsappCountryPriceAdjustment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | **string** | An uppercase ISO 3166-1 alpha-2 country code. Each country may appear only once in the list. |
**Multiplier** | **string** | Multiplier applied to the base maxBid for this country. Must be greater than 0 and at most 10, with at most 4 decimal places. |

## Methods

### NewWhatsappCountryPriceAdjustment

`func NewWhatsappCountryPriceAdjustment(countryCode string, multiplier string, ) *WhatsappCountryPriceAdjustment`

NewWhatsappCountryPriceAdjustment instantiates a new WhatsappCountryPriceAdjustment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappCountryPriceAdjustmentWithDefaults

`func NewWhatsappCountryPriceAdjustmentWithDefaults() *WhatsappCountryPriceAdjustment`

NewWhatsappCountryPriceAdjustmentWithDefaults instantiates a new WhatsappCountryPriceAdjustment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *WhatsappCountryPriceAdjustment) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *WhatsappCountryPriceAdjustment) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *WhatsappCountryPriceAdjustment) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetMultiplier

`func (o *WhatsappCountryPriceAdjustment) GetMultiplier() string`

GetMultiplier returns the Multiplier field if non-nil, zero value otherwise.

### GetMultiplierOk

`func (o *WhatsappCountryPriceAdjustment) GetMultiplierOk() (*string, bool)`

GetMultiplierOk returns a tuple with the Multiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplier

`func (o *WhatsappCountryPriceAdjustment) SetMultiplier(v string)`

SetMultiplier sets Multiplier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
