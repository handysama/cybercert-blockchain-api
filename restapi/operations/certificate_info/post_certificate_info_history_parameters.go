// Code generated by go-swagger; DO NOT EDIT.

package certificate_info

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"
)

// NewPostCertificateInfoHistoryParams creates a new PostCertificateInfoHistoryParams object
//
// There are no default values defined in the spec.
func NewPostCertificateInfoHistoryParams() PostCertificateInfoHistoryParams {

	return PostCertificateInfoHistoryParams{}
}

// PostCertificateInfoHistoryParams contains all the bound params for the post certificate info history operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostCertificateInfoHistory
type PostCertificateInfoHistoryParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: body
	*/
	GetHistory PostCertificateInfoHistoryBody
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostCertificateInfoHistoryParams() beforehand.
func (o *PostCertificateInfoHistoryParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body PostCertificateInfoHistoryBody
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("getHistory", "body", "", err))
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
				o.GetHistory = body
			}
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}