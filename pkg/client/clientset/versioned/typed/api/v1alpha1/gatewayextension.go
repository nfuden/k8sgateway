// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	context "context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"

	applyconfigurationapiv1alpha1 "github.com/kgateway-dev/kgateway/v2/api/applyconfiguration/api/v1alpha1"
	apiv1alpha1 "github.com/kgateway-dev/kgateway/v2/api/v1alpha1"
	scheme "github.com/kgateway-dev/kgateway/v2/pkg/client/clientset/versioned/scheme"
)

// GatewayExtensionsGetter has a method to return a GatewayExtensionInterface.
// A group's client should implement this interface.
type GatewayExtensionsGetter interface {
	GatewayExtensions(namespace string) GatewayExtensionInterface
}

// GatewayExtensionInterface has methods to work with GatewayExtension resources.
type GatewayExtensionInterface interface {
	Create(ctx context.Context, gatewayExtension *apiv1alpha1.GatewayExtension, opts v1.CreateOptions) (*apiv1alpha1.GatewayExtension, error)
	Update(ctx context.Context, gatewayExtension *apiv1alpha1.GatewayExtension, opts v1.UpdateOptions) (*apiv1alpha1.GatewayExtension, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, gatewayExtension *apiv1alpha1.GatewayExtension, opts v1.UpdateOptions) (*apiv1alpha1.GatewayExtension, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*apiv1alpha1.GatewayExtension, error)
	List(ctx context.Context, opts v1.ListOptions) (*apiv1alpha1.GatewayExtensionList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *apiv1alpha1.GatewayExtension, err error)
	Apply(ctx context.Context, gatewayExtension *applyconfigurationapiv1alpha1.GatewayExtensionApplyConfiguration, opts v1.ApplyOptions) (result *apiv1alpha1.GatewayExtension, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, gatewayExtension *applyconfigurationapiv1alpha1.GatewayExtensionApplyConfiguration, opts v1.ApplyOptions) (result *apiv1alpha1.GatewayExtension, err error)
	GatewayExtensionExpansion
}

// gatewayExtensions implements GatewayExtensionInterface
type gatewayExtensions struct {
	*gentype.ClientWithListAndApply[*apiv1alpha1.GatewayExtension, *apiv1alpha1.GatewayExtensionList, *applyconfigurationapiv1alpha1.GatewayExtensionApplyConfiguration]
}

// newGatewayExtensions returns a GatewayExtensions
func newGatewayExtensions(c *GatewayV1alpha1Client, namespace string) *gatewayExtensions {
	return &gatewayExtensions{
		gentype.NewClientWithListAndApply[*apiv1alpha1.GatewayExtension, *apiv1alpha1.GatewayExtensionList, *applyconfigurationapiv1alpha1.GatewayExtensionApplyConfiguration](
			"gatewayextensions",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *apiv1alpha1.GatewayExtension { return &apiv1alpha1.GatewayExtension{} },
			func() *apiv1alpha1.GatewayExtensionList { return &apiv1alpha1.GatewayExtensionList{} },
		),
	}
}
