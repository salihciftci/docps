FROM golang:alpine AS builder
LABEL MAINTAINER="Salih Çiftçi"

WORKDIR /go/src/liman
COPY . .

RUN apk add -U --no-cache git && \
    go get -d -v ./... && \
    go install -v ./...


FROM alpine:3.8

COPY --from=builder /go/bin/liman /liman
COPY --from=builder /go/src/liman/public /public
COPY --from=builder /go/src/liman/templates /templates

RUN apk add -U --no-cache ca-certificates docker

EXPOSE 8080
CMD /liman
