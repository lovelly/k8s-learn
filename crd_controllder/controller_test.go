package main

import (
	"encoding/json"
	"fmt"
	"k8s-learn/crd_controllder/pkg/client/clientset/versioned"
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestGet(t *testing.T) {
	config, err := buildConfig("", "./config")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	client, err := versioned.NewForConfig(config)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	nginx, err := client.MycompanyV1().Nginxes("default").Get("my-deployment", v1.GetOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}

	b, _ := json.MarshalIndent(nginx, "	", "	")
	fmt.Println(string(b))
}
