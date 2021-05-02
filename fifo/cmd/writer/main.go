package main

import (
	"context"
	"errors"
	"io"
	"os"
	"os/signal"
	"syscall"
)

const fileName = "fifo"

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

	err := syscall.Mkfifo(fileName, 0_600)
	if err != nil {
		panic(err)
	}
	f, err := os.OpenFile(fileName, syscall.O_WRONLY, os.ModeNamedPipe)
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
