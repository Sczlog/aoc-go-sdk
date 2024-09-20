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

// GetVMSnapshotsReader is a Reader for the GetVMSnapshots structure.
type GetVMSnapshotsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVMSnapshotsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVMSnapshotsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetVMSnapshotsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetVMSnapshotsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetVMSnapshotsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVMSnapshotsOK creates a GetVMSnapshotsOK with default headers values
func NewGetVMSnapshotsOK() *GetVMSnapshotsOK {
	return &GetVMSnapshotsOK{}
}

/* GetVMSnapshotsOK describes a response with status code 200, with default header values.

GetVMSnapshotsOK get Vm snapshots o k
*/
type GetVMSnapshotsOK struct {
	XAocRequestID string

	Payload []*models.VMSnapshot
}

func (o *GetVMSnapshotsOK) Error() string {
	return fmt.Sprintf("[POST /get-vm-snapshots][%d] getVmSnapshotsOK  %+v", 200, o.Payload)
}
func (o *GetVMSnapshotsOK) GetPayload() []*models.VMSnapshot {
	return o.Payload
}

func (o *GetVMSnapshotsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetVMSnapshotsBadRequest creates a GetVMSnapshotsBadRequest with default headers values
func NewGetVMSnapshotsBadRequest() *GetVMSnapshotsBadRequest {
	return &GetVMSnapshotsBadRequest{}
}

/* GetVMSnapshotsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetVMSnapshotsBadRequest struct {
	XAocRequestID string

	Payload *models.ErrorBody
}

func (o *GetVMSnapshotsBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-vm-snapshots][%d] getVmSnapshotsBadRequest  %+v", 400, o.Payload)
}
func (o *GetVMSnapshotsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVMSnapshotsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetVMSnapshotsNotFound creates a GetVMSnapshotsNotFound with default headers values
func NewGetVMSnapshotsNotFound() *GetVMSnapshotsNotFound {
	return &GetVMSnapshotsNotFound{}
}

/* GetVMSnapshotsNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetVMSnapshotsNotFound struct {
	XAocRequestID string

	Payload *models.ErrorBody
}

func (o *GetVMSnapshotsNotFound) Error() string {
	return fmt.Sprintf("[POST /get-vm-snapshots][%d] getVmSnapshotsNotFound  %+v", 404, o.Payload)
}
func (o *GetVMSnapshotsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVMSnapshotsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetVMSnapshotsInternalServerError creates a GetVMSnapshotsInternalServerError with default headers values
func NewGetVMSnapshotsInternalServerError() *GetVMSnapshotsInternalServerError {
	return &GetVMSnapshotsInternalServerError{}
}

/* GetVMSnapshotsInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetVMSnapshotsInternalServerError struct {
	XAocRequestID string

	Payload *models.ErrorBody
}

func (o *GetVMSnapshotsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-vm-snapshots][%d] getVmSnapshotsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetVMSnapshotsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVMSnapshotsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
