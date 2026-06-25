# WhatsappBusinessUsernameUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | Business Username to request for the phone number. Send the plain username without &#x60;@&#x60;.  YCloud trims leading and trailing whitespace and normalizes the value to lowercase before validation and submission. The value must be 3-35 characters, contain only English letters, numbers, periods, and underscores, and contain at least one English letter. It must not start or end with a period, contain consecutive periods, start with &#x60;www&#x60;, or end with common domain suffixes such as &#x60;.com&#x60;, &#x60;.org&#x60;, &#x60;.net&#x60;, &#x60;.int&#x60;, &#x60;.edu&#x60;, &#x60;.gov&#x60;, &#x60;.mil&#x60;, &#x60;.us&#x60;, &#x60;.in&#x60;, or &#x60;.html&#x60;. | 

## Methods

### NewWhatsappBusinessUsernameUpdateRequest

`func NewWhatsappBusinessUsernameUpdateRequest(username string, ) *WhatsappBusinessUsernameUpdateRequest`

NewWhatsappBusinessUsernameUpdateRequest instantiates a new WhatsappBusinessUsernameUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappBusinessUsernameUpdateRequestWithDefaults

`func NewWhatsappBusinessUsernameUpdateRequestWithDefaults() *WhatsappBusinessUsernameUpdateRequest`

NewWhatsappBusinessUsernameUpdateRequestWithDefaults instantiates a new WhatsappBusinessUsernameUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *WhatsappBusinessUsernameUpdateRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *WhatsappBusinessUsernameUpdateRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *WhatsappBusinessUsernameUpdateRequest) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


