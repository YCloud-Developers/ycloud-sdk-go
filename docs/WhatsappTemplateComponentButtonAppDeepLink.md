# WhatsappTemplateComponentButtonAppDeepLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetaAppId** | Pointer to **string** | Required if using a URL button mapped to a deep link. APP ID. | [optional] 
**AndroidDeepLink** | Pointer to **string** | Required if using a URL button component mapped to a deep link.The WhatsApp client will attempt to load this URI if the WhatsApp user taps the button on an Android device. | [optional] 
**AndroidFallbackPlaystoreUrl** | Pointer to **string** | Optional. URL of a website that the WhatsApp client will attempt to load in the device’s default web browser when the button is tapped but unable to load the Android deep link URI. | [optional] 

## Methods

### NewWhatsappTemplateComponentButtonAppDeepLink

`func NewWhatsappTemplateComponentButtonAppDeepLink() *WhatsappTemplateComponentButtonAppDeepLink`

NewWhatsappTemplateComponentButtonAppDeepLink instantiates a new WhatsappTemplateComponentButtonAppDeepLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappTemplateComponentButtonAppDeepLinkWithDefaults

`func NewWhatsappTemplateComponentButtonAppDeepLinkWithDefaults() *WhatsappTemplateComponentButtonAppDeepLink`

NewWhatsappTemplateComponentButtonAppDeepLinkWithDefaults instantiates a new WhatsappTemplateComponentButtonAppDeepLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetaAppId

`func (o *WhatsappTemplateComponentButtonAppDeepLink) GetMetaAppId() string`

GetMetaAppId returns the MetaAppId field if non-nil, zero value otherwise.

### GetMetaAppIdOk

`func (o *WhatsappTemplateComponentButtonAppDeepLink) GetMetaAppIdOk() (*string, bool)`

GetMetaAppIdOk returns a tuple with the MetaAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaAppId

`func (o *WhatsappTemplateComponentButtonAppDeepLink) SetMetaAppId(v string)`

SetMetaAppId sets MetaAppId field to given value.

### HasMetaAppId

`func (o *WhatsappTemplateComponentButtonAppDeepLink) HasMetaAppId() bool`

HasMetaAppId returns a boolean if a field has been set.

### GetAndroidDeepLink

`func (o *WhatsappTemplateComponentButtonAppDeepLink) GetAndroidDeepLink() string`

GetAndroidDeepLink returns the AndroidDeepLink field if non-nil, zero value otherwise.

### GetAndroidDeepLinkOk

`func (o *WhatsappTemplateComponentButtonAppDeepLink) GetAndroidDeepLinkOk() (*string, bool)`

GetAndroidDeepLinkOk returns a tuple with the AndroidDeepLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidDeepLink

`func (o *WhatsappTemplateComponentButtonAppDeepLink) SetAndroidDeepLink(v string)`

SetAndroidDeepLink sets AndroidDeepLink field to given value.

### HasAndroidDeepLink

`func (o *WhatsappTemplateComponentButtonAppDeepLink) HasAndroidDeepLink() bool`

HasAndroidDeepLink returns a boolean if a field has been set.

### GetAndroidFallbackPlaystoreUrl

`func (o *WhatsappTemplateComponentButtonAppDeepLink) GetAndroidFallbackPlaystoreUrl() string`

GetAndroidFallbackPlaystoreUrl returns the AndroidFallbackPlaystoreUrl field if non-nil, zero value otherwise.

### GetAndroidFallbackPlaystoreUrlOk

`func (o *WhatsappTemplateComponentButtonAppDeepLink) GetAndroidFallbackPlaystoreUrlOk() (*string, bool)`

GetAndroidFallbackPlaystoreUrlOk returns a tuple with the AndroidFallbackPlaystoreUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidFallbackPlaystoreUrl

`func (o *WhatsappTemplateComponentButtonAppDeepLink) SetAndroidFallbackPlaystoreUrl(v string)`

SetAndroidFallbackPlaystoreUrl sets AndroidFallbackPlaystoreUrl field to given value.

### HasAndroidFallbackPlaystoreUrl

`func (o *WhatsappTemplateComponentButtonAppDeepLink) HasAndroidFallbackPlaystoreUrl() bool`

HasAndroidFallbackPlaystoreUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


