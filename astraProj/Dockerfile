# syntax=docker/dockerfile:1

FROM golang:alpine AS builder

WORKDIR /app
COPY go.mod ./
COPY go.sum ./

RUN go clean -modcache
RUN go mod download
RUN go mod tidy && go mod vendor
#RUN apk update && apk upgrade -U -a
# RUN apt-get install -y ca-certificates
# RUN apk --no-cache add ca-certificates
# COPY ca-bundle.crt /etc/ssl/certs/ca-bundle.crt
# COPY ca-bundle.trust.crt /etc/ssl/certs/ca-bundle.trust.crt
COPY ./certs/* /etc/ssl/certs/
RUN apk --no-cache --no-progress --no-scripts add --update --allow-untrusted ca-certificates


COPY . ./
WORKDIR /app/
RUN  go build -o /main

#final stage
FROM alpine:latest
#RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app
ENTRYPOINT /app
LABEL Name=astraproj Version=0.0.1
EXPOSE 8120