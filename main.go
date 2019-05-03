package main

//go:generate go-bindata -pkg $GOPACKAGE static/

import (
	"encoding/json"
	"fmt"
	"gitlab.com/zullpro/core/1cclientgenerator.git/graphql"
	"gitlab.com/zullpro/core/1cclientgenerator.git/native"
	"gitlab.com/zullpro/core/1cclientgenerator.git/schema_loader"
	"io/ioutil"
	"log"
	"os"
)

func extractAsset(asset, path string) {
	data, err := Asset(asset)
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
	if len(os.Args) != 5 {
		log.Fatal("usage: 1cclientgenerator host 1СBase username password")
	}

	loader := schema_loader.NewSchemaLoader(fmt.Sprintf("http://%s/%s/odata/standard.odata/$metadata", os.Args[1], os.Args[2]), os.Args[3], os.Args[4])
	schema, err := loader.Load()
	if err != nil {
		log.Fatalln(err)
	}
	fields, _ := ioutil.ReadFile("fields.dat")
	types, _ := ioutil.ReadFile("types.dat")
	err = os.Mkdir("odata", os.ModePerm)
	if err != nil && err != os.ErrExist {
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

	graphqlGen := graphql.NewGenerator(*schema)
	json.Unmarshal(fields, &graphqlGen.NameMap)
	json.Unmarshal(types, &graphqlGen.TypeMap)

	graphqlGen.Start()

	clientGen := native.NewGenerator(*schema)
	json.Unmarshal(fields, &clientGen.NameMap)
	json.Unmarshal(types, &clientGen.TypeMap)
	clientGen.Start()
}
