// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"
)

// RouteFiltersOperations contains the methods for the RouteFilters group.
type RouteFiltersOperations interface {
	// BeginCreateOrUpdate - Creates or updates a route filter in a specified resource group.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, routeFilterName string, routeFilterParameters RouteFilter) (*RouteFilterPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (RouteFilterPoller, error)
	// BeginDelete - Deletes the specified route filter.
	BeginDelete(ctx context.Context, resourceGroupName string, routeFilterName string) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets the specified route filter.
	Get(ctx context.Context, resourceGroupName string, routeFilterName string, routeFiltersGetOptions *RouteFiltersGetOptions) (*RouteFilterResponse, error)
	// List - Gets all route filters in a subscription.
	List() (RouteFilterListResultPager, error)
	// ListByResourceGroup - Gets all route filters in a resource group.
	ListByResourceGroup(resourceGroupName string) (RouteFilterListResultPager, error)
	// UpdateTags - Updates tags of a route filter.
	UpdateTags(ctx context.Context, resourceGroupName string, routeFilterName string, parameters TagsObject) (*RouteFilterResponse, error)
}

// RouteFiltersClient implements the RouteFiltersOperations interface.
// Don't use this type directly, use NewRouteFiltersClient() instead.
type RouteFiltersClient struct {
	*Client
	subscriptionID string
}

// NewRouteFiltersClient creates a new instance of RouteFiltersClient with the specified values.
func NewRouteFiltersClient(c *Client, subscriptionID string) RouteFiltersOperations {
	return &RouteFiltersClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *RouteFiltersClient) Do(ctx context.Context, req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(ctx, req)
}

// CreateOrUpdate - Creates or updates a route filter in a specified resource group.
func (client *RouteFiltersClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, routeFilterName string, routeFilterParameters RouteFilter) (*RouteFilterPollerResponse, error) {
	req, err := client.CreateOrUpdateCreateRequest(resourceGroupName, routeFilterName, routeFilterParameters)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.CreateOrUpdateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := armcore.NewPoller("RouteFiltersClient.CreateOrUpdate", "azure-async-operation", resp, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &routeFilterPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*RouteFilterResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *RouteFiltersClient) ResumeCreateOrUpdate(token string) (RouteFilterPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("RouteFiltersClient.CreateOrUpdate", token, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &routeFilterPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// CreateOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *RouteFiltersClient) CreateOrUpdateCreateRequest(resourceGroupName string, routeFilterName string, routeFilterParameters RouteFilter) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{routeFilterName}", url.PathEscape(routeFilterName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(routeFilterParameters)
}

// CreateOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *RouteFiltersClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*RouteFilterPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated, http.StatusNoContent) {
		return nil, client.CreateOrUpdateHandleError(resp)
	}
	return &RouteFilterPollerResponse{RawResponse: resp.Response}, nil
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *RouteFiltersClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Delete - Deletes the specified route filter.
func (client *RouteFiltersClient) BeginDelete(ctx context.Context, resourceGroupName string, routeFilterName string) (*HTTPPollerResponse, error) {
	req, err := client.DeleteCreateRequest(resourceGroupName, routeFilterName)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.DeleteHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := armcore.NewPoller("RouteFiltersClient.Delete", "location", resp, client.DeleteHandleError)
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

func (client *RouteFiltersClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("RouteFiltersClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// DeleteCreateRequest creates the Delete request.
func (client *RouteFiltersClient) DeleteCreateRequest(resourceGroupName string, routeFilterName string) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{routeFilterName}", url.PathEscape(routeFilterName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodDelete, *u)
	return req, nil
}

// DeleteHandleResponse handles the Delete response.
func (client *RouteFiltersClient) DeleteHandleResponse(resp *azcore.Response) (*HTTPPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.DeleteHandleError(resp)
	}
	return &HTTPPollerResponse{RawResponse: resp.Response}, nil
}

// DeleteHandleError handles the Delete error response.
func (client *RouteFiltersClient) DeleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Get - Gets the specified route filter.
func (client *RouteFiltersClient) Get(ctx context.Context, resourceGroupName string, routeFilterName string, routeFiltersGetOptions *RouteFiltersGetOptions) (*RouteFilterResponse, error) {
	req, err := client.GetCreateRequest(resourceGroupName, routeFilterName, routeFiltersGetOptions)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCreateRequest creates the Get request.
func (client *RouteFiltersClient) GetCreateRequest(resourceGroupName string, routeFilterName string, routeFiltersGetOptions *RouteFiltersGetOptions) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{routeFilterName}", url.PathEscape(routeFilterName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	if routeFiltersGetOptions != nil && routeFiltersGetOptions.Expand != nil {
		query.Set("$expand", *routeFiltersGetOptions.Expand)
	}
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// GetHandleResponse handles the Get response.
func (client *RouteFiltersClient) GetHandleResponse(resp *azcore.Response) (*RouteFilterResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetHandleError(resp)
	}
	result := RouteFilterResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.RouteFilter)
}

// GetHandleError handles the Get error response.
func (client *RouteFiltersClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// List - Gets all route filters in a subscription.
func (client *RouteFiltersClient) List() (RouteFilterListResultPager, error) {
	req, err := client.ListCreateRequest()
	if err != nil {
		return nil, err
	}
	return &routeFilterListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.ListHandleResponse,
		advancer: func(resp *RouteFilterListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.RouteFilterListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.RouteFilterListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// ListCreateRequest creates the List request.
func (client *RouteFiltersClient) ListCreateRequest() (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/routeFilters"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// ListHandleResponse handles the List response.
func (client *RouteFiltersClient) ListHandleResponse(resp *azcore.Response) (*RouteFilterListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.ListHandleError(resp)
	}
	result := RouteFilterListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.RouteFilterListResult)
}

// ListHandleError handles the List error response.
func (client *RouteFiltersClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// ListByResourceGroup - Gets all route filters in a resource group.
func (client *RouteFiltersClient) ListByResourceGroup(resourceGroupName string) (RouteFilterListResultPager, error) {
	req, err := client.ListByResourceGroupCreateRequest(resourceGroupName)
	if err != nil {
		return nil, err
	}
	return &routeFilterListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.ListByResourceGroupHandleResponse,
		advancer: func(resp *RouteFilterListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.RouteFilterListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.RouteFilterListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// ListByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *RouteFiltersClient) ListByResourceGroupCreateRequest(resourceGroupName string) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// ListByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *RouteFiltersClient) ListByResourceGroupHandleResponse(resp *azcore.Response) (*RouteFilterListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.ListByResourceGroupHandleError(resp)
	}
	result := RouteFilterListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.RouteFilterListResult)
}

// ListByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *RouteFiltersClient) ListByResourceGroupHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// UpdateTags - Updates tags of a route filter.
func (client *RouteFiltersClient) UpdateTags(ctx context.Context, resourceGroupName string, routeFilterName string, parameters TagsObject) (*RouteFilterResponse, error) {
	req, err := client.UpdateTagsCreateRequest(resourceGroupName, routeFilterName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.UpdateTagsHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateTagsCreateRequest creates the UpdateTags request.
func (client *RouteFiltersClient) UpdateTagsCreateRequest(resourceGroupName string, routeFilterName string, parameters TagsObject) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{routeFilterName}", url.PathEscape(routeFilterName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPatch, *u)
	return req, req.MarshalAsJSON(parameters)
}

// UpdateTagsHandleResponse handles the UpdateTags response.
func (client *RouteFiltersClient) UpdateTagsHandleResponse(resp *azcore.Response) (*RouteFilterResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.UpdateTagsHandleError(resp)
	}
	result := RouteFilterResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.RouteFilter)
}

// UpdateTagsHandleError handles the UpdateTags error response.
func (client *RouteFiltersClient) UpdateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}