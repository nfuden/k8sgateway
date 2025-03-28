// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfiguration

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	testing "k8s.io/client-go/testing"

	apiv1alpha1 "github.com/kgateway-dev/kgateway/v2/api/applyconfiguration/api/v1alpha1"
	internal "github.com/kgateway-dev/kgateway/v2/api/applyconfiguration/internal"
	v1alpha1 "github.com/kgateway-dev/kgateway/v2/api/v1alpha1"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=gateway.kgateway.dev, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithKind("AccessLog"):
		return &apiv1alpha1.AccessLogApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("AccessLogFilter"):
		return &apiv1alpha1.AccessLogFilterApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("AIBackend"):
		return &apiv1alpha1.AIBackendApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("AiExtension"):
		return &apiv1alpha1.AiExtensionApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("AiExtensionStats"):
		return &apiv1alpha1.AiExtensionStatsApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("AIPromptEnrichment"):
		return &apiv1alpha1.AIPromptEnrichmentApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("AIPromptGuard"):
		return &apiv1alpha1.AIPromptGuardApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("AIRoutePolicy"):
		return &apiv1alpha1.AIRoutePolicyApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("AnthropicConfig"):
		return &apiv1alpha1.AnthropicConfigApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("AwsAuth"):
		return &apiv1alpha1.AwsAuthApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("AwsBackend"):
		return &apiv1alpha1.AwsBackendApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("AwsLambda"):
		return &apiv1alpha1.AwsLambdaApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("AzureOpenAIConfig"):
		return &apiv1alpha1.AzureOpenAIConfigApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Backend"):
		return &apiv1alpha1.BackendApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("BackendSpec"):
		return &apiv1alpha1.BackendSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("BackendStatus"):
		return &apiv1alpha1.BackendStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("BodyTransformation"):
		return &apiv1alpha1.BodyTransformationApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("BufferSettings"):
		return &apiv1alpha1.BufferSettingsApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("CELFilter"):
		return &apiv1alpha1.CELFilterApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("CustomLabel"):
		return &apiv1alpha1.CustomLabelApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("CustomResponse"):
		return &apiv1alpha1.CustomResponseApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("DirectResponse"):
		return &apiv1alpha1.DirectResponseApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("DirectResponseSpec"):
		return &apiv1alpha1.DirectResponseSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("DurationFilter"):
		return &apiv1alpha1.DurationFilterApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("EnvoyBootstrap"):
		return &apiv1alpha1.EnvoyBootstrapApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("EnvoyContainer"):
		return &apiv1alpha1.EnvoyContainerApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ExtAuthProvider"):
		return &apiv1alpha1.ExtAuthProviderApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ExtAuthRoutePolicy"):
		return &apiv1alpha1.ExtAuthRoutePolicyApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ExtProcGrpcService"):
		return &apiv1alpha1.ExtProcGrpcServiceApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ExtProcPolicy"):
		return &apiv1alpha1.ExtProcPolicyApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ExtProcProvider"):
		return &apiv1alpha1.ExtProcProviderApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("FieldDefault"):
		return &apiv1alpha1.FieldDefaultApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("FileSink"):
		return &apiv1alpha1.FileSinkApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("FilterType"):
		return &apiv1alpha1.FilterTypeApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("GatewayExtension"):
		return &apiv1alpha1.GatewayExtensionApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("GatewayExtensionSpec"):
		return &apiv1alpha1.GatewayExtensionSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("GatewayExtensionStatus"):
		return &apiv1alpha1.GatewayExtensionStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("GatewayParameters"):
		return &apiv1alpha1.GatewayParametersApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("GatewayParametersSpec"):
		return &apiv1alpha1.GatewayParametersSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("GeminiConfig"):
		return &apiv1alpha1.GeminiConfigApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("GracefulShutdownSpec"):
		return &apiv1alpha1.GracefulShutdownSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("GrpcService"):
		return &apiv1alpha1.GrpcServiceApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("GrpcStatusFilter"):
		return &apiv1alpha1.GrpcStatusFilterApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("HeaderFilter"):
		return &apiv1alpha1.HeaderFilterApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("HeaderTransformation"):
		return &apiv1alpha1.HeaderTransformationApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Host"):
		return &apiv1alpha1.HostApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("HTTPListenerPolicy"):
		return &apiv1alpha1.HTTPListenerPolicyApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("HTTPListenerPolicySpec"):
		return &apiv1alpha1.HTTPListenerPolicySpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Image"):
		return &apiv1alpha1.ImageApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("IstioContainer"):
		return &apiv1alpha1.IstioContainerApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("IstioIntegration"):
		return &apiv1alpha1.IstioIntegrationApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("KubernetesProxyConfig"):
		return &apiv1alpha1.KubernetesProxyConfigApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("LLMProvider"):
		return &apiv1alpha1.LLMProviderApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("LocalPolicyTargetReference"):
		return &apiv1alpha1.LocalPolicyTargetReferenceApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("LocalRateLimitPolicy"):
		return &apiv1alpha1.LocalRateLimitPolicyApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Message"):
		return &apiv1alpha1.MessageApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Moderation"):
		return &apiv1alpha1.ModerationApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("MultiPoolConfig"):
		return &apiv1alpha1.MultiPoolConfigApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("OpenAIConfig"):
		return &apiv1alpha1.OpenAIConfigApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Pod"):
		return &apiv1alpha1.PodApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Priority"):
		return &apiv1alpha1.PriorityApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ProcessingMode"):
		return &apiv1alpha1.ProcessingModeApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("PromptguardRequest"):
		return &apiv1alpha1.PromptguardRequestApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("PromptguardResponse"):
		return &apiv1alpha1.PromptguardResponseApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ProxyDeployment"):
		return &apiv1alpha1.ProxyDeploymentApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("RateLimit"):
		return &apiv1alpha1.RateLimitApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Regex"):
		return &apiv1alpha1.RegexApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("RegexMatch"):
		return &apiv1alpha1.RegexMatchApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ResponseFlagFilter"):
		return &apiv1alpha1.ResponseFlagFilterApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("RoutePolicy"):
		return &apiv1alpha1.RoutePolicyApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("RoutePolicySpec"):
		return &apiv1alpha1.RoutePolicySpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("SdsBootstrap"):
		return &apiv1alpha1.SdsBootstrapApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("SdsContainer"):
		return &apiv1alpha1.SdsContainerApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Service"):
		return &apiv1alpha1.ServiceApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ServiceAccount"):
		return &apiv1alpha1.ServiceAccountApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("SimpleStatus"):
		return &apiv1alpha1.SimpleStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("SingleAuthToken"):
		return &apiv1alpha1.SingleAuthTokenApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("StaticBackend"):
		return &apiv1alpha1.StaticBackendApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("StatsConfig"):
		return &apiv1alpha1.StatsConfigApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("StatusCodeFilter"):
		return &apiv1alpha1.StatusCodeFilterApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("SupportedLLMProvider"):
		return &apiv1alpha1.SupportedLLMProviderApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("TokenBucket"):
		return &apiv1alpha1.TokenBucketApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Transform"):
		return &apiv1alpha1.TransformApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("TransformationPolicy"):
		return &apiv1alpha1.TransformationPolicyApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("VertexAIConfig"):
		return &apiv1alpha1.VertexAIConfigApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Webhook"):
		return &apiv1alpha1.WebhookApplyConfiguration{}

	}
	return nil
}

func NewTypeConverter(scheme *runtime.Scheme) *testing.TypeConverter {
	return &testing.TypeConverter{Scheme: scheme, TypeResolver: internal.Parser()}
}
