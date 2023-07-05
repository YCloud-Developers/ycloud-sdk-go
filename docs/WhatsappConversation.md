# WhatsappConversation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique ID for the object. | [optional] 
**Type** | Pointer to [**WhatsappConversationType**](WhatsappConversationType.md) |  | [optional] 
**OriginType** | Pointer to [**WhatsappConversationOriginType**](WhatsappConversationOriginType.md) |  | [optional] 
**ExpireTime** | Pointer to **time.Time** | Date when the conversation expires, formatted in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339). e.g., &#x60;2022-06-01T12:00:00.000Z&#x60;. | [optional] 

## Methods

### NewWhatsappConversation

`func NewWhatsappConversation() *WhatsappConversation`

NewWhatsappConversation instantiates a new WhatsappConversation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappConversationWithDefaults

`func NewWhatsappConversationWithDefaults() *WhatsappConversation`

NewWhatsappConversationWithDefaults instantiates a new WhatsappConversation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WhatsappConversation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WhatsappConversation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WhatsappConversation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WhatsappConversation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *WhatsappConversation) GetType() WhatsappConversationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WhatsappConversation) GetTypeOk() (*WhatsappConversationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WhatsappConversation) SetType(v WhatsappConversationType)`

SetType sets Type field to given value.

### HasType

`func (o *WhatsappConversation) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOriginType

`func (o *WhatsappConversation) GetOriginType() WhatsappConversationOriginType`

GetOriginType returns the OriginType field if non-nil, zero value otherwise.

### GetOriginTypeOk

`func (o *WhatsappConversation) GetOriginTypeOk() (*WhatsappConversationOriginType, bool)`

GetOriginTypeOk returns a tuple with the OriginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginType

`func (o *WhatsappConversation) SetOriginType(v WhatsappConversationOriginType)`

SetOriginType sets OriginType field to given value.

### HasOriginType

`func (o *WhatsappConversation) HasOriginType() bool`

HasOriginType returns a boolean if a field has been set.

### GetExpireTime

`func (o *WhatsappConversation) GetExpireTime() time.Time`

GetExpireTime returns the ExpireTime field if non-nil, zero value otherwise.

### GetExpireTimeOk

`func (o *WhatsappConversation) GetExpireTimeOk() (*time.Time, bool)`

GetExpireTimeOk returns a tuple with the ExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireTime

`func (o *WhatsappConversation) SetExpireTime(v time.Time)`

SetExpireTime sets ExpireTime field to given value.

### HasExpireTime

`func (o *WhatsappConversation) HasExpireTime() bool`

HasExpireTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


