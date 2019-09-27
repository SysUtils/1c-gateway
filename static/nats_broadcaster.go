package odata

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"log"
	"reflect"
	"strings"
)

type NatsBroadcaster struct {
	watcher  *Watcher
	natsConn *nats.Conn
}

func NewNatsBroadcaster(watcher *Watcher, natsConn *nats.Conn) *NatsBroadcaster {
	return &NatsBroadcaster{watcher: watcher, natsConn: natsConn}
}

func (n *NatsBroadcaster) Start() {
	go n.loop()
}

func (n *NatsBroadcaster) broadcast(postFix string, e interface{}) {
	name := reflect.TypeOf(e)
	data, _ := json.Marshal(e)
	t := strings.TrimPrefix(name.String(), "*odata.")
	err := n.natsConn.Publish(t+"."+postFix, data)
	if err != nil {
		log.Printf("NATS broadcast error: %s", err)
	}
}

func (n *NatsBroadcaster) loop() {
	for {
		select {
		case e := <-n.watcher.Created:
			n.broadcast("Create", e)
		case e := <-n.watcher.Updated:
			n.broadcast("Update", e)
		case e := <-n.watcher.Deleted:
			n.broadcast("Delete", e)
		}
	}
}
