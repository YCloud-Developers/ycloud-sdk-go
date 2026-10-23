# ParameterBindingDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional]
**Value** | Pointer to **string** |  | [optional]
**Macro** | Pointer to **string** | Official binding macro. Accepting an Instagram macro does not enable Instagram entities or guarantee its value for WhatsApp consumers. | [optional]

## Methods

### NewParameterBindingDTO

`func NewParameterBindingDTO() *ParameterBindingDTO`

NewParameterBindingDTO instantiates a new ParameterBindingDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterBindingDTOWithDefaults

`func NewParameterBindingDTOWithDefaults() *ParameterBindingDTO`

NewParameterBindingDTOWithDefaults instantiates a new ParameterBindingDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ParameterBindingDTO) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ParameterBindingDTO) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ParameterBindingDTO) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ParameterBindingDTO) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetValue

`func (o *ParameterBindingDTO) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ParameterBindingDTO) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ParameterBindingDTO) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ParameterBindingDTO) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetMacro

`func (o *ParameterBindingDTO) GetMacro() string`

GetMacro returns the Macro field if non-nil, zero value otherwise.

### GetMacroOk

`func (o *ParameterBindingDTO) GetMacroOk() (*string, bool)`

GetMacroOk returns a tuple with the Macro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacro

`func (o *ParameterBindingDTO) SetMacro(v string)`

SetMacro sets Macro field to given value.

### HasMacro

`func (o *ParameterBindingDTO) HasMacro() bool`

HasMacro returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
