# MetaBusinessAgentConversationTurn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TurnId** | **string** |  |
**ConversationId** | **string** | App-scoped opaque identifier for the conversation containing this turn. |
**MessageId** | Pointer to **string** | WhatsApp Cloud API message ID associated with this turn, when available. | [optional]
**E2eLatencyMs** | Pointer to **int64** |  | [optional]
**Timestamp** | Pointer to **int64** | Turn timestamp as Unix epoch milliseconds. | [optional]
**Steps** | [**[]MetaBusinessAgentConversationTurnStep**](MetaBusinessAgentConversationTurnStep.md) |  |

## Methods

### NewMetaBusinessAgentConversationTurn

`func NewMetaBusinessAgentConversationTurn(turnId string, conversationId string, steps []MetaBusinessAgentConversationTurnStep, ) *MetaBusinessAgentConversationTurn`

NewMetaBusinessAgentConversationTurn instantiates a new MetaBusinessAgentConversationTurn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaBusinessAgentConversationTurnWithDefaults

`func NewMetaBusinessAgentConversationTurnWithDefaults() *MetaBusinessAgentConversationTurn`

NewMetaBusinessAgentConversationTurnWithDefaults instantiates a new MetaBusinessAgentConversationTurn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTurnId

`func (o *MetaBusinessAgentConversationTurn) GetTurnId() string`

GetTurnId returns the TurnId field if non-nil, zero value otherwise.

### GetTurnIdOk

`func (o *MetaBusinessAgentConversationTurn) GetTurnIdOk() (*string, bool)`

GetTurnIdOk returns a tuple with the TurnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTurnId

`func (o *MetaBusinessAgentConversationTurn) SetTurnId(v string)`

SetTurnId sets TurnId field to given value.


### GetConversationId

`func (o *MetaBusinessAgentConversationTurn) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *MetaBusinessAgentConversationTurn) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *MetaBusinessAgentConversationTurn) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.


### GetMessageId

`func (o *MetaBusinessAgentConversationTurn) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *MetaBusinessAgentConversationTurn) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *MetaBusinessAgentConversationTurn) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *MetaBusinessAgentConversationTurn) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetE2eLatencyMs

`func (o *MetaBusinessAgentConversationTurn) GetE2eLatencyMs() int64`

GetE2eLatencyMs returns the E2eLatencyMs field if non-nil, zero value otherwise.

### GetE2eLatencyMsOk

`func (o *MetaBusinessAgentConversationTurn) GetE2eLatencyMsOk() (*int64, bool)`

GetE2eLatencyMsOk returns a tuple with the E2eLatencyMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE2eLatencyMs

`func (o *MetaBusinessAgentConversationTurn) SetE2eLatencyMs(v int64)`

SetE2eLatencyMs sets E2eLatencyMs field to given value.

### HasE2eLatencyMs

`func (o *MetaBusinessAgentConversationTurn) HasE2eLatencyMs() bool`

HasE2eLatencyMs returns a boolean if a field has been set.

### GetTimestamp

`func (o *MetaBusinessAgentConversationTurn) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MetaBusinessAgentConversationTurn) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MetaBusinessAgentConversationTurn) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *MetaBusinessAgentConversationTurn) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetSteps

`func (o *MetaBusinessAgentConversationTurn) GetSteps() []MetaBusinessAgentConversationTurnStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *MetaBusinessAgentConversationTurn) GetStepsOk() (*[]MetaBusinessAgentConversationTurnStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *MetaBusinessAgentConversationTurn) SetSteps(v []MetaBusinessAgentConversationTurnStep)`

SetSteps sets Steps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
