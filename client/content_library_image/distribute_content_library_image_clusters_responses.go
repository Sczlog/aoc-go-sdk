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

// DistributeContentLibraryImageClustersReader is a Reader for the DistributeContentLibraryImageClusters structure.
type DistributeContentLibraryImageClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DistributeContentLibraryImageClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDistributeContentLibraryImageClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDistributeContentLibraryImageClustersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDistributeContentLibraryImageClustersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDistributeContentLibraryImageClustersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDistributeContentLibraryImageClustersOK creates a DistributeContentLibraryImageClustersOK with default headers values
func NewDistributeContentLibraryImageClustersOK() *DistributeContentLibraryImageClustersOK {
	return &DistributeContentLibraryImageClustersOK{}
}

/* DistributeContentLibraryImageClustersOK describes a response with status code 200, with default header values.

DistributeContentLibraryImageClustersOK distribute content library image clusters o k
*/
type DistributeContentLibraryImageClustersOK struct {
	XAocRequestID string

	Payload []*models.WithTaskContentLibraryImage
}

func (o *DistributeContentLibraryImageClustersOK) Error() string {
	return fmt.Sprintf("[POST /distribute-content-library-image-clusters][%d] distributeContentLibraryImageClustersOK  %+v", 200, o.Payload)
}
func (o *DistributeContentLibraryImageClustersOK) GetPayload() []*models.WithTaskContentLibraryImage {
	return o.Payload
}

func (o *DistributeContentLibraryImageClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDistributeContentLibraryImageClustersBadRequest creates a DistributeContentLibraryImageClustersBadRequest with default headers values
func NewDistributeContentLibraryImageClustersBadRequest() *DistributeContentLibraryImageClustersBadRequest {
	return &DistributeContentLibraryImageClustersBadRequest{}
}

/* DistributeContentLibraryImageClustersBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DistributeContentLibraryImageClustersBadRequest struct {
	XAocRequestID string

	Payload *models.ErrorBody
}

func (o *DistributeContentLibraryImageClustersBadRequest) Error() string {
	return fmt.Sprintf("[POST /distribute-content-library-image-clusters][%d] distributeContentLibraryImageClustersBadRequest  %+v", 400, o.Payload)
}
func (o *DistributeContentLibraryImageClustersBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DistributeContentLibraryImageClustersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDistributeContentLibraryImageClustersNotFound creates a DistributeContentLibraryImageClustersNotFound with default headers values
func NewDistributeContentLibraryImageClustersNotFound() *DistributeContentLibraryImageClustersNotFound {
	return &DistributeContentLibraryImageClustersNotFound{}
}

/* DistributeContentLibraryImageClustersNotFound describes a response with status code 404, with default header values.

Not found
*/
type DistributeContentLibraryImageClustersNotFound struct {
	XAocRequestID string

	Payload *models.ErrorBody
}

func (o *DistributeContentLibraryImageClustersNotFound) Error() string {
	return fmt.Sprintf("[POST /distribute-content-library-image-clusters][%d] distributeContentLibraryImageClustersNotFound  %+v", 404, o.Payload)
}
func (o *DistributeContentLibraryImageClustersNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DistributeContentLibraryImageClustersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDistributeContentLibraryImageClustersInternalServerError creates a DistributeContentLibraryImageClustersInternalServerError with default headers values
func NewDistributeContentLibraryImageClustersInternalServerError() *DistributeContentLibraryImageClustersInternalServerError {
	return &DistributeContentLibraryImageClustersInternalServerError{}
}

/* DistributeContentLibraryImageClustersInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type DistributeContentLibraryImageClustersInternalServerError struct {
	XAocRequestID string

	Payload *models.ErrorBody
}

func (o *DistributeContentLibraryImageClustersInternalServerError) Error() string {
	return fmt.Sprintf("[POST /distribute-content-library-image-clusters][%d] distributeContentLibraryImageClustersInternalServerError  %+v", 500, o.Payload)
}
func (o *DistributeContentLibraryImageClustersInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DistributeContentLibraryImageClustersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
