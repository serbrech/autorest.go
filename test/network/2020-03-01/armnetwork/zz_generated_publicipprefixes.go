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

// PublicIPPrefixesOperations contains the methods for the PublicIPPrefixes group.
type PublicIPPrefixesOperations interface {
	// BeginCreateOrUpdate - Creates or updates a static or dynamic public IP prefix.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, publicIPPrefixName string, parameters PublicIPPrefix, options *PublicIPPrefixesCreateOrUpdateOptions) (*PublicIPPrefixPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (PublicIPPrefixPoller, error)
	// BeginDelete - Deletes the specified public IP prefix.
	BeginDelete(ctx context.Context, resourceGroupName string, publicIPPrefixName string, options *PublicIPPrefixesDeleteOptions) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets the specified public IP prefix in a specified resource group.
	Get(ctx context.Context, resourceGroupName string, publicIPPrefixName string, options *PublicIPPrefixesGetOptions) (*PublicIPPrefixResponse, error)
	// List - Gets all public IP prefixes in a resource group.
	List(resourceGroupName string, options *PublicIPPrefixesListOptions) PublicIPPrefixListResultPager
	// ListAll - Gets all the public IP prefixes in a subscription.
	ListAll(options *PublicIPPrefixesListAllOptions) PublicIPPrefixListResultPager
	// UpdateTags - Updates public IP prefix tags.
	UpdateTags(ctx context.Context, resourceGroupName string, publicIPPrefixName string, parameters TagsObject, options *PublicIPPrefixesUpdateTagsOptions) (*PublicIPPrefixResponse, error)
}

// PublicIPPrefixesClient implements the PublicIPPrefixesOperations interface.
// Don't use this type directly, use NewPublicIPPrefixesClient() instead.
type PublicIPPrefixesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewPublicIPPrefixesClient creates a new instance of PublicIPPrefixesClient with the specified values.
func NewPublicIPPrefixesClient(con *armcore.Connection, subscriptionID string) PublicIPPrefixesOperations {
	return &PublicIPPrefixesClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client *PublicIPPrefixesClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

func (client *PublicIPPrefixesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, publicIPPrefixName string, parameters PublicIPPrefix, options *PublicIPPrefixesCreateOrUpdateOptions) (*PublicIPPrefixPollerResponse, error) {
	resp, err := client.CreateOrUpdate(ctx, resourceGroupName, publicIPPrefixName, parameters, options)
	if err != nil {
		return nil, err
	}
	result := &PublicIPPrefixPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("PublicIPPrefixesClient.CreateOrUpdate", "location", resp, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &publicIPPrefixPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*PublicIPPrefixResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *PublicIPPrefixesClient) ResumeCreateOrUpdate(token string) (PublicIPPrefixPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("PublicIPPrefixesClient.CreateOrUpdate", token, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &publicIPPrefixPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates or updates a static or dynamic public IP prefix.
func (client *PublicIPPrefixesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, publicIPPrefixName string, parameters PublicIPPrefix, options *PublicIPPrefixesCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.CreateOrUpdateCreateRequest(ctx, resourceGroupName, publicIPPrefixName, parameters, options)
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
func (client *PublicIPPrefixesClient) CreateOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, publicIPPrefixName string, parameters PublicIPPrefix, options *PublicIPPrefixesCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIpPrefixName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{publicIpPrefixName}", url.PathEscape(publicIPPrefixName))
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
func (client *PublicIPPrefixesClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*PublicIPPrefixResponse, error) {
	result := PublicIPPrefixResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PublicIPPrefix)
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *PublicIPPrefixesClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *PublicIPPrefixesClient) BeginDelete(ctx context.Context, resourceGroupName string, publicIPPrefixName string, options *PublicIPPrefixesDeleteOptions) (*HTTPPollerResponse, error) {
	resp, err := client.Delete(ctx, resourceGroupName, publicIPPrefixName, options)
	if err != nil {
		return nil, err
	}
	result := &HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("PublicIPPrefixesClient.Delete", "location", resp, client.DeleteHandleError)
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

func (client *PublicIPPrefixesClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("PublicIPPrefixesClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified public IP prefix.
func (client *PublicIPPrefixesClient) Delete(ctx context.Context, resourceGroupName string, publicIPPrefixName string, options *PublicIPPrefixesDeleteOptions) (*azcore.Response, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, publicIPPrefixName, options)
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
func (client *PublicIPPrefixesClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, publicIPPrefixName string, options *PublicIPPrefixesDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIpPrefixName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{publicIpPrefixName}", url.PathEscape(publicIPPrefixName))
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
func (client *PublicIPPrefixesClient) DeleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets the specified public IP prefix in a specified resource group.
func (client *PublicIPPrefixesClient) Get(ctx context.Context, resourceGroupName string, publicIPPrefixName string, options *PublicIPPrefixesGetOptions) (*PublicIPPrefixResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, publicIPPrefixName, options)
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
func (client *PublicIPPrefixesClient) GetCreateRequest(ctx context.Context, resourceGroupName string, publicIPPrefixName string, options *PublicIPPrefixesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIpPrefixName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{publicIpPrefixName}", url.PathEscape(publicIPPrefixName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	if options != nil && options.Expand != nil {
		query.Set("$expand", *options.Expand)
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetHandleResponse handles the Get response.
func (client *PublicIPPrefixesClient) GetHandleResponse(resp *azcore.Response) (*PublicIPPrefixResponse, error) {
	result := PublicIPPrefixResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PublicIPPrefix)
}

// GetHandleError handles the Get error response.
func (client *PublicIPPrefixesClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Gets all public IP prefixes in a resource group.
func (client *PublicIPPrefixesClient) List(resourceGroupName string, options *PublicIPPrefixesListOptions) PublicIPPrefixListResultPager {
	return &publicIPPrefixListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *PublicIPPrefixListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.PublicIPPrefixListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListCreateRequest creates the List request.
func (client *PublicIPPrefixesClient) ListCreateRequest(ctx context.Context, resourceGroupName string, options *PublicIPPrefixesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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
func (client *PublicIPPrefixesClient) ListHandleResponse(resp *azcore.Response) (*PublicIPPrefixListResultResponse, error) {
	result := PublicIPPrefixListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PublicIPPrefixListResult)
}

// ListHandleError handles the List error response.
func (client *PublicIPPrefixesClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListAll - Gets all the public IP prefixes in a subscription.
func (client *PublicIPPrefixesClient) ListAll(options *PublicIPPrefixesListAllOptions) PublicIPPrefixListResultPager {
	return &publicIPPrefixListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListAllCreateRequest(ctx, options)
		},
		responder: client.ListAllHandleResponse,
		errorer:   client.ListAllHandleError,
		advancer: func(ctx context.Context, resp *PublicIPPrefixListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.PublicIPPrefixListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListAllCreateRequest creates the ListAll request.
func (client *PublicIPPrefixesClient) ListAllCreateRequest(ctx context.Context, options *PublicIPPrefixesListAllOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/publicIPPrefixes"
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

// ListAllHandleResponse handles the ListAll response.
func (client *PublicIPPrefixesClient) ListAllHandleResponse(resp *azcore.Response) (*PublicIPPrefixListResultResponse, error) {
	result := PublicIPPrefixListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PublicIPPrefixListResult)
}

// ListAllHandleError handles the ListAll error response.
func (client *PublicIPPrefixesClient) ListAllHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// UpdateTags - Updates public IP prefix tags.
func (client *PublicIPPrefixesClient) UpdateTags(ctx context.Context, resourceGroupName string, publicIPPrefixName string, parameters TagsObject, options *PublicIPPrefixesUpdateTagsOptions) (*PublicIPPrefixResponse, error) {
	req, err := client.UpdateTagsCreateRequest(ctx, resourceGroupName, publicIPPrefixName, parameters, options)
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
func (client *PublicIPPrefixesClient) UpdateTagsCreateRequest(ctx context.Context, resourceGroupName string, publicIPPrefixName string, parameters TagsObject, options *PublicIPPrefixesUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIpPrefixName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{publicIpPrefixName}", url.PathEscape(publicIPPrefixName))
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
func (client *PublicIPPrefixesClient) UpdateTagsHandleResponse(resp *azcore.Response) (*PublicIPPrefixResponse, error) {
	result := PublicIPPrefixResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PublicIPPrefix)
}

// UpdateTagsHandleError handles the UpdateTags error response.
func (client *PublicIPPrefixesClient) UpdateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
