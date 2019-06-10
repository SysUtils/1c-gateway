// Package for generate 1c grpc gateway
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	generated "github.com/SysUtils/1c-gateway"
	"github.com/SysUtils/1c-gateway/grpc"
	"github.com/SysUtils/1c-gateway/native"
	"github.com/SysUtils/1c-gateway/schema_cleaner"
	"github.com/SysUtils/1c-gateway/schema_loader"
	"io/ioutil"
	"log"
	"os"
)

func extractAsset(asset, path string) {
	data, err := generated.Asset(asset)
	if err != nil {
		log.Panic(err)
	}
	f, err := os.Create(path)
	if err != nil {
		log.Panic(err)
	}
	_, err = f.Write(data)
	if err != nil {
		log.Panic(err)
	}
	err = f.Close()
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	var host = flag.String("host", "localhost:8091", "address and port of 1C web service")
	var base = flag.String("base", "MyBase", "address and port of 1C web service")
	var username = flag.String("username", "Administrator", "username for 1C web service")
	var password = flag.String("password", "password", "password for 1C web service")
	flag.Parse()

	if host == nil || base == nil || username == nil || password == nil {
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
	extractAsset("static/binary.go", "odata/binary.go")
	extractAsset("static/boolean.go", "odata/boolean.go")
	extractAsset("static/client.go", "odata/client.go")
	extractAsset("static/client_entity.go", "odata/client_entity.go")
	extractAsset("static/datetime.go", "odata/datetime.go")
	extractAsset("static/double.go", "odata/double.go")
	extractAsset("static/errors.go", "odata/errors.go")
	extractAsset("static/float.go", "odata/float.go")
	extractAsset("static/guid.go", "odata/guid.go")
	extractAsset("static/int.go", "odata/int.go")
	extractAsset("static/int16.go", "odata/int16.go")
	extractAsset("static/int64.go", "odata/int64.go")
	extractAsset("static/interfaces.go", "odata/interfaces.go")
	extractAsset("static/stream.go", "odata/stream.go")
	extractAsset("static/string.go", "odata/string.go")
	extractAsset("static/time.go", "odata/time.go")
	extractAsset("static/where.go", "odata/where.go")
	extractAsset("static/grpc_helper.go", "odata/grpc_helper.go")

	loader := schema_loader.NewSchemaLoader(fmt.Sprintf("http://%s/%s/odata/standard.odata/$metadata", *host, *base), *username, *password)
	schema, err := loader.Load()
	if err != nil {
		log.Fatalln(err)
	}
	fields, _ := ioutil.ReadFile("fields.dat")
	types, _ := ioutil.ReadFile("types.dat")

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

	grpcGen := grpc.NewGenerator(*schema)
	grpcGen.NameMap = nameMap
	grpcGen.TypeMap = typeMap
	grpcGen.Generate()
}
