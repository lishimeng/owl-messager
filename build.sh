#!/bin/bash
targetDir="pkg"
appName="owl-messager"

TAG=$(git describe --tags $(git rev-list --tags --max-count=1))

help
echo "Version:$TAG"

build_application

build_application() {
  docker build -t lishimeng/owl-messager:$TAG .
}

help() {
  echo 'build owl-messager'
}