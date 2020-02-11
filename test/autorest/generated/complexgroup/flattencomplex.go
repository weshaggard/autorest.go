// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package complexgroup

import (
	"context"
	azinternal "generatortests/autorest/generated/complexgroup/internal/complexgroup"
)

// FlattencomplexOperations contains the methods for the Flattencomplex group.
type FlattencomplexOperations interface {
	GetValid(ctx context.Context) (*FlattencomplexGetValidResponse, error)
}

type flattencomplexOperations struct {
	*Client
	azinternal.FlattencomplexOperations
}

func (client *flattencomplexOperations) GetValid(ctx context.Context) (*FlattencomplexGetValidResponse, error) {
	req, err := client.GetValidCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetValidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

var _ FlattencomplexOperations = (*flattencomplexOperations)(nil)