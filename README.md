# Screwdriver Client
[![Build Status][wercker-image]][wercker-url] [![Open Issues][issues-image]][issues-url]
[![Go Report Card][goreport-image]][goreport-url]

> The command line client for Screwdriver.

## Deprecated

**Please note that this code is no longer used by the screwdriver.cd team and has not been maintained in a while. You are welcome to use and/or contribute to it at your own risk.**

## Usage

```bash
$ client --help
$ client pipelines list  
$ client builds list
$ client jobs list
```


## Installation

### Building From Source
1. Ensure all dependencies (listed below) are met  
2. Get the code from github and add it to your GOPATH under src/github.com/screwdriver-cd/client  
3. Run ```bash $ swagger generate client <path to swagger spec>```
4. Run ```bash $ go get`` in the root directory of the Screwdriver Client  
5. Run ```bash $ go install ``` in the root directory of the Screwdriver Client  
6. Run ```bash $ sd --help``` to get started!  

## Testing

```bash
$ go get github.com/screwdriver-cd/client
$ go test -cover github.com/screwdriver-cd/client/...
```

## Dependencies  

1. Go-Swagger Binary - This is used in order to do the code generation following the Swagger (OpenAPI) Specification
2. Go-Swagger Tag 0.50.0 Dependency - This is a code dependency for the package go-swagger, in order to build from source, this is required.

## Caveats

* The code generation is picky about the packages it tries to pull down. Ensure that the code generated all uses the dependencies from go-swagger 

## License

Code licensed under the BSD 3-Clause license. See LICENSE file for terms.

[issues-image]: https://img.shields.io/github/issues/screwdriver-cd/client.svg
[issues-url]: https://github.com/screwdriver-cd/client/issues
[wercker-image]: https://app.wercker.com/status/822503b7af879d54018006aeafb317ae
[wercker-url]: https://app.wercker.com/project/bykey/822503b7af879d54018006aeafb317ae
[goreport-image]: https://goreportcard.com/badge/github.com/Screwdriver-cd/client
[goreport-url]: https://goreportcard.com/report/github.com/Screwdriver-cd/client`
