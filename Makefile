run:
	CGO_CXXFLAGS="-std=c++17" go run . -ldflags '-s -w'
	
build:
	go mod tidy
	CGO_CXXFLAGS="-std=c++17" go build -ldflags '-s -w'
