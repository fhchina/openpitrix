// Code generated by go-swagger; DO NOT EDIT.

package app_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDescribeAppVersionsParams creates a new DescribeAppVersionsParams object
// with the default values initialized.
func NewDescribeAppVersionsParams() *DescribeAppVersionsParams {
	var ()
	return &DescribeAppVersionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeAppVersionsParamsWithTimeout creates a new DescribeAppVersionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDescribeAppVersionsParamsWithTimeout(timeout time.Duration) *DescribeAppVersionsParams {
	var ()
	return &DescribeAppVersionsParams{

		timeout: timeout,
	}
}

// NewDescribeAppVersionsParamsWithContext creates a new DescribeAppVersionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDescribeAppVersionsParamsWithContext(ctx context.Context) *DescribeAppVersionsParams {
	var ()
	return &DescribeAppVersionsParams{

		Context: ctx,
	}
}

// NewDescribeAppVersionsParamsWithHTTPClient creates a new DescribeAppVersionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDescribeAppVersionsParamsWithHTTPClient(client *http.Client) *DescribeAppVersionsParams {
	var ()
	return &DescribeAppVersionsParams{
		HTTPClient: client,
	}
}

/*DescribeAppVersionsParams contains all the parameters to send to the API endpoint
for the describe app versions operation typically these are written to a http.Request
*/
type DescribeAppVersionsParams struct {

	/*AppID*/
	AppID []string
	/*Description*/
	Description []string
	/*DisplayColumns*/
	DisplayColumns []string
	/*Limit
	  default is 20, max value is 200.

	*/
	Limit *int64
	/*Name*/
	Name []string
	/*Offset
	  default is 0.

	*/
	Offset *int64
	/*OwnerPath*/
	OwnerPath []string
	/*PackageName*/
	PackageName []string
	/*Reverse*/
	Reverse *bool
	/*SearchWord*/
	SearchWord *string
	/*SortKey*/
	SortKey *string
	/*Status*/
	Status []string
	/*Type*/
	Type []string
	/*VersionID*/
	VersionID []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the describe app versions params
func (o *DescribeAppVersionsParams) WithTimeout(timeout time.Duration) *DescribeAppVersionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe app versions params
func (o *DescribeAppVersionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe app versions params
func (o *DescribeAppVersionsParams) WithContext(ctx context.Context) *DescribeAppVersionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe app versions params
func (o *DescribeAppVersionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe app versions params
func (o *DescribeAppVersionsParams) WithHTTPClient(client *http.Client) *DescribeAppVersionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe app versions params
func (o *DescribeAppVersionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the describe app versions params
func (o *DescribeAppVersionsParams) WithAppID(appID []string) *DescribeAppVersionsParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the describe app versions params
func (o *DescribeAppVersionsParams) SetAppID(appID []string) {
	o.AppID = appID
}

// WithDescription adds the description to the describe app versions params
func (o *DescribeAppVersionsParams) WithDescription(description []string) *DescribeAppVersionsParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the describe app versions params
func (o *DescribeAppVersionsParams) SetDescription(description []string) {
	o.Description = description
}

// WithDisplayColumns adds the displayColumns to the describe app versions params
func (o *DescribeAppVersionsParams) WithDisplayColumns(displayColumns []string) *DescribeAppVersionsParams {
	o.SetDisplayColumns(displayColumns)
	return o
}

// SetDisplayColumns adds the displayColumns to the describe app versions params
func (o *DescribeAppVersionsParams) SetDisplayColumns(displayColumns []string) {
	o.DisplayColumns = displayColumns
}

// WithLimit adds the limit to the describe app versions params
func (o *DescribeAppVersionsParams) WithLimit(limit *int64) *DescribeAppVersionsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the describe app versions params
func (o *DescribeAppVersionsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the describe app versions params
func (o *DescribeAppVersionsParams) WithName(name []string) *DescribeAppVersionsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the describe app versions params
func (o *DescribeAppVersionsParams) SetName(name []string) {
	o.Name = name
}

// WithOffset adds the offset to the describe app versions params
func (o *DescribeAppVersionsParams) WithOffset(offset *int64) *DescribeAppVersionsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the describe app versions params
func (o *DescribeAppVersionsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOwnerPath adds the ownerPath to the describe app versions params
func (o *DescribeAppVersionsParams) WithOwnerPath(ownerPath []string) *DescribeAppVersionsParams {
	o.SetOwnerPath(ownerPath)
	return o
}

// SetOwnerPath adds the ownerPath to the describe app versions params
func (o *DescribeAppVersionsParams) SetOwnerPath(ownerPath []string) {
	o.OwnerPath = ownerPath
}

// WithPackageName adds the packageName to the describe app versions params
func (o *DescribeAppVersionsParams) WithPackageName(packageName []string) *DescribeAppVersionsParams {
	o.SetPackageName(packageName)
	return o
}

// SetPackageName adds the packageName to the describe app versions params
func (o *DescribeAppVersionsParams) SetPackageName(packageName []string) {
	o.PackageName = packageName
}

// WithReverse adds the reverse to the describe app versions params
func (o *DescribeAppVersionsParams) WithReverse(reverse *bool) *DescribeAppVersionsParams {
	o.SetReverse(reverse)
	return o
}

// SetReverse adds the reverse to the describe app versions params
func (o *DescribeAppVersionsParams) SetReverse(reverse *bool) {
	o.Reverse = reverse
}

// WithSearchWord adds the searchWord to the describe app versions params
func (o *DescribeAppVersionsParams) WithSearchWord(searchWord *string) *DescribeAppVersionsParams {
	o.SetSearchWord(searchWord)
	return o
}

// SetSearchWord adds the searchWord to the describe app versions params
func (o *DescribeAppVersionsParams) SetSearchWord(searchWord *string) {
	o.SearchWord = searchWord
}

// WithSortKey adds the sortKey to the describe app versions params
func (o *DescribeAppVersionsParams) WithSortKey(sortKey *string) *DescribeAppVersionsParams {
	o.SetSortKey(sortKey)
	return o
}

// SetSortKey adds the sortKey to the describe app versions params
func (o *DescribeAppVersionsParams) SetSortKey(sortKey *string) {
	o.SortKey = sortKey
}

// WithStatus adds the status to the describe app versions params
func (o *DescribeAppVersionsParams) WithStatus(status []string) *DescribeAppVersionsParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the describe app versions params
func (o *DescribeAppVersionsParams) SetStatus(status []string) {
	o.Status = status
}

// WithType adds the typeVar to the describe app versions params
func (o *DescribeAppVersionsParams) WithType(typeVar []string) *DescribeAppVersionsParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the describe app versions params
func (o *DescribeAppVersionsParams) SetType(typeVar []string) {
	o.Type = typeVar
}

// WithVersionID adds the versionID to the describe app versions params
func (o *DescribeAppVersionsParams) WithVersionID(versionID []string) *DescribeAppVersionsParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the describe app versions params
func (o *DescribeAppVersionsParams) SetVersionID(versionID []string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeAppVersionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesAppID := o.AppID

	joinedAppID := swag.JoinByFormat(valuesAppID, "multi")
	// query array param app_id
	if err := r.SetQueryParam("app_id", joinedAppID...); err != nil {
		return err
	}

	valuesDescription := o.Description

	joinedDescription := swag.JoinByFormat(valuesDescription, "multi")
	// query array param description
	if err := r.SetQueryParam("description", joinedDescription...); err != nil {
		return err
	}

	valuesDisplayColumns := o.DisplayColumns

	joinedDisplayColumns := swag.JoinByFormat(valuesDisplayColumns, "multi")
	// query array param display_columns
	if err := r.SetQueryParam("display_columns", joinedDisplayColumns...); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	valuesName := o.Name

	joinedName := swag.JoinByFormat(valuesName, "multi")
	// query array param name
	if err := r.SetQueryParam("name", joinedName...); err != nil {
		return err
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	valuesOwnerPath := o.OwnerPath

	joinedOwnerPath := swag.JoinByFormat(valuesOwnerPath, "multi")
	// query array param owner_path
	if err := r.SetQueryParam("owner_path", joinedOwnerPath...); err != nil {
		return err
	}

	valuesPackageName := o.PackageName

	joinedPackageName := swag.JoinByFormat(valuesPackageName, "multi")
	// query array param package_name
	if err := r.SetQueryParam("package_name", joinedPackageName...); err != nil {
		return err
	}

	if o.Reverse != nil {

		// query param reverse
		var qrReverse bool
		if o.Reverse != nil {
			qrReverse = *o.Reverse
		}
		qReverse := swag.FormatBool(qrReverse)
		if qReverse != "" {
			if err := r.SetQueryParam("reverse", qReverse); err != nil {
				return err
			}
		}

	}

	if o.SearchWord != nil {

		// query param search_word
		var qrSearchWord string
		if o.SearchWord != nil {
			qrSearchWord = *o.SearchWord
		}
		qSearchWord := qrSearchWord
		if qSearchWord != "" {
			if err := r.SetQueryParam("search_word", qSearchWord); err != nil {
				return err
			}
		}

	}

	if o.SortKey != nil {

		// query param sort_key
		var qrSortKey string
		if o.SortKey != nil {
			qrSortKey = *o.SortKey
		}
		qSortKey := qrSortKey
		if qSortKey != "" {
			if err := r.SetQueryParam("sort_key", qSortKey); err != nil {
				return err
			}
		}

	}

	valuesStatus := o.Status

	joinedStatus := swag.JoinByFormat(valuesStatus, "multi")
	// query array param status
	if err := r.SetQueryParam("status", joinedStatus...); err != nil {
		return err
	}

	valuesType := o.Type

	joinedType := swag.JoinByFormat(valuesType, "multi")
	// query array param type
	if err := r.SetQueryParam("type", joinedType...); err != nil {
		return err
	}

	valuesVersionID := o.VersionID

	joinedVersionID := swag.JoinByFormat(valuesVersionID, "multi")
	// query array param version_id
	if err := r.SetQueryParam("version_id", joinedVersionID...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
