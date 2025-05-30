// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
)

// IstioContainerApplyConfiguration represents a declarative configuration of the IstioContainer type for use
// with apply.
type IstioContainerApplyConfiguration struct {
	Image                 *ImageApplyConfiguration `json:"image,omitempty"`
	SecurityContext       *v1.SecurityContext      `json:"securityContext,omitempty"`
	Resources             *v1.ResourceRequirements `json:"resources,omitempty"`
	LogLevel              *string                  `json:"logLevel,omitempty"`
	IstioDiscoveryAddress *string                  `json:"istioDiscoveryAddress,omitempty"`
	IstioMetaMeshId       *string                  `json:"istioMetaMeshId,omitempty"`
	IstioMetaClusterId    *string                  `json:"istioMetaClusterId,omitempty"`
}

// IstioContainerApplyConfiguration constructs a declarative configuration of the IstioContainer type for use with
// apply.
func IstioContainer() *IstioContainerApplyConfiguration {
	return &IstioContainerApplyConfiguration{}
}

// WithImage sets the Image field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Image field is set to the value of the last call.
func (b *IstioContainerApplyConfiguration) WithImage(value *ImageApplyConfiguration) *IstioContainerApplyConfiguration {
	b.Image = value
	return b
}

// WithSecurityContext sets the SecurityContext field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecurityContext field is set to the value of the last call.
func (b *IstioContainerApplyConfiguration) WithSecurityContext(value v1.SecurityContext) *IstioContainerApplyConfiguration {
	b.SecurityContext = &value
	return b
}

// WithResources sets the Resources field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Resources field is set to the value of the last call.
func (b *IstioContainerApplyConfiguration) WithResources(value v1.ResourceRequirements) *IstioContainerApplyConfiguration {
	b.Resources = &value
	return b
}

// WithLogLevel sets the LogLevel field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LogLevel field is set to the value of the last call.
func (b *IstioContainerApplyConfiguration) WithLogLevel(value string) *IstioContainerApplyConfiguration {
	b.LogLevel = &value
	return b
}

// WithIstioDiscoveryAddress sets the IstioDiscoveryAddress field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IstioDiscoveryAddress field is set to the value of the last call.
func (b *IstioContainerApplyConfiguration) WithIstioDiscoveryAddress(value string) *IstioContainerApplyConfiguration {
	b.IstioDiscoveryAddress = &value
	return b
}

// WithIstioMetaMeshId sets the IstioMetaMeshId field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IstioMetaMeshId field is set to the value of the last call.
func (b *IstioContainerApplyConfiguration) WithIstioMetaMeshId(value string) *IstioContainerApplyConfiguration {
	b.IstioMetaMeshId = &value
	return b
}

// WithIstioMetaClusterId sets the IstioMetaClusterId field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IstioMetaClusterId field is set to the value of the last call.
func (b *IstioContainerApplyConfiguration) WithIstioMetaClusterId(value string) *IstioContainerApplyConfiguration {
	b.IstioMetaClusterId = &value
	return b
}
