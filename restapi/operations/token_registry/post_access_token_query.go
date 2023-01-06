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
	"github.com/go-openapi/validate"
)

// PostAccessTokenQueryHandlerFunc turns a function with the right signature into a post access token query handler
type PostAccessTokenQueryHandlerFunc func(PostAccessTokenQueryParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostAccessTokenQueryHandlerFunc) Handle(params PostAccessTokenQueryParams) middleware.Responder {
	return fn(params)
}

// PostAccessTokenQueryHandler interface for that can handle valid post access token query params
type PostAccessTokenQueryHandler interface {
	Handle(PostAccessTokenQueryParams) middleware.Responder
}

// NewPostAccessTokenQuery creates a new http.Handler for the post access token query operation
func NewPostAccessTokenQuery(ctx *middleware.Context, handler PostAccessTokenQueryHandler) *PostAccessTokenQuery {
	return &PostAccessTokenQuery{Context: ctx, Handler: handler}
}

/*
	PostAccessTokenQuery swagger:route POST /access-token/query token_registry postAccessTokenQuery

PostAccessTokenQuery post access token query API
*/
type PostAccessTokenQuery struct {
	Context *middleware.Context
	Handler PostAccessTokenQueryHandler
}

func (o *PostAccessTokenQuery) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostAccessTokenQueryParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostAccessTokenQueryBody post access token query body
//
// swagger:model PostAccessTokenQueryBody
type PostAccessTokenQueryBody struct {

	// registered organization name on blockchain
	OrganizationName string `json:"organization_name,omitempty"`

	// token id
	// Required: true
	TokenID *string `json:"token_id"`
}

// Validate validates this post access token query body
func (o *PostAccessTokenQueryBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateTokenID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostAccessTokenQueryBody) validateTokenID(formats strfmt.Registry) error {

	if err := validate.Required("query_record"+"."+"token_id", "body", o.TokenID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post access token query body based on context it is used
func (o *PostAccessTokenQueryBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostAccessTokenQueryBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostAccessTokenQueryBody) UnmarshalBinary(b []byte) error {
	var res PostAccessTokenQueryBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}