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

// PrivateDNSZoneGroupsOperations contains the methods for the PrivateDNSZoneGroups group.
type PrivateDNSZoneGroupsOperations interface {
	// BeginCreateOrUpdate - Creates or updates a private dns zone group in the specified private endpoint.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, privateEndpointName string, privateDnsZoneGroupName string, parameters PrivateDNSZoneGroup) (*PrivateDNSZoneGroupPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (PrivateDNSZoneGroupPoller, error)
	// BeginDelete - Deletes the specified private dns zone group.
	BeginDelete(ctx context.Context, resourceGroupName string, privateEndpointName string, privateDnsZoneGroupName string) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets the private dns zone group resource by specified private dns zone group name.
	Get(ctx context.Context, resourceGroupName string, privateEndpointName string, privateDnsZoneGroupName string) (*PrivateDNSZoneGroupResponse, error)
	// List - Gets all private dns zone groups in a private endpoint.
	List(privateEndpointName string, resourceGroupName string) PrivateDNSZoneGroupListResultPager
}

// PrivateDNSZoneGroupsClient implements the PrivateDNSZoneGroupsOperations interface.
// Don't use this type directly, use NewPrivateDNSZoneGroupsClient() instead.
type PrivateDNSZoneGroupsClient struct {
	*Client
	subscriptionID string
}

// NewPrivateDNSZoneGroupsClient creates a new instance of PrivateDNSZoneGroupsClient with the specified values.
func NewPrivateDNSZoneGroupsClient(c *Client, subscriptionID string) PrivateDNSZoneGroupsOperations {
	return &PrivateDNSZoneGroupsClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *PrivateDNSZoneGroupsClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

func (client *PrivateDNSZoneGroupsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, privateEndpointName string, privateDnsZoneGroupName string, parameters PrivateDNSZoneGroup) (*PrivateDNSZoneGroupPollerResponse, error) {
	resp, err := client.CreateOrUpdate(ctx, resourceGroupName, privateEndpointName, privateDnsZoneGroupName, parameters)
	if err != nil {
		return nil, err
	}
	result := &PrivateDNSZoneGroupPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("PrivateDNSZoneGroupsClient.CreateOrUpdate", "azure-async-operation", resp, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &privateDnsZoneGroupPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*PrivateDNSZoneGroupResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *PrivateDNSZoneGroupsClient) ResumeCreateOrUpdate(token string) (PrivateDNSZoneGroupPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("PrivateDNSZoneGroupsClient.CreateOrUpdate", token, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &privateDnsZoneGroupPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates or updates a private dns zone group in the specified private endpoint.
func (client *PrivateDNSZoneGroupsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, privateEndpointName string, privateDnsZoneGroupName string, parameters PrivateDNSZoneGroup) (*azcore.Response, error) {
	req, err := client.CreateOrUpdateCreateRequest(ctx, resourceGroupName, privateEndpointName, privateDnsZoneGroupName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.CreateOrUpdateHandleError(resp)
	}
	return resp, nil
}

// CreateOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PrivateDNSZoneGroupsClient) CreateOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, privateEndpointName string, privateDnsZoneGroupName string, parameters PrivateDNSZoneGroup) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}/privateDnsZoneGroups/{privateDnsZoneGroupName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointName}", url.PathEscape(privateEndpointName))
	urlPath = strings.ReplaceAll(urlPath, "{privateDnsZoneGroupName}", url.PathEscape(privateDnsZoneGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
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
func (client *PrivateDNSZoneGroupsClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*PrivateDNSZoneGroupResponse, error) {
	result := PrivateDNSZoneGroupResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PrivateDNSZoneGroup)
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *PrivateDNSZoneGroupsClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *PrivateDNSZoneGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, privateEndpointName string, privateDnsZoneGroupName string) (*HTTPPollerResponse, error) {
	resp, err := client.Delete(ctx, resourceGroupName, privateEndpointName, privateDnsZoneGroupName)
	if err != nil {
		return nil, err
	}
	result := &HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("PrivateDNSZoneGroupsClient.Delete", "location", resp, client.DeleteHandleError)
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

func (client *PrivateDNSZoneGroupsClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("PrivateDNSZoneGroupsClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified private dns zone group.
func (client *PrivateDNSZoneGroupsClient) Delete(ctx context.Context, resourceGroupName string, privateEndpointName string, privateDnsZoneGroupName string) (*azcore.Response, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, privateEndpointName, privateDnsZoneGroupName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.DeleteHandleError(resp)
	}
	return resp, nil
}

// DeleteCreateRequest creates the Delete request.
func (client *PrivateDNSZoneGroupsClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, privateEndpointName string, privateDnsZoneGroupName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}/privateDnsZoneGroups/{privateDnsZoneGroupName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointName}", url.PathEscape(privateEndpointName))
	urlPath = strings.ReplaceAll(urlPath, "{privateDnsZoneGroupName}", url.PathEscape(privateDnsZoneGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.u, urlPath))
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
func (client *PrivateDNSZoneGroupsClient) DeleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets the private dns zone group resource by specified private dns zone group name.
func (client *PrivateDNSZoneGroupsClient) Get(ctx context.Context, resourceGroupName string, privateEndpointName string, privateDnsZoneGroupName string) (*PrivateDNSZoneGroupResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, privateEndpointName, privateDnsZoneGroupName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
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
func (client *PrivateDNSZoneGroupsClient) GetCreateRequest(ctx context.Context, resourceGroupName string, privateEndpointName string, privateDnsZoneGroupName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}/privateDnsZoneGroups/{privateDnsZoneGroupName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointName}", url.PathEscape(privateEndpointName))
	urlPath = strings.ReplaceAll(urlPath, "{privateDnsZoneGroupName}", url.PathEscape(privateDnsZoneGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
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
func (client *PrivateDNSZoneGroupsClient) GetHandleResponse(resp *azcore.Response) (*PrivateDNSZoneGroupResponse, error) {
	result := PrivateDNSZoneGroupResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PrivateDNSZoneGroup)
}

// GetHandleError handles the Get error response.
func (client *PrivateDNSZoneGroupsClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Gets all private dns zone groups in a private endpoint.
func (client *PrivateDNSZoneGroupsClient) List(privateEndpointName string, resourceGroupName string) PrivateDNSZoneGroupListResultPager {
	return &privateDnsZoneGroupListResultPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, privateEndpointName, resourceGroupName)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *PrivateDNSZoneGroupListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.PrivateDNSZoneGroupListResult.NextLink)
		},
	}
}

// ListCreateRequest creates the List request.
func (client *PrivateDNSZoneGroupsClient) ListCreateRequest(ctx context.Context, privateEndpointName string, resourceGroupName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}/privateDnsZoneGroups"
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointName}", url.PathEscape(privateEndpointName))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
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
func (client *PrivateDNSZoneGroupsClient) ListHandleResponse(resp *azcore.Response) (*PrivateDNSZoneGroupListResultResponse, error) {
	result := PrivateDNSZoneGroupListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PrivateDNSZoneGroupListResult)
}

// ListHandleError handles the List error response.
func (client *PrivateDNSZoneGroupsClient) ListHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
