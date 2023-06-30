#!/bin/bash
Org="lishimeng"

# shellcheck disable=SC2046
Version=$(git describe --tags $(git rev-list --tags --max-count=1))
# shellcheck disable=SC2154
GitCommit=$(git log --pretty=format:"%h" -1)
BuildTime=$(date +%FT%T%z)

checkout_tag(){
  git checkout "${Version}"
}

build_image(){
  local Name=$1
  local AppPath=$2
  print_app_info "${Name}" "${AppPath}"

  docker build -t "${Org}/${Name}:${Version}" \
  --build-arg NAME="${Name}" \
  --build-arg VERSION="${Version}" \
  --build-arg BUILD_TIME="${BuildTime}" \
  --build-arg COMMIT="${GitCommit}" \
  --build-arg APP_PATH="${AppPath}" -f "./${AppPath}/Dockerfile" .
}

print_app_info(){
  local Name=$1
  local AppPath=$2
  echo "****************************************"
  echo "App:${Org}:${Name}"
  echo "Version:${Version}"
  echo "Commit:${GitCommit}"
  echo "Build:${BuildTime}"
  echo "Main_Path:${AppPath}"
  echo "****************************************"
  echo ""
}

push_image(){
  local Name=$1
  echo "****************************************"
  echo "Push:${Org}:${Name}:${Version}"
  echo "****************************************"
  echo ""
  docker push "${Org}/${Name}:${Version}"
}

#
#@name
#@path
build_all(){
  checkout_tag
  build_image 'owl-messager' 'cmd/owl-messager'
  build_image 'owl-saas' 'cmd/saas'
}

push_all(){
  push_image 'owl-messager'
  push_image 'owl-saas'
}

case  $1 in
    push)
		push_all
        ;;
    *)
		build_all
        ;;
esac

