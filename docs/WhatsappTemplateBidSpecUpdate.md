# WhatsappTemplateBidSpecUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxBid** | Pointer to **NullableString** | Base maximum price per delivered marketing message, in your YCloud billing currency, before country and per-message multipliers. Must be greater than 0, with at most 4 decimal places. This is a price ceiling, not a fixed charge. | [optional]
**CountryPriceAdjustments** | Pointer to [**[]WhatsappCountryPriceAdjustment**](WhatsappCountryPriceAdjustment.md) | Complete list of overrides for up to 50 distinct countries. Unlisted countries use a multiplier of 1. The submitted list replaces the locally stored list; omitting it or sending an empty list does not explicitly clear country overrides already stored by Meta. To remove all bidding settings, edit the template with bidSpec: {}. | [optional]

## Methods

### NewWhatsappTemplateBidSpecUpdate

`func NewWhatsappTemplateBidSpecUpdate() *WhatsappTemplateBidSpecUpdate`

NewWhatsappTemplateBidSpecUpdate instantiates a new WhatsappTemplateBidSpecUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappTemplateBidSpecUpdateWithDefaults

`func NewWhatsappTemplateBidSpecUpdateWithDefaults() *WhatsappTemplateBidSpecUpdate`

NewWhatsappTemplateBidSpecUpdateWithDefaults instantiates a new WhatsappTemplateBidSpecUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxBid

`func (o *WhatsappTemplateBidSpecUpdate) GetMaxBid() string`

GetMaxBid returns the MaxBid field if non-nil, zero value otherwise.

### GetMaxBidOk

`func (o *WhatsappTemplateBidSpecUpdate) GetMaxBidOk() (*string, bool)`

GetMaxBidOk returns a tuple with the MaxBid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBid

`func (o *WhatsappTemplateBidSpecUpdate) SetMaxBid(v string)`

SetMaxBid sets MaxBid field to given value.

### HasMaxBid

`func (o *WhatsappTemplateBidSpecUpdate) HasMaxBid() bool`

HasMaxBid returns a boolean if a field has been set.

### SetMaxBidNil

`func (o *WhatsappTemplateBidSpecUpdate) SetMaxBidNil(b bool)`

 SetMaxBidNil sets the value for MaxBid to be an explicit nil

### UnsetMaxBid
`func (o *WhatsappTemplateBidSpecUpdate) UnsetMaxBid()`

UnsetMaxBid ensures that no value is present for MaxBid, not even an explicit nil
### GetCountryPriceAdjustments

`func (o *WhatsappTemplateBidSpecUpdate) GetCountryPriceAdjustments() []WhatsappCountryPriceAdjustment`

GetCountryPriceAdjustments returns the CountryPriceAdjustments field if non-nil, zero value otherwise.

### GetCountryPriceAdjustmentsOk

`func (o *WhatsappTemplateBidSpecUpdate) GetCountryPriceAdjustmentsOk() (*[]WhatsappCountryPriceAdjustment, bool)`

GetCountryPriceAdjustmentsOk returns a tuple with the CountryPriceAdjustments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryPriceAdjustments

`func (o *WhatsappTemplateBidSpecUpdate) SetCountryPriceAdjustments(v []WhatsappCountryPriceAdjustment)`

SetCountryPriceAdjustments sets CountryPriceAdjustments field to given value.

### HasCountryPriceAdjustments

`func (o *WhatsappTemplateBidSpecUpdate) HasCountryPriceAdjustments() bool`

HasCountryPriceAdjustments returns a boolean if a field has been set.

### SetCountryPriceAdjustmentsNil

`func (o *WhatsappTemplateBidSpecUpdate) SetCountryPriceAdjustmentsNil(b bool)`

 SetCountryPriceAdjustmentsNil sets the value for CountryPriceAdjustments to be an explicit nil

### UnsetCountryPriceAdjustments
`func (o *WhatsappTemplateBidSpecUpdate) UnsetCountryPriceAdjustments()`

UnsetCountryPriceAdjustments ensures that no value is present for CountryPriceAdjustments, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
