package main

import (
	"fmt"
	"gitlab.com/zullpro/core/1cclientgenerator.git/client"
	"gitlab.com/zullpro/core/1cclientgenerator.git/schema_loader"
	"log"
	"os"
)

func main() {
	//Startgqlserver()
	if len(os.Args) != 5 {
		log.Fatal("usage: 1cclientgenerator host 1СBase username password")
	}

	loader := schema_loader.NewSchemaLoader(fmt.Sprintf("http://%s/%s/odata/standard.odata/$metadata", os.Args[1], os.Args[2]), os.Args[3], os.Args[4])
	schema, err := loader.Load()
	if err != nil {
		log.Fatalln(err)
	}
	clientGen := client.NewGenerator(*schema)
	clientGen.Start()
	print()

}
