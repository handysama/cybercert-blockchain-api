// Code generated by go-swagger; DO NOT EDIT.

package certificate_info

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DeleteCertificateInfoRevokeHandlerFunc turns a function with the right signature into a delete certificate info revoke handler
type DeleteCertificateInfoRevokeHandlerFunc func(DeleteCertificateInfoRevokeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteCertificateInfoRevokeHandlerFunc) Handle(params DeleteCertificateInfoRevokeParams) middleware.Responder {
	return fn(params)
}

// DeleteCertificateInfoRevokeHandler interface for that can handle valid delete certificate info revoke params
type DeleteCertificateInfoRevokeHandler interface {
	Handle(DeleteCertificateInfoRevokeParams) middleware.Responder
}

// NewDeleteCertificateInfoRevoke creates a new http.Handler for the delete certificate info revoke operation
func NewDeleteCertificateInfoRevoke(ctx *middleware.Context, handler DeleteCertificateInfoRevokeHandler) *DeleteCertificateInfoRevoke {
	return &DeleteCertificateInfoRevoke{Context: ctx, Handler: handler}
}

/*
	DeleteCertificateInfoRevoke swagger:route DELETE /certificate/info/revoke certificate_info deleteCertificateInfoRevoke

DeleteCertificateInfoRevoke delete certificate info revoke API
*/
type DeleteCertificateInfoRevoke struct {
	Context *middleware.Context
	Handler DeleteCertificateInfoRevokeHandler
}

func (o *DeleteCertificateInfoRevoke) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteCertificateInfoRevokeParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// DeleteCertificateInfoRevokeBody delete certificate info revoke body
//
// swagger:model DeleteCertificateInfoRevokeBody
type DeleteCertificateInfoRevokeBody struct {

	// certificate id
	// Required: true
	CertificateID *string `json:"certificate_id"`

	// organization name
	OrganizationName string `json:"organization_name,omitempty"`
}

// Validate validates this delete certificate info revoke body
func (o *DeleteCertificateInfoRevokeBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCertificateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteCertificateInfoRevokeBody) validateCertificateID(formats strfmt.Registry) error {

	if err := validate.Required("revoke_certificate"+"."+"certificate_id", "body", o.CertificateID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this delete certificate info revoke body based on context it is used
func (o *DeleteCertificateInfoRevokeBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteCertificateInfoRevokeBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteCertificateInfoRevokeBody) UnmarshalBinary(b []byte) error {
	var res DeleteCertificateInfoRevokeBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}