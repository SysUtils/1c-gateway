package main

import (
	"encoding/json"
	"fmt"
	"gitlab.com/zullpro/core/1cgqlgenerator/generator"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	//Startgqlserver()
	if len(os.Args) != 7 {
		log.Fatal("usage: 1cclientgenerator host 1СBase username password typemap fieldmap")
	}
	types, err := ioutil.ReadFile(os.Args[5])
	if err != nil {
		log.Fatal(err)
	}
	fields, err := ioutil.ReadFile(os.Args[6])
	if err != nil {
		log.Fatal(err)
	}

	typemap := make(map[string]string)
	err = json.Unmarshal(types, &typemap)
	if err != nil {
		log.Fatal(err)
	}

	fieldmap := make(map[string]string)
	err = json.Unmarshal(fields, &fieldmap)
	if err != nil {
		log.Fatal(err)
	}

	gen, err := generator.NewGenerator(fmt.Sprintf("http://%s/%s/odata/standard.odata/$metadata", os.Args[1], os.Args[2]), os.Args[3], os.Args[4], typemap, fieldmap)
	if err != nil {
		log.Fatal(err)
	}
	err = gen.Bootstrap()
	if err != nil {
		log.Fatal(err)
	}
	gen.ExtractTypes()
	gen.ExtractAssociations()
	gen.PrepareComplexes()
	gen.PrepareEntities()
	gen.PrepareMethods()
	err = gen.Write()
	if err != nil {
		log.Fatal(err)
	}
}
