//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsignalr_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/signalr/armsignalr"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/signalr/resource-manager/Microsoft.SignalRService/preview/2021-06-01-preview/examples/SignalRPrivateLinkResources_List.json
func ExamplePrivateLinkResourcesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsignalr.NewPrivateLinkResourcesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("myResourceGroup", "mySignalRService", nil)
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
		// page.PrivateLinkResourceList = armsignalr.PrivateLinkResourceList{
		// 	Value: []*armsignalr.PrivateLinkResource{
		// 		{
		// 			Name: to.Ptr("myPrivateLink"),
		// 			Type: to.Ptr("privateLinkResources"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.SignalRService/SignalR/mySignalRService/privateLinkResources/myPrivateLink"),
		// 			Properties: &armsignalr.PrivateLinkResourceProperties{
		// 				GroupID: to.Ptr("signalr"),
		// 				RequiredMembers: []*string{
		// 					to.Ptr("signalr")},
		// 					RequiredZoneNames: []*string{
		// 						to.Ptr("privatelink.service.signalr.net")},
		// 						ShareablePrivateLinkResourceTypes: []*armsignalr.ShareablePrivateLinkResourceType{
		// 							{
		// 								Name: to.Ptr("site"),
		// 								Properties: &armsignalr.ShareablePrivateLinkResourceProperties{
		// 									Type: to.Ptr("Microsoft.Web/sites"),
		// 									Description: to.Ptr("Azure App Service can be used as an upstream"),
		// 									GroupID: to.Ptr("sites"),
		// 								},
		// 						}},
		// 					},
		// 			}},
		// 		}
	}
}
