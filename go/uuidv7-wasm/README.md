# uuidv7-wasm

A minimal sample that generates and displays a UUID v7 on button click in the
browser, using the `uuid` package newly added to the standard library in Go 1.27
RC1 (`uuid.NewV7()`), compiled to Wasm with `GOOS=js GOARCH=wasm`.

## Files

- `main.go` — Registers a global function `generateUUIDv7` via `syscall/js` that
  returns `uuid.NewV7().String()` (built with `GOOS=js GOARCH=wasm`).
- `index.html` — Loads `wasm_exec.js`, starts `main.wasm`, and on button click
  calls `generateUUIDv7()` and renders the result into `<pre>`.
- `wasm_exec.js` — Copied as-is from Go's `lib/wasm/wasm_exec.js`.
- `serve/main.go` — A static file server using `net/http`'s `FileServer`. Go's
  `mime` package serves `.wasm` as `application/wasm`, so
  `WebAssembly.instantiateStreaming` works.
- `Makefile` — `build` / `serve` targets.

## Usage

```sh
make build   # produce main.wasm
make serve   # serve at http://localhost:8080/
```

Open http://localhost:8080/ and click the "Generate UUID v7" button to display a
UUID v7. The value changes on every click.

## Notes

`go.mod` declares `go 1.27rc1`, so GOTOOLCHAIN automatically uses the go1.27rc1
toolchain (the standard library `uuid` package was introduced in this version).
