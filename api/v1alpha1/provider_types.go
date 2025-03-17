package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	gwv1 "sigs.k8s.io/gateway-api/apis/v1"
)

// +kubebuilder:rbac:groups=gateway.kgateway.dev,resources=ExternalProviders,verbs=get;list;watch
// +kubebuilder:rbac:groups=gateway.kgateway.dev,resources=ExternalProviders/status,verbs=get;update;patch

// +kubebuilder:printcolumn:name="Type",type=string,JSONPath=".spec.type",description="Which provider type?"
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=".metadata.creationTimestamp",description="The age of the ExternalProvider."

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:metadata:labels={app=kgateway,app.kubernetes.io/name=kgateway}
// +kubebuilder:resource:categories=kgateway
// +kubebuilder:subresource:status
type ExternalProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ExternalProviderSpec   `json:"spec,omitempty"`
	Status ExternalProviderStatus `json:"status,omitempty"`
}

// ExternalProviderType indicates the type of the ExternalProvider.
type ExternalProviderType string

const (
	// ExternalProviderTypeExtAuth is the type for Extauth providers.
	ExternalProviderTypeExtAuth ExternalProviderType = "ExtAuth"
	// ExternalProviderTypeExtProc is the type for ExtProc providers.
	ExternalProviderTypeExtProc ExternalProviderType = "ExtProc"
)

// FilterStageName represents the name of a filter stage.
// +kubebuilder:validation:Enum=FaultStage;CorsStage;WafStage;AuthNStage;AuthZStage;RateLimitStage;AcceptedStage;OutAuthStage;RouteStage
type FilterStageName string

const (
	FaultStage     FilterStageName = "FaultStage"
	CorsStage      FilterStageName = "CorsStage"
	WafStage       FilterStageName = "WafStage"
	AuthNStage     FilterStageName = "AuthNStage"
	AuthZStage     FilterStageName = "AuthZStage"
	RateLimitStage FilterStageName = "RateLimitStage"
	AcceptedStage  FilterStageName = "AcceptedStage"
	OutAuthStage   FilterStageName = "OutAuthStage"
	RouteStage     FilterStageName = "RouteStage"
)

// StageConfig defines the configuration for where a provider should be placed in the filter chain.
type StageConfig struct {
	// Name of the filter stage where the provider should be placed.
	// +kubebuilder:validation:Required
	Name FilterStageName `json:"name"`

	// Weight determines the relative order of providers within the same stage.
	// Lower weights are processed first.
	// In general ordering within a stage is considered not important.
	// +optional
	// +kubebuilder:validation:Minimum=-20
	// +kubebuilder:validation:Maximum=20
	Weight *int32 `json:"weight,omitempty"`
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

// ExternalProviderSpec defines the desired state of ExternalProvider.
type ExternalProviderSpec struct {
	// Type indicates the type of the ExternalProvider to be used.
	// +unionDiscriminator
	// +kubebuilder:validation:Enum=ExtAuth;ExtProc
	// +kubebuilder:validation:Required
	Type ExternalProviderType `json:"type"`

	// Stage configuration for where this provider should be placed in the filter chain.
	// If not specified, the provider will be placed based on the type of the provider.
	// For example Exauth will be place in the in AuthZStage by default.
	// +optional
	Stage StageConfig `json:"stage"`

	// ExtAuth configuration for ExtAuth provider type.
	// +optional
	ExtAuth *ExtAuthProvider `json:"extAuth,omitempty"`

	// ExtProc configuration for ExtProc provider type.
	// +optional
	ExtProc *ExtProcProvider `json:"extProc,omitempty"`
}

// ExternalProviderStatus defines the observed state of ExternalProvider.
type ExternalProviderStatus struct {
	// Conditions is the list of conditions for the ExternalProvider.
	// +optional
	// +listType=map
	// +listMapKey=type
	// +kubebuilder:validation:MaxItems=8
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
type ExternalProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExternalProvider `json:"items"`
}
