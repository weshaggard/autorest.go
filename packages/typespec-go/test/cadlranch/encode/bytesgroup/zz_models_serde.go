//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package bytesgroup

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// MarshalJSON implements the json.Marshaller interface for type Base64BytesProperty.
func (b Base64BytesProperty) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populateByteArray(objectMap, "value", b.Value, func() any {
		return runtime.EncodeByteArray(b.Value, runtime.Base64StdFormat)
	})
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Base64BytesProperty.
func (b *Base64BytesProperty) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", b, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "value":
			if val != nil && string(val) != "null" {
				err = runtime.DecodeByteArray(string(val), &b.Value, runtime.Base64StdFormat)
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", b, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Base64URLArrayBytesProperty.
func (b Base64URLArrayBytesProperty) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populateByteArray(objectMap, "value", b.Value, func() any {
		encodedValue := make([]string, len(b.Value))
		for i := 0; i < len(b.Value); i++ {
			encodedValue[i] = runtime.EncodeByteArray(b.Value[i], runtime.Base64URLFormat)
		}
		return encodedValue
	})
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Base64URLArrayBytesProperty.
func (b *Base64URLArrayBytesProperty) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", b, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "value":
			var encodedValue []string
			err = unpopulate(val, "Value", &encodedValue)
			if err == nil && len(encodedValue) > 0 {
				b.Value = make([][]byte, len(encodedValue))
				for i := 0; i < len(encodedValue) && err == nil; i++ {
					err = runtime.DecodeByteArray(encodedValue[i], &b.Value[i], runtime.Base64URLFormat)
				}
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", b, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Base64URLBytesProperty.
func (b Base64URLBytesProperty) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populateByteArray(objectMap, "value", b.Value, func() any {
		return runtime.EncodeByteArray(b.Value, runtime.Base64URLFormat)
	})
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Base64URLBytesProperty.
func (b *Base64URLBytesProperty) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", b, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "value":
			if val != nil && string(val) != "null" {
				err = runtime.DecodeByteArray(string(val), &b.Value, runtime.Base64URLFormat)
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", b, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DefaultBytesProperty.
func (d DefaultBytesProperty) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populateByteArray(objectMap, "value", d.Value, func() any {
		return runtime.EncodeByteArray(d.Value, runtime.Base64StdFormat)
	})
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DefaultBytesProperty.
func (d *DefaultBytesProperty) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "value":
			if val != nil && string(val) != "null" {
				err = runtime.DecodeByteArray(string(val), &d.Value, runtime.Base64StdFormat)
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

func populateByteArray[T any](m map[string]any, k string, b []T, convert func() any) {
	if azcore.IsNullValue(b) {
		m[k] = nil
	} else if len(b) == 0 {
		return
	} else {
		m[k] = convert()
	}
}

func unpopulate(data json.RawMessage, fn string, v any) error {
	if data == nil || string(data) == "null" {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}