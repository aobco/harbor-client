// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProjectSummary project summary
//
// swagger:model ProjectSummary
type ProjectSummary struct {

	// The total number of charts under this project.
	ChartCount int64 `json:"chart_count"`

	// The total number of developer members.
	DeveloperCount int64 `json:"developer_count,omitempty"`

	// The total number of guest members.
	GuestCount int64 `json:"guest_count,omitempty"`

	// The total number of limited guest members.
	LimitedGuestCount int64 `json:"limited_guest_count,omitempty"`

	// The total number of maintainer members.
	MaintainerCount int64 `json:"maintainer_count,omitempty"`

	// The total number of project admin members.
	ProjectAdminCount int64 `json:"project_admin_count,omitempty"`

	// quota
	Quota *ProjectSummaryQuota `json:"quota,omitempty"`

	// registry
	Registry *Registry `json:"registry,omitempty"`

	// The number of the repositories under this project.
	RepoCount int64 `json:"repo_count"`
}

// Validate validates this project summary
func (m *ProjectSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQuota(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegistry(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectSummary) validateQuota(formats strfmt.Registry) error {

	if swag.IsZero(m.Quota) { // not required
		return nil
	}

	if m.Quota != nil {
		if err := m.Quota.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quota")
			}
			return err
		}
	}

	return nil
}

func (m *ProjectSummary) validateRegistry(formats strfmt.Registry) error {

	if swag.IsZero(m.Registry) { // not required
		return nil
	}

	if m.Registry != nil {
		if err := m.Registry.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("registry")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProjectSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectSummary) UnmarshalBinary(b []byte) error {
	var res ProjectSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
