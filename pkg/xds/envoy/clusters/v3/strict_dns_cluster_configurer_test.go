package clusters_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	core_mesh "github.com/kumahq/kuma/pkg/core/resources/apis/mesh"
	"github.com/kumahq/kuma/pkg/core/xds"
	util_proto "github.com/kumahq/kuma/pkg/util/proto"
	"github.com/kumahq/kuma/pkg/xds/envoy"
	"github.com/kumahq/kuma/pkg/xds/envoy/clusters"
)

var _ = Describe("StrictDNSClusterConfigurer", func() {

	It("should generate proper Envoy config", func() {
		// given
		clusterName := "test:cluster"
		address := "google.com"
		port := uint32(80)
		expected := `
        altStatName: test_cluster
        connectTimeout: 5s
        dnsLookupFamily: V4_ONLY
        loadAssignment:
          clusterName: test:cluster
          endpoints:
          - lbEndpoints:
            - endpoint:
                address:
                  socketAddress:
                    address: google.com
                    portValue: 80
              loadBalancingWeight: 100
        name: test:cluster
        type: STRICT_DNS`

		// when
		cluster, err := clusters.NewClusterBuilder(envoy.APIV3).
			Configure(clusters.StrictDNSCluster(clusterName, []xds.Endpoint{
				{
					Target: address,
					Port:   port,
					Tags:   nil,
					Weight: 100,
					ExternalService: &xds.ExternalService{
						TLSEnabled: true,
					},
				},
			}, false)).
			Configure(clusters.Timeout(core_mesh.ProtocolTCP, DefaultTimeout())).
			Build()

		// then
		Expect(err).ToNot(HaveOccurred())

		actual, err := util_proto.ToYAML(cluster)
		Expect(err).ToNot(HaveOccurred())
		Expect(actual).To(MatchYAML(expected))
	})

	It("should generate proper Envoy config with IPv6", func() {
		// given
		clusterName := "test:cluster"
		address := "google.com"
		port := uint32(80)
		expected := `
        altStatName: test_cluster
        connectTimeout: 5s
        loadAssignment:
          clusterName: test:cluster
          endpoints:
          - lbEndpoints:
            - endpoint:
                address:
                  socketAddress:
                    address: google.com
                    portValue: 80
              loadBalancingWeight: 100
        name: test:cluster
        type: STRICT_DNS`

		// when
		cluster, err := clusters.NewClusterBuilder(envoy.APIV3).
			Configure(clusters.StrictDNSCluster(clusterName, []xds.Endpoint{
				{
					Target: address,
					Port:   port,
					Tags:   nil,
					Weight: 100,
					ExternalService: &xds.ExternalService{
						TLSEnabled: true,
					},
				},
			}, true)).
			Configure(clusters.Timeout(core_mesh.ProtocolTCP, DefaultTimeout())).
			Build()

		// then
		Expect(err).ToNot(HaveOccurred())

		actual, err := util_proto.ToYAML(cluster)
		Expect(err).ToNot(HaveOccurred())
		Expect(actual).To(MatchYAML(expected))
	})
})
