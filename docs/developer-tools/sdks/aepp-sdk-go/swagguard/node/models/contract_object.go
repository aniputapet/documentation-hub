// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ContractObject contract object
// swagger:model ContractObject
type ContractObject struct {

	// abi version
	// Required: true
	AbiVersion *uint16 `json:"abi_version"`

	// active
	// Required: true
	Active *bool `json:"active"`

	// deposit
	// Required: true
	Deposit *uint64 `json:"deposit"`

	// id
	// Required: true
	ID *string `json:"id"`

	// owner id
	// Required: true
	OwnerID *string `json:"owner_id"`

	// referrer ids
	// Required: true
	ReferrerIds []string `json:"referrer_ids"`

	// vm version
	// Required: true
	VMVersion *uint16 `json:"vm_version"`
}

// Validate validates this contract object
func (m *ContractObject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAbiVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateActive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeposit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwnerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReferrerIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContractObject) validateAbiVersion(formats strfmt.Registry) error {

	if err := validate.Required("abi_version", "body", m.AbiVersion); err != nil {
		return err
	}

	return nil
}

func (m *ContractObject) validateActive(formats strfmt.Registry) error {

	if err := validate.Required("active", "body", m.Active); err != nil {
		return err
	}

	return nil
}

func (m *ContractObject) validateDeposit(formats strfmt.Registry) error {

	if err := validate.Required("deposit", "body", m.Deposit); err != nil {
		return err
	}

	return nil
}

func (m *ContractObject) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ContractObject) validateOwnerID(formats strfmt.Registry) error {

	if err := validate.Required("owner_id", "body", m.OwnerID); err != nil {
		return err
	}

	return nil
}

func (m *ContractObject) validateReferrerIds(formats strfmt.Registry) error {

	if err := validate.Required("referrer_ids", "body", m.ReferrerIds); err != nil {
		return err
	}

	return nil
}

func (m *ContractObject) validateVMVersion(formats strfmt.Registry) error {

	if err := validate.Required("vm_version", "body", m.VMVersion); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContractObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContractObject) UnmarshalBinary(b []byte) error {
	var res ContractObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}