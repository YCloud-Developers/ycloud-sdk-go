# YCloud SDK for Go

The [YCloud](https://ycloud.com) API is organized around [REST](https://en.wikipedia.org/wiki/Representational_state_transfer). Our API is designed to have predictable, resource-oriented URLs, return [JSON](https://www.json.org) responses, and use standard HTTP response codes and verbs.

## Overview

- API version: v2
- Package version: 1.15.0

## Installation

Install the following dependencies:

```shell
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/ycloud-developers/ycloud-sdk-go
```

Add the following in import:

```golang
import ycloud "github.com/ycloud-developers/ycloud-sdk-go"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.ycloud.com/v2*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*BalanceApi* | [**Retrieve**](docs/BalanceApi.md#retrieve) | **Get** /balance | Retrieve balance
*ContactsApi* | [**Create**](docs/ContactsApi.md#create) | **Post** /contact/contacts | Create a contact
*ContactsApi* | [**Delete**](docs/ContactsApi.md#delete) | **Delete** /contact/contacts/{id} | Delete a contact
*ContactsApi* | [**List**](docs/ContactsApi.md#list) | **Get** /contact/contacts | List contacts
*ContactsApi* | [**Retrieve**](docs/ContactsApi.md#retrieve) | **Get** /contact/contacts/{id} | Retrieve a contact
*ContactsApi* | [**Update**](docs/ContactsApi.md#update) | **Patch** /contact/contacts/{id} | Update a contact
*CustomEventsApi* | [**CreateDefinition**](docs/CustomEventsApi.md#createdefinition) | **Post** /event/definitions | Create an event definition
*CustomEventsApi* | [**CreatePropertyDefinition**](docs/CustomEventsApi.md#createpropertydefinition) | **Post** /event/definitions/{name}/properties | Create an event property definition
*CustomEventsApi* | [**PropertyDefinition**](docs/CustomEventsApi.md#propertydefinition) | **Delete** /event/definitions/{name}/properties/{propertyName} | Delete an event property definition
*CustomEventsApi* | [**PropertyDefinition_0**](docs/CustomEventsApi.md#propertydefinition_0) | **Patch** /event/definitions/{name}/properties/{propertyName} | Update an event property definition
*CustomEventsApi* | [**RetrieveDefinition**](docs/CustomEventsApi.md#retrievedefinition) | **Get** /event/definitions/{name} | Retrieve an event definition
*CustomEventsApi* | [**SendEvent**](docs/CustomEventsApi.md#sendevent) | **Post** /event/events | Send an event
*CustomEventsApi* | [**UpdateDefinition**](docs/CustomEventsApi.md#updatedefinition) | **Patch** /event/definitions/{name} | Update an event definition
*EmailsApi* | [**Send**](docs/EmailsApi.md#send) | **Post** /emails | Send an email
*SmsApi* | [**List**](docs/SmsApi.md#list) | **Get** /sms | List SMS records
*SmsApi* | [**Send**](docs/SmsApi.md#send) | **Post** /sms | Send an SMS
*UnsubscribersApi* | [**Create**](docs/UnsubscribersApi.md#create) | **Post** /unsubscribers | Create an unsubscriber
*UnsubscribersApi* | [**DeleteByCustomerAndChannel**](docs/UnsubscribersApi.md#deletebycustomerandchannel) | **Delete** /unsubscribers/{customer}/{channel} | Delete an unsubscriber
*UnsubscribersApi* | [**List**](docs/UnsubscribersApi.md#list) | **Get** /unsubscribers | List unsubscribers
*UnsubscribersApi* | [**ListAllByCustomer**](docs/UnsubscribersApi.md#listallbycustomer) | **Get** /unsubscribers/{customer} | List all unsubscribers by customer
*UnsubscribersApi* | [**RetrieveByCustomerAndChannel**](docs/UnsubscribersApi.md#retrievebycustomerandchannel) | **Get** /unsubscribers/{customer}/{channel} | Retrieve an unsubscriber
*VerifyApi* | [**Check**](docs/VerifyApi.md#check) | **Post** /verify/verificationChecks | Check a verification
*VerifyApi* | [**Send**](docs/VerifyApi.md#send) | **Post** /verify/verifications | Start a verification
*VoicesApi* | [**List**](docs/VoicesApi.md#list) | **Get** /voices | List voice records
*VoicesApi* | [**Send**](docs/VoicesApi.md#send) | **Post** /voices | Send a voice code
*WebhookEndpointsApi* | [**Create**](docs/WebhookEndpointsApi.md#create) | **Post** /webhookEndpoints | Create a webhook endpoint
*WebhookEndpointsApi* | [**Delete**](docs/WebhookEndpointsApi.md#delete) | **Delete** /webhookEndpoints/{id} | Delete a webhook endpoint
*WebhookEndpointsApi* | [**List**](docs/WebhookEndpointsApi.md#list) | **Get** /webhookEndpoints | List webhook endpoints
*WebhookEndpointsApi* | [**Retrieve**](docs/WebhookEndpointsApi.md#retrieve) | **Get** /webhookEndpoints/{id} | Retrieve a webhook endpoint
*WebhookEndpointsApi* | [**RotateSecret**](docs/WebhookEndpointsApi.md#rotatesecret) | **Post** /webhookEndpoints/{id}/rotateSecret | Rotate a webhook endpoint secret
*WebhookEndpointsApi* | [**Update**](docs/WebhookEndpointsApi.md#update) | **Patch** /webhookEndpoints/{id} | Update a webhook endpoint
*WhatsappBusinessAccountsApi* | [**List**](docs/WhatsappBusinessAccountsApi.md#list) | **Get** /whatsapp/businessAccounts | List WABAs
*WhatsappBusinessAccountsApi* | [**Retrieve**](docs/WhatsappBusinessAccountsApi.md#retrieve) | **Get** /whatsapp/businessAccounts/{id} | Retrieve a WABA
*WhatsappInboundMessagesApi* | [**MarkAsRead**](docs/WhatsappInboundMessagesApi.md#markasread) | **Post** /whatsapp/inboundMessages/{id}/markAsRead | Mark message as read
*WhatsappMessagesApi* | [**Retrieve**](docs/WhatsappMessagesApi.md#retrieve) | **Get** /whatsapp/messages/{id} | Retrieve a message
*WhatsappMessagesApi* | [**Send**](docs/WhatsappMessagesApi.md#send) | **Post** /whatsapp/messages | Enqueue a message
*WhatsappMessagesApi* | [**SendDirectly**](docs/WhatsappMessagesApi.md#senddirectly) | **Post** /whatsapp/messages/sendDirectly | Send a message directly
*WhatsappPhoneNumbersApi* | [**List**](docs/WhatsappPhoneNumbersApi.md#list) | **Get** /whatsapp/phoneNumbers | List phone numbers
*WhatsappPhoneNumbersApi* | [**Register**](docs/WhatsappPhoneNumbersApi.md#register) | **Post** /whatsapp/phoneNumbers/{wabaId}/{phoneNumber}/register | Register a phone number
*WhatsappPhoneNumbersApi* | [**Retrieve**](docs/WhatsappPhoneNumbersApi.md#retrieve) | **Get** /whatsapp/phoneNumbers/{wabaId}/{phoneNumber} | Retrieve a phone number
*WhatsappPhoneNumbersApi* | [**RetrieveCommerceSettings**](docs/WhatsappPhoneNumbersApi.md#retrievecommercesettings) | **Get** /whatsapp/phoneNumbers/{wabaId}/{phoneNumber}/whatsappCommerceSettings | Retrieve commerce settings
*WhatsappPhoneNumbersApi* | [**RetrieveProfile**](docs/WhatsappPhoneNumbersApi.md#retrieveprofile) | **Get** /whatsapp/phoneNumbers/{wabaId}/{phoneNumber}/profile | Retrieve a phone number profile
*WhatsappPhoneNumbersApi* | [**UpdateCommerceSettings**](docs/WhatsappPhoneNumbersApi.md#updatecommercesettings) | **Patch** /whatsapp/phoneNumbers/{wabaId}/{phoneNumber}/whatsappCommerceSettings | Update commerce settings
*WhatsappPhoneNumbersApi* | [**UpdateProfile**](docs/WhatsappPhoneNumbersApi.md#updateprofile) | **Patch** /whatsapp/phoneNumbers/{wabaId}/{phoneNumber}/profile | Update a phone number profile
*WhatsappTemplatesApi* | [**Create**](docs/WhatsappTemplatesApi.md#create) | **Post** /whatsapp/templates | Create a template
*WhatsappTemplatesApi* | [**DeleteByName**](docs/WhatsappTemplatesApi.md#deletebyname) | **Delete** /whatsapp/templates/{wabaId}/{name} | Delete templates by name
*WhatsappTemplatesApi* | [**DeleteByNameAndLanguage**](docs/WhatsappTemplatesApi.md#deletebynameandlanguage) | **Delete** /whatsapp/templates/{wabaId}/{name}/{language} | Delete a template
*WhatsappTemplatesApi* | [**EditByNameAndLanguage**](docs/WhatsappTemplatesApi.md#editbynameandlanguage) | **Patch** /whatsapp/templates/{wabaId}/{name}/{language} | Edit a template
*WhatsappTemplatesApi* | [**List**](docs/WhatsappTemplatesApi.md#list) | **Get** /whatsapp/templates | List templates
*WhatsappTemplatesApi* | [**RetrieveByNameAndLanguage**](docs/WhatsappTemplatesApi.md#retrievebynameandlanguage) | **Get** /whatsapp/templates/{wabaId}/{name}/{language} | Retrieve a template


## Documentation For Models

 - [Balance](docs/Balance.md)
 - [Contact](docs/Contact.md)
 - [ContactCreateRequest](docs/ContactCreateRequest.md)
 - [ContactCustomAttribute](docs/ContactCustomAttribute.md)
 - [ContactPage](docs/ContactPage.md)
 - [ContactUpdateRequest](docs/ContactUpdateRequest.md)
 - [CustomEventDefinition](docs/CustomEventDefinition.md)
 - [CustomEventDefinitionCreateRequest](docs/CustomEventDefinitionCreateRequest.md)
 - [CustomEventDefinitionProperty](docs/CustomEventDefinitionProperty.md)
 - [CustomEventDefinitionPropertyCreateRequest](docs/CustomEventDefinitionPropertyCreateRequest.md)
 - [CustomEventDefinitionPropertyUpdateRequest](docs/CustomEventDefinitionPropertyUpdateRequest.md)
 - [CustomEventDefinitionUpdateRequest](docs/CustomEventDefinitionUpdateRequest.md)
 - [CustomEventSendRequest](docs/CustomEventSendRequest.md)
 - [Email](docs/Email.md)
 - [EmailContentType](docs/EmailContentType.md)
 - [EmailDelivery](docs/EmailDelivery.md)
 - [EmailSendRequest](docs/EmailSendRequest.md)
 - [Error](docs/Error.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [Event](docs/Event.md)
 - [EventType](docs/EventType.md)
 - [Mailbox](docs/Mailbox.md)
 - [MetaBusinessAccountVerificationStatus](docs/MetaBusinessAccountVerificationStatus.md)
 - [Page](docs/Page.md)
 - [PageCursor](docs/PageCursor.md)
 - [Sms](docs/Sms.md)
 - [SmsInbound](docs/SmsInbound.md)
 - [SmsPage](docs/SmsPage.md)
 - [SmsSendRequest](docs/SmsSendRequest.md)
 - [Unsubscriber](docs/Unsubscriber.md)
 - [UnsubscriberChannel](docs/UnsubscriberChannel.md)
 - [UnsubscriberCreateRequest](docs/UnsubscriberCreateRequest.md)
 - [UnsubscriberPage](docs/UnsubscriberPage.md)
 - [UnsubscriberType](docs/UnsubscriberType.md)
 - [Verification](docs/Verification.md)
 - [VerificationChannel](docs/VerificationChannel.md)
 - [VerificationCheck](docs/VerificationCheck.md)
 - [VerificationCheckRequest](docs/VerificationCheckRequest.md)
 - [VerificationFallback](docs/VerificationFallback.md)
 - [VerificationSendRequest](docs/VerificationSendRequest.md)
 - [VerificationStatus](docs/VerificationStatus.md)
 - [Voice](docs/Voice.md)
 - [VoicePage](docs/VoicePage.md)
 - [VoiceSendRequest](docs/VoiceSendRequest.md)
 - [WebhookEndpoint](docs/WebhookEndpoint.md)
 - [WebhookEndpointCreateRequest](docs/WebhookEndpointCreateRequest.md)
 - [WebhookEndpointPage](docs/WebhookEndpointPage.md)
 - [WebhookEndpointStatus](docs/WebhookEndpointStatus.md)
 - [WebhookEndpointUpdateRequest](docs/WebhookEndpointUpdateRequest.md)
 - [WhatsappApiError](docs/WhatsappApiError.md)
 - [WhatsappAuthIntlRateEligibilityCountry](docs/WhatsappAuthIntlRateEligibilityCountry.md)
 - [WhatsappBusinessAccount](docs/WhatsappBusinessAccount.md)
 - [WhatsappBusinessAccountBanState](docs/WhatsappBusinessAccountBanState.md)
 - [WhatsappBusinessAccountPage](docs/WhatsappBusinessAccountPage.md)
 - [WhatsappBusinessAccountRestrictionInfo](docs/WhatsappBusinessAccountRestrictionInfo.md)
 - [WhatsappBusinessAccountReviewStatus](docs/WhatsappBusinessAccountReviewStatus.md)
 - [WhatsappBusinessAccountUpdateEventEnum](docs/WhatsappBusinessAccountUpdateEventEnum.md)
 - [WhatsappCommerceSettings](docs/WhatsappCommerceSettings.md)
 - [WhatsappCommerceSettingsUpdateRequest](docs/WhatsappCommerceSettingsUpdateRequest.md)
 - [WhatsappConversation](docs/WhatsappConversation.md)
 - [WhatsappConversationOriginType](docs/WhatsappConversationOriginType.md)
 - [WhatsappConversationType](docs/WhatsappConversationType.md)
 - [WhatsappInboundMessage](docs/WhatsappInboundMessage.md)
 - [WhatsappInboundMessageButton](docs/WhatsappInboundMessageButton.md)
 - [WhatsappInboundMessageContext](docs/WhatsappInboundMessageContext.md)
 - [WhatsappInboundMessageError](docs/WhatsappInboundMessageError.md)
 - [WhatsappInboundMessageInteractive](docs/WhatsappInboundMessageInteractive.md)
 - [WhatsappInboundMessageInteractiveButtonReply](docs/WhatsappInboundMessageInteractiveButtonReply.md)
 - [WhatsappInboundMessageInteractiveListReply](docs/WhatsappInboundMessageInteractiveListReply.md)
 - [WhatsappInboundMessageLocation](docs/WhatsappInboundMessageLocation.md)
 - [WhatsappInboundMessageMedia](docs/WhatsappInboundMessageMedia.md)
 - [WhatsappInboundMessageOrder](docs/WhatsappInboundMessageOrder.md)
 - [WhatsappInboundMessageOrderProductItem](docs/WhatsappInboundMessageOrderProductItem.md)
 - [WhatsappInboundMessageReferral](docs/WhatsappInboundMessageReferral.md)
 - [WhatsappInboundMessageReferredProduct](docs/WhatsappInboundMessageReferredProduct.md)
 - [WhatsappInboundMessageSystem](docs/WhatsappInboundMessageSystem.md)
 - [WhatsappInboundMessageText](docs/WhatsappInboundMessageText.md)
 - [WhatsappInboundMessageType](docs/WhatsappInboundMessageType.md)
 - [WhatsappMessage](docs/WhatsappMessage.md)
 - [WhatsappMessageContact](docs/WhatsappMessageContact.md)
 - [WhatsappMessageContactAddress](docs/WhatsappMessageContactAddress.md)
 - [WhatsappMessageContactEmail](docs/WhatsappMessageContactEmail.md)
 - [WhatsappMessageContactName](docs/WhatsappMessageContactName.md)
 - [WhatsappMessageContactOrg](docs/WhatsappMessageContactOrg.md)
 - [WhatsappMessageContactPhone](docs/WhatsappMessageContactPhone.md)
 - [WhatsappMessageContactUrl](docs/WhatsappMessageContactUrl.md)
 - [WhatsappMessageContext](docs/WhatsappMessageContext.md)
 - [WhatsappMessageInteractive](docs/WhatsappMessageInteractive.md)
 - [WhatsappMessageInteractiveAction](docs/WhatsappMessageInteractiveAction.md)
 - [WhatsappMessageInteractiveActionButton](docs/WhatsappMessageInteractiveActionButton.md)
 - [WhatsappMessageInteractiveActionButtonReply](docs/WhatsappMessageInteractiveActionButtonReply.md)
 - [WhatsappMessageInteractiveActionParameters](docs/WhatsappMessageInteractiveActionParameters.md)
 - [WhatsappMessageInteractiveActionParametersFlowActionPayload](docs/WhatsappMessageInteractiveActionParametersFlowActionPayload.md)
 - [WhatsappMessageInteractiveActionSection](docs/WhatsappMessageInteractiveActionSection.md)
 - [WhatsappMessageInteractiveActionSectionProductItem](docs/WhatsappMessageInteractiveActionSectionProductItem.md)
 - [WhatsappMessageInteractiveActionSectionRow](docs/WhatsappMessageInteractiveActionSectionRow.md)
 - [WhatsappMessageInteractiveBody](docs/WhatsappMessageInteractiveBody.md)
 - [WhatsappMessageInteractiveFooter](docs/WhatsappMessageInteractiveFooter.md)
 - [WhatsappMessageInteractiveHeader](docs/WhatsappMessageInteractiveHeader.md)
 - [WhatsappMessageLocation](docs/WhatsappMessageLocation.md)
 - [WhatsappMessageMedia](docs/WhatsappMessageMedia.md)
 - [WhatsappMessageOrderAmount](docs/WhatsappMessageOrderAmount.md)
 - [WhatsappMessageOrderBeneficiary](docs/WhatsappMessageOrderBeneficiary.md)
 - [WhatsappMessageOrderDetails](docs/WhatsappMessageOrderDetails.md)
 - [WhatsappMessageOrderExpiration](docs/WhatsappMessageOrderExpiration.md)
 - [WhatsappMessageOrderInfo](docs/WhatsappMessageOrderInfo.md)
 - [WhatsappMessageOrderItem](docs/WhatsappMessageOrderItem.md)
 - [WhatsappMessageOrderPaymentGateway](docs/WhatsappMessageOrderPaymentGateway.md)
 - [WhatsappMessageOrderPaymentSetting](docs/WhatsappMessageOrderPaymentSetting.md)
 - [WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk](docs/WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk.md)
 - [WhatsappMessageOrderPaymentSettingPaymentGatewayPayu](docs/WhatsappMessageOrderPaymentSettingPaymentGatewayPayu.md)
 - [WhatsappMessageOrderPaymentSettingPaymentGatewayRazorpay](docs/WhatsappMessageOrderPaymentSettingPaymentGatewayRazorpay.md)
 - [WhatsappMessageOrderPaymentSettingPaymentGatewayZaakpay](docs/WhatsappMessageOrderPaymentSettingPaymentGatewayZaakpay.md)
 - [WhatsappMessageOrderStatus](docs/WhatsappMessageOrderStatus.md)
 - [WhatsappMessageOrderStatusEnum](docs/WhatsappMessageOrderStatusEnum.md)
 - [WhatsappMessageReaction](docs/WhatsappMessageReaction.md)
 - [WhatsappMessageSendRequest](docs/WhatsappMessageSendRequest.md)
 - [WhatsappMessageStatus](docs/WhatsappMessageStatus.md)
 - [WhatsappMessageTemplate](docs/WhatsappMessageTemplate.md)
 - [WhatsappMessageTemplateComponent](docs/WhatsappMessageTemplateComponent.md)
 - [WhatsappMessageTemplateComponentCard](docs/WhatsappMessageTemplateComponentCard.md)
 - [WhatsappMessageTemplateComponentCardComponent](docs/WhatsappMessageTemplateComponentCardComponent.md)
 - [WhatsappMessageTemplateComponentParameter](docs/WhatsappMessageTemplateComponentParameter.md)
 - [WhatsappMessageTemplateComponentParameterAction](docs/WhatsappMessageTemplateComponentParameterAction.md)
 - [WhatsappMessageTemplateComponentParameterActionSection](docs/WhatsappMessageTemplateComponentParameterActionSection.md)
 - [WhatsappMessageTemplateComponentParameterActionSectionProductItem](docs/WhatsappMessageTemplateComponentParameterActionSectionProductItem.md)
 - [WhatsappMessageTemplateComponentParameterLimitedTimeOffer](docs/WhatsappMessageTemplateComponentParameterLimitedTimeOffer.md)
 - [WhatsappMessageTemplateLanguage](docs/WhatsappMessageTemplateLanguage.md)
 - [WhatsappMessageText](docs/WhatsappMessageText.md)
 - [WhatsappMessageType](docs/WhatsappMessageType.md)
 - [WhatsappPayment](docs/WhatsappPayment.md)
 - [WhatsappPaymentStatus](docs/WhatsappPaymentStatus.md)
 - [WhatsappPaymentTransaction](docs/WhatsappPaymentTransaction.md)
 - [WhatsappPaymentTransactionError](docs/WhatsappPaymentTransactionError.md)
 - [WhatsappPhoneNumber](docs/WhatsappPhoneNumber.md)
 - [WhatsappPhoneNumberCodeVerificationStatus](docs/WhatsappPhoneNumberCodeVerificationStatus.md)
 - [WhatsappPhoneNumberNameStatus](docs/WhatsappPhoneNumberNameStatus.md)
 - [WhatsappPhoneNumberPage](docs/WhatsappPhoneNumberPage.md)
 - [WhatsappPhoneNumberProfile](docs/WhatsappPhoneNumberProfile.md)
 - [WhatsappPhoneNumberProfileUpdateRequest](docs/WhatsappPhoneNumberProfileUpdateRequest.md)
 - [WhatsappPhoneNumberProfileVertical](docs/WhatsappPhoneNumberProfileVertical.md)
 - [WhatsappPhoneNumberQualityRating](docs/WhatsappPhoneNumberQualityRating.md)
 - [WhatsappPhoneNumberQualityUpdateEventEnum](docs/WhatsappPhoneNumberQualityUpdateEventEnum.md)
 - [WhatsappPhoneNumberStatus](docs/WhatsappPhoneNumberStatus.md)
 - [WhatsappPricingCategory](docs/WhatsappPricingCategory.md)
 - [WhatsappProfile](docs/WhatsappProfile.md)
 - [WhatsappReviewDecision](docs/WhatsappReviewDecision.md)
 - [WhatsappTemplate](docs/WhatsappTemplate.md)
 - [WhatsappTemplateCategory](docs/WhatsappTemplateCategory.md)
 - [WhatsappTemplateComponent](docs/WhatsappTemplateComponent.md)
 - [WhatsappTemplateComponentButton](docs/WhatsappTemplateComponentButton.md)
 - [WhatsappTemplateComponentButtonOtpType](docs/WhatsappTemplateComponentButtonOtpType.md)
 - [WhatsappTemplateComponentButtonType](docs/WhatsappTemplateComponentButtonType.md)
 - [WhatsappTemplateComponentCard](docs/WhatsappTemplateComponentCard.md)
 - [WhatsappTemplateComponentCardComponent](docs/WhatsappTemplateComponentCardComponent.md)
 - [WhatsappTemplateComponentExample](docs/WhatsappTemplateComponentExample.md)
 - [WhatsappTemplateComponentLimitedTimeOffer](docs/WhatsappTemplateComponentLimitedTimeOffer.md)
 - [WhatsappTemplateCreateRequest](docs/WhatsappTemplateCreateRequest.md)
 - [WhatsappTemplateEditRequest](docs/WhatsappTemplateEditRequest.md)
 - [WhatsappTemplatePage](docs/WhatsappTemplatePage.md)
 - [WhatsappTemplateQualityRating](docs/WhatsappTemplateQualityRating.md)
 - [WhatsappTemplateStatus](docs/WhatsappTemplateStatus.md)
 - [WhatsappTemplateStatusUpdateEventEnum](docs/WhatsappTemplateStatusUpdateEventEnum.md)
 - [WhatsappTemplateSubCategory](docs/WhatsappTemplateSubCategory.md)


## Documentation For Authorization



### api_key

- **Type**: API key
- **API key parameter name**: X-API-Key
- **Location**: HTTP header

Configure API key:
```golang
import ycloud "github.com/ycloud-developers/ycloud-sdk-go"

configuration := ycloud.NewConfiguration()
configuration.AddDefaultHeader("X-API-Key", "YOUR_API_KEY")
```

## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`
