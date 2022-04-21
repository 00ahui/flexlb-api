// Copyright (c) 2022 Yaohui Wang (yaohuiwang@outlook.com)
// FlexLB is licensed under Mulan PubL v2.
// You can use this software according to the terms and conditions of the Mulan PubL v2.
// You may obtain a copy of Mulan PubL v2 at:
//         http://license.coscl.org.cn/MulanPubL-2.0
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
// EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
// MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PubL v2 for more details.

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Instance Instance
//
// swagger:model Instance
type Instance struct {

	// config
	// Required: true
	Config *InstanceConfig `json:"config"`

	// Instance ID
	// Required: true
	// Maximum: 255
	// Minimum: 0
	ID uint8 `json:"id"`

	// Last modified time
	// Minimum: 0
	LastModified int64 `json:"last_modified,omitempty"`

	// Instance status
	// Required: true
	Status map[string]string `json:"status"`
}

// Validate validates this instance
func (m *Instance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Instance) validateConfig(formats strfmt.Registry) error {

	if err := validate.Required("config", "body", m.Config); err != nil {
		return err
	}

	if m.Config != nil {
		if err := m.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

func (m *Instance) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", uint8(m.ID)); err != nil {
		return err
	}

	if err := validate.MinimumUint("id", "body", uint64(m.ID), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumUint("id", "body", uint64(m.ID), 255, false); err != nil {
		return err
	}

	return nil
}

func (m *Instance) validateLastModified(formats strfmt.Registry) error {
	if swag.IsZero(m.LastModified) { // not required
		return nil
	}

	if err := validate.MinimumInt("last_modified", "body", m.LastModified, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Instance) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this instance based on the context it is used
func (m *Instance) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Instance) contextValidateConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.Config != nil {
		if err := m.Config.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Instance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Instance) UnmarshalBinary(b []byte) error {
	var res Instance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
