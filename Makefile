BINARY_NAME = liman

build:
	go build -o $(BINARY_NAME) -v

run: 
	go build -o $(BINARY_NAME) -v 
	./$(BINARY_NAME)

test:
	go test -v .

linux-build:
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME) -v

darwin-build:
	GOOS=darwin GOARCH=amd64 go build -o $(BINARY_NAME) -v

docker:
	docker build -t $(BINARY_NAME) .

clean:
	go clean

.PHONY: build run test linux-build darwin-build docker clean
