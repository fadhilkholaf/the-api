all: run

migrate-create:
	@echo ---
	@echo Example: make migrate-create name=initial_migration
	@echo ---
	migrate create -ext sql -dir migrations -seq $(name)

migrate-up:
	migrate --database postgresql://postgres:Q2er5@936@localhost:5432/go-gorm?sslmode=disable -path migrations up

migrate-down:
	migrate --database postgresql://postgres:Q2er5@936@localhost:5432/go-gorm?sslmode=disable -path migrations down

migrate-auto:
	go run ./cmd/server/main.go

run:
	go run ./cmd/server/main.go

build:
	go build -ldflags '-s -w' -o ./bin/ ./cmd/server/main.go

build-windows:
	GOOS=windows GOARCH=amd64 GIN_MODE=release go build -ldflags '-s -w' -o ./bin/main ./cmd/server/main.go

build-linux:
	GOOS=linux GOARCH=amd64 GIN_MODE=release go build -ldflags '-s -w' -o ./bin/main ./cmd/server/main.go