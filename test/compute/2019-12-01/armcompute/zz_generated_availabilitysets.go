// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// AvailabilitySetsOperations contains the methods for the AvailabilitySets group.
type AvailabilitySetsOperations interface {
	// CreateOrUpdate - Create or update an availability set.
	CreateOrUpdate(ctx context.Context, resourceGroupName string, availabilitySetName string, parameters AvailabilitySet, options *AvailabilitySetsCreateOrUpdateOptions) (*AvailabilitySetResponse, error)
	// Delete - Delete an availability set.
	Delete(ctx context.Context, resourceGroupName string, availabilitySetName string, options *AvailabilitySetsDeleteOptions) (*http.Response, error)
	// Get - Retrieves information about an availability set.
	Get(ctx context.Context, resourceGroupName string, availabilitySetName string, options *AvailabilitySetsGetOptions) (*AvailabilitySetResponse, error)
	// List - Lists all availability sets in a resource group.
	List(resourceGroupName string, options *AvailabilitySetsListOptions) AvailabilitySetListResultPager
	// ListAvailableSizes - Lists all available virtual machine sizes that can be used to create a new virtual machine in an existing availability set.
	ListAvailableSizes(ctx context.Context, resourceGroupName string, availabilitySetName string, options *AvailabilitySetsListAvailableSizesOptions) (*VirtualMachineSizeListResultResponse, error)
	// ListBySubscription - Lists all availability sets in a subscription.
	ListBySubscription(options *AvailabilitySetsListBySubscriptionOptions) AvailabilitySetListResultPager
	// Update - Update an availability set.
	Update(ctx context.Context, resourceGroupName string, availabilitySetName string, parameters AvailabilitySetUpdate, options *AvailabilitySetsUpdateOptions) (*AvailabilitySetResponse, error)
}

// AvailabilitySetsClient implements the AvailabilitySetsOperations interface.
// Don't use this type directly, use NewAvailabilitySetsClient() instead.
type AvailabilitySetsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewAvailabilitySetsClient creates a new instance of AvailabilitySetsClient with the specified values.
func NewAvailabilitySetsClient(con *armcore.Connection, subscriptionID string) AvailabilitySetsOperations {
	return &AvailabilitySetsClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client *AvailabilitySetsClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// CreateOrUpdate - Create or update an availability set.
func (client *AvailabilitySetsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, availabilitySetName string, parameters AvailabilitySet, options *AvailabilitySetsCreateOrUpdateOptions) (*AvailabilitySetResponse, error) {
	req, err := client.CreateOrUpdateCreateRequest(ctx, resourceGroupName, availabilitySetName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.CreateOrUpdateHandleError(resp)
	}
	result, err := client.CreateOrUpdateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AvailabilitySetsClient) CreateOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, availabilitySetName string, parameters AvailabilitySet, options *AvailabilitySetsCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{availabilitySetName}", url.PathEscape(availabilitySetName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// CreateOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *AvailabilitySetsClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*AvailabilitySetResponse, error) {
	result := AvailabilitySetResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.AvailabilitySet)
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *AvailabilitySetsClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// Delete - Delete an availability set.
func (client *AvailabilitySetsClient) Delete(ctx context.Context, resourceGroupName string, availabilitySetName string, options *AvailabilitySetsDeleteOptions) (*http.Response, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, availabilitySetName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNoContent) {
		return nil, client.DeleteHandleError(resp)
	}
	return resp.Response, nil
}

// DeleteCreateRequest creates the Delete request.
func (client *AvailabilitySetsClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, availabilitySetName string, options *AvailabilitySetsDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{availabilitySetName}", url.PathEscape(availabilitySetName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	return req, nil
}

// DeleteHandleError handles the Delete error response.
func (client *AvailabilitySetsClient) DeleteHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// Get - Retrieves information about an availability set.
func (client *AvailabilitySetsClient) Get(ctx context.Context, resourceGroupName string, availabilitySetName string, options *AvailabilitySetsGetOptions) (*AvailabilitySetResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, availabilitySetName, options)
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
func (client *AvailabilitySetsClient) GetCreateRequest(ctx context.Context, resourceGroupName string, availabilitySetName string, options *AvailabilitySetsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{availabilitySetName}", url.PathEscape(availabilitySetName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetHandleResponse handles the Get response.
func (client *AvailabilitySetsClient) GetHandleResponse(resp *azcore.Response) (*AvailabilitySetResponse, error) {
	result := AvailabilitySetResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.AvailabilitySet)
}

// GetHandleError handles the Get error response.
func (client *AvailabilitySetsClient) GetHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// List - Lists all availability sets in a resource group.
func (client *AvailabilitySetsClient) List(resourceGroupName string, options *AvailabilitySetsListOptions) AvailabilitySetListResultPager {
	return &availabilitySetListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *AvailabilitySetListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.AvailabilitySetListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListCreateRequest creates the List request.
func (client *AvailabilitySetsClient) ListCreateRequest(ctx context.Context, resourceGroupName string, options *AvailabilitySetsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListHandleResponse handles the List response.
func (client *AvailabilitySetsClient) ListHandleResponse(resp *azcore.Response) (*AvailabilitySetListResultResponse, error) {
	result := AvailabilitySetListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.AvailabilitySetListResult)
}

// ListHandleError handles the List error response.
func (client *AvailabilitySetsClient) ListHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// ListAvailableSizes - Lists all available virtual machine sizes that can be used to create a new virtual machine in an existing availability set.
func (client *AvailabilitySetsClient) ListAvailableSizes(ctx context.Context, resourceGroupName string, availabilitySetName string, options *AvailabilitySetsListAvailableSizesOptions) (*VirtualMachineSizeListResultResponse, error) {
	req, err := client.ListAvailableSizesCreateRequest(ctx, resourceGroupName, availabilitySetName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.ListAvailableSizesHandleError(resp)
	}
	result, err := client.ListAvailableSizesHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListAvailableSizesCreateRequest creates the ListAvailableSizes request.
func (client *AvailabilitySetsClient) ListAvailableSizesCreateRequest(ctx context.Context, resourceGroupName string, availabilitySetName string, options *AvailabilitySetsListAvailableSizesOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}/vmSizes"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{availabilitySetName}", url.PathEscape(availabilitySetName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListAvailableSizesHandleResponse handles the ListAvailableSizes response.
func (client *AvailabilitySetsClient) ListAvailableSizesHandleResponse(resp *azcore.Response) (*VirtualMachineSizeListResultResponse, error) {
	result := VirtualMachineSizeListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualMachineSizeListResult)
}

// ListAvailableSizesHandleError handles the ListAvailableSizes error response.
func (client *AvailabilitySetsClient) ListAvailableSizesHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// ListBySubscription - Lists all availability sets in a subscription.
func (client *AvailabilitySetsClient) ListBySubscription(options *AvailabilitySetsListBySubscriptionOptions) AvailabilitySetListResultPager {
	return &availabilitySetListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListBySubscriptionCreateRequest(ctx, options)
		},
		responder: client.ListBySubscriptionHandleResponse,
		errorer:   client.ListBySubscriptionHandleError,
		advancer: func(ctx context.Context, resp *AvailabilitySetListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.AvailabilitySetListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *AvailabilitySetsClient) ListBySubscriptionCreateRequest(ctx context.Context, options *AvailabilitySetsListBySubscriptionOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/availabilitySets"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-12-01")
	if options != nil && options.Expand != nil {
		query.Set("$expand", *options.Expand)
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *AvailabilitySetsClient) ListBySubscriptionHandleResponse(resp *azcore.Response) (*AvailabilitySetListResultResponse, error) {
	result := AvailabilitySetListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.AvailabilitySetListResult)
}

// ListBySubscriptionHandleError handles the ListBySubscription error response.
func (client *AvailabilitySetsClient) ListBySubscriptionHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// Update - Update an availability set.
func (client *AvailabilitySetsClient) Update(ctx context.Context, resourceGroupName string, availabilitySetName string, parameters AvailabilitySetUpdate, options *AvailabilitySetsUpdateOptions) (*AvailabilitySetResponse, error) {
	req, err := client.UpdateCreateRequest(ctx, resourceGroupName, availabilitySetName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.UpdateHandleError(resp)
	}
	result, err := client.UpdateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateCreateRequest creates the Update request.
func (client *AvailabilitySetsClient) UpdateCreateRequest(ctx context.Context, resourceGroupName string, availabilitySetName string, parameters AvailabilitySetUpdate, options *AvailabilitySetsUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{availabilitySetName}", url.PathEscape(availabilitySetName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// UpdateHandleResponse handles the Update response.
func (client *AvailabilitySetsClient) UpdateHandleResponse(resp *azcore.Response) (*AvailabilitySetResponse, error) {
	result := AvailabilitySetResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.AvailabilitySet)
}

// UpdateHandleError handles the Update error response.
func (client *AvailabilitySetsClient) UpdateHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
