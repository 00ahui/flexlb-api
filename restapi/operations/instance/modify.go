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
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ModifyHandlerFunc turns a function with the right signature into a modify handler
type ModifyHandlerFunc func(ModifyParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ModifyHandlerFunc) Handle(params ModifyParams) middleware.Responder {
	return fn(params)
}

// ModifyHandler interface for that can handle valid modify params
type ModifyHandler interface {
	Handle(ModifyParams) middleware.Responder
}

// NewModify creates a new http.Handler for the modify operation
func NewModify(ctx *middleware.Context, handler ModifyHandler) *Modify {
	return &Modify{Context: ctx, Handler: handler}
}

/* Modify swagger:route PUT /instances Instance modify

Modify Instance

*/
type Modify struct {
	Context *middleware.Context
	Handler ModifyHandler
}

func (o *Modify) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewModifyParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// ModifyBadRequestBody modify bad request body
//
// swagger:model ModifyBadRequestBody
type ModifyBadRequestBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this modify bad request body
func (o *ModifyBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ModifyBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("modifyBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this modify bad request body based on context it is used
func (o *ModifyBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ModifyBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ModifyBadRequestBody) UnmarshalBinary(b []byte) error {
	var res ModifyBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// ModifyForbiddenBody modify forbidden body
//
// swagger:model ModifyForbiddenBody
type ModifyForbiddenBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this modify forbidden body
func (o *ModifyForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ModifyForbiddenBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("modifyForbidden"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this modify forbidden body based on context it is used
func (o *ModifyForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ModifyForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ModifyForbiddenBody) UnmarshalBinary(b []byte) error {
	var res ModifyForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// ModifyNotFoundBody modify not found body
//
// swagger:model ModifyNotFoundBody
type ModifyNotFoundBody struct {

	// error
	// Required: true
	Error *string `json:"error"`

	// status
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this modify not found body
func (o *ModifyNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ModifyNotFoundBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("modifyNotFound"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

func (o *ModifyNotFoundBody) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("modifyNotFound"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this modify not found body based on context it is used
func (o *ModifyNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ModifyNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ModifyNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ModifyNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
