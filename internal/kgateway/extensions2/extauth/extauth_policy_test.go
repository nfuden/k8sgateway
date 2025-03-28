package extauth

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/kgateway-dev/kgateway/v2/api/v1alpha1"
	"github.com/kgateway-dev/kgateway/v2/internal/kgateway/ir"
	"github.com/kgateway-dev/kgateway/v2/internal/kgateway/krtcollections"
	"github.com/kgateway-dev/kgateway/v2/internal/kgateway/pluginutils"
)

func TestExtAuthPolicy(t *testing.T) {
	// Create a test ExtAuthProvider
	extAuthProvider := &v1alpha1.ExtAuthProvider{
		Service: &v1alpha1.Service{
			Name:      "extauth",
			Namespace: "default",
			Port:      9000,
		},
		Timeout:          "1s",
		FailureModeAllow: true,
		WithRequestBody: &v1alpha1.BufferSettings{
			MaxRequestBytes: 8192,
		},
		ClearRouteCache:           true,
		MetadataContextNamespaces: []string{"jwt"},
		IncludePeerCertificate:    true,
		IncludeTLSSession:         true,
		EmitFilterStateStats:      true,
	}

	// Create a test GatewayExtension
	extension := &ir.GatewayExtension{
		ObjectSource: ir.ObjectSource{
			Group:     "gateway.kgateway.dev",
			Kind:      "GatewayPolicy",
			Namespace: "default",
			Name:      "test-extauth",
		},
		Type:    v1alpha1.GatewayPolicyTypeExtAuth,
		ExtAuth: extAuthProvider,
	}

	// Create a test ExtAuthPolicy
	policy := &v1alpha1.ExtAuthPolicy{
		ExtensionRef: &corev1.LocalObjectReference{
			Name: "test-extauth",
		},
		FailureModeAllow: true,
		WithRequestBody: &v1alpha1.BufferSettings{
			MaxRequestBytes: 8192,
		},
		ClearRouteCache:           true,
		MetadataContextNamespaces: []string{"jwt"},
		IncludePeerCertificate:    true,
		IncludeTLSSession:         true,
		EmitFilterStateStats:      true,
	}

	// Create a test ExtAuthRoutePolicy
	routePolicy := &v1alpha1.ExtAuthRoutePolicy{
		Enablement: v1alpha1.ExtAuthEnabledLastWins,
	}

	// Create a test ExtAuthPlugin
	plugin := &ExtAuthPlugin{
		extensions: &krtcollections.GatewayExtensionIndex{},
	}

	// Test the policy configuration
	t.Run("policy configuration", func(t *testing.T) {
		// Get the extension
		ext, err := pluginutils.GetGatewayExtension(plugin.extensions, nil, "test-extauth", "default")
		require.NoError(t, err)
		assert.Equal(t, extension, ext)

		// Verify the policy configuration
		assert.Equal(t, policy.FailureModeAllow, ext.ExtAuth.FailureModeAllow)
		assert.Equal(t, policy.WithRequestBody, ext.ExtAuth.WithRequestBody)
		assert.Equal(t, policy.ClearRouteCache, ext.ExtAuth.ClearRouteCache)
		assert.Equal(t, policy.MetadataContextNamespaces, ext.ExtAuth.MetadataContextNamespaces)
		assert.Equal(t, policy.IncludePeerCertificate, ext.ExtAuth.IncludePeerCertificate)
		assert.Equal(t, policy.IncludeTLSSession, ext.ExtAuth.IncludeTLSSession)
		assert.Equal(t, policy.EmitFilterStateStats, ext.ExtAuth.EmitFilterStateStats)
	})

	// Test the route policy configuration
	t.Run("route policy configuration", func(t *testing.T) {
		// Verify the route policy configuration
		assert.Equal(t, routePolicy.Enablement, v1alpha1.ExtAuthEnabledLastWins)
	})
}
