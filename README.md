# docps
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com) [![Build Status](https://travis-ci.org/salihciftci/docps.svg?branch=master)](https://travis-ci.org/salihciftci/docps) [![Go Report Card](https://goreportcard.com/badge/github.com/salihciftci/docps)](https://goreportcard.com/report/github.com/salihciftci/docps) [![Coverage Status](https://coveralls.io/repos/github/salihciftci/docps/badge.svg)](https://coveralls.io/github/salihciftci/docps)

docps is basic docker monitoring web application. Written with Go.

### Features
- Stats
- Process State
- Images
- Volumes

## Installation
### Host

``` bash
$ go get github.com/salihciftci/docps
$ go build
$ ./docps
```

### Docker

``` bash
$ docker build -t docps .
$ docker run -dit --name docps --restart always -p 8080:8080 -v /var/run/docker.sock:/var/run/docker.sock docps
```
Note: the `-v /var/run/docker.sock:/var/run/docker.sock` option can be used in Linux environments only. 

## Screenshots

![](https://img.salih.co/docps/stats.png)
![](https://img.salih.co/docps/containers.png)
![](https://img.salih.co/docps/images.png)
