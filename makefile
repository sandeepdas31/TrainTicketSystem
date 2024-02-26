generateproto:
	protoc --go_out=trainTicketProto --go_opt=paths=source_relative --go-grpc_out=trainTicketProto --go-grpc_opt=paths=source_relative TrainTicket.proto

test:
	go test -v ./...  -coverprofile=coverage.out
	go tool cover -html=coverage  

run:
	go run main.go
