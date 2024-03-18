// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdataboxedge

import "encoding/json"

// MarshalJSON implements the json.Marshaller interface for type AddonsClientCreateOrUpdateResponse.
func (a AddonsClientCreateOrUpdateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.AddonClassification)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AddonsClientCreateOrUpdateResponse.
func (a *AddonsClientCreateOrUpdateResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalAddonClassification(data)
	if err != nil {
		return err
	}
	a.AddonClassification = res
	return nil
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AddonsClientGetResponse.
func (a *AddonsClientGetResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalAddonClassification(data)
	if err != nil {
		return err
	}
	a.AddonClassification = res
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RolesClientCreateOrUpdateResponse.
func (r RolesClientCreateOrUpdateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.RoleClassification)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RolesClientCreateOrUpdateResponse.
func (r *RolesClientCreateOrUpdateResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalRoleClassification(data)
	if err != nil {
		return err
	}
	r.RoleClassification = res
	return nil
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RolesClientGetResponse.
func (r *RolesClientGetResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalRoleClassification(data)
	if err != nil {
		return err
	}
	r.RoleClassification = res
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type TriggersClientCreateOrUpdateResponse.
func (t TriggersClientCreateOrUpdateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.TriggerClassification)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type TriggersClientCreateOrUpdateResponse.
func (t *TriggersClientCreateOrUpdateResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalTriggerClassification(data)
	if err != nil {
		return err
	}
	t.TriggerClassification = res
	return nil
}

// UnmarshalJSON implements the json.Unmarshaller interface for type TriggersClientGetResponse.
func (t *TriggersClientGetResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalTriggerClassification(data)
	if err != nil {
		return err
	}
	t.TriggerClassification = res
	return nil
}