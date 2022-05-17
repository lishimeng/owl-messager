#!/bin/bash
targetDir="pkg"
appName="owl-messager"
GOOS=windows GOPROXY=https://goproxy.io GOARCH=amd64 go build -o "${targetDir}/${appName}.exe" cmd/${appName}/main.go
