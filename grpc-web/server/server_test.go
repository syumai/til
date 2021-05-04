package server_test

import (
	"context"
	"os"
	"testing"

	"github.com/syumai/til/grpc-web/server"
	"github.com/syumai/til/grpc-web/server/proto/echo"
)

var (
	echoServer *server.EchoServer
)

func TestMain(m *testing.M) {
	echoServer = &server.EchoServer{}
	os.Exit(m.Run())
}

func TestEchoServer_Echo(t *testing.T) {
	tests := []struct {
		name string
		body string
	}{
		{
			name: "Valid",
			body: "example",
		},
		{
			name: "Valid blank",
			body: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			req := &echo.EchoRequest{
				Body: tt.body,
			}
			res, err := echoServer.Echo(ctx, req)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if res.GetBody() != tt.body {
				t.Fatalf("TestEchoServer_Echo() = %s, want %s", res.GetBody(), tt.body)
			}
		})
	}
}

func TestEchoServer_EchoReverse(t *testing.T) {
	tests := []struct {
		name string
		body string
		want string
	}{
		{
			name: "Valid",
			body: "example",
			want: "elpmaxe",
		},
		{
			name: "Valid blank",
			body: "",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			req := &echo.EchoReverseRequest{
				Body: tt.body,
			}
			res, err := echoServer.EchoReverse(ctx, req)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if res.GetBody() != tt.want {
				t.Fatalf("TestEchoServer_EchoReverse() = %s, want %s", res.GetBody(), tt.want)
			}
		})
	}
}
