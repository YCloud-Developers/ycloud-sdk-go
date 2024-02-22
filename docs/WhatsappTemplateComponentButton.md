# WhatsappTemplateComponentButton

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**WhatsappTemplateComponentButtonType**](WhatsappTemplateComponentButtonType.md) |  | 
**Text** | Pointer to **string** | **Required for button type &#x60;PHONE_NUMBER&#x60; or &#x60;URL&#x60;.** Button text. For &#x60;CODE_CODE&#x60; buttons, the text is a pre-set value and cannot be customized. For &#x60;OTP&#x60; buttons, if omitted, the text will default to a pre-set value localized to the template&#39;s language. For example, &#x60;Copy Code&#x60; for English (US). If your template is using a one-tap autofill button and you supply this value, the authentication template message will display a copy code button with this text if we are unable to validate your [handshake](https://developers.facebook.com/docs/whatsapp/business-management-api/authentication-templates/autofill-button-authentication-templates#handshake). Maximum 25 characters. | [optional] 
**Url** | Pointer to **string** | **Required for button type &#x60;URL&#x60;.** URL of website. There can be at most 1 variable at the end of the URL. Example: &#x60;https://www.luckyshrub.com/shop?promo&#x3D;{{1}}&#x60;. 2000 characters maximum. | [optional] 
**PhoneNumber** | Pointer to **string** | **Required for button type &#x60;PHONE_NUMBER&#x60;.** Alphanumeric string. Business phone number to be (display phone number) called when the user taps the button. 20 characters maximum. | [optional] 
**OtpType** | Pointer to [**WhatsappTemplateComponentButtonOtpType**](WhatsappTemplateComponentButtonOtpType.md) |  | [optional] 
**AutofillText** | Pointer to **string** | **One-tap and zero-tap buttons only.** One-tap button text. Maximum 25 characters. | [optional] 
**PackageName** | Pointer to **string** | **One-tap and zero-tap buttons only.** Your Android app&#39;s package name. | [optional] 
**SignatureHash** | Pointer to **string** | **One-tap and zero-tap buttons only.** Your app signing key hash. See [App Signing Key Hash](https://developers.facebook.com/docs/whatsapp/business-management-api/authentication-templates/zero-tap-authentication-templates#app-signing-key-hash). | [optional] 
**ZeroTapTermsAccepted** | Pointer to **bool** | **Zero-tap buttons only.** Set to &#x60;true&#x60; to indicate that you understand that your use of zero-tap authentication is subject to the WhatsApp Business Terms of Service, and that it&#39;s your responsibility to ensure your customers expect that the code will be automatically filled in on their behalf when they choose to receive the zero-tap code through WhatsApp. If set to &#x60;false&#x60;, the template will not be created as you need to accept zero-tap terms before creating zero-tap enabled message templates. | [optional] 
**Example** | Pointer to **[]string** | Sample full URL for a &#x60;URL&#x60; button with a variable. | [optional] 

## Methods

### NewWhatsappTemplateComponentButton

`func NewWhatsappTemplateComponentButton(type_ WhatsappTemplateComponentButtonType, ) *WhatsappTemplateComponentButton`

NewWhatsappTemplateComponentButton instantiates a new WhatsappTemplateComponentButton object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappTemplateComponentButtonWithDefaults

`func NewWhatsappTemplateComponentButtonWithDefaults() *WhatsappTemplateComponentButton`

NewWhatsappTemplateComponentButtonWithDefaults instantiates a new WhatsappTemplateComponentButton object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WhatsappTemplateComponentButton) GetType() WhatsappTemplateComponentButtonType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WhatsappTemplateComponentButton) GetTypeOk() (*WhatsappTemplateComponentButtonType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WhatsappTemplateComponentButton) SetType(v WhatsappTemplateComponentButtonType)`

SetType sets Type field to given value.


### GetText

`func (o *WhatsappTemplateComponentButton) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *WhatsappTemplateComponentButton) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *WhatsappTemplateComponentButton) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *WhatsappTemplateComponentButton) HasText() bool`

HasText returns a boolean if a field has been set.

### GetUrl

`func (o *WhatsappTemplateComponentButton) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WhatsappTemplateComponentButton) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WhatsappTemplateComponentButton) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WhatsappTemplateComponentButton) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *WhatsappTemplateComponentButton) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *WhatsappTemplateComponentButton) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *WhatsappTemplateComponentButton) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *WhatsappTemplateComponentButton) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetOtpType

`func (o *WhatsappTemplateComponentButton) GetOtpType() WhatsappTemplateComponentButtonOtpType`

GetOtpType returns the OtpType field if non-nil, zero value otherwise.

### GetOtpTypeOk

`func (o *WhatsappTemplateComponentButton) GetOtpTypeOk() (*WhatsappTemplateComponentButtonOtpType, bool)`

GetOtpTypeOk returns a tuple with the OtpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpType

`func (o *WhatsappTemplateComponentButton) SetOtpType(v WhatsappTemplateComponentButtonOtpType)`

SetOtpType sets OtpType field to given value.

### HasOtpType

`func (o *WhatsappTemplateComponentButton) HasOtpType() bool`

HasOtpType returns a boolean if a field has been set.

### GetAutofillText

`func (o *WhatsappTemplateComponentButton) GetAutofillText() string`

GetAutofillText returns the AutofillText field if non-nil, zero value otherwise.

### GetAutofillTextOk

`func (o *WhatsappTemplateComponentButton) GetAutofillTextOk() (*string, bool)`

GetAutofillTextOk returns a tuple with the AutofillText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutofillText

`func (o *WhatsappTemplateComponentButton) SetAutofillText(v string)`

SetAutofillText sets AutofillText field to given value.

### HasAutofillText

`func (o *WhatsappTemplateComponentButton) HasAutofillText() bool`

HasAutofillText returns a boolean if a field has been set.

### GetPackageName

`func (o *WhatsappTemplateComponentButton) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *WhatsappTemplateComponentButton) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *WhatsappTemplateComponentButton) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.

### HasPackageName

`func (o *WhatsappTemplateComponentButton) HasPackageName() bool`

HasPackageName returns a boolean if a field has been set.

### GetSignatureHash

`func (o *WhatsappTemplateComponentButton) GetSignatureHash() string`

GetSignatureHash returns the SignatureHash field if non-nil, zero value otherwise.

### GetSignatureHashOk

`func (o *WhatsappTemplateComponentButton) GetSignatureHashOk() (*string, bool)`

GetSignatureHashOk returns a tuple with the SignatureHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureHash

`func (o *WhatsappTemplateComponentButton) SetSignatureHash(v string)`

SetSignatureHash sets SignatureHash field to given value.

### HasSignatureHash

`func (o *WhatsappTemplateComponentButton) HasSignatureHash() bool`

HasSignatureHash returns a boolean if a field has been set.

### GetZeroTapTermsAccepted

`func (o *WhatsappTemplateComponentButton) GetZeroTapTermsAccepted() bool`

GetZeroTapTermsAccepted returns the ZeroTapTermsAccepted field if non-nil, zero value otherwise.

### GetZeroTapTermsAcceptedOk

`func (o *WhatsappTemplateComponentButton) GetZeroTapTermsAcceptedOk() (*bool, bool)`

GetZeroTapTermsAcceptedOk returns a tuple with the ZeroTapTermsAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZeroTapTermsAccepted

`func (o *WhatsappTemplateComponentButton) SetZeroTapTermsAccepted(v bool)`

SetZeroTapTermsAccepted sets ZeroTapTermsAccepted field to given value.

### HasZeroTapTermsAccepted

`func (o *WhatsappTemplateComponentButton) HasZeroTapTermsAccepted() bool`

HasZeroTapTermsAccepted returns a boolean if a field has been set.

### GetExample

`func (o *WhatsappTemplateComponentButton) GetExample() []string`

GetExample returns the Example field if non-nil, zero value otherwise.

### GetExampleOk

`func (o *WhatsappTemplateComponentButton) GetExampleOk() (*[]string, bool)`

GetExampleOk returns a tuple with the Example field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExample

`func (o *WhatsappTemplateComponentButton) SetExample(v []string)`

SetExample sets Example field to given value.

### HasExample

`func (o *WhatsappTemplateComponentButton) HasExample() bool`

HasExample returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
