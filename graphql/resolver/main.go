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
	"strings"
	"time"
	"os"
	"github.com/google/uuid"
	"log"
	"strconv"
	"sync/atomic"
)

type GqlResolver struct {
	Client *Client
	createEvents chan interface{}
	updateEvents chan interface{}
	deleteEvents  chan interface{}
	subscribers chan *Subscriber
}

func NewGqlResolver(client *Client) *GqlResolver {
	r := &GqlResolver{
		Client:      client,
		createEvents:      make(chan interface{}),
		updateEvents:      make(chan interface{}),
		deleteEvents:      make(chan interface{}),
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
	uType string
}

type unsubscribeEvent struct {
	id    string
	uType string
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
