// Code generated by go-swagger; DO NOT EDIT.

package vm_folder

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/arcfraofficial/aoc-go-sdk/models"
)

// UpdateVMFolderReader is a Reader for the UpdateVMFolder structure.
type UpdateVMFolderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateVMFolderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateVMFolderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateVMFolderBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateVMFolderNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateVMFolderInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateVMFolderOK creates a UpdateVMFolderOK with default headers values
func NewUpdateVMFolderOK() *UpdateVMFolderOK {
	return &UpdateVMFolderOK{}
}

/* UpdateVMFolderOK describes a response with status code 200, with default header values.

UpdateVMFolderOK update Vm folder o k
*/
type UpdateVMFolderOK struct {
	XAocRequestID string

	Payload []*models.WithTaskVMFolder
}

func (o *UpdateVMFolderOK) Error() string {
	return fmt.Sprintf("[POST /update-vm-folder][%d] updateVmFolderOK  %+v", 200, o.Payload)
}
func (o *UpdateVMFolderOK) GetPayload() []*models.WithTaskVMFolder {
	return o.Payload
}

func (o *UpdateVMFolderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-aoc-request-id
	hdrXAocRequestID := response.GetHeader("x-aoc-request-id")

	if hdrXAocRequestID != "" {
		o.XAocRequestID = hdrXAocRequestID
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVMFolderBadRequest creates a UpdateVMFolderBadRequest with default headers values
func NewUpdateVMFolderBadRequest() *UpdateVMFolderBadRequest {
	return &UpdateVMFolderBadRequest{}
}

/* UpdateVMFolderBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type UpdateVMFolderBadRequest struct {
	XAocRequestID string

	Payload *models.ErrorBody
}

func (o *UpdateVMFolderBadRequest) Error() string {
	return fmt.Sprintf("[POST /update-vm-folder][%d] updateVmFolderBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateVMFolderBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateVMFolderBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-aoc-request-id
	hdrXAocRequestID := response.GetHeader("x-aoc-request-id")

	if hdrXAocRequestID != "" {
		o.XAocRequestID = hdrXAocRequestID
	}

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVMFolderNotFound creates a UpdateVMFolderNotFound with default headers values
func NewUpdateVMFolderNotFound() *UpdateVMFolderNotFound {
	return &UpdateVMFolderNotFound{}
}

/* UpdateVMFolderNotFound describes a response with status code 404, with default header values.

Not found
*/
type UpdateVMFolderNotFound struct {
	XAocRequestID string

	Payload *models.ErrorBody
}

func (o *UpdateVMFolderNotFound) Error() string {
	return fmt.Sprintf("[POST /update-vm-folder][%d] updateVmFolderNotFound  %+v", 404, o.Payload)
}
func (o *UpdateVMFolderNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateVMFolderNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-aoc-request-id
	hdrXAocRequestID := response.GetHeader("x-aoc-request-id")

	if hdrXAocRequestID != "" {
		o.XAocRequestID = hdrXAocRequestID
	}

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVMFolderInternalServerError creates a UpdateVMFolderInternalServerError with default headers values
func NewUpdateVMFolderInternalServerError() *UpdateVMFolderInternalServerError {
	return &UpdateVMFolderInternalServerError{}
}

/* UpdateVMFolderInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type UpdateVMFolderInternalServerError struct {
	XAocRequestID string

	Payload *models.ErrorBody
}

func (o *UpdateVMFolderInternalServerError) Error() string {
	return fmt.Sprintf("[POST /update-vm-folder][%d] updateVmFolderInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateVMFolderInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateVMFolderInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-aoc-request-id
	hdrXAocRequestID := response.GetHeader("x-aoc-request-id")

	if hdrXAocRequestID != "" {
		o.XAocRequestID = hdrXAocRequestID
	}

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
