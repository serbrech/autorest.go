// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// AvailablePrivateEndpointTypesOperations contains the methods for the AvailablePrivateEndpointTypes group.
type AvailablePrivateEndpointTypesOperations interface {
	// List - Returns all of the resource types that can be linked to a Private Endpoint in this subscription in this region.
	List(location string) AvailablePrivateEndpointTypesResultPager
	// ListByResourceGroup - Returns all of the resource types that can be linked to a Private Endpoint in this subscription in this region.
	ListByResourceGroup(location string, resourceGroupName string) AvailablePrivateEndpointTypesResultPager
}

// AvailablePrivateEndpointTypesClient implements the AvailablePrivateEndpointTypesOperations interface.
// Don't use this type directly, use NewAvailablePrivateEndpointTypesClient() instead.
type AvailablePrivateEndpointTypesClient struct {
	*Client
	subscriptionID string
}

// NewAvailablePrivateEndpointTypesClient creates a new instance of AvailablePrivateEndpointTypesClient with the specified values.
func NewAvailablePrivateEndpointTypesClient(c *Client, subscriptionID string) AvailablePrivateEndpointTypesOperations {
	return &AvailablePrivateEndpointTypesClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *AvailablePrivateEndpointTypesClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// List - Returns all of the resource types that can be linked to a Private Endpoint in this subscription in this region.
func (client *AvailablePrivateEndpointTypesClient) List(location string) AvailablePrivateEndpointTypesResultPager {
	return &availablePrivateEndpointTypesResultPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, location)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *AvailablePrivateEndpointTypesResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.AvailablePrivateEndpointTypesResult.NextLink)
		},
	}
}

// ListCreateRequest creates the List request.
func (client *AvailablePrivateEndpointTypesClient) ListCreateRequest(ctx context.Context, location string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/availablePrivateEndpointTypes"
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
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
func (client *AvailablePrivateEndpointTypesClient) ListHandleResponse(resp *azcore.Response) (*AvailablePrivateEndpointTypesResultResponse, error) {
	result := AvailablePrivateEndpointTypesResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.AvailablePrivateEndpointTypesResult)
}

// ListHandleError handles the List error response.
func (client *AvailablePrivateEndpointTypesClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListByResourceGroup - Returns all of the resource types that can be linked to a Private Endpoint in this subscription in this region.
func (client *AvailablePrivateEndpointTypesClient) ListByResourceGroup(location string, resourceGroupName string) AvailablePrivateEndpointTypesResultPager {
	return &availablePrivateEndpointTypesResultPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListByResourceGroupCreateRequest(ctx, location, resourceGroupName)
		},
		responder: client.ListByResourceGroupHandleResponse,
		errorer:   client.ListByResourceGroupHandleError,
		advancer: func(ctx context.Context, resp *AvailablePrivateEndpointTypesResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.AvailablePrivateEndpointTypesResult.NextLink)
		},
	}
}

// ListByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *AvailablePrivateEndpointTypesClient) ListByResourceGroupCreateRequest(ctx context.Context, location string, resourceGroupName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/locations/{location}/availablePrivateEndpointTypes"
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
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

// ListByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *AvailablePrivateEndpointTypesClient) ListByResourceGroupHandleResponse(resp *azcore.Response) (*AvailablePrivateEndpointTypesResultResponse, error) {
	result := AvailablePrivateEndpointTypesResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.AvailablePrivateEndpointTypesResult)
}

// ListByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *AvailablePrivateEndpointTypesClient) ListByResourceGroupHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
