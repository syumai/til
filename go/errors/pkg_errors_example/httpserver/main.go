package main

import (
	"context"
	"fmt"
	"net/http"

	pkg_errors "github.com/pkg/errors"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type loggerKey struct{}

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger, _ := zap.NewProduction()
		ctx := context.WithValue(r.Context(), loggerKey{}, logger)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func loggerFromContext(ctx context.Context) *zap.Logger {
	return ctx.Value(loggerKey{}).(*zap.Logger)
}

func A() error {
	return B()
}

func B() error {
	return pkg_errors.New("error happened")
}

func main() {
	r := chi.NewRouter()
	r.Use(loggerMiddleware)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		logger := loggerFromContext(r.Context())
		err := A()
		if err != nil {
			// {
			//   "level": "error",
			//   "ts": 1675966363.8471649,
			//   "caller": "httpserver/main.go:85",
			//   "msg": "unexpected error",
			//   "error": "error happened",
			//   "errorVerbose": "error happened
			// main.B
			//         /Users/syumai/go/src/github.com/syumai/til/go/errors/pkg_errors_example/httpserver/main.go:33
			// main.A
			//         /Users/syumai/go/src/github.com/syumai/til/go/errors/pkg_errors_example/httpserver/main.go:29
			// main.main.func1
			//         /Users/syumai/go/src/github.com/syumai/til/go/errors/pkg_errors_example/httpserver/main.go:41
			// net/http.HandlerFunc.ServeHTTP
			//         /opt/homebrew/Cellar/go/1.19.5/libexec/src/net/http/server.go:2109
			// github.com/go-chi/chi/v5.(*Mux).routeHTTP
			//         /Users/syumai/go/pkg/mod/github.com/go-chi/chi/v5@v5.0.8/mux.go:444
			// net/http.HandlerFunc.ServeHTTP
			//         /opt/homebrew/Cellar/go/1.19.5/libexec/src/net/http/server.go:2109
			// main.loggerMiddleware.func1
			//         /Users/syumai/go/src/github.com/syumai/til/go/errors/pkg_errors_example/httpserver/main.go:20
			// net/http.HandlerFunc.ServeHTTP
			//         /opt/homebrew/Cellar/go/1.19.5/libexec/src/net/http/server.go:2109
			// github.com/go-chi/chi/v5.(*Mux).ServeHTTP
			//         /Users/syumai/go/pkg/mod/github.com/go-chi/chi/v5@v5.0.8/mux.go:90
			// net/http.serverHandler.ServeHTTP
			//         /opt/homebrew/Cellar/go/1.19.5/libexec/src/net/http/server.go:2947
			// net/http.(*conn).serve
			//         /opt/homebrew/Cellar/go/1.19.5/libexec/src/net/http/server.go:1991
			// runtime.goexit
			//         /opt/homebrew/Cellar/go/1.19.5/libexec/src/runtime/asm_arm64.s:1172",
			//   "stacktrace": "main.main.func1
			//         /Users/syumai/go/src/github.com/syumai/til/go/errors/pkg_errors_example/httpserver/main.go:85
			// net/http.HandlerFunc.ServeHTTP
			//         /opt/homebrew/Cellar/go/1.19.5/libexec/src/net/http/server.go:2109
			// github.com/go-chi/chi/v5.(*Mux).routeHTTP
			//         /Users/syumai/go/pkg/mod/github.com/go-chi/chi/v5@v5.0.8/mux.go:444
			// net/http.HandlerFunc.ServeHTTP
			//         /opt/homebrew/Cellar/go/1.19.5/libexec/src/net/http/server.go:2109
			// main.loggerMiddleware.func1
			//         /Users/syumai/go/src/github.com/syumai/til/go/errors/pkg_errors_example/httpserver/main.go:20
			// net/http.HandlerFunc.ServeHTTP
			//         /opt/homebrew/Cellar/go/1.19.5/libexec/src/net/http/server.go:2109
			// github.com/go-chi/chi/v5.(*Mux).ServeHTTP
			//         /Users/syumai/go/pkg/mod/github.com/go-chi/chi/v5@v5.0.8/mux.go:90
			// net/http.serverHandler.ServeHTTP
			//         /opt/homebrew/Cellar/go/1.19.5/libexec/src/net/http/server.go:2947
			// net/http.(*conn).serve
			//         /opt/homebrew/Cellar/go/1.19.5/libexec/src/net/http/server.go:1991"
			// }
			logger.Error("unexpected error", zap.Error(err))
		}
	})
	const addr = ":9830"
	fmt.Printf("listening server on: localhost%s\n", addr)
	http.ListenAndServe(addr, r)
}
