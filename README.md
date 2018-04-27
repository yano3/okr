# okr

[![GitHub release](https://img.shields.io/github/release/yano3/okr.svg)](https://github.com/yano3/okr/releases)
[![CircleCI](https://circleci.com/gh/yano3/okr.svg?style=shield)](https://circleci.com/gh/yano3/okr)

Build okara URL with cli.

## Usage

```console
$ okr -c 300x200c -f webp https://example.com/path/to/your/image.jpg
```

## Installation

### macOS

If you use [Homebrew](https://brew.sh):

```
brew tap yano3/okr
brew install okr
```

### Other platforms

Download binary from [releases page](https://github.com/yano3/okr/releases) or use `go get` command.

```console
$ go get -u github.com/yano3/okr
```

## Configuration

Set environment variables bellow.

```
export OKARA_HOST=<Put your okara host beginning with "http" or "https">
export OKARA_SERVICE=<Put your service>
export OKARA_TYPE=<Put your type>
export OKARA_SECRET_TOKEN=<Put your okara secret token>
```
