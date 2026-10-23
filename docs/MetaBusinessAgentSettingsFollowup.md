# MetaBusinessAgentSettingsFollowup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  |
**FollowupIntervalInSeconds** | Pointer to **int32** | Follow-up interval in seconds when follow-up is enabled. | [optional]
**Message** | Pointer to **string** |  | [optional]

## Methods

### NewMetaBusinessAgentSettingsFollowup

`func NewMetaBusinessAgentSettingsFollowup(enabled bool, ) *MetaBusinessAgentSettingsFollowup`

NewMetaBusinessAgentSettingsFollowup instantiates a new MetaBusinessAgentSettingsFollowup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaBusinessAgentSettingsFollowupWithDefaults

`func NewMetaBusinessAgentSettingsFollowupWithDefaults() *MetaBusinessAgentSettingsFollowup`

NewMetaBusinessAgentSettingsFollowupWithDefaults instantiates a new MetaBusinessAgentSettingsFollowup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *MetaBusinessAgentSettingsFollowup) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MetaBusinessAgentSettingsFollowup) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MetaBusinessAgentSettingsFollowup) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetFollowupIntervalInSeconds

`func (o *MetaBusinessAgentSettingsFollowup) GetFollowupIntervalInSeconds() int32`

GetFollowupIntervalInSeconds returns the FollowupIntervalInSeconds field if non-nil, zero value otherwise.

### GetFollowupIntervalInSecondsOk

`func (o *MetaBusinessAgentSettingsFollowup) GetFollowupIntervalInSecondsOk() (*int32, bool)`

GetFollowupIntervalInSecondsOk returns a tuple with the FollowupIntervalInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowupIntervalInSeconds

`func (o *MetaBusinessAgentSettingsFollowup) SetFollowupIntervalInSeconds(v int32)`

SetFollowupIntervalInSeconds sets FollowupIntervalInSeconds field to given value.

### HasFollowupIntervalInSeconds

`func (o *MetaBusinessAgentSettingsFollowup) HasFollowupIntervalInSeconds() bool`

HasFollowupIntervalInSeconds returns a boolean if a field has been set.

### GetMessage

`func (o *MetaBusinessAgentSettingsFollowup) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MetaBusinessAgentSettingsFollowup) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MetaBusinessAgentSettingsFollowup) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MetaBusinessAgentSettingsFollowup) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
