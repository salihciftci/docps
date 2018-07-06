FROM golang:1.10
LABEL MAINTAINER="Salih Çiftçi"

WORKDIR /go/src/liman
COPY . .

RUN go get -d -v ./... && \
    go install -v ./... && \
    curl -sSL https://get.docker.com/ | sh

EXPOSE 8080
CMD ["liman"]
