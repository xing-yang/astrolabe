// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TaskInfo task info
// swagger:model TaskInfo
type TaskInfo struct {

	// completed
	// Required: true
	Completed *bool `json:"completed"`

	// details
	Details string `json:"details,omitempty"`

	// finished time
	FinishedTime string `json:"finishedTime,omitempty"`

	// id
	// Required: true
	ID TaskID `json:"id"`

	// progress
	// Required: true
	Progress *int64 `json:"progress"`

	// started time
	// Required: true
	StartedTime *string `json:"startedTime"`

	// status
	// Required: true
	// Enum: [running success failed cancelled]
	Status *string `json:"status"`
}

// Validate validates this task info
func (m *TaskInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompleted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProgress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskInfo) validateCompleted(formats strfmt.Registry) error {

	if err := validate.Required("completed", "body", m.Completed); err != nil {
		return err
	}

	return nil
}

func (m *TaskInfo) validateID(formats strfmt.Registry) error {

	if err := m.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *TaskInfo) validateProgress(formats strfmt.Registry) error {

	if err := validate.Required("progress", "body", m.Progress); err != nil {
		return err
	}

	return nil
}

func (m *TaskInfo) validateStartedTime(formats strfmt.Registry) error {

	if err := validate.Required("startedTime", "body", m.StartedTime); err != nil {
		return err
	}

	return nil
}

var taskInfoTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["running","success","failed","cancelled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		taskInfoTypeStatusPropEnum = append(taskInfoTypeStatusPropEnum, v)
	}
}

const (

	// TaskInfoStatusRunning captures enum value "running"
	TaskInfoStatusRunning string = "running"

	// TaskInfoStatusSuccess captures enum value "success"
	TaskInfoStatusSuccess string = "success"

	// TaskInfoStatusFailed captures enum value "failed"
	TaskInfoStatusFailed string = "failed"

	// TaskInfoStatusCancelled captures enum value "cancelled"
	TaskInfoStatusCancelled string = "cancelled"
)

// prop value enum
func (m *TaskInfo) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, taskInfoTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TaskInfo) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TaskInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskInfo) UnmarshalBinary(b []byte) error {
	var res TaskInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
