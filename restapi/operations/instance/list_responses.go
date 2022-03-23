// Code generated by go-swagger; DO NOT EDIT.

package instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitee.com/flexlb/flexlb-api/models"
)

// ListOKCode is the HTTP code returned for type ListOK
const ListOKCode int = 200

/*ListOK List Instances succeeded

swagger:response listOK
*/
type ListOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Instance `json:"body,omitempty"`
}

// NewListOK creates ListOK with default headers values
func NewListOK() *ListOK {

	return &ListOK{}
}

// WithPayload adds the payload to the list o k response
func (o *ListOK) WithPayload(payload []*models.Instance) *ListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list o k response
func (o *ListOK) SetPayload(payload []*models.Instance) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Instance, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListBadRequestCode is the HTTP code returned for type ListBadRequest
const ListBadRequestCode int = 400

/*ListBadRequest list bad request

swagger:response listBadRequest
*/
type ListBadRequest struct {

	/*
	  In: Body
	*/
	Payload *ListBadRequestBody `json:"body,omitempty"`
}

// NewListBadRequest creates ListBadRequest with default headers values
func NewListBadRequest() *ListBadRequest {

	return &ListBadRequest{}
}

// WithPayload adds the payload to the list bad request response
func (o *ListBadRequest) WithPayload(payload *ListBadRequestBody) *ListBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list bad request response
func (o *ListBadRequest) SetPayload(payload *ListBadRequestBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListUnauthorizedCode is the HTTP code returned for type ListUnauthorized
const ListUnauthorizedCode int = 401

/*ListUnauthorized list unauthorized

swagger:response listUnauthorized
*/
type ListUnauthorized struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewListUnauthorized creates ListUnauthorized with default headers values
func NewListUnauthorized() *ListUnauthorized {

	return &ListUnauthorized{}
}

// WithPayload adds the payload to the list unauthorized response
func (o *ListUnauthorized) WithPayload(payload interface{}) *ListUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list unauthorized response
func (o *ListUnauthorized) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListForbiddenCode is the HTTP code returned for type ListForbidden
const ListForbiddenCode int = 403

/*ListForbidden list forbidden

swagger:response listForbidden
*/
type ListForbidden struct {

	/*
	  In: Body
	*/
	Payload *ListForbiddenBody `json:"body,omitempty"`
}

// NewListForbidden creates ListForbidden with default headers values
func NewListForbidden() *ListForbidden {

	return &ListForbidden{}
}

// WithPayload adds the payload to the list forbidden response
func (o *ListForbidden) WithPayload(payload *ListForbiddenBody) *ListForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list forbidden response
func (o *ListForbidden) SetPayload(payload *ListForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListNotFoundCode is the HTTP code returned for type ListNotFound
const ListNotFoundCode int = 404

/*ListNotFound list not found

swagger:response listNotFound
*/
type ListNotFound struct {

	/*
	  In: Body
	*/
	Payload *ListNotFoundBody `json:"body,omitempty"`
}

// NewListNotFound creates ListNotFound with default headers values
func NewListNotFound() *ListNotFound {

	return &ListNotFound{}
}

// WithPayload adds the payload to the list not found response
func (o *ListNotFound) WithPayload(payload *ListNotFoundBody) *ListNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list not found response
func (o *ListNotFound) SetPayload(payload *ListNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListInternalServerErrorCode is the HTTP code returned for type ListInternalServerError
const ListInternalServerErrorCode int = 500

/*ListInternalServerError list internal server error

swagger:response listInternalServerError
*/
type ListInternalServerError struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewListInternalServerError creates ListInternalServerError with default headers values
func NewListInternalServerError() *ListInternalServerError {

	return &ListInternalServerError{}
}

// WithPayload adds the payload to the list internal server error response
func (o *ListInternalServerError) WithPayload(payload interface{}) *ListInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list internal server error response
func (o *ListInternalServerError) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
