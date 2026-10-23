# MetaBusinessAgentConversationTurnPaging

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursors** | [**MetaBusinessAgentConversationTurnCursors**](MetaBusinessAgentConversationTurnCursors.md) |  |
**Previous** | Pointer to **string** |  | [optional]
**Next** | Pointer to **string** | Present when another page is available, even if the current &#x60;data&#x60; array is empty. | [optional]

## Methods

### NewMetaBusinessAgentConversationTurnPaging

`func NewMetaBusinessAgentConversationTurnPaging(cursors MetaBusinessAgentConversationTurnCursors, ) *MetaBusinessAgentConversationTurnPaging`

NewMetaBusinessAgentConversationTurnPaging instantiates a new MetaBusinessAgentConversationTurnPaging object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaBusinessAgentConversationTurnPagingWithDefaults

`func NewMetaBusinessAgentConversationTurnPagingWithDefaults() *MetaBusinessAgentConversationTurnPaging`

NewMetaBusinessAgentConversationTurnPagingWithDefaults instantiates a new MetaBusinessAgentConversationTurnPaging object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursors

`func (o *MetaBusinessAgentConversationTurnPaging) GetCursors() MetaBusinessAgentConversationTurnCursors`

GetCursors returns the Cursors field if non-nil, zero value otherwise.

### GetCursorsOk

`func (o *MetaBusinessAgentConversationTurnPaging) GetCursorsOk() (*MetaBusinessAgentConversationTurnCursors, bool)`

GetCursorsOk returns a tuple with the Cursors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursors

`func (o *MetaBusinessAgentConversationTurnPaging) SetCursors(v MetaBusinessAgentConversationTurnCursors)`

SetCursors sets Cursors field to given value.


### GetPrevious

`func (o *MetaBusinessAgentConversationTurnPaging) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *MetaBusinessAgentConversationTurnPaging) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *MetaBusinessAgentConversationTurnPaging) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *MetaBusinessAgentConversationTurnPaging) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### GetNext

`func (o *MetaBusinessAgentConversationTurnPaging) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *MetaBusinessAgentConversationTurnPaging) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *MetaBusinessAgentConversationTurnPaging) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *MetaBusinessAgentConversationTurnPaging) HasNext() bool`

HasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
