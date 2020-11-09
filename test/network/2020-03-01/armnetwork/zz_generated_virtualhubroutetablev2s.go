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

// VirtualHubRouteTableV2SOperations contains the methods for the VirtualHubRouteTableV2S group.
type VirtualHubRouteTableV2SOperations interface {
	// BeginCreateOrUpdate - Creates a VirtualHubRouteTableV2 resource if it doesn't exist else updates the existing VirtualHubRouteTableV2.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, virtualHubRouteTableV2Parameters VirtualHubRouteTableV2, options *VirtualHubRouteTableV2SCreateOrUpdateOptions) (*VirtualHubRouteTableV2PollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (VirtualHubRouteTableV2Poller, error)
	// BeginDelete - Deletes a VirtualHubRouteTableV2.
	BeginDelete(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, options *VirtualHubRouteTableV2SDeleteOptions) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Retrieves the details of a VirtualHubRouteTableV2.
	Get(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, options *VirtualHubRouteTableV2SGetOptions) (*VirtualHubRouteTableV2Response, error)
	// List - Retrieves the details of all VirtualHubRouteTableV2s.
	List(resourceGroupName string, virtualHubName string, options *VirtualHubRouteTableV2SListOptions) ListVirtualHubRouteTableV2SResultPager
}

// VirtualHubRouteTableV2SClient implements the VirtualHubRouteTableV2SOperations interface.
// Don't use this type directly, use NewVirtualHubRouteTableV2SClient() instead.
type VirtualHubRouteTableV2SClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewVirtualHubRouteTableV2SClient creates a new instance of VirtualHubRouteTableV2SClient with the specified values.
func NewVirtualHubRouteTableV2SClient(con *armcore.Connection, subscriptionID string) VirtualHubRouteTableV2SOperations {
	return &VirtualHubRouteTableV2SClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client *VirtualHubRouteTableV2SClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

func (client *VirtualHubRouteTableV2SClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, virtualHubRouteTableV2Parameters VirtualHubRouteTableV2, options *VirtualHubRouteTableV2SCreateOrUpdateOptions) (*VirtualHubRouteTableV2PollerResponse, error) {
	resp, err := client.CreateOrUpdate(ctx, resourceGroupName, virtualHubName, routeTableName, virtualHubRouteTableV2Parameters, options)
	if err != nil {
		return nil, err
	}
	result := &VirtualHubRouteTableV2PollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualHubRouteTableV2SClient.CreateOrUpdate", "azure-async-operation", resp, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &virtualHubRouteTableV2Poller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*VirtualHubRouteTableV2Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *VirtualHubRouteTableV2SClient) ResumeCreateOrUpdate(token string) (VirtualHubRouteTableV2Poller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualHubRouteTableV2SClient.CreateOrUpdate", token, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &virtualHubRouteTableV2Poller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates a VirtualHubRouteTableV2 resource if it doesn't exist else updates the existing VirtualHubRouteTableV2.
func (client *VirtualHubRouteTableV2SClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, virtualHubRouteTableV2Parameters VirtualHubRouteTableV2, options *VirtualHubRouteTableV2SCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.CreateOrUpdateCreateRequest(ctx, resourceGroupName, virtualHubName, routeTableName, virtualHubRouteTableV2Parameters, options)
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
func (client *VirtualHubRouteTableV2SClient) CreateOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, virtualHubRouteTableV2Parameters VirtualHubRouteTableV2, options *VirtualHubRouteTableV2SCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/routeTables/{routeTableName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	urlPath = strings.ReplaceAll(urlPath, "{routeTableName}", url.PathEscape(routeTableName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(virtualHubRouteTableV2Parameters)
}

// CreateOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *VirtualHubRouteTableV2SClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*VirtualHubRouteTableV2Response, error) {
	result := VirtualHubRouteTableV2Response{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualHubRouteTableV2)
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *VirtualHubRouteTableV2SClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *VirtualHubRouteTableV2SClient) BeginDelete(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, options *VirtualHubRouteTableV2SDeleteOptions) (*HTTPPollerResponse, error) {
	resp, err := client.Delete(ctx, resourceGroupName, virtualHubName, routeTableName, options)
	if err != nil {
		return nil, err
	}
	result := &HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualHubRouteTableV2SClient.Delete", "location", resp, client.DeleteHandleError)
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

func (client *VirtualHubRouteTableV2SClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualHubRouteTableV2SClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes a VirtualHubRouteTableV2.
func (client *VirtualHubRouteTableV2SClient) Delete(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, options *VirtualHubRouteTableV2SDeleteOptions) (*azcore.Response, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, virtualHubName, routeTableName, options)
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
func (client *VirtualHubRouteTableV2SClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, options *VirtualHubRouteTableV2SDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/routeTables/{routeTableName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	urlPath = strings.ReplaceAll(urlPath, "{routeTableName}", url.PathEscape(routeTableName))
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
func (client *VirtualHubRouteTableV2SClient) DeleteHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Retrieves the details of a VirtualHubRouteTableV2.
func (client *VirtualHubRouteTableV2SClient) Get(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, options *VirtualHubRouteTableV2SGetOptions) (*VirtualHubRouteTableV2Response, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, virtualHubName, routeTableName, options)
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
func (client *VirtualHubRouteTableV2SClient) GetCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, options *VirtualHubRouteTableV2SGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/routeTables/{routeTableName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	urlPath = strings.ReplaceAll(urlPath, "{routeTableName}", url.PathEscape(routeTableName))
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
func (client *VirtualHubRouteTableV2SClient) GetHandleResponse(resp *azcore.Response) (*VirtualHubRouteTableV2Response, error) {
	result := VirtualHubRouteTableV2Response{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualHubRouteTableV2)
}

// GetHandleError handles the Get error response.
func (client *VirtualHubRouteTableV2SClient) GetHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Retrieves the details of all VirtualHubRouteTableV2s.
func (client *VirtualHubRouteTableV2SClient) List(resourceGroupName string, virtualHubName string, options *VirtualHubRouteTableV2SListOptions) ListVirtualHubRouteTableV2SResultPager {
	return &listVirtualHubRouteTableV2SResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, resourceGroupName, virtualHubName, options)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *ListVirtualHubRouteTableV2SResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ListVirtualHubRouteTableV2SResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListCreateRequest creates the List request.
func (client *VirtualHubRouteTableV2SClient) ListCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, options *VirtualHubRouteTableV2SListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/routeTables"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
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
func (client *VirtualHubRouteTableV2SClient) ListHandleResponse(resp *azcore.Response) (*ListVirtualHubRouteTableV2SResultResponse, error) {
	result := ListVirtualHubRouteTableV2SResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ListVirtualHubRouteTableV2SResult)
}

// ListHandleError handles the List error response.
func (client *VirtualHubRouteTableV2SClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
