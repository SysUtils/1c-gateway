// Package graphql/resolver provides generator for GraphQL gateway's resolver
package resolver

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

// NewGenerator returns initialized generator
func NewGenerator(schema shared.Schema) *Generator {
	return &Generator{schema: schema, TypeMap: make(map[string]string), NameMap: make(map[string]string)}
}

// Generate generates the GraphQL resolver and writes it to ./odata folder
func (g *Generator) Generate() {
	data := fmt.Sprintf(`package odata

%s`, g.genArgs(g.schema.Entities))
	g.writeGofile("Resolver_args.go", data)

	data = fmt.Sprintf(
		`package odata

import (
	"context"
	"time"
)

type GqlResolver struct {
	Client *Client
	events     chan *Event
	subscriber chan *Subscriber
}

func NewGqlResolver(client *Client) *GqlResolver {
	r := &GqlResolver{
		Client:         client,
		events:     make(chan *Event),
		subscriber: make(chan *Subscriber),
	}

	go r.broadcast()

	return r
}

type Subscriber struct {
	stop   <-chan struct{}
	events chan<- *Event
}

type Event struct {
	id        string
	eventType string
	payload   interface{}
}

func (r *GqlResolver) broadcast() {
	subscribers := map[string]*Subscriber{}
	unsubscribe := make(chan string)

	// NOTE: subscribing and sending events are at odds.
	for {
		select {
		case id := <-unsubscribe:
			delete(subscribers, id)
		case s := <-r.subscriber:
			subscribers[randomID()] = s
		case e := <-r.events:
			for id, s := range subscribers {
				go func(id string, s *Subscriber) {
					select {
					case <-s.stop:
						unsubscribe <- id
						return
					default:
					}

					select {
					case <-s.stop:
						unsubscribe <- id
					case s.events <- e:
					case <-time.After(time.Second):
					}
				}(id, s)
			}
		}
	}
}


%s
%s`, g.genResolvers(g.schema.Entities), g.genMutations(g.schema.Entities))
	g.writeGofile("Resolver.go", data)
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
