# ogc-capabilities-generator

![GitHub license](https://img.shields.io/github/license/PDOK/ogc-capabilities-generator)
[![GitHub release](https://img.shields.io/github/release/PDOK/ogc-capabilities-generator.svg)](https://github.com/PDOK/ogc-capabilities-generator/releases)
[![Go Report Card](https://goreportcard.com/badge/PDOK/ogc-capabilities-generator)](https://goreportcard.com/report/PDOK/ogc-capabilities-generator)
[![Docker Pulls](https://img.shields.io/docker/pulls/pdok/ogc-capabilities-gen.svg)](https://hub.docker.com/r/pdok/ogc-capabilities-gen)


## What will it do

This application will give the user/developer 'full' control in the generation of a GetCapabilities document, instead of letting it be dictated by a specific application implementation. By configuring the input YAML files, one can enable/disable certain operations, constraints or filters in the advertised OGC GetCapabilities response document. This generated file can then be used in a static http server like Apache/NGINX/Lighttpd/and so on. This solution make it also flexible to deploy a OGC in a 'micro service' like environment such as K8s.

## How to run

```go
go run . -c ./examples/kadastralekaart_v4.yaml
```

When the binary is build the application can be run with parsing the configuration in two ways:

1. parsing the configuration through -c {path/to/config.yaml}
2. setting the ENV var SERVICECONFIG={path/to/config.yaml}

## Test

```go
go test ./... -covermode=atomic
```

## Docker

```docker
docker build -t pdok/ogc-capabilities-generator .
docker run -d -v `pwd`/examples:/config -v `pwd`:/output -e SERVICECONFIG=/config/kadastralekaart_v4.yaml --name ocg pdok/ogc-capabilities-generator
```
