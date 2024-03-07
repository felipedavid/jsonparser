package main

import (
	"jsonparser/pkg/parser"
	"os"
)

func main() {
	buf, err := os.ReadFile("./cmd/input.json")
	if err != nil {
		panic(err)
	}

	//var value any
	//err = json.Unmarshal(buf, &value)
	//if err != nil {
	//	panic(err)
	//}

	var jsObject map[string]any
	err = parser.Parse(buf, &jsObject)
	if err != nil {
		panic(err)
	}

	os.Exit(0)
}
