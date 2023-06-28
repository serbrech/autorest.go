//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconsumption

// ChargeSummaryClassification provides polymorphic access to related types.
// Call the interface's GetChargeSummary() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *ChargeSummary, *LegacyChargeSummary, *ModernChargeSummary
type ChargeSummaryClassification interface {
	// GetChargeSummary returns the ChargeSummary content of the underlying type.
	GetChargeSummary() *ChargeSummary
}

// ReservationRecommendationClassification provides polymorphic access to related types.
// Call the interface's GetReservationRecommendation() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *LegacyReservationRecommendation, *ModernReservationRecommendation, *ReservationRecommendation
type ReservationRecommendationClassification interface {
	// GetReservationRecommendation returns the ReservationRecommendation content of the underlying type.
	GetReservationRecommendation() *ReservationRecommendation
}

// UsageDetailClassification provides polymorphic access to related types.
// Call the interface's GetUsageDetail() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *LegacyUsageDetail, *ModernUsageDetail, *UsageDetail
type UsageDetailClassification interface {
	// GetUsageDetail returns the UsageDetail content of the underlying type.
	GetUsageDetail() *UsageDetail
}