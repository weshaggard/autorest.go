// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"encoding/json"
	"enumdiscgroup"
)

func unmarshalDogClassification(rawMsg json.RawMessage) (enumdiscgroup.DogClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b enumdiscgroup.DogClassification
	switch m["kind"] {
	case string(enumdiscgroup.DogKindGolden):
		b = &enumdiscgroup.Golden{}
	default:
		b = &enumdiscgroup.Dog{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalSnakeClassification(rawMsg json.RawMessage) (enumdiscgroup.SnakeClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b enumdiscgroup.SnakeClassification
	switch m["kind"] {
	case string(enumdiscgroup.SnakeKindCobra):
		b = &enumdiscgroup.Cobra{}
	default:
		b = &enumdiscgroup.Snake{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}