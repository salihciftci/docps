# Liman
[![Build Status](https://travis-ci.org/salihciftci/liman.svg?branch=master)](https://travis-ci.org/salihciftci/liman) [![Go Report Card](https://goreportcard.com/badge/github.com/salihciftci/liman)](https://goreportcard.com/report/github.com/salihciftci/liman)

<img alt="logo" src="https://raw.githubusercontent.com/salihciftci/liman/master/public/img/liman.png" width=200>

Web application for monitoring docker. Monitor docker inside the docker. Written in Go.

----

## Features

* Monitoring docker
    * Containers
    * Logs
    * Images
    * Stats
    * Volumes
    * Networks
* Notifications
* Restful API

## Installation

[Download](https://github.com/salihciftci/liman/releases) and run the latest binary or ship with Docker.


```
docker run -it -v /var/run/docker.sock:/var/run/docker.sock salihciftci/liman
```

Note: the `-v /var/run/docker.sock:/var/run/docker.sock` option can be used in Linux environments only. 

## API Usage

Basic usage:
```
curl -i http://localhost:8080/api/status?key=XXX
```

More examples and all end points can be found in [wiki](https://github.com/salihciftci/liman/wiki/API-Usage).

## License
MIT
