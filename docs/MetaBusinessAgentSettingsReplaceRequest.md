# MetaBusinessAgentSettingsReplaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentId** | Pointer to **string** |  | [optional]
**Rollout** | Pointer to [**MetaBusinessAgentSettingsRollout**](MetaBusinessAgentSettingsRollout.md) |  | [optional]
**Handoff** | Pointer to [**MetaBusinessAgentSettingsHandoff**](MetaBusinessAgentSettingsHandoff.md) |  | [optional]
**Followup** | Pointer to [**MetaBusinessAgentSettingsFollowup**](MetaBusinessAgentSettingsFollowup.md) |  | [optional]
**AiAudience** | Pointer to **string** |  | [optional]
**NeverSayPhrases** | Pointer to **[]string** | Omit to preserve, send an empty array to clear, or send the complete replacement list. | [optional]

## Methods

### NewMetaBusinessAgentSettingsReplaceRequest

`func NewMetaBusinessAgentSettingsReplaceRequest() *MetaBusinessAgentSettingsReplaceRequest`

NewMetaBusinessAgentSettingsReplaceRequest instantiates a new MetaBusinessAgentSettingsReplaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaBusinessAgentSettingsReplaceRequestWithDefaults

`func NewMetaBusinessAgentSettingsReplaceRequestWithDefaults() *MetaBusinessAgentSettingsReplaceRequest`

NewMetaBusinessAgentSettingsReplaceRequestWithDefaults instantiates a new MetaBusinessAgentSettingsReplaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentId

`func (o *MetaBusinessAgentSettingsReplaceRequest) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *MetaBusinessAgentSettingsReplaceRequest) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *MetaBusinessAgentSettingsReplaceRequest) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *MetaBusinessAgentSettingsReplaceRequest) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetRollout

`func (o *MetaBusinessAgentSettingsReplaceRequest) GetRollout() MetaBusinessAgentSettingsRollout`

GetRollout returns the Rollout field if non-nil, zero value otherwise.

### GetRolloutOk

`func (o *MetaBusinessAgentSettingsReplaceRequest) GetRolloutOk() (*MetaBusinessAgentSettingsRollout, bool)`

GetRolloutOk returns a tuple with the Rollout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollout

`func (o *MetaBusinessAgentSettingsReplaceRequest) SetRollout(v MetaBusinessAgentSettingsRollout)`

SetRollout sets Rollout field to given value.

### HasRollout

`func (o *MetaBusinessAgentSettingsReplaceRequest) HasRollout() bool`

HasRollout returns a boolean if a field has been set.

### GetHandoff

`func (o *MetaBusinessAgentSettingsReplaceRequest) GetHandoff() MetaBusinessAgentSettingsHandoff`

GetHandoff returns the Handoff field if non-nil, zero value otherwise.

### GetHandoffOk

`func (o *MetaBusinessAgentSettingsReplaceRequest) GetHandoffOk() (*MetaBusinessAgentSettingsHandoff, bool)`

GetHandoffOk returns a tuple with the Handoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandoff

`func (o *MetaBusinessAgentSettingsReplaceRequest) SetHandoff(v MetaBusinessAgentSettingsHandoff)`

SetHandoff sets Handoff field to given value.

### HasHandoff

`func (o *MetaBusinessAgentSettingsReplaceRequest) HasHandoff() bool`

HasHandoff returns a boolean if a field has been set.

### GetFollowup

`func (o *MetaBusinessAgentSettingsReplaceRequest) GetFollowup() MetaBusinessAgentSettingsFollowup`

GetFollowup returns the Followup field if non-nil, zero value otherwise.

### GetFollowupOk

`func (o *MetaBusinessAgentSettingsReplaceRequest) GetFollowupOk() (*MetaBusinessAgentSettingsFollowup, bool)`

GetFollowupOk returns a tuple with the Followup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowup

`func (o *MetaBusinessAgentSettingsReplaceRequest) SetFollowup(v MetaBusinessAgentSettingsFollowup)`

SetFollowup sets Followup field to given value.

### HasFollowup

`func (o *MetaBusinessAgentSettingsReplaceRequest) HasFollowup() bool`

HasFollowup returns a boolean if a field has been set.

### GetAiAudience

`func (o *MetaBusinessAgentSettingsReplaceRequest) GetAiAudience() string`

GetAiAudience returns the AiAudience field if non-nil, zero value otherwise.

### GetAiAudienceOk

`func (o *MetaBusinessAgentSettingsReplaceRequest) GetAiAudienceOk() (*string, bool)`

GetAiAudienceOk returns a tuple with the AiAudience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiAudience

`func (o *MetaBusinessAgentSettingsReplaceRequest) SetAiAudience(v string)`

SetAiAudience sets AiAudience field to given value.

### HasAiAudience

`func (o *MetaBusinessAgentSettingsReplaceRequest) HasAiAudience() bool`

HasAiAudience returns a boolean if a field has been set.

### GetNeverSayPhrases

`func (o *MetaBusinessAgentSettingsReplaceRequest) GetNeverSayPhrases() []string`

GetNeverSayPhrases returns the NeverSayPhrases field if non-nil, zero value otherwise.

### GetNeverSayPhrasesOk

`func (o *MetaBusinessAgentSettingsReplaceRequest) GetNeverSayPhrasesOk() (*[]string, bool)`

GetNeverSayPhrasesOk returns a tuple with the NeverSayPhrases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeverSayPhrases

`func (o *MetaBusinessAgentSettingsReplaceRequest) SetNeverSayPhrases(v []string)`

SetNeverSayPhrases sets NeverSayPhrases field to given value.

### HasNeverSayPhrases

`func (o *MetaBusinessAgentSettingsReplaceRequest) HasNeverSayPhrases() bool`

HasNeverSayPhrases returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
