// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-operator/pkg/scyllaclient/internal/scylla/models"
)

// CacheServiceMetricsRowEntriesGetReader is a Reader for the CacheServiceMetricsRowEntriesGet structure.
type CacheServiceMetricsRowEntriesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CacheServiceMetricsRowEntriesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCacheServiceMetricsRowEntriesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCacheServiceMetricsRowEntriesGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCacheServiceMetricsRowEntriesGetOK creates a CacheServiceMetricsRowEntriesGetOK with default headers values
func NewCacheServiceMetricsRowEntriesGetOK() *CacheServiceMetricsRowEntriesGetOK {
	return &CacheServiceMetricsRowEntriesGetOK{}
}

/*
CacheServiceMetricsRowEntriesGetOK handles this case with default header values.

CacheServiceMetricsRowEntriesGetOK cache service metrics row entries get o k
*/
type CacheServiceMetricsRowEntriesGetOK struct {
	Payload int32
}

func (o *CacheServiceMetricsRowEntriesGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *CacheServiceMetricsRowEntriesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCacheServiceMetricsRowEntriesGetDefault creates a CacheServiceMetricsRowEntriesGetDefault with default headers values
func NewCacheServiceMetricsRowEntriesGetDefault(code int) *CacheServiceMetricsRowEntriesGetDefault {
	return &CacheServiceMetricsRowEntriesGetDefault{
		_statusCode: code,
	}
}

/*
CacheServiceMetricsRowEntriesGetDefault handles this case with default header values.

internal server error
*/
type CacheServiceMetricsRowEntriesGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the cache service metrics row entries get default response
func (o *CacheServiceMetricsRowEntriesGetDefault) Code() int {
	return o._statusCode
}

func (o *CacheServiceMetricsRowEntriesGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CacheServiceMetricsRowEntriesGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *CacheServiceMetricsRowEntriesGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
