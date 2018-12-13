/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2018 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * AUTHOR: Bob van Luijt (bob@kub.design)
 * See www.creativesoftwarefdn.org for details
 * Contact: @CreativeSofwFdn / bob@kub.design
 */
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

// ToolsMapRequest tools map request
// swagger:model ToolsMapRequest
type ToolsMapRequest struct {

	// Concept that should function as center. Should be one concept (e.g., car) or CamelCased (e.g, MovedTo)
	ConceptCenter string `json:"conceptCenter,omitempty"`

	// Only needs to be set when type is centerOfNetwork and should contain the name of the Weaviate that is taken as center.
	NetworkCenter string `json:"networkCenter,omitempty"`

	// What type of map should be generated?
	// Enum: [zeroPointPosition centerOfSelf centerOfNetwork centerOfConcept]
	Type string `json:"type,omitempty"`
}

// Validate validates this tools map request
func (m *ToolsMapRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var toolsMapRequestTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["zeroPointPosition","centerOfSelf","centerOfNetwork","centerOfConcept"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		toolsMapRequestTypeTypePropEnum = append(toolsMapRequestTypeTypePropEnum, v)
	}
}

const (

	// ToolsMapRequestTypeZeroPointPosition captures enum value "zeroPointPosition"
	ToolsMapRequestTypeZeroPointPosition string = "zeroPointPosition"

	// ToolsMapRequestTypeCenterOfSelf captures enum value "centerOfSelf"
	ToolsMapRequestTypeCenterOfSelf string = "centerOfSelf"

	// ToolsMapRequestTypeCenterOfNetwork captures enum value "centerOfNetwork"
	ToolsMapRequestTypeCenterOfNetwork string = "centerOfNetwork"

	// ToolsMapRequestTypeCenterOfConcept captures enum value "centerOfConcept"
	ToolsMapRequestTypeCenterOfConcept string = "centerOfConcept"
)

// prop value enum
func (m *ToolsMapRequest) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, toolsMapRequestTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ToolsMapRequest) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ToolsMapRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ToolsMapRequest) UnmarshalBinary(b []byte) error {
	var res ToolsMapRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}