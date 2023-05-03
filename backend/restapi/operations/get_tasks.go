// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetTasksHandlerFunc turns a function with the right signature into a get tasks handler
type GetTasksHandlerFunc func(GetTasksParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTasksHandlerFunc) Handle(params GetTasksParams) middleware.Responder {
	return fn(params)
}

// GetTasksHandler interface for that can handle valid get tasks params
type GetTasksHandler interface {
	Handle(GetTasksParams) middleware.Responder
}

// NewGetTasks creates a new http.Handler for the get tasks operation
func NewGetTasks(ctx *middleware.Context, handler GetTasksHandler) *GetTasks {
	return &GetTasks{Context: ctx, Handler: handler}
}

/*
	GetTasks swagger:route GET /tasks getTasks

Get a list of tasks
*/
type GetTasks struct {
	Context *middleware.Context
	Handler GetTasksHandler
}

func (o *GetTasks) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetTasksParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetTasksBadRequestBody get tasks bad request body
//
// swagger:model GetTasksBadRequestBody
type GetTasksBadRequestBody struct {

	// message
	// Example: Invalid input, please provide the required parameters.
	Message string `json:"message,omitempty"`
}

// Validate validates this get tasks bad request body
func (o *GetTasksBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get tasks bad request body based on context it is used
func (o *GetTasksBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetTasksBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTasksBadRequestBody) UnmarshalBinary(b []byte) error {
	var res GetTasksBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetTasksInternalServerErrorBody get tasks internal server error body
//
// swagger:model GetTasksInternalServerErrorBody
type GetTasksInternalServerErrorBody struct {

	// message
	// Example: Internal server error.
	Message string `json:"message,omitempty"`
}

// Validate validates this get tasks internal server error body
func (o *GetTasksInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get tasks internal server error body based on context it is used
func (o *GetTasksInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetTasksInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTasksInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res GetTasksInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
