# WhatsappGroupCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | **string** | The group subject. | 
**Description** | Pointer to **string** | The group description. | [optional] 
**JoinApprovalMode** | Pointer to [**WhatsappGroupJoinApprovalMode**](WhatsappGroupJoinApprovalMode.md) |  | [optional] 

## Methods

### NewWhatsappGroupCreateRequest

`func NewWhatsappGroupCreateRequest(subject string, ) *WhatsappGroupCreateRequest`

NewWhatsappGroupCreateRequest instantiates a new WhatsappGroupCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappGroupCreateRequestWithDefaults

`func NewWhatsappGroupCreateRequestWithDefaults() *WhatsappGroupCreateRequest`

NewWhatsappGroupCreateRequestWithDefaults instantiates a new WhatsappGroupCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *WhatsappGroupCreateRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *WhatsappGroupCreateRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *WhatsappGroupCreateRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetDescription

`func (o *WhatsappGroupCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WhatsappGroupCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WhatsappGroupCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WhatsappGroupCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJoinApprovalMode

`func (o *WhatsappGroupCreateRequest) GetJoinApprovalMode() WhatsappGroupJoinApprovalMode`

GetJoinApprovalMode returns the JoinApprovalMode field if non-nil, zero value otherwise.

### GetJoinApprovalModeOk

`func (o *WhatsappGroupCreateRequest) GetJoinApprovalModeOk() (*WhatsappGroupJoinApprovalMode, bool)`

GetJoinApprovalModeOk returns a tuple with the JoinApprovalMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinApprovalMode

`func (o *WhatsappGroupCreateRequest) SetJoinApprovalMode(v WhatsappGroupJoinApprovalMode)`

SetJoinApprovalMode sets JoinApprovalMode field to given value.

### HasJoinApprovalMode

`func (o *WhatsappGroupCreateRequest) HasJoinApprovalMode() bool`

HasJoinApprovalMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


