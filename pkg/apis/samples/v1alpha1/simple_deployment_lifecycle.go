/*
Copyright 2020 The Knative Authors

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

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"knative.dev/pkg/apis"
)

var simpleDeploymentCondSet = apis.NewLivingConditionSet()

// GetGroupVersionKind implements kmeta.OwnerRefable
func (*SimpleDeployment) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("AddressableService")
}

// GetConditionSet retrieves the condition set for this resource. Implements the KRShaped interface.
func (d *SimpleDeployment) GetConditionSet() apis.ConditionSet {
	return simpleDeploymentCondSet
}

// InitializeConditions sets the initial values to the conditions.
func (ds *SimpleDeploymentStatus) InitializeConditions() {
	simpleDeploymentCondSet.Manage(ds).InitializeConditions()
}

// MarkPodsNotReady makes the SimpleDeployment be not ready.
func (ds *SimpleDeploymentStatus) MarkPodsNotReady(n int32) {
	simpleDeploymentCondSet.Manage(ds).MarkFalse(
		SimpleDeploymentConditionReady,
		"PodsNotReady",
		"%d pods are not ready yet", n)
}

// MarkPodsReady makes the SimpleDeployment be ready.
func (ds *SimpleDeploymentStatus) MarkPodsReady() {
	condSet.Manage(ds).MarkTrue(SimpleDeploymentConditionReady)
}
