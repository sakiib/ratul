FROM golang:1.19-alpine AS build
COPY  . /demo
WORKDIR /demo

RUN go build -o app1 .

FROM alpine:latest
COPY --from=build /demo/app1 .

ENTRYPOINT ["./app1"]
