# WhatsappBusinessAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | WhatApp Business Account ID. | [optional] 
**Name** | Pointer to **string** | User-friendly name to differentiate WhatsApp Business Accounts. | [optional] 
**Currency** | Pointer to **string** | The currency in which the payment transactions for the WhatsApp Business Account will be processed. | [optional] 
**MessageTemplateNamespace** | Pointer to **string** | Namespace string for the message templates that belong to the WhatsApp Business Account. | [optional] 
**AccountReviewStatus** | Pointer to [**WhatsappBusinessAccountReviewStatus**](WhatsappBusinessAccountReviewStatus.md) |  | [optional] 
**BusinessVerificationStatus** | Pointer to [**MetaBusinessAccountVerificationStatus**](MetaBusinessAccountVerificationStatus.md) |  | [optional] 
**Country** | Pointer to **string** | Country of the WhatsApp Business Account&#39;s owning Meta Business account. | [optional] 
**OwnershipType** | Pointer to **string** | Ownership type of the WhatsApp Business Account. | [optional] 
**PaymentMethodAttached** | Pointer to **bool** | Whether we have attached a payment method to the WhatsApp Business Account. | [optional] 
**PrimaryFundingId** | Pointer to **string** | Primary funding ID for the WhatsApp Business Account paid service. | [optional] 
**PurchaseOrderNumber** | Pointer to **string** | The purchase order number supplied by the business for payment management purposes. | [optional] 
**TimezoneId** | Pointer to **string** | The timezone ID of the WhatsApp Business Account. See [Timezone IDs](https://developers.facebook.com/docs/marketing-api/reference/ad-account/timezone-ids). | [optional] 
**Decision** | Pointer to [**WhatsappReviewDecision**](WhatsappReviewDecision.md) |  | [optional] 
**UpdateEvent** | Pointer to [**WhatsappBusinessAccountUpdateEventEnum**](WhatsappBusinessAccountUpdateEventEnum.md) |  | [optional] 
**BanState** | Pointer to [**WhatsappBusinessAccountBanState**](WhatsappBusinessAccountBanState.md) |  | [optional] 
**BanDate** | Pointer to **string** | The date when the WABA is banned. | [optional] 
**ViolationType** | Pointer to **string** | Used to report violations imposed on the WABA. See also [WhatsApp Business Platform Policy Violations](https://developers.facebook.com/docs/whatsapp/overview/policy-enforcement/violations). | [optional] 
**Restrictions** | Pointer to [**[]WhatsappBusinessAccountRestrictionInfo**](WhatsappBusinessAccountRestrictionInfo.md) | Used to report restrictions imposed on the WABA, when that WABA violates [WhatsApp Business Platform policies](https://developers.facebook.com/docs/whatsapp/overview/policy-enforcement). | [optional] 
**AuthIntlRateEligibilityCountries** | Pointer to [**[]WhatsappAuthIntlRateEligibilityCountry**](WhatsappAuthIntlRateEligibilityCountry.md) | Starting June 1, 2024, we are updating our authentication rate card and introducing a new authentication-international rate. This rate will apply in the the following countries: - June 1, 2024 – Indonesia (country calling code +62, country code &#x60;ID&#x60;) - July 1, 2024 – India (country calling code +91, country code &#x60;IN&#x60;)  See also [Authentication-International Rates](https://developers.facebook.com/docs/whatsapp/pricing/authentication-international-rates). | [optional] 
**PrimaryBusinessLocation** | Pointer to **string** | Your primary business location is the country where your business is based. It will appear in the Business Manager under the Primary Business Location field starting May 1, 2024. [ISO 3166-1 alpha-2 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2). | [optional] 

## Methods

### NewWhatsappBusinessAccount

`func NewWhatsappBusinessAccount() *WhatsappBusinessAccount`

NewWhatsappBusinessAccount instantiates a new WhatsappBusinessAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappBusinessAccountWithDefaults

`func NewWhatsappBusinessAccountWithDefaults() *WhatsappBusinessAccount`

NewWhatsappBusinessAccountWithDefaults instantiates a new WhatsappBusinessAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WhatsappBusinessAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WhatsappBusinessAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WhatsappBusinessAccount) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WhatsappBusinessAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WhatsappBusinessAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WhatsappBusinessAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WhatsappBusinessAccount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WhatsappBusinessAccount) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCurrency

`func (o *WhatsappBusinessAccount) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *WhatsappBusinessAccount) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *WhatsappBusinessAccount) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *WhatsappBusinessAccount) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetMessageTemplateNamespace

`func (o *WhatsappBusinessAccount) GetMessageTemplateNamespace() string`

GetMessageTemplateNamespace returns the MessageTemplateNamespace field if non-nil, zero value otherwise.

### GetMessageTemplateNamespaceOk

`func (o *WhatsappBusinessAccount) GetMessageTemplateNamespaceOk() (*string, bool)`

GetMessageTemplateNamespaceOk returns a tuple with the MessageTemplateNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTemplateNamespace

`func (o *WhatsappBusinessAccount) SetMessageTemplateNamespace(v string)`

SetMessageTemplateNamespace sets MessageTemplateNamespace field to given value.

### HasMessageTemplateNamespace

`func (o *WhatsappBusinessAccount) HasMessageTemplateNamespace() bool`

HasMessageTemplateNamespace returns a boolean if a field has been set.

### GetAccountReviewStatus

`func (o *WhatsappBusinessAccount) GetAccountReviewStatus() WhatsappBusinessAccountReviewStatus`

GetAccountReviewStatus returns the AccountReviewStatus field if non-nil, zero value otherwise.

### GetAccountReviewStatusOk

`func (o *WhatsappBusinessAccount) GetAccountReviewStatusOk() (*WhatsappBusinessAccountReviewStatus, bool)`

GetAccountReviewStatusOk returns a tuple with the AccountReviewStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountReviewStatus

`func (o *WhatsappBusinessAccount) SetAccountReviewStatus(v WhatsappBusinessAccountReviewStatus)`

SetAccountReviewStatus sets AccountReviewStatus field to given value.

### HasAccountReviewStatus

`func (o *WhatsappBusinessAccount) HasAccountReviewStatus() bool`

HasAccountReviewStatus returns a boolean if a field has been set.

### GetBusinessVerificationStatus

`func (o *WhatsappBusinessAccount) GetBusinessVerificationStatus() MetaBusinessAccountVerificationStatus`

GetBusinessVerificationStatus returns the BusinessVerificationStatus field if non-nil, zero value otherwise.

### GetBusinessVerificationStatusOk

`func (o *WhatsappBusinessAccount) GetBusinessVerificationStatusOk() (*MetaBusinessAccountVerificationStatus, bool)`

GetBusinessVerificationStatusOk returns a tuple with the BusinessVerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessVerificationStatus

`func (o *WhatsappBusinessAccount) SetBusinessVerificationStatus(v MetaBusinessAccountVerificationStatus)`

SetBusinessVerificationStatus sets BusinessVerificationStatus field to given value.

### HasBusinessVerificationStatus

`func (o *WhatsappBusinessAccount) HasBusinessVerificationStatus() bool`

HasBusinessVerificationStatus returns a boolean if a field has been set.

### GetCountry

`func (o *WhatsappBusinessAccount) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *WhatsappBusinessAccount) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *WhatsappBusinessAccount) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *WhatsappBusinessAccount) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetOwnershipType

`func (o *WhatsappBusinessAccount) GetOwnershipType() string`

GetOwnershipType returns the OwnershipType field if non-nil, zero value otherwise.

### GetOwnershipTypeOk

`func (o *WhatsappBusinessAccount) GetOwnershipTypeOk() (*string, bool)`

GetOwnershipTypeOk returns a tuple with the OwnershipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnershipType

`func (o *WhatsappBusinessAccount) SetOwnershipType(v string)`

SetOwnershipType sets OwnershipType field to given value.

### HasOwnershipType

`func (o *WhatsappBusinessAccount) HasOwnershipType() bool`

HasOwnershipType returns a boolean if a field has been set.

### GetPaymentMethodAttached

`func (o *WhatsappBusinessAccount) GetPaymentMethodAttached() bool`

GetPaymentMethodAttached returns the PaymentMethodAttached field if non-nil, zero value otherwise.

### GetPaymentMethodAttachedOk

`func (o *WhatsappBusinessAccount) GetPaymentMethodAttachedOk() (*bool, bool)`

GetPaymentMethodAttachedOk returns a tuple with the PaymentMethodAttached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodAttached

`func (o *WhatsappBusinessAccount) SetPaymentMethodAttached(v bool)`

SetPaymentMethodAttached sets PaymentMethodAttached field to given value.

### HasPaymentMethodAttached

`func (o *WhatsappBusinessAccount) HasPaymentMethodAttached() bool`

HasPaymentMethodAttached returns a boolean if a field has been set.

### GetPrimaryFundingId

`func (o *WhatsappBusinessAccount) GetPrimaryFundingId() string`

GetPrimaryFundingId returns the PrimaryFundingId field if non-nil, zero value otherwise.

### GetPrimaryFundingIdOk

`func (o *WhatsappBusinessAccount) GetPrimaryFundingIdOk() (*string, bool)`

GetPrimaryFundingIdOk returns a tuple with the PrimaryFundingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryFundingId

`func (o *WhatsappBusinessAccount) SetPrimaryFundingId(v string)`

SetPrimaryFundingId sets PrimaryFundingId field to given value.

### HasPrimaryFundingId

`func (o *WhatsappBusinessAccount) HasPrimaryFundingId() bool`

HasPrimaryFundingId returns a boolean if a field has been set.

### GetPurchaseOrderNumber

`func (o *WhatsappBusinessAccount) GetPurchaseOrderNumber() string`

GetPurchaseOrderNumber returns the PurchaseOrderNumber field if non-nil, zero value otherwise.

### GetPurchaseOrderNumberOk

`func (o *WhatsappBusinessAccount) GetPurchaseOrderNumberOk() (*string, bool)`

GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderNumber

`func (o *WhatsappBusinessAccount) SetPurchaseOrderNumber(v string)`

SetPurchaseOrderNumber sets PurchaseOrderNumber field to given value.

### HasPurchaseOrderNumber

`func (o *WhatsappBusinessAccount) HasPurchaseOrderNumber() bool`

HasPurchaseOrderNumber returns a boolean if a field has been set.

### GetTimezoneId

`func (o *WhatsappBusinessAccount) GetTimezoneId() string`

GetTimezoneId returns the TimezoneId field if non-nil, zero value otherwise.

### GetTimezoneIdOk

`func (o *WhatsappBusinessAccount) GetTimezoneIdOk() (*string, bool)`

GetTimezoneIdOk returns a tuple with the TimezoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneId

`func (o *WhatsappBusinessAccount) SetTimezoneId(v string)`

SetTimezoneId sets TimezoneId field to given value.

### HasTimezoneId

`func (o *WhatsappBusinessAccount) HasTimezoneId() bool`

HasTimezoneId returns a boolean if a field has been set.

### GetDecision

`func (o *WhatsappBusinessAccount) GetDecision() WhatsappReviewDecision`

GetDecision returns the Decision field if non-nil, zero value otherwise.

### GetDecisionOk

`func (o *WhatsappBusinessAccount) GetDecisionOk() (*WhatsappReviewDecision, bool)`

GetDecisionOk returns a tuple with the Decision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecision

`func (o *WhatsappBusinessAccount) SetDecision(v WhatsappReviewDecision)`

SetDecision sets Decision field to given value.

### HasDecision

`func (o *WhatsappBusinessAccount) HasDecision() bool`

HasDecision returns a boolean if a field has been set.

### GetUpdateEvent

`func (o *WhatsappBusinessAccount) GetUpdateEvent() WhatsappBusinessAccountUpdateEventEnum`

GetUpdateEvent returns the UpdateEvent field if non-nil, zero value otherwise.

### GetUpdateEventOk

`func (o *WhatsappBusinessAccount) GetUpdateEventOk() (*WhatsappBusinessAccountUpdateEventEnum, bool)`

GetUpdateEventOk returns a tuple with the UpdateEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateEvent

`func (o *WhatsappBusinessAccount) SetUpdateEvent(v WhatsappBusinessAccountUpdateEventEnum)`

SetUpdateEvent sets UpdateEvent field to given value.

### HasUpdateEvent

`func (o *WhatsappBusinessAccount) HasUpdateEvent() bool`

HasUpdateEvent returns a boolean if a field has been set.

### GetBanState

`func (o *WhatsappBusinessAccount) GetBanState() WhatsappBusinessAccountBanState`

GetBanState returns the BanState field if non-nil, zero value otherwise.

### GetBanStateOk

`func (o *WhatsappBusinessAccount) GetBanStateOk() (*WhatsappBusinessAccountBanState, bool)`

GetBanStateOk returns a tuple with the BanState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanState

`func (o *WhatsappBusinessAccount) SetBanState(v WhatsappBusinessAccountBanState)`

SetBanState sets BanState field to given value.

### HasBanState

`func (o *WhatsappBusinessAccount) HasBanState() bool`

HasBanState returns a boolean if a field has been set.

### GetBanDate

`func (o *WhatsappBusinessAccount) GetBanDate() string`

GetBanDate returns the BanDate field if non-nil, zero value otherwise.

### GetBanDateOk

`func (o *WhatsappBusinessAccount) GetBanDateOk() (*string, bool)`

GetBanDateOk returns a tuple with the BanDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanDate

`func (o *WhatsappBusinessAccount) SetBanDate(v string)`

SetBanDate sets BanDate field to given value.

### HasBanDate

`func (o *WhatsappBusinessAccount) HasBanDate() bool`

HasBanDate returns a boolean if a field has been set.

### GetViolationType

`func (o *WhatsappBusinessAccount) GetViolationType() string`

GetViolationType returns the ViolationType field if non-nil, zero value otherwise.

### GetViolationTypeOk

`func (o *WhatsappBusinessAccount) GetViolationTypeOk() (*string, bool)`

GetViolationTypeOk returns a tuple with the ViolationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolationType

`func (o *WhatsappBusinessAccount) SetViolationType(v string)`

SetViolationType sets ViolationType field to given value.

### HasViolationType

`func (o *WhatsappBusinessAccount) HasViolationType() bool`

HasViolationType returns a boolean if a field has been set.

### GetRestrictions

`func (o *WhatsappBusinessAccount) GetRestrictions() []WhatsappBusinessAccountRestrictionInfo`

GetRestrictions returns the Restrictions field if non-nil, zero value otherwise.

### GetRestrictionsOk

`func (o *WhatsappBusinessAccount) GetRestrictionsOk() (*[]WhatsappBusinessAccountRestrictionInfo, bool)`

GetRestrictionsOk returns a tuple with the Restrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictions

`func (o *WhatsappBusinessAccount) SetRestrictions(v []WhatsappBusinessAccountRestrictionInfo)`

SetRestrictions sets Restrictions field to given value.

### HasRestrictions

`func (o *WhatsappBusinessAccount) HasRestrictions() bool`

HasRestrictions returns a boolean if a field has been set.

### GetAuthIntlRateEligibilityCountries

`func (o *WhatsappBusinessAccount) GetAuthIntlRateEligibilityCountries() []WhatsappAuthIntlRateEligibilityCountry`

GetAuthIntlRateEligibilityCountries returns the AuthIntlRateEligibilityCountries field if non-nil, zero value otherwise.

### GetAuthIntlRateEligibilityCountriesOk

`func (o *WhatsappBusinessAccount) GetAuthIntlRateEligibilityCountriesOk() (*[]WhatsappAuthIntlRateEligibilityCountry, bool)`

GetAuthIntlRateEligibilityCountriesOk returns a tuple with the AuthIntlRateEligibilityCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthIntlRateEligibilityCountries

`func (o *WhatsappBusinessAccount) SetAuthIntlRateEligibilityCountries(v []WhatsappAuthIntlRateEligibilityCountry)`

SetAuthIntlRateEligibilityCountries sets AuthIntlRateEligibilityCountries field to given value.

### HasAuthIntlRateEligibilityCountries

`func (o *WhatsappBusinessAccount) HasAuthIntlRateEligibilityCountries() bool`

HasAuthIntlRateEligibilityCountries returns a boolean if a field has been set.

### GetPrimaryBusinessLocation

`func (o *WhatsappBusinessAccount) GetPrimaryBusinessLocation() string`

GetPrimaryBusinessLocation returns the PrimaryBusinessLocation field if non-nil, zero value otherwise.

### GetPrimaryBusinessLocationOk

`func (o *WhatsappBusinessAccount) GetPrimaryBusinessLocationOk() (*string, bool)`

GetPrimaryBusinessLocationOk returns a tuple with the PrimaryBusinessLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryBusinessLocation

`func (o *WhatsappBusinessAccount) SetPrimaryBusinessLocation(v string)`

SetPrimaryBusinessLocation sets PrimaryBusinessLocation field to given value.

### HasPrimaryBusinessLocation

`func (o *WhatsappBusinessAccount) HasPrimaryBusinessLocation() bool`

HasPrimaryBusinessLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


