all: pretty test build run

pretty:
	gofmt -w .
test:
	go test -race -v .
coverage:
	go test -race -v -cover -coverprofile=coverage.out .
cover:
	go tool cover -html=coverage.out
build:
	go build
run:
	go run main.go
