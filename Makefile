create:
	protoc --proto_path=proto proto/account/*.proto --go_out=./
	protoc --proto_path=proto proto/account/*.proto --go-grpc_out=./

redis:
	docker-compose up

swag:
	swag init -g cmd/main.go