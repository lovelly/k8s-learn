package webhook

import (
	"context"
	"fmt"
	"net/http"

	"sigs.k8s.io/controller-runtime/pkg/webhook/admission/types"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:webhook:path=/validate-v1-pod,mutating=false,failurePolicy=fail,groups="",resources=pods,verbs=create;update,versions=v1,name=vpod.kb.io

// podValidator validates Pods
type podValidator struct {
	client  client.Client
	decoder *admission.DecodeFunc
}

// podValidator admits a pod iff a specific annotation exists.
func (v *podValidator) Handle(ctx context.Context, req types.Request) types.Response {
	pod := &corev1.Pod{}

	err := v.decoder.Decode(req, pod)
	if err != nil {
		return admission.ErrorResponse(http.StatusBadRequest, err)
	}

	key := "example-mutating-admission-webhook"
	anno, found := pod.Annotations[key]
	if !found {
		return admission.ValidationResponse(false, fmt.Sprintf("missing annotation %s", key))
	}
	if anno != "foo" {
		return admission.ValidationResponse(false, fmt.Sprintf("annotation %s did not have value %q", key, "foo"))
	}

	return admission.ValidationResponse(true, "")
}

// podValidator implements inject.Client.
// A client will be automatically injected.

// InjectClient injects the client.
func (v *podValidator) InjectClient(c client.Client) error {
	v.client = c
	return nil
}

// podValidator implements admission.DecoderInjector.
// A decoder will be automatically injected.

// InjectDecoder injects the decoder.
func (v *podValidator) InjectDecoder(d *admission.DecodeFunc) error {
	v.decoder = d
	return nil
}
