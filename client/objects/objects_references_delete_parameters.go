//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2020 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package objects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/semi-technologies/weaviate/entities/models"
)

// NewObjectsReferencesDeleteParams creates a new ObjectsReferencesDeleteParams object
// with the default values initialized.
func NewObjectsReferencesDeleteParams() *ObjectsReferencesDeleteParams {
	var ()
	return &ObjectsReferencesDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewObjectsReferencesDeleteParamsWithTimeout creates a new ObjectsReferencesDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewObjectsReferencesDeleteParamsWithTimeout(timeout time.Duration) *ObjectsReferencesDeleteParams {
	var ()
	return &ObjectsReferencesDeleteParams{

		timeout: timeout,
	}
}

// NewObjectsReferencesDeleteParamsWithContext creates a new ObjectsReferencesDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewObjectsReferencesDeleteParamsWithContext(ctx context.Context) *ObjectsReferencesDeleteParams {
	var ()
	return &ObjectsReferencesDeleteParams{

		Context: ctx,
	}
}

// NewObjectsReferencesDeleteParamsWithHTTPClient creates a new ObjectsReferencesDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewObjectsReferencesDeleteParamsWithHTTPClient(client *http.Client) *ObjectsReferencesDeleteParams {
	var ()
	return &ObjectsReferencesDeleteParams{
		HTTPClient: client,
	}
}

/*ObjectsReferencesDeleteParams contains all the parameters to send to the API endpoint
for the objects references delete operation typically these are written to a http.Request
*/
type ObjectsReferencesDeleteParams struct {

	/*Body*/
	Body *models.SingleRef
	/*ID
	  Unique ID of the Object.

	*/
	ID strfmt.UUID
	/*PropertyName
	  Unique name of the property related to the Object.

	*/
	PropertyName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the objects references delete params
func (o *ObjectsReferencesDeleteParams) WithTimeout(timeout time.Duration) *ObjectsReferencesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the objects references delete params
func (o *ObjectsReferencesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the objects references delete params
func (o *ObjectsReferencesDeleteParams) WithContext(ctx context.Context) *ObjectsReferencesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the objects references delete params
func (o *ObjectsReferencesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the objects references delete params
func (o *ObjectsReferencesDeleteParams) WithHTTPClient(client *http.Client) *ObjectsReferencesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the objects references delete params
func (o *ObjectsReferencesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the objects references delete params
func (o *ObjectsReferencesDeleteParams) WithBody(body *models.SingleRef) *ObjectsReferencesDeleteParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the objects references delete params
func (o *ObjectsReferencesDeleteParams) SetBody(body *models.SingleRef) {
	o.Body = body
}

// WithID adds the id to the objects references delete params
func (o *ObjectsReferencesDeleteParams) WithID(id strfmt.UUID) *ObjectsReferencesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the objects references delete params
func (o *ObjectsReferencesDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithPropertyName adds the propertyName to the objects references delete params
func (o *ObjectsReferencesDeleteParams) WithPropertyName(propertyName string) *ObjectsReferencesDeleteParams {
	o.SetPropertyName(propertyName)
	return o
}

// SetPropertyName adds the propertyName to the objects references delete params
func (o *ObjectsReferencesDeleteParams) SetPropertyName(propertyName string) {
	o.PropertyName = propertyName
}

// WriteToRequest writes these params to a swagger request
func (o *ObjectsReferencesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// path param propertyName
	if err := r.SetPathParam("propertyName", o.PropertyName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}