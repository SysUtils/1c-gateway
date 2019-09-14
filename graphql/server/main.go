// Package graphql/server provides GraphQL server for gateway
package server

import (
	"encoding/json"
	"fmt"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/graph-gophers/graphql-transport-ws/graphqlws"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"log"
	"net/http"
)

// Starts the server using the specified address, scheme and resolver
func Start(addr string, schemaBlob []byte, resolver interface{}, poolSize int, tokenManager ITokenManager) error {
	closer := initJaeger("1c-graphql-gateway")
	httpMetric := NewHttpMetric()
	defer closer.Close()
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers(), graphql.Tracer(NewTracer(httpMetric)), graphql.MaxParallelism(poolSize)}

	schema, err := graphql.ParseSchema(string(schemaBlob), resolver, opts...)
	if err != nil {
		return err
	}

	graphQLHandler := graphqlws.NewHandlerFunc(schema, &relay.Handler{Schema: schema})

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write(playground)
		if err != nil {
			log.Panic(err)
		}
	}))

	http.Handle("/token", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		const (
			Ok        = 0
			AuthError = 1
		)

		type res struct {
			Code  int
			Error string `json:",omitempty"`
			Token string `json:",omitempty"`
		}

		result := res{Code: Ok}
		login := r.FormValue("login")
		password := r.FormValue("password")
		token, err := tokenManager.Get(login, password)
		if err != nil {
			result.Error = err.Error()
			result.Code = AuthError
		}
		result.Token = token
		data, err := json.Marshal(result)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		w.Write(data)
	}))

	http.Handle("/metrics", Secure(promhttp.Handler(), tokenManager))

	http.Handle("/graphql", Secure(graphQLHandler, tokenManager))
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
