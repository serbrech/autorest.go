//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/compute/armcompute"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/UpdateCloudServiceUpdateDomain.json
func ExampleCloudServicesUpdateDomainClient_BeginWalkUpdateDomain() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewCloudServicesUpdateDomainClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginWalkUpdateDomain(ctx, "ConstosoRG", "{cs-name}", 1, armcompute.UpdateDomain{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/GetCloudServiceUpdateDomain.json
func ExampleCloudServicesUpdateDomainClient_GetUpdateDomain() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewCloudServicesUpdateDomainClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetUpdateDomain(ctx, "ConstosoRG", "{cs-name}", 1, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.UpdateDomain = armcompute.UpdateDomain{
	// 	Name: to.Ptr("1"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/ConstosoRG/providers/Microsoft.Compute/cloudServices/{cs-name}/updateDomains/1"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/ListCloudServiceUpdateDomains.json
func ExampleCloudServicesUpdateDomainClient_NewListUpdateDomainsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewCloudServicesUpdateDomainClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListUpdateDomainsPager("ConstosoRG", "{cs-name}", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.UpdateDomainListResult = armcompute.UpdateDomainListResult{
		// 	Value: []*armcompute.UpdateDomain{
		// 		{
		// 			Name: to.Ptr("0"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/ConstosoRG/providers/Microsoft.Compute/cloudServices/{cs-name}/updateDomains/0"),
		// 		},
		// 		{
		// 			Name: to.Ptr("1"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/ConstosoRG/providers/Microsoft.Compute/cloudServices/{cs-name}/updateDomains/1"),
		// 	}},
		// }
	}
}
