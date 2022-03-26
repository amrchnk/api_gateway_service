create:
	protoc --proto_path=proto proto/auth/*.proto --go_out=./
	protoc --proto_path=proto proto/auth/*.proto --go-grpc_out=./