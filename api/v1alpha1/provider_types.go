package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	gwv1 "sigs.k8s.io/gateway-api/apis/v1"
)

// +kubebuilder:rbac:groups=gateway.kgateway.dev,resources=GatewayPolicies,verbs=get;list;watch
// +kubebuilder:rbac:groups=gateway.kgateway.dev,resources=GatewayPolicies/status,verbs=get;update;patch

// +kubebuilder:printcolumn:name="Type",type=string,JSONPath=".spec.type",description="Which extension type?"
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=".metadata.creationTimestamp",description="The age of the GatewayExtensionPolicy."

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:metadata:labels={app=kgateway,app.kubernetes.io/name=kgateway}
// +kubebuilder:resource:categories=kgateway
// +kubebuilder:subresource:status
type GatewayExtensionPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GatewayExtensionSpec         `json:"spec,omitempty"`
	Status GatewayExtensionPolicyStatus `json:"status,omitempty"`
}

// GatewayExtensionPolicyType indicates the type of the GatewayExtensionPolicy.
type GatewayExtensionPolicyType string

const (
	// GatewayExtensionPolicyTypeExtAuth is the type for Extauth extensions.
	GatewayExtensionPolicyTypeExtAuth GatewayExtensionPolicyType = "ExtAuth"
	// GatewayExtensionPolicyTypeExtProc is the type for ExtProc extensions.
	GatewayExtensionPolicyTypeExtProc GatewayExtensionPolicyType = "ExtProc"

	// GatewayExtensionPolicyTypeExtended is the type for implementations outside of kgateway main.
	GatewayExtensionPolicyTypeExtended GatewayExtensionPolicyType = "Extended"
)

// FilterStageName represents the name of a filter stage.
// +kubebuilder:validation:Enum=FaultStage;CorsStage;WafStage;AuthNStage;AuthZStage;RateLimitStage;AcceptedStage;OutAuthStage;RouteStage
type FilterStageName string

const (
	FaultStage     FilterStageName = "FaultStage"
	AuthNStage     FilterStageName = "AuthNStage"
	AuthZStage     FilterStageName = "AuthZStage"
	RateLimitStage FilterStageName = "RateLimitStage"
	AcceptedStage  FilterStageName = "AcceptedStage"
	OutAuthStage   FilterStageName = "OutAuthStage"
	RouteStage     FilterStageName = "RouteStage"
)

// Placement defines the configuration for where a provider should be placed in the filter chain.
type Placement struct {
	// Name of the filter stage where the provider should be placed.
	// +kubebuilder:validation:Required
	Name FilterStageName `json:"name"`

	// Priority determines the relative order of providers within the same stage.
	// Lower priorities are processed first.
	// In general ordering within a stage is considered not important.
	// +optional
	// +kubebuilder:validation:Minimum=-10
	// +kubebuilder:validation:Maximum=10
	Priority *int32 `json:"priority,omitempty"`
}

// ExtAuthProvider defines the configuration for an ExtAuth provider.
type ExtAuthProvider struct {
	// BackendRef references the backend service that will handle the authentication.
	// +kubebuilder:validation:Required
	BackendRef *gwv1.BackendRef `json:"backendRef"`
}

// ExtProcProvider defines the configuration for an ExtProc provider.
type ExtProcProvider struct {
	// BackendRef references the backend service that will handle the processing.
	// +kubebuilder:validation:Required
	BackendRef *gwv1.BackendRef `json:"backendRef"`
}

// GatewayExtensionSpec defines the desired state of GatewayExtensionPolicy.
// +kubebuilder:validation:XValidation:message="ExtAuth must be set when type is ExtAuth",rule="self.type != 'ExtAuth' || has(self.extAuth)"
// +kubebuilder:validation:XValidation:message="ExtProc must be set when type is ExtProc",rule="self.type != 'ExtProc' || has(self.extProc)"
// +kubebuilder:validation:XValidation:message="ExtAuth must not be set when type is not ExtAuth",rule="self.type == 'ExtAuth' || !has(self.extAuth)"
// +kubebuilder:validation:XValidation:message="ExtProc must not be set when type is not ExtProc",rule="self.type == 'ExtProc' || !has(self.extProc)"
type GatewayExtensionSpec struct {
	// Type indicates the type of the GatewayExtensionPolicy to be used.
	// +unionDiscriminator
	// +kubebuilder:validation:Enum=ExtAuth;ExtProc;Extended
	// +kubebuilder:validation:Required
	Type GatewayExtensionPolicyType `json:"type"`

	// [unimplemented] TODO: add placement support or something else based on:
	//  https://github.com/kgateway-dev/kgateway/blob/main/design/10851.md and its iterations
	// Placement configuration for where this extension should be placed in the filter chain.
	// If not specified, the extension will be placed based on the type of the extension.
	// For example Exauth will be place in the in AuthZStage by default.
	// +optional
	// Placement Placement `json:"placement"`

	// ExtAuth configuration for ExtAuth extension type.
	// +optional
	// +unionMember:type=ExtAuth
	ExtAuth *ExtAuthProvider `json:"extAuth,omitempty"`

	// ExtProc configuration for ExtProc extension type.
	// +optional
	// +unionMember:type=ExtProc
	ExtProc *ExtProcProvider `json:"extProc,omitempty"`
}

// GatewayExtensionPolicyStatus defines the observed state of GatewayExtensionPolicy.
type GatewayExtensionPolicyStatus struct {
	// Conditions is the list of conditions for the GatewayExtensionPolicy.
	// +optional
	// +listType=map
	// +listMapKey=type
	// +kubebuilder:validation:MaxItems=8
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
type GatewayExtensionPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GatewayExtensionPolicy `json:"items"`
}
