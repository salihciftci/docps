
<h1 align="center">
  <br>
  <img src="https://img.salih.co/liman/logo.png" alt="Liman" width="200">
  <br>
  Liman
  <br>
</h1>

<h4 align="center">Basic docker monitoring web application. Written with Go.</h4>

<p align="center">
  <a href="https://travis-ci.org/salihciftci/liman">
    <img src="https://travis-ci.org/salihciftci/liman.svg?branch=master"
         alt="Travis-CI">
  </a>
  <a href="https://goreportcard.com/report/github.com/salihciftci/liman">
      <img src="https://goreportcard.com/badge/github.com/salihciftci/liman">
  </a>
  <a href="https://hub.docker.com/r/salihciftci/liman/">
    <img src="https://img.shields.io/docker/pulls/salihciftci/liman.svg">
  </a>
</p>

![screenshot](https://img.salih.co/liman/featured.png)

## Features

* Containers
* Stats
* Logs
* Networks
* Images
* Volumes

## Prerequisites

To keep our data safe Liman has login page. We need to export environment variable to set login password.

```
#In Docker use -e option
-e pass=MYLOGINPASS

#In host you need to export pass variable
$ export pass=MYLOGINPASS
```

## Installation

Liman works with host and docker container.

### Docker

```
docker pull salihciftci/liman
docker run -dit --name liman -e pass=PASS -v /var/run/docker.sock:/var/run/docker.sock salihciftci/liman
```

Note: the `-v /var/run/docker.sock:/var/run/docker.sock` option can be used in Linux environments only. 

### Host

You can [download](https://github.com/salihciftci/liman/releases) the lastest version of liman from releases or you can build with Go.

```
go get github.com/salihciftci/liman
make build
```

## Screenshots

 <table>
  <tr>
    <td><img src="https://img.salih.co/liman/dashboard.png"></td>
    <td><img src="https://img.salih.co/liman/containers.png"></td>
  </tr>
  <tr>
    <td><img src="https://img.salih.co/liman/logs.png"></td>
    <td><img src="https://img.salih.co/liman/stats.png"></td>
  </tr>
  <tr>
    <td><img src="https://img.salih.co/liman/images.png"></td>
    <td><img src="https://img.salih.co/liman/volumes.png"></td>
  </tr>
  <tr>
    <td><img src="https://img.salih.co/liman/networks.png"></td>
    <td><img src="https://img.salih.co/liman/login.png"></td>
  </tr>
</table> 

## License

MIT