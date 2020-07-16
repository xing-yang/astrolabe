// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware-tanzu/astrolabe/gen/models"
)

// ListTaskNexusReader is a Reader for the ListTaskNexus structure.
type ListTaskNexusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListTaskNexusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListTaskNexusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListTaskNexusOK creates a ListTaskNexusOK with default headers values
func NewListTaskNexusOK() *ListTaskNexusOK {
	return &ListTaskNexusOK{}
}

/*ListTaskNexusOK handles this case with default header values.

Task nexus list
*/
type ListTaskNexusOK struct {
	Payload models.TaskNexusList
}

func (o *ListTaskNexusOK) Error() string {
	return fmt.Sprintf("[GET /astrolabe/tasks/nexus][%d] listTaskNexusOK  %+v", 200, o.Payload)
}

func (o *ListTaskNexusOK) GetPayload() models.TaskNexusList {
	return o.Payload
}

func (o *ListTaskNexusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}