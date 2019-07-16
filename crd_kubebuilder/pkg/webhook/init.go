package webhook

import (
	"os"

	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

func init() {
	AddToManagerFuncs = append(AddToManagerFuncs, addHook)
}

func addHook(mgr manager.Manager) error {
	hookServer, _ := webhook.NewServer("ff", mgr, webhook.ServerOptions{
		Port:                          1887,
		CertDir:                       "",
		Client:                        nil,
		DisableWebhookConfigInstaller: nil,
		BootstrapOptions:              nil,
	})
	if err := mgr.Add(hookServer); err != nil {
		os.Exit(1)
	}

	err := hookServer.Register(&admission.Webhook{
		Path:     "/validate-v1-pod",
		Handlers: []admission.Handler{&podValidator{}}})
	if err != nil {
		panic(err)
	}
	return nil
}
