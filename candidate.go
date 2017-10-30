package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	core "k8s.io/api/core/v1"
	kerr "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/leaderelection"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	"k8s.io/client-go/tools/record"
)

func main() {
	flag.Parse()

	leaderElectionLease := 3 * time.Second
	configMapName := "lock-leader"
	namespace := "default"
	podName := "my-pod-" + flag.Arg(0)

	fmt.Println("Leader Election for pod:", podName)

	kubeConfigPath := os.Getenv("HOME") + "/.kube/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		log.Fatalf(err.Error())
	}
	kubeClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf(err.Error())
	}

	configMap := &core.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      configMapName,
			Namespace: namespace,
		},
	}
	if _, err := kubeClient.CoreV1().ConfigMaps(namespace).Create(configMap); err != nil && !kerr.IsAlreadyExists(err) {
		log.Fatal(err)
	}

	resLock := &resourcelock.ConfigMapLock{
		ConfigMapMeta: configMap.ObjectMeta,
		Client: kubeClient.CoreV1(),
		LockConfig: resourcelock.ResourceLockConfig{
			Identity:      podName,
			EventRecorder: &record.FakeRecorder{},
		},
	}

	stopLeading := make(chan struct{})

	go func() {
		leaderelection.RunOrDie(leaderelection.LeaderElectionConfig{
			Lock:          resLock,
			LeaseDuration: leaderElectionLease,
			RenewDeadline: leaderElectionLease * 2 / 3,
			RetryPeriod:   leaderElectionLease / 3,
			Callbacks: leaderelection.LeaderCallbacks{
				OnStartedLeading: func(stop <-chan struct{}) {
					go restOfCode(stopLeading)
				},
				OnStoppedLeading: func() {
					stopLeading <- struct{}{}
				},
			},
		})
	}()

	select {}
}

func restOfCode(stop chan struct{}) {
	fmt.Println("Got leadership, running rest of code...")
	<-stop
	fmt.Println("Lost leadership, exiting rest of code...")
}

/*
go run candidate.go 1 -v=10
go run candidate.go 2 -v=10
kubectl describe configmap lock-leader
*/
