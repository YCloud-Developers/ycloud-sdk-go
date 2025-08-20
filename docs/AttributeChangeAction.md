# AttributeChangeAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | The type of change action performed. | 
**Id** | Pointer to **string** | The ID of the item when the attribute is &#39;tags&#39;.  This field is only present for tag-related changes. | [optional] 
**Value** | Pointer to **string** | The value of the item when the attribute is &#39;tags&#39;.  This field is only present for tag-related changes. | [optional] 

## Methods

### NewAttributeChangeAction

`func NewAttributeChangeAction(action string, ) *AttributeChangeAction`

NewAttributeChangeAction instantiates a new AttributeChangeAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributeChangeActionWithDefaults

`func NewAttributeChangeActionWithDefaults() *AttributeChangeAction`

NewAttributeChangeActionWithDefaults instantiates a new AttributeChangeAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *AttributeChangeAction) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AttributeChangeAction) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AttributeChangeAction) SetAction(v string)`

SetAction sets Action field to given value.


### GetId

`func (o *AttributeChangeAction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AttributeChangeAction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AttributeChangeAction) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AttributeChangeAction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetValue

`func (o *AttributeChangeAction) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AttributeChangeAction) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AttributeChangeAction) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *AttributeChangeAction) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


