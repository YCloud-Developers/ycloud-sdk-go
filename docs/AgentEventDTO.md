# AgentEventDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Stable business event type, up to 256 characters. |
**Description** | **string** | Plain-language meaning of the event, up to 1,024 characters. |
**Payload** | **string** | Opaque JSON serialized as a string, up to 4,096 characters. |

## Methods

### NewAgentEventDTO

`func NewAgentEventDTO(type_ string, description string, payload string, ) *AgentEventDTO`

NewAgentEventDTO instantiates a new AgentEventDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentEventDTOWithDefaults

`func NewAgentEventDTOWithDefaults() *AgentEventDTO`

NewAgentEventDTOWithDefaults instantiates a new AgentEventDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AgentEventDTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AgentEventDTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AgentEventDTO) SetType(v string)`

SetType sets Type field to given value.


### GetDescription

`func (o *AgentEventDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AgentEventDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AgentEventDTO) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPayload

`func (o *AgentEventDTO) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *AgentEventDTO) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *AgentEventDTO) SetPayload(v string)`

SetPayload sets Payload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
