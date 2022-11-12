# doctools

A set of tools to automate standard project documentation chores.


## Table of Contents

	- [Quick Start](#quick-start)
	    - [Building](#building)
	    - [Running](#running)
	    - [Testing](#testing)
	- [Installation](#installation)


## Quick Start


### Building

```console
$ go build -o ./ doctools/cmd/...
```


### Running

```console
$ go run doctools/cmd/dt
$ go run doctools/cmd/dt-adr
$ go run doctools/cmd/dt-chglg
$ go run doctools/cmd/dt-license
$ go run doctools/cmd/dt-rdme
```


### Testing

```console
$ go test ./...
```


## Installation

```console
$ go build -o ./ doctools/cmd/...
$ go install doctools/cmd/...
```
