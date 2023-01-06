// Code generated by go-swagger; DO NOT EDIT.

package token_registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"cybercert-blockchain-api/models"
)

// PostAccessTokenChangeTokenOwnerOKCode is the HTTP code returned for type PostAccessTokenChangeTokenOwnerOK
const PostAccessTokenChangeTokenOwnerOKCode int = 200

/*
PostAccessTokenChangeTokenOwnerOK OK

swagger:response postAccessTokenChangeTokenOwnerOK
*/
type PostAccessTokenChangeTokenOwnerOK struct {

	/*
	  In: Body
	*/
	Payload *models.CreatedResponse `json:"body,omitempty"`
}

// NewPostAccessTokenChangeTokenOwnerOK creates PostAccessTokenChangeTokenOwnerOK with default headers values
func NewPostAccessTokenChangeTokenOwnerOK() *PostAccessTokenChangeTokenOwnerOK {

	return &PostAccessTokenChangeTokenOwnerOK{}
}

// WithPayload adds the payload to the post access token change token owner o k response
func (o *PostAccessTokenChangeTokenOwnerOK) WithPayload(payload *models.CreatedResponse) *PostAccessTokenChangeTokenOwnerOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post access token change token owner o k response
func (o *PostAccessTokenChangeTokenOwnerOK) SetPayload(payload *models.CreatedResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostAccessTokenChangeTokenOwnerOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostAccessTokenChangeTokenOwnerUnauthorizedCode is the HTTP code returned for type PostAccessTokenChangeTokenOwnerUnauthorized
const PostAccessTokenChangeTokenOwnerUnauthorizedCode int = 401

/*
PostAccessTokenChangeTokenOwnerUnauthorized Unauthorized

swagger:response postAccessTokenChangeTokenOwnerUnauthorized
*/
type PostAccessTokenChangeTokenOwnerUnauthorized struct {
}

// NewPostAccessTokenChangeTokenOwnerUnauthorized creates PostAccessTokenChangeTokenOwnerUnauthorized with default headers values
func NewPostAccessTokenChangeTokenOwnerUnauthorized() *PostAccessTokenChangeTokenOwnerUnauthorized {

	return &PostAccessTokenChangeTokenOwnerUnauthorized{}
}

// WriteResponse to the client
func (o *PostAccessTokenChangeTokenOwnerUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// PostAccessTokenChangeTokenOwnerForbiddenCode is the HTTP code returned for type PostAccessTokenChangeTokenOwnerForbidden
const PostAccessTokenChangeTokenOwnerForbiddenCode int = 403

/*
PostAccessTokenChangeTokenOwnerForbidden Forbidden

swagger:response postAccessTokenChangeTokenOwnerForbidden
*/
type PostAccessTokenChangeTokenOwnerForbidden struct {
}

// NewPostAccessTokenChangeTokenOwnerForbidden creates PostAccessTokenChangeTokenOwnerForbidden with default headers values
func NewPostAccessTokenChangeTokenOwnerForbidden() *PostAccessTokenChangeTokenOwnerForbidden {

	return &PostAccessTokenChangeTokenOwnerForbidden{}
}

// WriteResponse to the client
func (o *PostAccessTokenChangeTokenOwnerForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// PostAccessTokenChangeTokenOwnerInternalServerErrorCode is the HTTP code returned for type PostAccessTokenChangeTokenOwnerInternalServerError
const PostAccessTokenChangeTokenOwnerInternalServerErrorCode int = 500

/*
PostAccessTokenChangeTokenOwnerInternalServerError Internal Server Error

swagger:response postAccessTokenChangeTokenOwnerInternalServerError
*/
type PostAccessTokenChangeTokenOwnerInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewPostAccessTokenChangeTokenOwnerInternalServerError creates PostAccessTokenChangeTokenOwnerInternalServerError with default headers values
func NewPostAccessTokenChangeTokenOwnerInternalServerError() *PostAccessTokenChangeTokenOwnerInternalServerError {

	return &PostAccessTokenChangeTokenOwnerInternalServerError{}
}

// WithPayload adds the payload to the post access token change token owner internal server error response
func (o *PostAccessTokenChangeTokenOwnerInternalServerError) WithPayload(payload *models.APIResponse) *PostAccessTokenChangeTokenOwnerInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post access token change token owner internal server error response
func (o *PostAccessTokenChangeTokenOwnerInternalServerError) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostAccessTokenChangeTokenOwnerInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}