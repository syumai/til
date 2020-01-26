CLIENT_PROTO_OUT_DIR=client/proto/echo
SERVER_PROTO_OUT_DIR=server/proto/echo

gen/proto:
	cd ./proto && protoc \
	  --go_out=plugins=grpc:../$(SERVER_PROTO_OUT_DIR) \
	  --js_out=import_style=commonjs:../$(CLIENT_PROTO_OUT_DIR) \
	  --grpc-web_out=import_style=typescript,mode=grpcwebtext:../$(CLIENT_PROTO_OUT_DIR) \
	  ./*.proto

.PHONY: server
server:
	cd ./server && go run cmd/echo/*.go

.PHONY: client
client:
	echo "unimplemented"

