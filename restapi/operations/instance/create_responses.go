// Copyright (c) 2022 Yaohui Wang (yaohuiwang@outlook.com)
// FlexLB is licensed under Mulan PubL v2.
// You can use this software according to the terms and conditions of the Mulan PubL v2.
// You may obtain a copy of Mulan PubL v2 at:
//         http://license.coscl.org.cn/MulanPubL-2.0
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
// EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
// MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PubL v2 for more details.

// Code generated by go-swagger; DO NOT EDIT.

package instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitee.com/flexlb/flexlb-api/models"
)

// CreateOKCode is the HTTP code returned for type CreateOK
const CreateOKCode int = 200

/*CreateOK Create Instance succeeded

swagger:response createOK
*/
type CreateOK struct {

	/*
	  In: Body
	*/
	Payload *models.Instance `json:"body,omitempty"`
}

// NewCreateOK creates CreateOK with default headers values
func NewCreateOK() *CreateOK {

	return &CreateOK{}
}

// WithPayload adds the payload to the create o k response
func (o *CreateOK) WithPayload(payload *models.Instance) *CreateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create o k response
func (o *CreateOK) SetPayload(payload *models.Instance) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateBadRequestCode is the HTTP code returned for type CreateBadRequest
const CreateBadRequestCode int = 400

/*CreateBadRequest create bad request

swagger:response createBadRequest
*/
type CreateBadRequest struct {

	/*
	  In: Body
	*/
	Payload *CreateBadRequestBody `json:"body,omitempty"`
}

// NewCreateBadRequest creates CreateBadRequest with default headers values
func NewCreateBadRequest() *CreateBadRequest {

	return &CreateBadRequest{}
}

// WithPayload adds the payload to the create bad request response
func (o *CreateBadRequest) WithPayload(payload *CreateBadRequestBody) *CreateBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create bad request response
func (o *CreateBadRequest) SetPayload(payload *CreateBadRequestBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateUnauthorizedCode is the HTTP code returned for type CreateUnauthorized
const CreateUnauthorizedCode int = 401

/*CreateUnauthorized create unauthorized

swagger:response createUnauthorized
*/
type CreateUnauthorized struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewCreateUnauthorized creates CreateUnauthorized with default headers values
func NewCreateUnauthorized() *CreateUnauthorized {

	return &CreateUnauthorized{}
}

// WithPayload adds the payload to the create unauthorized response
func (o *CreateUnauthorized) WithPayload(payload interface{}) *CreateUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create unauthorized response
func (o *CreateUnauthorized) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// CreateForbiddenCode is the HTTP code returned for type CreateForbidden
const CreateForbiddenCode int = 403

/*CreateForbidden create forbidden

swagger:response createForbidden
*/
type CreateForbidden struct {

	/*
	  In: Body
	*/
	Payload *CreateForbiddenBody `json:"body,omitempty"`
}

// NewCreateForbidden creates CreateForbidden with default headers values
func NewCreateForbidden() *CreateForbidden {

	return &CreateForbidden{}
}

// WithPayload adds the payload to the create forbidden response
func (o *CreateForbidden) WithPayload(payload *CreateForbiddenBody) *CreateForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create forbidden response
func (o *CreateForbidden) SetPayload(payload *CreateForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateNotFoundCode is the HTTP code returned for type CreateNotFound
const CreateNotFoundCode int = 404

/*CreateNotFound create not found

swagger:response createNotFound
*/
type CreateNotFound struct {

	/*
	  In: Body
	*/
	Payload *CreateNotFoundBody `json:"body,omitempty"`
}

// NewCreateNotFound creates CreateNotFound with default headers values
func NewCreateNotFound() *CreateNotFound {

	return &CreateNotFound{}
}

// WithPayload adds the payload to the create not found response
func (o *CreateNotFound) WithPayload(payload *CreateNotFoundBody) *CreateNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create not found response
func (o *CreateNotFound) SetPayload(payload *CreateNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateInternalServerErrorCode is the HTTP code returned for type CreateInternalServerError
const CreateInternalServerErrorCode int = 500

/*CreateInternalServerError create internal server error

swagger:response createInternalServerError
*/
type CreateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewCreateInternalServerError creates CreateInternalServerError with default headers values
func NewCreateInternalServerError() *CreateInternalServerError {

	return &CreateInternalServerError{}
}

// WithPayload adds the payload to the create internal server error response
func (o *CreateInternalServerError) WithPayload(payload interface{}) *CreateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create internal server error response
func (o *CreateInternalServerError) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
