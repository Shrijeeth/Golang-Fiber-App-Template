setup:
	go install github.com/axw/gocov/gocov@latest
	go install github.com/t-yuki/gocover-cobertura@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.1
	go install github.com/vektra/mockery/v2@v2.20.0

build: setup
	go mod tidy
	go mod vendor
	go build -o ./out/server ./app/internal/server.go

migrate:
	go run ./platform/migrations/run_migrations/migrate.go

run:
	go run ./app/internal/server.go

lint:
	golangci-lint run