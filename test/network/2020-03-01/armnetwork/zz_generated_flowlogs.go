// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// FlowLogsOperations contains the methods for the FlowLogs group.
type FlowLogsOperations interface {
	// BeginCreateOrUpdate - Create or update a flow log for the specified network security group.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, networkWatcherName string, flowLogName string, parameters FlowLog, options *FlowLogsCreateOrUpdateOptions) (*FlowLogPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (FlowLogPoller, error)
	// BeginDelete - Deletes the specified flow log resource.
	BeginDelete(ctx context.Context, resourceGroupName string, networkWatcherName string, flowLogName string, options *FlowLogsDeleteOptions) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets a flow log resource by name.
	Get(ctx context.Context, resourceGroupName string, networkWatcherName string, flowLogName string, options *FlowLogsGetOptions) (*FlowLogResponse, error)
	// List - Lists all flow log resources for the specified Network Watcher.
	List(resourceGroupName string, networkWatcherName string, options *FlowLogsListOptions) FlowLogListResultPager
}

// FlowLogsClient implements the FlowLogsOperations interface.
// Don't use this type directly, use NewFlowLogsClient() instead.
type FlowLogsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewFlowLogsClient creates a new instance of FlowLogsClient with the specified values.
func NewFlowLogsClient(con *armcore.Connection, subscriptionID string) FlowLogsOperations {
	return &FlowLogsClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client *FlowLogsClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

func (client *FlowLogsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, networkWatcherName string, flowLogName string, parameters FlowLog, options *FlowLogsCreateOrUpdateOptions) (*FlowLogPollerResponse, error) {
	resp, err := client.CreateOrUpdate(ctx, resourceGroupName, networkWatcherName, flowLogName, parameters, options)
	if err != nil {
		return nil, err
	}
	result := &FlowLogPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("FlowLogsClient.CreateOrUpdate", "azure-async-operation", resp, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &flowLogPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*FlowLogResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *FlowLogsClient) ResumeCreateOrUpdate(token string) (FlowLogPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("FlowLogsClient.CreateOrUpdate", token, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &flowLogPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Create or update a flow log for the specified network security group.
func (client *FlowLogsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, networkWatcherName string, flowLogName string, parameters FlowLog, options *FlowLogsCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.CreateOrUpdateCreateRequest(ctx, resourceGroupName, networkWatcherName, flowLogName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.CreateOrUpdateHandleError(resp)
	}
	return resp, nil
}

// CreateOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *FlowLogsClient) CreateOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, flowLogName string, parameters FlowLog, options *FlowLogsCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/flowLogs/{flowLogName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	urlPath = strings.ReplaceAll(urlPath, "{flowLogName}", url.PathEscape(flowLogName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// CreateOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *FlowLogsClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*FlowLogResponse, error) {
	result := FlowLogResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.FlowLog)
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *FlowLogsClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *FlowLogsClient) BeginDelete(ctx context.Context, resourceGroupName string, networkWatcherName string, flowLogName string, options *FlowLogsDeleteOptions) (*HTTPPollerResponse, error) {
	resp, err := client.Delete(ctx, resourceGroupName, networkWatcherName, flowLogName, options)
	if err != nil {
		return nil, err
	}
	result := &HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("FlowLogsClient.Delete", "location", resp, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *FlowLogsClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("FlowLogsClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified flow log resource.
func (client *FlowLogsClient) Delete(ctx context.Context, resourceGroupName string, networkWatcherName string, flowLogName string, options *FlowLogsDeleteOptions) (*azcore.Response, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, networkWatcherName, flowLogName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusAccepted, http.StatusNoContent) {
		return nil, client.DeleteHandleError(resp)
	}
	return resp, nil
}

// DeleteCreateRequest creates the Delete request.
func (client *FlowLogsClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, flowLogName string, options *FlowLogsDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/flowLogs/{flowLogName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	urlPath = strings.ReplaceAll(urlPath, "{flowLogName}", url.PathEscape(flowLogName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// DeleteHandleError handles the Delete error response.
func (client *FlowLogsClient) DeleteHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets a flow log resource by name.
func (client *FlowLogsClient) Get(ctx context.Context, resourceGroupName string, networkWatcherName string, flowLogName string, options *FlowLogsGetOptions) (*FlowLogResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, networkWatcherName, flowLogName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetHandleError(resp)
	}
	result, err := client.GetHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCreateRequest creates the Get request.
func (client *FlowLogsClient) GetCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, flowLogName string, options *FlowLogsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/flowLogs/{flowLogName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	urlPath = strings.ReplaceAll(urlPath, "{flowLogName}", url.PathEscape(flowLogName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetHandleResponse handles the Get response.
func (client *FlowLogsClient) GetHandleResponse(resp *azcore.Response) (*FlowLogResponse, error) {
	result := FlowLogResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.FlowLog)
}

// GetHandleError handles the Get error response.
func (client *FlowLogsClient) GetHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Lists all flow log resources for the specified Network Watcher.
func (client *FlowLogsClient) List(resourceGroupName string, networkWatcherName string, options *FlowLogsListOptions) FlowLogListResultPager {
	return &flowLogListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, resourceGroupName, networkWatcherName, options)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *FlowLogListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.FlowLogListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListCreateRequest creates the List request.
func (client *FlowLogsClient) ListCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, options *FlowLogsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/flowLogs"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListHandleResponse handles the List response.
func (client *FlowLogsClient) ListHandleResponse(resp *azcore.Response) (*FlowLogListResultResponse, error) {
	result := FlowLogListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.FlowLogListResult)
}

// ListHandleError handles the List error response.
func (client *FlowLogsClient) ListHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
