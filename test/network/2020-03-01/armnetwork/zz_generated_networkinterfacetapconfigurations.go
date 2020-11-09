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

// NetworkInterfaceTapConfigurationsOperations contains the methods for the NetworkInterfaceTapConfigurations group.
type NetworkInterfaceTapConfigurationsOperations interface {
	// BeginCreateOrUpdate - Creates or updates a Tap configuration in the specified NetworkInterface.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, networkInterfaceName string, tapConfigurationName string, tapConfigurationParameters NetworkInterfaceTapConfiguration, options *NetworkInterfaceTapConfigurationsCreateOrUpdateOptions) (*NetworkInterfaceTapConfigurationPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (NetworkInterfaceTapConfigurationPoller, error)
	// BeginDelete - Deletes the specified tap configuration from the NetworkInterface.
	BeginDelete(ctx context.Context, resourceGroupName string, networkInterfaceName string, tapConfigurationName string, options *NetworkInterfaceTapConfigurationsDeleteOptions) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Get the specified tap configuration on a network interface.
	Get(ctx context.Context, resourceGroupName string, networkInterfaceName string, tapConfigurationName string, options *NetworkInterfaceTapConfigurationsGetOptions) (*NetworkInterfaceTapConfigurationResponse, error)
	// List - Get all Tap configurations in a network interface.
	List(resourceGroupName string, networkInterfaceName string, options *NetworkInterfaceTapConfigurationsListOptions) NetworkInterfaceTapConfigurationListResultPager
}

// NetworkInterfaceTapConfigurationsClient implements the NetworkInterfaceTapConfigurationsOperations interface.
// Don't use this type directly, use NewNetworkInterfaceTapConfigurationsClient() instead.
type NetworkInterfaceTapConfigurationsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewNetworkInterfaceTapConfigurationsClient creates a new instance of NetworkInterfaceTapConfigurationsClient with the specified values.
func NewNetworkInterfaceTapConfigurationsClient(con *armcore.Connection, subscriptionID string) NetworkInterfaceTapConfigurationsOperations {
	return &NetworkInterfaceTapConfigurationsClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client *NetworkInterfaceTapConfigurationsClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

func (client *NetworkInterfaceTapConfigurationsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, networkInterfaceName string, tapConfigurationName string, tapConfigurationParameters NetworkInterfaceTapConfiguration, options *NetworkInterfaceTapConfigurationsCreateOrUpdateOptions) (*NetworkInterfaceTapConfigurationPollerResponse, error) {
	resp, err := client.CreateOrUpdate(ctx, resourceGroupName, networkInterfaceName, tapConfigurationName, tapConfigurationParameters, options)
	if err != nil {
		return nil, err
	}
	result := &NetworkInterfaceTapConfigurationPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("NetworkInterfaceTapConfigurationsClient.CreateOrUpdate", "azure-async-operation", resp, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &networkInterfaceTapConfigurationPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*NetworkInterfaceTapConfigurationResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *NetworkInterfaceTapConfigurationsClient) ResumeCreateOrUpdate(token string) (NetworkInterfaceTapConfigurationPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("NetworkInterfaceTapConfigurationsClient.CreateOrUpdate", token, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &networkInterfaceTapConfigurationPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates or updates a Tap configuration in the specified NetworkInterface.
func (client *NetworkInterfaceTapConfigurationsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, networkInterfaceName string, tapConfigurationName string, tapConfigurationParameters NetworkInterfaceTapConfiguration, options *NetworkInterfaceTapConfigurationsCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.CreateOrUpdateCreateRequest(ctx, resourceGroupName, networkInterfaceName, tapConfigurationName, tapConfigurationParameters, options)
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
func (client *NetworkInterfaceTapConfigurationsClient) CreateOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, networkInterfaceName string, tapConfigurationName string, tapConfigurationParameters NetworkInterfaceTapConfiguration, options *NetworkInterfaceTapConfigurationsCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/tapConfigurations/{tapConfigurationName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkInterfaceName}", url.PathEscape(networkInterfaceName))
	urlPath = strings.ReplaceAll(urlPath, "{tapConfigurationName}", url.PathEscape(tapConfigurationName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(tapConfigurationParameters)
}

// CreateOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *NetworkInterfaceTapConfigurationsClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*NetworkInterfaceTapConfigurationResponse, error) {
	result := NetworkInterfaceTapConfigurationResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.NetworkInterfaceTapConfiguration)
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *NetworkInterfaceTapConfigurationsClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *NetworkInterfaceTapConfigurationsClient) BeginDelete(ctx context.Context, resourceGroupName string, networkInterfaceName string, tapConfigurationName string, options *NetworkInterfaceTapConfigurationsDeleteOptions) (*HTTPPollerResponse, error) {
	resp, err := client.Delete(ctx, resourceGroupName, networkInterfaceName, tapConfigurationName, options)
	if err != nil {
		return nil, err
	}
	result := &HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("NetworkInterfaceTapConfigurationsClient.Delete", "location", resp, client.DeleteHandleError)
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

func (client *NetworkInterfaceTapConfigurationsClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("NetworkInterfaceTapConfigurationsClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified tap configuration from the NetworkInterface.
func (client *NetworkInterfaceTapConfigurationsClient) Delete(ctx context.Context, resourceGroupName string, networkInterfaceName string, tapConfigurationName string, options *NetworkInterfaceTapConfigurationsDeleteOptions) (*azcore.Response, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, networkInterfaceName, tapConfigurationName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.DeleteHandleError(resp)
	}
	return resp, nil
}

// DeleteCreateRequest creates the Delete request.
func (client *NetworkInterfaceTapConfigurationsClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, networkInterfaceName string, tapConfigurationName string, options *NetworkInterfaceTapConfigurationsDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/tapConfigurations/{tapConfigurationName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkInterfaceName}", url.PathEscape(networkInterfaceName))
	urlPath = strings.ReplaceAll(urlPath, "{tapConfigurationName}", url.PathEscape(tapConfigurationName))
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
func (client *NetworkInterfaceTapConfigurationsClient) DeleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Get the specified tap configuration on a network interface.
func (client *NetworkInterfaceTapConfigurationsClient) Get(ctx context.Context, resourceGroupName string, networkInterfaceName string, tapConfigurationName string, options *NetworkInterfaceTapConfigurationsGetOptions) (*NetworkInterfaceTapConfigurationResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, networkInterfaceName, tapConfigurationName, options)
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
func (client *NetworkInterfaceTapConfigurationsClient) GetCreateRequest(ctx context.Context, resourceGroupName string, networkInterfaceName string, tapConfigurationName string, options *NetworkInterfaceTapConfigurationsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/tapConfigurations/{tapConfigurationName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkInterfaceName}", url.PathEscape(networkInterfaceName))
	urlPath = strings.ReplaceAll(urlPath, "{tapConfigurationName}", url.PathEscape(tapConfigurationName))
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
func (client *NetworkInterfaceTapConfigurationsClient) GetHandleResponse(resp *azcore.Response) (*NetworkInterfaceTapConfigurationResponse, error) {
	result := NetworkInterfaceTapConfigurationResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.NetworkInterfaceTapConfiguration)
}

// GetHandleError handles the Get error response.
func (client *NetworkInterfaceTapConfigurationsClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Get all Tap configurations in a network interface.
func (client *NetworkInterfaceTapConfigurationsClient) List(resourceGroupName string, networkInterfaceName string, options *NetworkInterfaceTapConfigurationsListOptions) NetworkInterfaceTapConfigurationListResultPager {
	return &networkInterfaceTapConfigurationListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, resourceGroupName, networkInterfaceName, options)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *NetworkInterfaceTapConfigurationListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.NetworkInterfaceTapConfigurationListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListCreateRequest creates the List request.
func (client *NetworkInterfaceTapConfigurationsClient) ListCreateRequest(ctx context.Context, resourceGroupName string, networkInterfaceName string, options *NetworkInterfaceTapConfigurationsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/tapConfigurations"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkInterfaceName}", url.PathEscape(networkInterfaceName))
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
func (client *NetworkInterfaceTapConfigurationsClient) ListHandleResponse(resp *azcore.Response) (*NetworkInterfaceTapConfigurationListResultResponse, error) {
	result := NetworkInterfaceTapConfigurationListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.NetworkInterfaceTapConfigurationListResult)
}

// ListHandleError handles the List error response.
func (client *NetworkInterfaceTapConfigurationsClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
