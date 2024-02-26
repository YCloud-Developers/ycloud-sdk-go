# WebhookEndpointCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The URL of the webhook endpoint. | 
**EnabledEvents** | [**[]EventType**](EventType.md) | The list of events to enable for this endpoint. | 
**Description** | Pointer to **string** | An optional description of what the webhook is used for. | [optional] 
**Status** | Pointer to [**WebhookEndpointStatus**](WebhookEndpointStatus.md) |  | [optional] 

## Methods

### NewWebhookEndpointCreateRequest

`func NewWebhookEndpointCreateRequest(url string, enabledEvents []EventType, ) *WebhookEndpointCreateRequest`

NewWebhookEndpointCreateRequest instantiates a new WebhookEndpointCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookEndpointCreateRequestWithDefaults

`func NewWebhookEndpointCreateRequestWithDefaults() *WebhookEndpointCreateRequest`

NewWebhookEndpointCreateRequestWithDefaults instantiates a new WebhookEndpointCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *WebhookEndpointCreateRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookEndpointCreateRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookEndpointCreateRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetEnabledEvents

`func (o *WebhookEndpointCreateRequest) GetEnabledEvents() []EventType`

GetEnabledEvents returns the EnabledEvents field if non-nil, zero value otherwise.

### GetEnabledEventsOk

`func (o *WebhookEndpointCreateRequest) GetEnabledEventsOk() (*[]EventType, bool)`

GetEnabledEventsOk returns a tuple with the EnabledEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledEvents

`func (o *WebhookEndpointCreateRequest) SetEnabledEvents(v []EventType)`

SetEnabledEvents sets EnabledEvents field to given value.


### GetDescription

`func (o *WebhookEndpointCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebhookEndpointCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebhookEndpointCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WebhookEndpointCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *WebhookEndpointCreateRequest) GetStatus() WebhookEndpointStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebhookEndpointCreateRequest) GetStatusOk() (*WebhookEndpointStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebhookEndpointCreateRequest) SetStatus(v WebhookEndpointStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WebhookEndpointCreateRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


