package server

import (
	"context"
	"log"

	"github.com/syumai/til/grpc-web/server/proto/echo"
	"github.com/syumai/til/grpc-web/server/strutil"
)

var (
	_ echo.EchoServer = &EchoServer{}
)

// EchoServer is server implementation of echo.
type EchoServer struct{}

// Echo returns given string as response.
func (s *EchoServer) Echo(ctx context.Context, req *echo.EchoRequest) (*echo.EchoResponse, error) {
	body := req.GetBody()
	log.Println("body:", body)
	return &echo.EchoResponse{Body: body}, nil
}

// EchoReverse reverses given string and returns as response.
func (s *EchoServer) EchoReverse(ctx context.Context, req *echo.EchoReverseRequest) (*echo.EchoReverseResponse, error) {
	body := req.GetBody()
	reversedBody := strutil.Reverse(body)
	log.Println("reversed body:", reversedBody)
	return &echo.EchoReverseResponse{Body: reversedBody}, nil
}
