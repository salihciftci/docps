# Liman
[![Build Status](https://travis-ci.org/salihciftci/liman.svg?branch=master)](https://travis-ci.org/salihciftci/liman) [![Go Report Card](https://goreportcard.com/badge/github.com/salihciftci/liman)](https://goreportcard.com/report/github.com/salihciftci/liman)

![alt text](https://img.salih.co/liman/v0.6/logo.png "Liman")

Web application for monitoring docker. Monitor docker inside the docker. Written in Go.

----
![screenshot](https://img.salih.co/liman/v0.6/dashboard.png)

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

### Persist root password
Liman stores the root password setup during install in the FileSystem, so if you want to persist the 
password between containers, you should mount the /data folder to a volume.

The below run command will run Liman with the data folder mounted to a local data folder in the current
directory.
```
docker run -it -v /var/run/docker.sock:/var/run/docker.sock -v ./data:/data salihciftci/liman
```

## API Usage

Basic usage:
```
curl -i http://localhost:8080/api/status?key=XXX
```

More examples and all end points can be found in [wiki](https://github.com/salihciftci/liman/wiki/API-Usage).


| Screenshots | | |
|:-------------:|:-------:|:-------:|
|![Dashboard](https://img.salih.co/liman/v0.6/dashboard.png)|![Containers](https://img.salih.co/liman/v0.6/containers.png)|![Images](https://img.salih.co/liman/v0.6/images.png)|
|![Stats](https://img.salih.co/liman/v0.6/stats.png)|![Volumes](https://img.salih.co/liman/v0.6/volumes.png)|![Networks](https://img.salih.co/liman/v0.6/networks.png)|
|![Logs](https://img.salih.co/liman/v0.6/logs.png)|![Notifications](https://img.salih.co/liman/v0.6/notifications.png)|![Settings](https://img.salih.co/liman/v0.6/settings.png)|

## License
MIT
