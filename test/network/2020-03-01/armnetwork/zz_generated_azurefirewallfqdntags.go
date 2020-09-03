// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"path"
	"strings"
)

// AzureFirewallFqdnTagsOperations contains the methods for the AzureFirewallFqdnTags group.
type AzureFirewallFqdnTagsOperations interface {
	// ListAll - Gets all the Azure Firewall FQDN Tags in a subscription.
	ListAll() (AzureFirewallFqdnTagListResultPager, error)
}

// AzureFirewallFqdnTagsClient implements the AzureFirewallFqdnTagsOperations interface.
// Don't use this type directly, use NewAzureFirewallFqdnTagsClient() instead.
type AzureFirewallFqdnTagsClient struct {
	*Client
	subscriptionID string
}

// NewAzureFirewallFqdnTagsClient creates a new instance of AzureFirewallFqdnTagsClient with the specified values.
func NewAzureFirewallFqdnTagsClient(c *Client, subscriptionID string) AzureFirewallFqdnTagsOperations {
	return &AzureFirewallFqdnTagsClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *AzureFirewallFqdnTagsClient) Do(ctx context.Context, req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(ctx, req)
}

// ListAll - Gets all the Azure Firewall FQDN Tags in a subscription.
func (client *AzureFirewallFqdnTagsClient) ListAll() (AzureFirewallFqdnTagListResultPager, error) {
	req, err := client.ListAllCreateRequest()
	if err != nil {
		return nil, err
	}
	return &azureFirewallFqdnTagListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.ListAllHandleResponse,
		advancer: func(resp *AzureFirewallFqdnTagListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.AzureFirewallFqdnTagListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.AzureFirewallFqdnTagListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// ListAllCreateRequest creates the ListAll request.
func (client *AzureFirewallFqdnTagsClient) ListAllCreateRequest() (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/azureFirewallFqdnTags"
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

// ListAllHandleResponse handles the ListAll response.
func (client *AzureFirewallFqdnTagsClient) ListAllHandleResponse(resp *azcore.Response) (*AzureFirewallFqdnTagListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.ListAllHandleError(resp)
	}
	result := AzureFirewallFqdnTagListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.AzureFirewallFqdnTagListResult)
}

// ListAllHandleError handles the ListAll error response.
func (client *AzureFirewallFqdnTagsClient) ListAllHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}