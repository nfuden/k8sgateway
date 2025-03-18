// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	apiv1alpha1 "github.com/kgateway-dev/kgateway/v2/api/v1alpha1"
)

// HeaderTransformationApplyConfiguration represents a declarative configuration of the HeaderTransformation type for use
// with apply.
type HeaderTransformationApplyConfiguration struct {
	Name  *apiv1alpha1.HeaderName   `json:"name,omitempty"`
	Value *apiv1alpha1.InjaTemplate `json:"value,omitempty"`
}

// HeaderTransformationApplyConfiguration constructs a declarative configuration of the HeaderTransformation type for use with
// apply.
func HeaderTransformation() *HeaderTransformationApplyConfiguration {
	return &HeaderTransformationApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *HeaderTransformationApplyConfiguration) WithName(value apiv1alpha1.HeaderName) *HeaderTransformationApplyConfiguration {
	b.Name = &value
	return b
}

// WithValue sets the Value field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Value field is set to the value of the last call.
func (b *HeaderTransformationApplyConfiguration) WithValue(value apiv1alpha1.InjaTemplate) *HeaderTransformationApplyConfiguration {
	b.Value = &value
	return b
}
