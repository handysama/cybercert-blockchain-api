// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IssueStandardTokenPayload issue standard token payload
//
// swagger:model IssueStandardTokenPayload
type IssueStandardTokenPayload struct {

	// access quota
	AccessQuota int64 `json:"access_quota,omitempty"`

	// amount
	Amount int64 `json:"amount,omitempty"`

	// unix timestamp in second (UTC)
	ExpiryDate int64 `json:"expiry_date,omitempty"`

	// issuer token id
	IssuerTokenID string `json:"issuer_token_id,omitempty"`

	// recipient
	Recipient string `json:"recipient,omitempty"`

	// token id
	TokenID string `json:"token_id,omitempty"`
}

// Validate validates this issue standard token payload
func (m *IssueStandardTokenPayload) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this issue standard token payload based on context it is used
func (m *IssueStandardTokenPayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IssueStandardTokenPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IssueStandardTokenPayload) UnmarshalBinary(b []byte) error {
	var res IssueStandardTokenPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
