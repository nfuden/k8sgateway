// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RateLimitProviderApplyConfiguration represents a declarative configuration of the RateLimitProvider type for use
// with apply.
type RateLimitProviderApplyConfiguration struct {
	GrpcService *ExtGrpcServiceApplyConfiguration `json:"grpcService,omitempty"`
	Domain      *string                           `json:"domain,omitempty"`
	FailOpen    *bool                             `json:"failOpen,omitempty"`
	Timeout     *v1.Duration                      `json:"timeout,omitempty"`
}

// RateLimitProviderApplyConfiguration constructs a declarative configuration of the RateLimitProvider type for use with
// apply.
func RateLimitProvider() *RateLimitProviderApplyConfiguration {
	return &RateLimitProviderApplyConfiguration{}
}

// WithGrpcService sets the GrpcService field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GrpcService field is set to the value of the last call.
func (b *RateLimitProviderApplyConfiguration) WithGrpcService(value *ExtGrpcServiceApplyConfiguration) *RateLimitProviderApplyConfiguration {
	b.GrpcService = value
	return b
}

// WithDomain sets the Domain field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Domain field is set to the value of the last call.
func (b *RateLimitProviderApplyConfiguration) WithDomain(value string) *RateLimitProviderApplyConfiguration {
	b.Domain = &value
	return b
}

// WithFailOpen sets the FailOpen field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FailOpen field is set to the value of the last call.
func (b *RateLimitProviderApplyConfiguration) WithFailOpen(value bool) *RateLimitProviderApplyConfiguration {
	b.FailOpen = &value
	return b
}

// WithTimeout sets the Timeout field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Timeout field is set to the value of the last call.
func (b *RateLimitProviderApplyConfiguration) WithTimeout(value v1.Duration) *RateLimitProviderApplyConfiguration {
	b.Timeout = &value
	return b
}
