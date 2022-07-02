FROM node:18.4.0 as ui
WORKDIR /ui_build
ADD ui .
RUN npm install && npm run build

FROM golang:1.18 as build
ENV GOPROXY=https://goproxy.cn,direct
WORKDIR /release
ADD . .

COPY --from=ui /ui_build/dist /release/static/

RUN go mod download && go mod verify

#RUN go generate static/static.go

RUN go build -v -o owl-messager cmd/owl-messager/main.go


FROM ubuntu:22.04 as prod
EXPOSE 80/tcp
WORKDIR /
COPY --from=build /release/owl-messager /
ENTRYPOINT ["/owl-messager"]