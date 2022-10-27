# YCloud SDK for Go

The [YCloud](https://ycloud.com) API is organized around [REST](https://en.wikipedia.org/wiki/Representational_state_transfer). Our API is designed to have predictable, resource-oriented URLs, return [JSON](https://www.json.org) responses, and use standard HTTP response codes and verbs.

## Overview

- API version: v2
- Package version: 1.2.0

## Installation

Install the following dependencies:

```shell
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/ycloud-cpaas/ycloud-sdk-go
```

Add the following in import:

```golang
import ycloud "github.com/ycloud-cpaas/ycloud-sdk-go"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.ycloud.com/v2*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*BalanceApi* | [**Retrieve**](docs/BalanceApi.md#retrieve) | **Get** /balance | Retrieve balance
*EmailsApi* | [**Send**](docs/EmailsApi.md#send) | **Post** /emails | Send an email
*SmsApi* | [**List**](docs/SmsApi.md#list) | **Get** /sms | List SMS records
*SmsApi* | [**Send**](docs/SmsApi.md#send) | **Post** /sms | Send an SMS
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
*WhatsappMessagesApi* | [**Retrieve**](docs/WhatsappMessagesApi.md#retrieve) | **Get** /whatsapp/messages/{id} | Retrieve a WhatsApp message
*WhatsappMessagesApi* | [**Send**](docs/WhatsappMessagesApi.md#send) | **Post** /whatsapp/messages | Send a WhatsApp message
*WhatsappPhoneNumbersApi* | [**List**](docs/WhatsappPhoneNumbersApi.md#list) | **Get** /whatsapp/phoneNumbers | List WhatsApp phone numbers
*WhatsappPhoneNumbersApi* | [**Retrieve**](docs/WhatsappPhoneNumbersApi.md#retrieve) | **Get** /whatsapp/phoneNumbers/{wabaId}/{phoneNumber} | Retrieve a WhatsApp phone number
*WhatsappTemplatesApi* | [**Create**](docs/WhatsappTemplatesApi.md#create) | **Post** /whatsapp/templates | Create a WhatsApp template
*WhatsappTemplatesApi* | [**DeleteByName**](docs/WhatsappTemplatesApi.md#deletebyname) | **Delete** /whatsapp/templates/{wabaId}/{name} | Delete WhatsApp templates by name
*WhatsappTemplatesApi* | [**List**](docs/WhatsappTemplatesApi.md#list) | **Get** /whatsapp/templates | List WhatsApp templates
*WhatsappTemplatesApi* | [**RetrieveByNameAndLanguage**](docs/WhatsappTemplatesApi.md#retrievebynameandlanguage) | **Get** /whatsapp/templates/{wabaId}/{name}/{language} | Retrieve a WhatsApp template


## Documentation For Models

 - [Balance](docs/Balance.md)
 - [Email](docs/Email.md)
 - [EmailContentType](docs/EmailContentType.md)
 - [EmailDelivery](docs/EmailDelivery.md)
 - [EmailSendRequest](docs/EmailSendRequest.md)
 - [Error](docs/Error.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [Event](docs/Event.md)
 - [EventType](docs/EventType.md)
 - [Mailbox](docs/Mailbox.md)
 - [Page](docs/Page.md)
 - [Sms](docs/Sms.md)
 - [SmsPage](docs/SmsPage.md)
 - [SmsSendRequest](docs/SmsSendRequest.md)
 - [Verification](docs/Verification.md)
 - [VerificationChannel](docs/VerificationChannel.md)
 - [VerificationCheck](docs/VerificationCheck.md)
 - [VerificationCheckRequest](docs/VerificationCheckRequest.md)
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
 - [WhatsappInboundMessage](docs/WhatsappInboundMessage.md)
 - [WhatsappInboundMessageButton](docs/WhatsappInboundMessageButton.md)
 - [WhatsappInboundMessageInteractive](docs/WhatsappInboundMessageInteractive.md)
 - [WhatsappInboundMessageInteractiveButtonReply](docs/WhatsappInboundMessageInteractiveButtonReply.md)
 - [WhatsappInboundMessageInteractiveListReply](docs/WhatsappInboundMessageInteractiveListReply.md)
 - [WhatsappInboundMessageLocation](docs/WhatsappInboundMessageLocation.md)
 - [WhatsappInboundMessageMedia](docs/WhatsappInboundMessageMedia.md)
 - [WhatsappInboundMessageText](docs/WhatsappInboundMessageText.md)
 - [WhatsappInboundMessageType](docs/WhatsappInboundMessageType.md)
 - [WhatsappMessage](docs/WhatsappMessage.md)
 - [WhatsappMessageContact](docs/WhatsappMessageContact.md)
 - [WhatsappMessageContactAddressesInner](docs/WhatsappMessageContactAddressesInner.md)
 - [WhatsappMessageContactEmailsInner](docs/WhatsappMessageContactEmailsInner.md)
 - [WhatsappMessageContactName](docs/WhatsappMessageContactName.md)
 - [WhatsappMessageContactOrg](docs/WhatsappMessageContactOrg.md)
 - [WhatsappMessageContactPhonesInner](docs/WhatsappMessageContactPhonesInner.md)
 - [WhatsappMessageContactUrlsInner](docs/WhatsappMessageContactUrlsInner.md)
 - [WhatsappMessageInteractive](docs/WhatsappMessageInteractive.md)
 - [WhatsappMessageInteractiveAction](docs/WhatsappMessageInteractiveAction.md)
 - [WhatsappMessageInteractiveActionButtonsInner](docs/WhatsappMessageInteractiveActionButtonsInner.md)
 - [WhatsappMessageInteractiveActionButtonsInnerReply](docs/WhatsappMessageInteractiveActionButtonsInnerReply.md)
 - [WhatsappMessageInteractiveActionSectionsInner](docs/WhatsappMessageInteractiveActionSectionsInner.md)
 - [WhatsappMessageInteractiveActionSectionsInnerRowsInner](docs/WhatsappMessageInteractiveActionSectionsInnerRowsInner.md)
 - [WhatsappMessageInteractiveBody](docs/WhatsappMessageInteractiveBody.md)
 - [WhatsappMessageInteractiveFooter](docs/WhatsappMessageInteractiveFooter.md)
 - [WhatsappMessageInteractiveHeader](docs/WhatsappMessageInteractiveHeader.md)
 - [WhatsappMessageLocation](docs/WhatsappMessageLocation.md)
 - [WhatsappMessageMedia](docs/WhatsappMessageMedia.md)
 - [WhatsappMessageSendRequest](docs/WhatsappMessageSendRequest.md)
 - [WhatsappMessageStatus](docs/WhatsappMessageStatus.md)
 - [WhatsappMessageTemplate](docs/WhatsappMessageTemplate.md)
 - [WhatsappMessageTemplateComponentsInner](docs/WhatsappMessageTemplateComponentsInner.md)
 - [WhatsappMessageTemplateComponentsInnerParametersInner](docs/WhatsappMessageTemplateComponentsInnerParametersInner.md)
 - [WhatsappMessageTemplateLanguage](docs/WhatsappMessageTemplateLanguage.md)
 - [WhatsappMessageText](docs/WhatsappMessageText.md)
 - [WhatsappMessageType](docs/WhatsappMessageType.md)
 - [WhatsappPhoneNumber](docs/WhatsappPhoneNumber.md)
 - [WhatsappPhoneNumberCodeVerificationStatus](docs/WhatsappPhoneNumberCodeVerificationStatus.md)
 - [WhatsappPhoneNumberNameStatus](docs/WhatsappPhoneNumberNameStatus.md)
 - [WhatsappPhoneNumberPage](docs/WhatsappPhoneNumberPage.md)
 - [WhatsappPhoneNumberQualityRating](docs/WhatsappPhoneNumberQualityRating.md)
 - [WhatsappPhoneNumberStatus](docs/WhatsappPhoneNumberStatus.md)
 - [WhatsappTemplate](docs/WhatsappTemplate.md)
 - [WhatsappTemplateCategory](docs/WhatsappTemplateCategory.md)
 - [WhatsappTemplateComponent](docs/WhatsappTemplateComponent.md)
 - [WhatsappTemplateComponentButton](docs/WhatsappTemplateComponentButton.md)
 - [WhatsappTemplateComponentButtonType](docs/WhatsappTemplateComponentButtonType.md)
 - [WhatsappTemplateComponentExample](docs/WhatsappTemplateComponentExample.md)
 - [WhatsappTemplateCreateRequest](docs/WhatsappTemplateCreateRequest.md)
 - [WhatsappTemplatePage](docs/WhatsappTemplatePage.md)
 - [WhatsappTemplateStatus](docs/WhatsappTemplateStatus.md)


## Documentation For Authorization



### api_key

- **Type**: API key
- **API key parameter name**: X-API-Key
- **Location**: HTTP header

Configure API key:
```golang
import ycloud "github.com/ycloud-cpaas/ycloud-sdk-go"

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
