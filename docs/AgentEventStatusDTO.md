# AgentEventStatusDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  |
**EventType** | **string** |  |
**ErrorMessage** | Pointer to **string** |  | [optional]
**SkippedReason** | Pointer to **string** |  | [optional]
**CreatedAt** | **string** |  |
**UpdatedAt** | **string** |  |
**ConversationId** | Pointer to **string** | Optional app-scoped opaque conversation ID, also returned by Conversation Turns. Omitted when unavailable; not an event ID or proof of message delivery. | [optional]

## Methods

### NewAgentEventStatusDTO

`func NewAgentEventStatusDTO(status string, eventType string, createdAt string, updatedAt string, ) *AgentEventStatusDTO`

NewAgentEventStatusDTO instantiates a new AgentEventStatusDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentEventStatusDTOWithDefaults

`func NewAgentEventStatusDTOWithDefaults() *AgentEventStatusDTO`

NewAgentEventStatusDTOWithDefaults instantiates a new AgentEventStatusDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AgentEventStatusDTO) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AgentEventStatusDTO) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AgentEventStatusDTO) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetEventType

`func (o *AgentEventStatusDTO) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *AgentEventStatusDTO) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *AgentEventStatusDTO) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetErrorMessage

`func (o *AgentEventStatusDTO) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *AgentEventStatusDTO) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *AgentEventStatusDTO) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *AgentEventStatusDTO) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetSkippedReason

`func (o *AgentEventStatusDTO) GetSkippedReason() string`

GetSkippedReason returns the SkippedReason field if non-nil, zero value otherwise.

### GetSkippedReasonOk

`func (o *AgentEventStatusDTO) GetSkippedReasonOk() (*string, bool)`

GetSkippedReasonOk returns a tuple with the SkippedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkippedReason

`func (o *AgentEventStatusDTO) SetSkippedReason(v string)`

SetSkippedReason sets SkippedReason field to given value.

### HasSkippedReason

`func (o *AgentEventStatusDTO) HasSkippedReason() bool`

HasSkippedReason returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AgentEventStatusDTO) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AgentEventStatusDTO) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AgentEventStatusDTO) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AgentEventStatusDTO) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AgentEventStatusDTO) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AgentEventStatusDTO) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetConversationId

`func (o *AgentEventStatusDTO) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *AgentEventStatusDTO) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *AgentEventStatusDTO) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.

### HasConversationId

`func (o *AgentEventStatusDTO) HasConversationId() bool`

HasConversationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
