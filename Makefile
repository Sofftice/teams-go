run:
	CGO_CXXFLAGS="-std=c++17" go run . -ldflags '-s -w'
	
build-mac:
	go mod tidy

	rm -rf dist/
	mkdir -p dist/

	GOOS="darwin" GOARCH="arm64" CGO_CXXFLAGS="-std=c++17" go build -ldflags '-s -w' -o dist/teams-go
	mkdir -p "dist/Sofftice Teams.app/Contents/MacOS"
	mkdir "dist/Sofftice Teams.app/Contents/Resources"
	
	mv dist/teams-go "dist/Sofftice Teams.app/Contents/MacOS"
	cp building/Info.plist "dist/Sofftice Teams.app/Contents/Info.plist"
	cp building/icon.icns "dist/Sofftice Teams.app/Contents/Resources/icon.icns"