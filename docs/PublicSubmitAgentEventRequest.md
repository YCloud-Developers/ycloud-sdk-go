# PublicSubmitAgentEventRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | **string** | WhatsApp consumer E.164 phone number or ordinary business-scoped user ID (BSUID). Parent BSUIDs and Instagram recipients are not supported. |
**Event** | [**AgentEventDTO**](AgentEventDTO.md) |  |

## Methods

### NewPublicSubmitAgentEventRequest

`func NewPublicSubmitAgentEventRequest(to string, event AgentEventDTO, ) *PublicSubmitAgentEventRequest`

NewPublicSubmitAgentEventRequest instantiates a new PublicSubmitAgentEventRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicSubmitAgentEventRequestWithDefaults

`func NewPublicSubmitAgentEventRequestWithDefaults() *PublicSubmitAgentEventRequest`

NewPublicSubmitAgentEventRequestWithDefaults instantiates a new PublicSubmitAgentEventRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTo

`func (o *PublicSubmitAgentEventRequest) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *PublicSubmitAgentEventRequest) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *PublicSubmitAgentEventRequest) SetTo(v string)`

SetTo sets To field to given value.


### GetEvent

`func (o *PublicSubmitAgentEventRequest) GetEvent() AgentEventDTO`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *PublicSubmitAgentEventRequest) GetEventOk() (*AgentEventDTO, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *PublicSubmitAgentEventRequest) SetEvent(v AgentEventDTO)`

SetEvent sets Event field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
