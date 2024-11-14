# EmailSendRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | **string** | - The sender&#39;s email. Its domain should be one that has been registered and activated in your account. - The sender&#39;s email address is required while the sender&#39;s name is optional. For example, both &#x60;support@example.com&#x60; and &#x60;Sender&lt;support@example.com&gt;&#x60; work. | 
**To** | **string** | - The intended recipients&#39; email addresses. - Supports a comma-separated list of one or more addresses. Max items: 100. | 
**Subject** | **string** | The email subject, which contains a short string identifying the topic of the message. Max length: 255. | 
**Content** | **string** | - The email body. Max size: 150 KB. - Variables in the form of &#x60;#var_1#&#x60; are supported, they should be used together with the &#x60;variables&#x60; parameter. Variable keys only support letters, digits, and the underline character (&#x60;_&#x60;). - You can use the [Test Templates](https://help.ycloud.com/en/articles/6006545) provided by YCloud for testing. | 
**ContentType** | Pointer to [**EmailContentType**](EmailContentType.md) |  | [optional] 
**Variables** | Pointer to **[]map[string]string** | - The variable key-value pairs that will replace the variable placeholders in &#x60;content&#x60; for each recipient. Variable keys are those that are wrapped with &#x60;#&#x60; as placeholders (e.g., &#x60;#var_1#&#x60;) in &#x60;content&#x60;. The placeholders will be replaced by variable values when sending the email. - The size of the array must be the same as the number of recipients in &#x60;to&#x60;. Be aware that &#x60;cc&#x60; and &#x60;bcc&#x60; addresses are excluded, and they can not receive emails that contain variables. - This parameter&#39;s size will be calculated together with the parameter &#x60;content&#x60;. The whole size must not exceed 150 KB. | [optional] 
**Cc** | Pointer to **string** | Recipients who will receive a copy of the email. | [optional] 
**Bcc** | Pointer to **string** | Recipients who will receive a blind carbon copy of the email. | [optional] 
**ReplyTo** | Pointer to **string** | If this field exists, then the reply should go to the addresses indicated in that field and not to the address(es) indicated in the &#x60;from&#x60; field. | [optional] 
**Summary** | Pointer to **string** | This is a summary of your email. Max length: 70. | [optional] 
**ExternalId** | Pointer to **string** | A unique (recommended) string to reference the object. This can be an order number or similar, and can be used to reconcile the object with your internal systems. | [optional] 
**CallbackUrl** | Pointer to **string** | Delivery report URL. You can provide a URL, and we will push the updated status report to your server in time. e.g., https://httpbin.org/anything?tag&#x3D;api. Note: We recommend configuring Webhook Endpoints instead. | [optional] 

## Methods

### NewEmailSendRequest

`func NewEmailSendRequest(from string, to string, subject string, content string, ) *EmailSendRequest`

NewEmailSendRequest instantiates a new EmailSendRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailSendRequestWithDefaults

`func NewEmailSendRequestWithDefaults() *EmailSendRequest`

NewEmailSendRequestWithDefaults instantiates a new EmailSendRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *EmailSendRequest) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *EmailSendRequest) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *EmailSendRequest) SetFrom(v string)`

SetFrom sets From field to given value.


### GetTo

`func (o *EmailSendRequest) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *EmailSendRequest) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *EmailSendRequest) SetTo(v string)`

SetTo sets To field to given value.


### GetSubject

`func (o *EmailSendRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *EmailSendRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *EmailSendRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetContent

`func (o *EmailSendRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *EmailSendRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *EmailSendRequest) SetContent(v string)`

SetContent sets Content field to given value.


### GetContentType

`func (o *EmailSendRequest) GetContentType() EmailContentType`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *EmailSendRequest) GetContentTypeOk() (*EmailContentType, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *EmailSendRequest) SetContentType(v EmailContentType)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *EmailSendRequest) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetVariables

`func (o *EmailSendRequest) GetVariables() []map[string]string`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *EmailSendRequest) GetVariablesOk() (*[]map[string]string, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *EmailSendRequest) SetVariables(v []map[string]string)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *EmailSendRequest) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetCc

`func (o *EmailSendRequest) GetCc() string`

GetCc returns the Cc field if non-nil, zero value otherwise.

### GetCcOk

`func (o *EmailSendRequest) GetCcOk() (*string, bool)`

GetCcOk returns a tuple with the Cc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCc

`func (o *EmailSendRequest) SetCc(v string)`

SetCc sets Cc field to given value.

### HasCc

`func (o *EmailSendRequest) HasCc() bool`

HasCc returns a boolean if a field has been set.

### GetBcc

`func (o *EmailSendRequest) GetBcc() string`

GetBcc returns the Bcc field if non-nil, zero value otherwise.

### GetBccOk

`func (o *EmailSendRequest) GetBccOk() (*string, bool)`

GetBccOk returns a tuple with the Bcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcc

`func (o *EmailSendRequest) SetBcc(v string)`

SetBcc sets Bcc field to given value.

### HasBcc

`func (o *EmailSendRequest) HasBcc() bool`

HasBcc returns a boolean if a field has been set.

### GetReplyTo

`func (o *EmailSendRequest) GetReplyTo() string`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *EmailSendRequest) GetReplyToOk() (*string, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *EmailSendRequest) SetReplyTo(v string)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *EmailSendRequest) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetSummary

`func (o *EmailSendRequest) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *EmailSendRequest) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *EmailSendRequest) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *EmailSendRequest) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetExternalId

`func (o *EmailSendRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *EmailSendRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *EmailSendRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *EmailSendRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetCallbackUrl

`func (o *EmailSendRequest) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *EmailSendRequest) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *EmailSendRequest) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *EmailSendRequest) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


