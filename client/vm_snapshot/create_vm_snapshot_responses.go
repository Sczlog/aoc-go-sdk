// Code generated by go-swagger; DO NOT EDIT.

package vm_snapshot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/arcfraofficial/aoc-go-sdk/models"
)

// CreateVMSnapshotReader is a Reader for the CreateVMSnapshot structure.
type CreateVMSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateVMSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateVMSnapshotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateVMSnapshotBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateVMSnapshotNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateVMSnapshotInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateVMSnapshotOK creates a CreateVMSnapshotOK with default headers values
func NewCreateVMSnapshotOK() *CreateVMSnapshotOK {
	return &CreateVMSnapshotOK{}
}

/* CreateVMSnapshotOK describes a response with status code 200, with default header values.

CreateVMSnapshotOK create Vm snapshot o k
*/
type CreateVMSnapshotOK struct {
	XAocRequestID string

	Payload []*models.WithTaskVMSnapshot
}

func (o *CreateVMSnapshotOK) Error() string {
	return fmt.Sprintf("[POST /create-vm-snapshot][%d] createVmSnapshotOK  %+v", 200, o.Payload)
}
func (o *CreateVMSnapshotOK) GetPayload() []*models.WithTaskVMSnapshot {
	return o.Payload
}

func (o *CreateVMSnapshotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateVMSnapshotBadRequest creates a CreateVMSnapshotBadRequest with default headers values
func NewCreateVMSnapshotBadRequest() *CreateVMSnapshotBadRequest {
	return &CreateVMSnapshotBadRequest{}
}

/* CreateVMSnapshotBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateVMSnapshotBadRequest struct {
	XAocRequestID string

	Payload *models.ErrorBody
}

func (o *CreateVMSnapshotBadRequest) Error() string {
	return fmt.Sprintf("[POST /create-vm-snapshot][%d] createVmSnapshotBadRequest  %+v", 400, o.Payload)
}
func (o *CreateVMSnapshotBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateVMSnapshotBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateVMSnapshotNotFound creates a CreateVMSnapshotNotFound with default headers values
func NewCreateVMSnapshotNotFound() *CreateVMSnapshotNotFound {
	return &CreateVMSnapshotNotFound{}
}

/* CreateVMSnapshotNotFound describes a response with status code 404, with default header values.

Not found
*/
type CreateVMSnapshotNotFound struct {
	XAocRequestID string

	Payload *models.ErrorBody
}

func (o *CreateVMSnapshotNotFound) Error() string {
	return fmt.Sprintf("[POST /create-vm-snapshot][%d] createVmSnapshotNotFound  %+v", 404, o.Payload)
}
func (o *CreateVMSnapshotNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateVMSnapshotNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateVMSnapshotInternalServerError creates a CreateVMSnapshotInternalServerError with default headers values
func NewCreateVMSnapshotInternalServerError() *CreateVMSnapshotInternalServerError {
	return &CreateVMSnapshotInternalServerError{}
}

/* CreateVMSnapshotInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type CreateVMSnapshotInternalServerError struct {
	XAocRequestID string

	Payload *models.ErrorBody
}

func (o *CreateVMSnapshotInternalServerError) Error() string {
	return fmt.Sprintf("[POST /create-vm-snapshot][%d] createVmSnapshotInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateVMSnapshotInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateVMSnapshotInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
