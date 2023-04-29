// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"golang-nextjs-app/restapi/models"
)

// GetTasksOKCode is the HTTP code returned for type GetTasksOK
const GetTasksOKCode int = 200

/*
GetTasksOK OK

swagger:response getTasksOK
*/
type GetTasksOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Task `json:"body,omitempty"`
}

// NewGetTasksOK creates GetTasksOK with default headers values
func NewGetTasksOK() *GetTasksOK {

	return &GetTasksOK{}
}

// WithPayload adds the payload to the get tasks o k response
func (o *GetTasksOK) WithPayload(payload []*models.Task) *GetTasksOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get tasks o k response
func (o *GetTasksOK) SetPayload(payload []*models.Task) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTasksOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Task, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
