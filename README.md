# docps
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com) [![Build Status](https://travis-ci.org/salihciftci/docps.svg?branch=master)](https://travis-ci.org/salihciftci/docps) [![Go Report Card](https://goreportcard.com/badge/github.com/salihciftci/docps)](https://goreportcard.com/report/github.com/salihciftci/docps)

docps is basic docker process state web application. Written with Go.

It's basically gathering `docker ps -a` but without Commands and Ports and serving.

### Why?
-Too lazy to ssh.
-Want to check with my phone.

## Installation

### Dockerfile

``` bash
$ docker build -t docps .
$ docker run -dit --name docps --restart always -p 8080:8080 -v /var/run/docker.sock:/var/run/docker.sock docps
```
Note: the `-v /var/run/docker.sock:/var/run/docker.sock` option can be used in Linux environments only. 

## Screenshots

<table>
<tr>
<td valign="top">
<img src="https://raw.githubusercontent.com/salihciftci/docps/master/screenshots/web.png">
</td>
<td valign="top">
<img src="https://raw.githubusercontent.com/salihciftci/docps/master/screenshots/mobil.png" height="290">
</td>
</tr>
</table>
