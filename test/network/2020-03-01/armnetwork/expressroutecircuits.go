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

// ExpressRouteCircuitsOperations contains the methods for the ExpressRouteCircuits group.
type ExpressRouteCircuitsOperations interface {
	// BeginCreateOrUpdate - Creates or updates an express route circuit.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, circuitName string, parameters ExpressRouteCircuit) (*ExpressRouteCircuitPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (ExpressRouteCircuitPoller, error)
	// BeginDelete - Deletes the specified express route circuit.
	BeginDelete(ctx context.Context, resourceGroupName string, circuitName string) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets information about the specified express route circuit.
	Get(ctx context.Context, resourceGroupName string, circuitName string) (*ExpressRouteCircuitResponse, error)
	// GetPeeringStats - Gets all stats from an express route circuit in a resource group.
	GetPeeringStats(ctx context.Context, resourceGroupName string, circuitName string, peeringName string) (*ExpressRouteCircuitStatsResponse, error)
	// GetStats - Gets all the stats from an express route circuit in a resource group.
	GetStats(ctx context.Context, resourceGroupName string, circuitName string) (*ExpressRouteCircuitStatsResponse, error)
	// List - Gets all the express route circuits in a resource group.
	List(resourceGroupName string) (ExpressRouteCircuitListResultPager, error)
	// ListAll - Gets all the express route circuits in a subscription.
	ListAll() (ExpressRouteCircuitListResultPager, error)
	// BeginListArpTable - Gets the currently advertised ARP table associated with the express route circuit in a resource group.
	BeginListArpTable(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string) (*ExpressRouteCircuitsArpTableListResultPollerResponse, error)
	// ResumeListArpTable - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeListArpTable(token string) (ExpressRouteCircuitsArpTableListResultPoller, error)
	// BeginListRoutesTable - Gets the currently advertised routes table associated with the express route circuit in a resource group.
	BeginListRoutesTable(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string) (*ExpressRouteCircuitsRoutesTableListResultPollerResponse, error)
	// ResumeListRoutesTable - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeListRoutesTable(token string) (ExpressRouteCircuitsRoutesTableListResultPoller, error)
	// BeginListRoutesTableSummary - Gets the currently advertised routes table summary associated with the express route circuit in a resource group.
	BeginListRoutesTableSummary(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string) (*ExpressRouteCircuitsRoutesTableSummaryListResultPollerResponse, error)
	// ResumeListRoutesTableSummary - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeListRoutesTableSummary(token string) (ExpressRouteCircuitsRoutesTableSummaryListResultPoller, error)
	// UpdateTags - Updates an express route circuit tags.
	UpdateTags(ctx context.Context, resourceGroupName string, circuitName string, parameters TagsObject) (*ExpressRouteCircuitResponse, error)
}

// expressRouteCircuitsOperations implements the ExpressRouteCircuitsOperations interface.
type expressRouteCircuitsOperations struct {
	*Client
	subscriptionID string
}

// CreateOrUpdate - Creates or updates an express route circuit.
func (client *expressRouteCircuitsOperations) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, circuitName string, parameters ExpressRouteCircuit) (*ExpressRouteCircuitPollerResponse, error) {
	req, err := client.createOrUpdateCreateRequest(resourceGroupName, circuitName, parameters)
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
	pt, err := armcore.NewPoller("expressRouteCircuitsOperations.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &expressRouteCircuitPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ExpressRouteCircuitResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *expressRouteCircuitsOperations) ResumeCreateOrUpdate(token string) (ExpressRouteCircuitPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("expressRouteCircuitsOperations.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &expressRouteCircuitPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *expressRouteCircuitsOperations) createOrUpdateCreateRequest(resourceGroupName string, circuitName string, parameters ExpressRouteCircuit) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
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
func (client *expressRouteCircuitsOperations) createOrUpdateHandleResponse(resp *azcore.Response) (*ExpressRouteCircuitPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated, http.StatusNoContent) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return &ExpressRouteCircuitPollerResponse{RawResponse: resp.Response}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *expressRouteCircuitsOperations) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Delete - Deletes the specified express route circuit.
func (client *expressRouteCircuitsOperations) BeginDelete(ctx context.Context, resourceGroupName string, circuitName string) (*HTTPPollerResponse, error) {
	req, err := client.deleteCreateRequest(resourceGroupName, circuitName)
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
	pt, err := armcore.NewPoller("expressRouteCircuitsOperations.Delete", "location", resp, client.deleteHandleError)
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

func (client *expressRouteCircuitsOperations) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("expressRouteCircuitsOperations.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *expressRouteCircuitsOperations) deleteCreateRequest(resourceGroupName string, circuitName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
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
func (client *expressRouteCircuitsOperations) deleteHandleResponse(resp *azcore.Response) (*HTTPPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return &HTTPPollerResponse{RawResponse: resp.Response}, nil
}

// deleteHandleError handles the Delete error response.
func (client *expressRouteCircuitsOperations) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Get - Gets information about the specified express route circuit.
func (client *expressRouteCircuitsOperations) Get(ctx context.Context, resourceGroupName string, circuitName string) (*ExpressRouteCircuitResponse, error) {
	req, err := client.getCreateRequest(resourceGroupName, circuitName)
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
func (client *expressRouteCircuitsOperations) getCreateRequest(resourceGroupName string, circuitName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
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
func (client *expressRouteCircuitsOperations) getHandleResponse(resp *azcore.Response) (*ExpressRouteCircuitResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getHandleError(resp)
	}
	result := ExpressRouteCircuitResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCircuit)
}

// getHandleError handles the Get error response.
func (client *expressRouteCircuitsOperations) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetPeeringStats - Gets all stats from an express route circuit in a resource group.
func (client *expressRouteCircuitsOperations) GetPeeringStats(ctx context.Context, resourceGroupName string, circuitName string, peeringName string) (*ExpressRouteCircuitStatsResponse, error) {
	req, err := client.getPeeringStatsCreateRequest(resourceGroupName, circuitName, peeringName)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getPeeringStatsHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getPeeringStatsCreateRequest creates the GetPeeringStats request.
func (client *expressRouteCircuitsOperations) getPeeringStatsCreateRequest(resourceGroupName string, circuitName string, peeringName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/stats"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
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

// getPeeringStatsHandleResponse handles the GetPeeringStats response.
func (client *expressRouteCircuitsOperations) getPeeringStatsHandleResponse(resp *azcore.Response) (*ExpressRouteCircuitStatsResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getPeeringStatsHandleError(resp)
	}
	result := ExpressRouteCircuitStatsResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCircuitStats)
}

// getPeeringStatsHandleError handles the GetPeeringStats error response.
func (client *expressRouteCircuitsOperations) getPeeringStatsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetStats - Gets all the stats from an express route circuit in a resource group.
func (client *expressRouteCircuitsOperations) GetStats(ctx context.Context, resourceGroupName string, circuitName string) (*ExpressRouteCircuitStatsResponse, error) {
	req, err := client.getStatsCreateRequest(resourceGroupName, circuitName)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getStatsHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getStatsCreateRequest creates the GetStats request.
func (client *expressRouteCircuitsOperations) getStatsCreateRequest(resourceGroupName string, circuitName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/stats"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
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

// getStatsHandleResponse handles the GetStats response.
func (client *expressRouteCircuitsOperations) getStatsHandleResponse(resp *azcore.Response) (*ExpressRouteCircuitStatsResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getStatsHandleError(resp)
	}
	result := ExpressRouteCircuitStatsResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCircuitStats)
}

// getStatsHandleError handles the GetStats error response.
func (client *expressRouteCircuitsOperations) getStatsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// List - Gets all the express route circuits in a resource group.
func (client *expressRouteCircuitsOperations) List(resourceGroupName string) (ExpressRouteCircuitListResultPager, error) {
	req, err := client.listCreateRequest(resourceGroupName)
	if err != nil {
		return nil, err
	}
	return &expressRouteCircuitListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listHandleResponse,
		advancer: func(resp *ExpressRouteCircuitListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.ExpressRouteCircuitListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.ExpressRouteCircuitListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listCreateRequest creates the List request.
func (client *expressRouteCircuitsOperations) listCreateRequest(resourceGroupName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits"
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
func (client *expressRouteCircuitsOperations) listHandleResponse(resp *azcore.Response) (*ExpressRouteCircuitListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listHandleError(resp)
	}
	result := ExpressRouteCircuitListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCircuitListResult)
}

// listHandleError handles the List error response.
func (client *expressRouteCircuitsOperations) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// ListAll - Gets all the express route circuits in a subscription.
func (client *expressRouteCircuitsOperations) ListAll() (ExpressRouteCircuitListResultPager, error) {
	req, err := client.listAllCreateRequest()
	if err != nil {
		return nil, err
	}
	return &expressRouteCircuitListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listAllHandleResponse,
		advancer: func(resp *ExpressRouteCircuitListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.ExpressRouteCircuitListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.ExpressRouteCircuitListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listAllCreateRequest creates the ListAll request.
func (client *expressRouteCircuitsOperations) listAllCreateRequest() (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteCircuits"
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
func (client *expressRouteCircuitsOperations) listAllHandleResponse(resp *azcore.Response) (*ExpressRouteCircuitListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listAllHandleError(resp)
	}
	result := ExpressRouteCircuitListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCircuitListResult)
}

// listAllHandleError handles the ListAll error response.
func (client *expressRouteCircuitsOperations) listAllHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// ListArpTable - Gets the currently advertised ARP table associated with the express route circuit in a resource group.
func (client *expressRouteCircuitsOperations) BeginListArpTable(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string) (*ExpressRouteCircuitsArpTableListResultPollerResponse, error) {
	req, err := client.listArpTableCreateRequest(resourceGroupName, circuitName, peeringName, devicePath)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.listArpTableHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := armcore.NewPoller("expressRouteCircuitsOperations.ListArpTable", "location", resp, client.listArpTableHandleError)
	if err != nil {
		return nil, err
	}
	poller := &expressRouteCircuitsArpTableListResultPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ExpressRouteCircuitsArpTableListResultResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *expressRouteCircuitsOperations) ResumeListArpTable(token string) (ExpressRouteCircuitsArpTableListResultPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("expressRouteCircuitsOperations.ListArpTable", token, client.listArpTableHandleError)
	if err != nil {
		return nil, err
	}
	return &expressRouteCircuitsArpTableListResultPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// listArpTableCreateRequest creates the ListArpTable request.
func (client *expressRouteCircuitsOperations) listArpTableCreateRequest(resourceGroupName string, circuitName string, peeringName string, devicePath string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/arpTables/{devicePath}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	urlPath = strings.ReplaceAll(urlPath, "{devicePath}", url.PathEscape(devicePath))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPost, *u)
	return req, nil
}

// listArpTableHandleResponse handles the ListArpTable response.
func (client *expressRouteCircuitsOperations) listArpTableHandleResponse(resp *azcore.Response) (*ExpressRouteCircuitsArpTableListResultPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.listArpTableHandleError(resp)
	}
	return &ExpressRouteCircuitsArpTableListResultPollerResponse{RawResponse: resp.Response}, nil
}

// listArpTableHandleError handles the ListArpTable error response.
func (client *expressRouteCircuitsOperations) listArpTableHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// ListRoutesTable - Gets the currently advertised routes table associated with the express route circuit in a resource group.
func (client *expressRouteCircuitsOperations) BeginListRoutesTable(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string) (*ExpressRouteCircuitsRoutesTableListResultPollerResponse, error) {
	req, err := client.listRoutesTableCreateRequest(resourceGroupName, circuitName, peeringName, devicePath)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.listRoutesTableHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := armcore.NewPoller("expressRouteCircuitsOperations.ListRoutesTable", "location", resp, client.listRoutesTableHandleError)
	if err != nil {
		return nil, err
	}
	poller := &expressRouteCircuitsRoutesTableListResultPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ExpressRouteCircuitsRoutesTableListResultResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *expressRouteCircuitsOperations) ResumeListRoutesTable(token string) (ExpressRouteCircuitsRoutesTableListResultPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("expressRouteCircuitsOperations.ListRoutesTable", token, client.listRoutesTableHandleError)
	if err != nil {
		return nil, err
	}
	return &expressRouteCircuitsRoutesTableListResultPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// listRoutesTableCreateRequest creates the ListRoutesTable request.
func (client *expressRouteCircuitsOperations) listRoutesTableCreateRequest(resourceGroupName string, circuitName string, peeringName string, devicePath string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/routeTables/{devicePath}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	urlPath = strings.ReplaceAll(urlPath, "{devicePath}", url.PathEscape(devicePath))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPost, *u)
	return req, nil
}

// listRoutesTableHandleResponse handles the ListRoutesTable response.
func (client *expressRouteCircuitsOperations) listRoutesTableHandleResponse(resp *azcore.Response) (*ExpressRouteCircuitsRoutesTableListResultPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.listRoutesTableHandleError(resp)
	}
	return &ExpressRouteCircuitsRoutesTableListResultPollerResponse{RawResponse: resp.Response}, nil
}

// listRoutesTableHandleError handles the ListRoutesTable error response.
func (client *expressRouteCircuitsOperations) listRoutesTableHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// ListRoutesTableSummary - Gets the currently advertised routes table summary associated with the express route circuit in a resource group.
func (client *expressRouteCircuitsOperations) BeginListRoutesTableSummary(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string) (*ExpressRouteCircuitsRoutesTableSummaryListResultPollerResponse, error) {
	req, err := client.listRoutesTableSummaryCreateRequest(resourceGroupName, circuitName, peeringName, devicePath)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.listRoutesTableSummaryHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := armcore.NewPoller("expressRouteCircuitsOperations.ListRoutesTableSummary", "location", resp, client.listRoutesTableSummaryHandleError)
	if err != nil {
		return nil, err
	}
	poller := &expressRouteCircuitsRoutesTableSummaryListResultPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ExpressRouteCircuitsRoutesTableSummaryListResultResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *expressRouteCircuitsOperations) ResumeListRoutesTableSummary(token string) (ExpressRouteCircuitsRoutesTableSummaryListResultPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("expressRouteCircuitsOperations.ListRoutesTableSummary", token, client.listRoutesTableSummaryHandleError)
	if err != nil {
		return nil, err
	}
	return &expressRouteCircuitsRoutesTableSummaryListResultPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// listRoutesTableSummaryCreateRequest creates the ListRoutesTableSummary request.
func (client *expressRouteCircuitsOperations) listRoutesTableSummaryCreateRequest(resourceGroupName string, circuitName string, peeringName string, devicePath string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/routeTablesSummary/{devicePath}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	urlPath = strings.ReplaceAll(urlPath, "{devicePath}", url.PathEscape(devicePath))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPost, *u)
	return req, nil
}

// listRoutesTableSummaryHandleResponse handles the ListRoutesTableSummary response.
func (client *expressRouteCircuitsOperations) listRoutesTableSummaryHandleResponse(resp *azcore.Response) (*ExpressRouteCircuitsRoutesTableSummaryListResultPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.listRoutesTableSummaryHandleError(resp)
	}
	return &ExpressRouteCircuitsRoutesTableSummaryListResultPollerResponse{RawResponse: resp.Response}, nil
}

// listRoutesTableSummaryHandleError handles the ListRoutesTableSummary error response.
func (client *expressRouteCircuitsOperations) listRoutesTableSummaryHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// UpdateTags - Updates an express route circuit tags.
func (client *expressRouteCircuitsOperations) UpdateTags(ctx context.Context, resourceGroupName string, circuitName string, parameters TagsObject) (*ExpressRouteCircuitResponse, error) {
	req, err := client.updateTagsCreateRequest(resourceGroupName, circuitName, parameters)
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
func (client *expressRouteCircuitsOperations) updateTagsCreateRequest(resourceGroupName string, circuitName string, parameters TagsObject) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
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
func (client *expressRouteCircuitsOperations) updateTagsHandleResponse(resp *azcore.Response) (*ExpressRouteCircuitResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.updateTagsHandleError(resp)
	}
	result := ExpressRouteCircuitResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCircuit)
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *expressRouteCircuitsOperations) updateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
