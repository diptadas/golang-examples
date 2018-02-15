package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"log"

	"reflect"

	core_v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	util_runtime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

const maxRetries = 5

// Controller object
type Controller struct {
	clientSet    kubernetes.Interface
	queue        workqueue.RateLimitingInterface
	informer     cache.SharedIndexInformer
	eventHandler Handler
}

func Start() {
	kubeClient := GetClientOutOfCluster()

	c := newController(kubeClient, new(Default))
	stopCh := make(chan struct{})
	defer close(stopCh)

	go c.Run(stopCh)

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGTERM)
	signal.Notify(sigterm, syscall.SIGINT)
	<-sigterm

}

func newController(client kubernetes.Interface, eventHandler Handler) *Controller {
	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

	informer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				return client.CoreV1().ConfigMaps(meta_v1.NamespaceAll).List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				return client.CoreV1().ConfigMaps(meta_v1.NamespaceAll).Watch(options)
			},
		},
		&core_v1.ConfigMap{},
		0, // skip re-sync
		cache.Indexers{},
	)

	// Warm up the cache for initial synchronization.
	if list, err := client.CoreV1().ConfigMaps(meta_v1.NamespaceAll).List(meta_v1.ListOptions{}); err == nil {
		for i := range list.Items {
			log.Println("\n**********debug@dipta***********\n", "Previously Added\n", &list.Items[i], "\n==============================")
			informer.GetIndexer().Add(&list.Items[i])
		}
	} else {
		log.Println(err)
	}

	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			if key, err := cache.MetaNamespaceKeyFunc(obj); err == nil {
				log.Println("Queued Add event")
				queue.Add(key)
			} else {
				log.Println(err)
			}
		},
		UpdateFunc: func(old, new interface{}) {
			if oldMap, oldOK := old.(*core_v1.ConfigMap); oldOK {
				if newMap, newOK := new.(*core_v1.ConfigMap); newOK {
					if !reflect.DeepEqual(oldMap.Data, newMap.Data) {
						if key, err := cache.MetaNamespaceKeyFunc(new); err == nil {
							log.Println("Queued Update event", key)
							queue.Add(key)
						} else {
							log.Println(err)
						}
					}
				}
			}
		},
		DeleteFunc: func(obj interface{}) {
			if key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj); err == nil {
				log.Println("Queued Delete event", key)
				queue.Add(key)
			}
		},
	})

	return &Controller{
		clientSet:    client,
		informer:     informer,
		queue:        queue,
		eventHandler: eventHandler,
	}
}

// Run starts the kubewatch controller
func (c *Controller) Run(stopCh <-chan struct{}) {
	defer util_runtime.HandleCrash()
	defer c.queue.ShutDown()

	log.Println("Starting kubewatch controller")

	go c.informer.Run(stopCh)

	if !cache.WaitForCacheSync(stopCh, c.HasSynced) {
		util_runtime.HandleError(fmt.Errorf("Timed out waiting for caches to sync"))
		return
	}

	log.Println("Kubewatch controller synced and ready")

	wait.Until(c.runWorker, time.Second, stopCh)
}

// HasSynced is required for the cache.Controller interface.
func (c *Controller) HasSynced() bool {
	return c.informer.HasSynced()
}

// LastSyncResourceVersion is required for the cache.Controller interface.
func (c *Controller) LastSyncResourceVersion() string {
	return c.informer.LastSyncResourceVersion()
}

func (c *Controller) runWorker() {
	for c.processNextItem() {
		// continue looping
	}
}

func (c *Controller) processNextItem() bool {
	key, quit := c.queue.Get()
	if quit {
		return false
	}
	defer c.queue.Done(key)

	err := c.processItem(key.(string))
	if err == nil {
		c.queue.Forget(key)
	} else if c.queue.NumRequeues(key) < maxRetries {
		log.Printf("Error processing %s (will retry): %v\n", key, err)
		c.queue.AddRateLimited(key)
	} else {
		log.Printf("Error processing %s (giving up): %v\n", key, err)
		c.queue.Forget(key)
		util_runtime.HandleError(err)
	}

	return true
}

func (c *Controller) processItem(key string) error {
	log.Printf("Processing change to Configmap %s\n", key)

	obj, exists, err := c.informer.GetIndexer().GetByKey(key)
	if err != nil {
		return fmt.Errorf("Error fetching object with key %s from store: %v", key, err)
	}

	if !exists {
		c.eventHandler.ObjectDeleted(obj)
		return nil
	}

	c.eventHandler.ObjectCreated(obj)
	return nil
}
