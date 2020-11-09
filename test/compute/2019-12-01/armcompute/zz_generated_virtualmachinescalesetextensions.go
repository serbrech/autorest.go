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
	"time"
)

// VirtualMachineScaleSetExtensionsOperations contains the methods for the VirtualMachineScaleSetExtensions group.
type VirtualMachineScaleSetExtensionsOperations interface {
	// BeginCreateOrUpdate - The operation to create or update an extension.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, extensionParameters VirtualMachineScaleSetExtension, options *VirtualMachineScaleSetExtensionsCreateOrUpdateOptions) (*VirtualMachineScaleSetExtensionPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (VirtualMachineScaleSetExtensionPoller, error)
	// BeginDelete - The operation to delete the extension.
	BeginDelete(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, options *VirtualMachineScaleSetExtensionsDeleteOptions) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - The operation to get the extension.
	Get(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, options *VirtualMachineScaleSetExtensionsGetOptions) (*VirtualMachineScaleSetExtensionResponse, error)
	// List - Gets a list of all extensions in a VM scale set.
	List(resourceGroupName string, vmScaleSetName string, options *VirtualMachineScaleSetExtensionsListOptions) VirtualMachineScaleSetExtensionListResultPager
	// BeginUpdate - The operation to update an extension.
	BeginUpdate(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, extensionParameters VirtualMachineScaleSetExtensionUpdate, options *VirtualMachineScaleSetExtensionsUpdateOptions) (*VirtualMachineScaleSetExtensionPollerResponse, error)
	// ResumeUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeUpdate(token string) (VirtualMachineScaleSetExtensionPoller, error)
}

// VirtualMachineScaleSetExtensionsClient implements the VirtualMachineScaleSetExtensionsOperations interface.
// Don't use this type directly, use NewVirtualMachineScaleSetExtensionsClient() instead.
type VirtualMachineScaleSetExtensionsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewVirtualMachineScaleSetExtensionsClient creates a new instance of VirtualMachineScaleSetExtensionsClient with the specified values.
func NewVirtualMachineScaleSetExtensionsClient(con *armcore.Connection, subscriptionID string) VirtualMachineScaleSetExtensionsOperations {
	return &VirtualMachineScaleSetExtensionsClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client *VirtualMachineScaleSetExtensionsClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

func (client *VirtualMachineScaleSetExtensionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, extensionParameters VirtualMachineScaleSetExtension, options *VirtualMachineScaleSetExtensionsCreateOrUpdateOptions) (*VirtualMachineScaleSetExtensionPollerResponse, error) {
	resp, err := client.CreateOrUpdate(ctx, resourceGroupName, vmScaleSetName, vmssExtensionName, extensionParameters, options)
	if err != nil {
		return nil, err
	}
	result := &VirtualMachineScaleSetExtensionPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualMachineScaleSetExtensionsClient.CreateOrUpdate", "", resp, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &virtualMachineScaleSetExtensionPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*VirtualMachineScaleSetExtensionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *VirtualMachineScaleSetExtensionsClient) ResumeCreateOrUpdate(token string) (VirtualMachineScaleSetExtensionPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualMachineScaleSetExtensionsClient.CreateOrUpdate", token, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &virtualMachineScaleSetExtensionPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - The operation to create or update an extension.
func (client *VirtualMachineScaleSetExtensionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, extensionParameters VirtualMachineScaleSetExtension, options *VirtualMachineScaleSetExtensionsCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.CreateOrUpdateCreateRequest(ctx, resourceGroupName, vmScaleSetName, vmssExtensionName, extensionParameters, options)
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
func (client *VirtualMachineScaleSetExtensionsClient) CreateOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, extensionParameters VirtualMachineScaleSetExtension, options *VirtualMachineScaleSetExtensionsCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/extensions/{vmssExtensionName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	urlPath = strings.ReplaceAll(urlPath, "{vmssExtensionName}", url.PathEscape(vmssExtensionName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(extensionParameters)
}

// CreateOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *VirtualMachineScaleSetExtensionsClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*VirtualMachineScaleSetExtensionResponse, error) {
	result := VirtualMachineScaleSetExtensionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualMachineScaleSetExtension)
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *VirtualMachineScaleSetExtensionsClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

func (client *VirtualMachineScaleSetExtensionsClient) BeginDelete(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, options *VirtualMachineScaleSetExtensionsDeleteOptions) (*HTTPPollerResponse, error) {
	resp, err := client.Delete(ctx, resourceGroupName, vmScaleSetName, vmssExtensionName, options)
	if err != nil {
		return nil, err
	}
	result := &HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualMachineScaleSetExtensionsClient.Delete", "", resp, client.DeleteHandleError)
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

func (client *VirtualMachineScaleSetExtensionsClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualMachineScaleSetExtensionsClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - The operation to delete the extension.
func (client *VirtualMachineScaleSetExtensionsClient) Delete(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, options *VirtualMachineScaleSetExtensionsDeleteOptions) (*azcore.Response, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, vmScaleSetName, vmssExtensionName, options)
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
func (client *VirtualMachineScaleSetExtensionsClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, options *VirtualMachineScaleSetExtensionsDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/extensions/{vmssExtensionName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	urlPath = strings.ReplaceAll(urlPath, "{vmssExtensionName}", url.PathEscape(vmssExtensionName))
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
func (client *VirtualMachineScaleSetExtensionsClient) DeleteHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// Get - The operation to get the extension.
func (client *VirtualMachineScaleSetExtensionsClient) Get(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, options *VirtualMachineScaleSetExtensionsGetOptions) (*VirtualMachineScaleSetExtensionResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, vmScaleSetName, vmssExtensionName, options)
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
func (client *VirtualMachineScaleSetExtensionsClient) GetCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, options *VirtualMachineScaleSetExtensionsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/extensions/{vmssExtensionName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	urlPath = strings.ReplaceAll(urlPath, "{vmssExtensionName}", url.PathEscape(vmssExtensionName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	if options != nil && options.Expand != nil {
		query.Set("$expand", *options.Expand)
	}
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetHandleResponse handles the Get response.
func (client *VirtualMachineScaleSetExtensionsClient) GetHandleResponse(resp *azcore.Response) (*VirtualMachineScaleSetExtensionResponse, error) {
	result := VirtualMachineScaleSetExtensionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualMachineScaleSetExtension)
}

// GetHandleError handles the Get error response.
func (client *VirtualMachineScaleSetExtensionsClient) GetHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// List - Gets a list of all extensions in a VM scale set.
func (client *VirtualMachineScaleSetExtensionsClient) List(resourceGroupName string, vmScaleSetName string, options *VirtualMachineScaleSetExtensionsListOptions) VirtualMachineScaleSetExtensionListResultPager {
	return &virtualMachineScaleSetExtensionListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, resourceGroupName, vmScaleSetName, options)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *VirtualMachineScaleSetExtensionListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.VirtualMachineScaleSetExtensionListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListCreateRequest creates the List request.
func (client *VirtualMachineScaleSetExtensionsClient) ListCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, options *VirtualMachineScaleSetExtensionsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/extensions"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
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
func (client *VirtualMachineScaleSetExtensionsClient) ListHandleResponse(resp *azcore.Response) (*VirtualMachineScaleSetExtensionListResultResponse, error) {
	result := VirtualMachineScaleSetExtensionListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualMachineScaleSetExtensionListResult)
}

// ListHandleError handles the List error response.
func (client *VirtualMachineScaleSetExtensionsClient) ListHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

func (client *VirtualMachineScaleSetExtensionsClient) BeginUpdate(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, extensionParameters VirtualMachineScaleSetExtensionUpdate, options *VirtualMachineScaleSetExtensionsUpdateOptions) (*VirtualMachineScaleSetExtensionPollerResponse, error) {
	resp, err := client.Update(ctx, resourceGroupName, vmScaleSetName, vmssExtensionName, extensionParameters, options)
	if err != nil {
		return nil, err
	}
	result := &VirtualMachineScaleSetExtensionPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualMachineScaleSetExtensionsClient.Update", "", resp, client.UpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &virtualMachineScaleSetExtensionPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*VirtualMachineScaleSetExtensionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *VirtualMachineScaleSetExtensionsClient) ResumeUpdate(token string) (VirtualMachineScaleSetExtensionPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualMachineScaleSetExtensionsClient.Update", token, client.UpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &virtualMachineScaleSetExtensionPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Update - The operation to update an extension.
func (client *VirtualMachineScaleSetExtensionsClient) Update(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, extensionParameters VirtualMachineScaleSetExtensionUpdate, options *VirtualMachineScaleSetExtensionsUpdateOptions) (*azcore.Response, error) {
	req, err := client.UpdateCreateRequest(ctx, resourceGroupName, vmScaleSetName, vmssExtensionName, extensionParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.UpdateHandleError(resp)
	}
	return resp, nil
}

// UpdateCreateRequest creates the Update request.
func (client *VirtualMachineScaleSetExtensionsClient) UpdateCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, extensionParameters VirtualMachineScaleSetExtensionUpdate, options *VirtualMachineScaleSetExtensionsUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/extensions/{vmssExtensionName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	urlPath = strings.ReplaceAll(urlPath, "{vmssExtensionName}", url.PathEscape(vmssExtensionName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(extensionParameters)
}

// UpdateHandleResponse handles the Update response.
func (client *VirtualMachineScaleSetExtensionsClient) UpdateHandleResponse(resp *azcore.Response) (*VirtualMachineScaleSetExtensionResponse, error) {
	result := VirtualMachineScaleSetExtensionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualMachineScaleSetExtension)
}

// UpdateHandleError handles the Update error response.
func (client *VirtualMachineScaleSetExtensionsClient) UpdateHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
