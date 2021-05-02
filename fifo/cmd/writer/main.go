package main

import (
	"context"
	"errors"
	"io"
	"os"
	"os/signal"
	"syscall"

	"github.com/containerd/fifo"
)

const fileName = "fifo"

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

	f, err := fifo.OpenFifo(ctx, fileName, syscall.O_CREAT|syscall.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}

	go func() {
		<-ctx.Done()
		_ = f.Close()
	}()

	_, err = io.Copy(f, os.Stdin)
	if err != nil && !errors.Is(err, os.ErrClosed) {
		panic(err)
	}

	err = os.Remove(fileName)
	if err != nil {
		panic(err)
	}
}
