# ContactNote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID for the contact note. |
**ContactId** | **string** | Unique ID of the contact that owns this note. |
**Content** | **string** | Note content. |
**OperatorId** | Pointer to **string** | ID of the actor who created the note. | [optional]
**UpdateOperatorId** | Pointer to **string** | ID of the actor who last updated the note. | [optional]
**CreateTime** | Pointer to **time.Time** | The time at which the note was created, formatted in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339). e.g., &#x60;2022-06-01T12:00:00.000Z&#x60;. | [optional]
**UpdateTime** | Pointer to **time.Time** | The time at which the note was last updated, formatted in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339). e.g., &#x60;2022-06-01T12:00:00.000Z&#x60;. | [optional]

## Methods

### NewContactNote

`func NewContactNote(id string, contactId string, content string, ) *ContactNote`

NewContactNote instantiates a new ContactNote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactNoteWithDefaults

`func NewContactNoteWithDefaults() *ContactNote`

NewContactNoteWithDefaults instantiates a new ContactNote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContactNote) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContactNote) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContactNote) SetId(v string)`

SetId sets Id field to given value.


### GetContactId

`func (o *ContactNote) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *ContactNote) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *ContactNote) SetContactId(v string)`

SetContactId sets ContactId field to given value.


### GetContent

`func (o *ContactNote) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ContactNote) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ContactNote) SetContent(v string)`

SetContent sets Content field to given value.


### GetOperatorId

`func (o *ContactNote) GetOperatorId() string`

GetOperatorId returns the OperatorId field if non-nil, zero value otherwise.

### GetOperatorIdOk

`func (o *ContactNote) GetOperatorIdOk() (*string, bool)`

GetOperatorIdOk returns a tuple with the OperatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorId

`func (o *ContactNote) SetOperatorId(v string)`

SetOperatorId sets OperatorId field to given value.

### HasOperatorId

`func (o *ContactNote) HasOperatorId() bool`

HasOperatorId returns a boolean if a field has been set.

### GetUpdateOperatorId

`func (o *ContactNote) GetUpdateOperatorId() string`

GetUpdateOperatorId returns the UpdateOperatorId field if non-nil, zero value otherwise.

### GetUpdateOperatorIdOk

`func (o *ContactNote) GetUpdateOperatorIdOk() (*string, bool)`

GetUpdateOperatorIdOk returns a tuple with the UpdateOperatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateOperatorId

`func (o *ContactNote) SetUpdateOperatorId(v string)`

SetUpdateOperatorId sets UpdateOperatorId field to given value.

### HasUpdateOperatorId

`func (o *ContactNote) HasUpdateOperatorId() bool`

HasUpdateOperatorId returns a boolean if a field has been set.

### GetCreateTime

`func (o *ContactNote) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *ContactNote) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *ContactNote) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *ContactNote) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *ContactNote) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *ContactNote) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *ContactNote) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *ContactNote) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
