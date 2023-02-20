# EmailDelivery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailId** | **string** | Unique ID for the related email you&#39;ve previously sent. | 
**RecipientAddress** | **string** | A recipient&#39;s email address. | 
**Status** | Pointer to **string** | Delivery status of the email to the specific recipient address. - &#x60;sending&#x60;: The messaging request is accepted by our system. - &#x60;failed&#x60;: The message failed to be sent from our system. - &#x60;sent&#x60;: The message has been sent from YCloud. - &#x60;delivered&#x60;: YCloud has received a delivery receipt indicating that message is delivered. - &#x60;undelivered&#x60;: YCloud has received a delivery receipt indicating that message is not delivered. | [optional] 
**ErrorCode** | Pointer to **string** | Error code when the email is undeliverable. | [optional] 
**ErrorMessage** | Pointer to **string** | Error message when the email is undeliverable. | [optional] 
**ExternalId** | Pointer to **string** | The &#x60;externalId&#x60; you specified when you sent the email. | [optional] 
**BizType** | Pointer to **string** | This can be either empty or one of &#x60;email&#x60;, or &#x60;verify&#x60;. Defaults to &#x60;email&#x60;. - &#x60;email&#x60;: Indicates that the message is sent via [Email](https://www.ycloud.com/email) product. - &#x60;verify&#x60;: Indicates that the message is sent via [Verify](https://www.ycloud.com/verify) product. | [optional] 
**VerificationId** | Pointer to **string** | The verification ID. Included only when &#x60;bizType&#x60; is &#x60;verify&#x60;. | [optional] 

## Methods

### NewEmailDelivery

`func NewEmailDelivery(emailId string, recipientAddress string, ) *EmailDelivery`

NewEmailDelivery instantiates a new EmailDelivery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailDeliveryWithDefaults

`func NewEmailDeliveryWithDefaults() *EmailDelivery`

NewEmailDeliveryWithDefaults instantiates a new EmailDelivery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailId

`func (o *EmailDelivery) GetEmailId() string`

GetEmailId returns the EmailId field if non-nil, zero value otherwise.

### GetEmailIdOk

`func (o *EmailDelivery) GetEmailIdOk() (*string, bool)`

GetEmailIdOk returns a tuple with the EmailId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailId

`func (o *EmailDelivery) SetEmailId(v string)`

SetEmailId sets EmailId field to given value.


### GetRecipientAddress

`func (o *EmailDelivery) GetRecipientAddress() string`

GetRecipientAddress returns the RecipientAddress field if non-nil, zero value otherwise.

### GetRecipientAddressOk

`func (o *EmailDelivery) GetRecipientAddressOk() (*string, bool)`

GetRecipientAddressOk returns a tuple with the RecipientAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientAddress

`func (o *EmailDelivery) SetRecipientAddress(v string)`

SetRecipientAddress sets RecipientAddress field to given value.


### GetStatus

`func (o *EmailDelivery) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EmailDelivery) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EmailDelivery) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EmailDelivery) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrorCode

`func (o *EmailDelivery) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *EmailDelivery) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *EmailDelivery) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *EmailDelivery) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *EmailDelivery) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *EmailDelivery) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *EmailDelivery) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *EmailDelivery) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetExternalId

`func (o *EmailDelivery) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *EmailDelivery) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *EmailDelivery) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *EmailDelivery) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetBizType

`func (o *EmailDelivery) GetBizType() string`

GetBizType returns the BizType field if non-nil, zero value otherwise.

### GetBizTypeOk

`func (o *EmailDelivery) GetBizTypeOk() (*string, bool)`

GetBizTypeOk returns a tuple with the BizType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBizType

`func (o *EmailDelivery) SetBizType(v string)`

SetBizType sets BizType field to given value.

### HasBizType

`func (o *EmailDelivery) HasBizType() bool`

HasBizType returns a boolean if a field has been set.

### GetVerificationId

`func (o *EmailDelivery) GetVerificationId() string`

GetVerificationId returns the VerificationId field if non-nil, zero value otherwise.

### GetVerificationIdOk

`func (o *EmailDelivery) GetVerificationIdOk() (*string, bool)`

GetVerificationIdOk returns a tuple with the VerificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationId

`func (o *EmailDelivery) SetVerificationId(v string)`

SetVerificationId sets VerificationId field to given value.

### HasVerificationId

`func (o *EmailDelivery) HasVerificationId() bool`

HasVerificationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


