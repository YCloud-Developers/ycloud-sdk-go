# Voice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID for the object. | 
**To** | **string** | The recipient&#39;s phone number in E.164 format. | 
**VerificationCode** | Pointer to **string** | The verification code to be sent, 4 to 6 digits. | [optional] 
**Language** | Pointer to **string** | [ISO 639 Language Code](https://www.iso.org/iso-639-language-codes.html). | [optional] 
**RegionCode** | Pointer to **string** | [ISO 3166-1 alpha-2 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2). | [optional] 
**TotalSegments** | Pointer to **int32** | Number of message segments. It&#39;s always 1 for voice calls. | [optional] 
**TotalPrice** | Pointer to **float64** | Total price of this message. | [optional] 
**Currency** | Pointer to **string** | [ISO 4217 currency code](https://en.wikipedia.org/wiki/ISO_4217) | [optional] 
**Status** | Pointer to **string** | Delivery status. One of &#x60;accepted&#x60;, &#x60;sent&#x60;, &#x60;delivered&#x60;, &#x60;undelivered&#x60;. | [optional] 
**ErrorCode** | Pointer to **string** | Error code when the message is undeliverable. | [optional] 
**CreateTime** | Pointer to **time.Time** | The time at which this message was created, formatted in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339). e.g., &#x60;2022-03-01T12:00:00.000Z&#x60;. | [optional] 
**UpdateTime** | Pointer to **time.Time** | The time at which the delivery report for this message was updated, formatted in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339). e.g., &#x60;2022-03-01T12:00:00.000Z&#x60;. | [optional] 
**ExternalId** | Pointer to **string** | A unique string to reference the object. This can be an order number or similar, and can be used to reconcile the object with your internal systems. | [optional] 
**CallbackUrl** | Pointer to **string** | Delivery report URL. You can provide a URL, and we will push the updated status report to your server in time. e.g., https://httpbin.org/anything?tag&#x3D;api. Note: We recommend configuring Webhook Endpoints instead. | [optional] 

## Methods

### NewVoice

`func NewVoice(id string, to string, ) *Voice`

NewVoice instantiates a new Voice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoiceWithDefaults

`func NewVoiceWithDefaults() *Voice`

NewVoiceWithDefaults instantiates a new Voice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Voice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Voice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Voice) SetId(v string)`

SetId sets Id field to given value.


### GetTo

`func (o *Voice) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *Voice) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *Voice) SetTo(v string)`

SetTo sets To field to given value.


### GetVerificationCode

`func (o *Voice) GetVerificationCode() string`

GetVerificationCode returns the VerificationCode field if non-nil, zero value otherwise.

### GetVerificationCodeOk

`func (o *Voice) GetVerificationCodeOk() (*string, bool)`

GetVerificationCodeOk returns a tuple with the VerificationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationCode

`func (o *Voice) SetVerificationCode(v string)`

SetVerificationCode sets VerificationCode field to given value.

### HasVerificationCode

`func (o *Voice) HasVerificationCode() bool`

HasVerificationCode returns a boolean if a field has been set.

### GetLanguage

`func (o *Voice) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *Voice) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *Voice) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *Voice) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetRegionCode

`func (o *Voice) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *Voice) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *Voice) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *Voice) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetTotalSegments

`func (o *Voice) GetTotalSegments() int32`

GetTotalSegments returns the TotalSegments field if non-nil, zero value otherwise.

### GetTotalSegmentsOk

`func (o *Voice) GetTotalSegmentsOk() (*int32, bool)`

GetTotalSegmentsOk returns a tuple with the TotalSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSegments

`func (o *Voice) SetTotalSegments(v int32)`

SetTotalSegments sets TotalSegments field to given value.

### HasTotalSegments

`func (o *Voice) HasTotalSegments() bool`

HasTotalSegments returns a boolean if a field has been set.

### GetTotalPrice

`func (o *Voice) GetTotalPrice() float64`

GetTotalPrice returns the TotalPrice field if non-nil, zero value otherwise.

### GetTotalPriceOk

`func (o *Voice) GetTotalPriceOk() (*float64, bool)`

GetTotalPriceOk returns a tuple with the TotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrice

`func (o *Voice) SetTotalPrice(v float64)`

SetTotalPrice sets TotalPrice field to given value.

### HasTotalPrice

`func (o *Voice) HasTotalPrice() bool`

HasTotalPrice returns a boolean if a field has been set.

### GetCurrency

`func (o *Voice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Voice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Voice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Voice) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetStatus

`func (o *Voice) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Voice) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Voice) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Voice) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrorCode

`func (o *Voice) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *Voice) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *Voice) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *Voice) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetCreateTime

`func (o *Voice) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *Voice) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *Voice) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *Voice) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *Voice) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *Voice) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *Voice) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *Voice) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetExternalId

`func (o *Voice) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Voice) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Voice) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *Voice) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetCallbackUrl

`func (o *Voice) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *Voice) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *Voice) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *Voice) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


