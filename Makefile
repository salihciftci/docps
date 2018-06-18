BINARY_NAME = liman

build:
	go build -o $(BINARY_NAME) -v

run: 
	go build -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)

test:
	go test -v ./...

linux-build:
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME)-linux -v

darwin-build:
	GOOS=darwin GOARCH=amd64 go build -o $(BINARY_NAME)-darwin -v

windows-build:
	GOOS=windows GOARCH=amd64 go build -o $(BINARY_NAME)-windows.exe -v
	
docker:
	docker build -t $(BINARY_NAME) .
	
clean:
	go clean
