package subscriptions

import (
	"fmt"
	"github.com/SysUtils/1c-gateway/shared"
	"log"
	"os"
)

type Generator struct {
	TypeMap map[string]string
	NameMap map[string]string
	schema  shared.Schema
}

func NewGenerator(schema shared.Schema) *Generator {
	return &Generator{schema: schema, TypeMap: make(map[string]string), NameMap: make(map[string]string)}
}

func (g *Generator) Start() {
	shared.ExtractAsset("static/splitter.go", "odata/splitter.go")
	shared.ExtractAsset("static/nats_broadcaster.go", "odata/nats_broadcaster.go")

	create, update := filterEntities(g.schema.Entities)
	data := fmt.Sprintf(
		`package odata

import (
	"sync"
	"time"
	"log"
)

type EventType int

const (
	None EventType = iota
%s
)

type Watcher struct {
	listenersCounts map[EventType]int
	listenersCountsMutex *sync.RWMutex
	Client *Client
	Deleted chan interface{}
	Updated chan interface{}
	Created chan interface{}
}

func NewWatcher(client *Client) *Watcher {
	r := &Watcher{
		Client:          client,
		listenersCounts: make(map[EventType]int, 0),
		listenersCountsMutex: &sync.RWMutex{},
		Deleted:         make(chan interface{}),
		Updated:         make(chan interface{}),
		Created:         make(chan interface{}),
	}
	return r
}

func (r *Watcher) InternalCopy() *Watcher {
	return &Watcher {
		listenersCounts: r.listenersCounts,
		listenersCountsMutex: r. listenersCountsMutex,
		Deleted:         make(chan interface{}),
		Updated:         make(chan interface{}),
		Created:         make(chan interface{}),
	}
}

func (r *Watcher) AddEventListener(t EventType, delta int) {
	r.listenersCountsMutex.Lock()
	defer r.listenersCountsMutex.Unlock()
	r.listenersCounts[t] += delta
}

func (r *Watcher) Start(offset int64) {
	go r.watch(time.Duration(offset))
}

%s`, g.genTypeConsts(create, update), g.genWatchers(create, update))
	g.writeGofile("Watcher.go", data)
}

func (g *Generator) writeGofile(filename, data string) {
	f, err := os.Create("odata/" + filename)
	if err != nil {
		log.Panic(err)
	}
	_, err = f.WriteString(data)
	if err != nil {
		log.Panic(err)
	}

	err = f.Close()
	if err != nil {
		log.Panic(err)
	}
}
