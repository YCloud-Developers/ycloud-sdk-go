# ContactPageCursor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**After** | **string** | Contact ID of the last public item in this page. Pass this value unchanged as &#x60;pageAfter&#x60; to fetch the next page. |

## Methods

### NewContactPageCursor

`func NewContactPageCursor(after string, ) *ContactPageCursor`

NewContactPageCursor instantiates a new ContactPageCursor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactPageCursorWithDefaults

`func NewContactPageCursorWithDefaults() *ContactPageCursor`

NewContactPageCursorWithDefaults instantiates a new ContactPageCursor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfter

`func (o *ContactPageCursor) GetAfter() string`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *ContactPageCursor) GetAfterOk() (*string, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *ContactPageCursor) SetAfter(v string)`

SetAfter sets After field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
