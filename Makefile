run:
	go run . -ldflags '-s -w'
	
build:
	go mod tidy
	go build -ldflags '-s -w'
