# WhatsappGroupListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | WhatsApp group ID. | [optional] 
**Subject** | Pointer to **string** | The group subject. | [optional] 
**CreatedAt** | Pointer to **int64** | Unix timestamp indicating when the group was created. | [optional] 

## Methods

### NewWhatsappGroupListItem

`func NewWhatsappGroupListItem() *WhatsappGroupListItem`

NewWhatsappGroupListItem instantiates a new WhatsappGroupListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappGroupListItemWithDefaults

`func NewWhatsappGroupListItemWithDefaults() *WhatsappGroupListItem`

NewWhatsappGroupListItemWithDefaults instantiates a new WhatsappGroupListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *WhatsappGroupListItem) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *WhatsappGroupListItem) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *WhatsappGroupListItem) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *WhatsappGroupListItem) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetSubject

`func (o *WhatsappGroupListItem) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *WhatsappGroupListItem) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *WhatsappGroupListItem) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *WhatsappGroupListItem) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetCreatedAt

`func (o *WhatsappGroupListItem) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WhatsappGroupListItem) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WhatsappGroupListItem) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WhatsappGroupListItem) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


