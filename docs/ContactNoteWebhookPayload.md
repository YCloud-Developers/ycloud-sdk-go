# ContactNoteWebhookPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique 24-character ObjectId for the contact note. IDs remain unchanged after storage migration. |
**ContactId** | **string** | Unique ID of the contact that owns this note. |
**Username** | **NullableString** | Username of the contact that owns this note, without the leading &#x60;@&#x60;. |
**PhoneNumber** | **NullableString** | Phone number of the contact that owns this note in E.164 format. |
**Content** | **string** | Note content. |
**CreateTime** | **time.Time** | The time at which the note was created, formatted in RFC 3339. |
**UpdateTime** | **time.Time** | The time at which the note was last updated, formatted in RFC 3339. |

## Methods

### NewContactNoteWebhookPayload

`func NewContactNoteWebhookPayload(id string, contactId string, username NullableString, phoneNumber NullableString, content string, createTime time.Time, updateTime time.Time, ) *ContactNoteWebhookPayload`

NewContactNoteWebhookPayload instantiates a new ContactNoteWebhookPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactNoteWebhookPayloadWithDefaults

`func NewContactNoteWebhookPayloadWithDefaults() *ContactNoteWebhookPayload`

NewContactNoteWebhookPayloadWithDefaults instantiates a new ContactNoteWebhookPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContactNoteWebhookPayload) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContactNoteWebhookPayload) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContactNoteWebhookPayload) SetId(v string)`

SetId sets Id field to given value.


### GetContactId

`func (o *ContactNoteWebhookPayload) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *ContactNoteWebhookPayload) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *ContactNoteWebhookPayload) SetContactId(v string)`

SetContactId sets ContactId field to given value.


### GetUsername

`func (o *ContactNoteWebhookPayload) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ContactNoteWebhookPayload) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ContactNoteWebhookPayload) SetUsername(v string)`

SetUsername sets Username field to given value.


### SetUsernameNil

`func (o *ContactNoteWebhookPayload) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *ContactNoteWebhookPayload) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPhoneNumber

`func (o *ContactNoteWebhookPayload) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *ContactNoteWebhookPayload) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *ContactNoteWebhookPayload) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.


### SetPhoneNumberNil

`func (o *ContactNoteWebhookPayload) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *ContactNoteWebhookPayload) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetContent

`func (o *ContactNoteWebhookPayload) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ContactNoteWebhookPayload) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ContactNoteWebhookPayload) SetContent(v string)`

SetContent sets Content field to given value.


### GetCreateTime

`func (o *ContactNoteWebhookPayload) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *ContactNoteWebhookPayload) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *ContactNoteWebhookPayload) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.


### GetUpdateTime

`func (o *ContactNoteWebhookPayload) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *ContactNoteWebhookPayload) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *ContactNoteWebhookPayload) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
