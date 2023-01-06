// Code generated by go-swagger; DO NOT EDIT.

package certificate_template

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

// PostCertificateTemplateQueryHandlerFunc turns a function with the right signature into a post certificate template query handler
type PostCertificateTemplateQueryHandlerFunc func(PostCertificateTemplateQueryParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostCertificateTemplateQueryHandlerFunc) Handle(params PostCertificateTemplateQueryParams) middleware.Responder {
	return fn(params)
}

// PostCertificateTemplateQueryHandler interface for that can handle valid post certificate template query params
type PostCertificateTemplateQueryHandler interface {
	Handle(PostCertificateTemplateQueryParams) middleware.Responder
}

// NewPostCertificateTemplateQuery creates a new http.Handler for the post certificate template query operation
func NewPostCertificateTemplateQuery(ctx *middleware.Context, handler PostCertificateTemplateQueryHandler) *PostCertificateTemplateQuery {
	return &PostCertificateTemplateQuery{Context: ctx, Handler: handler}
}

/*
	PostCertificateTemplateQuery swagger:route POST /certificate/template/query certificate_template postCertificateTemplateQuery

query certificate template by key
*/
type PostCertificateTemplateQuery struct {
	Context *middleware.Context
	Handler PostCertificateTemplateQueryHandler
}

func (o *PostCertificateTemplateQuery) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostCertificateTemplateQueryParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostCertificateTemplateQueryBody post certificate template query body
//
// swagger:model PostCertificateTemplateQueryBody
type PostCertificateTemplateQueryBody struct {

	// organization name
	OrganizationName string `json:"organization_name,omitempty"`

	// template key
	// Required: true
	TemplateKey *string `json:"template_key"`
}

// Validate validates this post certificate template query body
func (o *PostCertificateTemplateQueryBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateTemplateKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostCertificateTemplateQueryBody) validateTemplateKey(formats strfmt.Registry) error {

	if err := validate.Required("query_template"+"."+"template_key", "body", o.TemplateKey); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post certificate template query body based on context it is used
func (o *PostCertificateTemplateQueryBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostCertificateTemplateQueryBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostCertificateTemplateQueryBody) UnmarshalBinary(b []byte) error {
	var res PostCertificateTemplateQueryBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}