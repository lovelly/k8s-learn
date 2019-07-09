package main

import (
	"fmt"
	"k8s-learn/crd_controllder/pkg/client/clientset/versioned"
	"k8s-learn/crd_controllder/pkg/client/informers/externalversions"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
	"k8s.io/sample-controller/pkg/signals"
)

func main() {
	stopCh := signals.SetupSignalHandler()
	config, err := buildConfig("", "./config")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	kubeClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		klog.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	crdClientset, err := versioned.NewForConfig(config)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	fmt.Println("11111111")
	informer := externalversions.NewSharedInformerFactory(crdClientset, 0)
	c := NewController(kubeClient, informer.Mycompany().V1().Nginxes())

	if err = c.Run(2, stopCh); err != nil {
		klog.Error(err)
	}
}

func buildConfig(master, kubeconfig string) (*rest.Config, error) {
	if master != "" || kubeconfig != "" {
		return clientcmd.BuildConfigFromFlags(master, kubeconfig)
	}
	return rest.InClusterConfig()
}
