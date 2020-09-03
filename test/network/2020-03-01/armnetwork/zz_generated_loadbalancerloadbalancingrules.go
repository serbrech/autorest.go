// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"path"
	"strings"
)

// LoadBalancerLoadBalancingRulesOperations contains the methods for the LoadBalancerLoadBalancingRules group.
type LoadBalancerLoadBalancingRulesOperations interface {
	// Get - Gets the specified load balancer load balancing rule.
	Get(ctx context.Context, resourceGroupName string, loadBalancerName string, loadBalancingRuleName string) (*LoadBalancingRuleResponse, error)
	// List - Gets all the load balancing rules in a load balancer.
	List(resourceGroupName string, loadBalancerName string) (LoadBalancerLoadBalancingRuleListResultPager, error)
}

// LoadBalancerLoadBalancingRulesClient implements the LoadBalancerLoadBalancingRulesOperations interface.
// Don't use this type directly, use NewLoadBalancerLoadBalancingRulesClient() instead.
type LoadBalancerLoadBalancingRulesClient struct {
	*Client
	subscriptionID string
}

// NewLoadBalancerLoadBalancingRulesClient creates a new instance of LoadBalancerLoadBalancingRulesClient with the specified values.
func NewLoadBalancerLoadBalancingRulesClient(c *Client, subscriptionID string) LoadBalancerLoadBalancingRulesOperations {
	return &LoadBalancerLoadBalancingRulesClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *LoadBalancerLoadBalancingRulesClient) Do(ctx context.Context, req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(ctx, req)
}

// Get - Gets the specified load balancer load balancing rule.
func (client *LoadBalancerLoadBalancingRulesClient) Get(ctx context.Context, resourceGroupName string, loadBalancerName string, loadBalancingRuleName string) (*LoadBalancingRuleResponse, error) {
	req, err := client.GetCreateRequest(resourceGroupName, loadBalancerName, loadBalancingRuleName)
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
func (client *LoadBalancerLoadBalancingRulesClient) GetCreateRequest(resourceGroupName string, loadBalancerName string, loadBalancingRuleName string) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/loadBalancingRules/{loadBalancingRuleName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancingRuleName}", url.PathEscape(loadBalancingRuleName))
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
func (client *LoadBalancerLoadBalancingRulesClient) GetHandleResponse(resp *azcore.Response) (*LoadBalancingRuleResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetHandleError(resp)
	}
	result := LoadBalancingRuleResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.LoadBalancingRule)
}

// GetHandleError handles the Get error response.
func (client *LoadBalancerLoadBalancingRulesClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// List - Gets all the load balancing rules in a load balancer.
func (client *LoadBalancerLoadBalancingRulesClient) List(resourceGroupName string, loadBalancerName string) (LoadBalancerLoadBalancingRuleListResultPager, error) {
	req, err := client.ListCreateRequest(resourceGroupName, loadBalancerName)
	if err != nil {
		return nil, err
	}
	return &loadBalancerLoadBalancingRuleListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.ListHandleResponse,
		advancer: func(resp *LoadBalancerLoadBalancingRuleListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.LoadBalancerLoadBalancingRuleListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.LoadBalancerLoadBalancingRuleListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// ListCreateRequest creates the List request.
func (client *LoadBalancerLoadBalancingRulesClient) ListCreateRequest(resourceGroupName string, loadBalancerName string) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/loadBalancingRules"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
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
func (client *LoadBalancerLoadBalancingRulesClient) ListHandleResponse(resp *azcore.Response) (*LoadBalancerLoadBalancingRuleListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.ListHandleError(resp)
	}
	result := LoadBalancerLoadBalancingRuleListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.LoadBalancerLoadBalancingRuleListResult)
}

// ListHandleError handles the List error response.
func (client *LoadBalancerLoadBalancingRulesClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}