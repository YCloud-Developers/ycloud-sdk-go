# WhatsappMessageOrderAmount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Offset** | **int32** | Must be &#x60;100&#x60; for &#x60;INR&#x60;. | 
**Value** | **int32** | Positive integer representing the amount value multiplied by offset.  For example, ₹12.34 has value 1234. | 
**Description** | Pointer to **string** | Use only for &#x60;tax&#x60;, &#x60;shipping&#x60;, or &#x60;discount&#x60;. Description of the amount. Max character limit is 60 characters. | [optional] 
**DiscountProgramName** | Pointer to **string** | Use only for &#x60;discount&#x60;. Text used for defining incentivised orders. If order is incentivised, the merchant needs to define this information. Max character limit is 60 characters. | [optional] 

## Methods

### NewWhatsappMessageOrderAmount

`func NewWhatsappMessageOrderAmount(offset int32, value int32, ) *WhatsappMessageOrderAmount`

NewWhatsappMessageOrderAmount instantiates a new WhatsappMessageOrderAmount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappMessageOrderAmountWithDefaults

`func NewWhatsappMessageOrderAmountWithDefaults() *WhatsappMessageOrderAmount`

NewWhatsappMessageOrderAmountWithDefaults instantiates a new WhatsappMessageOrderAmount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffset

`func (o *WhatsappMessageOrderAmount) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *WhatsappMessageOrderAmount) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *WhatsappMessageOrderAmount) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetValue

`func (o *WhatsappMessageOrderAmount) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *WhatsappMessageOrderAmount) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *WhatsappMessageOrderAmount) SetValue(v int32)`

SetValue sets Value field to given value.


### GetDescription

`func (o *WhatsappMessageOrderAmount) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WhatsappMessageOrderAmount) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WhatsappMessageOrderAmount) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WhatsappMessageOrderAmount) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiscountProgramName

`func (o *WhatsappMessageOrderAmount) GetDiscountProgramName() string`

GetDiscountProgramName returns the DiscountProgramName field if non-nil, zero value otherwise.

### GetDiscountProgramNameOk

`func (o *WhatsappMessageOrderAmount) GetDiscountProgramNameOk() (*string, bool)`

GetDiscountProgramNameOk returns a tuple with the DiscountProgramName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountProgramName

`func (o *WhatsappMessageOrderAmount) SetDiscountProgramName(v string)`

SetDiscountProgramName sets DiscountProgramName field to given value.

### HasDiscountProgramName

`func (o *WhatsappMessageOrderAmount) HasDiscountProgramName() bool`

HasDiscountProgramName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


