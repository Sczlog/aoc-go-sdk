// Code generated by go-swagger; DO NOT EDIT.

package vm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/arcfraofficial/aoc-go-sdk/models"
)

// DeleteVMReader is a Reader for the DeleteVM structure.
type DeleteVMReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVMReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteVMOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteVMBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteVMNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteVMInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteVMOK creates a DeleteVMOK with default headers values
func NewDeleteVMOK() *DeleteVMOK {
	return &DeleteVMOK{}
}

/* DeleteVMOK describes a response with status code 200, with default header values.

DeleteVMOK delete Vm o k
*/
type DeleteVMOK struct {
	XAocRequestID string

	Payload []*models.WithTaskDeleteVM
}

func (o *DeleteVMOK) Error() string {
	return fmt.Sprintf("[POST /delete-vm][%d] deleteVmOK  %+v", 200, o.Payload)
}
func (o *DeleteVMOK) GetPayload() []*models.WithTaskDeleteVM {
	return o.Payload
}

func (o *DeleteVMOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteVMBadRequest creates a DeleteVMBadRequest with default headers values
func NewDeleteVMBadRequest() *DeleteVMBadRequest {
	return &DeleteVMBadRequest{}
}

/* DeleteVMBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeleteVMBadRequest struct {
	XAocRequestID string

	Payload *models.ErrorBody
}

func (o *DeleteVMBadRequest) Error() string {
	return fmt.Sprintf("[POST /delete-vm][%d] deleteVmBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteVMBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteVMBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteVMNotFound creates a DeleteVMNotFound with default headers values
func NewDeleteVMNotFound() *DeleteVMNotFound {
	return &DeleteVMNotFound{}
}

/* DeleteVMNotFound describes a response with status code 404, with default header values.

Not found
*/
type DeleteVMNotFound struct {
	XAocRequestID string

	Payload *models.ErrorBody
}

func (o *DeleteVMNotFound) Error() string {
	return fmt.Sprintf("[POST /delete-vm][%d] deleteVmNotFound  %+v", 404, o.Payload)
}
func (o *DeleteVMNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteVMNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteVMInternalServerError creates a DeleteVMInternalServerError with default headers values
func NewDeleteVMInternalServerError() *DeleteVMInternalServerError {
	return &DeleteVMInternalServerError{}
}

/* DeleteVMInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type DeleteVMInternalServerError struct {
	XAocRequestID string

	Payload *models.ErrorBody
}

func (o *DeleteVMInternalServerError) Error() string {
	return fmt.Sprintf("[POST /delete-vm][%d] deleteVmInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteVMInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteVMInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
