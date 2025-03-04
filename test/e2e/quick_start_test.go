/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package e2e

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/utils/ptr"
	capi_e2e "sigs.k8s.io/cluster-api/test/e2e"

	"sigs.k8s.io/cluster-api-provider-vsphere/test/e2e/ipam"
)

var _ = Describe("Cluster Creation using Cluster API quick-start test", func() {
	var (
		testSpecificClusterctlConfigPath string
		testSpecificIPAddressClaims      ipam.IPAddressClaims
	)
	BeforeEach(func() {
		testSpecificClusterctlConfigPath, testSpecificIPAddressClaims = ipamHelper.ClaimIPs(ctx, clusterctlConfigPath)
	})
	defer AfterEach(func() {
		Expect(ipamHelper.Cleanup(ctx, testSpecificIPAddressClaims)).To(Succeed())
	})

	capi_e2e.QuickStartSpec(ctx, func() capi_e2e.QuickStartSpecInput {
		return capi_e2e.QuickStartSpecInput{
			E2EConfig:             e2eConfig,
			ClusterctlConfigPath:  testSpecificClusterctlConfigPath,
			BootstrapClusterProxy: bootstrapClusterProxy,
			ArtifactFolder:        artifactFolder,
			SkipCleanup:           skipCleanup,
		}
	})
})

var _ = Describe("ClusterClass Creation using Cluster API quick-start test [PR-Blocking] [ClusterClass]", func() {
	var (
		testSpecificClusterctlConfigPath string
		testSpecificIPAddressClaims      ipam.IPAddressClaims
	)
	BeforeEach(func() {
		testSpecificClusterctlConfigPath, testSpecificIPAddressClaims = ipamHelper.ClaimIPs(ctx, clusterctlConfigPath)
	})
	defer AfterEach(func() {
		Expect(ipamHelper.Cleanup(ctx, testSpecificIPAddressClaims)).To(Succeed())
	})

	capi_e2e.QuickStartSpec(ctx, func() capi_e2e.QuickStartSpecInput {
		return capi_e2e.QuickStartSpecInput{
			E2EConfig:             e2eConfig,
			ClusterctlConfigPath:  testSpecificClusterctlConfigPath,
			BootstrapClusterProxy: bootstrapClusterProxy,
			ArtifactFolder:        artifactFolder,
			SkipCleanup:           skipCleanup,
			Flavor:                ptr.To("topology"),
		}
	})
})

var _ = Describe("Cluster creation with [Ignition] bootstrap [PR-Blocking]", func() {
	var (
		testSpecificClusterctlConfigPath string
		testSpecificIPAddressClaims      ipam.IPAddressClaims
	)
	BeforeEach(func() {
		testSpecificClusterctlConfigPath, testSpecificIPAddressClaims = ipamHelper.ClaimIPs(ctx, clusterctlConfigPath)
	})
	defer AfterEach(func() {
		Expect(ipamHelper.Cleanup(ctx, testSpecificIPAddressClaims)).To(Succeed())
	})

	capi_e2e.QuickStartSpec(ctx, func() capi_e2e.QuickStartSpecInput {
		return capi_e2e.QuickStartSpecInput{
			E2EConfig:             e2eConfig,
			ClusterctlConfigPath:  testSpecificClusterctlConfigPath,
			BootstrapClusterProxy: bootstrapClusterProxy,
			ArtifactFolder:        artifactFolder,
			SkipCleanup:           skipCleanup,
			Flavor:                ptr.To("ignition"),
		}
	})
})
