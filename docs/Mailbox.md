# Mailbox

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the mailbox. | [optional] 
**Address** | Pointer to **string** | Address of the mailbox. | [optional] 

## Methods

### NewMailbox

`func NewMailbox() *Mailbox`

NewMailbox instantiates a new Mailbox object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMailboxWithDefaults

`func NewMailboxWithDefaults() *Mailbox`

NewMailboxWithDefaults instantiates a new Mailbox object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Mailbox) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Mailbox) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Mailbox) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Mailbox) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddress

`func (o *Mailbox) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Mailbox) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Mailbox) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Mailbox) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
