# VerificationCheckRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VerificationId** | Pointer to **string** | ID of the verification to be checked. One of &#x60;verificationId&#x60; or &#x60;to&#x60; is required. | [optional] 
**To** | Pointer to **string** | The recipient&#39;s phone number or email address. One of &#x60;verificationId&#x60; or &#x60;to&#x60; is required. | [optional] 
**Code** | Pointer to **string** | The verification to be checked. | [optional] 

## Methods

### NewVerificationCheckRequest

`func NewVerificationCheckRequest() *VerificationCheckRequest`

NewVerificationCheckRequest instantiates a new VerificationCheckRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerificationCheckRequestWithDefaults

`func NewVerificationCheckRequestWithDefaults() *VerificationCheckRequest`

NewVerificationCheckRequestWithDefaults instantiates a new VerificationCheckRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerificationId

`func (o *VerificationCheckRequest) GetVerificationId() string`

GetVerificationId returns the VerificationId field if non-nil, zero value otherwise.

### GetVerificationIdOk

`func (o *VerificationCheckRequest) GetVerificationIdOk() (*string, bool)`

GetVerificationIdOk returns a tuple with the VerificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationId

`func (o *VerificationCheckRequest) SetVerificationId(v string)`

SetVerificationId sets VerificationId field to given value.

### HasVerificationId

`func (o *VerificationCheckRequest) HasVerificationId() bool`

HasVerificationId returns a boolean if a field has been set.

### GetTo

`func (o *VerificationCheckRequest) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *VerificationCheckRequest) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *VerificationCheckRequest) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *VerificationCheckRequest) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetCode

`func (o *VerificationCheckRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *VerificationCheckRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *VerificationCheckRequest) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *VerificationCheckRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


