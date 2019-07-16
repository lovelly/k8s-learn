/*
Copyright 2017 The Kubernetes Authors.

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
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MyController
// +k8s:openapi-gen=true
// +resource:path=mycontrollers,strategy=MyControllerStrategy
type MyController struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MyControllerSpec   `json:"spec,omitempty"`
	Status MyControllerStatus `json:"status,omitempty"`
}

// MyControllerSpec defines the desired state of MyController
type MyControllerSpec struct {
}

// MyControllerStatus defines the observed state of MyController
type MyControllerStatus struct {
}

// DefaultingFunction sets default MyController field values
func (MyControllerSchemeFns) DefaultingFunction(o interface{}) {
	obj := o.(*MyController)
	// set default field values here
	log.Printf("Defaulting fields for MyController %s\n", obj.Name)
}
