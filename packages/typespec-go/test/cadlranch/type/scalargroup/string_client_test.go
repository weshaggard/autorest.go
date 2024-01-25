//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package scalargroup_test

import (
	"context"
	"scalargroup"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStringClient_Get(t *testing.T) {
	client, err := scalargroup.NewStringClient(nil)
	require.NoError(t, err)
	resp, err := client.Get(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, "test", *resp.Value)
}

func TestStringClient_Put(t *testing.T) {
	client, err := scalargroup.NewStringClient(nil)
	require.NoError(t, err)
	resp, err := client.Put(context.Background(), "test", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
