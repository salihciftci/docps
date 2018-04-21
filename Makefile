BINARY_NAME = liman

run: 
	go build -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)

build:
	go build -o $(BINARY_NAME) -v

test:
	go test -v ./...

linux-build:
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME)-linux -v

darwin-build:
	GOOS=darwin GOARCH=amd64 go build -o $(BINARY_NAME)-darwin -v

docker:
	docker build -t $(BINARY_NAME) .
	
clean:
	go clean