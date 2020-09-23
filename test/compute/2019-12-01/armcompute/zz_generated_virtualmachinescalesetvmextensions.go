// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// VirtualMachineScaleSetVMExtensionsOperations contains the methods for the VirtualMachineScaleSetVMExtensions group.
type VirtualMachineScaleSetVMExtensionsOperations interface {
	// BeginCreateOrUpdate - The operation to create or update the VMSS VM extension.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceId string, vmExtensionName string, extensionParameters VirtualMachineExtension) (*VirtualMachineExtensionPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (VirtualMachineExtensionPoller, error)
	// BeginDelete - The operation to delete the VMSS VM extension.
	BeginDelete(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceId string, vmExtensionName string) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - The operation to get the VMSS VM extension.
	Get(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceId string, vmExtensionName string, virtualMachineScaleSetVMExtensionsGetOptions *VirtualMachineScaleSetVMExtensionsGetOptions) (*VirtualMachineExtensionResponse, error)
	// List - The operation to get all extensions of an instance in Virtual Machine Scaleset.
	List(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceId string, virtualMachineScaleSetVMExtensionsListOptions *VirtualMachineScaleSetVMExtensionsListOptions) (*VirtualMachineExtensionsListResultResponse, error)
	// BeginUpdate - The operation to update the VMSS VM extension.
	BeginUpdate(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceId string, vmExtensionName string, extensionParameters VirtualMachineExtensionUpdate) (*VirtualMachineExtensionPollerResponse, error)
	// ResumeUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeUpdate(token string) (VirtualMachineExtensionPoller, error)
}

// VirtualMachineScaleSetVMExtensionsClient implements the VirtualMachineScaleSetVMExtensionsOperations interface.
// Don't use this type directly, use NewVirtualMachineScaleSetVMExtensionsClient() instead.
type VirtualMachineScaleSetVMExtensionsClient struct {
	*Client
	subscriptionID string
}

// NewVirtualMachineScaleSetVMExtensionsClient creates a new instance of VirtualMachineScaleSetVMExtensionsClient with the specified values.
func NewVirtualMachineScaleSetVMExtensionsClient(c *Client, subscriptionID string) VirtualMachineScaleSetVMExtensionsOperations {
	return &VirtualMachineScaleSetVMExtensionsClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *VirtualMachineScaleSetVMExtensionsClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

func (client *VirtualMachineScaleSetVMExtensionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceId string, vmExtensionName string, extensionParameters VirtualMachineExtension) (*VirtualMachineExtensionPollerResponse, error) {
	resp, err := client.CreateOrUpdate(ctx, resourceGroupName, vmScaleSetName, instanceId, vmExtensionName, extensionParameters)
	if err != nil {
		return nil, err
	}
	result := &VirtualMachineExtensionPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualMachineScaleSetVMExtensionsClient.CreateOrUpdate", "", resp, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &virtualMachineExtensionPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*VirtualMachineExtensionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *VirtualMachineScaleSetVMExtensionsClient) ResumeCreateOrUpdate(token string) (VirtualMachineExtensionPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualMachineScaleSetVMExtensionsClient.CreateOrUpdate", token, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &virtualMachineExtensionPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// CreateOrUpdate - The operation to create or update the VMSS VM extension.
func (client *VirtualMachineScaleSetVMExtensionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceId string, vmExtensionName string, extensionParameters VirtualMachineExtension) (*azcore.Response, error) {
	req, err := client.CreateOrUpdateCreateRequest(ctx, resourceGroupName, vmScaleSetName, instanceId, vmExtensionName, extensionParameters)
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
func (client *VirtualMachineScaleSetVMExtensionsClient) CreateOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceId string, vmExtensionName string, extensionParameters VirtualMachineExtension) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/extensions/{vmExtensionName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceId))
	urlPath = strings.ReplaceAll(urlPath, "{vmExtensionName}", url.PathEscape(vmExtensionName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
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
func (client *VirtualMachineScaleSetVMExtensionsClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*VirtualMachineExtensionResponse, error) {
	result := VirtualMachineExtensionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualMachineExtension)
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *VirtualMachineScaleSetVMExtensionsClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *VirtualMachineScaleSetVMExtensionsClient) BeginDelete(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceId string, vmExtensionName string) (*HTTPPollerResponse, error) {
	resp, err := client.Delete(ctx, resourceGroupName, vmScaleSetName, instanceId, vmExtensionName)
	if err != nil {
		return nil, err
	}
	result := &HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualMachineScaleSetVMExtensionsClient.Delete", "", resp, client.DeleteHandleError)
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

func (client *VirtualMachineScaleSetVMExtensionsClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualMachineScaleSetVMExtensionsClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// Delete - The operation to delete the VMSS VM extension.
func (client *VirtualMachineScaleSetVMExtensionsClient) Delete(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceId string, vmExtensionName string) (*azcore.Response, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, vmScaleSetName, instanceId, vmExtensionName)
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
func (client *VirtualMachineScaleSetVMExtensionsClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceId string, vmExtensionName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/extensions/{vmExtensionName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceId))
	urlPath = strings.ReplaceAll(urlPath, "{vmExtensionName}", url.PathEscape(vmExtensionName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// DeleteHandleError handles the Delete error response.
func (client *VirtualMachineScaleSetVMExtensionsClient) DeleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - The operation to get the VMSS VM extension.
func (client *VirtualMachineScaleSetVMExtensionsClient) Get(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceId string, vmExtensionName string, virtualMachineScaleSetVMExtensionsGetOptions *VirtualMachineScaleSetVMExtensionsGetOptions) (*VirtualMachineExtensionResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, vmScaleSetName, instanceId, vmExtensionName, virtualMachineScaleSetVMExtensionsGetOptions)
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
func (client *VirtualMachineScaleSetVMExtensionsClient) GetCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceId string, vmExtensionName string, virtualMachineScaleSetVMExtensionsGetOptions *VirtualMachineScaleSetVMExtensionsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/extensions/{vmExtensionName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceId))
	urlPath = strings.ReplaceAll(urlPath, "{vmExtensionName}", url.PathEscape(vmExtensionName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	if virtualMachineScaleSetVMExtensionsGetOptions != nil && virtualMachineScaleSetVMExtensionsGetOptions.Expand != nil {
		query.Set("$expand", *virtualMachineScaleSetVMExtensionsGetOptions.Expand)
	}
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetHandleResponse handles the Get response.
func (client *VirtualMachineScaleSetVMExtensionsClient) GetHandleResponse(resp *azcore.Response) (*VirtualMachineExtensionResponse, error) {
	result := VirtualMachineExtensionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualMachineExtension)
}

// GetHandleError handles the Get error response.
func (client *VirtualMachineScaleSetVMExtensionsClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - The operation to get all extensions of an instance in Virtual Machine Scaleset.
func (client *VirtualMachineScaleSetVMExtensionsClient) List(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceId string, virtualMachineScaleSetVMExtensionsListOptions *VirtualMachineScaleSetVMExtensionsListOptions) (*VirtualMachineExtensionsListResultResponse, error) {
	req, err := client.ListCreateRequest(ctx, resourceGroupName, vmScaleSetName, instanceId, virtualMachineScaleSetVMExtensionsListOptions)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
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
func (client *VirtualMachineScaleSetVMExtensionsClient) ListCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceId string, virtualMachineScaleSetVMExtensionsListOptions *VirtualMachineScaleSetVMExtensionsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/extensions"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceId))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	if virtualMachineScaleSetVMExtensionsListOptions != nil && virtualMachineScaleSetVMExtensionsListOptions.Expand != nil {
		query.Set("$expand", *virtualMachineScaleSetVMExtensionsListOptions.Expand)
	}
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListHandleResponse handles the List response.
func (client *VirtualMachineScaleSetVMExtensionsClient) ListHandleResponse(resp *azcore.Response) (*VirtualMachineExtensionsListResultResponse, error) {
	result := VirtualMachineExtensionsListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualMachineExtensionsListResult)
}

// ListHandleError handles the List error response.
func (client *VirtualMachineScaleSetVMExtensionsClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *VirtualMachineScaleSetVMExtensionsClient) BeginUpdate(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceId string, vmExtensionName string, extensionParameters VirtualMachineExtensionUpdate) (*VirtualMachineExtensionPollerResponse, error) {
	resp, err := client.Update(ctx, resourceGroupName, vmScaleSetName, instanceId, vmExtensionName, extensionParameters)
	if err != nil {
		return nil, err
	}
	result := &VirtualMachineExtensionPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualMachineScaleSetVMExtensionsClient.Update", "", resp, client.UpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &virtualMachineExtensionPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*VirtualMachineExtensionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *VirtualMachineScaleSetVMExtensionsClient) ResumeUpdate(token string) (VirtualMachineExtensionPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualMachineScaleSetVMExtensionsClient.Update", token, client.UpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &virtualMachineExtensionPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// Update - The operation to update the VMSS VM extension.
func (client *VirtualMachineScaleSetVMExtensionsClient) Update(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceId string, vmExtensionName string, extensionParameters VirtualMachineExtensionUpdate) (*azcore.Response, error) {
	req, err := client.UpdateCreateRequest(ctx, resourceGroupName, vmScaleSetName, instanceId, vmExtensionName, extensionParameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.UpdateHandleError(resp)
	}
	return resp, nil
}

// UpdateCreateRequest creates the Update request.
func (client *VirtualMachineScaleSetVMExtensionsClient) UpdateCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceId string, vmExtensionName string, extensionParameters VirtualMachineExtensionUpdate) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/extensions/{vmExtensionName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceId))
	urlPath = strings.ReplaceAll(urlPath, "{vmExtensionName}", url.PathEscape(vmExtensionName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.u, urlPath))
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
func (client *VirtualMachineScaleSetVMExtensionsClient) UpdateHandleResponse(resp *azcore.Response) (*VirtualMachineExtensionResponse, error) {
	result := VirtualMachineExtensionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualMachineExtension)
}

// UpdateHandleError handles the Update error response.
func (client *VirtualMachineScaleSetVMExtensionsClient) UpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
