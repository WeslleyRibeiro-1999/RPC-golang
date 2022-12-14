server: 
	go run cmd/server/main.go

client:
	go run cmd/client/main.go

proto:
	protoc --go_out=. --go-grpc_out=. proto/*.proto