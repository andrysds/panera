dep:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure -v

pretty:
	gofmt -w .

compile:
	go build -o panera main.go
	go build -o panera-scheduler scheduler/main.go

run: pretty compile
	./panera
