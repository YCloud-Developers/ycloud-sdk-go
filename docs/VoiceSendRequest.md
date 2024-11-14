# VoiceSendRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | **string** | The recipient&#39;s phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format. | 
**VerificationCode** | **string** | The verification code to be sent, 4 to 6 digits. | 
**Language** | Pointer to **string** | [ISO 639 Language Code](https://www.iso.org/iso-639-language-codes.html). If not specified, language will be set as &#x60;en&#x60; by default. Notably, in certain countries or regions, language will be automatically set as the local language due to the regional restrictions. Applicable languages: &#x60;ar&#x60;: Arabic &#x60;de&#x60;: German &#x60;en&#x60;: English &#x60;es&#x60;: Spanish &#x60;fr&#x60;: French &#x60;id&#x60;: Indonesian &#x60;it&#x60;: Italian &#x60;pt&#x60;: Portuguese &#x60;ru&#x60;: Russian &#x60;tr&#x60;: Turkish &#x60;vi&#x60;: Vietnamese &#x60;zh&#x60;: Chinese | [optional] 
**ExternalId** | Pointer to **string** | A unique (recommended) string to reference the object. This can be an order number or similar, and can be used to reconcile the object with your internal systems. | [optional] 
**CallbackUrl** | Pointer to **string** | Delivery report URL. You can provide a URL, and we will push the updated status report to your server in time. e.g., https://httpbin.org/anything?tag&#x3D;api. Note: We recommend configuring Webhook Endpoints instead. | [optional] 

## Methods

### NewVoiceSendRequest

`func NewVoiceSendRequest(to string, verificationCode string, ) *VoiceSendRequest`

NewVoiceSendRequest instantiates a new VoiceSendRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoiceSendRequestWithDefaults

`func NewVoiceSendRequestWithDefaults() *VoiceSendRequest`

NewVoiceSendRequestWithDefaults instantiates a new VoiceSendRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTo

`func (o *VoiceSendRequest) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *VoiceSendRequest) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *VoiceSendRequest) SetTo(v string)`

SetTo sets To field to given value.


### GetVerificationCode

`func (o *VoiceSendRequest) GetVerificationCode() string`

GetVerificationCode returns the VerificationCode field if non-nil, zero value otherwise.

### GetVerificationCodeOk

`func (o *VoiceSendRequest) GetVerificationCodeOk() (*string, bool)`

GetVerificationCodeOk returns a tuple with the VerificationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationCode

`func (o *VoiceSendRequest) SetVerificationCode(v string)`

SetVerificationCode sets VerificationCode field to given value.


### GetLanguage

`func (o *VoiceSendRequest) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *VoiceSendRequest) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *VoiceSendRequest) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *VoiceSendRequest) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetExternalId

`func (o *VoiceSendRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *VoiceSendRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *VoiceSendRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *VoiceSendRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetCallbackUrl

`func (o *VoiceSendRequest) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *VoiceSendRequest) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *VoiceSendRequest) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *VoiceSendRequest) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


