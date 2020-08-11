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

// VirtualNetworksOperations contains the methods for the VirtualNetworks group.
type VirtualNetworksOperations interface {
	// CheckIPAddressAvailability - Checks whether a private IP address is available for use.
	CheckIPAddressAvailability(ctx context.Context, resourceGroupName string, virtualNetworkName string, ipAddress string) (*IPAddressAvailabilityResultResponse, error)
	// BeginCreateOrUpdate - Creates or updates a virtual network in the specified resource group.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkName string, parameters VirtualNetwork) (*VirtualNetworkPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (VirtualNetworkPoller, error)
	// BeginDelete - Deletes the specified virtual network.
	BeginDelete(ctx context.Context, resourceGroupName string, virtualNetworkName string) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets the specified virtual network by resource group.
	Get(ctx context.Context, resourceGroupName string, virtualNetworkName string, virtualNetworksGetOptions *VirtualNetworksGetOptions) (*VirtualNetworkResponse, error)
	// List - Gets all virtual networks in a resource group.
	List(resourceGroupName string) (VirtualNetworkListResultPager, error)
	// ListAll - Gets all virtual networks in a subscription.
	ListAll() (VirtualNetworkListResultPager, error)
	// ListUsage - Lists usage stats.
	ListUsage(resourceGroupName string, virtualNetworkName string) (VirtualNetworkListUsageResultPager, error)
	// UpdateTags - Updates a virtual network tags.
	UpdateTags(ctx context.Context, resourceGroupName string, virtualNetworkName string, parameters TagsObject) (*VirtualNetworkResponse, error)
}

// virtualNetworksOperations implements the VirtualNetworksOperations interface.
type virtualNetworksOperations struct {
	*Client
	subscriptionID string
}

// CheckIPAddressAvailability - Checks whether a private IP address is available for use.
func (client *virtualNetworksOperations) CheckIPAddressAvailability(ctx context.Context, resourceGroupName string, virtualNetworkName string, ipAddress string) (*IPAddressAvailabilityResultResponse, error) {
	req, err := client.checkIPAddressAvailabilityCreateRequest(resourceGroupName, virtualNetworkName, ipAddress)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.checkIPAddressAvailabilityHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// checkIPAddressAvailabilityCreateRequest creates the CheckIPAddressAvailability request.
func (client *virtualNetworksOperations) checkIPAddressAvailabilityCreateRequest(resourceGroupName string, virtualNetworkName string, ipAddress string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/CheckIPAddressAvailability"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("ipAddress", ipAddress)
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// checkIPAddressAvailabilityHandleResponse handles the CheckIPAddressAvailability response.
func (client *virtualNetworksOperations) checkIPAddressAvailabilityHandleResponse(resp *azcore.Response) (*IPAddressAvailabilityResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.checkIPAddressAvailabilityHandleError(resp)
	}
	result := IPAddressAvailabilityResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.IPAddressAvailabilityResult)
}

// checkIPAddressAvailabilityHandleError handles the CheckIPAddressAvailability error response.
func (client *virtualNetworksOperations) checkIPAddressAvailabilityHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// CreateOrUpdate - Creates or updates a virtual network in the specified resource group.
func (client *virtualNetworksOperations) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkName string, parameters VirtualNetwork) (*VirtualNetworkPollerResponse, error) {
	req, err := client.createOrUpdateCreateRequest(resourceGroupName, virtualNetworkName, parameters)
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
	pt, err := armcore.NewPoller("virtualNetworksOperations.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &virtualNetworkPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*VirtualNetworkResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *virtualNetworksOperations) ResumeCreateOrUpdate(token string) (VirtualNetworkPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("virtualNetworksOperations.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &virtualNetworkPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *virtualNetworksOperations) createOrUpdateCreateRequest(resourceGroupName string, virtualNetworkName string, parameters VirtualNetwork) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
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
func (client *virtualNetworksOperations) createOrUpdateHandleResponse(resp *azcore.Response) (*VirtualNetworkPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated, http.StatusNoContent) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return &VirtualNetworkPollerResponse{RawResponse: resp.Response}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *virtualNetworksOperations) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Delete - Deletes the specified virtual network.
func (client *virtualNetworksOperations) BeginDelete(ctx context.Context, resourceGroupName string, virtualNetworkName string) (*HTTPPollerResponse, error) {
	req, err := client.deleteCreateRequest(resourceGroupName, virtualNetworkName)
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
	pt, err := armcore.NewPoller("virtualNetworksOperations.Delete", "location", resp, client.deleteHandleError)
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

func (client *virtualNetworksOperations) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("virtualNetworksOperations.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *virtualNetworksOperations) deleteCreateRequest(resourceGroupName string, virtualNetworkName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
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
func (client *virtualNetworksOperations) deleteHandleResponse(resp *azcore.Response) (*HTTPPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return &HTTPPollerResponse{RawResponse: resp.Response}, nil
}

// deleteHandleError handles the Delete error response.
func (client *virtualNetworksOperations) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Get - Gets the specified virtual network by resource group.
func (client *virtualNetworksOperations) Get(ctx context.Context, resourceGroupName string, virtualNetworkName string, virtualNetworksGetOptions *VirtualNetworksGetOptions) (*VirtualNetworkResponse, error) {
	req, err := client.getCreateRequest(resourceGroupName, virtualNetworkName, virtualNetworksGetOptions)
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
func (client *virtualNetworksOperations) getCreateRequest(resourceGroupName string, virtualNetworkName string, virtualNetworksGetOptions *VirtualNetworksGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	if virtualNetworksGetOptions != nil && virtualNetworksGetOptions.Expand != nil {
		query.Set("$expand", *virtualNetworksGetOptions.Expand)
	}
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *virtualNetworksOperations) getHandleResponse(resp *azcore.Response) (*VirtualNetworkResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getHandleError(resp)
	}
	result := VirtualNetworkResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualNetwork)
}

// getHandleError handles the Get error response.
func (client *virtualNetworksOperations) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// List - Gets all virtual networks in a resource group.
func (client *virtualNetworksOperations) List(resourceGroupName string) (VirtualNetworkListResultPager, error) {
	req, err := client.listCreateRequest(resourceGroupName)
	if err != nil {
		return nil, err
	}
	return &virtualNetworkListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listHandleResponse,
		advancer: func(resp *VirtualNetworkListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.VirtualNetworkListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.VirtualNetworkListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listCreateRequest creates the List request.
func (client *virtualNetworksOperations) listCreateRequest(resourceGroupName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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
func (client *virtualNetworksOperations) listHandleResponse(resp *azcore.Response) (*VirtualNetworkListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listHandleError(resp)
	}
	result := VirtualNetworkListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualNetworkListResult)
}

// listHandleError handles the List error response.
func (client *virtualNetworksOperations) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// ListAll - Gets all virtual networks in a subscription.
func (client *virtualNetworksOperations) ListAll() (VirtualNetworkListResultPager, error) {
	req, err := client.listAllCreateRequest()
	if err != nil {
		return nil, err
	}
	return &virtualNetworkListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listAllHandleResponse,
		advancer: func(resp *VirtualNetworkListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.VirtualNetworkListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.VirtualNetworkListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listAllCreateRequest creates the ListAll request.
func (client *virtualNetworksOperations) listAllCreateRequest() (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/virtualNetworks"
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

// listAllHandleResponse handles the ListAll response.
func (client *virtualNetworksOperations) listAllHandleResponse(resp *azcore.Response) (*VirtualNetworkListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listAllHandleError(resp)
	}
	result := VirtualNetworkListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualNetworkListResult)
}

// listAllHandleError handles the ListAll error response.
func (client *virtualNetworksOperations) listAllHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// ListUsage - Lists usage stats.
func (client *virtualNetworksOperations) ListUsage(resourceGroupName string, virtualNetworkName string) (VirtualNetworkListUsageResultPager, error) {
	req, err := client.listUsageCreateRequest(resourceGroupName, virtualNetworkName)
	if err != nil {
		return nil, err
	}
	return &virtualNetworkListUsageResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listUsageHandleResponse,
		advancer: func(resp *VirtualNetworkListUsageResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.VirtualNetworkListUsageResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.VirtualNetworkListUsageResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listUsageCreateRequest creates the ListUsage request.
func (client *virtualNetworksOperations) listUsageCreateRequest(resourceGroupName string, virtualNetworkName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/usages"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
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

// listUsageHandleResponse handles the ListUsage response.
func (client *virtualNetworksOperations) listUsageHandleResponse(resp *azcore.Response) (*VirtualNetworkListUsageResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listUsageHandleError(resp)
	}
	result := VirtualNetworkListUsageResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualNetworkListUsageResult)
}

// listUsageHandleError handles the ListUsage error response.
func (client *virtualNetworksOperations) listUsageHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// UpdateTags - Updates a virtual network tags.
func (client *virtualNetworksOperations) UpdateTags(ctx context.Context, resourceGroupName string, virtualNetworkName string, parameters TagsObject) (*VirtualNetworkResponse, error) {
	req, err := client.updateTagsCreateRequest(resourceGroupName, virtualNetworkName, parameters)
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
func (client *virtualNetworksOperations) updateTagsCreateRequest(resourceGroupName string, virtualNetworkName string, parameters TagsObject) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
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
func (client *virtualNetworksOperations) updateTagsHandleResponse(resp *azcore.Response) (*VirtualNetworkResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.updateTagsHandleError(resp)
	}
	result := VirtualNetworkResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualNetwork)
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *virtualNetworksOperations) updateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
