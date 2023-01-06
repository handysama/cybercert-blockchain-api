// Code generated by go-swagger; DO NOT EDIT.

package token_registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"cybercert-blockchain-api/models"
)

// PostAccessTokenQueryOKCode is the HTTP code returned for type PostAccessTokenQueryOK
const PostAccessTokenQueryOKCode int = 200

/*
PostAccessTokenQueryOK OK

swagger:response postAccessTokenQueryOK
*/
type PostAccessTokenQueryOK struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewPostAccessTokenQueryOK creates PostAccessTokenQueryOK with default headers values
func NewPostAccessTokenQueryOK() *PostAccessTokenQueryOK {

	return &PostAccessTokenQueryOK{}
}

// WithPayload adds the payload to the post access token query o k response
func (o *PostAccessTokenQueryOK) WithPayload(payload interface{}) *PostAccessTokenQueryOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post access token query o k response
func (o *PostAccessTokenQueryOK) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostAccessTokenQueryOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PostAccessTokenQueryUnauthorizedCode is the HTTP code returned for type PostAccessTokenQueryUnauthorized
const PostAccessTokenQueryUnauthorizedCode int = 401

/*
PostAccessTokenQueryUnauthorized Unauthorized

swagger:response postAccessTokenQueryUnauthorized
*/
type PostAccessTokenQueryUnauthorized struct {
}

// NewPostAccessTokenQueryUnauthorized creates PostAccessTokenQueryUnauthorized with default headers values
func NewPostAccessTokenQueryUnauthorized() *PostAccessTokenQueryUnauthorized {

	return &PostAccessTokenQueryUnauthorized{}
}

// WriteResponse to the client
func (o *PostAccessTokenQueryUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// PostAccessTokenQueryForbiddenCode is the HTTP code returned for type PostAccessTokenQueryForbidden
const PostAccessTokenQueryForbiddenCode int = 403

/*
PostAccessTokenQueryForbidden Forbidden

swagger:response postAccessTokenQueryForbidden
*/
type PostAccessTokenQueryForbidden struct {
}

// NewPostAccessTokenQueryForbidden creates PostAccessTokenQueryForbidden with default headers values
func NewPostAccessTokenQueryForbidden() *PostAccessTokenQueryForbidden {

	return &PostAccessTokenQueryForbidden{}
}

// WriteResponse to the client
func (o *PostAccessTokenQueryForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// PostAccessTokenQueryInternalServerErrorCode is the HTTP code returned for type PostAccessTokenQueryInternalServerError
const PostAccessTokenQueryInternalServerErrorCode int = 500

/*
PostAccessTokenQueryInternalServerError Internal Server Error

swagger:response postAccessTokenQueryInternalServerError
*/
type PostAccessTokenQueryInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewPostAccessTokenQueryInternalServerError creates PostAccessTokenQueryInternalServerError with default headers values
func NewPostAccessTokenQueryInternalServerError() *PostAccessTokenQueryInternalServerError {

	return &PostAccessTokenQueryInternalServerError{}
}

// WithPayload adds the payload to the post access token query internal server error response
func (o *PostAccessTokenQueryInternalServerError) WithPayload(payload *models.APIResponse) *PostAccessTokenQueryInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post access token query internal server error response
func (o *PostAccessTokenQueryInternalServerError) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostAccessTokenQueryInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
