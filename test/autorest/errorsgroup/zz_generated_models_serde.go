//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package errorsgroup

import (
	"encoding/json"
	"fmt"
)

// UnmarshalJSON implements the json.Unmarshaller interface for type AnimalNotFound.
func (a *AnimalNotFound) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "name":
			err = unpopulate(val, "Name", &a.Name)
			delete(rawMsg, key)
		case "reason":
			err = unpopulate(val, "Reason", &a.Reason)
			delete(rawMsg, key)
		case "someBaseProp":
			err = unpopulate(val, "SomeBaseProp", &a.SomeBaseProp)
			delete(rawMsg, key)
		case "whatNotFound":
			err = unpopulate(val, "WhatNotFound", &a.WhatNotFound)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LinkNotFound.
func (l *LinkNotFound) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "reason":
			err = unpopulate(val, "Reason", &l.Reason)
			delete(rawMsg, key)
		case "someBaseProp":
			err = unpopulate(val, "SomeBaseProp", &l.SomeBaseProp)
			delete(rawMsg, key)
		case "whatNotFound":
			err = unpopulate(val, "WhatNotFound", &l.WhatNotFound)
			delete(rawMsg, key)
		case "whatSubAddress":
			err = unpopulate(val, "WhatSubAddress", &l.WhatSubAddress)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PetHungryOrThirstyError.
func (p *PetHungryOrThirstyError) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "actionResponse":
			err = unpopulate(val, "ActionResponse", &p.ActionResponse)
			delete(rawMsg, key)
		case "errorMessage":
			err = unpopulate(val, "ErrorMessage", &p.ErrorMessage)
			delete(rawMsg, key)
		case "errorType":
			err = unpopulate(val, "ErrorType", &p.ErrorType)
			delete(rawMsg, key)
		case "hungryOrThirsty":
			err = unpopulate(val, "HungryOrThirsty", &p.HungryOrThirsty)
			delete(rawMsg, key)
		case "reason":
			err = unpopulate(val, "Reason", &p.Reason)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PetSadError.
func (p *PetSadError) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "actionResponse":
			err = unpopulate(val, "ActionResponse", &p.ActionResponse)
			delete(rawMsg, key)
		case "errorMessage":
			err = unpopulate(val, "ErrorMessage", &p.ErrorMessage)
			delete(rawMsg, key)
		case "errorType":
			err = unpopulate(val, "ErrorType", &p.ErrorType)
			delete(rawMsg, key)
		case "reason":
			err = unpopulate(val, "Reason", &p.Reason)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

func unpopulate(data json.RawMessage, fn string, v interface{}) error {
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}
