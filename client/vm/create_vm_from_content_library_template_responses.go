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

// CreateVMFromContentLibraryTemplateReader is a Reader for the CreateVMFromContentLibraryTemplate structure.
type CreateVMFromContentLibraryTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateVMFromContentLibraryTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateVMFromContentLibraryTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateVMFromContentLibraryTemplateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateVMFromContentLibraryTemplateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateVMFromContentLibraryTemplateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateVMFromContentLibraryTemplateOK creates a CreateVMFromContentLibraryTemplateOK with default headers values
func NewCreateVMFromContentLibraryTemplateOK() *CreateVMFromContentLibraryTemplateOK {
	return &CreateVMFromContentLibraryTemplateOK{}
}

/* CreateVMFromContentLibraryTemplateOK describes a response with status code 200, with default header values.

CreateVMFromContentLibraryTemplateOK create Vm from content library template o k
*/
type CreateVMFromContentLibraryTemplateOK struct {
	XAocRequestID string

	Payload []*models.WithTaskVM
}

func (o *CreateVMFromContentLibraryTemplateOK) Error() string {
	return fmt.Sprintf("[POST /create-vm-from-content-library-template][%d] createVmFromContentLibraryTemplateOK  %+v", 200, o.Payload)
}
func (o *CreateVMFromContentLibraryTemplateOK) GetPayload() []*models.WithTaskVM {
	return o.Payload
}

func (o *CreateVMFromContentLibraryTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateVMFromContentLibraryTemplateBadRequest creates a CreateVMFromContentLibraryTemplateBadRequest with default headers values
func NewCreateVMFromContentLibraryTemplateBadRequest() *CreateVMFromContentLibraryTemplateBadRequest {
	return &CreateVMFromContentLibraryTemplateBadRequest{}
}

/* CreateVMFromContentLibraryTemplateBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateVMFromContentLibraryTemplateBadRequest struct {
	XAocRequestID string

	Payload *models.ErrorBody
}

func (o *CreateVMFromContentLibraryTemplateBadRequest) Error() string {
	return fmt.Sprintf("[POST /create-vm-from-content-library-template][%d] createVmFromContentLibraryTemplateBadRequest  %+v", 400, o.Payload)
}
func (o *CreateVMFromContentLibraryTemplateBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateVMFromContentLibraryTemplateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateVMFromContentLibraryTemplateNotFound creates a CreateVMFromContentLibraryTemplateNotFound with default headers values
func NewCreateVMFromContentLibraryTemplateNotFound() *CreateVMFromContentLibraryTemplateNotFound {
	return &CreateVMFromContentLibraryTemplateNotFound{}
}

/* CreateVMFromContentLibraryTemplateNotFound describes a response with status code 404, with default header values.

Not found
*/
type CreateVMFromContentLibraryTemplateNotFound struct {
	XAocRequestID string

	Payload *models.ErrorBody
}

func (o *CreateVMFromContentLibraryTemplateNotFound) Error() string {
	return fmt.Sprintf("[POST /create-vm-from-content-library-template][%d] createVmFromContentLibraryTemplateNotFound  %+v", 404, o.Payload)
}
func (o *CreateVMFromContentLibraryTemplateNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateVMFromContentLibraryTemplateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateVMFromContentLibraryTemplateInternalServerError creates a CreateVMFromContentLibraryTemplateInternalServerError with default headers values
func NewCreateVMFromContentLibraryTemplateInternalServerError() *CreateVMFromContentLibraryTemplateInternalServerError {
	return &CreateVMFromContentLibraryTemplateInternalServerError{}
}

/* CreateVMFromContentLibraryTemplateInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type CreateVMFromContentLibraryTemplateInternalServerError struct {
	XAocRequestID string

	Payload *models.ErrorBody
}

func (o *CreateVMFromContentLibraryTemplateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /create-vm-from-content-library-template][%d] createVmFromContentLibraryTemplateInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateVMFromContentLibraryTemplateInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateVMFromContentLibraryTemplateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
