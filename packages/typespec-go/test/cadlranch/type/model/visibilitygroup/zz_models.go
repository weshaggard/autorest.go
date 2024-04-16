// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package visibilitygroup

// VisibilityModel - Output model with visibility properties.
type VisibilityModel struct {
	// REQUIRED; Required string[], illustrating a create property.
	CreateProp []*string

	// REQUIRED; Required bool, illustrating a delete property.
	DeleteProp *bool

	// REQUIRED; Required int32, illustrating a query property.
	QueryProp *int32

	// REQUIRED; Required string, illustrating a readonly property.
	ReadProp *string

	// REQUIRED; Required int32[], illustrating a update property.
	UpdateProp []*int32
}
