// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RunList run list
// swagger:model runList
type RunList []*RunListItems0

// Validate validates this run list
func (m RunList) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {
		if swag.IsZero(m[i]) { // not required
			continue
		}

		if m[i] != nil {
			if err := m[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// RunListItems0 run list items0
// swagger:model RunListItems0
type RunListItems0 struct {

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// finished at
	// Format: date-time
	FinishedAt *strfmt.DateTime `json:"finished_at,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// settings
	Settings *RunListItems0Settings `json:"settings,omitempty"`

	// started at
	// Format: date-time
	StartedAt *strfmt.DateTime `json:"started_at,omitempty"`

	// status
	Status *bool `json:"status,omitempty"`

	// task
	Task *RunListItems0Task `json:"task,omitempty"`
}

// Validate validates this run list items0
func (m *RunListItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFinishedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTask(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RunListItems0) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RunListItems0) validateFinishedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.FinishedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("finished_at", "body", "date-time", m.FinishedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RunListItems0) validateSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.Settings) { // not required
		return nil
	}

	if m.Settings != nil {
		if err := m.Settings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("settings")
			}
			return err
		}
	}

	return nil
}

func (m *RunListItems0) validateStartedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.StartedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("started_at", "body", "date-time", m.StartedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RunListItems0) validateTask(formats strfmt.Registry) error {

	if swag.IsZero(m.Task) { // not required
		return nil
	}

	if m.Task != nil {
		if err := m.Task.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("task")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RunListItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RunListItems0) UnmarshalBinary(b []byte) error {
	var res RunListItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RunListItems0Settings run list items0 settings
// swagger:model RunListItems0Settings
type RunListItems0Settings struct {

	// command
	Command []string `json:"command"`

	// image
	Image string `json:"image,omitempty"`

	// metadata
	Metadata interface{} `json:"metadata,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// queue
	Queue string `json:"queue,omitempty"`

	// timeout
	Timeout int64 `json:"timeout,omitempty"`
}

// Validate validates this run list items0 settings
func (m *RunListItems0Settings) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RunListItems0Settings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RunListItems0Settings) UnmarshalBinary(b []byte) error {
	var res RunListItems0Settings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RunListItems0Task run list items0 task
// swagger:model RunListItems0Task
type RunListItems0Task struct {

	// base sha
	BaseSha string `json:"base_sha,omitempty"`

	// canceled
	Canceled bool `json:"canceled,omitempty"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// finished at
	// Format: date-time
	FinishedAt *strfmt.DateTime `json:"finished_at,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// parent
	Parent *RunListItems0TaskParent `json:"parent,omitempty"`

	// path
	Path string `json:"path,omitempty"`

	// pull request id
	PullRequestID int64 `json:"pull_request_id,omitempty"`

	// ref
	Ref *RunListItems0TaskRef `json:"ref,omitempty"`

	// settings
	Settings *RunListItems0TaskSettings `json:"settings,omitempty"`

	// started at
	// Format: date-time
	StartedAt *strfmt.DateTime `json:"started_at,omitempty"`

	// status
	Status *bool `json:"status,omitempty"`
}

// Validate validates this run list items0 task
func (m *RunListItems0Task) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFinishedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RunListItems0Task) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("task"+"."+"created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RunListItems0Task) validateFinishedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.FinishedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("task"+"."+"finished_at", "body", "date-time", m.FinishedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RunListItems0Task) validateParent(formats strfmt.Registry) error {

	if swag.IsZero(m.Parent) { // not required
		return nil
	}

	if m.Parent != nil {
		if err := m.Parent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("task" + "." + "parent")
			}
			return err
		}
	}

	return nil
}

func (m *RunListItems0Task) validateRef(formats strfmt.Registry) error {

	if swag.IsZero(m.Ref) { // not required
		return nil
	}

	if m.Ref != nil {
		if err := m.Ref.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("task" + "." + "ref")
			}
			return err
		}
	}

	return nil
}

func (m *RunListItems0Task) validateSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.Settings) { // not required
		return nil
	}

	if m.Settings != nil {
		if err := m.Settings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("task" + "." + "settings")
			}
			return err
		}
	}

	return nil
}

func (m *RunListItems0Task) validateStartedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.StartedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("task"+"."+"started_at", "body", "date-time", m.StartedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RunListItems0Task) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RunListItems0Task) UnmarshalBinary(b []byte) error {
	var res RunListItems0Task
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RunListItems0TaskParent run list items0 task parent
// swagger:model RunListItems0TaskParent
type RunListItems0TaskParent struct {

	// auto created
	AutoCreated bool `json:"auto_created,omitempty"`

	// disabled
	Disabled bool `json:"disabled,omitempty"`

	// github
	Github interface{} `json:"github,omitempty"`

	// hook secret
	HookSecret string `json:"hook_secret,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// private
	Private bool `json:"private,omitempty"`
}

// Validate validates this run list items0 task parent
func (m *RunListItems0TaskParent) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RunListItems0TaskParent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RunListItems0TaskParent) UnmarshalBinary(b []byte) error {
	var res RunListItems0TaskParent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RunListItems0TaskRef run list items0 task ref
// swagger:model RunListItems0TaskRef
type RunListItems0TaskRef struct {

	// id
	ID int64 `json:"id,omitempty"`

	// ref name
	RefName string `json:"ref_name,omitempty"`

	// repository
	Repository *RunListItems0TaskRefRepository `json:"repository,omitempty"`

	// sha
	Sha string `json:"sha,omitempty"`
}

// Validate validates this run list items0 task ref
func (m *RunListItems0TaskRef) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRepository(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RunListItems0TaskRef) validateRepository(formats strfmt.Registry) error {

	if swag.IsZero(m.Repository) { // not required
		return nil
	}

	if m.Repository != nil {
		if err := m.Repository.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("task" + "." + "ref" + "." + "repository")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RunListItems0TaskRef) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RunListItems0TaskRef) UnmarshalBinary(b []byte) error {
	var res RunListItems0TaskRef
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RunListItems0TaskRefRepository run list items0 task ref repository
// swagger:model RunListItems0TaskRefRepository
type RunListItems0TaskRefRepository struct {

	// auto created
	AutoCreated bool `json:"auto_created,omitempty"`

	// disabled
	Disabled bool `json:"disabled,omitempty"`

	// github
	Github interface{} `json:"github,omitempty"`

	// hook secret
	HookSecret string `json:"hook_secret,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// private
	Private bool `json:"private,omitempty"`
}

// Validate validates this run list items0 task ref repository
func (m *RunListItems0TaskRefRepository) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RunListItems0TaskRefRepository) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RunListItems0TaskRefRepository) UnmarshalBinary(b []byte) error {
	var res RunListItems0TaskRefRepository
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RunListItems0TaskSettings run list items0 task settings
// swagger:model RunListItems0TaskSettings
type RunListItems0TaskSettings struct {

	// config
	Config *RunListItems0TaskSettingsConfig `json:"config,omitempty"`

	// default image
	DefaultImage string `json:"default_image,omitempty"`

	// default queue
	DefaultQueue string `json:"default_queue,omitempty"`

	// the default timeout; in nanoseconds
	DefaultTimeout int64 `json:"default_timeout,omitempty"`

	// env
	Env []string `json:"env"`

	// metadata
	Metadata interface{} `json:"metadata,omitempty"`

	// mountpoint
	Mountpoint string `json:"mountpoint,omitempty"`

	// runs
	Runs map[string]RunListItems0TaskSettingsRunsAnon `json:"runs,omitempty"`

	// workdir
	Workdir string `json:"workdir,omitempty"`
}

// Validate validates this run list items0 task settings
func (m *RunListItems0TaskSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuns(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RunListItems0TaskSettings) validateConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.Config) { // not required
		return nil
	}

	if m.Config != nil {
		if err := m.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("task" + "." + "settings" + "." + "config")
			}
			return err
		}
	}

	return nil
}

func (m *RunListItems0TaskSettings) validateRuns(formats strfmt.Registry) error {

	if swag.IsZero(m.Runs) { // not required
		return nil
	}

	for k := range m.Runs {

		if swag.IsZero(m.Runs[k]) { // not required
			continue
		}
		if val, ok := m.Runs[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RunListItems0TaskSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RunListItems0TaskSettings) UnmarshalBinary(b []byte) error {
	var res RunListItems0TaskSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RunListItems0TaskSettingsConfig run list items0 task settings config
// swagger:model RunListItems0TaskSettingsConfig
type RunListItems0TaskSettingsConfig struct {

	// global timeout
	GlobalTimeout int64 `json:"global_timeout,omitempty"`

	// override queue
	OverrideQueue bool `json:"override_queue,omitempty"`

	// override timeout
	OverrideTimeout bool `json:"override_timeout,omitempty"`

	// queue
	Queue string `json:"queue,omitempty"`

	// workdir
	Workdir string `json:"workdir,omitempty"`
}

// Validate validates this run list items0 task settings config
func (m *RunListItems0TaskSettingsConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RunListItems0TaskSettingsConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RunListItems0TaskSettingsConfig) UnmarshalBinary(b []byte) error {
	var res RunListItems0TaskSettingsConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RunListItems0TaskSettingsRunsAnon run list items0 task settings runs anon
// swagger:model RunListItems0TaskSettingsRunsAnon
type RunListItems0TaskSettingsRunsAnon struct {

	// command
	Command []string `json:"command"`

	// image
	Image string `json:"image,omitempty"`

	// metadata
	Metadata interface{} `json:"metadata,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// queue
	Queue string `json:"queue,omitempty"`

	// timeout
	Timeout int64 `json:"timeout,omitempty"`
}

// Validate validates this run list items0 task settings runs anon
func (m *RunListItems0TaskSettingsRunsAnon) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RunListItems0TaskSettingsRunsAnon) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RunListItems0TaskSettingsRunsAnon) UnmarshalBinary(b []byte) error {
	var res RunListItems0TaskSettingsRunsAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
