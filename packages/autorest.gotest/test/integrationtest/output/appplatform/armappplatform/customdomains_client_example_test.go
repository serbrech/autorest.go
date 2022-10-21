//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2020-11-01-preview/examples/CustomDomains_Get.json
func ExampleCustomDomainsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armappplatform.NewCustomDomainsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "myResourceGroup", "myservice", "myapp", "mydomain.com", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CustomDomainResource = armappplatform.CustomDomainResource{
	// 	Name: to.Ptr("mydomain.com"),
	// 	Type: to.Ptr("Microsoft.AppPlatform/Spring/apps/domains"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/apps/myapp/domains/mydomain.com"),
	// 	Properties: &armappplatform.CustomDomainProperties{
	// 		AppName: to.Ptr("myapp"),
	// 		CertName: to.Ptr("mycert"),
	// 		Thumbprint: to.Ptr("934367bf1c97033f877db0f15cb1b586957d3133"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2020-11-01-preview/examples/CustomDomains_CreateOrUpdate.json
func ExampleCustomDomainsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armappplatform.NewCustomDomainsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "myResourceGroup", "myservice", "myapp", "mydomain.com", armappplatform.CustomDomainResource{
		Properties: &armappplatform.CustomDomainProperties{
			CertName:   to.Ptr("mycert"),
			Thumbprint: to.Ptr("934367bf1c97033f877db0f15cb1b586957d3133"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CustomDomainResource = armappplatform.CustomDomainResource{
	// 	Name: to.Ptr("mydomain.com"),
	// 	Type: to.Ptr("Microsoft.AppPlatform/Spring/apps/domains"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/apps/myapp/domains/mydomain.com"),
	// 	Properties: &armappplatform.CustomDomainProperties{
	// 		AppName: to.Ptr("myapp"),
	// 		CertName: to.Ptr("mycert"),
	// 		Thumbprint: to.Ptr("934367bf1c97033f877db0f15cb1b586957d3133"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2020-11-01-preview/examples/CustomDomains_Delete.json
func ExampleCustomDomainsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armappplatform.NewCustomDomainsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx, "myResourceGroup", "myservice", "myapp", "mydomain.com", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2020-11-01-preview/examples/CustomDomains_Update.json
func ExampleCustomDomainsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armappplatform.NewCustomDomainsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx, "myResourceGroup", "myservice", "myapp", "mydomain.com", armappplatform.CustomDomainResource{
		Properties: &armappplatform.CustomDomainProperties{
			CertName:   to.Ptr("mycert"),
			Thumbprint: to.Ptr("934367bf1c97033f877db0f15cb1b586957d3133"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CustomDomainResource = armappplatform.CustomDomainResource{
	// 	Name: to.Ptr("mydomain.com"),
	// 	Type: to.Ptr("Microsoft.AppPlatform/Spring/apps/domains"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/apps/myapp/domains/mydomain.com"),
	// 	Properties: &armappplatform.CustomDomainProperties{
	// 		AppName: to.Ptr("myapp"),
	// 		CertName: to.Ptr("mycert"),
	// 		Thumbprint: to.Ptr("934367bf1c97033f877db0f15cb1b586957d3133"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2020-11-01-preview/examples/CustomDomains_List.json
func ExampleCustomDomainsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armappplatform.NewCustomDomainsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("myResourceGroup", "myservice", "myapp", nil)
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
		// page.CustomDomainResourceCollection = armappplatform.CustomDomainResourceCollection{
		// 	Value: []*armappplatform.CustomDomainResource{
		// 		{
		// 			Name: to.Ptr("mydomain.com"),
		// 			Type: to.Ptr("Microsoft.AppPlatform/Spring/apps/domains"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/apps/myapp/domains/mydomain.com"),
		// 			Properties: &armappplatform.CustomDomainProperties{
		// 				AppName: to.Ptr("myapp"),
		// 				CertName: to.Ptr("mycert"),
		// 				Thumbprint: to.Ptr("934367bf1c97033f877db0f15cb1b586957d3133"),
		// 			},
		// 	}},
		// }
	}
}
