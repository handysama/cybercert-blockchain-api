// Code generated by go-swagger; DO NOT EDIT.

package token_registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"cybercert-blockchain-api/models"
)

// PostAccessTokenIssueStandardHandlerFunc turns a function with the right signature into a post access token issue standard handler
type PostAccessTokenIssueStandardHandlerFunc func(PostAccessTokenIssueStandardParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostAccessTokenIssueStandardHandlerFunc) Handle(params PostAccessTokenIssueStandardParams) middleware.Responder {
	return fn(params)
}

// PostAccessTokenIssueStandardHandler interface for that can handle valid post access token issue standard params
type PostAccessTokenIssueStandardHandler interface {
	Handle(PostAccessTokenIssueStandardParams) middleware.Responder
}

// NewPostAccessTokenIssueStandard creates a new http.Handler for the post access token issue standard operation
func NewPostAccessTokenIssueStandard(ctx *middleware.Context, handler PostAccessTokenIssueStandardHandler) *PostAccessTokenIssueStandard {
	return &PostAccessTokenIssueStandard{Context: ctx, Handler: handler}
}

/*
	PostAccessTokenIssueStandard swagger:route POST /access-token/issue/standard token_registry postAccessTokenIssueStandard

PostAccessTokenIssueStandard post access token issue standard API
*/
type PostAccessTokenIssueStandard struct {
	Context *middleware.Context
	Handler PostAccessTokenIssueStandardHandler
}

func (o *PostAccessTokenIssueStandard) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostAccessTokenIssueStandardParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostAccessTokenIssueStandardBody post access token issue standard body
//
// swagger:model PostAccessTokenIssueStandardBody
type PostAccessTokenIssueStandardBody struct {

	// args
	Args *models.IssueStandardTokenPayload `json:"args,omitempty"`

	// registered organization name on blockchain
	OrganizationName string `json:"organization_name,omitempty"`
}

// Validate validates this post access token issue standard body
func (o *PostAccessTokenIssueStandardBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateArgs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostAccessTokenIssueStandardBody) validateArgs(formats strfmt.Registry) error {
	if swag.IsZero(o.Args) { // not required
		return nil
	}

	if o.Args != nil {
		if err := o.Args.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("token_record" + "." + "args")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("token_record" + "." + "args")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post access token issue standard body based on the context it is used
func (o *PostAccessTokenIssueStandardBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateArgs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostAccessTokenIssueStandardBody) contextValidateArgs(ctx context.Context, formats strfmt.Registry) error {

	if o.Args != nil {
		if err := o.Args.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("token_record" + "." + "args")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("token_record" + "." + "args")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostAccessTokenIssueStandardBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostAccessTokenIssueStandardBody) UnmarshalBinary(b []byte) error {
	var res PostAccessTokenIssueStandardBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
