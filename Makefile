APP_NAME=packapp

.PHONY: dev build-linux build-mac-intel build-mac-arm build-windows test

dev:
	docker compose up

test:
	go test -v ./...

build-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/$(APP_NAME)-linux ./cmd/app/main.go

# Build for macOS (Intel)
build-mac-intel:
	GOOS=darwin GOARCH=amd64 go build -o bin/$(APP_NAME)-mac ./cmd/app/main.go

# Build for macOS (Apple Silicon)
build-mac-arm:
	GOOS=darwin GOARCH=arm64 go build -o bin/$(APP_NAME)-mac-arm ./cmd/app/main.go

build-windows:
	GOOS=windows GOARCH=amd64 go build -o bin/$(APP_NAME).exe ./cmd/app/main.go