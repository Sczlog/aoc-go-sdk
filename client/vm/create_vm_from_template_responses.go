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

// CreateVMFromTemplateReader is a Reader for the CreateVMFromTemplate structure.
type CreateVMFromTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateVMFromTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateVMFromTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateVMFromTemplateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateVMFromTemplateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateVMFromTemplateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateVMFromTemplateOK creates a CreateVMFromTemplateOK with default headers values
func NewCreateVMFromTemplateOK() *CreateVMFromTemplateOK {
	return &CreateVMFromTemplateOK{}
}

/* CreateVMFromTemplateOK describes a response with status code 200, with default header values.

CreateVMFromTemplateOK create Vm from template o k
*/
type CreateVMFromTemplateOK struct {
	XAocRequestID string

	Payload []*models.WithTaskVM
}

func (o *CreateVMFromTemplateOK) Error() string {
	return fmt.Sprintf("[POST /create-vm-from-template][%d] createVmFromTemplateOK  %+v", 200, o.Payload)
}
func (o *CreateVMFromTemplateOK) GetPayload() []*models.WithTaskVM {
	return o.Payload
}

func (o *CreateVMFromTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateVMFromTemplateBadRequest creates a CreateVMFromTemplateBadRequest with default headers values
func NewCreateVMFromTemplateBadRequest() *CreateVMFromTemplateBadRequest {
	return &CreateVMFromTemplateBadRequest{}
}

/* CreateVMFromTemplateBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateVMFromTemplateBadRequest struct {
	XAocRequestID string

	Payload *models.ErrorBody
}

func (o *CreateVMFromTemplateBadRequest) Error() string {
	return fmt.Sprintf("[POST /create-vm-from-template][%d] createVmFromTemplateBadRequest  %+v", 400, o.Payload)
}
func (o *CreateVMFromTemplateBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateVMFromTemplateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateVMFromTemplateNotFound creates a CreateVMFromTemplateNotFound with default headers values
func NewCreateVMFromTemplateNotFound() *CreateVMFromTemplateNotFound {
	return &CreateVMFromTemplateNotFound{}
}

/* CreateVMFromTemplateNotFound describes a response with status code 404, with default header values.

Not found
*/
type CreateVMFromTemplateNotFound struct {
	XAocRequestID string

	Payload *models.ErrorBody
}

func (o *CreateVMFromTemplateNotFound) Error() string {
	return fmt.Sprintf("[POST /create-vm-from-template][%d] createVmFromTemplateNotFound  %+v", 404, o.Payload)
}
func (o *CreateVMFromTemplateNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateVMFromTemplateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateVMFromTemplateInternalServerError creates a CreateVMFromTemplateInternalServerError with default headers values
func NewCreateVMFromTemplateInternalServerError() *CreateVMFromTemplateInternalServerError {
	return &CreateVMFromTemplateInternalServerError{}
}

/* CreateVMFromTemplateInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type CreateVMFromTemplateInternalServerError struct {
	XAocRequestID string

	Payload *models.ErrorBody
}

func (o *CreateVMFromTemplateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /create-vm-from-template][%d] createVmFromTemplateInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateVMFromTemplateInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateVMFromTemplateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
