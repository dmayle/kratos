// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ory/kratos/internal/httpclient/models"
)

// DeleteIdentityReader is a Reader for the DeleteIdentity structure.
type DeleteIdentityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteIdentityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteIdentityNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteIdentityNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteIdentityInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteIdentityNoContent creates a DeleteIdentityNoContent with default headers values
func NewDeleteIdentityNoContent() *DeleteIdentityNoContent {
	return &DeleteIdentityNoContent{}
}

/*DeleteIdentityNoContent handles this case with default header values.

Empty responses are sent when, for example, resources are deleted. The HTTP status code for empty responses is
typically 201.
*/
type DeleteIdentityNoContent struct {
}

func (o *DeleteIdentityNoContent) Error() string {
	return fmt.Sprintf("[DELETE /identities/{id}][%d] deleteIdentityNoContent ", 204)
}

func (o *DeleteIdentityNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIdentityNotFound creates a DeleteIdentityNotFound with default headers values
func NewDeleteIdentityNotFound() *DeleteIdentityNotFound {
	return &DeleteIdentityNotFound{}
}

/*DeleteIdentityNotFound handles this case with default header values.

genericError
*/
type DeleteIdentityNotFound struct {
	Payload *models.GenericError
}

func (o *DeleteIdentityNotFound) Error() string {
	return fmt.Sprintf("[DELETE /identities/{id}][%d] deleteIdentityNotFound  %+v", 404, o.Payload)
}

func (o *DeleteIdentityNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *DeleteIdentityNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityInternalServerError creates a DeleteIdentityInternalServerError with default headers values
func NewDeleteIdentityInternalServerError() *DeleteIdentityInternalServerError {
	return &DeleteIdentityInternalServerError{}
}

/*DeleteIdentityInternalServerError handles this case with default header values.

genericError
*/
type DeleteIdentityInternalServerError struct {
	Payload *models.GenericError
}

func (o *DeleteIdentityInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /identities/{id}][%d] deleteIdentityInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteIdentityInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *DeleteIdentityInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
