.PHONY: serve
serve:
	go tool goexec 'http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`.`)))'

.PHONY: build
build:
	GOOS=js GOARCH=wasm go build -o ./main.wasm .
