ifeq ($(OS),Windows_NT)
    # Windows
    RM := del /s /q
    MKDIR := mkdir
else
    # Linux/Unix
    RM := rm -rf
    MKDIR := mkdir -p
endif

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

test:
	$(RM) coverage
	$(MKDIR) coverage
	go test -race ./app/src/tests/... -count=1 -p 1 -covermode=atomic -coverprofile=coverage/coverage.out

test.cover: test
	gocov convert coverage/coverage.out | gocov report 2>&1

test.report: test.cover
	go tool cover -html coverage/coverage.out -o coverage/coverage.html
	gocover-cobertura < coverage/coverage.out > coverage/coverage.xml