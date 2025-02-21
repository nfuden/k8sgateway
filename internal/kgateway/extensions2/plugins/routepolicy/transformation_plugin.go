package routepolicy

import (
	"fmt"
	"os"

	"github.com/kgateway-dev/kgateway/v2/api/v1alpha1"
	"github.com/kgateway-dev/kgateway/v2/internal/kgateway/utils"
	transformationpb "github.com/solo-io/envoy-gloo/go/config/filter/http/transformation/v2"
	"google.golang.org/protobuf/types/known/anypb"
)

func toTraditionalTransform(t *v1alpha1.Transform) *transformationpb.Transformation_TransformationTemplate {
	if t == nil {
		return nil
	}
	hasTransform := false
	tt := &transformationpb.Transformation_TransformationTemplate{
		TransformationTemplate: &transformationpb.TransformationTemplate{
			Headers: map[string]*transformationpb.InjaTemplate{},
		},
	}
	for _, h := range t.Set {
		tt.TransformationTemplate.GetHeaders()[string(h.Name)] = &transformationpb.InjaTemplate{
			Text: string(h.Value),
		}
		tt.TransformationTemplate.ParseBodyBehavior = transformationpb.TransformationTemplate_DontParse
		hasTransform = true
	}

	for _, h := range t.Add {
		tt.TransformationTemplate.HeadersToAppend = append(tt.TransformationTemplate.GetHeadersToAppend(), &transformationpb.TransformationTemplate_HeaderToAppend{
			Key: string(h.Name),
			Value: &transformationpb.InjaTemplate{
				Text: string(h.Value),
			},
		})
		tt.TransformationTemplate.ParseBodyBehavior = transformationpb.TransformationTemplate_DontParse
		hasTransform = true
	}

	// tt.TransformationTemplate.HeadersToRemove = make([]string, 0, len(t.Remove))
	// for _, h := range t.Remove {
	// 	tt.TransformationTemplate.HeadersToRemove = append(tt.TransformationTemplate.HeadersToRemove, string(h))
	// 	hasTransform = true
	// }

	//BODY
	// if t.Body == nil {
	// 	tt.TransformationTemplate.BodyTransformation = &transformationpb.TransformationTemplate_Passthrough{
	// 		Passthrough: &transformationpb.Passthrough{},
	// 	}
	// } else {
	// 	if t.Body.ParseAs == v1alpha1.BodyParseBehaviorAsString {
	// 		tt.TransformationTemplate.ParseBodyBehavior = transformationpb.TransformationTemplate_DontParse
	// 	}
	// 	if value := t.Body.Value; value != nil {
	// 		hasTransform = true
	// 		tt.TransformationTemplate.BodyTransformation = &transformationpb.TransformationTemplate_Body{
	// 			Body: &transformationpb.InjaTemplate{
	// 				Text: string(*value),
	// 			},
	// 		}
	// 	}
	// }

	if !hasTransform {
		return nil
	}
	return tt
}

func toTransformFilterConfig(t *v1alpha1.TransformationPolicy) (*anypb.Any, error) {
	if t == nil || *t == (v1alpha1.TransformationPolicy{}) {
		return nil, nil
	}

	toTransform := toTraditionalTransform
	if os.Getenv("USE_RUSTFORMATION") == "true" {
		fmt.Println("lets use rusformation")
	} else {
		fmt.Println("using legacy transformation")
	}
	var reqt *transformationpb.Transformation
	var respt *transformationpb.Transformation

	if rtt := toTransform(t.Request); rtt != nil {
		reqt = &transformationpb.Transformation{
			TransformationType: rtt,
		}
	}
	if rtt := toTransform(t.Response); rtt != nil {
		respt = &transformationpb.Transformation{
			TransformationType: rtt,
		}
	}
	if reqt == nil && respt == nil {
		return nil, nil
	}

	reqm := &transformationpb.RouteTransformations_RouteTransformation_RequestMatch{
		RequestTransformation:  reqt,
		ResponseTransformation: respt,
	}

	envoyT := &transformationpb.RouteTransformations{
		Transformations: []*transformationpb.RouteTransformations_RouteTransformation{
			{

				Match: &transformationpb.RouteTransformations_RouteTransformation_RequestMatch_{
					RequestMatch: reqm,
				},
			},
		},
	}
	return utils.MessageToAny(envoyT)
}

/*
   - name: dynamic_modules/header_mutation
     typed_config:
       # https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/dynamic_modules/v3/dynamic_modules.proto#envoy-v3-api-msg-extensions-dynamic-modules-v3-dynamicmoduleconfig
       "@type": type.googleapis.com/envoy.extensions.filters.http.dynamic_modules.v3.DynamicModuleFilter
       dynamic_module_config:
         name: rust_module
       filter_name: http_simple_mutations
       filter_config: |
         {
           "request_headers_setter": [["X-Envoy-Header", "envoy-header{{substring("ENVOYPROXY", 2, 3)}}"], ["X-Envoy-Header2", "envoy-header2"]],
           "response_headers_setter": [["Foo", "bar"], ["Foo2", "bar2"]]
         }
*/
