# AgentEventResultDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Initial event processing status returned by Meta. |
**AgentEventId** | Pointer to **string** | ID used to poll event processing status. | [optional]

## Methods

### NewAgentEventResultDTO

`func NewAgentEventResultDTO(status string, ) *AgentEventResultDTO`

NewAgentEventResultDTO instantiates a new AgentEventResultDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentEventResultDTOWithDefaults

`func NewAgentEventResultDTOWithDefaults() *AgentEventResultDTO`

NewAgentEventResultDTOWithDefaults instantiates a new AgentEventResultDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AgentEventResultDTO) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AgentEventResultDTO) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AgentEventResultDTO) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAgentEventId

`func (o *AgentEventResultDTO) GetAgentEventId() string`

GetAgentEventId returns the AgentEventId field if non-nil, zero value otherwise.

### GetAgentEventIdOk

`func (o *AgentEventResultDTO) GetAgentEventIdOk() (*string, bool)`

GetAgentEventIdOk returns a tuple with the AgentEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentEventId

`func (o *AgentEventResultDTO) SetAgentEventId(v string)`

SetAgentEventId sets AgentEventId field to given value.

### HasAgentEventId

`func (o *AgentEventResultDTO) HasAgentEventId() bool`

HasAgentEventId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
