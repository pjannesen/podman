// Code generated by go generate; DO NOT EDIT.
package network

import (
	"net"
	"net/url"

	"github.com/containers/podman/v4/pkg/bindings/internal/util"
)

// Changed returns true if named field has been set
func (o *CreateOptions) Changed(fieldName string) bool {
	return util.Changed(o, fieldName)
}

// ToParams formats struct fields to be passed to API service
func (o *CreateOptions) ToParams() (url.Values, error) {
	return util.ToParams(o)
}

// WithDisableDNS set field DisableDNS to given value
func (o *CreateOptions) WithDisableDNS(value bool) *CreateOptions {
	o.DisableDNS = &value
	return o
}

// GetDisableDNS returns value of field DisableDNS
func (o *CreateOptions) GetDisableDNS() bool {
	if o.DisableDNS == nil {
		var z bool
		return z
	}
	return *o.DisableDNS
}

// WithDriver set field Driver to given value
func (o *CreateOptions) WithDriver(value string) *CreateOptions {
	o.Driver = &value
	return o
}

// GetDriver returns value of field Driver
func (o *CreateOptions) GetDriver() string {
	if o.Driver == nil {
		var z string
		return z
	}
	return *o.Driver
}

// WithGateway set field Gateway to given value
func (o *CreateOptions) WithGateway(value net.IP) *CreateOptions {
	o.Gateway = &value
	return o
}

// GetGateway returns value of field Gateway
func (o *CreateOptions) GetGateway() net.IP {
	if o.Gateway == nil {
		var z net.IP
		return z
	}
	return *o.Gateway
}

// WithInternal set field Internal to given value
func (o *CreateOptions) WithInternal(value bool) *CreateOptions {
	o.Internal = &value
	return o
}

// GetInternal returns value of field Internal
func (o *CreateOptions) GetInternal() bool {
	if o.Internal == nil {
		var z bool
		return z
	}
	return *o.Internal
}

// WithLabels set field Labels to given value
func (o *CreateOptions) WithLabels(value map[string]string) *CreateOptions {
	o.Labels = value
	return o
}

// GetLabels returns value of field Labels
func (o *CreateOptions) GetLabels() map[string]string {
	if o.Labels == nil {
		var z map[string]string
		return z
	}
	return o.Labels
}

// WithMacVLAN set field MacVLAN to given value
func (o *CreateOptions) WithMacVLAN(value string) *CreateOptions {
	o.MacVLAN = &value
	return o
}

// GetMacVLAN returns value of field MacVLAN
func (o *CreateOptions) GetMacVLAN() string {
	if o.MacVLAN == nil {
		var z string
		return z
	}
	return *o.MacVLAN
}

// WithIPRange set field IPRange to given value
func (o *CreateOptions) WithIPRange(value net.IPNet) *CreateOptions {
	o.IPRange = &value
	return o
}

// GetIPRange returns value of field IPRange
func (o *CreateOptions) GetIPRange() net.IPNet {
	if o.IPRange == nil {
		var z net.IPNet
		return z
	}
	return *o.IPRange
}

// WithSubnet set field Subnet to given value
func (o *CreateOptions) WithSubnet(value net.IPNet) *CreateOptions {
	o.Subnet = &value
	return o
}

// GetSubnet returns value of field Subnet
func (o *CreateOptions) GetSubnet() net.IPNet {
	if o.Subnet == nil {
		var z net.IPNet
		return z
	}
	return *o.Subnet
}

// WithIPv6 set field IPv6 to given value
func (o *CreateOptions) WithIPv6(value bool) *CreateOptions {
	o.IPv6 = &value
	return o
}

// GetIPv6 returns value of field IPv6
func (o *CreateOptions) GetIPv6() bool {
	if o.IPv6 == nil {
		var z bool
		return z
	}
	return *o.IPv6
}

// WithOptions set field Options to given value
func (o *CreateOptions) WithOptions(value map[string]string) *CreateOptions {
	o.Options = value
	return o
}

// GetOptions returns value of field Options
func (o *CreateOptions) GetOptions() map[string]string {
	if o.Options == nil {
		var z map[string]string
		return z
	}
	return o.Options
}

// WithName set field Name to given value
func (o *CreateOptions) WithName(value string) *CreateOptions {
	o.Name = &value
	return o
}

// GetName returns value of field Name
func (o *CreateOptions) GetName() string {
	if o.Name == nil {
		var z string
		return z
	}
	return *o.Name
}
