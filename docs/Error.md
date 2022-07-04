# Error

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int32** | HTTP status code, [RFC 7231, Section 6](https://datatracker.ietf.org/doc/html/rfc7231#section-6). | 
**Code** | **string** | One of a server-defined set of error codes, which could be handled programmatically. | 
**Message** | Pointer to **string** | A human-readable representation of the error. It is intended as an aid to developers and is not suitable for exposure to end users. | [optional] 
**Target** | Pointer to **string** | The target of the error. | [optional] 
**DocUrl** | Pointer to **string** | A URL to more information about the error. | [optional] 

## Methods

### NewError

`func NewError(status int32, code string, ) *Error`

NewError instantiates a new Error object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorWithDefaults

`func NewErrorWithDefaults() *Error`

NewErrorWithDefaults instantiates a new Error object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Error) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Error) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Error) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetCode

`func (o *Error) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Error) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Error) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *Error) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Error) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Error) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Error) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTarget

`func (o *Error) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *Error) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *Error) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *Error) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetDocUrl

`func (o *Error) GetDocUrl() string`

GetDocUrl returns the DocUrl field if non-nil, zero value otherwise.

### GetDocUrlOk

`func (o *Error) GetDocUrlOk() (*string, bool)`

GetDocUrlOk returns a tuple with the DocUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocUrl

`func (o *Error) SetDocUrl(v string)`

SetDocUrl sets DocUrl field to given value.

### HasDocUrl

`func (o *Error) HasDocUrl() bool`

HasDocUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


