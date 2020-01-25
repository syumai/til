gen/proto:
	cd ./proto && protoc --go_out=plugins=grpc:../server/proto/echo ./*.proto

.PHONY: server
server:
	cd ./server && go run cmd/echo/*.go
