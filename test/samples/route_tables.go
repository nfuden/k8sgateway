package samples

import (
	v1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/core/matchers"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

func RouteTableWithLabelsAndPrefix(name, namespace, prefix string, labels map[string]string) *v1.RouteTable {
	return &v1.RouteTable{
		Metadata: &core.Metadata{
			Name:      name,
			Namespace: namespace,
			Labels:    labels,
		},
		Routes: []*v1.Route{
			{
				Matchers: []*matchers.Matcher{
					{
						PathSpecifier: &matchers.Matcher_Prefix{
							Prefix: prefix,
						},
					},
				},
				Action: &v1.Route_DirectResponseAction{
					DirectResponseAction: &gloov1.DirectResponseAction{Status: 200}},
			},
		},
	}
}
