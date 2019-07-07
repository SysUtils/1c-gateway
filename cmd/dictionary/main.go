// Package for generate 1c odata client
package main

import (
	"flag"
	"fmt"
	generated "github.com/SysUtils/1c-gateway"
	"github.com/SysUtils/1c-gateway/dictionary"
	"github.com/SysUtils/1c-gateway/schema_loader"
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

	loader := schema_loader.NewSchemaLoader(fmt.Sprintf("http://%s/%s/odata/standard.odata/$metadata", *host, *base), *username, *password)
	schema, err := loader.Load()
	if err != nil {
		log.Fatalln(err)
	}
	gen := &dictionary.Generator{}
	gen.GenerateDictionary(schema)

}
