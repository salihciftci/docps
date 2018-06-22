
<h1 align="center">
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
* Images
* Networks
* Volumes
* API

## Installation

[Download](https://github.com/salihciftci/liman/releases) and run the latest binary or use it inside the Docker.

### Docker

```
docker run -it -v /var/run/docker.sock:/var/run/docker.sock salihciftci/liman
```

Note: the `-v /var/run/docker.sock:/var/run/docker.sock` option can be used in Linux environments only. 

## API Usage

API only allow **GET** requests, other requests will rejected.

We need a key to use Liman API. We can have that from settings in home page.

Basic usage:
```
curl -i http://localhost:8080/api/status?key=xxx
```

More examples and all end points can be found in [wiki](https://github.com/salihciftci/liman/wiki/API-Usage).

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
</table> 

## License

MIT