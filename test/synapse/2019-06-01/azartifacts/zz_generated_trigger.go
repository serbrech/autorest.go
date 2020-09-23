// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

type triggerClient struct {
	*client
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *triggerClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// CreateOrUpdateTrigger - Creates or updates a trigger.
func (client *triggerClient) CreateOrUpdateTrigger(ctx context.Context, triggerName string, trigger TriggerResource, triggerCreateOrUpdateTriggerOptions *TriggerCreateOrUpdateTriggerOptions) (*azcore.Response, error) {
	req, err := client.CreateOrUpdateTriggerCreateRequest(ctx, triggerName, trigger, triggerCreateOrUpdateTriggerOptions)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.CreateOrUpdateTriggerHandleError(resp)
	}
	return resp, nil
}

// CreateOrUpdateTriggerCreateRequest creates the CreateOrUpdateTrigger request.
func (client *triggerClient) CreateOrUpdateTriggerCreateRequest(ctx context.Context, triggerName string, trigger TriggerResource, triggerCreateOrUpdateTriggerOptions *TriggerCreateOrUpdateTriggerOptions) (*azcore.Request, error) {
	urlPath := "/triggers/{triggerName}"
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	if triggerCreateOrUpdateTriggerOptions != nil && triggerCreateOrUpdateTriggerOptions.IfMatch != nil {
		req.Header.Set("If-Match", *triggerCreateOrUpdateTriggerOptions.IfMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(trigger)
}

// CreateOrUpdateTriggerHandleResponse handles the CreateOrUpdateTrigger response.
func (client *triggerClient) CreateOrUpdateTriggerHandleResponse(resp *azcore.Response) (*TriggerResourceResponse, error) {
	result := TriggerResourceResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.TriggerResource)
}

// CreateOrUpdateTriggerHandleError handles the CreateOrUpdateTrigger error response.
func (client *triggerClient) CreateOrUpdateTriggerHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// DeleteTrigger - Deletes a trigger.
func (client *triggerClient) DeleteTrigger(ctx context.Context, triggerName string) (*azcore.Response, error) {
	req, err := client.DeleteTriggerCreateRequest(ctx, triggerName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.DeleteTriggerHandleError(resp)
	}
	return resp, nil
}

// DeleteTriggerCreateRequest creates the DeleteTrigger request.
func (client *triggerClient) DeleteTriggerCreateRequest(ctx context.Context, triggerName string) (*azcore.Request, error) {
	urlPath := "/triggers/{triggerName}"
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// DeleteTriggerHandleError handles the DeleteTrigger error response.
func (client *triggerClient) DeleteTriggerHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetEventSubscriptionStatus - Get a trigger's event subscription status.
func (client *triggerClient) GetEventSubscriptionStatus(ctx context.Context, triggerName string) (*TriggerSubscriptionOperationStatusResponse, error) {
	req, err := client.GetEventSubscriptionStatusCreateRequest(ctx, triggerName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetEventSubscriptionStatusHandleError(resp)
	}
	result, err := client.GetEventSubscriptionStatusHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetEventSubscriptionStatusCreateRequest creates the GetEventSubscriptionStatus request.
func (client *triggerClient) GetEventSubscriptionStatusCreateRequest(ctx context.Context, triggerName string) (*azcore.Request, error) {
	urlPath := "/triggers/{triggerName}/getEventSubscriptionStatus"
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetEventSubscriptionStatusHandleResponse handles the GetEventSubscriptionStatus response.
func (client *triggerClient) GetEventSubscriptionStatusHandleResponse(resp *azcore.Response) (*TriggerSubscriptionOperationStatusResponse, error) {
	result := TriggerSubscriptionOperationStatusResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.TriggerSubscriptionOperationStatus)
}

// GetEventSubscriptionStatusHandleError handles the GetEventSubscriptionStatus error response.
func (client *triggerClient) GetEventSubscriptionStatusHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetTrigger - Gets a trigger.
func (client *triggerClient) GetTrigger(ctx context.Context, triggerName string, triggerGetTriggerOptions *TriggerGetTriggerOptions) (*TriggerResourceResponse, error) {
	req, err := client.GetTriggerCreateRequest(ctx, triggerName, triggerGetTriggerOptions)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNotModified) {
		return nil, client.GetTriggerHandleError(resp)
	}
	result, err := client.GetTriggerHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetTriggerCreateRequest creates the GetTrigger request.
func (client *triggerClient) GetTriggerCreateRequest(ctx context.Context, triggerName string, triggerGetTriggerOptions *TriggerGetTriggerOptions) (*azcore.Request, error) {
	urlPath := "/triggers/{triggerName}"
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	if triggerGetTriggerOptions != nil && triggerGetTriggerOptions.IfNoneMatch != nil {
		req.Header.Set("If-None-Match", *triggerGetTriggerOptions.IfNoneMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetTriggerHandleResponse handles the GetTrigger response.
func (client *triggerClient) GetTriggerHandleResponse(resp *azcore.Response) (*TriggerResourceResponse, error) {
	result := TriggerResourceResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.TriggerResource)
}

// GetTriggerHandleError handles the GetTrigger error response.
func (client *triggerClient) GetTriggerHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetTriggersByWorkspace - Lists triggers.
func (client *triggerClient) GetTriggersByWorkspace() TriggerListResponsePager {
	return &triggerListResponsePager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.GetTriggersByWorkspaceCreateRequest(ctx)
		},
		responder: client.GetTriggersByWorkspaceHandleResponse,
		errorer:   client.GetTriggersByWorkspaceHandleError,
		advancer: func(ctx context.Context, resp *TriggerListResponseResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.TriggerListResponse.NextLink)
		},
	}
}

// GetTriggersByWorkspaceCreateRequest creates the GetTriggersByWorkspace request.
func (client *triggerClient) GetTriggersByWorkspaceCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/triggers"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetTriggersByWorkspaceHandleResponse handles the GetTriggersByWorkspace response.
func (client *triggerClient) GetTriggersByWorkspaceHandleResponse(resp *azcore.Response) (*TriggerListResponseResponse, error) {
	result := TriggerListResponseResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.TriggerListResponse)
}

// GetTriggersByWorkspaceHandleError handles the GetTriggersByWorkspace error response.
func (client *triggerClient) GetTriggersByWorkspaceHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// StartTrigger - Starts a trigger.
func (client *triggerClient) StartTrigger(ctx context.Context, triggerName string) (*azcore.Response, error) {
	req, err := client.StartTriggerCreateRequest(ctx, triggerName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.StartTriggerHandleError(resp)
	}
	return resp, nil
}

// StartTriggerCreateRequest creates the StartTrigger request.
func (client *triggerClient) StartTriggerCreateRequest(ctx context.Context, triggerName string) (*azcore.Request, error) {
	urlPath := "/triggers/{triggerName}/start"
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// StartTriggerHandleError handles the StartTrigger error response.
func (client *triggerClient) StartTriggerHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// StopTrigger - Stops a trigger.
func (client *triggerClient) StopTrigger(ctx context.Context, triggerName string) (*azcore.Response, error) {
	req, err := client.StopTriggerCreateRequest(ctx, triggerName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.StopTriggerHandleError(resp)
	}
	return resp, nil
}

// StopTriggerCreateRequest creates the StopTrigger request.
func (client *triggerClient) StopTriggerCreateRequest(ctx context.Context, triggerName string) (*azcore.Request, error) {
	urlPath := "/triggers/{triggerName}/stop"
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// StopTriggerHandleError handles the StopTrigger error response.
func (client *triggerClient) StopTriggerHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// SubscribeTriggerToEvents - Subscribe event trigger to events.
func (client *triggerClient) SubscribeTriggerToEvents(ctx context.Context, triggerName string) (*azcore.Response, error) {
	req, err := client.SubscribeTriggerToEventsCreateRequest(ctx, triggerName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.SubscribeTriggerToEventsHandleError(resp)
	}
	return resp, nil
}

// SubscribeTriggerToEventsCreateRequest creates the SubscribeTriggerToEvents request.
func (client *triggerClient) SubscribeTriggerToEventsCreateRequest(ctx context.Context, triggerName string) (*azcore.Request, error) {
	urlPath := "/triggers/{triggerName}/subscribeToEvents"
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// SubscribeTriggerToEventsHandleResponse handles the SubscribeTriggerToEvents response.
func (client *triggerClient) SubscribeTriggerToEventsHandleResponse(resp *azcore.Response) (*TriggerSubscriptionOperationStatusResponse, error) {
	result := TriggerSubscriptionOperationStatusResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.TriggerSubscriptionOperationStatus)
}

// SubscribeTriggerToEventsHandleError handles the SubscribeTriggerToEvents error response.
func (client *triggerClient) SubscribeTriggerToEventsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// UnsubscribeTriggerFromEvents - Unsubscribe event trigger from events.
func (client *triggerClient) UnsubscribeTriggerFromEvents(ctx context.Context, triggerName string) (*azcore.Response, error) {
	req, err := client.UnsubscribeTriggerFromEventsCreateRequest(ctx, triggerName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.UnsubscribeTriggerFromEventsHandleError(resp)
	}
	return resp, nil
}

// UnsubscribeTriggerFromEventsCreateRequest creates the UnsubscribeTriggerFromEvents request.
func (client *triggerClient) UnsubscribeTriggerFromEventsCreateRequest(ctx context.Context, triggerName string) (*azcore.Request, error) {
	urlPath := "/triggers/{triggerName}/unsubscribeFromEvents"
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// UnsubscribeTriggerFromEventsHandleResponse handles the UnsubscribeTriggerFromEvents response.
func (client *triggerClient) UnsubscribeTriggerFromEventsHandleResponse(resp *azcore.Response) (*TriggerSubscriptionOperationStatusResponse, error) {
	result := TriggerSubscriptionOperationStatusResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.TriggerSubscriptionOperationStatus)
}

// UnsubscribeTriggerFromEventsHandleError handles the UnsubscribeTriggerFromEvents error response.
func (client *triggerClient) UnsubscribeTriggerFromEventsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
