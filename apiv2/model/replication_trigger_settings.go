// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ReplicationTriggerSettings replication trigger settings
//
// swagger:model ReplicationTriggerSettings
type ReplicationTriggerSettings struct {

	// The cron string for scheduled trigger
	Cron string `json:"cron,omitempty"`
}

// Validate validates this replication trigger settings
func (m *ReplicationTriggerSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReplicationTriggerSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReplicationTriggerSettings) UnmarshalBinary(b []byte) error {
	var res ReplicationTriggerSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
