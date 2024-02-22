# WebhookEndpointUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The URL of the webhook endpoint. | [optional] 
**EnabledEvents** | Pointer to [**[]EventType**](EventType.md) | The list of events to enable for this endpoint. | [optional] 
**Description** | Pointer to **string** | An optional description of what the webhook is used for. | [optional] 
**Status** | Pointer to [**WebhookEndpointStatus**](WebhookEndpointStatus.md) |  | [optional] 

## Methods

### NewWebhookEndpointUpdateRequest

`func NewWebhookEndpointUpdateRequest() *WebhookEndpointUpdateRequest`

NewWebhookEndpointUpdateRequest instantiates a new WebhookEndpointUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookEndpointUpdateRequestWithDefaults

`func NewWebhookEndpointUpdateRequestWithDefaults() *WebhookEndpointUpdateRequest`

NewWebhookEndpointUpdateRequestWithDefaults instantiates a new WebhookEndpointUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *WebhookEndpointUpdateRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookEndpointUpdateRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookEndpointUpdateRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WebhookEndpointUpdateRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetEnabledEvents

`func (o *WebhookEndpointUpdateRequest) GetEnabledEvents() []EventType`

GetEnabledEvents returns the EnabledEvents field if non-nil, zero value otherwise.

### GetEnabledEventsOk

`func (o *WebhookEndpointUpdateRequest) GetEnabledEventsOk() (*[]EventType, bool)`

GetEnabledEventsOk returns a tuple with the EnabledEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledEvents

`func (o *WebhookEndpointUpdateRequest) SetEnabledEvents(v []EventType)`

SetEnabledEvents sets EnabledEvents field to given value.

### HasEnabledEvents

`func (o *WebhookEndpointUpdateRequest) HasEnabledEvents() bool`

HasEnabledEvents returns a boolean if a field has been set.

### GetDescription

`func (o *WebhookEndpointUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebhookEndpointUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebhookEndpointUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WebhookEndpointUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *WebhookEndpointUpdateRequest) GetStatus() WebhookEndpointStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebhookEndpointUpdateRequest) GetStatusOk() (*WebhookEndpointStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebhookEndpointUpdateRequest) SetStatus(v WebhookEndpointStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WebhookEndpointUpdateRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
