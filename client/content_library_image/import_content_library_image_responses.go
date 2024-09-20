// Code generated by go-swagger; DO NOT EDIT.

package content_library_image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/arcfraofficial/aoc-go-sdk/models"
)

// ImportContentLibraryImageReader is a Reader for the ImportContentLibraryImage structure.
type ImportContentLibraryImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImportContentLibraryImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImportContentLibraryImageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewImportContentLibraryImageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewImportContentLibraryImageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImportContentLibraryImageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewImportContentLibraryImageOK creates a ImportContentLibraryImageOK with default headers values
func NewImportContentLibraryImageOK() *ImportContentLibraryImageOK {
	return &ImportContentLibraryImageOK{}
}

/* ImportContentLibraryImageOK describes a response with status code 200, with default header values.

ImportContentLibraryImageOK import content library image o k
*/
type ImportContentLibraryImageOK struct {
	XAocRequestID string

	Payload *models.UploadTask
}

func (o *ImportContentLibraryImageOK) Error() string {
	return fmt.Sprintf("[POST /import-content-library-image][%d] importContentLibraryImageOK  %+v", 200, o.Payload)
}
func (o *ImportContentLibraryImageOK) GetPayload() *models.UploadTask {
	return o.Payload
}

func (o *ImportContentLibraryImageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-aoc-request-id
	hdrXAocRequestID := response.GetHeader("x-aoc-request-id")

	if hdrXAocRequestID != "" {
		o.XAocRequestID = hdrXAocRequestID
	}

	o.Payload = new(models.UploadTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportContentLibraryImageBadRequest creates a ImportContentLibraryImageBadRequest with default headers values
func NewImportContentLibraryImageBadRequest() *ImportContentLibraryImageBadRequest {
	return &ImportContentLibraryImageBadRequest{}
}

/* ImportContentLibraryImageBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ImportContentLibraryImageBadRequest struct {
	XAocRequestID string

	Payload *models.ErrorBody
}

func (o *ImportContentLibraryImageBadRequest) Error() string {
	return fmt.Sprintf("[POST /import-content-library-image][%d] importContentLibraryImageBadRequest  %+v", 400, o.Payload)
}
func (o *ImportContentLibraryImageBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *ImportContentLibraryImageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewImportContentLibraryImageNotFound creates a ImportContentLibraryImageNotFound with default headers values
func NewImportContentLibraryImageNotFound() *ImportContentLibraryImageNotFound {
	return &ImportContentLibraryImageNotFound{}
}

/* ImportContentLibraryImageNotFound describes a response with status code 404, with default header values.

Not found
*/
type ImportContentLibraryImageNotFound struct {
	XAocRequestID string

	Payload *models.ErrorBody
}

func (o *ImportContentLibraryImageNotFound) Error() string {
	return fmt.Sprintf("[POST /import-content-library-image][%d] importContentLibraryImageNotFound  %+v", 404, o.Payload)
}
func (o *ImportContentLibraryImageNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *ImportContentLibraryImageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewImportContentLibraryImageInternalServerError creates a ImportContentLibraryImageInternalServerError with default headers values
func NewImportContentLibraryImageInternalServerError() *ImportContentLibraryImageInternalServerError {
	return &ImportContentLibraryImageInternalServerError{}
}

/* ImportContentLibraryImageInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type ImportContentLibraryImageInternalServerError struct {
	XAocRequestID string

	Payload *models.ErrorBody
}

func (o *ImportContentLibraryImageInternalServerError) Error() string {
	return fmt.Sprintf("[POST /import-content-library-image][%d] importContentLibraryImageInternalServerError  %+v", 500, o.Payload)
}
func (o *ImportContentLibraryImageInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *ImportContentLibraryImageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
