// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"flexlb/models"
)

// ReadyzOKCode is the HTTP code returned for type ReadyzOK
const ReadyzOKCode int = 200

/*ReadyzOK Ready status

swagger:response readyzOK
*/
type ReadyzOK struct {

	/*
	  In: Body
	*/
	Payload *models.ReadyStatus `json:"body,omitempty"`
}

// NewReadyzOK creates ReadyzOK with default headers values
func NewReadyzOK() *ReadyzOK {

	return &ReadyzOK{}
}

// WithPayload adds the payload to the readyz o k response
func (o *ReadyzOK) WithPayload(payload *models.ReadyStatus) *ReadyzOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the readyz o k response
func (o *ReadyzOK) SetPayload(payload *models.ReadyStatus) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadyzOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
