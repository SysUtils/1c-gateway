package shared

import (
	generated "github.com/SysUtils/1c-gateway"
	"log"
	"os"
)

func ExtractAsset(asset, path string) {
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
