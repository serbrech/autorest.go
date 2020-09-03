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

// DefaultSecurityRulesOperations contains the methods for the DefaultSecurityRules group.
type DefaultSecurityRulesOperations interface {
	// Get - Get the specified default network security rule.
	Get(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, defaultSecurityRuleName string) (*SecurityRuleResponse, error)
	// List - Gets all default security rules in a network security group.
	List(resourceGroupName string, networkSecurityGroupName string) (SecurityRuleListResultPager, error)
}

// DefaultSecurityRulesClient implements the DefaultSecurityRulesOperations interface.
// Don't use this type directly, use NewDefaultSecurityRulesClient() instead.
type DefaultSecurityRulesClient struct {
	*Client
	subscriptionID string
}

// NewDefaultSecurityRulesClient creates a new instance of DefaultSecurityRulesClient with the specified values.
func NewDefaultSecurityRulesClient(c *Client, subscriptionID string) DefaultSecurityRulesOperations {
	return &DefaultSecurityRulesClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *DefaultSecurityRulesClient) Do(ctx context.Context, req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(ctx, req)
}

// Get - Get the specified default network security rule.
func (client *DefaultSecurityRulesClient) Get(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, defaultSecurityRuleName string) (*SecurityRuleResponse, error) {
	req, err := client.GetCreateRequest(resourceGroupName, networkSecurityGroupName, defaultSecurityRuleName)
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
func (client *DefaultSecurityRulesClient) GetCreateRequest(resourceGroupName string, networkSecurityGroupName string, defaultSecurityRuleName string) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/defaultSecurityRules/{defaultSecurityRuleName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityGroupName}", url.PathEscape(networkSecurityGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{defaultSecurityRuleName}", url.PathEscape(defaultSecurityRuleName))
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
func (client *DefaultSecurityRulesClient) GetHandleResponse(resp *azcore.Response) (*SecurityRuleResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetHandleError(resp)
	}
	result := SecurityRuleResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.SecurityRule)
}

// GetHandleError handles the Get error response.
func (client *DefaultSecurityRulesClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// List - Gets all default security rules in a network security group.
func (client *DefaultSecurityRulesClient) List(resourceGroupName string, networkSecurityGroupName string) (SecurityRuleListResultPager, error) {
	req, err := client.ListCreateRequest(resourceGroupName, networkSecurityGroupName)
	if err != nil {
		return nil, err
	}
	return &securityRuleListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.ListHandleResponse,
		advancer: func(resp *SecurityRuleListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.SecurityRuleListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.SecurityRuleListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// ListCreateRequest creates the List request.
func (client *DefaultSecurityRulesClient) ListCreateRequest(resourceGroupName string, networkSecurityGroupName string) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/defaultSecurityRules"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityGroupName}", url.PathEscape(networkSecurityGroupName))
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
func (client *DefaultSecurityRulesClient) ListHandleResponse(resp *azcore.Response) (*SecurityRuleListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.ListHandleError(resp)
	}
	result := SecurityRuleListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.SecurityRuleListResult)
}

// ListHandleError handles the List error response.
func (client *DefaultSecurityRulesClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}