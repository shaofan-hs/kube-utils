/**
 * Copyright 2024 KusionStack Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package clusterprovider

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/rest"
)

var _ ClusterManager = &TestClusterManager{}

// TestClusterManager is a test implementation of ClusterProviderInfo
type TestClusterManager struct {
	schema.GroupVersionResource
	ClusterNameToConfig map[string]*rest.Config // Map from cluster name to kubeconfig
}

func (p *TestClusterManager) Init(config *rest.Config) {
	// Do nothing
}

func (p *TestClusterManager) GetClusterMangementGVR() schema.GroupVersionResource {
	return p.GroupVersionResource
}

func (p *TestClusterManager) GetClusterName(obj *unstructured.Unstructured) string {
	if obj == nil {
		return ""
	}
	return obj.GetName() // Use resource name as cluster name
}

func (p *TestClusterManager) GetClusterConfig(obj *unstructured.Unstructured) *rest.Config {
	if obj == nil || p.ClusterNameToConfig == nil {
		return nil
	}
	config, ok := p.ClusterNameToConfig[obj.GetName()]
	if !ok {
		return nil
	}
	return config
}