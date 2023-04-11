// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteUsersIDOKCode is the HTTP code returned for type DeleteUsersIDOK
const DeleteUsersIDOKCode int = 200

/*
DeleteUsersIDOK OK

swagger:response deleteUsersIdOK
*/
type DeleteUsersIDOK struct {
}

// NewDeleteUsersIDOK creates DeleteUsersIDOK with default headers values
func NewDeleteUsersIDOK() *DeleteUsersIDOK {

	return &DeleteUsersIDOK{}
}

// WriteResponse to the client
func (o *DeleteUsersIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}
