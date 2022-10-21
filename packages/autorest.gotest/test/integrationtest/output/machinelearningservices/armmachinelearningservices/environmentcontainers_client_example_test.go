//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmachinelearningservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearningservices/armmachinelearningservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/EnvironmentContainer/list.json
func ExampleEnvironmentContainersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearningservices.NewEnvironmentContainersClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("testrg123", "testworkspace", &armmachinelearningservices.EnvironmentContainersClientListOptions{Skip: nil,
		ListViewType: nil,
	})
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
		// page.EnvironmentContainerResourceArmPaginatedResult = armmachinelearningservices.EnvironmentContainerResourceArmPaginatedResult{
		// 	Value: []*armmachinelearningservices.EnvironmentContainerData{
		// 		{
		// 			Name: to.Ptr("testEnvironment"),
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/environments"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg123/providers/Microsoft.MachineLearningServices/workspaces/testworkspace/environments/testEnvironment"),
		// 			SystemData: &armmachinelearningservices.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-01T12:00:00.000Z"); return t}()),
		// 				CreatedBy: to.Ptr("John Smith"),
		// 				CreatedByType: to.Ptr(armmachinelearningservices.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-01T12:00:00.000Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("John Smith"),
		// 				LastModifiedByType: to.Ptr(armmachinelearningservices.CreatedByTypeUser),
		// 			},
		// 			Properties: &armmachinelearningservices.EnvironmentContainerDetails{
		// 				Description: to.Ptr("string"),
		// 				Tags: map[string]*string{
		// 					"tag1": to.Ptr("value1"),
		// 					"tag2": to.Ptr("value2"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/EnvironmentContainer/delete.json
func ExampleEnvironmentContainersClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearningservices.NewEnvironmentContainersClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx, "testrg123", "testworkspace", "testContainer", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/EnvironmentContainer/get.json
func ExampleEnvironmentContainersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearningservices.NewEnvironmentContainersClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "testrg123", "testworkspace", "testEnvironment", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.EnvironmentContainerData = armmachinelearningservices.EnvironmentContainerData{
	// 	Name: to.Ptr("testEnvironment"),
	// 	Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/environments"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg123/providers/Microsoft.MachineLearningServices/workspaces/testworkspace/environments/testEnvironment"),
	// 	SystemData: &armmachinelearningservices.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-01T12:00:00.000Z"); return t}()),
	// 		CreatedBy: to.Ptr("John Smith"),
	// 		CreatedByType: to.Ptr(armmachinelearningservices.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-01T12:00:00.000Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("John Smith"),
	// 		LastModifiedByType: to.Ptr(armmachinelearningservices.CreatedByTypeUser),
	// 	},
	// 	Properties: &armmachinelearningservices.EnvironmentContainerDetails{
	// 		Description: to.Ptr("string"),
	// 		Tags: map[string]*string{
	// 			"tag1": to.Ptr("value1"),
	// 			"tag2": to.Ptr("value2"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/EnvironmentContainer/createOrUpdate.json
func ExampleEnvironmentContainersClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearningservices.NewEnvironmentContainersClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "testrg123", "testworkspace", "testEnvironment", armmachinelearningservices.EnvironmentContainerData{
		Properties: &armmachinelearningservices.EnvironmentContainerDetails{
			Description: to.Ptr("string"),
			Properties: map[string]*string{
				"additionalProp1": to.Ptr("string"),
				"additionalProp2": to.Ptr("string"),
				"additionalProp3": to.Ptr("string"),
			},
			Tags: map[string]*string{
				"additionalProp1": to.Ptr("string"),
				"additionalProp2": to.Ptr("string"),
				"additionalProp3": to.Ptr("string"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.EnvironmentContainerData = armmachinelearningservices.EnvironmentContainerData{
	// 	Name: to.Ptr("testEnvironment"),
	// 	Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/environments"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg123/providers/Microsoft.MachineLearningServices/workspaces/testworkspace/environments/testEnvironment"),
	// 	SystemData: &armmachinelearningservices.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-04T03:39:11.300Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armmachinelearningservices.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-04T03:39:11.300Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armmachinelearningservices.CreatedByTypeUser),
	// 	},
	// 	Properties: &armmachinelearningservices.EnvironmentContainerDetails{
	// 		Description: to.Ptr("string"),
	// 		Properties: map[string]*string{
	// 			"additionalProp1": to.Ptr("string"),
	// 			"additionalProp2": to.Ptr("string"),
	// 			"additionalProp3": to.Ptr("string"),
	// 		},
	// 		Tags: map[string]*string{
	// 			"additionalProp1": to.Ptr("string"),
	// 			"additionalProp2": to.Ptr("string"),
	// 			"additionalProp3": to.Ptr("string"),
	// 		},
	// 	},
	// }
}
