# WhatsappMessageMediaOneOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Link** | **string** | Required when sending media directly from your server.  The protocol and URL of the media to be sent. Use only with HTTP/HTTPS URLs. Note: WhatsApp Cloud API caches media resources for 10 minutes. To ensure latest content, add random query strings to the URL. | 

## Methods

### NewWhatsappMessageMediaOneOf1

`func NewWhatsappMessageMediaOneOf1(link string, ) *WhatsappMessageMediaOneOf1`

NewWhatsappMessageMediaOneOf1 instantiates a new WhatsappMessageMediaOneOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappMessageMediaOneOf1WithDefaults

`func NewWhatsappMessageMediaOneOf1WithDefaults() *WhatsappMessageMediaOneOf1`

NewWhatsappMessageMediaOneOf1WithDefaults instantiates a new WhatsappMessageMediaOneOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLink

`func (o *WhatsappMessageMediaOneOf1) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *WhatsappMessageMediaOneOf1) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *WhatsappMessageMediaOneOf1) SetLink(v string)`

SetLink sets Link field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


