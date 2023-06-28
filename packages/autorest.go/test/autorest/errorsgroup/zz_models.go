//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package errorsgroup

type Pet struct {
	AniType *string

	// READ-ONLY; Gets the Pet by id.
	Name *string
}

type PetAction struct {
	// action feedback
	ActionResponse *string
}