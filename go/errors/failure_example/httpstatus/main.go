package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/morikuni/failure"
	"github.com/syumai/til/go/errors/failure_example/codes"
)

func A() error {
	return failure.New(codes.InvalidArgument, failure.Message("invalid argument error"))
}

func handleError(err error, w http.ResponseWriter) {
	switch {
	case failure.Is(err, codes.InvalidArgument):
		w.WriteHeader(http.StatusBadRequest)
	case failure.Is(err, codes.NotFound):
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("internal error: %+v", err)
	}
}

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		err := A()
		if err != nil {
			// $ curl -v localhost:9840
			// < HTTP/1.1 400 Bad Request
			handleError(err, w)
			return
		}
	})
	const addr = ":9840"
	fmt.Printf("listening server on: localhost%s\n", addr)
	http.ListenAndServe(addr, r)
}
