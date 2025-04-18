// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// ModerationApplyConfiguration represents a declarative configuration of the Moderation type for use
// with apply.
type ModerationApplyConfiguration struct {
	OpenAIModeration *OpenAIConfigApplyConfiguration `json:"openAIModeration,omitempty"`
}

// ModerationApplyConfiguration constructs a declarative configuration of the Moderation type for use with
// apply.
func Moderation() *ModerationApplyConfiguration {
	return &ModerationApplyConfiguration{}
}

// WithOpenAIModeration sets the OpenAIModeration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OpenAIModeration field is set to the value of the last call.
func (b *ModerationApplyConfiguration) WithOpenAIModeration(value *OpenAIConfigApplyConfiguration) *ModerationApplyConfiguration {
	b.OpenAIModeration = value
	return b
}
