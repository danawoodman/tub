# BestWay hottub CLI

![tub](gopher-tub.png)

> A Golang powered CLI for controlling and monitoring BestWay hottubs

## Installation

```sh
go install github.com/danawoodman/tub
```

This provides a `tub` command that you can use to control your BestWay hottub (see usage below).

## Usage

```sh
# view the docs
tub --help

# first, login using your BestWay credentials
# these credentials are stored in your ~/.bestway.yaml file
# and are used to generate API tokens to authenticate with the API
# service that controls your BestWay tub
tub login -u <USERNAME> -p <PASSWORD>

# check your login status
tub whoami

# get the status of your tub (temp, power, etc.)
tub status

# set the temperature of your tub
tub temp 104

# other commands:
tub filter off|on
tub jets off|low|high
tub heat off|on
tub lock off|on
tub power off|on
```

## Development

```sh
# using makefile, install the tub CLI to your $GOPATH
make install
```

## Troubleshooting

### 502 Bad Gateway

The API service (Gizwits) seems to be quite unstable, so sometimes you'll see a `502 Bad Gateway` error. If you do, just try running the command again, usually it will work.
