package main

import (
	"fmt"
	"github.com/lishimeng/app-starter/buildscript"
)

func main() {
	err := buildscript.Generate("owl-messager",
		"lishimeng",
		"cmd/owl-messager/main.go", true)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ok")
	}
}
