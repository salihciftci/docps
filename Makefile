BINARY_NAME = docps

build:
	go build -o $(BINARY_NAME) -v
test:
	go test -v ./...
run: 
	go build -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
clean:
	go clean