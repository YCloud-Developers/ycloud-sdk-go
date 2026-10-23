# ContactNoteWriteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** | Note content. Leading and trailing whitespace is removed before validation and storage. |

## Methods

### NewContactNoteWriteRequest

`func NewContactNoteWriteRequest(content string, ) *ContactNoteWriteRequest`

NewContactNoteWriteRequest instantiates a new ContactNoteWriteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactNoteWriteRequestWithDefaults

`func NewContactNoteWriteRequestWithDefaults() *ContactNoteWriteRequest`

NewContactNoteWriteRequestWithDefaults instantiates a new ContactNoteWriteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *ContactNoteWriteRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ContactNoteWriteRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ContactNoteWriteRequest) SetContent(v string)`

SetContent sets Content field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
