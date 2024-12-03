package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func runServer(ctx context.Context, port string) error {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	srv := &http.Server{Addr: fmt.Sprintf(":%s", port), Handler: h}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("error on serve: %v", err)
		}
	}()

	// 外から渡されたContextの完了を待ってサーバーをShutdownする
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return srv.Shutdown(ctx)
}

func main() {
	if err := runServer(context.Background(), "8080"); err != nil {
		log.Fatal(err)
	}
}
