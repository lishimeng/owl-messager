#!/bin/bash
APP_NAME="owl-messager"

# shellcheck disable=SC2046
TAG=$(git describe --tags $(git rev-list --tags --max-count=1))

build_application(){
  docker build -t lishimeng/owl-messager:"$TAG" .
}

help_print(){
  echo "build ${APP_NAME}"
  echo "Version:$TAG"
}

help_print
build_application