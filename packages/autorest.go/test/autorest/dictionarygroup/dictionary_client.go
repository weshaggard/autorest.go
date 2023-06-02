// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package dictionarygroup

import (
	"generatortests"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewDictionaryClient(options *azcore.ClientOptions) (*DictionaryClient, error) {
	client, err := azcore.NewClient("dictionarygroup.DictionaryClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &DictionaryClient{internal: client}, nil
}
