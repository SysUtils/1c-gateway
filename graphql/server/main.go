package server

import (
	"fmt"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/graph-gophers/graphql-transport-ws/graphqlws"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"log"
	"net/http"
)

func Start(addr string, schemaBlob []byte, resolver interface{}) error {
	closer := initJaeger("1c-graphql-gateway")
	defer closer.Close()
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}

	schema, err := graphql.ParseSchema(string(schemaBlob), resolver, opts...)
	if err != nil {
		return err
	}

	graphQLHandler := graphqlws.NewHandlerFunc(schema, &relay.Handler{Schema: schema})

	http.Handle("/graphql", graphQLHandler)
	return http.ListenAndServe(addr, nil)
}

func initJaeger(service string) io.Closer {
	cfg, err := config.FromEnv()
	if err != nil {
		log.Panic(err)
	}
	if service != "" {
		cfg.ServiceName = service
	}
	cfg.Reporter.LogSpans = true
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	opentracing.SetGlobalTracer(
		tracer,
	)
	return closer
}
