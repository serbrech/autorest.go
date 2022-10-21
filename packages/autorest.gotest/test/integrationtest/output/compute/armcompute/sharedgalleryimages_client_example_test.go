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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/compute/resource-manager/Microsoft.Compute/stable/2020-09-30/examples/ListSharedGalleryImages.json
func ExampleSharedGalleryImagesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewSharedGalleryImagesClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("myLocation", "galleryUniqueName", &armcompute.SharedGalleryImagesClientListOptions{SharedTo: nil})
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
		// page.SharedGalleryImageList = armcompute.SharedGalleryImageList{
		// 	Value: []*armcompute.SharedGalleryImage{
		// 		{
		// 			Name: to.Ptr("myGalleryImageName"),
		// 			Location: to.Ptr("myLocation"),
		// 			Identifier: &armcompute.SharedGalleryIdentifier{
		// 				UniqueID: to.Ptr("/SharedGalleries/galleryUniqueName/Images/myGalleryImageName"),
		// 			},
		// 			Properties: &armcompute.SharedGalleryImageProperties{
		// 				HyperVGeneration: to.Ptr(armcompute.HyperVGenerationV1),
		// 				Identifier: &armcompute.GalleryImageIdentifier{
		// 					Offer: to.Ptr("myOfferName"),
		// 					Publisher: to.Ptr("myPublisherName"),
		// 					SKU: to.Ptr("mySkuName"),
		// 				},
		// 				OSState: to.Ptr(armcompute.OperatingSystemStateTypesGeneralized),
		// 				OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/compute/resource-manager/Microsoft.Compute/stable/2020-09-30/examples/GetASharedGalleryImage.json
func ExampleSharedGalleryImagesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewSharedGalleryImagesClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "myLocation", "galleryUniqueName", "myGalleryImageName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SharedGalleryImage = armcompute.SharedGalleryImage{
	// 	Name: to.Ptr("myGalleryImageName"),
	// 	Location: to.Ptr("myLocation"),
	// 	Identifier: &armcompute.SharedGalleryIdentifier{
	// 		UniqueID: to.Ptr("/SharedGalleries/galleryUniqueName/Images/myGalleryImageName"),
	// 	},
	// 	Properties: &armcompute.SharedGalleryImageProperties{
	// 		HyperVGeneration: to.Ptr(armcompute.HyperVGenerationV1),
	// 		Identifier: &armcompute.GalleryImageIdentifier{
	// 			Offer: to.Ptr("myOfferName"),
	// 			Publisher: to.Ptr("myPublisherName"),
	// 			SKU: to.Ptr("mySkuName"),
	// 		},
	// 		OSState: to.Ptr(armcompute.OperatingSystemStateTypesGeneralized),
	// 		OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
	// 	},
	// }
}
