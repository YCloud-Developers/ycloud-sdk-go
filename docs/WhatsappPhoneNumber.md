# WhatsappPhoneNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Phone number ID. | [optional] 
**PhoneNumber** | Pointer to **string** | Phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format. | [optional] 
**DisplayPhoneNumber** | Pointer to **string** | Display phone number. | [optional] 
**WabaId** | Pointer to **string** | WhatsApp Business Account ID. | [optional] 
**QualityRating** | Pointer to [**WhatsappPhoneNumberQualityRating**](WhatsappPhoneNumberQualityRating.md) |  | [optional] 
**MessagingLimit** | Pointer to **string** | Messaging limits determine the maximum number of business-initiated conversations each phone number can start in a rolling 24-hour period. See also [Messaging Limits](https://developers.facebook.com/docs/whatsapp/messaging-limits#messaging-limits). - &#x60;TIER_NOT_SET&#x60;: Unknown limit. - &#x60;TIER_50&#x60;: 50 business-initiated conversations in a rolling 24-hour period. - &#x60;TIER_250&#x60;: 250 business-initiated conversations in a rolling 24-hour period. - &#x60;TIER_1K&#x60;: 1K business-initiated conversations with unique customers in a rolling 24-hour period. - &#x60;TIER_10K&#x60;: 10K business-initiated conversations with unique customers in a rolling 24-hour period. - &#x60;TIER_100K&#x60;: 100K business-initiated conversations with unique customers in a rolling 24-hour period. - &#x60;TIER_UNLIMITED&#x60;: An unlimited number of business-initiated conversations in a rolling 24-hour period. | [optional] 
**VerifiedName** | Pointer to **string** | Verified name. | [optional] 
**CodeVerificationStatus** | Pointer to [**WhatsappPhoneNumberCodeVerificationStatus**](WhatsappPhoneNumberCodeVerificationStatus.md) |  | [optional] 
**IsOfficialBusinessAccount** | Pointer to **bool** | Whether this phone number is an official business account or not. An official business account has a green checkmark badge in its profile and chat thread headers. See [Official Business Account](https://developers.facebook.com/docs/whatsapp/overview/business-accounts#official-business-account) for more information. | [optional] 
**Status** | Pointer to [**WhatsappPhoneNumberStatus**](WhatsappPhoneNumberStatus.md) |  | [optional] 
**NameStatus** | Pointer to [**WhatsappPhoneNumberNameStatus**](WhatsappPhoneNumberNameStatus.md) |  | [optional] 
**NewNameStatus** | Pointer to [**WhatsappPhoneNumberNameStatus**](WhatsappPhoneNumberNameStatus.md) |  | [optional] 
**Decision** | Pointer to [**WhatsappReviewDecision**](WhatsappReviewDecision.md) |  | [optional] 
**RequestedVerifiedName** | Pointer to **string** | Last requested verified name. See [Phone Number Name Update](https://developers.facebook.com/docs/graph-api/webhooks/reference/whatsapp-business-account/#phone_number_name_update). | [optional] 
**RejectionReason** | Pointer to **string** | Rejection reason. See [Phone Number Name Update](https://developers.facebook.com/docs/graph-api/webhooks/reference/whatsapp-business-account/#phone_number_name_update). | [optional] 
**QualityUpdateEvent** | Pointer to [**WhatsappPhoneNumberQualityUpdateEventEnum**](WhatsappPhoneNumberQualityUpdateEventEnum.md) |  | [optional] 

## Methods

### NewWhatsappPhoneNumber

`func NewWhatsappPhoneNumber() *WhatsappPhoneNumber`

NewWhatsappPhoneNumber instantiates a new WhatsappPhoneNumber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappPhoneNumberWithDefaults

`func NewWhatsappPhoneNumberWithDefaults() *WhatsappPhoneNumber`

NewWhatsappPhoneNumberWithDefaults instantiates a new WhatsappPhoneNumber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WhatsappPhoneNumber) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WhatsappPhoneNumber) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WhatsappPhoneNumber) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WhatsappPhoneNumber) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *WhatsappPhoneNumber) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *WhatsappPhoneNumber) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *WhatsappPhoneNumber) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *WhatsappPhoneNumber) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetDisplayPhoneNumber

`func (o *WhatsappPhoneNumber) GetDisplayPhoneNumber() string`

GetDisplayPhoneNumber returns the DisplayPhoneNumber field if non-nil, zero value otherwise.

### GetDisplayPhoneNumberOk

`func (o *WhatsappPhoneNumber) GetDisplayPhoneNumberOk() (*string, bool)`

GetDisplayPhoneNumberOk returns a tuple with the DisplayPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayPhoneNumber

`func (o *WhatsappPhoneNumber) SetDisplayPhoneNumber(v string)`

SetDisplayPhoneNumber sets DisplayPhoneNumber field to given value.

### HasDisplayPhoneNumber

`func (o *WhatsappPhoneNumber) HasDisplayPhoneNumber() bool`

HasDisplayPhoneNumber returns a boolean if a field has been set.

### GetWabaId

`func (o *WhatsappPhoneNumber) GetWabaId() string`

GetWabaId returns the WabaId field if non-nil, zero value otherwise.

### GetWabaIdOk

`func (o *WhatsappPhoneNumber) GetWabaIdOk() (*string, bool)`

GetWabaIdOk returns a tuple with the WabaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWabaId

`func (o *WhatsappPhoneNumber) SetWabaId(v string)`

SetWabaId sets WabaId field to given value.

### HasWabaId

`func (o *WhatsappPhoneNumber) HasWabaId() bool`

HasWabaId returns a boolean if a field has been set.

### GetQualityRating

`func (o *WhatsappPhoneNumber) GetQualityRating() WhatsappPhoneNumberQualityRating`

GetQualityRating returns the QualityRating field if non-nil, zero value otherwise.

### GetQualityRatingOk

`func (o *WhatsappPhoneNumber) GetQualityRatingOk() (*WhatsappPhoneNumberQualityRating, bool)`

GetQualityRatingOk returns a tuple with the QualityRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityRating

`func (o *WhatsappPhoneNumber) SetQualityRating(v WhatsappPhoneNumberQualityRating)`

SetQualityRating sets QualityRating field to given value.

### HasQualityRating

`func (o *WhatsappPhoneNumber) HasQualityRating() bool`

HasQualityRating returns a boolean if a field has been set.

### GetMessagingLimit

`func (o *WhatsappPhoneNumber) GetMessagingLimit() string`

GetMessagingLimit returns the MessagingLimit field if non-nil, zero value otherwise.

### GetMessagingLimitOk

`func (o *WhatsappPhoneNumber) GetMessagingLimitOk() (*string, bool)`

GetMessagingLimitOk returns a tuple with the MessagingLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingLimit

`func (o *WhatsappPhoneNumber) SetMessagingLimit(v string)`

SetMessagingLimit sets MessagingLimit field to given value.

### HasMessagingLimit

`func (o *WhatsappPhoneNumber) HasMessagingLimit() bool`

HasMessagingLimit returns a boolean if a field has been set.

### GetVerifiedName

`func (o *WhatsappPhoneNumber) GetVerifiedName() string`

GetVerifiedName returns the VerifiedName field if non-nil, zero value otherwise.

### GetVerifiedNameOk

`func (o *WhatsappPhoneNumber) GetVerifiedNameOk() (*string, bool)`

GetVerifiedNameOk returns a tuple with the VerifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedName

`func (o *WhatsappPhoneNumber) SetVerifiedName(v string)`

SetVerifiedName sets VerifiedName field to given value.

### HasVerifiedName

`func (o *WhatsappPhoneNumber) HasVerifiedName() bool`

HasVerifiedName returns a boolean if a field has been set.

### GetCodeVerificationStatus

`func (o *WhatsappPhoneNumber) GetCodeVerificationStatus() WhatsappPhoneNumberCodeVerificationStatus`

GetCodeVerificationStatus returns the CodeVerificationStatus field if non-nil, zero value otherwise.

### GetCodeVerificationStatusOk

`func (o *WhatsappPhoneNumber) GetCodeVerificationStatusOk() (*WhatsappPhoneNumberCodeVerificationStatus, bool)`

GetCodeVerificationStatusOk returns a tuple with the CodeVerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeVerificationStatus

`func (o *WhatsappPhoneNumber) SetCodeVerificationStatus(v WhatsappPhoneNumberCodeVerificationStatus)`

SetCodeVerificationStatus sets CodeVerificationStatus field to given value.

### HasCodeVerificationStatus

`func (o *WhatsappPhoneNumber) HasCodeVerificationStatus() bool`

HasCodeVerificationStatus returns a boolean if a field has been set.

### GetIsOfficialBusinessAccount

`func (o *WhatsappPhoneNumber) GetIsOfficialBusinessAccount() bool`

GetIsOfficialBusinessAccount returns the IsOfficialBusinessAccount field if non-nil, zero value otherwise.

### GetIsOfficialBusinessAccountOk

`func (o *WhatsappPhoneNumber) GetIsOfficialBusinessAccountOk() (*bool, bool)`

GetIsOfficialBusinessAccountOk returns a tuple with the IsOfficialBusinessAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOfficialBusinessAccount

`func (o *WhatsappPhoneNumber) SetIsOfficialBusinessAccount(v bool)`

SetIsOfficialBusinessAccount sets IsOfficialBusinessAccount field to given value.

### HasIsOfficialBusinessAccount

`func (o *WhatsappPhoneNumber) HasIsOfficialBusinessAccount() bool`

HasIsOfficialBusinessAccount returns a boolean if a field has been set.

### GetStatus

`func (o *WhatsappPhoneNumber) GetStatus() WhatsappPhoneNumberStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WhatsappPhoneNumber) GetStatusOk() (*WhatsappPhoneNumberStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WhatsappPhoneNumber) SetStatus(v WhatsappPhoneNumberStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WhatsappPhoneNumber) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetNameStatus

`func (o *WhatsappPhoneNumber) GetNameStatus() WhatsappPhoneNumberNameStatus`

GetNameStatus returns the NameStatus field if non-nil, zero value otherwise.

### GetNameStatusOk

`func (o *WhatsappPhoneNumber) GetNameStatusOk() (*WhatsappPhoneNumberNameStatus, bool)`

GetNameStatusOk returns a tuple with the NameStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameStatus

`func (o *WhatsappPhoneNumber) SetNameStatus(v WhatsappPhoneNumberNameStatus)`

SetNameStatus sets NameStatus field to given value.

### HasNameStatus

`func (o *WhatsappPhoneNumber) HasNameStatus() bool`

HasNameStatus returns a boolean if a field has been set.

### GetNewNameStatus

`func (o *WhatsappPhoneNumber) GetNewNameStatus() WhatsappPhoneNumberNameStatus`

GetNewNameStatus returns the NewNameStatus field if non-nil, zero value otherwise.

### GetNewNameStatusOk

`func (o *WhatsappPhoneNumber) GetNewNameStatusOk() (*WhatsappPhoneNumberNameStatus, bool)`

GetNewNameStatusOk returns a tuple with the NewNameStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewNameStatus

`func (o *WhatsappPhoneNumber) SetNewNameStatus(v WhatsappPhoneNumberNameStatus)`

SetNewNameStatus sets NewNameStatus field to given value.

### HasNewNameStatus

`func (o *WhatsappPhoneNumber) HasNewNameStatus() bool`

HasNewNameStatus returns a boolean if a field has been set.

### GetDecision

`func (o *WhatsappPhoneNumber) GetDecision() WhatsappReviewDecision`

GetDecision returns the Decision field if non-nil, zero value otherwise.

### GetDecisionOk

`func (o *WhatsappPhoneNumber) GetDecisionOk() (*WhatsappReviewDecision, bool)`

GetDecisionOk returns a tuple with the Decision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecision

`func (o *WhatsappPhoneNumber) SetDecision(v WhatsappReviewDecision)`

SetDecision sets Decision field to given value.

### HasDecision

`func (o *WhatsappPhoneNumber) HasDecision() bool`

HasDecision returns a boolean if a field has been set.

### GetRequestedVerifiedName

`func (o *WhatsappPhoneNumber) GetRequestedVerifiedName() string`

GetRequestedVerifiedName returns the RequestedVerifiedName field if non-nil, zero value otherwise.

### GetRequestedVerifiedNameOk

`func (o *WhatsappPhoneNumber) GetRequestedVerifiedNameOk() (*string, bool)`

GetRequestedVerifiedNameOk returns a tuple with the RequestedVerifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedVerifiedName

`func (o *WhatsappPhoneNumber) SetRequestedVerifiedName(v string)`

SetRequestedVerifiedName sets RequestedVerifiedName field to given value.

### HasRequestedVerifiedName

`func (o *WhatsappPhoneNumber) HasRequestedVerifiedName() bool`

HasRequestedVerifiedName returns a boolean if a field has been set.

### GetRejectionReason

`func (o *WhatsappPhoneNumber) GetRejectionReason() string`

GetRejectionReason returns the RejectionReason field if non-nil, zero value otherwise.

### GetRejectionReasonOk

`func (o *WhatsappPhoneNumber) GetRejectionReasonOk() (*string, bool)`

GetRejectionReasonOk returns a tuple with the RejectionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionReason

`func (o *WhatsappPhoneNumber) SetRejectionReason(v string)`

SetRejectionReason sets RejectionReason field to given value.

### HasRejectionReason

`func (o *WhatsappPhoneNumber) HasRejectionReason() bool`

HasRejectionReason returns a boolean if a field has been set.

### GetQualityUpdateEvent

`func (o *WhatsappPhoneNumber) GetQualityUpdateEvent() WhatsappPhoneNumberQualityUpdateEventEnum`

GetQualityUpdateEvent returns the QualityUpdateEvent field if non-nil, zero value otherwise.

### GetQualityUpdateEventOk

`func (o *WhatsappPhoneNumber) GetQualityUpdateEventOk() (*WhatsappPhoneNumberQualityUpdateEventEnum, bool)`

GetQualityUpdateEventOk returns a tuple with the QualityUpdateEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityUpdateEvent

`func (o *WhatsappPhoneNumber) SetQualityUpdateEvent(v WhatsappPhoneNumberQualityUpdateEventEnum)`

SetQualityUpdateEvent sets QualityUpdateEvent field to given value.

### HasQualityUpdateEvent

`func (o *WhatsappPhoneNumber) HasQualityUpdateEvent() bool`

HasQualityUpdateEvent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
