# ContactAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the contact attribute. | 
**Name** | **string** | Display name of the contact attribute. | 
**Key** | **string** | Key name used to reference this attribute. | 
**Type** | **string** | Data type of the contact attribute. | 
**Desc** | Pointer to **string** | Description of the contact attribute. | [optional] 
**Values** | Pointer to **[]string** | Array of possible values for this attribute. Only present when type is \&quot;ARRAY\&quot;. | [optional] 

## Methods

### NewContactAttribute

`func NewContactAttribute(id string, name string, key string, type_ string, ) *ContactAttribute`

NewContactAttribute instantiates a new ContactAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactAttributeWithDefaults

`func NewContactAttributeWithDefaults() *ContactAttribute`

NewContactAttributeWithDefaults instantiates a new ContactAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContactAttribute) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContactAttribute) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContactAttribute) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ContactAttribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContactAttribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContactAttribute) SetName(v string)`

SetName sets Name field to given value.


### GetKey

`func (o *ContactAttribute) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ContactAttribute) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ContactAttribute) SetKey(v string)`

SetKey sets Key field to given value.


### GetType

`func (o *ContactAttribute) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContactAttribute) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContactAttribute) SetType(v string)`

SetType sets Type field to given value.


### GetDesc

`func (o *ContactAttribute) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *ContactAttribute) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *ContactAttribute) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *ContactAttribute) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetValues

`func (o *ContactAttribute) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ContactAttribute) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ContactAttribute) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *ContactAttribute) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


