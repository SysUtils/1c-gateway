// Package graphql/server provides GraphQL server for gateway
package odata

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/graph-gophers/graphql-transport-ws/graphqlws"
	"net/http"
)

// Starts the server using the specified address, scheme and resolver
func HandlerFunc(resolver interface{}, opts []graphql.SchemaOpt) (http.HandlerFunc, error) {
	opts = append(opts, graphql.UseFieldResolvers())

	schema, err := graphql.ParseSchema(gqlSchema, resolver, opts...)
	if err != nil {
		return nil, err
	}

	graphQLHandler := graphqlws.NewHandlerFunc(schema, &relay.Handler{Schema: schema})

	return graphQLHandler, nil
}
