# WhatsappMessageInteractiveActionParametersFlowActionPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Screen** | Pointer to **string** | The ID of the screen displayed first. It needs to be an **entry** screen. | [optional] 
**Data** | Pointer to **map[string]map[string]interface{}** | Optional input data for the first screen of the Flow. If provided, this must be a non-empty object. | [optional] 

## Methods

### NewWhatsappMessageInteractiveActionParametersFlowActionPayload

`func NewWhatsappMessageInteractiveActionParametersFlowActionPayload() *WhatsappMessageInteractiveActionParametersFlowActionPayload`

NewWhatsappMessageInteractiveActionParametersFlowActionPayload instantiates a new WhatsappMessageInteractiveActionParametersFlowActionPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappMessageInteractiveActionParametersFlowActionPayloadWithDefaults

`func NewWhatsappMessageInteractiveActionParametersFlowActionPayloadWithDefaults() *WhatsappMessageInteractiveActionParametersFlowActionPayload`

NewWhatsappMessageInteractiveActionParametersFlowActionPayloadWithDefaults instantiates a new WhatsappMessageInteractiveActionParametersFlowActionPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScreen

`func (o *WhatsappMessageInteractiveActionParametersFlowActionPayload) GetScreen() string`

GetScreen returns the Screen field if non-nil, zero value otherwise.

### GetScreenOk

`func (o *WhatsappMessageInteractiveActionParametersFlowActionPayload) GetScreenOk() (*string, bool)`

GetScreenOk returns a tuple with the Screen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreen

`func (o *WhatsappMessageInteractiveActionParametersFlowActionPayload) SetScreen(v string)`

SetScreen sets Screen field to given value.

### HasScreen

`func (o *WhatsappMessageInteractiveActionParametersFlowActionPayload) HasScreen() bool`

HasScreen returns a boolean if a field has been set.

### GetData

`func (o *WhatsappMessageInteractiveActionParametersFlowActionPayload) GetData() map[string]map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WhatsappMessageInteractiveActionParametersFlowActionPayload) GetDataOk() (*map[string]map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WhatsappMessageInteractiveActionParametersFlowActionPayload) SetData(v map[string]map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *WhatsappMessageInteractiveActionParametersFlowActionPayload) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


