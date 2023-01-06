// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IssueRootTokenPayload issue root token payload
//
// swagger:model IssueRootTokenPayload
type IssueRootTokenPayload struct {

	// certificate id
	CertificateID string `json:"certificate_id,omitempty"`

	// owner
	Owner string `json:"owner,omitempty"`

	// token id
	TokenID string `json:"token_id,omitempty"`
}

// Validate validates this issue root token payload
func (m *IssueRootTokenPayload) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this issue root token payload based on context it is used
func (m *IssueRootTokenPayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IssueRootTokenPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IssueRootTokenPayload) UnmarshalBinary(b []byte) error {
	var res IssueRootTokenPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
