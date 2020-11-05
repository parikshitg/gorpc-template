grpc:
	protoc --go_out=plugins=grpc:. ./proto/user.proto
build:
	go build
run:
	go run main.go
docker:
	docker build . 