FROM golang:alpine AS builder
LABEL MAINTAINER="Salih Çiftçi"

WORKDIR /go/src/liman
COPY . .

ENV GO111MODULE=on

RUN apk add -U --no-cache git gcc musl-dev && \
    go mod download && \
    go install -v ./...

FROM alpine:3.8

COPY --from=builder /go/bin/liman /liman
COPY --from=builder /go/src/liman/public /public
COPY --from=builder /go/src/liman/templates /templates

RUN apk add -U --no-cache ca-certificates docker

EXPOSE 8080
CMD /liman
