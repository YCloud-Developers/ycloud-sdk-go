# MetaBusinessAgentSettingsHandoff

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  |
**MessageSelection** | Pointer to **string** | &#x60;message&#x60; is required when this value is &#x60;CUSTOM&#x60;. | [optional]
**Message** | Pointer to **string** |  | [optional]

## Methods

### NewMetaBusinessAgentSettingsHandoff

`func NewMetaBusinessAgentSettingsHandoff(enabled bool, ) *MetaBusinessAgentSettingsHandoff`

NewMetaBusinessAgentSettingsHandoff instantiates a new MetaBusinessAgentSettingsHandoff object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaBusinessAgentSettingsHandoffWithDefaults

`func NewMetaBusinessAgentSettingsHandoffWithDefaults() *MetaBusinessAgentSettingsHandoff`

NewMetaBusinessAgentSettingsHandoffWithDefaults instantiates a new MetaBusinessAgentSettingsHandoff object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *MetaBusinessAgentSettingsHandoff) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MetaBusinessAgentSettingsHandoff) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MetaBusinessAgentSettingsHandoff) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMessageSelection

`func (o *MetaBusinessAgentSettingsHandoff) GetMessageSelection() string`

GetMessageSelection returns the MessageSelection field if non-nil, zero value otherwise.

### GetMessageSelectionOk

`func (o *MetaBusinessAgentSettingsHandoff) GetMessageSelectionOk() (*string, bool)`

GetMessageSelectionOk returns a tuple with the MessageSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageSelection

`func (o *MetaBusinessAgentSettingsHandoff) SetMessageSelection(v string)`

SetMessageSelection sets MessageSelection field to given value.

### HasMessageSelection

`func (o *MetaBusinessAgentSettingsHandoff) HasMessageSelection() bool`

HasMessageSelection returns a boolean if a field has been set.

### GetMessage

`func (o *MetaBusinessAgentSettingsHandoff) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MetaBusinessAgentSettingsHandoff) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MetaBusinessAgentSettingsHandoff) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MetaBusinessAgentSettingsHandoff) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
