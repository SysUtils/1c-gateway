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
}
