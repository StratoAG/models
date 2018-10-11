// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Server Server
//
// HAProxy backend server configuration
// swagger:model server
type Server struct {

	// address
	Address string `json:"address,omitempty"`

	// check
	// Enum: [enabled]
	Check string `json:"check,omitempty"`

	// http cookie id
	HTTPCookieID string `json:"http-cookie-id,omitempty"`

	// maintenance
	// Enum: [enabled]
	Maintenance string `json:"maintenance,omitempty"`

	// max connections
	MaxConnections *int64 `json:"max-connections,omitempty"`

	// name
	// Required: true
	Name string `json:"name"`

	// port
	// Maximum: 65535
	// Minimum: 0
	Port *int64 `json:"port,omitempty"`

	// sorry
	// Enum: [enabled]
	Sorry string `json:"sorry,omitempty"`

	// ssl
	// Enum: [enabled]
	Ssl string `json:"ssl,omitempty"`

	// ssl cafile
	SslCafile string `json:"ssl_cafile,omitempty"`

	// ssl certificate
	SslCertificate string `json:"ssl_certificate,omitempty"`

	// weight
	Weight *int64 `json:"weight,omitempty"`
}

// Validate validates this server
func (m *Server) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCheck(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaintenance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSorry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSsl(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var serverTypeCheckPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serverTypeCheckPropEnum = append(serverTypeCheckPropEnum, v)
	}
}

const (

	// ServerCheckEnabled captures enum value "enabled"
	ServerCheckEnabled string = "enabled"
)

// prop value enum
func (m *Server) validateCheckEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, serverTypeCheckPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Server) validateCheck(formats strfmt.Registry) error {

	if swag.IsZero(m.Check) { // not required
		return nil
	}

	// value enum
	if err := m.validateCheckEnum("check", "body", m.Check); err != nil {
		return err
	}

	return nil
}

var serverTypeMaintenancePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serverTypeMaintenancePropEnum = append(serverTypeMaintenancePropEnum, v)
	}
}

const (

	// ServerMaintenanceEnabled captures enum value "enabled"
	ServerMaintenanceEnabled string = "enabled"
)

// prop value enum
func (m *Server) validateMaintenanceEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, serverTypeMaintenancePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Server) validateMaintenance(formats strfmt.Registry) error {

	if swag.IsZero(m.Maintenance) { // not required
		return nil
	}

	// value enum
	if err := m.validateMaintenanceEnum("maintenance", "body", m.Maintenance); err != nil {
		return err
	}

	return nil
}

func (m *Server) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *Server) validatePort(formats strfmt.Registry) error {

	if swag.IsZero(m.Port) { // not required
		return nil
	}

	if err := validate.MinimumInt("port", "body", int64(*m.Port), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("port", "body", int64(*m.Port), 65535, false); err != nil {
		return err
	}

	return nil
}

var serverTypeSorryPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serverTypeSorryPropEnum = append(serverTypeSorryPropEnum, v)
	}
}

const (

	// ServerSorryEnabled captures enum value "enabled"
	ServerSorryEnabled string = "enabled"
)

// prop value enum
func (m *Server) validateSorryEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, serverTypeSorryPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Server) validateSorry(formats strfmt.Registry) error {

	if swag.IsZero(m.Sorry) { // not required
		return nil
	}

	// value enum
	if err := m.validateSorryEnum("sorry", "body", m.Sorry); err != nil {
		return err
	}

	return nil
}

var serverTypeSslPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serverTypeSslPropEnum = append(serverTypeSslPropEnum, v)
	}
}

const (

	// ServerSslEnabled captures enum value "enabled"
	ServerSslEnabled string = "enabled"
)

// prop value enum
func (m *Server) validateSslEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, serverTypeSslPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Server) validateSsl(formats strfmt.Registry) error {

	if swag.IsZero(m.Ssl) { // not required
		return nil
	}

	// value enum
	if err := m.validateSslEnum("ssl", "body", m.Ssl); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Server) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Server) UnmarshalBinary(b []byte) error {
	var res Server
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
