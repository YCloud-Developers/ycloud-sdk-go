# PublicUpdateFaqRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Question** | **string** |  |
**Answer** | **string** |  |
**Metadata** | Pointer to **map[string]string** |  | [optional]

## Methods

### NewPublicUpdateFaqRequest

`func NewPublicUpdateFaqRequest(question string, answer string, ) *PublicUpdateFaqRequest`

NewPublicUpdateFaqRequest instantiates a new PublicUpdateFaqRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicUpdateFaqRequestWithDefaults

`func NewPublicUpdateFaqRequestWithDefaults() *PublicUpdateFaqRequest`

NewPublicUpdateFaqRequestWithDefaults instantiates a new PublicUpdateFaqRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuestion

`func (o *PublicUpdateFaqRequest) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *PublicUpdateFaqRequest) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *PublicUpdateFaqRequest) SetQuestion(v string)`

SetQuestion sets Question field to given value.


### GetAnswer

`func (o *PublicUpdateFaqRequest) GetAnswer() string`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *PublicUpdateFaqRequest) GetAnswerOk() (*string, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *PublicUpdateFaqRequest) SetAnswer(v string)`

SetAnswer sets Answer field to given value.


### GetMetadata

`func (o *PublicUpdateFaqRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PublicUpdateFaqRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PublicUpdateFaqRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PublicUpdateFaqRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
