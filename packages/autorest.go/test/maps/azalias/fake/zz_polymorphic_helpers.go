//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"azalias"
	"encoding/json"
)

func unmarshalGeoJSONObjectClassification(rawMsg json.RawMessage) (azalias.GeoJSONObjectClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b azalias.GeoJSONObjectClassification
	switch m["type"] {
	case string(azalias.GeoJSONObjectTypeGeoJSONFeature):
		b = &azalias.GeoJSONFeature{}
	default:
		b = &azalias.GeoJSONObject{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}
