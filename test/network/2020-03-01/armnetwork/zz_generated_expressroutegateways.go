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
	"path"
	"strings"
	"time"
)

// ExpressRouteGatewaysOperations contains the methods for the ExpressRouteGateways group.
type ExpressRouteGatewaysOperations interface {
	// BeginCreateOrUpdate - Creates or updates a ExpressRoute gateway in a specified resource group.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, putExpressRouteGatewayParameters ExpressRouteGateway) (*ExpressRouteGatewayPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (ExpressRouteGatewayPoller, error)
	// BeginDelete - Deletes the specified ExpressRoute gateway in a resource group. An ExpressRoute gateway resource can only be deleted when there are no connection subresources.
	BeginDelete(ctx context.Context, resourceGroupName string, expressRouteGatewayName string) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Fetches the details of a ExpressRoute gateway in a resource group.
	Get(ctx context.Context, resourceGroupName string, expressRouteGatewayName string) (*ExpressRouteGatewayResponse, error)
	// ListByResourceGroup - Lists ExpressRoute gateways in a given resource group.
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (*ExpressRouteGatewayListResponse, error)
	// ListBySubscription - Lists ExpressRoute gateways under a given subscription.
	ListBySubscription(ctx context.Context) (*ExpressRouteGatewayListResponse, error)
}

// ExpressRouteGatewaysClient implements the ExpressRouteGatewaysOperations interface.
// Don't use this type directly, use NewExpressRouteGatewaysClient() instead.
type ExpressRouteGatewaysClient struct {
	*Client
	subscriptionID string
}

// NewExpressRouteGatewaysClient creates a new instance of ExpressRouteGatewaysClient with the specified values.
func NewExpressRouteGatewaysClient(c *Client, subscriptionID string) ExpressRouteGatewaysOperations {
	return &ExpressRouteGatewaysClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *ExpressRouteGatewaysClient) Do(ctx context.Context, req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(ctx, req)
}

// CreateOrUpdate - Creates or updates a ExpressRoute gateway in a specified resource group.
func (client *ExpressRouteGatewaysClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, putExpressRouteGatewayParameters ExpressRouteGateway) (*ExpressRouteGatewayPollerResponse, error) {
	req, err := client.CreateOrUpdateCreateRequest(resourceGroupName, expressRouteGatewayName, putExpressRouteGatewayParameters)
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
	pt, err := armcore.NewPoller("ExpressRouteGatewaysClient.CreateOrUpdate", "azure-async-operation", resp, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &expressRouteGatewayPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ExpressRouteGatewayResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *ExpressRouteGatewaysClient) ResumeCreateOrUpdate(token string) (ExpressRouteGatewayPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ExpressRouteGatewaysClient.CreateOrUpdate", token, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &expressRouteGatewayPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// CreateOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ExpressRouteGatewaysClient) CreateOrUpdateCreateRequest(resourceGroupName string, expressRouteGatewayName string, putExpressRouteGatewayParameters ExpressRouteGateway) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{expressRouteGatewayName}", url.PathEscape(expressRouteGatewayName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(putExpressRouteGatewayParameters)
}

// CreateOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ExpressRouteGatewaysClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*ExpressRouteGatewayPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated, http.StatusNoContent) {
		return nil, client.CreateOrUpdateHandleError(resp)
	}
	return &ExpressRouteGatewayPollerResponse{RawResponse: resp.Response}, nil
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ExpressRouteGatewaysClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Delete - Deletes the specified ExpressRoute gateway in a resource group. An ExpressRoute gateway resource can only be deleted when there are no connection subresources.
func (client *ExpressRouteGatewaysClient) BeginDelete(ctx context.Context, resourceGroupName string, expressRouteGatewayName string) (*HTTPPollerResponse, error) {
	req, err := client.DeleteCreateRequest(resourceGroupName, expressRouteGatewayName)
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
	pt, err := armcore.NewPoller("ExpressRouteGatewaysClient.Delete", "location", resp, client.DeleteHandleError)
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

func (client *ExpressRouteGatewaysClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ExpressRouteGatewaysClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// DeleteCreateRequest creates the Delete request.
func (client *ExpressRouteGatewaysClient) DeleteCreateRequest(resourceGroupName string, expressRouteGatewayName string) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{expressRouteGatewayName}", url.PathEscape(expressRouteGatewayName))
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
func (client *ExpressRouteGatewaysClient) DeleteHandleResponse(resp *azcore.Response) (*HTTPPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.DeleteHandleError(resp)
	}
	return &HTTPPollerResponse{RawResponse: resp.Response}, nil
}

// DeleteHandleError handles the Delete error response.
func (client *ExpressRouteGatewaysClient) DeleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Get - Fetches the details of a ExpressRoute gateway in a resource group.
func (client *ExpressRouteGatewaysClient) Get(ctx context.Context, resourceGroupName string, expressRouteGatewayName string) (*ExpressRouteGatewayResponse, error) {
	req, err := client.GetCreateRequest(resourceGroupName, expressRouteGatewayName)
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
func (client *ExpressRouteGatewaysClient) GetCreateRequest(resourceGroupName string, expressRouteGatewayName string) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{expressRouteGatewayName}", url.PathEscape(expressRouteGatewayName))
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

// GetHandleResponse handles the Get response.
func (client *ExpressRouteGatewaysClient) GetHandleResponse(resp *azcore.Response) (*ExpressRouteGatewayResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetHandleError(resp)
	}
	result := ExpressRouteGatewayResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteGateway)
}

// GetHandleError handles the Get error response.
func (client *ExpressRouteGatewaysClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// ListByResourceGroup - Lists ExpressRoute gateways in a given resource group.
func (client *ExpressRouteGatewaysClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (*ExpressRouteGatewayListResponse, error) {
	req, err := client.ListByResourceGroupCreateRequest(resourceGroupName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.ListByResourceGroupHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ExpressRouteGatewaysClient) ListByResourceGroupCreateRequest(resourceGroupName string) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways"
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
func (client *ExpressRouteGatewaysClient) ListByResourceGroupHandleResponse(resp *azcore.Response) (*ExpressRouteGatewayListResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.ListByResourceGroupHandleError(resp)
	}
	result := ExpressRouteGatewayListResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteGatewayList)
}

// ListByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *ExpressRouteGatewaysClient) ListByResourceGroupHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// ListBySubscription - Lists ExpressRoute gateways under a given subscription.
func (client *ExpressRouteGatewaysClient) ListBySubscription(ctx context.Context) (*ExpressRouteGatewayListResponse, error) {
	req, err := client.ListBySubscriptionCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.ListBySubscriptionHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ExpressRouteGatewaysClient) ListBySubscriptionCreateRequest() (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteGateways"
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

// ListBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ExpressRouteGatewaysClient) ListBySubscriptionHandleResponse(resp *azcore.Response) (*ExpressRouteGatewayListResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.ListBySubscriptionHandleError(resp)
	}
	result := ExpressRouteGatewayListResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteGatewayList)
}

// ListBySubscriptionHandleError handles the ListBySubscription error response.
func (client *ExpressRouteGatewaysClient) ListBySubscriptionHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}