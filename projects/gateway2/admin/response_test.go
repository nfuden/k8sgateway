package admin_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rotisserie/eris"
	"github.com/solo-io/gloo/projects/gateway2/admin"
	crdv1 "github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd/solo.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("SnapshotResponseData", func() {

	DescribeTable("MarshalJSONString",
		func(response admin.SnapshotResponseData, expectedString string) {
			responseStr := response.MarshalJSONString()
			Expect(responseStr).To(Equal(expectedString))
		},
		Entry("successful response can be formatted as json",
			admin.SnapshotResponseData{
				Data:  "my data",
				Error: nil,
			},
			"{\"data\":\"my data\",\"error\":\"\"}"),
		Entry("errored response can be formatted as json",
			admin.SnapshotResponseData{
				Data:  "",
				Error: eris.New("one error"),
			},
			"{\"data\":\"\",\"error\":\"one error\"}"),
		Entry("CR list can be formatted as json",
			admin.SnapshotResponseData{
				Data: []crdv1.Resource{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "name",
							Namespace: "namespace",
							ManagedFields: []metav1.ManagedFieldsEntry{{
								Manager: "manager",
							}},
						},
						TypeMeta: metav1.TypeMeta{
							Kind:       "kind",
							APIVersion: "version",
						},
					},
				},
				Error: nil,
			},
			"{\"data\":[{\"kind\":\"kind\",\"apiVersion\":\"version\",\"metadata\":{\"name\":\"name\",\"namespace\":\"namespace\",\"creationTimestamp\":null,\"managedFields\":[{\"manager\":\"manager\"}]},\"status\":null,\"spec\":null}],\"error\":\"\"}"),
	)
})