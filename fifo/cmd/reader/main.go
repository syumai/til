package main

import (
	"context"
	"io"
	"os"
	"os/signal"
	"syscall"
)

const fileName = "fifo"

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

	f, err := os.OpenFile(fileName, syscall.O_RDONLY, os.ModeNamedPipe)
	if err != nil {
		panic(err)
	}

	go func() {
		<-ctx.Done()
		_ = f.Close()
	}()

	_, err = io.Copy(os.Stdout, f)
	if err != nil {
		panic(err)
	}
}
