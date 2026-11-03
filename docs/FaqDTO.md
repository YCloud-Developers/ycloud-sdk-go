# FaqDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | FAQ resource ID. |
**Question** | **string** | Customer question written in natural language. |
**Answer** | **string** | Factual answer the agent should use. |
**CreatedAt** | Pointer to **int64** | Meta timestamp. When an update changes the FAQ ID, this resets to the update time. | [optional]
**Metadata** | Pointer to **map[string]string** | String metadata retained with the FAQ. | [optional]

## Methods

### NewFaqDTO

`func NewFaqDTO(id string, question string, answer string, ) *FaqDTO`

NewFaqDTO instantiates a new FaqDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFaqDTOWithDefaults

`func NewFaqDTOWithDefaults() *FaqDTO`

NewFaqDTOWithDefaults instantiates a new FaqDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FaqDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FaqDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FaqDTO) SetId(v string)`

SetId sets Id field to given value.


### GetQuestion

`func (o *FaqDTO) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *FaqDTO) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *FaqDTO) SetQuestion(v string)`

SetQuestion sets Question field to given value.


### GetAnswer

`func (o *FaqDTO) GetAnswer() string`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *FaqDTO) GetAnswerOk() (*string, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *FaqDTO) SetAnswer(v string)`

SetAnswer sets Answer field to given value.


### GetCreatedAt

`func (o *FaqDTO) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FaqDTO) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FaqDTO) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FaqDTO) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetMetadata

`func (o *FaqDTO) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FaqDTO) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FaqDTO) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FaqDTO) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
