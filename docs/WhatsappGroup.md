# WhatsappGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | WhatsApp group ID. | [optional] 
**Subject** | Pointer to **string** | The group subject. | [optional] 
**Description** | Pointer to **string** | The group description. | [optional] 
**JoinApprovalMode** | Pointer to [**WhatsappGroupJoinApprovalMode**](WhatsappGroupJoinApprovalMode.md) |  | [optional] 
**Suspended** | Pointer to **bool** | Whether the group is suspended. | [optional] 
**CreationTimestamp** | Pointer to **int64** | Unix timestamp indicating when the group was created. | [optional] 
**TotalParticipantCount** | Pointer to **int32** | Total number of participants in the group. | [optional] 
**Participants** | Pointer to [**[]WhatsappGroupParticipant**](WhatsappGroupParticipant.md) | Group participants. | [optional] 

## Methods

### NewWhatsappGroup

`func NewWhatsappGroup() *WhatsappGroup`

NewWhatsappGroup instantiates a new WhatsappGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappGroupWithDefaults

`func NewWhatsappGroupWithDefaults() *WhatsappGroup`

NewWhatsappGroupWithDefaults instantiates a new WhatsappGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *WhatsappGroup) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *WhatsappGroup) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *WhatsappGroup) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *WhatsappGroup) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetSubject

`func (o *WhatsappGroup) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *WhatsappGroup) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *WhatsappGroup) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *WhatsappGroup) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetDescription

`func (o *WhatsappGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WhatsappGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WhatsappGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WhatsappGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJoinApprovalMode

`func (o *WhatsappGroup) GetJoinApprovalMode() WhatsappGroupJoinApprovalMode`

GetJoinApprovalMode returns the JoinApprovalMode field if non-nil, zero value otherwise.

### GetJoinApprovalModeOk

`func (o *WhatsappGroup) GetJoinApprovalModeOk() (*WhatsappGroupJoinApprovalMode, bool)`

GetJoinApprovalModeOk returns a tuple with the JoinApprovalMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinApprovalMode

`func (o *WhatsappGroup) SetJoinApprovalMode(v WhatsappGroupJoinApprovalMode)`

SetJoinApprovalMode sets JoinApprovalMode field to given value.

### HasJoinApprovalMode

`func (o *WhatsappGroup) HasJoinApprovalMode() bool`

HasJoinApprovalMode returns a boolean if a field has been set.

### GetSuspended

`func (o *WhatsappGroup) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *WhatsappGroup) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *WhatsappGroup) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *WhatsappGroup) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### GetCreationTimestamp

`func (o *WhatsappGroup) GetCreationTimestamp() int64`

GetCreationTimestamp returns the CreationTimestamp field if non-nil, zero value otherwise.

### GetCreationTimestampOk

`func (o *WhatsappGroup) GetCreationTimestampOk() (*int64, bool)`

GetCreationTimestampOk returns a tuple with the CreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimestamp

`func (o *WhatsappGroup) SetCreationTimestamp(v int64)`

SetCreationTimestamp sets CreationTimestamp field to given value.

### HasCreationTimestamp

`func (o *WhatsappGroup) HasCreationTimestamp() bool`

HasCreationTimestamp returns a boolean if a field has been set.

### GetTotalParticipantCount

`func (o *WhatsappGroup) GetTotalParticipantCount() int32`

GetTotalParticipantCount returns the TotalParticipantCount field if non-nil, zero value otherwise.

### GetTotalParticipantCountOk

`func (o *WhatsappGroup) GetTotalParticipantCountOk() (*int32, bool)`

GetTotalParticipantCountOk returns a tuple with the TotalParticipantCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalParticipantCount

`func (o *WhatsappGroup) SetTotalParticipantCount(v int32)`

SetTotalParticipantCount sets TotalParticipantCount field to given value.

### HasTotalParticipantCount

`func (o *WhatsappGroup) HasTotalParticipantCount() bool`

HasTotalParticipantCount returns a boolean if a field has been set.

### GetParticipants

`func (o *WhatsappGroup) GetParticipants() []WhatsappGroupParticipant`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *WhatsappGroup) GetParticipantsOk() (*[]WhatsappGroupParticipant, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *WhatsappGroup) SetParticipants(v []WhatsappGroupParticipant)`

SetParticipants sets Participants field to given value.

### HasParticipants

`func (o *WhatsappGroup) HasParticipants() bool`

HasParticipants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


