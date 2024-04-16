// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package nullablegroup

import "time"

// BytesProperty - Template type for testing models with nullable property. Pass in the type of the property you are looking
// for
type BytesProperty struct {
	// REQUIRED; Property
	NullableProperty []byte

	// REQUIRED; Required property
	RequiredProperty *string
}

// CollectionsByteProperty - Model with collection bytes properties
type CollectionsByteProperty struct {
	// REQUIRED; Property
	NullableProperty [][]byte

	// REQUIRED; Required property
	RequiredProperty *string
}

// CollectionsModelProperty - Model with collection models properties
type CollectionsModelProperty struct {
	// REQUIRED; Property
	NullableProperty []*InnerModel

	// REQUIRED; Required property
	RequiredProperty *string
}

// DatetimeProperty - Model with a datetime property
type DatetimeProperty struct {
	// REQUIRED; Property
	NullableProperty *time.Time

	// REQUIRED; Required property
	RequiredProperty *string
}

// DurationProperty - Model with a duration property
type DurationProperty struct {
	// REQUIRED; Property
	NullableProperty *string

	// REQUIRED; Required property
	RequiredProperty *string
}

// InnerModel - Inner model used in collections model property
type InnerModel struct {
	// REQUIRED; Inner model property
	Property *string
}

// StringProperty - Template type for testing models with nullable property. Pass in the type of the property you are looking
// for
type StringProperty struct {
	// REQUIRED; Property
	NullableProperty *string

	// REQUIRED; Required property
	RequiredProperty *string
}
