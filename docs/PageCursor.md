# PageCursor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**After** | Pointer to **string** | A cursor to fetch the next page in cursor pagination. For example, if you make a list request, receive 100 objects and &#x60;cursor.after&#x3D;id:foo&#x60;, your subsequent call can include &#x60;pageAfter&#x3D;id:foo&#x60; in order to fetch the next page of the list. This field is returned only if there are more items in the list. | [optional] 

## Methods

### NewPageCursor

`func NewPageCursor() *PageCursor`

NewPageCursor instantiates a new PageCursor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageCursorWithDefaults

`func NewPageCursorWithDefaults() *PageCursor`

NewPageCursorWithDefaults instantiates a new PageCursor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfter

`func (o *PageCursor) GetAfter() string`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *PageCursor) GetAfterOk() (*string, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *PageCursor) SetAfter(v string)`

SetAfter sets After field to given value.

### HasAfter

`func (o *PageCursor) HasAfter() bool`

HasAfter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
