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

type TokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

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
	if tokenManager != nil {
		http.Handle("/token", disableCors(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

			params := TokenRequest{}

			defer r.Body.Close()
			dec := json.NewDecoder(r.Body)
			for {
				if err := dec.Decode(&params); err == io.EOF {
					break
				} else if err != nil {
					log.Println(err)
					result.Error = err.Error()
					result.Code = AuthError
				}
			}

			token, err := tokenManager.Get(params.Username, params.Password)
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
		})))

		http.Handle("/metrics", promhttp.Handler())
		http.Handle("/graphql", disableCors(Secure(graphQLHandler, tokenManager)))
	} else {
		http.Handle("/metrics", promhttp.Handler())
		http.Handle("/graphql", graphQLHandler)
	}

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

func disableCors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, Accept-Encoding")
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Max-Age", "86400")
			w.WriteHeader(http.StatusOK)
			return
		}
		h.ServeHTTP(w, r)
	})
}
