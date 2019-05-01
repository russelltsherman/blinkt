# blinkt

This is a re-implementation of [Pimoroni's Python library](https://github.com/pimoroni/blinkt), written in Go.

The [Pimoroni Blinkt](https://shop.pimoroni.com/products/blinkt) is a low-profile strip of eight super-bright, color LED indicators that plugs directly onto the Raspberry Pi's GPIO header.

Examples are included via a cli app, just pass `--help` flag to see the available options.

## load dependencies

```sh
GO111MODULE=on go mod download
```

## run

```sh
go run app.go --help
```
