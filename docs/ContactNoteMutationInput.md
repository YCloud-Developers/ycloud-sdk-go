# ContactNoteMutationInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Existing note ID. Omit to create a new note. | [optional]
**Content** | **string** | Note content. Leading and trailing whitespace is removed before validation and storage. |

## Methods

### NewContactNoteMutationInput

`func NewContactNoteMutationInput(content string, ) *ContactNoteMutationInput`

NewContactNoteMutationInput instantiates a new ContactNoteMutationInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactNoteMutationInputWithDefaults

`func NewContactNoteMutationInputWithDefaults() *ContactNoteMutationInput`

NewContactNoteMutationInputWithDefaults instantiates a new ContactNoteMutationInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContactNoteMutationInput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContactNoteMutationInput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContactNoteMutationInput) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ContactNoteMutationInput) HasId() bool`

HasId returns a boolean if a field has been set.

### GetContent

`func (o *ContactNoteMutationInput) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ContactNoteMutationInput) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ContactNoteMutationInput) SetContent(v string)`

SetContent sets Content field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
