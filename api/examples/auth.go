package examples

import (
	"fmt"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/kubernetes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"time"
)


func AuthIn(){
	for  {
		time.Sleep(10 * time.Second)
		// creates the in-cluster config
		config, err := rest.InClusterConfig()
		if err != nil {
			log.Println(err)
			continue
		}
		// creates the clientset
		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			log.Println(err)
			continue
		}
		for {
			time.Sleep(10 * time.Second)
			pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
			if err != nil {
				log.Println(err)
				continue
			}
			fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

			// Examples for error handling:
			// - Use helper functions like e.g. errors.IsNotFound()
			// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
			_, err = clientset.CoreV1().Pods("default").Get("example-xxxxx", metav1.GetOptions{})
			if errors.IsNotFound(err) {
				fmt.Printf("Pod not found\n")
			} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
				fmt.Printf("Error getting pod %v\n", statusError.ErrStatus.Message)
			} else if err != nil {
				log.Println(err.Error())
			} else {
				fmt.Printf("Found pod\n")
			}
		}
	}
}

func AuthOut(){
	config, err := clientcmd.BuildConfigFromFlags("", "./config")
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	for {
		pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(pods.Items)
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

		// Examples for error handling:
		// - Use helper functions like e.g. errors.IsNotFound()
		// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
		namespace := "default"
		pod := "task-pv-pod"
		_, err = clientset.CoreV1().Pods(namespace).Get(pod, metav1.GetOptions{})
		if errors.IsNotFound(err) {
			fmt.Printf("Pod %s in namespace %s not found\n", pod, namespace)
		} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
			fmt.Printf("Error getting pod %s in namespace %s: %v\n",
				pod, namespace, statusError.ErrStatus.Message)
		} else if err != nil {
			panic(err.Error())
		} else {
			fmt.Printf("Found pod %s in namespace %s\n", pod, namespace)
		}

		time.Sleep(10 * time.Second)
	}
}
