FROM golang:1.19-alpine AS build
COPY  . /demo
WORKDIR /demo

RUN go build -o app2 .

FROM alpine:latest
COPY --from=build /demo/app2 .

ENTRYPOINT ["./app2"]
