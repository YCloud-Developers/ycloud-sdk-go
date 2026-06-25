# WhatsappGroupWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WabaId** | Pointer to **string** | WhatsApp Business Account ID. | [optional] 
**DisplayPhoneNumber** | Pointer to **string** | The display phone number from WhatsApp webhook metadata. | [optional] 
**PhoneNumberId** | Pointer to **string** | WhatsApp phone number ID. | [optional] 
**Field** | Pointer to [**WhatsappGroupWebhookField**](WhatsappGroupWebhookField.md) |  | [optional] 
**Type** | Pointer to [**WhatsappGroupWebhookType**](WhatsappGroupWebhookType.md) |  | [optional] 
**RequestId** | Pointer to **string** | The request ID returned by an asynchronous group API operation. | [optional] 
**Status** | Pointer to [**WhatsappGroupWebhookStatus**](WhatsappGroupWebhookStatus.md) |  | [optional] 
**GroupId** | Pointer to **string** | WhatsApp group ID. | [optional] 
**InviteLink** | Pointer to **string** | The group invite link. | [optional] 
**Reason** | Pointer to **string** | The reason for a participant, join request, or removal event. | [optional] 
**InitiatedBy** | Pointer to **string** | Indicates who initiated a participant removal event. | [optional] 
**JoinRequestId** | Pointer to **string** | The join request ID. | [optional] 
**WaId** | Pointer to **string** | WhatsApp user ID for a single participant event. | [optional] 
**RecipientUserId** | Pointer to **string** | Business-scoped user ID for a single participant event. | [optional] 
**ParentRecipientUserId** | Pointer to **string** | Parent business-scoped user ID for a single participant event. | [optional] 
**CustomerProfile** | Pointer to [**WhatsappGroupCustomerProfile**](WhatsappGroupCustomerProfile.md) |  | [optional] 
**Subject** | Pointer to **string** | The group subject. | [optional] 
**Description** | Pointer to **string** | The group description. | [optional] 
**JoinApprovalMode** | Pointer to [**WhatsappGroupJoinApprovalMode**](WhatsappGroupJoinApprovalMode.md) |  | [optional] 
**AddedParticipants** | Pointer to [**[]WhatsappGroupWebhookParticipant**](WhatsappGroupWebhookParticipant.md) | Participants added to the group. | [optional] 
**RemovedParticipants** | Pointer to [**[]WhatsappGroupWebhookParticipant**](WhatsappGroupWebhookParticipant.md) | Participants removed from the group. | [optional] 
**FailedParticipants** | Pointer to [**[]WhatsappGroupWebhookParticipant**](WhatsappGroupWebhookParticipant.md) | Participants that failed to be added or removed. | [optional] 
**Settings** | Pointer to [**[]WhatsappGroupWebhookSetting**](WhatsappGroupWebhookSetting.md) | Group setting update details. | [optional] 
**Errors** | Pointer to **[]map[string]interface{}** | Errors returned by WhatsApp. | [optional] 
**Contacts** | Pointer to [**[]WhatsappGroupWebhookStatusContact**](WhatsappGroupWebhookStatusContact.md) | Contacts included in group message status webhooks. | [optional] 
**Statuses** | Pointer to [**[]WhatsappGroupWebhookMessageStatus**](WhatsappGroupWebhookMessageStatus.md) | Group message status details. | [optional] 
**WebhookTime** | Pointer to **time.Time** | The time at which WhatsApp triggered this webhook. | [optional] 
**DedupeKey** | Pointer to **string** | Idempotency key for deduplicating group webhook events. | [optional] 

## Methods

### NewWhatsappGroupWebhook

`func NewWhatsappGroupWebhook() *WhatsappGroupWebhook`

NewWhatsappGroupWebhook instantiates a new WhatsappGroupWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappGroupWebhookWithDefaults

`func NewWhatsappGroupWebhookWithDefaults() *WhatsappGroupWebhook`

NewWhatsappGroupWebhookWithDefaults instantiates a new WhatsappGroupWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWabaId

`func (o *WhatsappGroupWebhook) GetWabaId() string`

GetWabaId returns the WabaId field if non-nil, zero value otherwise.

### GetWabaIdOk

`func (o *WhatsappGroupWebhook) GetWabaIdOk() (*string, bool)`

GetWabaIdOk returns a tuple with the WabaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWabaId

`func (o *WhatsappGroupWebhook) SetWabaId(v string)`

SetWabaId sets WabaId field to given value.

### HasWabaId

`func (o *WhatsappGroupWebhook) HasWabaId() bool`

HasWabaId returns a boolean if a field has been set.

### GetDisplayPhoneNumber

`func (o *WhatsappGroupWebhook) GetDisplayPhoneNumber() string`

GetDisplayPhoneNumber returns the DisplayPhoneNumber field if non-nil, zero value otherwise.

### GetDisplayPhoneNumberOk

`func (o *WhatsappGroupWebhook) GetDisplayPhoneNumberOk() (*string, bool)`

GetDisplayPhoneNumberOk returns a tuple with the DisplayPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayPhoneNumber

`func (o *WhatsappGroupWebhook) SetDisplayPhoneNumber(v string)`

SetDisplayPhoneNumber sets DisplayPhoneNumber field to given value.

### HasDisplayPhoneNumber

`func (o *WhatsappGroupWebhook) HasDisplayPhoneNumber() bool`

HasDisplayPhoneNumber returns a boolean if a field has been set.

### GetPhoneNumberId

`func (o *WhatsappGroupWebhook) GetPhoneNumberId() string`

GetPhoneNumberId returns the PhoneNumberId field if non-nil, zero value otherwise.

### GetPhoneNumberIdOk

`func (o *WhatsappGroupWebhook) GetPhoneNumberIdOk() (*string, bool)`

GetPhoneNumberIdOk returns a tuple with the PhoneNumberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberId

`func (o *WhatsappGroupWebhook) SetPhoneNumberId(v string)`

SetPhoneNumberId sets PhoneNumberId field to given value.

### HasPhoneNumberId

`func (o *WhatsappGroupWebhook) HasPhoneNumberId() bool`

HasPhoneNumberId returns a boolean if a field has been set.

### GetField

`func (o *WhatsappGroupWebhook) GetField() WhatsappGroupWebhookField`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *WhatsappGroupWebhook) GetFieldOk() (*WhatsappGroupWebhookField, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *WhatsappGroupWebhook) SetField(v WhatsappGroupWebhookField)`

SetField sets Field field to given value.

### HasField

`func (o *WhatsappGroupWebhook) HasField() bool`

HasField returns a boolean if a field has been set.

### GetType

`func (o *WhatsappGroupWebhook) GetType() WhatsappGroupWebhookType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WhatsappGroupWebhook) GetTypeOk() (*WhatsappGroupWebhookType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WhatsappGroupWebhook) SetType(v WhatsappGroupWebhookType)`

SetType sets Type field to given value.

### HasType

`func (o *WhatsappGroupWebhook) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRequestId

`func (o *WhatsappGroupWebhook) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *WhatsappGroupWebhook) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *WhatsappGroupWebhook) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *WhatsappGroupWebhook) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetStatus

`func (o *WhatsappGroupWebhook) GetStatus() WhatsappGroupWebhookStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WhatsappGroupWebhook) GetStatusOk() (*WhatsappGroupWebhookStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WhatsappGroupWebhook) SetStatus(v WhatsappGroupWebhookStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WhatsappGroupWebhook) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetGroupId

`func (o *WhatsappGroupWebhook) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *WhatsappGroupWebhook) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *WhatsappGroupWebhook) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *WhatsappGroupWebhook) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetInviteLink

`func (o *WhatsappGroupWebhook) GetInviteLink() string`

GetInviteLink returns the InviteLink field if non-nil, zero value otherwise.

### GetInviteLinkOk

`func (o *WhatsappGroupWebhook) GetInviteLinkOk() (*string, bool)`

GetInviteLinkOk returns a tuple with the InviteLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteLink

`func (o *WhatsappGroupWebhook) SetInviteLink(v string)`

SetInviteLink sets InviteLink field to given value.

### HasInviteLink

`func (o *WhatsappGroupWebhook) HasInviteLink() bool`

HasInviteLink returns a boolean if a field has been set.

### GetReason

`func (o *WhatsappGroupWebhook) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *WhatsappGroupWebhook) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *WhatsappGroupWebhook) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *WhatsappGroupWebhook) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetInitiatedBy

`func (o *WhatsappGroupWebhook) GetInitiatedBy() string`

GetInitiatedBy returns the InitiatedBy field if non-nil, zero value otherwise.

### GetInitiatedByOk

`func (o *WhatsappGroupWebhook) GetInitiatedByOk() (*string, bool)`

GetInitiatedByOk returns a tuple with the InitiatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatedBy

`func (o *WhatsappGroupWebhook) SetInitiatedBy(v string)`

SetInitiatedBy sets InitiatedBy field to given value.

### HasInitiatedBy

`func (o *WhatsappGroupWebhook) HasInitiatedBy() bool`

HasInitiatedBy returns a boolean if a field has been set.

### GetJoinRequestId

`func (o *WhatsappGroupWebhook) GetJoinRequestId() string`

GetJoinRequestId returns the JoinRequestId field if non-nil, zero value otherwise.

### GetJoinRequestIdOk

`func (o *WhatsappGroupWebhook) GetJoinRequestIdOk() (*string, bool)`

GetJoinRequestIdOk returns a tuple with the JoinRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinRequestId

`func (o *WhatsappGroupWebhook) SetJoinRequestId(v string)`

SetJoinRequestId sets JoinRequestId field to given value.

### HasJoinRequestId

`func (o *WhatsappGroupWebhook) HasJoinRequestId() bool`

HasJoinRequestId returns a boolean if a field has been set.

### GetWaId

`func (o *WhatsappGroupWebhook) GetWaId() string`

GetWaId returns the WaId field if non-nil, zero value otherwise.

### GetWaIdOk

`func (o *WhatsappGroupWebhook) GetWaIdOk() (*string, bool)`

GetWaIdOk returns a tuple with the WaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaId

`func (o *WhatsappGroupWebhook) SetWaId(v string)`

SetWaId sets WaId field to given value.

### HasWaId

`func (o *WhatsappGroupWebhook) HasWaId() bool`

HasWaId returns a boolean if a field has been set.

### GetRecipientUserId

`func (o *WhatsappGroupWebhook) GetRecipientUserId() string`

GetRecipientUserId returns the RecipientUserId field if non-nil, zero value otherwise.

### GetRecipientUserIdOk

`func (o *WhatsappGroupWebhook) GetRecipientUserIdOk() (*string, bool)`

GetRecipientUserIdOk returns a tuple with the RecipientUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientUserId

`func (o *WhatsappGroupWebhook) SetRecipientUserId(v string)`

SetRecipientUserId sets RecipientUserId field to given value.

### HasRecipientUserId

`func (o *WhatsappGroupWebhook) HasRecipientUserId() bool`

HasRecipientUserId returns a boolean if a field has been set.

### GetParentRecipientUserId

`func (o *WhatsappGroupWebhook) GetParentRecipientUserId() string`

GetParentRecipientUserId returns the ParentRecipientUserId field if non-nil, zero value otherwise.

### GetParentRecipientUserIdOk

`func (o *WhatsappGroupWebhook) GetParentRecipientUserIdOk() (*string, bool)`

GetParentRecipientUserIdOk returns a tuple with the ParentRecipientUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRecipientUserId

`func (o *WhatsappGroupWebhook) SetParentRecipientUserId(v string)`

SetParentRecipientUserId sets ParentRecipientUserId field to given value.

### HasParentRecipientUserId

`func (o *WhatsappGroupWebhook) HasParentRecipientUserId() bool`

HasParentRecipientUserId returns a boolean if a field has been set.

### GetCustomerProfile

`func (o *WhatsappGroupWebhook) GetCustomerProfile() WhatsappGroupCustomerProfile`

GetCustomerProfile returns the CustomerProfile field if non-nil, zero value otherwise.

### GetCustomerProfileOk

`func (o *WhatsappGroupWebhook) GetCustomerProfileOk() (*WhatsappGroupCustomerProfile, bool)`

GetCustomerProfileOk returns a tuple with the CustomerProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerProfile

`func (o *WhatsappGroupWebhook) SetCustomerProfile(v WhatsappGroupCustomerProfile)`

SetCustomerProfile sets CustomerProfile field to given value.

### HasCustomerProfile

`func (o *WhatsappGroupWebhook) HasCustomerProfile() bool`

HasCustomerProfile returns a boolean if a field has been set.

### GetSubject

`func (o *WhatsappGroupWebhook) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *WhatsappGroupWebhook) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *WhatsappGroupWebhook) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *WhatsappGroupWebhook) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetDescription

`func (o *WhatsappGroupWebhook) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WhatsappGroupWebhook) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WhatsappGroupWebhook) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WhatsappGroupWebhook) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJoinApprovalMode

`func (o *WhatsappGroupWebhook) GetJoinApprovalMode() WhatsappGroupJoinApprovalMode`

GetJoinApprovalMode returns the JoinApprovalMode field if non-nil, zero value otherwise.

### GetJoinApprovalModeOk

`func (o *WhatsappGroupWebhook) GetJoinApprovalModeOk() (*WhatsappGroupJoinApprovalMode, bool)`

GetJoinApprovalModeOk returns a tuple with the JoinApprovalMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinApprovalMode

`func (o *WhatsappGroupWebhook) SetJoinApprovalMode(v WhatsappGroupJoinApprovalMode)`

SetJoinApprovalMode sets JoinApprovalMode field to given value.

### HasJoinApprovalMode

`func (o *WhatsappGroupWebhook) HasJoinApprovalMode() bool`

HasJoinApprovalMode returns a boolean if a field has been set.

### GetAddedParticipants

`func (o *WhatsappGroupWebhook) GetAddedParticipants() []WhatsappGroupWebhookParticipant`

GetAddedParticipants returns the AddedParticipants field if non-nil, zero value otherwise.

### GetAddedParticipantsOk

`func (o *WhatsappGroupWebhook) GetAddedParticipantsOk() (*[]WhatsappGroupWebhookParticipant, bool)`

GetAddedParticipantsOk returns a tuple with the AddedParticipants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedParticipants

`func (o *WhatsappGroupWebhook) SetAddedParticipants(v []WhatsappGroupWebhookParticipant)`

SetAddedParticipants sets AddedParticipants field to given value.

### HasAddedParticipants

`func (o *WhatsappGroupWebhook) HasAddedParticipants() bool`

HasAddedParticipants returns a boolean if a field has been set.

### GetRemovedParticipants

`func (o *WhatsappGroupWebhook) GetRemovedParticipants() []WhatsappGroupWebhookParticipant`

GetRemovedParticipants returns the RemovedParticipants field if non-nil, zero value otherwise.

### GetRemovedParticipantsOk

`func (o *WhatsappGroupWebhook) GetRemovedParticipantsOk() (*[]WhatsappGroupWebhookParticipant, bool)`

GetRemovedParticipantsOk returns a tuple with the RemovedParticipants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedParticipants

`func (o *WhatsappGroupWebhook) SetRemovedParticipants(v []WhatsappGroupWebhookParticipant)`

SetRemovedParticipants sets RemovedParticipants field to given value.

### HasRemovedParticipants

`func (o *WhatsappGroupWebhook) HasRemovedParticipants() bool`

HasRemovedParticipants returns a boolean if a field has been set.

### GetFailedParticipants

`func (o *WhatsappGroupWebhook) GetFailedParticipants() []WhatsappGroupWebhookParticipant`

GetFailedParticipants returns the FailedParticipants field if non-nil, zero value otherwise.

### GetFailedParticipantsOk

`func (o *WhatsappGroupWebhook) GetFailedParticipantsOk() (*[]WhatsappGroupWebhookParticipant, bool)`

GetFailedParticipantsOk returns a tuple with the FailedParticipants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedParticipants

`func (o *WhatsappGroupWebhook) SetFailedParticipants(v []WhatsappGroupWebhookParticipant)`

SetFailedParticipants sets FailedParticipants field to given value.

### HasFailedParticipants

`func (o *WhatsappGroupWebhook) HasFailedParticipants() bool`

HasFailedParticipants returns a boolean if a field has been set.

### GetSettings

`func (o *WhatsappGroupWebhook) GetSettings() []WhatsappGroupWebhookSetting`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *WhatsappGroupWebhook) GetSettingsOk() (*[]WhatsappGroupWebhookSetting, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *WhatsappGroupWebhook) SetSettings(v []WhatsappGroupWebhookSetting)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *WhatsappGroupWebhook) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetErrors

`func (o *WhatsappGroupWebhook) GetErrors() []map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *WhatsappGroupWebhook) GetErrorsOk() (*[]map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *WhatsappGroupWebhook) SetErrors(v []map[string]interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *WhatsappGroupWebhook) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetContacts

`func (o *WhatsappGroupWebhook) GetContacts() []WhatsappGroupWebhookStatusContact`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *WhatsappGroupWebhook) GetContactsOk() (*[]WhatsappGroupWebhookStatusContact, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *WhatsappGroupWebhook) SetContacts(v []WhatsappGroupWebhookStatusContact)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *WhatsappGroupWebhook) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### GetStatuses

`func (o *WhatsappGroupWebhook) GetStatuses() []WhatsappGroupWebhookMessageStatus`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *WhatsappGroupWebhook) GetStatusesOk() (*[]WhatsappGroupWebhookMessageStatus, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *WhatsappGroupWebhook) SetStatuses(v []WhatsappGroupWebhookMessageStatus)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *WhatsappGroupWebhook) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### GetWebhookTime

`func (o *WhatsappGroupWebhook) GetWebhookTime() time.Time`

GetWebhookTime returns the WebhookTime field if non-nil, zero value otherwise.

### GetWebhookTimeOk

`func (o *WhatsappGroupWebhook) GetWebhookTimeOk() (*time.Time, bool)`

GetWebhookTimeOk returns a tuple with the WebhookTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookTime

`func (o *WhatsappGroupWebhook) SetWebhookTime(v time.Time)`

SetWebhookTime sets WebhookTime field to given value.

### HasWebhookTime

`func (o *WhatsappGroupWebhook) HasWebhookTime() bool`

HasWebhookTime returns a boolean if a field has been set.

### GetDedupeKey

`func (o *WhatsappGroupWebhook) GetDedupeKey() string`

GetDedupeKey returns the DedupeKey field if non-nil, zero value otherwise.

### GetDedupeKeyOk

`func (o *WhatsappGroupWebhook) GetDedupeKeyOk() (*string, bool)`

GetDedupeKeyOk returns a tuple with the DedupeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedupeKey

`func (o *WhatsappGroupWebhook) SetDedupeKey(v string)`

SetDedupeKey sets DedupeKey field to given value.

### HasDedupeKey

`func (o *WhatsappGroupWebhook) HasDedupeKey() bool`

HasDedupeKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


