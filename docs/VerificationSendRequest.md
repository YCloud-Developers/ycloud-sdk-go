# VerificationSendRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | [**VerificationChannel**](VerificationChannel.md) |  | 
**To** | **string** | The recipient&#39;s phone number or email address depending on &#x60;channel&#x60;. - Phone number: In [E.164](https://en.wikipedia.org/wiki/E.164) format. Applicable when &#x60;channel&#x60; is &#x60;sms&#x60; or &#x60;voice&#x60;. - Email address: For example, &#x60;tom@example.com&#x60;. Applicable when &#x60;channel&#x60; is &#x60;email_code&#x60;. | 
**Code** | Pointer to **string** | Verification code to be sent. This field is optional. If not provided, we will automatically generate a code. | [optional] 
**SenderId** | Pointer to **string** | [Sender ID](https://help.ycloud.com/en/articles/3080386) to be used. | [optional] 
**Signature** | Pointer to **string** | This parameter is only required for Chinese mainland SMS messages. You must specify an approved signature such as &#x60;Brand&#x60;. It will be added to the beginning of SMS body and wrapped with &#x60;【】&#x60;, e.g. &#x60;【Brand】Your verification code is 123456&#x60;. | [optional] 
**Language** | Pointer to **string** | [ISO 639 Language Code](https://www.iso.org/iso-639-language-codes.html). If not specified, language will be set as &#x60;en&#x60; by default. Notably, in certain countries or regions, language will be automatically set as the local language due to the regional restrictions. Applicable languages: &#x60;ar&#x60;: Arabic &#x60;de&#x60;: German &#x60;en&#x60;: English &#x60;es&#x60;: Spanish &#x60;fr&#x60;: French &#x60;id&#x60;: Indonesian &#x60;it&#x60;: Italian &#x60;pt_BR&#x60;: Portuguese &#x60;ru&#x60;: Russian &#x60;tr&#x60;: Turkish &#x60;vi&#x60;: Vietnamese &#x60;zh_CN&#x60;: Simplified Chinese &#x60;zh_HK&#x60;: Traditional Chinese | [optional] 
**ExternalId** | Pointer to **string** | A unique string to reference the object. This can be an order number or similar, and can be used to reconcile the object with your internal systems. If present, this value will also be attached to the &#x60;externalId&#x60; of message objects. | [optional] 

## Methods

### NewVerificationSendRequest

`func NewVerificationSendRequest(channel VerificationChannel, to string, ) *VerificationSendRequest`

NewVerificationSendRequest instantiates a new VerificationSendRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerificationSendRequestWithDefaults

`func NewVerificationSendRequestWithDefaults() *VerificationSendRequest`

NewVerificationSendRequestWithDefaults instantiates a new VerificationSendRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *VerificationSendRequest) GetChannel() VerificationChannel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *VerificationSendRequest) GetChannelOk() (*VerificationChannel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *VerificationSendRequest) SetChannel(v VerificationChannel)`

SetChannel sets Channel field to given value.


### GetTo

`func (o *VerificationSendRequest) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *VerificationSendRequest) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *VerificationSendRequest) SetTo(v string)`

SetTo sets To field to given value.


### GetCode

`func (o *VerificationSendRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *VerificationSendRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *VerificationSendRequest) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *VerificationSendRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetSenderId

`func (o *VerificationSendRequest) GetSenderId() string`

GetSenderId returns the SenderId field if non-nil, zero value otherwise.

### GetSenderIdOk

`func (o *VerificationSendRequest) GetSenderIdOk() (*string, bool)`

GetSenderIdOk returns a tuple with the SenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderId

`func (o *VerificationSendRequest) SetSenderId(v string)`

SetSenderId sets SenderId field to given value.

### HasSenderId

`func (o *VerificationSendRequest) HasSenderId() bool`

HasSenderId returns a boolean if a field has been set.

### GetSignature

`func (o *VerificationSendRequest) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *VerificationSendRequest) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *VerificationSendRequest) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *VerificationSendRequest) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetLanguage

`func (o *VerificationSendRequest) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *VerificationSendRequest) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *VerificationSendRequest) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *VerificationSendRequest) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetExternalId

`func (o *VerificationSendRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *VerificationSendRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *VerificationSendRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *VerificationSendRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
