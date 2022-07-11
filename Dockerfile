FROM node:18.4.0 as ui
WORKDIR /ui_build
ADD ui .
RUN npm install && npm run build

FROM golang:1.18 as build
ARG APP_VERSION
ENV GOPROXY=https://goproxy.cn,direct
WORKDIR /release
ADD . .

COPY --from=ui /ui_build/dist /release/static/

RUN go mod download && go mod verify

RUN go build -v --ldflags "-X cmd.Version=${APP_VERSION}" -o owl-messager cmd/owl-messager/main.go


FROM ubuntu:22.04 as prod
EXPOSE 80/tcp
WORKDIR /
COPY --from=build /release/owl-messager /
ENTRYPOINT ["/owl-messager"]