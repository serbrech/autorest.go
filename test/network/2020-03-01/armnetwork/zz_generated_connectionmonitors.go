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

// ConnectionMonitorsOperations contains the methods for the ConnectionMonitors group.
type ConnectionMonitorsOperations interface {
	// BeginCreateOrUpdate - Create or update a connection monitor.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, parameters ConnectionMonitor, options *ConnectionMonitorsCreateOrUpdateOptions) (*ConnectionMonitorResultPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (ConnectionMonitorResultPoller, error)
	// BeginDelete - Deletes the specified connection monitor.
	BeginDelete(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *ConnectionMonitorsDeleteOptions) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets a connection monitor by name.
	Get(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *ConnectionMonitorsGetOptions) (*ConnectionMonitorResultResponse, error)
	// List - Lists all connection monitors for the specified Network Watcher.
	List(ctx context.Context, resourceGroupName string, networkWatcherName string, options *ConnectionMonitorsListOptions) (*ConnectionMonitorListResultResponse, error)
	// BeginQuery - Query a snapshot of the most recent connection states.
	BeginQuery(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *ConnectionMonitorsQueryOptions) (*ConnectionMonitorQueryResultPollerResponse, error)
	// ResumeQuery - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeQuery(token string) (ConnectionMonitorQueryResultPoller, error)
	// BeginStart - Starts the specified connection monitor.
	BeginStart(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *ConnectionMonitorsStartOptions) (*HTTPPollerResponse, error)
	// ResumeStart - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeStart(token string) (HTTPPoller, error)
	// BeginStop - Stops the specified connection monitor.
	BeginStop(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *ConnectionMonitorsStopOptions) (*HTTPPollerResponse, error)
	// ResumeStop - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeStop(token string) (HTTPPoller, error)
	// UpdateTags - Update tags of the specified connection monitor.
	UpdateTags(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, parameters TagsObject, options *ConnectionMonitorsUpdateTagsOptions) (*ConnectionMonitorResultResponse, error)
}

// ConnectionMonitorsClient implements the ConnectionMonitorsOperations interface.
// Don't use this type directly, use NewConnectionMonitorsClient() instead.
type ConnectionMonitorsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewConnectionMonitorsClient creates a new instance of ConnectionMonitorsClient with the specified values.
func NewConnectionMonitorsClient(con *armcore.Connection, subscriptionID string) ConnectionMonitorsOperations {
	return &ConnectionMonitorsClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client *ConnectionMonitorsClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

func (client *ConnectionMonitorsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, parameters ConnectionMonitor, options *ConnectionMonitorsCreateOrUpdateOptions) (*ConnectionMonitorResultPollerResponse, error) {
	resp, err := client.CreateOrUpdate(ctx, resourceGroupName, networkWatcherName, connectionMonitorName, parameters, options)
	if err != nil {
		return nil, err
	}
	result := &ConnectionMonitorResultPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ConnectionMonitorsClient.CreateOrUpdate", "azure-async-operation", resp, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &connectionMonitorResultPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ConnectionMonitorResultResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *ConnectionMonitorsClient) ResumeCreateOrUpdate(token string) (ConnectionMonitorResultPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ConnectionMonitorsClient.CreateOrUpdate", token, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &connectionMonitorResultPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Create or update a connection monitor.
func (client *ConnectionMonitorsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, parameters ConnectionMonitor, options *ConnectionMonitorsCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.CreateOrUpdateCreateRequest(ctx, resourceGroupName, networkWatcherName, connectionMonitorName, parameters, options)
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
func (client *ConnectionMonitorsClient) CreateOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, parameters ConnectionMonitor, options *ConnectionMonitorsCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/connectionMonitors/{connectionMonitorName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	urlPath = strings.ReplaceAll(urlPath, "{connectionMonitorName}", url.PathEscape(connectionMonitorName))
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
func (client *ConnectionMonitorsClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*ConnectionMonitorResultResponse, error) {
	result := ConnectionMonitorResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ConnectionMonitorResult)
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ConnectionMonitorsClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *ConnectionMonitorsClient) BeginDelete(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *ConnectionMonitorsDeleteOptions) (*HTTPPollerResponse, error) {
	resp, err := client.Delete(ctx, resourceGroupName, networkWatcherName, connectionMonitorName, options)
	if err != nil {
		return nil, err
	}
	result := &HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ConnectionMonitorsClient.Delete", "location", resp, client.DeleteHandleError)
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

func (client *ConnectionMonitorsClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ConnectionMonitorsClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified connection monitor.
func (client *ConnectionMonitorsClient) Delete(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *ConnectionMonitorsDeleteOptions) (*azcore.Response, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, networkWatcherName, connectionMonitorName, options)
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
func (client *ConnectionMonitorsClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *ConnectionMonitorsDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/connectionMonitors/{connectionMonitorName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	urlPath = strings.ReplaceAll(urlPath, "{connectionMonitorName}", url.PathEscape(connectionMonitorName))
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
func (client *ConnectionMonitorsClient) DeleteHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets a connection monitor by name.
func (client *ConnectionMonitorsClient) Get(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *ConnectionMonitorsGetOptions) (*ConnectionMonitorResultResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, networkWatcherName, connectionMonitorName, options)
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
func (client *ConnectionMonitorsClient) GetCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *ConnectionMonitorsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/connectionMonitors/{connectionMonitorName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	urlPath = strings.ReplaceAll(urlPath, "{connectionMonitorName}", url.PathEscape(connectionMonitorName))
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
func (client *ConnectionMonitorsClient) GetHandleResponse(resp *azcore.Response) (*ConnectionMonitorResultResponse, error) {
	result := ConnectionMonitorResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ConnectionMonitorResult)
}

// GetHandleError handles the Get error response.
func (client *ConnectionMonitorsClient) GetHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Lists all connection monitors for the specified Network Watcher.
func (client *ConnectionMonitorsClient) List(ctx context.Context, resourceGroupName string, networkWatcherName string, options *ConnectionMonitorsListOptions) (*ConnectionMonitorListResultResponse, error) {
	req, err := client.ListCreateRequest(ctx, resourceGroupName, networkWatcherName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.ListHandleError(resp)
	}
	result, err := client.ListHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListCreateRequest creates the List request.
func (client *ConnectionMonitorsClient) ListCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, options *ConnectionMonitorsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/connectionMonitors"
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
func (client *ConnectionMonitorsClient) ListHandleResponse(resp *azcore.Response) (*ConnectionMonitorListResultResponse, error) {
	result := ConnectionMonitorListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ConnectionMonitorListResult)
}

// ListHandleError handles the List error response.
func (client *ConnectionMonitorsClient) ListHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *ConnectionMonitorsClient) BeginQuery(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *ConnectionMonitorsQueryOptions) (*ConnectionMonitorQueryResultPollerResponse, error) {
	resp, err := client.Query(ctx, resourceGroupName, networkWatcherName, connectionMonitorName, options)
	if err != nil {
		return nil, err
	}
	result := &ConnectionMonitorQueryResultPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ConnectionMonitorsClient.Query", "location", resp, client.QueryHandleError)
	if err != nil {
		return nil, err
	}
	poller := &connectionMonitorQueryResultPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ConnectionMonitorQueryResultResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *ConnectionMonitorsClient) ResumeQuery(token string) (ConnectionMonitorQueryResultPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ConnectionMonitorsClient.Query", token, client.QueryHandleError)
	if err != nil {
		return nil, err
	}
	return &connectionMonitorQueryResultPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Query - Query a snapshot of the most recent connection states.
func (client *ConnectionMonitorsClient) Query(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *ConnectionMonitorsQueryOptions) (*azcore.Response, error) {
	req, err := client.QueryCreateRequest(ctx, resourceGroupName, networkWatcherName, connectionMonitorName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.QueryHandleError(resp)
	}
	return resp, nil
}

// QueryCreateRequest creates the Query request.
func (client *ConnectionMonitorsClient) QueryCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *ConnectionMonitorsQueryOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/connectionMonitors/{connectionMonitorName}/query"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	urlPath = strings.ReplaceAll(urlPath, "{connectionMonitorName}", url.PathEscape(connectionMonitorName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// QueryHandleResponse handles the Query response.
func (client *ConnectionMonitorsClient) QueryHandleResponse(resp *azcore.Response) (*ConnectionMonitorQueryResultResponse, error) {
	result := ConnectionMonitorQueryResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ConnectionMonitorQueryResult)
}

// QueryHandleError handles the Query error response.
func (client *ConnectionMonitorsClient) QueryHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *ConnectionMonitorsClient) BeginStart(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *ConnectionMonitorsStartOptions) (*HTTPPollerResponse, error) {
	resp, err := client.Start(ctx, resourceGroupName, networkWatcherName, connectionMonitorName, options)
	if err != nil {
		return nil, err
	}
	result := &HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ConnectionMonitorsClient.Start", "location", resp, client.StartHandleError)
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

func (client *ConnectionMonitorsClient) ResumeStart(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ConnectionMonitorsClient.Start", token, client.StartHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Start - Starts the specified connection monitor.
func (client *ConnectionMonitorsClient) Start(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *ConnectionMonitorsStartOptions) (*azcore.Response, error) {
	req, err := client.StartCreateRequest(ctx, resourceGroupName, networkWatcherName, connectionMonitorName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.StartHandleError(resp)
	}
	return resp, nil
}

// StartCreateRequest creates the Start request.
func (client *ConnectionMonitorsClient) StartCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *ConnectionMonitorsStartOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/connectionMonitors/{connectionMonitorName}/start"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	urlPath = strings.ReplaceAll(urlPath, "{connectionMonitorName}", url.PathEscape(connectionMonitorName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// StartHandleError handles the Start error response.
func (client *ConnectionMonitorsClient) StartHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *ConnectionMonitorsClient) BeginStop(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *ConnectionMonitorsStopOptions) (*HTTPPollerResponse, error) {
	resp, err := client.Stop(ctx, resourceGroupName, networkWatcherName, connectionMonitorName, options)
	if err != nil {
		return nil, err
	}
	result := &HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ConnectionMonitorsClient.Stop", "location", resp, client.StopHandleError)
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

func (client *ConnectionMonitorsClient) ResumeStop(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ConnectionMonitorsClient.Stop", token, client.StopHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Stop - Stops the specified connection monitor.
func (client *ConnectionMonitorsClient) Stop(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *ConnectionMonitorsStopOptions) (*azcore.Response, error) {
	req, err := client.StopCreateRequest(ctx, resourceGroupName, networkWatcherName, connectionMonitorName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.StopHandleError(resp)
	}
	return resp, nil
}

// StopCreateRequest creates the Stop request.
func (client *ConnectionMonitorsClient) StopCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *ConnectionMonitorsStopOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/connectionMonitors/{connectionMonitorName}/stop"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	urlPath = strings.ReplaceAll(urlPath, "{connectionMonitorName}", url.PathEscape(connectionMonitorName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// StopHandleError handles the Stop error response.
func (client *ConnectionMonitorsClient) StopHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// UpdateTags - Update tags of the specified connection monitor.
func (client *ConnectionMonitorsClient) UpdateTags(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, parameters TagsObject, options *ConnectionMonitorsUpdateTagsOptions) (*ConnectionMonitorResultResponse, error) {
	req, err := client.UpdateTagsCreateRequest(ctx, resourceGroupName, networkWatcherName, connectionMonitorName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.UpdateTagsHandleError(resp)
	}
	result, err := client.UpdateTagsHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateTagsCreateRequest creates the UpdateTags request.
func (client *ConnectionMonitorsClient) UpdateTagsCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, parameters TagsObject, options *ConnectionMonitorsUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/connectionMonitors/{connectionMonitorName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	urlPath = strings.ReplaceAll(urlPath, "{connectionMonitorName}", url.PathEscape(connectionMonitorName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// UpdateTagsHandleResponse handles the UpdateTags response.
func (client *ConnectionMonitorsClient) UpdateTagsHandleResponse(resp *azcore.Response) (*ConnectionMonitorResultResponse, error) {
	result := ConnectionMonitorResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ConnectionMonitorResult)
}

// UpdateTagsHandleError handles the UpdateTags error response.
func (client *ConnectionMonitorsClient) UpdateTagsHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
