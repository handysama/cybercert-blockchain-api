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

// PostCertificateInfoHistoryHandlerFunc turns a function with the right signature into a post certificate info history handler
type PostCertificateInfoHistoryHandlerFunc func(PostCertificateInfoHistoryParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostCertificateInfoHistoryHandlerFunc) Handle(params PostCertificateInfoHistoryParams) middleware.Responder {
	return fn(params)
}

// PostCertificateInfoHistoryHandler interface for that can handle valid post certificate info history params
type PostCertificateInfoHistoryHandler interface {
	Handle(PostCertificateInfoHistoryParams) middleware.Responder
}

// NewPostCertificateInfoHistory creates a new http.Handler for the post certificate info history operation
func NewPostCertificateInfoHistory(ctx *middleware.Context, handler PostCertificateInfoHistoryHandler) *PostCertificateInfoHistory {
	return &PostCertificateInfoHistory{Context: ctx, Handler: handler}
}

/*
	PostCertificateInfoHistory swagger:route POST /certificate/info/history certificate_info postCertificateInfoHistory

PostCertificateInfoHistory post certificate info history API
*/
type PostCertificateInfoHistory struct {
	Context *middleware.Context
	Handler PostCertificateInfoHistoryHandler
}

func (o *PostCertificateInfoHistory) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostCertificateInfoHistoryParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostCertificateInfoHistoryBody post certificate info history body
//
// swagger:model PostCertificateInfoHistoryBody
type PostCertificateInfoHistoryBody struct {

	// certificate id
	// Required: true
	CertificateID *string `json:"certificate_id"`

	// organization name
	OrganizationName string `json:"organization_name,omitempty"`
}

// Validate validates this post certificate info history body
func (o *PostCertificateInfoHistoryBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCertificateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostCertificateInfoHistoryBody) validateCertificateID(formats strfmt.Registry) error {

	if err := validate.Required("get_history"+"."+"certificate_id", "body", o.CertificateID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post certificate info history body based on context it is used
func (o *PostCertificateInfoHistoryBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostCertificateInfoHistoryBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostCertificateInfoHistoryBody) UnmarshalBinary(b []byte) error {
	var res PostCertificateInfoHistoryBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}