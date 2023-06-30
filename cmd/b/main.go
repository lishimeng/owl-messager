package main

import (
	"fmt"
	"github.com/lishimeng/app-starter/buildscript"
)

func main() {
	err := buildscript.Generate("owl-messager",
		"lishimeng",
		"cmd/owl-messager/main.go", false)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ok")
	}
}
