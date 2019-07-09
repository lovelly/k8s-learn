package main

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/klog"
)

func (nginxController *Controller) OnAddPod(obj interface{}) {
	nginxController.handlePodObject(obj)
}

func (nginxController *Controller) OnUpdatePod(oldObj, newObj interface{}) {
	nginxController.handlePodObject(newObj)
}

func (nginxController *Controller) OnDeletePod(obj interface{}) {
	nginxController.handlePodObject(obj)
}

// 消费workqueue
func (nginxController *Controller) runPodWorker() {
	// 什么也不做
}

func (nginxController *Controller) handlePodObject(obj interface{}) (err error) {
	var (
		pod      *v1.Pod
		ok       bool
		nginxKey string
		hasLabel bool
	)

	// 反解出Pod
	if pod, ok = obj.(*v1.Pod); !ok {
		return
	}

	// 确认属于nginx部署的POD
	if nginxKey, hasLabel = pod.Labels["nginxKey"]; !hasLabel {
		return // 不属于nginx部署的POD， 忽略
	}

	klog.Info("add key ", nginxKey, obj)
	// 投递给nginx的workqueue
	nginxController.NginxWorkqueue.AddRateLimited(nginxKey)
	klog.Infoln("[POD - 更新]", nginxKey, pod.Name)
	return
}
