//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"armnetwork"
	"encoding/json"
)

func unmarshalBaseAdminRuleClassification(rawMsg json.RawMessage) (armnetwork.BaseAdminRuleClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b armnetwork.BaseAdminRuleClassification
	switch m["kind"] {
	case string(armnetwork.AdminRuleKindCustom):
		b = &armnetwork.AdminRule{}
	case string(armnetwork.AdminRuleKindDefault):
		b = &armnetwork.DefaultAdminRule{}
	default:
		b = &armnetwork.BaseAdminRule{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}
