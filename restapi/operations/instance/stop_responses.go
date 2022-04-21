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

// StopOKCode is the HTTP code returned for type StopOK
const StopOKCode int = 200

/*StopOK Stop Instance succeeded

swagger:response stopOK
*/
type StopOK struct {

	/*
	  In: Body
	*/
	Payload *models.Instance `json:"body,omitempty"`
}

// NewStopOK creates StopOK with default headers values
func NewStopOK() *StopOK {

	return &StopOK{}
}

// WithPayload adds the payload to the stop o k response
func (o *StopOK) WithPayload(payload *models.Instance) *StopOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the stop o k response
func (o *StopOK) SetPayload(payload *models.Instance) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *StopOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// StopBadRequestCode is the HTTP code returned for type StopBadRequest
const StopBadRequestCode int = 400

/*StopBadRequest stop bad request

swagger:response stopBadRequest
*/
type StopBadRequest struct {

	/*
	  In: Body
	*/
	Payload *StopBadRequestBody `json:"body,omitempty"`
}

// NewStopBadRequest creates StopBadRequest with default headers values
func NewStopBadRequest() *StopBadRequest {

	return &StopBadRequest{}
}

// WithPayload adds the payload to the stop bad request response
func (o *StopBadRequest) WithPayload(payload *StopBadRequestBody) *StopBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the stop bad request response
func (o *StopBadRequest) SetPayload(payload *StopBadRequestBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *StopBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// StopUnauthorizedCode is the HTTP code returned for type StopUnauthorized
const StopUnauthorizedCode int = 401

/*StopUnauthorized stop unauthorized

swagger:response stopUnauthorized
*/
type StopUnauthorized struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewStopUnauthorized creates StopUnauthorized with default headers values
func NewStopUnauthorized() *StopUnauthorized {

	return &StopUnauthorized{}
}

// WithPayload adds the payload to the stop unauthorized response
func (o *StopUnauthorized) WithPayload(payload interface{}) *StopUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the stop unauthorized response
func (o *StopUnauthorized) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *StopUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// StopForbiddenCode is the HTTP code returned for type StopForbidden
const StopForbiddenCode int = 403

/*StopForbidden stop forbidden

swagger:response stopForbidden
*/
type StopForbidden struct {

	/*
	  In: Body
	*/
	Payload *StopForbiddenBody `json:"body,omitempty"`
}

// NewStopForbidden creates StopForbidden with default headers values
func NewStopForbidden() *StopForbidden {

	return &StopForbidden{}
}

// WithPayload adds the payload to the stop forbidden response
func (o *StopForbidden) WithPayload(payload *StopForbiddenBody) *StopForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the stop forbidden response
func (o *StopForbidden) SetPayload(payload *StopForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *StopForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// StopNotFoundCode is the HTTP code returned for type StopNotFound
const StopNotFoundCode int = 404

/*StopNotFound stop not found

swagger:response stopNotFound
*/
type StopNotFound struct {

	/*
	  In: Body
	*/
	Payload *StopNotFoundBody `json:"body,omitempty"`
}

// NewStopNotFound creates StopNotFound with default headers values
func NewStopNotFound() *StopNotFound {

	return &StopNotFound{}
}

// WithPayload adds the payload to the stop not found response
func (o *StopNotFound) WithPayload(payload *StopNotFoundBody) *StopNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the stop not found response
func (o *StopNotFound) SetPayload(payload *StopNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *StopNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// StopInternalServerErrorCode is the HTTP code returned for type StopInternalServerError
const StopInternalServerErrorCode int = 500

/*StopInternalServerError stop internal server error

swagger:response stopInternalServerError
*/
type StopInternalServerError struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewStopInternalServerError creates StopInternalServerError with default headers values
func NewStopInternalServerError() *StopInternalServerError {

	return &StopInternalServerError{}
}

// WithPayload adds the payload to the stop internal server error response
func (o *StopInternalServerError) WithPayload(payload interface{}) *StopInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the stop internal server error response
func (o *StopInternalServerError) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *StopInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
