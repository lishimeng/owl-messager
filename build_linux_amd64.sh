#!/bin/bash
targetDir="pkg"
appName="owl-messager"
GOOS=linux GOPROXY=https://goproxy.io GOARCH=amd64 go build -o "${targetDir}/${appName}" cmd/${appName}/main.go
