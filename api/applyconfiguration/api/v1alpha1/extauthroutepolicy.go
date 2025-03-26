// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"

	apiv1alpha1 "github.com/kgateway-dev/kgateway/v2/api/v1alpha1"
)

// ExtAuthRoutePolicyApplyConfiguration represents a declarative configuration of the ExtAuthRoutePolicy type for use
// with apply.
type ExtAuthRoutePolicyApplyConfiguration struct {
	ExtensionRef              *v1.LocalObjectReference          `json:"extensionRef,omitempty"`
	Enablement                *apiv1alpha1.ExtAuthEnabled       `json:"enablement,omitempty"`
	FailureModeAllow          *bool                             `json:"failureModeAllow,omitempty"`
	WithRequestBody           *BufferSettingsApplyConfiguration `json:"withRequestBody,omitempty"`
	ClearRouteCache           *bool                             `json:"clearRouteCache,omitempty"`
	MetadataContextNamespaces []string                          `json:"metadataContextNamespaces,omitempty"`
	IncludePeerCertificate    *bool                             `json:"includePeerCertificate,omitempty"`
	IncludeTLSSession         *bool                             `json:"includeTLSSession,omitempty"`
	EmitFilterStateStats      *bool                             `json:"emitFilterStateStats,omitempty"`
}

// ExtAuthRoutePolicyApplyConfiguration constructs a declarative configuration of the ExtAuthRoutePolicy type for use with
// apply.
func ExtAuthRoutePolicy() *ExtAuthRoutePolicyApplyConfiguration {
	return &ExtAuthRoutePolicyApplyConfiguration{}
}

// WithExtensionRef sets the ExtensionRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ExtensionRef field is set to the value of the last call.
func (b *ExtAuthRoutePolicyApplyConfiguration) WithExtensionRef(value v1.LocalObjectReference) *ExtAuthRoutePolicyApplyConfiguration {
	b.ExtensionRef = &value
	return b
}

// WithEnablement sets the Enablement field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Enablement field is set to the value of the last call.
func (b *ExtAuthRoutePolicyApplyConfiguration) WithEnablement(value apiv1alpha1.ExtAuthEnabled) *ExtAuthRoutePolicyApplyConfiguration {
	b.Enablement = &value
	return b
}

// WithFailureModeAllow sets the FailureModeAllow field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FailureModeAllow field is set to the value of the last call.
func (b *ExtAuthRoutePolicyApplyConfiguration) WithFailureModeAllow(value bool) *ExtAuthRoutePolicyApplyConfiguration {
	b.FailureModeAllow = &value
	return b
}

// WithWithRequestBody sets the WithRequestBody field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the WithRequestBody field is set to the value of the last call.
func (b *ExtAuthRoutePolicyApplyConfiguration) WithWithRequestBody(value *BufferSettingsApplyConfiguration) *ExtAuthRoutePolicyApplyConfiguration {
	b.WithRequestBody = value
	return b
}

// WithClearRouteCache sets the ClearRouteCache field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ClearRouteCache field is set to the value of the last call.
func (b *ExtAuthRoutePolicyApplyConfiguration) WithClearRouteCache(value bool) *ExtAuthRoutePolicyApplyConfiguration {
	b.ClearRouteCache = &value
	return b
}

// WithMetadataContextNamespaces adds the given value to the MetadataContextNamespaces field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the MetadataContextNamespaces field.
func (b *ExtAuthRoutePolicyApplyConfiguration) WithMetadataContextNamespaces(values ...string) *ExtAuthRoutePolicyApplyConfiguration {
	for i := range values {
		b.MetadataContextNamespaces = append(b.MetadataContextNamespaces, values[i])
	}
	return b
}

// WithIncludePeerCertificate sets the IncludePeerCertificate field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IncludePeerCertificate field is set to the value of the last call.
func (b *ExtAuthRoutePolicyApplyConfiguration) WithIncludePeerCertificate(value bool) *ExtAuthRoutePolicyApplyConfiguration {
	b.IncludePeerCertificate = &value
	return b
}

// WithIncludeTLSSession sets the IncludeTLSSession field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IncludeTLSSession field is set to the value of the last call.
func (b *ExtAuthRoutePolicyApplyConfiguration) WithIncludeTLSSession(value bool) *ExtAuthRoutePolicyApplyConfiguration {
	b.IncludeTLSSession = &value
	return b
}

// WithEmitFilterStateStats sets the EmitFilterStateStats field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EmitFilterStateStats field is set to the value of the last call.
func (b *ExtAuthRoutePolicyApplyConfiguration) WithEmitFilterStateStats(value bool) *ExtAuthRoutePolicyApplyConfiguration {
	b.EmitFilterStateStats = &value
	return b
}
