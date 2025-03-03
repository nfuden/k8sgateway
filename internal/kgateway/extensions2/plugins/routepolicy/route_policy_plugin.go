package routepolicy

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	exteniondynamicmodulev3 "github.com/envoyproxy/go-control-plane/envoy/extensions/dynamic_modules/v3"
	dynamicmodulesv3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/dynamic_modules/v3"
	envoyhttp "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
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
	"github.com/kgateway-dev/kgateway/v2/internal/kgateway/utils"
	"github.com/kgateway-dev/kgateway/v2/internal/kgateway/utils/krtutil"

	// TODO(nfuden): remove

	transformationpb "github.com/solo-io/envoy-gloo/go/config/filter/http/transformation/v2"
)

const transformationFilterNamePrefix = "transformation"
const setFilterStateFilterName = "setfilterstate"
const setMetadataFilterName = "setmetadata"
const rustformationFilterNamePrefix = "composite/rustformation"

const hackKey = "kgateway.route.hack"

// const hackKey = "io.solo.transformation"

var (
	pluginStage = plugins.AfterStage(plugins.AuthZStage)
)

type routePolicy struct {
	ct   time.Time
	spec routeSpecIr
}

type routeSpecIr struct {
	timeout            *durationpb.Duration
	transform          *anypb.Any
	rustformation      *anypb.Any
	rustformationStash map[string]string
	errors             []error
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
	// TODO(nfuden): dont abuse httplevel filter in favor of route level
	rustformationStash map[string]string
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

	rustformation, toStash, err := torustformFilterConfig(&spec.Transformation)
	if err != nil {
		ret.errors = append(ret.errors, err)
	}
	ret.rustformation = rustformation
	if ret.rustformationStash == nil {
		ret.rustformationStash = make(map[string]string)
	}
	ret.rustformationStash[toStash] = string(toStash)

	if ret.errors != nil {
		panic(fmt.Sprintf("failed to create d filter config %v", ret.errors))
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
		if policy.spec.transform != nil {
			fmt.Println("setting transformation in route")
			outputRoute.GetTypedPerFilterConfig()[transformationFilterNamePrefix] = policy.spec.transform
		}
		if policy.spec.rustformation != nil {

			// we cant do this in the route so have extra shared state
			fmt.Println("rusting =")

			// outputRoute.GetTypedPerFilterConfig()["composite"] = policy.spec.rustformation

			// TODO(nfuden): get back to this path once we have valid perroute
			// outputRoute.GetTypedPerFilterConfig()["dynamic_modules/simple_mutations"] = policy.spec.rustformation

			// Hack around not having route level.
			// Note this is really really bad and rather fragile
			routeHash := strconv.Itoa(int(utils.HashProto(outputRoute)))
			if p.rustformationStash == nil {
				p.rustformationStash = make(map[string]string)
			}

			p.rustformationStash[routeHash] = string(policy.spec.rustformation.GetValue())

			// setMetaCfgRaw := set_metadatav3.Config{
			// 	MetadataNamespace: "kgateway",
			// 	Metadata: []*set_metadatav3.Metadata{
			// 		{
			// 			MetadataNamespace: "kgateway",
			// 			AllowOverwrite:    true,

			// 			Value: &structpb.Struct{Fields: map[string]*structpb.Value{
			// 				hackKey: structpb.NewStringValue(routeHash),
			// 			}},
			// 		},
			// 	},
			// }

			// setCfgRaw := set_filter_statev3.Config{OnRequestHeaders: []*common_set_filter_statev3.FilterStateValue{
			// 	&common_set_filter_statev3.FilterStateValue{
			// 		Key:        &common_set_filter_statev3.FilterStateValue_ObjectKey{ObjectKey: hackKey},
			// 		FactoryKey: "envoy.string",
			// 		Value: &common_set_filter_statev3.FilterStateValue_FormatString{
			// 			FormatString: &corev3.SubstitutionFormatString{
			// 				Format: &corev3.SubstitutionFormatString_TextFormat{
			// 					TextFormat: (routeHash),
			// 				},
			// 			},
			// 		},
			// 	},
			// }}
			// setCfg, _ := utils.MessageToAny(&setCfgRaw)

			// outputRoute.GetTypedPerFilterConfig()[setFilterStateFilterName] = setCfg

			meta := &transformationpb.Transformation{
				TransformationType: &transformationpb.Transformation_TransformationTemplate{
					TransformationTemplate: &transformationpb.TransformationTemplate{
						ParseBodyBehavior: transformationpb.TransformationTemplate_DontParse, // Default is to try for JSON... Its kinda nice but failure is bad...
						DynamicMetadataValues: []*transformationpb.TransformationTemplate_DynamicMetadataValue{
							&transformationpb.TransformationTemplate_DynamicMetadataValue{
								MetadataNamespace: "kgateway",
								Key:               "route",
								Value: &transformationpb.InjaTemplate{
									Text: routeHash,
								},
							},
						},
					},
				},
			}

			reqm := &transformationpb.RouteTransformations_RouteTransformation_RequestMatch{
				RequestTransformation:  meta,
				ResponseTransformation: meta,
			}

			setmetaTransform := &transformationpb.RouteTransformations{
				Transformations: []*transformationpb.RouteTransformations_RouteTransformation{
					{

						Match: &transformationpb.RouteTransformations_RouteTransformation_RequestMatch_{
							RequestMatch: reqm,
						},
					},
				},
			}
			outputRoute.GetTypedPerFilterConfig()["helper/perroute/transform"], _ = utils.MessageToAny(setmetaTransform)

		}
		fmt.Println("setting transformation in chain")
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

func mustMessageToAny(msgIn protoreflect.ProtoMessage) *anypb.Any {
	anyOut, _ := utils.MessageToAny(msgIn)
	return anyOut
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

		// ---------------
		// | END CLASSIC |
		// ---------------

		// TODO(nfuden/yuvalk): how to do route level correctly probably contribute to dynamic module upstream
		// smash together configuration
		filterRouteHashConfig := map[string]string{}

		for k, v := range p.rustformationStash {
			fmt.Println("k", k, "v", v)
			filterRouteHashConfig[k] = v
		}

		filterConfig, _ := json.Marshal(filterRouteHashConfig)

		rustCfg := dynamicmodulesv3.DynamicModuleFilter{
			DynamicModuleConfig: &exteniondynamicmodulev3.DynamicModuleConfig{
				Name: "rust_module",
			},
			FilterName:   "http_simple_mutations",
			FilterConfig: fmt.Sprintf(`{ "request_headers_setter": [],"response_headers_setter": [], "route_specific": %s}`, filterConfig),
		}
		// rustCfgAny := mustMessageToAny(&rustCfg)

		// emptyComposite := mustMessageToAny(&compositev3.Composite{})

		// actionCfg := mustMessageToAny(&compositev3.ExecuteFilterAction{
		// 	TypedConfig: &corev3.TypedExtensionConfig{
		// 		Name:        rustformationFilterNamePrefix,
		// 		TypedConfig: rustCfgAny,
		// 	}})

		// matcherCfg := &cncftypev3.Matcher{
		// 	OnNoMatch: &cncftypev3.Matcher_OnMatch{
		// 		OnMatch: &cncftypev3.Matcher_OnMatch_Action{
		// 			Action: &cncfcorev3.TypedExtensionConfig{
		// 				Name:        "composite-action",
		// 				TypedConfig: actionCfg,
		// 			},
		// 		},
		// 	},
		// }

		// populate the OnMatch with a tree
		// matcherCfg.MatcherType = &cncfmatcherv3.Matcher_MatcherTree_{
		// 	MatcherTree: &cncfmatcherv3.Matcher_MatcherTree{
		// 		// Input: ,
		// 		TreeType: &cncfmatcherv3.Matcher_MatcherTree_ExactMatchMap{{}
		// 	},

		// }

		// compositeCfg := extensionmatcherv3.ExtensionWithMatcher{
		// 	ExtensionConfig: &corev3.TypedExtensionConfig{
		// 		Name:        "composite",
		// 		TypedConfig: emptyComposite,
		// 	},
		// 	XdsMatcher: matcherCfg,
		// }

		filters = append(filters, plugins.MustNewStagedFilter("dynamic_modules/simple_mutations",
			&rustCfg,
			plugins.BeforeStage(plugins.AcceptedStage)))

		// filters = append(filters, plugins.MustNewStagedFilter(setFilterStateFilterName,
		// 	&set_filter_statev3.Config{}, plugins.AfterStage(plugins.FaultStage)))
		filters = append(filters, plugins.MustNewStagedFilter("helper/perroute/transform",
			&transformationpb.FilterTransformations{},
			plugins.AfterStage(plugins.FaultStage)))

		// setMetaCfgRaw := set_metadatav3.Config{
		// 	MetadataNamespace: "kgateway",
		// }
		// filters = append(filters, plugins.MustNewStagedFilter(setMetadataFilterName,
		// 	&setMetaCfgRaw, plugins.AfterStage(plugins.FaultStage)),
		// )

		// TODO(nfuden): find a cleaner way to do route level configuration
		// filters = append(filters, plugins.MustNewStagedFilter("composite",
		// 	&compositeCfg,
		// 	plugins.BeforeStage(plugins.AcceptedStage)))
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
