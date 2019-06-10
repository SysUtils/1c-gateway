// Package grpc/server provides generator for GRPC gateway's server
package server

import (
	"github.com/SysUtils/1c-gateway/shared"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

type Generator struct {
	TypeMap        map[string]string
	NameMap        map[string]string
	ReverseNameMap map[string]string
	Associations   map[string]map[string]string
	schema         shared.Schema
}

// NewGenerator returns initialized generator
func NewGenerator(schema shared.Schema) *Generator {
	return &Generator{schema: schema, TypeMap: make(map[string]string), NameMap: make(map[string]string), Associations: map[string]map[string]string{}}
}

// Generate generates the grpc server and writes it to ./odata folder
func (g *Generator) Start() {
	cmd := exec.Command("protoc", "--go_out=plugins=grpc:.", "odata/grpc.proto")
	cmd.Stdout = os.Stdout // cmd.Stdout -> stdout
	cmd.Stderr = os.Stderr // cmd.Stderr -> stderr

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	source, err := ioutil.ReadFile("odata/grpc.pb.go")
	if err != nil {
		log.Fatal(err)
	}
	strSource := string(source)
	strSource = g.fixJson(strSource)
	err = ioutil.WriteFile("odata/grpc.pb.go", []byte(strSource), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}
