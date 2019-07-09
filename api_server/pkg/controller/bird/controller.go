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

package bird

import (
	"log"

	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"

	"k8s-learn/api_server/pkg/apis/demo/v1alpha1"
	listers "k8s-learn/api_server/pkg/client/listers_generated/demo/v1alpha1"
	"k8s-learn/api_server/pkg/controller/sharedinformers"
)

// +controller:group=demo,version=v1alpha1,kind=Bird,resource=birds
type BirdControllerImpl struct {
	builders.DefaultControllerFns

	// lister indexes properties about Bird
	lister listers.BirdLister
}

// Init initializes the controller and is called by the generated code
// Register watches for additional resource types here.
func (c *BirdControllerImpl) Init(arguments sharedinformers.ControllerInitArguments) {
	// Use the lister for indexing birds labels
	c.lister = arguments.GetSharedInformers().Factory.Demo().V1alpha1().Birds().Lister()
}

// Reconcile handles enqueued messages
func (c *BirdControllerImpl) Reconcile(u *v1alpha1.Bird) error {
	// Implement controller logic here
	log.Printf("Running reconcile Bird for %s\n", u.Name)
	return nil
}

func (c *BirdControllerImpl) Get(namespace, name string) (*v1alpha1.Bird, error) {
	return c.lister.Birds(namespace).Get(name)
}
