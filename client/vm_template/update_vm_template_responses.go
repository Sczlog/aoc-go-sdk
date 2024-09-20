// Code generated by go-swagger; DO NOT EDIT.

package vm_template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/arcfraofficial/aoc-go-sdk/models"
)

// UpdateVMTemplateReader is a Reader for the UpdateVMTemplate structure.
type UpdateVMTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateVMTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateVMTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateVMTemplateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateVMTemplateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateVMTemplateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateVMTemplateOK creates a UpdateVMTemplateOK with default headers values
func NewUpdateVMTemplateOK() *UpdateVMTemplateOK {
	return &UpdateVMTemplateOK{}
}

/* UpdateVMTemplateOK describes a response with status code 200, with default header values.

UpdateVMTemplateOK update Vm template o k
*/
type UpdateVMTemplateOK struct {
	XAocRequestID string

	Payload []*models.WithTaskVMTemplate
}

func (o *UpdateVMTemplateOK) Error() string {
	return fmt.Sprintf("[POST /update-vm-template][%d] updateVmTemplateOK  %+v", 200, o.Payload)
}
func (o *UpdateVMTemplateOK) GetPayload() []*models.WithTaskVMTemplate {
	return o.Payload
}

func (o *UpdateVMTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateVMTemplateBadRequest creates a UpdateVMTemplateBadRequest with default headers values
func NewUpdateVMTemplateBadRequest() *UpdateVMTemplateBadRequest {
	return &UpdateVMTemplateBadRequest{}
}

/* UpdateVMTemplateBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type UpdateVMTemplateBadRequest struct {
	XAocRequestID string

	Payload *models.ErrorBody
}

func (o *UpdateVMTemplateBadRequest) Error() string {
	return fmt.Sprintf("[POST /update-vm-template][%d] updateVmTemplateBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateVMTemplateBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateVMTemplateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateVMTemplateNotFound creates a UpdateVMTemplateNotFound with default headers values
func NewUpdateVMTemplateNotFound() *UpdateVMTemplateNotFound {
	return &UpdateVMTemplateNotFound{}
}

/* UpdateVMTemplateNotFound describes a response with status code 404, with default header values.

Not found
*/
type UpdateVMTemplateNotFound struct {
	XAocRequestID string

	Payload *models.ErrorBody
}

func (o *UpdateVMTemplateNotFound) Error() string {
	return fmt.Sprintf("[POST /update-vm-template][%d] updateVmTemplateNotFound  %+v", 404, o.Payload)
}
func (o *UpdateVMTemplateNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateVMTemplateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateVMTemplateInternalServerError creates a UpdateVMTemplateInternalServerError with default headers values
func NewUpdateVMTemplateInternalServerError() *UpdateVMTemplateInternalServerError {
	return &UpdateVMTemplateInternalServerError{}
}

/* UpdateVMTemplateInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type UpdateVMTemplateInternalServerError struct {
	XAocRequestID string

	Payload *models.ErrorBody
}

func (o *UpdateVMTemplateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /update-vm-template][%d] updateVmTemplateInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateVMTemplateInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateVMTemplateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
