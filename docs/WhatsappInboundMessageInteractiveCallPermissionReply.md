# WhatsappInboundMessageInteractiveCallPermissionReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Response** | Pointer to **string** | The customer&#39;s response to the call permission request. - &#x60;accept&#x60;: User granted permission for business to call back - &#x60;reject&#x60;: User rejected permission for business to call back | [optional] 
**ExpirationTimestamp** | Pointer to **int64** | The timestamp (in seconds) when the call permission expires. Only present when response is \&quot;accept\&quot; and is_permanent is false. | [optional] 
**IsPermanent** | Pointer to **bool** | Whether the permission is permanent or temporary. - &#x60;true&#x60;: Permanent authorization (no expiration) - &#x60;false&#x60;: Temporary authorization (expires at expiration_timestamp) | [optional] 
**ResponseSource** | Pointer to **string** | The source of this permission response. - &#x60;user_action&#x60;: User explicitly approved or rejected the permission - &#x60;automatic&#x60;: Automatic permission approval due to the WhatsApp user initiating the call | [optional] 

## Methods

### NewWhatsappInboundMessageInteractiveCallPermissionReply

`func NewWhatsappInboundMessageInteractiveCallPermissionReply() *WhatsappInboundMessageInteractiveCallPermissionReply`

NewWhatsappInboundMessageInteractiveCallPermissionReply instantiates a new WhatsappInboundMessageInteractiveCallPermissionReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappInboundMessageInteractiveCallPermissionReplyWithDefaults

`func NewWhatsappInboundMessageInteractiveCallPermissionReplyWithDefaults() *WhatsappInboundMessageInteractiveCallPermissionReply`

NewWhatsappInboundMessageInteractiveCallPermissionReplyWithDefaults instantiates a new WhatsappInboundMessageInteractiveCallPermissionReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponse

`func (o *WhatsappInboundMessageInteractiveCallPermissionReply) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *WhatsappInboundMessageInteractiveCallPermissionReply) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *WhatsappInboundMessageInteractiveCallPermissionReply) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *WhatsappInboundMessageInteractiveCallPermissionReply) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetExpirationTimestamp

`func (o *WhatsappInboundMessageInteractiveCallPermissionReply) GetExpirationTimestamp() int64`

GetExpirationTimestamp returns the ExpirationTimestamp field if non-nil, zero value otherwise.

### GetExpirationTimestampOk

`func (o *WhatsappInboundMessageInteractiveCallPermissionReply) GetExpirationTimestampOk() (*int64, bool)`

GetExpirationTimestampOk returns a tuple with the ExpirationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTimestamp

`func (o *WhatsappInboundMessageInteractiveCallPermissionReply) SetExpirationTimestamp(v int64)`

SetExpirationTimestamp sets ExpirationTimestamp field to given value.

### HasExpirationTimestamp

`func (o *WhatsappInboundMessageInteractiveCallPermissionReply) HasExpirationTimestamp() bool`

HasExpirationTimestamp returns a boolean if a field has been set.

### GetIsPermanent

`func (o *WhatsappInboundMessageInteractiveCallPermissionReply) GetIsPermanent() bool`

GetIsPermanent returns the IsPermanent field if non-nil, zero value otherwise.

### GetIsPermanentOk

`func (o *WhatsappInboundMessageInteractiveCallPermissionReply) GetIsPermanentOk() (*bool, bool)`

GetIsPermanentOk returns a tuple with the IsPermanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPermanent

`func (o *WhatsappInboundMessageInteractiveCallPermissionReply) SetIsPermanent(v bool)`

SetIsPermanent sets IsPermanent field to given value.

### HasIsPermanent

`func (o *WhatsappInboundMessageInteractiveCallPermissionReply) HasIsPermanent() bool`

HasIsPermanent returns a boolean if a field has been set.

### GetResponseSource

`func (o *WhatsappInboundMessageInteractiveCallPermissionReply) GetResponseSource() string`

GetResponseSource returns the ResponseSource field if non-nil, zero value otherwise.

### GetResponseSourceOk

`func (o *WhatsappInboundMessageInteractiveCallPermissionReply) GetResponseSourceOk() (*string, bool)`

GetResponseSourceOk returns a tuple with the ResponseSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseSource

`func (o *WhatsappInboundMessageInteractiveCallPermissionReply) SetResponseSource(v string)`

SetResponseSource sets ResponseSource field to given value.

### HasResponseSource

`func (o *WhatsappInboundMessageInteractiveCallPermissionReply) HasResponseSource() bool`

HasResponseSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


