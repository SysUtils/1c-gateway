package generated

import (
	"encoding/json"
	"github.com/SysUtils/1c-gateway/shared"
	"os"
	"strings"
)

var fields = map[string]string{}
var types = map[string]string{}

func TranslateField(source string) string {
	if val, ok := types[source]; ok {
		return val
	}

	val := strings.Replace(source, "_", "", -1)
	fields[source] = val
	return val
}

func TranslateType(source string) string {
	if val, ok := types[source]; ok {
		return val
	}
	if strings.HasPrefix(source, "Collection(") && strings.HasSuffix(source, ")") {
		return "[]" + TranslateType(source[11:len(source)-1])
	}
	val := strings.Replace(source, "_", "", -1)
	types[source] = val
	return val
}

func SaveToFile(data map[string]string, path string) {
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	e := json.NewEncoder(f)
	// Encoding the map
	err = e.Encode(data)
	if err != nil {
		panic(err)
	}
	_ = f.Close()
}

func LoadFromFile(path string) map[string]string {
	res := make(map[string]string)
	f, err := os.Open(path)
	if err != nil {
		return res
	}

	e := json.NewDecoder(f)
	// Encoding the map
	err = e.Decode(&res)
	if err != nil {
		return res
	}
	_ = f.Close()
	return res
}

func _(schema shared.Schema) {
	types = LoadFromFile("types.dat")
	fields = LoadFromFile("fields.dat")
	for i, entity := range schema.Entities {
		println(i, "/", len(schema.Entities))
		TranslateType(entity.Name)
		for _, prop := range entity.Properties {
			TranslateField(prop.Name)
		}
		SaveToFile(types, "types.dat")
		SaveToFile(fields, "fields.dat")
	}

	for i, entity := range schema.Complexes {
		println(i, "/", len(schema.Entities))
		TranslateType(entity.Name)
		for _, prop := range entity.Properties {
			TranslateField(prop.Name)
		}
		SaveToFile(types, "types.dat")
		SaveToFile(fields, "fields.dat")
	}
}
