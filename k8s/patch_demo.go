package main

import (
	"log"
	"os"

	core_util "github.com/appscode/kutil/core/v1"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"fmt"
)

func main() {
	c, err := getKubeClient()
	if err != nil {
		log.Fatal(err)
	}

	meta := metav1.ObjectMeta{
		Name: "my-svc",
		Namespace: core.NamespaceDefault,
	}

	cfg, _, err := core_util.CreateOrPatchService(c, meta, func(obj *core.Service) *core.Service {
		obj.Annotations = map[string]string{
			"id": "one",
		}
		obj.Spec.Ports = []core.ServicePort{
			{
				Port: 8080,
			},
		}
		return obj
	})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(cfg)
	}

	cfg, _, err = core_util.CreateOrPatchService(c, meta, func(obj *core.Service) *core.Service {
		obj.Annotations = map[string]string{
			"id": "two",
		}
		return obj
	})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(cfg)
	}
}

func getKubeClient() (kubernetes.Interface, error) {
	kubeConfigPath := os.Getenv("HOME") + "/.kube/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(config)
}
