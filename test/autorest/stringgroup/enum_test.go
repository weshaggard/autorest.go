// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package stringgroup

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func newEnumClient() *EnumClient {
	return NewEnumClient(nil)
}

func TestEnumGetNotExpandable(t *testing.T) {
	client := newEnumClient()
	result, err := client.GetNotExpandable(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, ColorsRedColor.ToPtr()); r != "" {
		t.Fatal(r)
	}
}

func TestEnumGetReferenced(t *testing.T) {
	client := newEnumClient()
	result, err := client.GetReferenced(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, ColorsRedColor.ToPtr()); r != "" {
		t.Fatal(r)
	}
}

func TestEnumGetReferencedConstant(t *testing.T) {
	client := newEnumClient()
	result, err := client.GetReferencedConstant(context.Background(), nil)
	require.NoError(t, err)
	val := "Sample String"
	if r := cmp.Diff(result.RefColorConstant, RefColorConstant{Field1: &val}); r != "" {
		t.Fatal(r)
	}
}

func TestEnumPutNotExpandable(t *testing.T) {
	client := newEnumClient()
	result, err := client.PutNotExpandable(context.Background(), ColorsRedColor, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestEnumPutReferenced(t *testing.T) {
	client := newEnumClient()
	result, err := client.PutReferenced(context.Background(), ColorsRedColor, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestEnumPutReferencedConstant(t *testing.T) {
	client := newEnumClient()
	val := string(ColorsGreenColor)
	result, err := client.PutReferencedConstant(context.Background(), RefColorConstant{ColorConstant: &val}, nil)
	require.NoError(t, err)
	require.Zero(t, result)

}
