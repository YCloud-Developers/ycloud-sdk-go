/*
YCloud API

The [YCloud](https://ycloud.com) API is organized around [REST](https://en.wikipedia.org/wiki/Representational_state_transfer). Our API is designed to have predictable, resource-oriented URLs, return [JSON](https://www.json.org) responses, and use standard HTTP response codes and verbs.

API version: v2
Contact: service@ycloud.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ycloud

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// CustomEventsApiService CustomEventsApi service
type CustomEventsApiService service

type CustomEventsApiCreateDefinitionRequest struct {
	ctx context.Context
	ApiService *CustomEventsApiService
	customEventDefinitionCreateRequest *CustomEventDefinitionCreateRequest
}

func (r CustomEventsApiCreateDefinitionRequest) CustomEventDefinitionCreateRequest(customEventDefinitionCreateRequest CustomEventDefinitionCreateRequest) CustomEventsApiCreateDefinitionRequest {
	r.customEventDefinitionCreateRequest = &customEventDefinitionCreateRequest
	return r
}

func (r CustomEventsApiCreateDefinitionRequest) Execute() (*CustomEventDefinition, *http.Response, error) {
	return r.ApiService.CreateDefinitionExecute(r)
}

/*
CreateDefinition Create an event definition

Creates a custom event definition.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return CustomEventsApiCreateDefinitionRequest
*/
func (a *CustomEventsApiService) CreateDefinition(ctx context.Context) CustomEventsApiCreateDefinitionRequest {
	return CustomEventsApiCreateDefinitionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CustomEventDefinition
func (a *CustomEventsApiService) CreateDefinitionExecute(r CustomEventsApiCreateDefinitionRequest) (*CustomEventDefinition, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CustomEventDefinition
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomEventsApiService.CreateDefinition")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/event/definitions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.customEventDefinitionCreateRequest == nil {
		return localVarReturnValue, nil, reportError("customEventDefinitionCreateRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.customEventDefinitionCreateRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type CustomEventsApiCreatePropertyDefinitionRequest struct {
	ctx context.Context
	ApiService *CustomEventsApiService
	name string
	customEventDefinitionPropertyCreateRequest *CustomEventDefinitionPropertyCreateRequest
}

func (r CustomEventsApiCreatePropertyDefinitionRequest) CustomEventDefinitionPropertyCreateRequest(customEventDefinitionPropertyCreateRequest CustomEventDefinitionPropertyCreateRequest) CustomEventsApiCreatePropertyDefinitionRequest {
	r.customEventDefinitionPropertyCreateRequest = &customEventDefinitionPropertyCreateRequest
	return r
}

func (r CustomEventsApiCreatePropertyDefinitionRequest) Execute() (*CustomEventDefinitionProperty, *http.Response, error) {
	return r.ApiService.CreatePropertyDefinitionExecute(r)
}

/*
CreatePropertyDefinition Create an event property definition

Defines a new property for the event definition.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name Name of the custom event.
 @return CustomEventsApiCreatePropertyDefinitionRequest
*/
func (a *CustomEventsApiService) CreatePropertyDefinition(ctx context.Context, name string) CustomEventsApiCreatePropertyDefinitionRequest {
	return CustomEventsApiCreatePropertyDefinitionRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
	}
}

// Execute executes the request
//  @return CustomEventDefinitionProperty
func (a *CustomEventsApiService) CreatePropertyDefinitionExecute(r CustomEventsApiCreatePropertyDefinitionRequest) (*CustomEventDefinitionProperty, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CustomEventDefinitionProperty
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomEventsApiService.CreatePropertyDefinition")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/event/definitions/{name}/properties"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterToString(r.name, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.customEventDefinitionPropertyCreateRequest == nil {
		return localVarReturnValue, nil, reportError("customEventDefinitionPropertyCreateRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.customEventDefinitionPropertyCreateRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type CustomEventsApiPropertyDefinitionRequest struct {
	ctx context.Context
	ApiService *CustomEventsApiService
	name string
	propertyName string
}

func (r CustomEventsApiPropertyDefinitionRequest) Execute() (*http.Response, error) {
	return r.ApiService.PropertyDefinitionExecute(r)
}

/*
PropertyDefinition Delete an event property definition

Deletes a property of the event definition.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name Name of the custom event.
 @param propertyName Name of the custom event property.
 @return CustomEventsApiPropertyDefinitionRequest
*/
func (a *CustomEventsApiService) PropertyDefinition(ctx context.Context, name string, propertyName string) CustomEventsApiPropertyDefinitionRequest {
	return CustomEventsApiPropertyDefinitionRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
		propertyName: propertyName,
	}
}

// Execute executes the request
func (a *CustomEventsApiService) PropertyDefinitionExecute(r CustomEventsApiPropertyDefinitionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomEventsApiService.PropertyDefinition")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/event/definitions/{name}/properties/{propertyName}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterToString(r.name, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"propertyName"+"}", url.PathEscape(parameterToString(r.propertyName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type CustomEventsApiPropertyDefinition_0Request struct {
	ctx context.Context
	ApiService *CustomEventsApiService
	name string
	propertyName string
	customEventDefinitionPropertyUpdateRequest *CustomEventDefinitionPropertyUpdateRequest
}

func (r CustomEventsApiPropertyDefinition_0Request) CustomEventDefinitionPropertyUpdateRequest(customEventDefinitionPropertyUpdateRequest CustomEventDefinitionPropertyUpdateRequest) CustomEventsApiPropertyDefinition_0Request {
	r.customEventDefinitionPropertyUpdateRequest = &customEventDefinitionPropertyUpdateRequest
	return r
}

func (r CustomEventsApiPropertyDefinition_0Request) Execute() (*CustomEventDefinitionProperty, *http.Response, error) {
	return r.ApiService.PropertyDefinition_1Execute(r)
}

/*
PropertyDefinition_0 Update an event property definition

Updates an event property definition's label and description.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name Name of the custom event.
 @param propertyName Name of the custom event property.
 @return CustomEventsApiPropertyDefinition_0Request
*/
func (a *CustomEventsApiService) PropertyDefinition_1(ctx context.Context, name string, propertyName string) CustomEventsApiPropertyDefinition_0Request {
	return CustomEventsApiPropertyDefinition_0Request{
		ApiService: a,
		ctx: ctx,
		name: name,
		propertyName: propertyName,
	}
}

// Execute executes the request
//  @return CustomEventDefinitionProperty
func (a *CustomEventsApiService) PropertyDefinition_1Execute(r CustomEventsApiPropertyDefinition_0Request) (*CustomEventDefinitionProperty, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CustomEventDefinitionProperty
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomEventsApiService.PropertyDefinition_1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/event/definitions/{name}/properties/{propertyName}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterToString(r.name, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"propertyName"+"}", url.PathEscape(parameterToString(r.propertyName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.customEventDefinitionPropertyUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("customEventDefinitionPropertyUpdateRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.customEventDefinitionPropertyUpdateRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type CustomEventsApiRetrieveDefinitionRequest struct {
	ctx context.Context
	ApiService *CustomEventsApiService
	name string
}

func (r CustomEventsApiRetrieveDefinitionRequest) Execute() (*CustomEventDefinition, *http.Response, error) {
	return r.ApiService.RetrieveDefinitionExecute(r)
}

/*
RetrieveDefinition Retrieve an event definition

Retrieves a custom event definition you previously created.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name Name of the custom event.
 @return CustomEventsApiRetrieveDefinitionRequest
*/
func (a *CustomEventsApiService) RetrieveDefinition(ctx context.Context, name string) CustomEventsApiRetrieveDefinitionRequest {
	return CustomEventsApiRetrieveDefinitionRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
	}
}

// Execute executes the request
//  @return CustomEventDefinition
func (a *CustomEventsApiService) RetrieveDefinitionExecute(r CustomEventsApiRetrieveDefinitionRequest) (*CustomEventDefinition, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CustomEventDefinition
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomEventsApiService.RetrieveDefinition")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/event/definitions/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterToString(r.name, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type CustomEventsApiSendEventRequest struct {
	ctx context.Context
	ApiService *CustomEventsApiService
	customEventSendRequest *CustomEventSendRequest
}

func (r CustomEventsApiSendEventRequest) CustomEventSendRequest(customEventSendRequest CustomEventSendRequest) CustomEventsApiSendEventRequest {
	r.customEventSendRequest = &customEventSendRequest
	return r
}

func (r CustomEventsApiSendEventRequest) Execute() (*http.Response, error) {
	return r.ApiService.SendEventExecute(r)
}

/*
SendEvent Send an event

Sends an event.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return CustomEventsApiSendEventRequest
*/
func (a *CustomEventsApiService) SendEvent(ctx context.Context) CustomEventsApiSendEventRequest {
	return CustomEventsApiSendEventRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *CustomEventsApiService) SendEventExecute(r CustomEventsApiSendEventRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomEventsApiService.SendEvent")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/event/events"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.customEventSendRequest == nil {
		return nil, reportError("customEventSendRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.customEventSendRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type CustomEventsApiUpdateDefinitionRequest struct {
	ctx context.Context
	ApiService *CustomEventsApiService
	name string
	customEventDefinitionUpdateRequest *CustomEventDefinitionUpdateRequest
}

func (r CustomEventsApiUpdateDefinitionRequest) CustomEventDefinitionUpdateRequest(customEventDefinitionUpdateRequest CustomEventDefinitionUpdateRequest) CustomEventsApiUpdateDefinitionRequest {
	r.customEventDefinitionUpdateRequest = &customEventDefinitionUpdateRequest
	return r
}

func (r CustomEventsApiUpdateDefinitionRequest) Execute() (*CustomEventDefinition, *http.Response, error) {
	return r.ApiService.UpdateDefinitionExecute(r)
}

/*
UpdateDefinition Update an event definition

Updates an event definition's label and description.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name Name of the custom event.
 @return CustomEventsApiUpdateDefinitionRequest
*/
func (a *CustomEventsApiService) UpdateDefinition(ctx context.Context, name string) CustomEventsApiUpdateDefinitionRequest {
	return CustomEventsApiUpdateDefinitionRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
	}
}

// Execute executes the request
//  @return CustomEventDefinition
func (a *CustomEventsApiService) UpdateDefinitionExecute(r CustomEventsApiUpdateDefinitionRequest) (*CustomEventDefinition, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CustomEventDefinition
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomEventsApiService.UpdateDefinition")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/event/definitions/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterToString(r.name, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.customEventDefinitionUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("customEventDefinitionUpdateRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.customEventDefinitionUpdateRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
