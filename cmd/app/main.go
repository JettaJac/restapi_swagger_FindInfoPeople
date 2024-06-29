package main

// go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=config.yaml ../../api.yaml
import (
	"fmt"
	// _ "github.com/oapi-codegen/oapi-codegen"
	// _ "github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen"
	_ "github.com/oapi-codegen/oapi-codegen/v2"
)

func main() {
	fmt.Println("hello world")
}
