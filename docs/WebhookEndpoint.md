# WebhookEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID for the object. | 
**Url** | Pointer to **string** | The URL of the webhook endpoint. | [optional] 
**EnabledEvents** | Pointer to **[]string** | The list of events to enable for this endpoint. | [optional] 
**Description** | Pointer to **string** | An optional description of what the webhook is used for. | [optional] 
**Status** | Pointer to [**WebhookEndpointStatus**](WebhookEndpointStatus.md) |  | [optional] 
**Secret** | Pointer to **string** | The endpoint&#39;s secret, used to generate webhook signatures. | [optional] 
**CreateTime** | Pointer to **time.Time** | The time at which this object was created, formatted in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339). e.g., &#x60;2022-06-01T12:00:00.000Z&#x60;. | [optional] 
**UpdateTime** | Pointer to **time.Time** | The time at which this object was updated, formatted in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339). e.g., &#x60;2022-06-01T12:00:00.000Z&#x60;. | [optional] 

## Methods

### NewWebhookEndpoint

`func NewWebhookEndpoint(id string, ) *WebhookEndpoint`

NewWebhookEndpoint instantiates a new WebhookEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookEndpointWithDefaults

`func NewWebhookEndpointWithDefaults() *WebhookEndpoint`

NewWebhookEndpointWithDefaults instantiates a new WebhookEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookEndpoint) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookEndpoint) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookEndpoint) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *WebhookEndpoint) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookEndpoint) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookEndpoint) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WebhookEndpoint) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetEnabledEvents

`func (o *WebhookEndpoint) GetEnabledEvents() []string`

GetEnabledEvents returns the EnabledEvents field if non-nil, zero value otherwise.

### GetEnabledEventsOk

`func (o *WebhookEndpoint) GetEnabledEventsOk() (*[]string, bool)`

GetEnabledEventsOk returns a tuple with the EnabledEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledEvents

`func (o *WebhookEndpoint) SetEnabledEvents(v []string)`

SetEnabledEvents sets EnabledEvents field to given value.

### HasEnabledEvents

`func (o *WebhookEndpoint) HasEnabledEvents() bool`

HasEnabledEvents returns a boolean if a field has been set.

### GetDescription

`func (o *WebhookEndpoint) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebhookEndpoint) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebhookEndpoint) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WebhookEndpoint) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *WebhookEndpoint) GetStatus() WebhookEndpointStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebhookEndpoint) GetStatusOk() (*WebhookEndpointStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebhookEndpoint) SetStatus(v WebhookEndpointStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WebhookEndpoint) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSecret

`func (o *WebhookEndpoint) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *WebhookEndpoint) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *WebhookEndpoint) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *WebhookEndpoint) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetCreateTime

`func (o *WebhookEndpoint) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *WebhookEndpoint) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *WebhookEndpoint) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *WebhookEndpoint) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *WebhookEndpoint) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *WebhookEndpoint) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *WebhookEndpoint) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *WebhookEndpoint) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
