package main

import (
	"log"
)

// Handler is implemented by any handler.
// The Handle method is used to process event
type Handler interface {
	Init(config interface{}) error
	ObjectCreated(obj interface{})
	ObjectDeleted(obj interface{})
	ObjectUpdated(oldObj, newObj interface{})
}

// Default handler implements Handler interface,
// print each event with JSON format
type Default struct {
}

// Init initializes handler configuration
// Do nothing for default handler
func (d *Default) Init(config interface{}) error {
	return nil
}

func (d *Default) ObjectCreated(obj interface{}) {
	log.Println("\n**********debug@dipta***********\n", "Created\n", obj, "\n==============================")
}

func (d *Default) ObjectDeleted(obj interface{}) {
	log.Println("\n**********debug@dipta***********\n", "Deleted\n", obj, "\n==============================")
}

func (d *Default) ObjectUpdated(oldObj, newObj interface{}) {
	log.Println("\n**********debug@dipta***********\n", "Updated\n", oldObj, newObj, "\n==============================")
}
