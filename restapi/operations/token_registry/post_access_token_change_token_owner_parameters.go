// Code generated by go-swagger; DO NOT EDIT.

package token_registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"
)

// NewPostAccessTokenChangeTokenOwnerParams creates a new PostAccessTokenChangeTokenOwnerParams object
//
// There are no default values defined in the spec.
func NewPostAccessTokenChangeTokenOwnerParams() PostAccessTokenChangeTokenOwnerParams {

	return PostAccessTokenChangeTokenOwnerParams{}
}

// PostAccessTokenChangeTokenOwnerParams contains all the bound params for the post access token change token owner operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostAccessTokenChangeTokenOwner
type PostAccessTokenChangeTokenOwnerParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: body
	*/
	TokenRecord PostAccessTokenChangeTokenOwnerBody
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostAccessTokenChangeTokenOwnerParams() beforehand.
func (o *PostAccessTokenChangeTokenOwnerParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body PostAccessTokenChangeTokenOwnerBody
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("tokenRecord", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(r.Context())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.TokenRecord = body
			}
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
