# WhatsappPhoneNumberProfileUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**About** | Pointer to **string** | The business&#39;s **About** text. This text appears in the business&#39;s profile, beneath its profile image, phone number, and contact buttons. - String cannot be empty. - Strings must be between 1 and 139 characters. - Rendered emojis are supported however their unicode values are not. Emoji unicode values must be Java- or JavaScript-escape encoded. - Hyperlinks can be included but will not render as clickable links. - Markdown is not supported. | [optional] 
**Address** | Pointer to **string** | Address of the business. Character limit 256. | [optional] 
**Description** | Pointer to **string** | Description of the business. Character limit 512. | [optional] 
**Email** | Pointer to **string** | The contact email address (in valid email format) of the business. Character limit 128. | [optional] 
**ProfilePictureUrl** | Pointer to **string** | URL of the profile picture that was uploaded to Meta. | [optional] 
**Vertical** | Pointer to [**WhatsappPhoneNumberProfileVertical**](WhatsappPhoneNumberProfileVertical.md) |  | [optional] 
**Websites** | Pointer to **[]string** | The URLs associated with the business. For instance, a website, Facebook Page, or Instagram. You must include the http:// or https:// portion of the URL. There is a maximum of 2 websites with a maximum of 255 characters each. | [optional] 

## Methods

### NewWhatsappPhoneNumberProfileUpdateRequest

`func NewWhatsappPhoneNumberProfileUpdateRequest() *WhatsappPhoneNumberProfileUpdateRequest`

NewWhatsappPhoneNumberProfileUpdateRequest instantiates a new WhatsappPhoneNumberProfileUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappPhoneNumberProfileUpdateRequestWithDefaults

`func NewWhatsappPhoneNumberProfileUpdateRequestWithDefaults() *WhatsappPhoneNumberProfileUpdateRequest`

NewWhatsappPhoneNumberProfileUpdateRequestWithDefaults instantiates a new WhatsappPhoneNumberProfileUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbout

`func (o *WhatsappPhoneNumberProfileUpdateRequest) GetAbout() string`

GetAbout returns the About field if non-nil, zero value otherwise.

### GetAboutOk

`func (o *WhatsappPhoneNumberProfileUpdateRequest) GetAboutOk() (*string, bool)`

GetAboutOk returns a tuple with the About field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbout

`func (o *WhatsappPhoneNumberProfileUpdateRequest) SetAbout(v string)`

SetAbout sets About field to given value.

### HasAbout

`func (o *WhatsappPhoneNumberProfileUpdateRequest) HasAbout() bool`

HasAbout returns a boolean if a field has been set.

### GetAddress

`func (o *WhatsappPhoneNumberProfileUpdateRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *WhatsappPhoneNumberProfileUpdateRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *WhatsappPhoneNumberProfileUpdateRequest) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *WhatsappPhoneNumberProfileUpdateRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetDescription

`func (o *WhatsappPhoneNumberProfileUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WhatsappPhoneNumberProfileUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WhatsappPhoneNumberProfileUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WhatsappPhoneNumberProfileUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEmail

`func (o *WhatsappPhoneNumberProfileUpdateRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *WhatsappPhoneNumberProfileUpdateRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *WhatsappPhoneNumberProfileUpdateRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *WhatsappPhoneNumberProfileUpdateRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetProfilePictureUrl

`func (o *WhatsappPhoneNumberProfileUpdateRequest) GetProfilePictureUrl() string`

GetProfilePictureUrl returns the ProfilePictureUrl field if non-nil, zero value otherwise.

### GetProfilePictureUrlOk

`func (o *WhatsappPhoneNumberProfileUpdateRequest) GetProfilePictureUrlOk() (*string, bool)`

GetProfilePictureUrlOk returns a tuple with the ProfilePictureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePictureUrl

`func (o *WhatsappPhoneNumberProfileUpdateRequest) SetProfilePictureUrl(v string)`

SetProfilePictureUrl sets ProfilePictureUrl field to given value.

### HasProfilePictureUrl

`func (o *WhatsappPhoneNumberProfileUpdateRequest) HasProfilePictureUrl() bool`

HasProfilePictureUrl returns a boolean if a field has been set.

### GetVertical

`func (o *WhatsappPhoneNumberProfileUpdateRequest) GetVertical() WhatsappPhoneNumberProfileVertical`

GetVertical returns the Vertical field if non-nil, zero value otherwise.

### GetVerticalOk

`func (o *WhatsappPhoneNumberProfileUpdateRequest) GetVerticalOk() (*WhatsappPhoneNumberProfileVertical, bool)`

GetVerticalOk returns a tuple with the Vertical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVertical

`func (o *WhatsappPhoneNumberProfileUpdateRequest) SetVertical(v WhatsappPhoneNumberProfileVertical)`

SetVertical sets Vertical field to given value.

### HasVertical

`func (o *WhatsappPhoneNumberProfileUpdateRequest) HasVertical() bool`

HasVertical returns a boolean if a field has been set.

### GetWebsites

`func (o *WhatsappPhoneNumberProfileUpdateRequest) GetWebsites() []string`

GetWebsites returns the Websites field if non-nil, zero value otherwise.

### GetWebsitesOk

`func (o *WhatsappPhoneNumberProfileUpdateRequest) GetWebsitesOk() (*[]string, bool)`

GetWebsitesOk returns a tuple with the Websites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsites

`func (o *WhatsappPhoneNumberProfileUpdateRequest) SetWebsites(v []string)`

SetWebsites sets Websites field to given value.

### HasWebsites

`func (o *WhatsappPhoneNumberProfileUpdateRequest) HasWebsites() bool`

HasWebsites returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
