package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"k8s.io/api/admissionregistration/v1alpha1"
	apps "k8s.io/api/apps/v1beta1"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ktypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/strategicpatch"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	kubeClient kubernetes.Interface
)

func main() {
	flag.Parse()

	kubeConfigPath := os.Getenv("HOME") + "/.kube/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		log.Fatalf(err.Error())
	}
	kubeClient, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf(err.Error())
	}

	go RunController()

	initConfig := InitializerForResources([]string{"deployments"})
	fmt.Println("Creating InitializerConfigurations:", initConfig.Name)
	_, err = kubeClient.AdmissionregistrationV1alpha1().InitializerConfigurations().Create(&initConfig)
	if err != nil {
		log.Fatalf(err.Error())
	}

	time.Sleep(3 * time.Second)

	deploy := Deployment()
	fmt.Println("Creating Deployment:", deploy.Name)
	_, err = kubeClient.AppsV1beta1().Deployments(deploy.Namespace).Create(&deploy)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println("Getting deploy:", deploy.Name)
	obj, err := kubeClient.AppsV1beta1().Deployments(deploy.Namespace).Get(deploy.Name, metav1.GetOptions{})
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println("Status:", obj.Status.String())
	select {}
}

func RunController() {
	lw := &cache.ListWatch{
		ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
			options.IncludeUninitialized = true
			return kubeClient.AppsV1beta1().Deployments(core.NamespaceAll).List(options)
		},
		WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
			options.IncludeUninitialized = true
			return kubeClient.AppsV1beta1().Deployments(core.NamespaceAll).Watch(options)
		},
	}

	resyncPeriod := 30 * time.Second

	_, controller := cache.NewInformer(lw, &apps.Deployment{}, resyncPeriod,
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				deploy := obj.(*apps.Deployment)
				fmt.Println("Added deployment:", deploy.Name)
				if deploy.GetInitializers() != nil && len(deploy.GetInitializers().Pending) > 0 {
					InitializeDeploy(deploy)
				}
			},
		},
	)

	fmt.Println("Running controller")
	go controller.Run(make(chan struct{}))
}

func InitializeDeploy(deploy *apps.Deployment) {
	fmt.Println("Patching deployment:", deploy.Name)
	obj, err := PatchDeployment(kubeClient, deploy, func(deploy *apps.Deployment) *apps.Deployment {
		pending := deploy.Initializers.Pending
		fmt.Println("Removing pending initializer:", pending[0])
		if len(pending) == 1 {
			deploy.ObjectMeta.Initializers = nil
		} else {
			deploy.ObjectMeta.Initializers.Pending = append(pending[:0], pending[1:]...)
		}
		return deploy
	})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Patched Initializer:", obj.Initializers)
}

func InitializerForResources(resources []string) v1alpha1.InitializerConfiguration {
	return v1alpha1.InitializerConfiguration{
		ObjectMeta: metav1.ObjectMeta{
			Name: "my-initializer-config",
			Labels: map[string]string{
				"app": "my-app",
			},
		},
		Initializers: []v1alpha1.Initializer{
			{
				Name: "com.my.initializer",
				Rules: []v1alpha1.Rule{
					{
						APIGroups:   []string{"*"},
						APIVersions: []string{"*"},
						Resources:   resources,
					},
				},
			},
		},
	}
}

func PodTemplate() core.PodTemplateSpec {
	return core.PodTemplateSpec{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app": "my-app",
			},
		},
		Spec: core.PodSpec{
			Containers: []core.Container{
				{
					Name:            "busybox",
					Image:           "busybox",
					ImagePullPolicy: core.PullIfNotPresent,
					Command: []string{
						"sleep",
						"3600",
					},
				},
			},
		},
	}
}

func Deployment() apps.Deployment {
	return apps.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-deploy",
			Namespace: "default",
		},
		Spec: apps.DeploymentSpec{
			Replicas: func() *int32 { i := int32(1); return &i }(),
			Template: PodTemplate(),
		},
	}
}

func PatchDeployment(c kubernetes.Interface, cur *apps.Deployment, transform func(*apps.Deployment) *apps.Deployment) (*apps.Deployment, error) {
	curJson, err := json.Marshal(cur)
	if err != nil {
		return nil, err
	}

	modJson, err := json.Marshal(transform(cur))
	if err != nil {
		return nil, err
	}

	patch, err := strategicpatch.CreateTwoWayMergePatch(curJson, modJson, apps.Deployment{})
	if err != nil {
		return nil, err
	}
	if len(patch) == 0 || string(patch) == "{}" {
		return cur, nil
	}
	return c.AppsV1beta1().Deployments(cur.Namespace).Patch(cur.Name, ktypes.StrategicMergePatchType, patch)
}

/*
minikube start --extra-config=apiserver.Admission.PluginNames="Initializers,NamespaceLifecycle,LimitRanger,ServiceAccount,DefaultStorageClass,ResourceQuota"

kubectl delete initializerconfigurations/my-initializer-config deploy/my-deploy

goimports -w k8s_initializer.go; go run k8s_initializer.go
*/
