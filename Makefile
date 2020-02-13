.PHONY: build
build:
		go build -v ./cmd/apiserver
.PHONY: test
test:
		go test -v -race -timeout 30s ./...

.PHONY: migrate-dev
migrate-dev:
		migrate -path migrations -database "postgres://localhost/restapi_dev?sslmode=disable&user=postgres&password=12345" up

.PHONY: migrate-test
migrate-test:
		migrate -path migrations -database "postgres://localhost/restapi_test?sslmode=disable&user=postgres&password=12345" up

.DEFAULT_GOAL := build