// Code generated by go-swagger; DO NOT EDIT.

package certificate_info

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"cybercert-blockchain-api/models"
)

// PostCertificateInfoIssueCreatedCode is the HTTP code returned for type PostCertificateInfoIssueCreated
const PostCertificateInfoIssueCreatedCode int = 201

/*
PostCertificateInfoIssueCreated OK

swagger:response postCertificateInfoIssueCreated
*/
type PostCertificateInfoIssueCreated struct {

	/*
	  In: Body
	*/
	Payload *models.CreatedResponse `json:"body,omitempty"`
}

// NewPostCertificateInfoIssueCreated creates PostCertificateInfoIssueCreated with default headers values
func NewPostCertificateInfoIssueCreated() *PostCertificateInfoIssueCreated {

	return &PostCertificateInfoIssueCreated{}
}

// WithPayload adds the payload to the post certificate info issue created response
func (o *PostCertificateInfoIssueCreated) WithPayload(payload *models.CreatedResponse) *PostCertificateInfoIssueCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post certificate info issue created response
func (o *PostCertificateInfoIssueCreated) SetPayload(payload *models.CreatedResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostCertificateInfoIssueCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostCertificateInfoIssueUnauthorizedCode is the HTTP code returned for type PostCertificateInfoIssueUnauthorized
const PostCertificateInfoIssueUnauthorizedCode int = 401

/*
PostCertificateInfoIssueUnauthorized Unauthorized

swagger:response postCertificateInfoIssueUnauthorized
*/
type PostCertificateInfoIssueUnauthorized struct {
}

// NewPostCertificateInfoIssueUnauthorized creates PostCertificateInfoIssueUnauthorized with default headers values
func NewPostCertificateInfoIssueUnauthorized() *PostCertificateInfoIssueUnauthorized {

	return &PostCertificateInfoIssueUnauthorized{}
}

// WriteResponse to the client
func (o *PostCertificateInfoIssueUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// PostCertificateInfoIssueInternalServerErrorCode is the HTTP code returned for type PostCertificateInfoIssueInternalServerError
const PostCertificateInfoIssueInternalServerErrorCode int = 500

/*
PostCertificateInfoIssueInternalServerError Internal Server Error

swagger:response postCertificateInfoIssueInternalServerError
*/
type PostCertificateInfoIssueInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewPostCertificateInfoIssueInternalServerError creates PostCertificateInfoIssueInternalServerError with default headers values
func NewPostCertificateInfoIssueInternalServerError() *PostCertificateInfoIssueInternalServerError {

	return &PostCertificateInfoIssueInternalServerError{}
}

// WithPayload adds the payload to the post certificate info issue internal server error response
func (o *PostCertificateInfoIssueInternalServerError) WithPayload(payload *models.APIResponse) *PostCertificateInfoIssueInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post certificate info issue internal server error response
func (o *PostCertificateInfoIssueInternalServerError) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostCertificateInfoIssueInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}