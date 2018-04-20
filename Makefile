BINARY_NAME = docps

build:
	go build -o $(BINARY_NAME) -v

test:
	go test -v ./...

run: 
	go build -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)

linux-build:
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME)-linux -v

darwin-build:
	GOOS=darwin GOARCH=amd64 go build -o $(BINARY_NAME)-darwin -v
	
clean:
	go clean