// Code generated by go-swagger; DO NOT EDIT.

// /*
// Copyright The Rekor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */
//

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

// TufSignatureV001Schema TUF signature v0.0.1 Schema
//
// Schema for tuf signatures
//
// swagger:model tufSignatureV001Schema
type TufSignatureV001Schema struct {

	// Key identifier
	// Required: true
	Keyid *string `json:"keyid"`

	// Signature bytes
	// Required: true
	Sig *string `json:"sig"`
}

// Validate validates this tuf signature v001 schema
func (m *TufSignatureV001Schema) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKeyid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TufSignatureV001Schema) validateKeyid(formats strfmt.Registry) error {

	if err := validate.Required("keyid", "body", m.Keyid); err != nil {
		return err
	}

	return nil
}

func (m *TufSignatureV001Schema) validateSig(formats strfmt.Registry) error {

	if err := validate.Required("sig", "body", m.Sig); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this tuf signature v001 schema based on context it is used
func (m *TufSignatureV001Schema) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TufSignatureV001Schema) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TufSignatureV001Schema) UnmarshalBinary(b []byte) error {
	var res TufSignatureV001Schema
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}