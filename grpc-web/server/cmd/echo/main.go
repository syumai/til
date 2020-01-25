package main

import (
	"log"
	"net"
	"os"

	"github.com/syumai/til/grpc-web/server"
	"github.com/syumai/til/grpc-web/server/proto/echo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const addr = ":5000"

func main() {
	os.Exit(realMain())
}

func realMain() int {
	s := grpc.NewServer()
	echo.RegisterEchoServer(s, &server.EchoServer{})
	reflection.Register(s)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Println(err)
		return 1
	}
	defer lis.Close()

	if err = s.Serve(lis); err != nil {
		log.Println(err)
		return 1
	}
	return 0
}
