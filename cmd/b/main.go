package main

import (
	"fmt"
	"github.com/lishimeng/app-starter/buildscript"
)

func main() {
	err := buildscript.Generate(
		"lishimeng",
		buildscript.Application{
			Name:    "owl-messager",
			AppPath: "cmd/owl-messager",
			HasUI:   false,
		},
		buildscript.Application{
			Name:    "owl-console",
			AppPath: "cmd/console",
			HasUI:   true,
		},
	)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ok")
	}
}
