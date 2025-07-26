run:
	go run ./cmd/server/main.go

build:
	go build -ldflags '-s -w' -o ./bin/ ./cmd/server/main.go

build-windows:
	GOOS=windows GOARCH=amd64 GIN_MODE=release go build -ldflags '-s -w' -o ./bin/main ./cmd/server/main.go

build-linux:
	GOOS=linux GOARCH=amd64 GIN_MODE=release go build -ldflags '-s -w' -o ./bin/main ./cmd/server/main.go