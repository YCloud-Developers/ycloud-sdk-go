# WhatsappApiError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | A human-readable description of the error. | 
**Code** | **string** | An error code. | 
**Type** | Pointer to **string** | Error type. | [optional] 
**ErrorSubcode** | Pointer to **string** | Additional code about the error. | [optional] 
**ErrorUserMsg** | Pointer to **string** | The message to display to the user. The language of the message is based on the locale of the API request. | [optional] 
**ErrorUserTitle** | Pointer to **string** | The title of the dialog, if shown. The language of the message is based on the locale of the API request. | [optional] 
**FbtraceId** | Pointer to **string** | Internal support identifier. When reporting a bug related to a Graph API call, include the fbtrace_id to help us find log data for debugging. | [optional] 
**ErrorData** | Pointer to **map[string]interface{}** | Additional data about the error. A string or map. - For template APIs, this field is a string describing the reason for the error.   - For message APIs, this field is a map with property &#x60;details&#x60; describing the reason for the error. | [optional] 

## Methods

### NewWhatsappApiError

`func NewWhatsappApiError(message string, code string, ) *WhatsappApiError`

NewWhatsappApiError instantiates a new WhatsappApiError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappApiErrorWithDefaults

`func NewWhatsappApiErrorWithDefaults() *WhatsappApiError`

NewWhatsappApiErrorWithDefaults instantiates a new WhatsappApiError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *WhatsappApiError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *WhatsappApiError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *WhatsappApiError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetCode

`func (o *WhatsappApiError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *WhatsappApiError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *WhatsappApiError) SetCode(v string)`

SetCode sets Code field to given value.


### GetType

`func (o *WhatsappApiError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WhatsappApiError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WhatsappApiError) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WhatsappApiError) HasType() bool`

HasType returns a boolean if a field has been set.

### GetErrorSubcode

`func (o *WhatsappApiError) GetErrorSubcode() string`

GetErrorSubcode returns the ErrorSubcode field if non-nil, zero value otherwise.

### GetErrorSubcodeOk

`func (o *WhatsappApiError) GetErrorSubcodeOk() (*string, bool)`

GetErrorSubcodeOk returns a tuple with the ErrorSubcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorSubcode

`func (o *WhatsappApiError) SetErrorSubcode(v string)`

SetErrorSubcode sets ErrorSubcode field to given value.

### HasErrorSubcode

`func (o *WhatsappApiError) HasErrorSubcode() bool`

HasErrorSubcode returns a boolean if a field has been set.

### GetErrorUserMsg

`func (o *WhatsappApiError) GetErrorUserMsg() string`

GetErrorUserMsg returns the ErrorUserMsg field if non-nil, zero value otherwise.

### GetErrorUserMsgOk

`func (o *WhatsappApiError) GetErrorUserMsgOk() (*string, bool)`

GetErrorUserMsgOk returns a tuple with the ErrorUserMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorUserMsg

`func (o *WhatsappApiError) SetErrorUserMsg(v string)`

SetErrorUserMsg sets ErrorUserMsg field to given value.

### HasErrorUserMsg

`func (o *WhatsappApiError) HasErrorUserMsg() bool`

HasErrorUserMsg returns a boolean if a field has been set.

### GetErrorUserTitle

`func (o *WhatsappApiError) GetErrorUserTitle() string`

GetErrorUserTitle returns the ErrorUserTitle field if non-nil, zero value otherwise.

### GetErrorUserTitleOk

`func (o *WhatsappApiError) GetErrorUserTitleOk() (*string, bool)`

GetErrorUserTitleOk returns a tuple with the ErrorUserTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorUserTitle

`func (o *WhatsappApiError) SetErrorUserTitle(v string)`

SetErrorUserTitle sets ErrorUserTitle field to given value.

### HasErrorUserTitle

`func (o *WhatsappApiError) HasErrorUserTitle() bool`

HasErrorUserTitle returns a boolean if a field has been set.

### GetFbtraceId

`func (o *WhatsappApiError) GetFbtraceId() string`

GetFbtraceId returns the FbtraceId field if non-nil, zero value otherwise.

### GetFbtraceIdOk

`func (o *WhatsappApiError) GetFbtraceIdOk() (*string, bool)`

GetFbtraceIdOk returns a tuple with the FbtraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFbtraceId

`func (o *WhatsappApiError) SetFbtraceId(v string)`

SetFbtraceId sets FbtraceId field to given value.

### HasFbtraceId

`func (o *WhatsappApiError) HasFbtraceId() bool`

HasFbtraceId returns a boolean if a field has been set.

### GetErrorData

`func (o *WhatsappApiError) GetErrorData() map[string]interface{}`

GetErrorData returns the ErrorData field if non-nil, zero value otherwise.

### GetErrorDataOk

`func (o *WhatsappApiError) GetErrorDataOk() (*map[string]interface{}, bool)`

GetErrorDataOk returns a tuple with the ErrorData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorData

`func (o *WhatsappApiError) SetErrorData(v map[string]interface{})`

SetErrorData sets ErrorData field to given value.

### HasErrorData

`func (o *WhatsappApiError) HasErrorData() bool`

HasErrorData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


