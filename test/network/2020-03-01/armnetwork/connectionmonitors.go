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
	"path"
	"strings"
	"time"
)

// ConnectionMonitorsOperations contains the methods for the ConnectionMonitors group.
type ConnectionMonitorsOperations interface {
	// BeginCreateOrUpdate - Create or update a connection monitor.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, parameters ConnectionMonitor) (*ConnectionMonitorResultPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (ConnectionMonitorResultPoller, error)
	// BeginDelete - Deletes the specified connection monitor.
	BeginDelete(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets a connection monitor by name.
	Get(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string) (*ConnectionMonitorResultResponse, error)
	// List - Lists all connection monitors for the specified Network Watcher.
	List(ctx context.Context, resourceGroupName string, networkWatcherName string) (*ConnectionMonitorListResultResponse, error)
	// BeginQuery - Query a snapshot of the most recent connection states.
	BeginQuery(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string) (*ConnectionMonitorQueryResultPollerResponse, error)
	// ResumeQuery - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeQuery(token string) (ConnectionMonitorQueryResultPoller, error)
	// BeginStart - Starts the specified connection monitor.
	BeginStart(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string) (*HTTPPollerResponse, error)
	// ResumeStart - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeStart(token string) (HTTPPoller, error)
	// BeginStop - Stops the specified connection monitor.
	BeginStop(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string) (*HTTPPollerResponse, error)
	// ResumeStop - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeStop(token string) (HTTPPoller, error)
	// UpdateTags - Update tags of the specified connection monitor.
	UpdateTags(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, parameters TagsObject) (*ConnectionMonitorResultResponse, error)
}

// connectionMonitorsOperations implements the ConnectionMonitorsOperations interface.
type connectionMonitorsOperations struct {
	*Client
	subscriptionID string
}

// CreateOrUpdate - Create or update a connection monitor.
func (client *connectionMonitorsOperations) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, parameters ConnectionMonitor) (*ConnectionMonitorResultPollerResponse, error) {
	req, err := client.createOrUpdateCreateRequest(resourceGroupName, networkWatcherName, connectionMonitorName, parameters)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.createOrUpdateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := armcore.NewPoller("connectionMonitorsOperations.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &connectionMonitorResultPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ConnectionMonitorResultResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *connectionMonitorsOperations) ResumeCreateOrUpdate(token string) (ConnectionMonitorResultPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("connectionMonitorsOperations.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &connectionMonitorResultPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *connectionMonitorsOperations) createOrUpdateCreateRequest(resourceGroupName string, networkWatcherName string, connectionMonitorName string, parameters ConnectionMonitor) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/connectionMonitors/{connectionMonitorName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	urlPath = strings.ReplaceAll(urlPath, "{connectionMonitorName}", url.PathEscape(connectionMonitorName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *connectionMonitorsOperations) createOrUpdateHandleResponse(resp *azcore.Response) (*ConnectionMonitorResultPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated, http.StatusNoContent) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return &ConnectionMonitorResultPollerResponse{RawResponse: resp.Response}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *connectionMonitorsOperations) createOrUpdateHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Delete - Deletes the specified connection monitor.
func (client *connectionMonitorsOperations) BeginDelete(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string) (*HTTPPollerResponse, error) {
	req, err := client.deleteCreateRequest(resourceGroupName, networkWatcherName, connectionMonitorName)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.deleteHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := armcore.NewPoller("connectionMonitorsOperations.Delete", "location", resp, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *connectionMonitorsOperations) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("connectionMonitorsOperations.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *connectionMonitorsOperations) deleteCreateRequest(resourceGroupName string, networkWatcherName string, connectionMonitorName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/connectionMonitors/{connectionMonitorName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	urlPath = strings.ReplaceAll(urlPath, "{connectionMonitorName}", url.PathEscape(connectionMonitorName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodDelete, *u)
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *connectionMonitorsOperations) deleteHandleResponse(resp *azcore.Response) (*HTTPPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return &HTTPPollerResponse{RawResponse: resp.Response}, nil
}

// deleteHandleError handles the Delete error response.
func (client *connectionMonitorsOperations) deleteHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Get - Gets a connection monitor by name.
func (client *connectionMonitorsOperations) Get(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string) (*ConnectionMonitorResultResponse, error) {
	req, err := client.getCreateRequest(resourceGroupName, networkWatcherName, connectionMonitorName)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getCreateRequest creates the Get request.
func (client *connectionMonitorsOperations) getCreateRequest(resourceGroupName string, networkWatcherName string, connectionMonitorName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/connectionMonitors/{connectionMonitorName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	urlPath = strings.ReplaceAll(urlPath, "{connectionMonitorName}", url.PathEscape(connectionMonitorName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *connectionMonitorsOperations) getHandleResponse(resp *azcore.Response) (*ConnectionMonitorResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getHandleError(resp)
	}
	result := ConnectionMonitorResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ConnectionMonitorResult)
}

// getHandleError handles the Get error response.
func (client *connectionMonitorsOperations) getHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// List - Lists all connection monitors for the specified Network Watcher.
func (client *connectionMonitorsOperations) List(ctx context.Context, resourceGroupName string, networkWatcherName string) (*ConnectionMonitorListResultResponse, error) {
	req, err := client.listCreateRequest(resourceGroupName, networkWatcherName)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.listHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// listCreateRequest creates the List request.
func (client *connectionMonitorsOperations) listCreateRequest(resourceGroupName string, networkWatcherName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/connectionMonitors"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// listHandleResponse handles the List response.
func (client *connectionMonitorsOperations) listHandleResponse(resp *azcore.Response) (*ConnectionMonitorListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listHandleError(resp)
	}
	result := ConnectionMonitorListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ConnectionMonitorListResult)
}

// listHandleError handles the List error response.
func (client *connectionMonitorsOperations) listHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Query - Query a snapshot of the most recent connection states.
func (client *connectionMonitorsOperations) BeginQuery(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string) (*ConnectionMonitorQueryResultPollerResponse, error) {
	req, err := client.queryCreateRequest(resourceGroupName, networkWatcherName, connectionMonitorName)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.queryHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := armcore.NewPoller("connectionMonitorsOperations.Query", "location", resp, client.queryHandleError)
	if err != nil {
		return nil, err
	}
	poller := &connectionMonitorQueryResultPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ConnectionMonitorQueryResultResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *connectionMonitorsOperations) ResumeQuery(token string) (ConnectionMonitorQueryResultPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("connectionMonitorsOperations.Query", token, client.queryHandleError)
	if err != nil {
		return nil, err
	}
	return &connectionMonitorQueryResultPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// queryCreateRequest creates the Query request.
func (client *connectionMonitorsOperations) queryCreateRequest(resourceGroupName string, networkWatcherName string, connectionMonitorName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/connectionMonitors/{connectionMonitorName}/query"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	urlPath = strings.ReplaceAll(urlPath, "{connectionMonitorName}", url.PathEscape(connectionMonitorName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPost, *u)
	return req, nil
}

// queryHandleResponse handles the Query response.
func (client *connectionMonitorsOperations) queryHandleResponse(resp *azcore.Response) (*ConnectionMonitorQueryResultPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.queryHandleError(resp)
	}
	return &ConnectionMonitorQueryResultPollerResponse{RawResponse: resp.Response}, nil
}

// queryHandleError handles the Query error response.
func (client *connectionMonitorsOperations) queryHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Start - Starts the specified connection monitor.
func (client *connectionMonitorsOperations) BeginStart(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string) (*HTTPPollerResponse, error) {
	req, err := client.startCreateRequest(resourceGroupName, networkWatcherName, connectionMonitorName)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.startHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := armcore.NewPoller("connectionMonitorsOperations.Start", "location", resp, client.startHandleError)
	if err != nil {
		return nil, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *connectionMonitorsOperations) ResumeStart(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("connectionMonitorsOperations.Start", token, client.startHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// startCreateRequest creates the Start request.
func (client *connectionMonitorsOperations) startCreateRequest(resourceGroupName string, networkWatcherName string, connectionMonitorName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/connectionMonitors/{connectionMonitorName}/start"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	urlPath = strings.ReplaceAll(urlPath, "{connectionMonitorName}", url.PathEscape(connectionMonitorName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPost, *u)
	return req, nil
}

// startHandleResponse handles the Start response.
func (client *connectionMonitorsOperations) startHandleResponse(resp *azcore.Response) (*HTTPPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.startHandleError(resp)
	}
	return &HTTPPollerResponse{RawResponse: resp.Response}, nil
}

// startHandleError handles the Start error response.
func (client *connectionMonitorsOperations) startHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Stop - Stops the specified connection monitor.
func (client *connectionMonitorsOperations) BeginStop(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string) (*HTTPPollerResponse, error) {
	req, err := client.stopCreateRequest(resourceGroupName, networkWatcherName, connectionMonitorName)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.stopHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := armcore.NewPoller("connectionMonitorsOperations.Stop", "location", resp, client.stopHandleError)
	if err != nil {
		return nil, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *connectionMonitorsOperations) ResumeStop(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("connectionMonitorsOperations.Stop", token, client.stopHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// stopCreateRequest creates the Stop request.
func (client *connectionMonitorsOperations) stopCreateRequest(resourceGroupName string, networkWatcherName string, connectionMonitorName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/connectionMonitors/{connectionMonitorName}/stop"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	urlPath = strings.ReplaceAll(urlPath, "{connectionMonitorName}", url.PathEscape(connectionMonitorName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPost, *u)
	return req, nil
}

// stopHandleResponse handles the Stop response.
func (client *connectionMonitorsOperations) stopHandleResponse(resp *azcore.Response) (*HTTPPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.stopHandleError(resp)
	}
	return &HTTPPollerResponse{RawResponse: resp.Response}, nil
}

// stopHandleError handles the Stop error response.
func (client *connectionMonitorsOperations) stopHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// UpdateTags - Update tags of the specified connection monitor.
func (client *connectionMonitorsOperations) UpdateTags(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, parameters TagsObject) (*ConnectionMonitorResultResponse, error) {
	req, err := client.updateTagsCreateRequest(resourceGroupName, networkWatcherName, connectionMonitorName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.updateTagsHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *connectionMonitorsOperations) updateTagsCreateRequest(resourceGroupName string, networkWatcherName string, connectionMonitorName string, parameters TagsObject) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/connectionMonitors/{connectionMonitorName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	urlPath = strings.ReplaceAll(urlPath, "{connectionMonitorName}", url.PathEscape(connectionMonitorName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPatch, *u)
	return req, req.MarshalAsJSON(parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *connectionMonitorsOperations) updateTagsHandleResponse(resp *azcore.Response) (*ConnectionMonitorResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.updateTagsHandleError(resp)
	}
	result := ConnectionMonitorResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ConnectionMonitorResult)
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *connectionMonitorsOperations) updateTagsHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
