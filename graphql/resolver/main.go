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
	"os"
	"github.com/google/uuid"
	"log"
	"strconv"
	"sync/atomic"
)

type EventClass int

const (
	Create EventClass = iota
	Update
	Delete
)

type GqlResolver struct {
	Client *Client
	watcher *Watcher
	subscribers chan *Subscriber
}

func NewGqlResolver(client *Client, watcher *Watcher) *GqlResolver {
	r := &GqlResolver{
		Client:      client,
		watcher:     watcher,
		subscribers: make(chan *Subscriber),
	}

	offset, err := strconv.ParseInt(os.Getenv("C_OFFSET"), 10, 32)
	if err != nil {
		log.Println(err)
		log.Println(os.Getenv("C_OFFSET"))
	}

	go r.broadcast(time.Duration(offset))

	return r
}

type Subscriber struct {
	stop   <-chan struct{}
	events interface{}
	uType EventType
	uClass EventClass
}

type unsubscribeEvent struct {
	id    string
	uType EventType
}

%s
%s
%s`, g.genResolvers(g.schema.Entities), g.genMutations(g.schema.Entities, g.schema.Functions), g.GenSubAll(g.schema.Entities))
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
