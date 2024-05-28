# WhatsappAuthIntlRateEligibilityCountry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | Pointer to **string** | [ISO 3166-1 alpha-2 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2). | [optional] 
**StartTime** | Pointer to **time.Time** | Date when newly-opened authentication conversations are subject to authentication-international rates, formatted in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339). e.g., &#x60;2024-07-01T00:00:00.000Z&#x60;. | [optional] 

## Methods

### NewWhatsappAuthIntlRateEligibilityCountry

`func NewWhatsappAuthIntlRateEligibilityCountry() *WhatsappAuthIntlRateEligibilityCountry`

NewWhatsappAuthIntlRateEligibilityCountry instantiates a new WhatsappAuthIntlRateEligibilityCountry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappAuthIntlRateEligibilityCountryWithDefaults

`func NewWhatsappAuthIntlRateEligibilityCountryWithDefaults() *WhatsappAuthIntlRateEligibilityCountry`

NewWhatsappAuthIntlRateEligibilityCountryWithDefaults instantiates a new WhatsappAuthIntlRateEligibilityCountry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *WhatsappAuthIntlRateEligibilityCountry) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *WhatsappAuthIntlRateEligibilityCountry) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *WhatsappAuthIntlRateEligibilityCountry) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *WhatsappAuthIntlRateEligibilityCountry) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetStartTime

`func (o *WhatsappAuthIntlRateEligibilityCountry) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *WhatsappAuthIntlRateEligibilityCountry) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *WhatsappAuthIntlRateEligibilityCountry) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *WhatsappAuthIntlRateEligibilityCountry) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


