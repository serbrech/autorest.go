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

// VirtualNetworkGatewayConnectionsOperations contains the methods for the VirtualNetworkGatewayConnections group.
type VirtualNetworkGatewayConnectionsOperations interface {
	// BeginCreateOrUpdate - Creates or updates a virtual network gateway connection in the specified resource group.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters VirtualNetworkGatewayConnection) (*VirtualNetworkGatewayConnectionPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (VirtualNetworkGatewayConnectionPoller, error)
	// BeginDelete - Deletes the specified virtual network Gateway connection.
	BeginDelete(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets the specified virtual network gateway connection by resource group.
	Get(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string) (*VirtualNetworkGatewayConnectionResponse, error)
	// GetSharedKey - The Get VirtualNetworkGatewayConnectionSharedKey operation retrieves information about the specified virtual network gateway connection shared key through Network resource provider.
	GetSharedKey(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string) (*ConnectionSharedKeyResponse, error)
	// List - The List VirtualNetworkGatewayConnections operation retrieves all the virtual network gateways connections created.
	List(resourceGroupName string) (VirtualNetworkGatewayConnectionListResultPager, error)
	// BeginResetSharedKey - The VirtualNetworkGatewayConnectionResetSharedKey operation resets the virtual network gateway connection shared key for passed virtual network gateway connection in the specified resource group through Network resource provider.
	BeginResetSharedKey(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters ConnectionResetSharedKey) (*ConnectionResetSharedKeyPollerResponse, error)
	// ResumeResetSharedKey - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeResetSharedKey(token string) (ConnectionResetSharedKeyPoller, error)
	// BeginSetSharedKey - The Put VirtualNetworkGatewayConnectionSharedKey operation sets the virtual network gateway connection shared key for passed virtual network gateway connection in the specified resource group through Network resource provider.
	BeginSetSharedKey(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters ConnectionSharedKey) (*ConnectionSharedKeyPollerResponse, error)
	// ResumeSetSharedKey - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeSetSharedKey(token string) (ConnectionSharedKeyPoller, error)
	// BeginStartPacketCapture - Starts packet capture on virtual network gateway connection in the specified resource group.
	BeginStartPacketCapture(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, virtualNetworkGatewayConnectionsStartPacketCaptureOptions *VirtualNetworkGatewayConnectionsStartPacketCaptureOptions) (*StringPollerResponse, error)
	// ResumeStartPacketCapture - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeStartPacketCapture(token string) (StringPoller, error)
	// BeginStopPacketCapture - Stops packet capture on virtual network gateway connection in the specified resource group.
	BeginStopPacketCapture(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters VpnPacketCaptureStopParameters) (*StringPollerResponse, error)
	// ResumeStopPacketCapture - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeStopPacketCapture(token string) (StringPoller, error)
	// BeginUpdateTags - Updates a virtual network gateway connection tags.
	BeginUpdateTags(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters TagsObject) (*VirtualNetworkGatewayConnectionPollerResponse, error)
	// ResumeUpdateTags - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeUpdateTags(token string) (VirtualNetworkGatewayConnectionPoller, error)
}

// virtualNetworkGatewayConnectionsOperations implements the VirtualNetworkGatewayConnectionsOperations interface.
type virtualNetworkGatewayConnectionsOperations struct {
	*Client
	subscriptionID string
}

// CreateOrUpdate - Creates or updates a virtual network gateway connection in the specified resource group.
func (client *virtualNetworkGatewayConnectionsOperations) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters VirtualNetworkGatewayConnection) (*VirtualNetworkGatewayConnectionPollerResponse, error) {
	req, err := client.createOrUpdateCreateRequest(resourceGroupName, virtualNetworkGatewayConnectionName, parameters)
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
	pt, err := armcore.NewPoller("virtualNetworkGatewayConnectionsOperations.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &virtualNetworkGatewayConnectionPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*VirtualNetworkGatewayConnectionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *virtualNetworkGatewayConnectionsOperations) ResumeCreateOrUpdate(token string) (VirtualNetworkGatewayConnectionPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("virtualNetworkGatewayConnectionsOperations.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &virtualNetworkGatewayConnectionPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *virtualNetworkGatewayConnectionsOperations) createOrUpdateCreateRequest(resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters VirtualNetworkGatewayConnection) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkGatewayConnectionName}", url.PathEscape(virtualNetworkGatewayConnectionName))
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
func (client *virtualNetworkGatewayConnectionsOperations) createOrUpdateHandleResponse(resp *azcore.Response) (*VirtualNetworkGatewayConnectionPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated, http.StatusNoContent) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return &VirtualNetworkGatewayConnectionPollerResponse{RawResponse: resp.Response}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *virtualNetworkGatewayConnectionsOperations) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Delete - Deletes the specified virtual network Gateway connection.
func (client *virtualNetworkGatewayConnectionsOperations) BeginDelete(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string) (*HTTPPollerResponse, error) {
	req, err := client.deleteCreateRequest(resourceGroupName, virtualNetworkGatewayConnectionName)
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
	pt, err := armcore.NewPoller("virtualNetworkGatewayConnectionsOperations.Delete", "location", resp, client.deleteHandleError)
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

func (client *virtualNetworkGatewayConnectionsOperations) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("virtualNetworkGatewayConnectionsOperations.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *virtualNetworkGatewayConnectionsOperations) deleteCreateRequest(resourceGroupName string, virtualNetworkGatewayConnectionName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkGatewayConnectionName}", url.PathEscape(virtualNetworkGatewayConnectionName))
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
func (client *virtualNetworkGatewayConnectionsOperations) deleteHandleResponse(resp *azcore.Response) (*HTTPPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return &HTTPPollerResponse{RawResponse: resp.Response}, nil
}

// deleteHandleError handles the Delete error response.
func (client *virtualNetworkGatewayConnectionsOperations) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Get - Gets the specified virtual network gateway connection by resource group.
func (client *virtualNetworkGatewayConnectionsOperations) Get(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string) (*VirtualNetworkGatewayConnectionResponse, error) {
	req, err := client.getCreateRequest(resourceGroupName, virtualNetworkGatewayConnectionName)
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
func (client *virtualNetworkGatewayConnectionsOperations) getCreateRequest(resourceGroupName string, virtualNetworkGatewayConnectionName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkGatewayConnectionName}", url.PathEscape(virtualNetworkGatewayConnectionName))
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

// getHandleResponse handles the Get response.
func (client *virtualNetworkGatewayConnectionsOperations) getHandleResponse(resp *azcore.Response) (*VirtualNetworkGatewayConnectionResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getHandleError(resp)
	}
	result := VirtualNetworkGatewayConnectionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualNetworkGatewayConnection)
}

// getHandleError handles the Get error response.
func (client *virtualNetworkGatewayConnectionsOperations) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetSharedKey - The Get VirtualNetworkGatewayConnectionSharedKey operation retrieves information about the specified virtual network gateway connection shared key through Network resource provider.
func (client *virtualNetworkGatewayConnectionsOperations) GetSharedKey(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string) (*ConnectionSharedKeyResponse, error) {
	req, err := client.getSharedKeyCreateRequest(resourceGroupName, virtualNetworkGatewayConnectionName)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getSharedKeyHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getSharedKeyCreateRequest creates the GetSharedKey request.
func (client *virtualNetworkGatewayConnectionsOperations) getSharedKeyCreateRequest(resourceGroupName string, virtualNetworkGatewayConnectionName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}/sharedkey"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkGatewayConnectionName}", url.PathEscape(virtualNetworkGatewayConnectionName))
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

// getSharedKeyHandleResponse handles the GetSharedKey response.
func (client *virtualNetworkGatewayConnectionsOperations) getSharedKeyHandleResponse(resp *azcore.Response) (*ConnectionSharedKeyResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getSharedKeyHandleError(resp)
	}
	result := ConnectionSharedKeyResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ConnectionSharedKey)
}

// getSharedKeyHandleError handles the GetSharedKey error response.
func (client *virtualNetworkGatewayConnectionsOperations) getSharedKeyHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// List - The List VirtualNetworkGatewayConnections operation retrieves all the virtual network gateways connections created.
func (client *virtualNetworkGatewayConnectionsOperations) List(resourceGroupName string) (VirtualNetworkGatewayConnectionListResultPager, error) {
	req, err := client.listCreateRequest(resourceGroupName)
	if err != nil {
		return nil, err
	}
	return &virtualNetworkGatewayConnectionListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listHandleResponse,
		advancer: func(resp *VirtualNetworkGatewayConnectionListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.VirtualNetworkGatewayConnectionListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.VirtualNetworkGatewayConnectionListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listCreateRequest creates the List request.
func (client *virtualNetworkGatewayConnectionsOperations) listCreateRequest(resourceGroupName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections"
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
func (client *virtualNetworkGatewayConnectionsOperations) listHandleResponse(resp *azcore.Response) (*VirtualNetworkGatewayConnectionListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listHandleError(resp)
	}
	result := VirtualNetworkGatewayConnectionListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualNetworkGatewayConnectionListResult)
}

// listHandleError handles the List error response.
func (client *virtualNetworkGatewayConnectionsOperations) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// ResetSharedKey - The VirtualNetworkGatewayConnectionResetSharedKey operation resets the virtual network gateway connection shared key for passed virtual network gateway connection in the specified resource group through Network resource provider.
func (client *virtualNetworkGatewayConnectionsOperations) BeginResetSharedKey(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters ConnectionResetSharedKey) (*ConnectionResetSharedKeyPollerResponse, error) {
	req, err := client.resetSharedKeyCreateRequest(resourceGroupName, virtualNetworkGatewayConnectionName, parameters)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.resetSharedKeyHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := armcore.NewPoller("virtualNetworkGatewayConnectionsOperations.ResetSharedKey", "location", resp, client.resetSharedKeyHandleError)
	if err != nil {
		return nil, err
	}
	poller := &connectionResetSharedKeyPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ConnectionResetSharedKeyResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *virtualNetworkGatewayConnectionsOperations) ResumeResetSharedKey(token string) (ConnectionResetSharedKeyPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("virtualNetworkGatewayConnectionsOperations.ResetSharedKey", token, client.resetSharedKeyHandleError)
	if err != nil {
		return nil, err
	}
	return &connectionResetSharedKeyPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// resetSharedKeyCreateRequest creates the ResetSharedKey request.
func (client *virtualNetworkGatewayConnectionsOperations) resetSharedKeyCreateRequest(resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters ConnectionResetSharedKey) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}/sharedkey/reset"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkGatewayConnectionName}", url.PathEscape(virtualNetworkGatewayConnectionName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPost, *u)
	return req, req.MarshalAsJSON(parameters)
}

// resetSharedKeyHandleResponse handles the ResetSharedKey response.
func (client *virtualNetworkGatewayConnectionsOperations) resetSharedKeyHandleResponse(resp *azcore.Response) (*ConnectionResetSharedKeyPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.resetSharedKeyHandleError(resp)
	}
	return &ConnectionResetSharedKeyPollerResponse{RawResponse: resp.Response}, nil
}

// resetSharedKeyHandleError handles the ResetSharedKey error response.
func (client *virtualNetworkGatewayConnectionsOperations) resetSharedKeyHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// SetSharedKey - The Put VirtualNetworkGatewayConnectionSharedKey operation sets the virtual network gateway connection shared key for passed virtual network gateway connection in the specified resource group through Network resource provider.
func (client *virtualNetworkGatewayConnectionsOperations) BeginSetSharedKey(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters ConnectionSharedKey) (*ConnectionSharedKeyPollerResponse, error) {
	req, err := client.setSharedKeyCreateRequest(resourceGroupName, virtualNetworkGatewayConnectionName, parameters)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.setSharedKeyHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := armcore.NewPoller("virtualNetworkGatewayConnectionsOperations.SetSharedKey", "azure-async-operation", resp, client.setSharedKeyHandleError)
	if err != nil {
		return nil, err
	}
	poller := &connectionSharedKeyPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ConnectionSharedKeyResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *virtualNetworkGatewayConnectionsOperations) ResumeSetSharedKey(token string) (ConnectionSharedKeyPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("virtualNetworkGatewayConnectionsOperations.SetSharedKey", token, client.setSharedKeyHandleError)
	if err != nil {
		return nil, err
	}
	return &connectionSharedKeyPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// setSharedKeyCreateRequest creates the SetSharedKey request.
func (client *virtualNetworkGatewayConnectionsOperations) setSharedKeyCreateRequest(resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters ConnectionSharedKey) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}/sharedkey"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkGatewayConnectionName}", url.PathEscape(virtualNetworkGatewayConnectionName))
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

// setSharedKeyHandleResponse handles the SetSharedKey response.
func (client *virtualNetworkGatewayConnectionsOperations) setSharedKeyHandleResponse(resp *azcore.Response) (*ConnectionSharedKeyPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated, http.StatusNoContent) {
		return nil, client.setSharedKeyHandleError(resp)
	}
	return &ConnectionSharedKeyPollerResponse{RawResponse: resp.Response}, nil
}

// setSharedKeyHandleError handles the SetSharedKey error response.
func (client *virtualNetworkGatewayConnectionsOperations) setSharedKeyHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// StartPacketCapture - Starts packet capture on virtual network gateway connection in the specified resource group.
func (client *virtualNetworkGatewayConnectionsOperations) BeginStartPacketCapture(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, virtualNetworkGatewayConnectionsStartPacketCaptureOptions *VirtualNetworkGatewayConnectionsStartPacketCaptureOptions) (*StringPollerResponse, error) {
	req, err := client.startPacketCaptureCreateRequest(resourceGroupName, virtualNetworkGatewayConnectionName, virtualNetworkGatewayConnectionsStartPacketCaptureOptions)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.startPacketCaptureHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := armcore.NewPoller("virtualNetworkGatewayConnectionsOperations.StartPacketCapture", "location", resp, client.startPacketCaptureHandleError)
	if err != nil {
		return nil, err
	}
	poller := &stringPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*StringResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *virtualNetworkGatewayConnectionsOperations) ResumeStartPacketCapture(token string) (StringPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("virtualNetworkGatewayConnectionsOperations.StartPacketCapture", token, client.startPacketCaptureHandleError)
	if err != nil {
		return nil, err
	}
	return &stringPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// startPacketCaptureCreateRequest creates the StartPacketCapture request.
func (client *virtualNetworkGatewayConnectionsOperations) startPacketCaptureCreateRequest(resourceGroupName string, virtualNetworkGatewayConnectionName string, virtualNetworkGatewayConnectionsStartPacketCaptureOptions *VirtualNetworkGatewayConnectionsStartPacketCaptureOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}/startPacketCapture"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkGatewayConnectionName}", url.PathEscape(virtualNetworkGatewayConnectionName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPost, *u)
	if virtualNetworkGatewayConnectionsStartPacketCaptureOptions != nil {
		return req, req.MarshalAsJSON(virtualNetworkGatewayConnectionsStartPacketCaptureOptions.Parameters)
	}
	return req, nil
}

// startPacketCaptureHandleResponse handles the StartPacketCapture response.
func (client *virtualNetworkGatewayConnectionsOperations) startPacketCaptureHandleResponse(resp *azcore.Response) (*StringPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.startPacketCaptureHandleError(resp)
	}
	return &StringPollerResponse{RawResponse: resp.Response}, nil
}

// startPacketCaptureHandleError handles the StartPacketCapture error response.
func (client *virtualNetworkGatewayConnectionsOperations) startPacketCaptureHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// StopPacketCapture - Stops packet capture on virtual network gateway connection in the specified resource group.
func (client *virtualNetworkGatewayConnectionsOperations) BeginStopPacketCapture(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters VpnPacketCaptureStopParameters) (*StringPollerResponse, error) {
	req, err := client.stopPacketCaptureCreateRequest(resourceGroupName, virtualNetworkGatewayConnectionName, parameters)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.stopPacketCaptureHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := armcore.NewPoller("virtualNetworkGatewayConnectionsOperations.StopPacketCapture", "location", resp, client.stopPacketCaptureHandleError)
	if err != nil {
		return nil, err
	}
	poller := &stringPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*StringResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *virtualNetworkGatewayConnectionsOperations) ResumeStopPacketCapture(token string) (StringPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("virtualNetworkGatewayConnectionsOperations.StopPacketCapture", token, client.stopPacketCaptureHandleError)
	if err != nil {
		return nil, err
	}
	return &stringPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// stopPacketCaptureCreateRequest creates the StopPacketCapture request.
func (client *virtualNetworkGatewayConnectionsOperations) stopPacketCaptureCreateRequest(resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters VpnPacketCaptureStopParameters) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}/stopPacketCapture"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkGatewayConnectionName}", url.PathEscape(virtualNetworkGatewayConnectionName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPost, *u)
	return req, req.MarshalAsJSON(parameters)
}

// stopPacketCaptureHandleResponse handles the StopPacketCapture response.
func (client *virtualNetworkGatewayConnectionsOperations) stopPacketCaptureHandleResponse(resp *azcore.Response) (*StringPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.stopPacketCaptureHandleError(resp)
	}
	return &StringPollerResponse{RawResponse: resp.Response}, nil
}

// stopPacketCaptureHandleError handles the StopPacketCapture error response.
func (client *virtualNetworkGatewayConnectionsOperations) stopPacketCaptureHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// UpdateTags - Updates a virtual network gateway connection tags.
func (client *virtualNetworkGatewayConnectionsOperations) BeginUpdateTags(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters TagsObject) (*VirtualNetworkGatewayConnectionPollerResponse, error) {
	req, err := client.updateTagsCreateRequest(resourceGroupName, virtualNetworkGatewayConnectionName, parameters)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.updateTagsHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := armcore.NewPoller("virtualNetworkGatewayConnectionsOperations.UpdateTags", "azure-async-operation", resp, client.updateTagsHandleError)
	if err != nil {
		return nil, err
	}
	poller := &virtualNetworkGatewayConnectionPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*VirtualNetworkGatewayConnectionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *virtualNetworkGatewayConnectionsOperations) ResumeUpdateTags(token string) (VirtualNetworkGatewayConnectionPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("virtualNetworkGatewayConnectionsOperations.UpdateTags", token, client.updateTagsHandleError)
	if err != nil {
		return nil, err
	}
	return &virtualNetworkGatewayConnectionPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *virtualNetworkGatewayConnectionsOperations) updateTagsCreateRequest(resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters TagsObject) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkGatewayConnectionName}", url.PathEscape(virtualNetworkGatewayConnectionName))
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
func (client *virtualNetworkGatewayConnectionsOperations) updateTagsHandleResponse(resp *azcore.Response) (*VirtualNetworkGatewayConnectionPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.updateTagsHandleError(resp)
	}
	return &VirtualNetworkGatewayConnectionPollerResponse{RawResponse: resp.Response}, nil
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *virtualNetworkGatewayConnectionsOperations) updateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
