// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixBindUserRoleRequest openpitrix bind user role request
// swagger:model openpitrixBindUserRoleRequest
type OpenpitrixBindUserRoleRequest struct {

	// role id
	RoleID []string `json:"role_id"`

	// user id
	UserID []string `json:"user_id"`
}

// Validate validates this openpitrix bind user role request
func (m *OpenpitrixBindUserRoleRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRoleID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixBindUserRoleRequest) validateRoleID(formats strfmt.Registry) error {

	if swag.IsZero(m.RoleID) { // not required
		return nil
	}

	return nil
}

func (m *OpenpitrixBindUserRoleRequest) validateUserID(formats strfmt.Registry) error {

	if swag.IsZero(m.UserID) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixBindUserRoleRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixBindUserRoleRequest) UnmarshalBinary(b []byte) error {
	var res OpenpitrixBindUserRoleRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
