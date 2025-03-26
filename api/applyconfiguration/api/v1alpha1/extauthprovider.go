// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "sigs.k8s.io/gateway-api/apis/v1"
)

// ExtAuthProviderApplyConfiguration represents a declarative configuration of the ExtAuthProvider type for use
// with apply.
type ExtAuthProviderApplyConfiguration struct {
	BackendRef *v1.BackendRef `json:"backendRef,omitempty"`
}

// ExtAuthProviderApplyConfiguration constructs a declarative configuration of the ExtAuthProvider type for use with
// apply.
func ExtAuthProvider() *ExtAuthProviderApplyConfiguration {
	return &ExtAuthProviderApplyConfiguration{}
}

// WithBackendRef sets the BackendRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BackendRef field is set to the value of the last call.
func (b *ExtAuthProviderApplyConfiguration) WithBackendRef(value v1.BackendRef) *ExtAuthProviderApplyConfiguration {
	b.BackendRef = &value
	return b
}
