proto_build:
	protoc --go_out=. --go-grpc_out=. proto/greet.proto
