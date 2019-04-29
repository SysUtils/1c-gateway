package main

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

func main() {
	if len(os.Args) != 5 {
		log.Fatal("usage: 1cclientgenerator host 1СBase username password")
	}

	loader := schema_loader.NewSchemaLoader(fmt.Sprintf("http://%s/%s/odata/standard.odata/$metadata", os.Args[1], os.Args[2]), os.Args[3], os.Args[4])
	schema, err := loader.Load()
	if err != nil {
		log.Fatalln(err)
	}
	graphqlGen := graphql.NewGenerator(*schema)
	fields, _ := ioutil.ReadFile("fields.dat")
	types, _ := ioutil.ReadFile("types.dat")
	json.Unmarshal(fields, &graphqlGen.NameMap)
	json.Unmarshal(types, &graphqlGen.TypeMap)

	graphqlGen.Start()

	clientGen := native.NewGenerator(*schema)
	json.Unmarshal(fields, &clientGen.NameMap)
	json.Unmarshal(types, &clientGen.TypeMap)

	clientGen.Start()
	//	StartServer()
}
