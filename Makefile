gen:
	protoc --go_out=. --go-grpc_out=. proto/*.proto
clean:
	rm pb/*.go
server_start:
	go run server/main.go
client_start:
	go run client/main.go