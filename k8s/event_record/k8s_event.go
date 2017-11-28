package main

import (
	"log"
	"os"

	core "k8s.io/api/core/v1"
	kerr "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/record"
	"time"
)

func main() {
	kubeClient, err := getKubeClient()
	if err != nil {
		log.Fatalf(err.Error())
	}

	resource := &core.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: "example-configmap",
		},
	}

	log.Println("Deleting old resource")
	err = kubeClient.CoreV1().ConfigMaps("default").Delete(resource.Name, &metav1.DeleteOptions{})
	if err != nil && !kerr.IsNotFound(err) {
		log.Fatalf(err.Error())
	}

	log.Println("Creating new resource")
	resource, err = kubeClient.CoreV1().ConfigMaps("default").Create(resource)
	if err != nil {
		log.Fatalf(err.Error())
	}

	log.Println("Recording events")
	_, recorder := NewEventRecorder(kubeClient, "golang-examples")

	go func() {
		for i := 0; i < 5; i++ {
			recorder.Event(resource, core.EventTypeNormal, "event-test", "new event is recorded")
			time.Sleep(time.Second)
		}
	}()

	select {}
}

func NewEventRecorder(client kubernetes.Interface, component string) (watch.Interface, record.EventRecorder) {
	// Event Broadcaster
	broadcaster := record.NewBroadcaster()
	watcher := broadcaster.StartEventWatcher(
		func(event *core.Event) {
			if _, err := client.CoreV1().Events(event.Namespace).Create(event); err != nil {
				log.Println(err)
			} else {
				log.Println("Event recorded:", event.Name)
			}
		},
	)
	// Event Recorder
	return watcher, broadcaster.NewRecorder(scheme.Scheme, core.EventSource{Component: component})
}

func getKubeClient() (kubernetes.Interface, error) {
	kubeConfigPath := os.Getenv("HOME") + "/.kube/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(config)
}
