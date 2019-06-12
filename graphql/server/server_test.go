package server

import (
	"github.com/SysUtils/1c-gateway/graphql/server/mock"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	schema, err := mock.Asset("schema.graphql")
	if err != nil {
		t.Fatal(err)
	}
	go func() {
		if err := Start(":4000", schema, &mock.Resolver{}); err != nil {
			t.Fatal(err)
		}
	}()
	time.Sleep(time.Minute)
}
