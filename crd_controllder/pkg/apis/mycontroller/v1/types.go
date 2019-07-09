package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Nginx struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NginxSpec `json:"spec"`
	Status            MyStatus  `json:"status"`
}

type NginxSpec struct {
	DeploymentName string `json:"deploymentName"`
	Replicas       int    `json:"replicas"`
	MyAge          int    `json:"age"`
}

type MyStatus struct {
	AvailableReplicas int32  `json:"availableReplicas"`
	Message           string `json:"message"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type NginxList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Nginx `json:"items"`
}
