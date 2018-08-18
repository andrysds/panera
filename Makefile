dep:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure -v

pretty:
	gofmt -w .

compile:
	go build -o panera .

run: pretty compile
	./panera
