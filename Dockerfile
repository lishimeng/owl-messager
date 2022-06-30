FROM golang:1.18 as build
ENV GOPROXY=https://goproxy.cn,direct
WORKDIR /release
ADD . .
RUN "go build -o owl-messager cmd/owl-messager/main.go"


FROM ubuntu:22.04 as prod
EXPOSE 80/tcp
WORKDIR /
COPY --from:build /release/owl-messager /
ENTRYPOINT ["/owl-messager"]