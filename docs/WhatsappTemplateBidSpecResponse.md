# WhatsappTemplateBidSpecResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxBid** | **string** | Base maximum price per delivered marketing message, in your YCloud billing currency, before country and per-message multipliers. Must be greater than 0, with at most 4 decimal places. This is a price ceiling, not a fixed charge. |
**CountryPriceAdjustments** | Pointer to [**[]WhatsappCountryPriceAdjustment**](WhatsappCountryPriceAdjustment.md) | Complete list of overrides for up to 50 distinct countries. Unlisted countries use a multiplier of 1. The submitted list replaces the locally stored list; omitting it or sending an empty list does not explicitly clear country overrides already stored by Meta. To remove all bidding settings, edit the template with bidSpec: {}. | [optional]

## Methods

### NewWhatsappTemplateBidSpecResponse

`func NewWhatsappTemplateBidSpecResponse(maxBid string, ) *WhatsappTemplateBidSpecResponse`

NewWhatsappTemplateBidSpecResponse instantiates a new WhatsappTemplateBidSpecResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappTemplateBidSpecResponseWithDefaults

`func NewWhatsappTemplateBidSpecResponseWithDefaults() *WhatsappTemplateBidSpecResponse`

NewWhatsappTemplateBidSpecResponseWithDefaults instantiates a new WhatsappTemplateBidSpecResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxBid

`func (o *WhatsappTemplateBidSpecResponse) GetMaxBid() string`

GetMaxBid returns the MaxBid field if non-nil, zero value otherwise.

### GetMaxBidOk

`func (o *WhatsappTemplateBidSpecResponse) GetMaxBidOk() (*string, bool)`

GetMaxBidOk returns a tuple with the MaxBid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBid

`func (o *WhatsappTemplateBidSpecResponse) SetMaxBid(v string)`

SetMaxBid sets MaxBid field to given value.


### GetCountryPriceAdjustments

`func (o *WhatsappTemplateBidSpecResponse) GetCountryPriceAdjustments() []WhatsappCountryPriceAdjustment`

GetCountryPriceAdjustments returns the CountryPriceAdjustments field if non-nil, zero value otherwise.

### GetCountryPriceAdjustmentsOk

`func (o *WhatsappTemplateBidSpecResponse) GetCountryPriceAdjustmentsOk() (*[]WhatsappCountryPriceAdjustment, bool)`

GetCountryPriceAdjustmentsOk returns a tuple with the CountryPriceAdjustments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryPriceAdjustments

`func (o *WhatsappTemplateBidSpecResponse) SetCountryPriceAdjustments(v []WhatsappCountryPriceAdjustment)`

SetCountryPriceAdjustments sets CountryPriceAdjustments field to given value.

### HasCountryPriceAdjustments

`func (o *WhatsappTemplateBidSpecResponse) HasCountryPriceAdjustments() bool`

HasCountryPriceAdjustments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
