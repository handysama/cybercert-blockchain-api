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

// NewPostAccessTokenIssueRootParams creates a new PostAccessTokenIssueRootParams object
//
// There are no default values defined in the spec.
func NewPostAccessTokenIssueRootParams() PostAccessTokenIssueRootParams {

	return PostAccessTokenIssueRootParams{}
}

// PostAccessTokenIssueRootParams contains all the bound params for the post access token issue root operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostAccessTokenIssueRoot
type PostAccessTokenIssueRootParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: body
	*/
	TokenRecord PostAccessTokenIssueRootBody
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostAccessTokenIssueRootParams() beforehand.
func (o *PostAccessTokenIssueRootParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body PostAccessTokenIssueRootBody
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
