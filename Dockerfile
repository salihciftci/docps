FROM golang:1.10
LABEL MAINTAINER="Salih Çiftçi"

WORKDIR /go/src/liman
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN curl -sSL https://get.docker.com/ | sh

EXPOSE 8080
CMD ["liman"]
