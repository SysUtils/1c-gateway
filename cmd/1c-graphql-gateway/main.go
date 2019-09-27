// Package for generate 1c graphql gateway
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/SysUtils/1c-gateway/graphql"
	"github.com/SysUtils/1c-gateway/native"
	"github.com/SysUtils/1c-gateway/schema_cleaner"
	"github.com/SysUtils/1c-gateway/schema_loader"
	"github.com/SysUtils/1c-gateway/shared"
	"github.com/SysUtils/1c-gateway/subscriptions"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var host = flag.String("host", "localhost:8091", "address and port of 1C web service")
	var base = flag.String("base", "MyBase", "address and port of 1C web service")
	var username = flag.String("username", "Administrator", "username for 1C web service")
	var password = flag.String("password", "password", "password for 1C web service")
	flag.Parse()

	if host == nil || base == nil || username == nil || password == nil || *host == "" || *base == "" || *username == "" || *password == "" {
		log.Fatal("arguments host, base, username and password is required")
	}

	err := os.Mkdir("odata", os.ModePerm)
	if err != nil && err.Error() != "mkdir odata: file exists" {
		log.Fatalln(err)
	}
	// static/binary.go
	// static/boolean.go
	// static/client.go
	// static/client_entity.go
	// static/datetime.go
	// static/double.go
	// static/errors.go
	// static/float.go
	// static/guid.go
	// static/int.go
	// static/int16.go
	// static/int64.go
	// static/interfaces.go
	// static/stream.go
	// static/string.go
	// static/time.go
	// static/where.go
	shared.ExtractAsset("static/binary.go", "odata/binary.go")
	shared.ExtractAsset("static/boolean.go", "odata/boolean.go")
	shared.ExtractAsset("static/client.go", "odata/client.go")
	shared.ExtractAsset("static/client_entity.go", "odata/client_entity.go")
	shared.ExtractAsset("static/datetime.go", "odata/datetime.go")
	shared.ExtractAsset("static/double.go", "odata/double.go")
	shared.ExtractAsset("static/errors.go", "odata/errors.go")
	shared.ExtractAsset("static/float.go", "odata/float.go")
	shared.ExtractAsset("static/guid.go", "odata/guid.go")
	shared.ExtractAsset("static/int.go", "odata/int.go")
	shared.ExtractAsset("static/int16.go", "odata/int16.go")
	shared.ExtractAsset("static/int64.go", "odata/int64.go")
	shared.ExtractAsset("static/interfaces.go", "odata/interfaces.go")
	shared.ExtractAsset("static/stream.go", "odata/stream.go")
	shared.ExtractAsset("static/string.go", "odata/string.go")
	shared.ExtractAsset("static/time.go", "odata/time.go")
	shared.ExtractAsset("static/where.go", "odata/where.go")
	shared.ExtractAsset("static/grpc_helper.go", "odata/grpc_helper.go")

	loader := schema_loader.NewSchemaLoader(fmt.Sprintf("http://%s/%s/odata/standard.odata/$metadata", *host, *base), *username, *password)
	schema, err := loader.Load()
	if err != nil {
		log.Fatalln(err)
	}
	fields, _ := ioutil.ReadFile("fields.json")
	types, _ := ioutil.ReadFile("types.json")

	nameMap := make(map[string]string)
	typeMap := make(map[string]string)
	if err := json.Unmarshal(fields, &nameMap); err != nil {
		log.Panic(err)
	}
	if err := json.Unmarshal(types, &typeMap); err != nil {
		log.Panic(err)
	}

	schema = schema_cleaner.Clean(schema, typeMap)

	clientGen := native.NewGenerator(*schema)
	clientGen.NameMap = nameMap
	clientGen.TypeMap = typeMap
	clientGen.Generate()

	graphqlGen := graphql.NewGenerator(*schema)
	graphqlGen.NameMap = nameMap
	graphqlGen.TypeMap = typeMap
	graphqlGen.Start()

	watchersGen := subscriptions.NewGenerator(*schema)
	watchersGen.NameMap = nameMap
	watchersGen.TypeMap = typeMap
	watchersGen.Start()
}
