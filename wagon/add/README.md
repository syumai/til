# Add example

* This is a example package to execute WebAssembly in Go package.
* This command adds two numbers using Wasm loaded in Go package.

## Install

```
go get github.com/syumai/til/wagon/add/cmd/add
```

## Usage

```
$ add 1 3
4
```

## Development

### Editing Wasm file

1. Edit Wat in `module` directory
2. convert Wat into Wasm using [`wat2wasm`](https://github.com/WebAssembly/wabt)
3. bundle Wasm into `statik` package by running `go generate`

### Test

```
go test ./...
```
