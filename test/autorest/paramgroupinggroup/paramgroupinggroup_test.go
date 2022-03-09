// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package paramgroupinggroup

import (
	"context"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func newParameterGroupingClient() *ParameterGroupingClient {
	return NewParameterGroupingClient(nil)
}

// PostMultiParamGroups - Post parameters from multiple different parameter groups
func TestPostMultiParamGroups(t *testing.T) {
	client := newParameterGroupingClient()
	result, err := client.PostMultiParamGroups(context.Background(), nil, nil)
	require.NoError(t, err)
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

// PostOptional - Post a bunch of optional parameters grouped
func TestPostOptional(t *testing.T) {
	client := newParameterGroupingClient()
	result, err := client.PostOptional(context.Background(), nil)
	require.NoError(t, err)
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

// PostRequired - Post a bunch of required parameters grouped
func TestPostRequired(t *testing.T) {
	client := newParameterGroupingClient()
	result, err := client.PostRequired(context.Background(), ParameterGroupingClientPostRequiredParameters{
		Body: 1234,
		Path: "path",
	})
	require.NoError(t, err)
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

// PostSharedParameterGroupObject - Post parameters with a shared parameter group object
func TestPostSharedParameterGroupObject(t *testing.T) {
	client := newParameterGroupingClient()
	result, err := client.PostSharedParameterGroupObject(context.Background(), nil)
	require.NoError(t, err)
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}
