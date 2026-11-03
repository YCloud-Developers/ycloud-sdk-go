# AgentAllowlistEntryDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Allowlist entry ID. |
**ConsumerPhoneNumber** | Pointer to **string** | Allowlisted consumer phone number in E.164 format. | [optional]
**ConsumerBsuid** | Pointer to **string** | Actual BSUID used to create this entry, if returned. | [optional]
**ConsumerInstagramUsername** | Pointer to **string** | Optional read-only compatibility field. Instagram management is not supported. | [optional]

## Methods

### NewAgentAllowlistEntryDTO

`func NewAgentAllowlistEntryDTO(id string, ) *AgentAllowlistEntryDTO`

NewAgentAllowlistEntryDTO instantiates a new AgentAllowlistEntryDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentAllowlistEntryDTOWithDefaults

`func NewAgentAllowlistEntryDTOWithDefaults() *AgentAllowlistEntryDTO`

NewAgentAllowlistEntryDTOWithDefaults instantiates a new AgentAllowlistEntryDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgentAllowlistEntryDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentAllowlistEntryDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentAllowlistEntryDTO) SetId(v string)`

SetId sets Id field to given value.


### GetConsumerPhoneNumber

`func (o *AgentAllowlistEntryDTO) GetConsumerPhoneNumber() string`

GetConsumerPhoneNumber returns the ConsumerPhoneNumber field if non-nil, zero value otherwise.

### GetConsumerPhoneNumberOk

`func (o *AgentAllowlistEntryDTO) GetConsumerPhoneNumberOk() (*string, bool)`

GetConsumerPhoneNumberOk returns a tuple with the ConsumerPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerPhoneNumber

`func (o *AgentAllowlistEntryDTO) SetConsumerPhoneNumber(v string)`

SetConsumerPhoneNumber sets ConsumerPhoneNumber field to given value.

### HasConsumerPhoneNumber

`func (o *AgentAllowlistEntryDTO) HasConsumerPhoneNumber() bool`

HasConsumerPhoneNumber returns a boolean if a field has been set.

### GetConsumerBsuid

`func (o *AgentAllowlistEntryDTO) GetConsumerBsuid() string`

GetConsumerBsuid returns the ConsumerBsuid field if non-nil, zero value otherwise.

### GetConsumerBsuidOk

`func (o *AgentAllowlistEntryDTO) GetConsumerBsuidOk() (*string, bool)`

GetConsumerBsuidOk returns a tuple with the ConsumerBsuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerBsuid

`func (o *AgentAllowlistEntryDTO) SetConsumerBsuid(v string)`

SetConsumerBsuid sets ConsumerBsuid field to given value.

### HasConsumerBsuid

`func (o *AgentAllowlistEntryDTO) HasConsumerBsuid() bool`

HasConsumerBsuid returns a boolean if a field has been set.

### GetConsumerInstagramUsername

`func (o *AgentAllowlistEntryDTO) GetConsumerInstagramUsername() string`

GetConsumerInstagramUsername returns the ConsumerInstagramUsername field if non-nil, zero value otherwise.

### GetConsumerInstagramUsernameOk

`func (o *AgentAllowlistEntryDTO) GetConsumerInstagramUsernameOk() (*string, bool)`

GetConsumerInstagramUsernameOk returns a tuple with the ConsumerInstagramUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerInstagramUsername

`func (o *AgentAllowlistEntryDTO) SetConsumerInstagramUsername(v string)`

SetConsumerInstagramUsername sets ConsumerInstagramUsername field to given value.

### HasConsumerInstagramUsername

`func (o *AgentAllowlistEntryDTO) HasConsumerInstagramUsername() bool`

HasConsumerInstagramUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
