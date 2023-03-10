// Code generated by go-swagger; DO NOT EDIT.

package token_registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"cybercert-blockchain-api/models"
)

// PostAccessTokenIssueStandardCreatedCode is the HTTP code returned for type PostAccessTokenIssueStandardCreated
const PostAccessTokenIssueStandardCreatedCode int = 201

/*
PostAccessTokenIssueStandardCreated OK

swagger:response postAccessTokenIssueStandardCreated
*/
type PostAccessTokenIssueStandardCreated struct {

	/*
	  In: Body
	*/
	Payload *models.CreatedResponse `json:"body,omitempty"`
}

// NewPostAccessTokenIssueStandardCreated creates PostAccessTokenIssueStandardCreated with default headers values
func NewPostAccessTokenIssueStandardCreated() *PostAccessTokenIssueStandardCreated {

	return &PostAccessTokenIssueStandardCreated{}
}

// WithPayload adds the payload to the post access token issue standard created response
func (o *PostAccessTokenIssueStandardCreated) WithPayload(payload *models.CreatedResponse) *PostAccessTokenIssueStandardCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post access token issue standard created response
func (o *PostAccessTokenIssueStandardCreated) SetPayload(payload *models.CreatedResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostAccessTokenIssueStandardCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostAccessTokenIssueStandardUnauthorizedCode is the HTTP code returned for type PostAccessTokenIssueStandardUnauthorized
const PostAccessTokenIssueStandardUnauthorizedCode int = 401

/*
PostAccessTokenIssueStandardUnauthorized Unauthorized

swagger:response postAccessTokenIssueStandardUnauthorized
*/
type PostAccessTokenIssueStandardUnauthorized struct {
}

// NewPostAccessTokenIssueStandardUnauthorized creates PostAccessTokenIssueStandardUnauthorized with default headers values
func NewPostAccessTokenIssueStandardUnauthorized() *PostAccessTokenIssueStandardUnauthorized {

	return &PostAccessTokenIssueStandardUnauthorized{}
}

// WriteResponse to the client
func (o *PostAccessTokenIssueStandardUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// PostAccessTokenIssueStandardInternalServerErrorCode is the HTTP code returned for type PostAccessTokenIssueStandardInternalServerError
const PostAccessTokenIssueStandardInternalServerErrorCode int = 500

/*
PostAccessTokenIssueStandardInternalServerError Internal Server Error

swagger:response postAccessTokenIssueStandardInternalServerError
*/
type PostAccessTokenIssueStandardInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewPostAccessTokenIssueStandardInternalServerError creates PostAccessTokenIssueStandardInternalServerError with default headers values
func NewPostAccessTokenIssueStandardInternalServerError() *PostAccessTokenIssueStandardInternalServerError {

	return &PostAccessTokenIssueStandardInternalServerError{}
}

// WithPayload adds the payload to the post access token issue standard internal server error response
func (o *PostAccessTokenIssueStandardInternalServerError) WithPayload(payload *models.APIResponse) *PostAccessTokenIssueStandardInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post access token issue standard internal server error response
func (o *PostAccessTokenIssueStandardInternalServerError) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostAccessTokenIssueStandardInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
