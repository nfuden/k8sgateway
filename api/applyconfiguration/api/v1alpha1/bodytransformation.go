// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	apiv1alpha1 "github.com/kgateway-dev/kgateway/v2/api/v1alpha1"
)

// BodyTransformationApplyConfiguration represents a declarative configuration of the BodyTransformation type for use
// with apply.
type BodyTransformationApplyConfiguration struct {
	ParseAs *apiv1alpha1.BodyParseBehavior `json:"parseAs,omitempty"`
	Value   *apiv1alpha1.InjaTemplate      `json:"value,omitempty"`
}

// BodyTransformationApplyConfiguration constructs a declarative configuration of the BodyTransformation type for use with
// apply.
func BodyTransformation() *BodyTransformationApplyConfiguration {
	return &BodyTransformationApplyConfiguration{}
}

// WithParseAs sets the ParseAs field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ParseAs field is set to the value of the last call.
func (b *BodyTransformationApplyConfiguration) WithParseAs(value apiv1alpha1.BodyParseBehavior) *BodyTransformationApplyConfiguration {
	b.ParseAs = &value
	return b
}

// WithValue sets the Value field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Value field is set to the value of the last call.
func (b *BodyTransformationApplyConfiguration) WithValue(value apiv1alpha1.InjaTemplate) *BodyTransformationApplyConfiguration {
	b.Value = &value
	return b
}
