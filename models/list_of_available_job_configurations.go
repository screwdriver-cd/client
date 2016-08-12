package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*ListOfAvailableJobConfigurations list of available job configurations

swagger:model List of available job configurations
*/
type ListOfAvailableJobConfigurations struct {

	/* main

	Required: true
	*/
	Main []string `json:"main"`
}

// Validate validates this list of available job configurations
func (m *ListOfAvailableJobConfigurations) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMain(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListOfAvailableJobConfigurations) validateMain(formats strfmt.Registry) error {

	if err := validate.Required("main", "body", m.Main); err != nil {
		return err
	}

	for i := 0; i < len(m.Main); i++ {

		if err := validate.RequiredString("main"+"."+strconv.Itoa(i), "body", string(m.Main[i])); err != nil {
			return err
		}

	}

	return nil
}
