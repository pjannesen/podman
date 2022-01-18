// Code generated by go generate; DO NOT EDIT.
package volumes

import (
	"net/url"

	"github.com/containers/podman/v4/pkg/bindings/internal/util"
)

// Changed returns true if named field has been set
func (o *PruneOptions) Changed(fieldName string) bool {
	return util.Changed(o, fieldName)
}

// ToParams formats struct fields to be passed to API service
func (o *PruneOptions) ToParams() (url.Values, error) {
	return util.ToParams(o)
}

// WithFilters set field Filters to given value
func (o *PruneOptions) WithFilters(value map[string][]string) *PruneOptions {
	o.Filters = value
	return o
}

// GetFilters returns value of field Filters
func (o *PruneOptions) GetFilters() map[string][]string {
	if o.Filters == nil {
		var z map[string][]string
		return z
	}
	return o.Filters
}
