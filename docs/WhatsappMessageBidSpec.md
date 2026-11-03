# WhatsappMessageBidSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PerMessageBidMultiplier** | **string** | Supplemental bid multiplier for this message, greater than 0 and at most 999.999, with at most 3 decimal places. Values above 1 raise the ceiling; values below 1 lower it. Required whenever the message bidSpec object is provided. |

## Methods

### NewWhatsappMessageBidSpec

`func NewWhatsappMessageBidSpec(perMessageBidMultiplier string, ) *WhatsappMessageBidSpec`

NewWhatsappMessageBidSpec instantiates a new WhatsappMessageBidSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappMessageBidSpecWithDefaults

`func NewWhatsappMessageBidSpecWithDefaults() *WhatsappMessageBidSpec`

NewWhatsappMessageBidSpecWithDefaults instantiates a new WhatsappMessageBidSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPerMessageBidMultiplier

`func (o *WhatsappMessageBidSpec) GetPerMessageBidMultiplier() string`

GetPerMessageBidMultiplier returns the PerMessageBidMultiplier field if non-nil, zero value otherwise.

### GetPerMessageBidMultiplierOk

`func (o *WhatsappMessageBidSpec) GetPerMessageBidMultiplierOk() (*string, bool)`

GetPerMessageBidMultiplierOk returns a tuple with the PerMessageBidMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerMessageBidMultiplier

`func (o *WhatsappMessageBidSpec) SetPerMessageBidMultiplier(v string)`

SetPerMessageBidMultiplier sets PerMessageBidMultiplier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
