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

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/endpoints/request"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"

	"k8s-learn/api_server/pkg/apis/demo"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Bird
// +k8s:openapi-gen=true
// +resource:path=birds,strategy=BirdStrategy
type Bird struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BirdSpec   `json:"spec,omitempty"`
	Status BirdStatus `json:"status,omitempty"`
}

// BirdSpec defines the desired state of Bird
type BirdSpec struct {
}

// BirdStatus defines the observed state of Bird
type BirdStatus struct {
}

// Validate checks that an instance of Bird is well formed
func (BirdStrategy) Validate(ctx request.Context, obj runtime.Object) field.ErrorList {
	o := obj.(*demo.Bird)
	log.Printf("Validating fields for Bird %s\n", o.Name)
	errors := field.ErrorList{}
	// perform validation here and add to errors using field.Invalid
	return errors
}

// DefaultingFunction sets default Bird field values
func (BirdSchemeFns) DefaultingFunction(o interface{}) {
	obj := o.(*Bird)
	// set default field values here
	log.Printf("Defaulting fields for Bird %s\n", obj.Name)
}
