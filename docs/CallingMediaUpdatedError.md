# CallingMediaUpdatedError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Stable failure code. |
**Retryable** | **bool** | Whether retrying the upstream media operation may succeed. |

## Methods

### NewCallingMediaUpdatedError

`func NewCallingMediaUpdatedError(code string, retryable bool, ) *CallingMediaUpdatedError`

NewCallingMediaUpdatedError instantiates a new CallingMediaUpdatedError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallingMediaUpdatedErrorWithDefaults

`func NewCallingMediaUpdatedErrorWithDefaults() *CallingMediaUpdatedError`

NewCallingMediaUpdatedErrorWithDefaults instantiates a new CallingMediaUpdatedError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CallingMediaUpdatedError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CallingMediaUpdatedError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CallingMediaUpdatedError) SetCode(v string)`

SetCode sets Code field to given value.


### GetRetryable

`func (o *CallingMediaUpdatedError) GetRetryable() bool`

GetRetryable returns the Retryable field if non-nil, zero value otherwise.

### GetRetryableOk

`func (o *CallingMediaUpdatedError) GetRetryableOk() (*bool, bool)`

GetRetryableOk returns a tuple with the Retryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryable

`func (o *CallingMediaUpdatedError) SetRetryable(v bool)`

SetRetryable sets Retryable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
