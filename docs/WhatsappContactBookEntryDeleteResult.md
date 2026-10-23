# WhatsappContactBookEntryDeleteResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | Always &#x60;true&#x60; in an HTTP 200 response. |
**Deleted** | **bool** | Meta&#39;s deletion result. &#x60;true&#x60; means Meta reports that it deleted a matching entry. &#x60;false&#x60; means Meta processed the request but found no matching entry to delete. |

## Methods

### NewWhatsappContactBookEntryDeleteResult

`func NewWhatsappContactBookEntryDeleteResult(success bool, deleted bool, ) *WhatsappContactBookEntryDeleteResult`

NewWhatsappContactBookEntryDeleteResult instantiates a new WhatsappContactBookEntryDeleteResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappContactBookEntryDeleteResultWithDefaults

`func NewWhatsappContactBookEntryDeleteResultWithDefaults() *WhatsappContactBookEntryDeleteResult`

NewWhatsappContactBookEntryDeleteResultWithDefaults instantiates a new WhatsappContactBookEntryDeleteResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *WhatsappContactBookEntryDeleteResult) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *WhatsappContactBookEntryDeleteResult) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *WhatsappContactBookEntryDeleteResult) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetDeleted

`func (o *WhatsappContactBookEntryDeleteResult) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *WhatsappContactBookEntryDeleteResult) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *WhatsappContactBookEntryDeleteResult) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
