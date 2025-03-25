package routepolicy

import (
	envoy_ext_authz_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/ext_authz/v3"

	"github.com/kgateway-dev/kgateway/v2/api/v1alpha1"
)

type extAuthIR struct {
	filter       *envoy_ext_authz_v3.ExtAuthz
	providerName string
	enabled      v1alpha1.ExtAuthEnabled
}

// extAuthForSpec translates the ExtAuthz spec into the Envoy configuration
func extAuthForSpec(routeSpec *v1alpha1.RoutePolicySpec, out *routeSpecIr) {
	if routeSpec == nil || routeSpec.ExtAuth == nil {
		return
	}
	spec := routeSpec.ExtAuth
	// Create the ExtAuthz configuration
	extAuth := &envoy_ext_authz_v3.ExtAuthz{}
	if spec.FailureModeAllow != nil {
		extAuth.FailureModeAllow = *spec.FailureModeAllow
	}
	if spec.ClearRouteCache != nil {
		extAuth.ClearRouteCache = *spec.ClearRouteCache
	}
	if spec.IncludePeerCertificate != nil {
		extAuth.IncludePeerCertificate = *spec.IncludePeerCertificate
	}
	if spec.IncludeTLSSession != nil {
		extAuth.IncludeTlsSession = *spec.IncludeTLSSession
	}

	// Configure metadata context namespaces if specified
	if len(spec.MetadataContextNamespaces) > 0 {
		extAuth.MetadataContextNamespaces = spec.MetadataContextNamespaces
	}

	// Configure request body buffering if specified
	if spec.WithRequestBody != nil {
		extAuth.WithRequestBody = &envoy_ext_authz_v3.BufferSettings{
			MaxRequestBytes: spec.WithRequestBody.MaxRequestBytes,
		}
		if spec.WithRequestBody.AllowPartialMessage != nil {
			extAuth.GetWithRequestBody().AllowPartialMessage = *spec.WithRequestBody.AllowPartialMessage
		}
		if spec.WithRequestBody.PackAsBytes != nil {
			extAuth.GetWithRequestBody().PackAsBytes = *spec.WithRequestBody.PackAsBytes
		}
	}

	if spec.ExtensionRef != nil {

	}

	nameOrPlaceholder := ""
	if spec.ExtensionRef != nil {
		nameOrPlaceholder = string(spec.ExtensionRef.Name)
	}

	out.extAuth = &extAuthIR{
		filter:       extAuth,
		providerName: nameOrPlaceholder,
		enabled:      spec.Enablement,
	}
}
