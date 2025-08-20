# ContactAttributeChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OldValue** | Pointer to [**ContactAttributeChangeOldValue**](ContactAttributeChangeOldValue.md) |  | [optional] 
**NewValue** | Pointer to [**ContactAttributeChangeNewValue**](ContactAttributeChangeNewValue.md) |  | [optional] 
**Extra** | Pointer to [**[]AttributeChangeAction**](AttributeChangeAction.md) | An array of change actions that describe what operations were performed on this attribute. | [optional] 

## Methods

### NewContactAttributeChange

`func NewContactAttributeChange() *ContactAttributeChange`

NewContactAttributeChange instantiates a new ContactAttributeChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactAttributeChangeWithDefaults

`func NewContactAttributeChangeWithDefaults() *ContactAttributeChange`

NewContactAttributeChangeWithDefaults instantiates a new ContactAttributeChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOldValue

`func (o *ContactAttributeChange) GetOldValue() ContactAttributeChangeOldValue`

GetOldValue returns the OldValue field if non-nil, zero value otherwise.

### GetOldValueOk

`func (o *ContactAttributeChange) GetOldValueOk() (*ContactAttributeChangeOldValue, bool)`

GetOldValueOk returns a tuple with the OldValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldValue

`func (o *ContactAttributeChange) SetOldValue(v ContactAttributeChangeOldValue)`

SetOldValue sets OldValue field to given value.

### HasOldValue

`func (o *ContactAttributeChange) HasOldValue() bool`

HasOldValue returns a boolean if a field has been set.

### GetNewValue

`func (o *ContactAttributeChange) GetNewValue() ContactAttributeChangeNewValue`

GetNewValue returns the NewValue field if non-nil, zero value otherwise.

### GetNewValueOk

`func (o *ContactAttributeChange) GetNewValueOk() (*ContactAttributeChangeNewValue, bool)`

GetNewValueOk returns a tuple with the NewValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewValue

`func (o *ContactAttributeChange) SetNewValue(v ContactAttributeChangeNewValue)`

SetNewValue sets NewValue field to given value.

### HasNewValue

`func (o *ContactAttributeChange) HasNewValue() bool`

HasNewValue returns a boolean if a field has been set.

### GetExtra

`func (o *ContactAttributeChange) GetExtra() []AttributeChangeAction`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *ContactAttributeChange) GetExtraOk() (*[]AttributeChangeAction, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *ContactAttributeChange) SetExtra(v []AttributeChangeAction)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *ContactAttributeChange) HasExtra() bool`

HasExtra returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


