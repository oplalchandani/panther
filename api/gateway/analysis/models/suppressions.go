// Code generated by go-swagger; DO NOT EDIT.

package models

/**
 * Panther is a scalable, powerful, cloud-native SIEM written in Golang/React.
 * Copyright (C) 2020 Panther Labs Inc
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// Suppressions List of resource ID regexes that are excepted from this policy. The policy will still be evaluated, but failures will not trigger alerts nor remediations.
//
// swagger:model suppressions
type Suppressions []string

// Validate validates this suppressions
func (m Suppressions) Validate(formats strfmt.Registry) error {
	var res []error

	iSuppressionsSize := int64(len(m))

	if err := validate.MaxItems("", "body", iSuppressionsSize, 500); err != nil {
		return err
	}

	if err := validate.UniqueItems("", "body", m); err != nil {
		return err
	}

	for i := 0; i < len(m); i++ {

		if err := validate.MaxLength(strconv.Itoa(i), "body", string(m[i]), 1000); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}