# ogc-capabilities-generator

[![Build](https://github.com/PDOK/ogc-capabilities-generator/actions/workflows/build-and-publish-image.yml/badge.svg)](https://github.com/PDOK/ogc-capabilities-generator/actions/workflows/build-and-publish-image.yml)
[![Lint](https://github.com/PDOK/ogc-capabilities-generator/actions/workflows/lint.yml/badge.svg)](https://github.com/PDOK/ogc-capabilities-generator/actions/workflows/lint.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/PDOK/ogc-capabilities-generator)](https://goreportcard.com/report/github.com/PDOK/ogc-capabilities-generator)
[![GitHub license](https://img.shields.io/github/license/PDOK/ogc-capabilities-generator)](https://github.com/PDOK/ogc-capabilities-generator/blob/master/LICENSE)
[![Docker Pulls](https://img.shields.io/docker/pulls/pdok/ogc-capabilities-generator.svg)](https://hub.docker.com/r/pdok/ogc-capabilities-generator)

## What will it do

This application will give the user/developer 'full' control in the generation
of a GetCapabilities document, instead of letting it be dictated by a specific
application implementation. By configuring the input YAML files, one can
enable/disable certain operations, constraints or filters in the advertised OGC
GetCapabilities response document. This generated file can then be used in a
static http server like Apache/NGINX/Lighttpd/and so on. This solution make it
also flexible to deploy a OGC in a 'micro service' like environment such as K8s.

## How to run

Building the binary

```sh
go build .
```

When the binary is built the application can be run with parsing the
configuration in two ways:

1. parsing the configuration through -c {path/to/config.yaml}
2. setting the ENV var SERVICECONFIG={path/to/config.yaml}

```sh
XML_CATALOG_FILES=./xml-catalog/ogc-catalog.xml ./ogc-capabilities-generator -c ./examples/config/wmts_1_0_0.yaml
```

In dev mode through `go run`:

```sh
XML_CATALOG_FILES=./xml-catalog/ogc-catalog.xml go run . -c ./examples/config/wms_1_3_0.yaml
XML_CATALOG_FILES=./xml-catalog/ogc-catalog.xml go run . -c ./examples/config/wfs_2_0_0.yaml
XML_CATALOG_FILES=./xml-catalog/ogc-catalog.xml go run . -c ./examples/config/wmts_1_0_0.yaml
```

## How to configure

The ogc-capabilities-generator makes usage of the
[ogc-specifications](https://github.com/PDOK/ogc-specifications) package

## Test

```sh
XML_CATALOG_FILES=./xml-catalog/ogc-catalog.xml go test ./... -covermode=atomic -v
```

## Docker

```sh
docker build -t pdok/ogc-capabilities-generator .

docker run --rm -d -v `pwd`/examples/config:/config -v `pwd`:/output \
-e SERVICECONFIG=/config/wfs_2_0_0.yaml --name ogc pdok/ogc-capabilities-generator
```


## XML catalog

### Install
You should have libxml2-utils installed: 

```shell
apt-get install libxml2-dev libxml2-utils
```

### Build

Download the relevant xsd schemas and its children into the `xml-catalog/schemas` directory
([this tool](https://github.com/n-a-t-e/xsd_download) does the job).

Add a line to the `build-catalog.sh` script for each downloaded schema url and its relative file location.
Execute it:

```shell
./build-catalog.sh
```

### Usage
Now you can use the command with the catalog:

```
XML_CATALOG_FILES=./xml-catalog/ogc-catalog.xml ./ogc-capabilities-generator -c ./examples/config/wmts_1_0_0.yaml
```