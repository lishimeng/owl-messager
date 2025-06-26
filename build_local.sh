#!/bin/bash

Version=''
GitCommit=''
BuildTime=''

Base="github.com/lishimeng/app-starter/version"
GOPROXY=https://goproxy.cn,direct
if [ -n "${GOOS}" ]; then
  echo "custom GOOS"
else
  GOOS=linux
fi
echo "GOOS:${GOOS}"

if [ -n "${GOARCH}" ]; then
  echo "custom GOARCH"
else
  GOARCH=amd64
fi
echo "GOARCH:${GOARCH}"

meta(){
  # shellcheck disable=SC2046
  Version=$(git describe --tags $(git rev-list --tags --max-count=1))
  # shellcheck disable=SC2154
  GitCommit=$(git log --pretty=format:"%h" -1)
  BuildTime=$(date +%FT%T%z)
}

meta_dev(){
  # shellcheck disable=SC2046
  Version="snapshot_"$(git log --pretty=format:"%s" -1)
  # shellcheck disable=SC2154
  GitCommit=$(git log --pretty=format:"%h" -1)
  BuildTime=$(date +%FT%T%z)
}

checkout_tag(){
  git checkout "${Version}"
}

common(){
  echo ""
}

print_app_info(){
  local Name=$1
  local AppPath=$2
  echo "****************************************"
  echo "App:${Name}"
  echo "Version:${Version}"
  echo "Commit:${GitCommit}"
  echo "Build:${BuildTime}"
  echo "Main_Path:${AppPath}"
  echo "Build_Os:${GOOS}"
  echo "Build_Arch:${GOARCH}"
  echo "GoProxy:${GOPROXY}"
  echo "****************************************"
  echo ""
}

build(){
  local Name=$1
  local AppPath=$2
  local LdFlags=" \
            -X ${Base}.AppName=${Name} \
            -X ${Base}.Version=${Version} \
            -X ${Base}.Commit=${GitCommit} \
            -X ${Base}.Build=${BuildTime} \
            "
  go mod tidy && go mod vendor
  GOOS="${GOOS}" GOARCH="${GOARCH}" go build -v --ldflags "${LdFlags} -X ${Base}.Compiler=$(go version | sed 's/[ ][ ]*/_/g')" -o "${Name}" "${AppPath}"/main.go
}

build_app(){
  local Name=$1
  local AppPath=$2
  if [ "${GOOS}" == "windows" ]; then
      Name="${Name}.exe"
  fi
  print_app_info "${Name}" "${AppPath}"
  build "${Name}" "${AppPath}"
}

build_all(){
  common
  build_app 'owl-messager' 'cmd/owl-messager'
  build_app 'owl-console' 'cmd/console'
}

# command
case  $1 in
    release)
    meta
    checkout_tag
		build_all
        ;;
    *)
		meta_dev
		build_all
        ;;
esac