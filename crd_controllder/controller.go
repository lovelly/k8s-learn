package main

import (
	"fmt"
	"k8s-learn/crd_controllder/pkg/client/informers/externalversions/mycontroller/v1"
	"time"

	"k8s.io/klog"

	"k8s.io/client-go/informers"
	core_v1 "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

type Controller struct {
	Clientset     *kubernetes.Clientset
	NginxInformer v1.NginxInformer
	PodInformer   core_v1.PodInformer

	NginxWorkqueue workqueue.RateLimitingInterface
	PodWorkqueue   workqueue.RateLimitingInterface
}

func NewController(clientset *kubernetes.Clientset, NginxInformer v1.NginxInformer) *Controller {
	informerFactory := informers.NewSharedInformerFactory(clientset, time.Second*120)

	c := &Controller{
		Clientset:      clientset,
		NginxInformer:  NginxInformer,
		PodInformer:    informerFactory.Core().V1().Pods(),
		PodWorkqueue:   workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "Pod"),
		NginxWorkqueue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "Nginx"),
	}

	c.PodInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			c.OnAddPod(obj)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			c.OnUpdatePod(oldObj, newObj)
		},
		DeleteFunc: func(obj interface{}) {
			c.OnDeletePod(obj)
		},
	})

	c.NginxInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			c.OnAddNginx(obj)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			c.OnUpdateNginx(oldObj, newObj)
		},
		DeleteFunc: func(obj interface{}) {
			c.OnDeleteNginx(obj)
		},
	})

	return c
}

func (c *Controller) Run(threadiness int, stopCh <-chan struct{}) error {
	defer c.NginxWorkqueue.ShutDown()
	defer c.PodWorkqueue.ShutDown()

	// nginx informer开始拉取事件，存到local cache，并回调event handler
	go c.NginxInformer.Informer().Run(stopCh)
	// pod informer开始拉取事件，存到local cache，并回调event handler
	go c.PodInformer.Informer().Run(stopCh)

	klog.Info("wait object ")
	fmt.Println("22222222")
	// 等待etcd已有数据都下载回来, 再启动事件处理线程, 这样local cache可以反馈出贴近准实时的etcd数据，供逻辑决策准确
	if syncOk := cache.WaitForCacheSync(stopCh, c.NginxInformer.Informer().HasSynced, c.PodInformer.Informer().HasSynced); !syncOk {
		err := fmt.Errorf("sync失败")
		return err
	}

	// 启动nginx event processor
	for i := 0; i < threadiness; i++ {
		go c.runNginxWorker()
		go c.runPodWorker()
	}

	klog.Info("wait close ... ")
	<-stopCh
	return nil
}
