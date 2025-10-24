# WhatsappPhoneNumberProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**About** | Pointer to **string** | The business&#39;s **About** text. This text appears in the business&#39;s profile, beneath its profile image, phone number, and contact buttons. | [optional] 
**VerifiedName** | Pointer to **string** | The verified name | [optional] 
**NameStatus** | Pointer to [**WhatsappPhoneNumberNameStatus**](WhatsappPhoneNumberNameStatus.md) |  | [optional] 
**NewName** | Pointer to **string** | The modified name | [optional] 
**NewNameStatus** | Pointer to [**WhatsappPhoneNumberNameStatus**](WhatsappPhoneNumberNameStatus.md) |  | [optional] 
**Address** | Pointer to **string** | Address of the business. Character limit 256. | [optional] 
**Description** | Pointer to **string** | Description of the business. Character limit 512. | [optional] 
**Email** | Pointer to **string** | The contact email address (in valid email format) of the business. Character limit 128. | [optional] 
**ProfilePictureUrl** | Pointer to **string** | URL of the profile picture used to upload to Meta. | [optional] 
**Vertical** | Pointer to [**WhatsappPhoneNumberProfileVertical**](WhatsappPhoneNumberProfileVertical.md) |  | [optional] 
**Websites** | Pointer to **[]string** | The URLs associated with the business. For instance, a website, Facebook Page, or Instagram. You must include the http:// or https:// portion of the URL. There is a maximum of 2 websites with a maximum of 255 characters each. | [optional] 

## Methods

### NewWhatsappPhoneNumberProfile

`func NewWhatsappPhoneNumberProfile() *WhatsappPhoneNumberProfile`

NewWhatsappPhoneNumberProfile instantiates a new WhatsappPhoneNumberProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappPhoneNumberProfileWithDefaults

`func NewWhatsappPhoneNumberProfileWithDefaults() *WhatsappPhoneNumberProfile`

NewWhatsappPhoneNumberProfileWithDefaults instantiates a new WhatsappPhoneNumberProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbout

`func (o *WhatsappPhoneNumberProfile) GetAbout() string`

GetAbout returns the About field if non-nil, zero value otherwise.

### GetAboutOk

`func (o *WhatsappPhoneNumberProfile) GetAboutOk() (*string, bool)`

GetAboutOk returns a tuple with the About field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbout

`func (o *WhatsappPhoneNumberProfile) SetAbout(v string)`

SetAbout sets About field to given value.

### HasAbout

`func (o *WhatsappPhoneNumberProfile) HasAbout() bool`

HasAbout returns a boolean if a field has been set.

### GetVerifiedName

`func (o *WhatsappPhoneNumberProfile) GetVerifiedName() string`

GetVerifiedName returns the VerifiedName field if non-nil, zero value otherwise.

### GetVerifiedNameOk

`func (o *WhatsappPhoneNumberProfile) GetVerifiedNameOk() (*string, bool)`

GetVerifiedNameOk returns a tuple with the VerifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedName

`func (o *WhatsappPhoneNumberProfile) SetVerifiedName(v string)`

SetVerifiedName sets VerifiedName field to given value.

### HasVerifiedName

`func (o *WhatsappPhoneNumberProfile) HasVerifiedName() bool`

HasVerifiedName returns a boolean if a field has been set.

### GetNameStatus

`func (o *WhatsappPhoneNumberProfile) GetNameStatus() WhatsappPhoneNumberNameStatus`

GetNameStatus returns the NameStatus field if non-nil, zero value otherwise.

### GetNameStatusOk

`func (o *WhatsappPhoneNumberProfile) GetNameStatusOk() (*WhatsappPhoneNumberNameStatus, bool)`

GetNameStatusOk returns a tuple with the NameStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameStatus

`func (o *WhatsappPhoneNumberProfile) SetNameStatus(v WhatsappPhoneNumberNameStatus)`

SetNameStatus sets NameStatus field to given value.

### HasNameStatus

`func (o *WhatsappPhoneNumberProfile) HasNameStatus() bool`

HasNameStatus returns a boolean if a field has been set.

### GetNewName

`func (o *WhatsappPhoneNumberProfile) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *WhatsappPhoneNumberProfile) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *WhatsappPhoneNumberProfile) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *WhatsappPhoneNumberProfile) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetNewNameStatus

`func (o *WhatsappPhoneNumberProfile) GetNewNameStatus() WhatsappPhoneNumberNameStatus`

GetNewNameStatus returns the NewNameStatus field if non-nil, zero value otherwise.

### GetNewNameStatusOk

`func (o *WhatsappPhoneNumberProfile) GetNewNameStatusOk() (*WhatsappPhoneNumberNameStatus, bool)`

GetNewNameStatusOk returns a tuple with the NewNameStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewNameStatus

`func (o *WhatsappPhoneNumberProfile) SetNewNameStatus(v WhatsappPhoneNumberNameStatus)`

SetNewNameStatus sets NewNameStatus field to given value.

### HasNewNameStatus

`func (o *WhatsappPhoneNumberProfile) HasNewNameStatus() bool`

HasNewNameStatus returns a boolean if a field has been set.

### GetAddress

`func (o *WhatsappPhoneNumberProfile) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *WhatsappPhoneNumberProfile) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *WhatsappPhoneNumberProfile) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *WhatsappPhoneNumberProfile) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetDescription

`func (o *WhatsappPhoneNumberProfile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WhatsappPhoneNumberProfile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WhatsappPhoneNumberProfile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WhatsappPhoneNumberProfile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEmail

`func (o *WhatsappPhoneNumberProfile) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *WhatsappPhoneNumberProfile) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *WhatsappPhoneNumberProfile) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *WhatsappPhoneNumberProfile) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetProfilePictureUrl

`func (o *WhatsappPhoneNumberProfile) GetProfilePictureUrl() string`

GetProfilePictureUrl returns the ProfilePictureUrl field if non-nil, zero value otherwise.

### GetProfilePictureUrlOk

`func (o *WhatsappPhoneNumberProfile) GetProfilePictureUrlOk() (*string, bool)`

GetProfilePictureUrlOk returns a tuple with the ProfilePictureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePictureUrl

`func (o *WhatsappPhoneNumberProfile) SetProfilePictureUrl(v string)`

SetProfilePictureUrl sets ProfilePictureUrl field to given value.

### HasProfilePictureUrl

`func (o *WhatsappPhoneNumberProfile) HasProfilePictureUrl() bool`

HasProfilePictureUrl returns a boolean if a field has been set.

### GetVertical

`func (o *WhatsappPhoneNumberProfile) GetVertical() WhatsappPhoneNumberProfileVertical`

GetVertical returns the Vertical field if non-nil, zero value otherwise.

### GetVerticalOk

`func (o *WhatsappPhoneNumberProfile) GetVerticalOk() (*WhatsappPhoneNumberProfileVertical, bool)`

GetVerticalOk returns a tuple with the Vertical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVertical

`func (o *WhatsappPhoneNumberProfile) SetVertical(v WhatsappPhoneNumberProfileVertical)`

SetVertical sets Vertical field to given value.

### HasVertical

`func (o *WhatsappPhoneNumberProfile) HasVertical() bool`

HasVertical returns a boolean if a field has been set.

### GetWebsites

`func (o *WhatsappPhoneNumberProfile) GetWebsites() []string`

GetWebsites returns the Websites field if non-nil, zero value otherwise.

### GetWebsitesOk

`func (o *WhatsappPhoneNumberProfile) GetWebsitesOk() (*[]string, bool)`

GetWebsitesOk returns a tuple with the Websites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsites

`func (o *WhatsappPhoneNumberProfile) SetWebsites(v []string)`

SetWebsites sets Websites field to given value.

### HasWebsites

`func (o *WhatsappPhoneNumberProfile) HasWebsites() bool`

HasWebsites returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


