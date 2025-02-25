package routepolicy

import (
	"context"
	"time"

	envoyhttp "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"k8s.io/apimachinery/pkg/runtime/schema"

	envoy_config_listener_v3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	envoy_config_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	"istio.io/istio/pkg/kube/krt"

	"github.com/kgateway-dev/kgateway/v2/api/v1alpha1"
	"github.com/kgateway-dev/kgateway/v2/internal/kgateway/extensions2/common"
	extensionplug "github.com/kgateway-dev/kgateway/v2/internal/kgateway/extensions2/plugin"
	extensionsplug "github.com/kgateway-dev/kgateway/v2/internal/kgateway/extensions2/plugin"
	"github.com/kgateway-dev/kgateway/v2/internal/kgateway/ir"
	"github.com/kgateway-dev/kgateway/v2/internal/kgateway/plugins"
	"github.com/kgateway-dev/kgateway/v2/internal/kgateway/utils/krtutil"

	exteniondynamicmodulev3 "github.com/envoyproxy/go-control-plane/envoy/extensions/dynamic_modules/v3"
	dynamicmodulesv3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/dynamic_modules/v3"

	// TODO(nfuden): remove
	transformationpb "github.com/solo-io/envoy-gloo/go/config/filter/http/transformation/v2"
)

const transformationFilterNamePrefix = "transformation"

var (
	pluginStage = plugins.AfterStage(plugins.AuthZStage)
)

type routePolicy struct {
	ct   time.Time
	spec routeSpecIr
}

type routeSpecIr struct {
	timeout   *durationpb.Duration
	transform *anypb.Any
	errors    []error
}

func (d *routePolicy) CreationTime() time.Time {
	return d.ct
}

func (d *routePolicy) Equals(in any) bool {
	d2, ok := in.(*routePolicy)
	if !ok {
		return false
	}

	if !proto.Equal(d.spec.timeout, d2.spec.timeout) {
		return false
	}
	if !proto.Equal(d.spec.transform, d2.spec.transform) {
		return false
	}

	return true
}

type routePolicyPluginGwPass struct {
	setTransformationInChain bool // TODO(nfuden): mae this multi stage
}

func (p *routePolicyPluginGwPass) ApplyHCM(ctx context.Context, pCtx *ir.HcmContext, out *envoyhttp.HttpConnectionManager) error {
	// no op
	return nil
}

func NewPlugin(ctx context.Context, commoncol *common.CommonCollections) extensionplug.Plugin {
	col := krtutil.SetupCollectionDynamic[v1alpha1.RoutePolicy](
		ctx,
		commoncol.Client,
		v1alpha1.SchemeGroupVersion.WithResource("routepolicies"),
		commoncol.KrtOpts.ToOptions("RoutePolicy")...,
	)
	gk := v1alpha1.RoutePolicyGVK.GroupKind()
	policyCol := krt.NewCollection(col, func(krtctx krt.HandlerContext, policyCR *v1alpha1.RoutePolicy) *ir.PolicyWrapper {
		var pol = &ir.PolicyWrapper{
			ObjectSource: ir.ObjectSource{
				Group:     gk.Group,
				Kind:      gk.Kind,
				Namespace: policyCR.Namespace,
				Name:      policyCR.Name,
			},
			Policy:     policyCR,
			PolicyIR:   &routePolicy{ct: policyCR.CreationTimestamp.Time, spec: toSpec(policyCR.Spec)},
			TargetRefs: convert(policyCR.Spec.TargetRef),
		}
		return pol
	})

	return extensionplug.Plugin{
		ContributesPolicies: map[schema.GroupKind]extensionsplug.PolicyPlugin{
			v1alpha1.RoutePolicyGVK.GroupKind(): {
				//AttachmentPoints: []ir.AttachmentPoints{ir.HttpAttachmentPoint},
				NewGatewayTranslationPass: NewGatewayTranslationPass,
				Policies:                  policyCol,
			},
		},
	}
}

func toSpec(spec v1alpha1.RoutePolicySpec) routeSpecIr {
	var ret routeSpecIr

	if spec.Timeout > 0 {
		ret.timeout = durationpb.New(time.Second * time.Duration(spec.Timeout))
	}
	var err error
	ret.transform, err = toTransformFilterConfig(&spec.Transformation)
	if err != nil {
		ret.errors = append(ret.errors, err)
	}

	return ret
}

func convert(targetRef v1alpha1.LocalPolicyTargetReference) []ir.PolicyTargetRef {
	return []ir.PolicyTargetRef{{
		Kind:  string(targetRef.Kind),
		Name:  string(targetRef.Name),
		Group: string(targetRef.Group),
	}}
}

func NewGatewayTranslationPass(ctx context.Context, tctx ir.GwTranslationCtx) ir.ProxyTranslationPass {
	return &routePolicyPluginGwPass{}
}
func (p *routePolicy) Name() string {
	return "routepolicies"
}

// called 1 time for each listener
func (p *routePolicyPluginGwPass) ApplyListenerPlugin(ctx context.Context, pCtx *ir.ListenerContext, out *envoy_config_listener_v3.Listener) {

}

func (p *routePolicyPluginGwPass) ApplyVhostPlugin(ctx context.Context, pCtx *ir.VirtualHostContext, out *envoy_config_route_v3.VirtualHost) {
}

// called 0 or more times
func (p *routePolicyPluginGwPass) ApplyForRoute(ctx context.Context, pCtx *ir.RouteContext, outputRoute *envoy_config_route_v3.Route) error {
	policy, ok := pCtx.Policy.(*routePolicy)
	if !ok {
		return nil
	}
	if policy.spec.timeout != nil && outputRoute.GetRoute() != nil {
		outputRoute.GetRoute().Timeout = policy.spec.timeout
	}

	if policy.spec.transform != nil {
		if outputRoute.GetTypedPerFilterConfig() == nil {
			outputRoute.TypedPerFilterConfig = make(map[string]*anypb.Any)
		}
		outputRoute.GetTypedPerFilterConfig()[transformationFilterNamePrefix] = policy.spec.transform
		p.setTransformationInChain = true
	}

	return nil
}

func (p *routePolicyPluginGwPass) ApplyForRouteBackend(
	ctx context.Context,
	policy ir.PolicyIR,
	pCtx *ir.RouteBackendContext,
) error {
	return nil
}

// called 1 time per listener
// if a plugin emits new filters, they must be with a plugin unique name.
// any filter returned from route config must be disabled, so it doesnt impact other routes.
func (p *routePolicyPluginGwPass) HttpFilters(ctx context.Context, fcc ir.FilterChainCommon) ([]plugins.StagedHttpFilter, error) {
	if p.setTransformationInChain {

		// TODO(nfuden): support stages such as early
		filters := []plugins.StagedHttpFilter{}
		// first register classic
		filters = append(filters, plugins.MustNewStagedFilter(transformationFilterNamePrefix,
			&transformationpb.FilterTransformations{},
			plugins.BeforeStage(plugins.AcceptedStage)))

		// TODO(nfuden/yuvalk): how to do route level correctly
		rustCfg := dynamicmodulesv3.DynamicModuleFilter{
			DynamicModuleConfig: &exteniondynamicmodulev3.DynamicModuleConfig{
				Name: "rust_module",
			},
			FilterName: "http_simple_mutations",
			FilterConfig: `{ "request_headers_setter": [],
			"response_headers_setter": [["x-conditional-donor", "x-donorcontent{{ substring( header(\"x-donor\"), 0, 7)}}"], ["X-emptyish", "ish"]]} `,
		}

		filters = append(filters, plugins.MustNewStagedFilter("dynamic_modules/simple_mutations",
			&rustCfg,
			plugins.BeforeStage(plugins.AcceptedStage)))

		return filters, nil

	}
	return nil, nil
}

func (p *routePolicyPluginGwPass) UpstreamHttpFilters(ctx context.Context) ([]plugins.StagedUpstreamHttpFilter, error) {
	return nil, nil
}

func (p *routePolicyPluginGwPass) NetworkFilters(ctx context.Context) ([]plugins.StagedNetworkFilter, error) {
	return nil, nil
}

// called 1 time (per envoy proxy). replaces GeneratedResources
func (p *routePolicyPluginGwPass) ResourcesToAdd(ctx context.Context) ir.Resources {
	return ir.Resources{}
}
