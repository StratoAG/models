// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BackendBalance backend balance
// swagger:model backendBalance
type BackendBalance struct {

	// algorithm
	// Enum: [roundrobin static-rr leastconn first source uri url_param random]
	Algorithm string `json:"algorithm,omitempty"`

	// arguments
	Arguments []string `json:"arguments"`
}

// Validate validates this backend balance
func (m *BackendBalance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlgorithm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateArguments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var backendBalanceTypeAlgorithmPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["roundrobin","static-rr","leastconn","first","source","uri","url_param","random"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendBalanceTypeAlgorithmPropEnum = append(backendBalanceTypeAlgorithmPropEnum, v)
	}
}

const (

	// BackendBalanceAlgorithmRoundrobin captures enum value "roundrobin"
	BackendBalanceAlgorithmRoundrobin string = "roundrobin"

	// BackendBalanceAlgorithmStaticRr captures enum value "static-rr"
	BackendBalanceAlgorithmStaticRr string = "static-rr"

	// BackendBalanceAlgorithmLeastconn captures enum value "leastconn"
	BackendBalanceAlgorithmLeastconn string = "leastconn"

	// BackendBalanceAlgorithmFirst captures enum value "first"
	BackendBalanceAlgorithmFirst string = "first"

	// BackendBalanceAlgorithmSource captures enum value "source"
	BackendBalanceAlgorithmSource string = "source"

	// BackendBalanceAlgorithmURI captures enum value "uri"
	BackendBalanceAlgorithmURI string = "uri"

	// BackendBalanceAlgorithmURLParam captures enum value "url_param"
	BackendBalanceAlgorithmURLParam string = "url_param"

	// BackendBalanceAlgorithmRandom captures enum value "random"
	BackendBalanceAlgorithmRandom string = "random"
)

// prop value enum
func (m *BackendBalance) validateAlgorithmEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendBalanceTypeAlgorithmPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *BackendBalance) validateAlgorithm(formats strfmt.Registry) error {

	if swag.IsZero(m.Algorithm) { // not required
		return nil
	}

	// value enum
	if err := m.validateAlgorithmEnum("algorithm", "body", m.Algorithm); err != nil {
		return err
	}

	return nil
}

func (m *BackendBalance) validateArguments(formats strfmt.Registry) error {

	if swag.IsZero(m.Arguments) { // not required
		return nil
	}

	for i := 0; i < len(m.Arguments); i++ {

		if err := validate.Pattern("arguments"+"."+strconv.Itoa(i), "body", string(m.Arguments[i]), `^[^\s]+$`); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BackendBalance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackendBalance) UnmarshalBinary(b []byte) error {
	var res BackendBalance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
