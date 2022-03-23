// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InstanceConfig Instance config
//
// swagger:model InstanceConfig
type InstanceConfig struct {

	// backend check commands
	BackendCheckCommands *InstanceConfigBackendCheckCommands `json:"backend_check_commands,omitempty"`

	// Backend default server options
	BackendDefaultServer *string `json:"backend_default_server,omitempty"`

	// Backend options
	// Example: ["httpchk GET /"]
	BackendOptions []string `json:"backend_options"`

	// Backend servers
	// Required: true
	BackendServers []*BackendServer `json:"backend_servers"`

	// Balance algorithm
	// Required: true
	Balance string `json:"balance"`

	// Frontend network interface
	// Example: eth0
	// Required: true
	// Pattern: ^[A-Za-z0-9\-_.]{1,32}$
	FrontendInterface string `json:"frontend_interface"`

	// Frontend IP address
	// Example: 192.168.1.2
	// Required: true
	// Pattern: ((^\s*((([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]))\s*$)|(^\s*((([0-9A-Fa-f]{1,4}:){7}([0-9A-Fa-f]{1,4}|:))|(([0-9A-Fa-f]{1,4}:){6}(:[0-9A-Fa-f]{1,4}|((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){1,2})|:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){1,3})|((:[0-9A-Fa-f]{1,4})?:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){1,4})|((:[0-9A-Fa-f]{1,4}){0,2}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){1,5})|((:[0-9A-Fa-f]{1,4}){0,3}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){1,6})|((:[0-9A-Fa-f]{1,4}){0,4}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(:(((:[0-9A-Fa-f]{1,4}){1,7})|((:[0-9A-Fa-f]{1,4}){0,5}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:)))(%.+)?\s*$))
	FrontendIpaddress string `json:"frontend_ipaddress"`

	// Frontend network prefix
	// Example: 24
	// Required: true
	// Maximum: 32
	// Minimum: 8
	FrontendNetPrefix uint8 `json:"frontend_net_prefix"`

	// Frontend options
	// Example: ssl
	FrontendOptions *string `json:"frontend_options,omitempty"`

	// Frontend port
	// Example: 443
	// Required: true
	FrontendPort uint16 `json:"frontend_port"`

	// Protocol mode
	// Required: true
	// Enum: [tcp udp http]
	Mode string `json:"mode"`

	// Instance name
	// Required: true
	// Pattern: ^[A-Za-z0-9\-_.]{1,32}$
	Name string `json:"name"`

	// Instance priority
	// Required: true
	// Maximum: 100
	// Minimum: 1
	Priority uint8 `json:"priority"`

	// Instance state
	// Required: true
	// Enum: [MASTER BACKUP]
	State string `json:"state"`
}

// Validate validates this instance config
func (m *InstanceConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackendCheckCommands(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBackendServers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBalance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrontendInterface(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrontendIpaddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrontendNetPrefix(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrontendPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriority(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstanceConfig) validateBackendCheckCommands(formats strfmt.Registry) error {
	if swag.IsZero(m.BackendCheckCommands) { // not required
		return nil
	}

	if m.BackendCheckCommands != nil {
		if err := m.BackendCheckCommands.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backend_check_commands")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("backend_check_commands")
			}
			return err
		}
	}

	return nil
}

func (m *InstanceConfig) validateBackendServers(formats strfmt.Registry) error {

	if err := validate.Required("backend_servers", "body", m.BackendServers); err != nil {
		return err
	}

	for i := 0; i < len(m.BackendServers); i++ {
		if swag.IsZero(m.BackendServers[i]) { // not required
			continue
		}

		if m.BackendServers[i] != nil {
			if err := m.BackendServers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("backend_servers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("backend_servers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InstanceConfig) validateBalance(formats strfmt.Registry) error {

	if err := validate.RequiredString("balance", "body", m.Balance); err != nil {
		return err
	}

	return nil
}

func (m *InstanceConfig) validateFrontendInterface(formats strfmt.Registry) error {

	if err := validate.RequiredString("frontend_interface", "body", m.FrontendInterface); err != nil {
		return err
	}

	if err := validate.Pattern("frontend_interface", "body", m.FrontendInterface, `^[A-Za-z0-9\-_.]{1,32}$`); err != nil {
		return err
	}

	return nil
}

func (m *InstanceConfig) validateFrontendIpaddress(formats strfmt.Registry) error {

	if err := validate.RequiredString("frontend_ipaddress", "body", m.FrontendIpaddress); err != nil {
		return err
	}

	if err := validate.Pattern("frontend_ipaddress", "body", m.FrontendIpaddress, `((^\s*((([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]))\s*$)|(^\s*((([0-9A-Fa-f]{1,4}:){7}([0-9A-Fa-f]{1,4}|:))|(([0-9A-Fa-f]{1,4}:){6}(:[0-9A-Fa-f]{1,4}|((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){1,2})|:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){1,3})|((:[0-9A-Fa-f]{1,4})?:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){1,4})|((:[0-9A-Fa-f]{1,4}){0,2}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){1,5})|((:[0-9A-Fa-f]{1,4}){0,3}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){1,6})|((:[0-9A-Fa-f]{1,4}){0,4}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(:(((:[0-9A-Fa-f]{1,4}){1,7})|((:[0-9A-Fa-f]{1,4}){0,5}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:)))(%.+)?\s*$))`); err != nil {
		return err
	}

	return nil
}

func (m *InstanceConfig) validateFrontendNetPrefix(formats strfmt.Registry) error {

	if err := validate.Required("frontend_net_prefix", "body", uint8(m.FrontendNetPrefix)); err != nil {
		return err
	}

	if err := validate.MinimumUint("frontend_net_prefix", "body", uint64(m.FrontendNetPrefix), 8, false); err != nil {
		return err
	}

	if err := validate.MaximumUint("frontend_net_prefix", "body", uint64(m.FrontendNetPrefix), 32, false); err != nil {
		return err
	}

	return nil
}

func (m *InstanceConfig) validateFrontendPort(formats strfmt.Registry) error {

	if err := validate.Required("frontend_port", "body", uint16(m.FrontendPort)); err != nil {
		return err
	}

	return nil
}

var instanceConfigTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["tcp","udp","http"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		instanceConfigTypeModePropEnum = append(instanceConfigTypeModePropEnum, v)
	}
}

const (

	// InstanceConfigModeTCP captures enum value "tcp"
	InstanceConfigModeTCP string = "tcp"

	// InstanceConfigModeUDP captures enum value "udp"
	InstanceConfigModeUDP string = "udp"

	// InstanceConfigModeHTTP captures enum value "http"
	InstanceConfigModeHTTP string = "http"
)

// prop value enum
func (m *InstanceConfig) validateModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, instanceConfigTypeModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InstanceConfig) validateMode(formats strfmt.Registry) error {

	if err := validate.RequiredString("mode", "body", m.Mode); err != nil {
		return err
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", m.Mode); err != nil {
		return err
	}

	return nil
}

func (m *InstanceConfig) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", m.Name, `^[A-Za-z0-9\-_.]{1,32}$`); err != nil {
		return err
	}

	return nil
}

func (m *InstanceConfig) validatePriority(formats strfmt.Registry) error {

	if err := validate.Required("priority", "body", uint8(m.Priority)); err != nil {
		return err
	}

	if err := validate.MinimumUint("priority", "body", uint64(m.Priority), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumUint("priority", "body", uint64(m.Priority), 100, false); err != nil {
		return err
	}

	return nil
}

var instanceConfigTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["MASTER","BACKUP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		instanceConfigTypeStatePropEnum = append(instanceConfigTypeStatePropEnum, v)
	}
}

const (

	// InstanceConfigStateMASTER captures enum value "MASTER"
	InstanceConfigStateMASTER string = "MASTER"

	// InstanceConfigStateBACKUP captures enum value "BACKUP"
	InstanceConfigStateBACKUP string = "BACKUP"
)

// prop value enum
func (m *InstanceConfig) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, instanceConfigTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InstanceConfig) validateState(formats strfmt.Registry) error {

	if err := validate.RequiredString("state", "body", m.State); err != nil {
		return err
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this instance config based on the context it is used
func (m *InstanceConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBackendCheckCommands(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBackendServers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstanceConfig) contextValidateBackendCheckCommands(ctx context.Context, formats strfmt.Registry) error {

	if m.BackendCheckCommands != nil {
		if err := m.BackendCheckCommands.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backend_check_commands")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("backend_check_commands")
			}
			return err
		}
	}

	return nil
}

func (m *InstanceConfig) contextValidateBackendServers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BackendServers); i++ {

		if m.BackendServers[i] != nil {
			if err := m.BackendServers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("backend_servers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("backend_servers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *InstanceConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstanceConfig) UnmarshalBinary(b []byte) error {
	var res InstanceConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// InstanceConfigBackendCheckCommands Backend check commands
//
// swagger:model InstanceConfigBackendCheckCommands
type InstanceConfigBackendCheckCommands struct {

	// check type
	// Example: http-check
	// Enum: [http-check tcp-check]
	CheckType string `json:"check_type,omitempty"`

	// Backend TCP check commands
	// Example: ["expect status 200"]
	Commands []string `json:"commands"`
}

// Validate validates this instance config backend check commands
func (m *InstanceConfigBackendCheckCommands) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCheckType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var instanceConfigBackendCheckCommandsTypeCheckTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["http-check","tcp-check"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		instanceConfigBackendCheckCommandsTypeCheckTypePropEnum = append(instanceConfigBackendCheckCommandsTypeCheckTypePropEnum, v)
	}
}

const (

	// InstanceConfigBackendCheckCommandsCheckTypeHTTPDashCheck captures enum value "http-check"
	InstanceConfigBackendCheckCommandsCheckTypeHTTPDashCheck string = "http-check"

	// InstanceConfigBackendCheckCommandsCheckTypeTCPDashCheck captures enum value "tcp-check"
	InstanceConfigBackendCheckCommandsCheckTypeTCPDashCheck string = "tcp-check"
)

// prop value enum
func (m *InstanceConfigBackendCheckCommands) validateCheckTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, instanceConfigBackendCheckCommandsTypeCheckTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InstanceConfigBackendCheckCommands) validateCheckType(formats strfmt.Registry) error {
	if swag.IsZero(m.CheckType) { // not required
		return nil
	}

	// value enum
	if err := m.validateCheckTypeEnum("backend_check_commands"+"."+"check_type", "body", m.CheckType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this instance config backend check commands based on context it is used
func (m *InstanceConfigBackendCheckCommands) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InstanceConfigBackendCheckCommands) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstanceConfigBackendCheckCommands) UnmarshalBinary(b []byte) error {
	var res InstanceConfigBackendCheckCommands
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
