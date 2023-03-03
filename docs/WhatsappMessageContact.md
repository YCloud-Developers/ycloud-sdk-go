# WhatsappMessageContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | Pointer to [**[]WhatsappMessageContactAddress**](WhatsappMessageContactAddress.md) |  | [optional] 
**Birthday** | Pointer to **string** | &#x60;YYYY-MM-DD&#x60; formatted string. | [optional] 
**Emails** | Pointer to [**[]WhatsappMessageContactEmail**](WhatsappMessageContactEmail.md) |  | [optional] 
**Name** | [**WhatsappMessageContactName**](WhatsappMessageContactName.md) |  | 
**Org** | Pointer to [**WhatsappMessageContactOrg**](WhatsappMessageContactOrg.md) |  | [optional] 
**Phones** | Pointer to [**[]WhatsappMessageContactPhone**](WhatsappMessageContactPhone.md) | Contact phone number(s) formatted as a phone object. | [optional] 
**Urls** | Pointer to [**[]WhatsappMessageContactUrl**](WhatsappMessageContactUrl.md) | Contact URL(s) formatted as a urls object. | [optional] 

## Methods

### NewWhatsappMessageContact

`func NewWhatsappMessageContact(name WhatsappMessageContactName, ) *WhatsappMessageContact`

NewWhatsappMessageContact instantiates a new WhatsappMessageContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappMessageContactWithDefaults

`func NewWhatsappMessageContactWithDefaults() *WhatsappMessageContact`

NewWhatsappMessageContactWithDefaults instantiates a new WhatsappMessageContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *WhatsappMessageContact) GetAddresses() []WhatsappMessageContactAddress`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *WhatsappMessageContact) GetAddressesOk() (*[]WhatsappMessageContactAddress, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *WhatsappMessageContact) SetAddresses(v []WhatsappMessageContactAddress)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *WhatsappMessageContact) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetBirthday

`func (o *WhatsappMessageContact) GetBirthday() string`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *WhatsappMessageContact) GetBirthdayOk() (*string, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *WhatsappMessageContact) SetBirthday(v string)`

SetBirthday sets Birthday field to given value.

### HasBirthday

`func (o *WhatsappMessageContact) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### GetEmails

`func (o *WhatsappMessageContact) GetEmails() []WhatsappMessageContactEmail`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *WhatsappMessageContact) GetEmailsOk() (*[]WhatsappMessageContactEmail, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *WhatsappMessageContact) SetEmails(v []WhatsappMessageContactEmail)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *WhatsappMessageContact) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetName

`func (o *WhatsappMessageContact) GetName() WhatsappMessageContactName`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WhatsappMessageContact) GetNameOk() (*WhatsappMessageContactName, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WhatsappMessageContact) SetName(v WhatsappMessageContactName)`

SetName sets Name field to given value.


### GetOrg

`func (o *WhatsappMessageContact) GetOrg() WhatsappMessageContactOrg`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *WhatsappMessageContact) GetOrgOk() (*WhatsappMessageContactOrg, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *WhatsappMessageContact) SetOrg(v WhatsappMessageContactOrg)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *WhatsappMessageContact) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetPhones

`func (o *WhatsappMessageContact) GetPhones() []WhatsappMessageContactPhone`

GetPhones returns the Phones field if non-nil, zero value otherwise.

### GetPhonesOk

`func (o *WhatsappMessageContact) GetPhonesOk() (*[]WhatsappMessageContactPhone, bool)`

GetPhonesOk returns a tuple with the Phones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhones

`func (o *WhatsappMessageContact) SetPhones(v []WhatsappMessageContactPhone)`

SetPhones sets Phones field to given value.

### HasPhones

`func (o *WhatsappMessageContact) HasPhones() bool`

HasPhones returns a boolean if a field has been set.

### GetUrls

`func (o *WhatsappMessageContact) GetUrls() []WhatsappMessageContactUrl`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *WhatsappMessageContact) GetUrlsOk() (*[]WhatsappMessageContactUrl, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *WhatsappMessageContact) SetUrls(v []WhatsappMessageContactUrl)`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *WhatsappMessageContact) HasUrls() bool`

HasUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


