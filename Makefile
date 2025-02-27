download-deps:
	@go mod download

run: download-deps
	@go run cmd/*

install-mockgen:
	@which mockgen || go install go.uber.org/mock/mockgen@v0.5.0

test:
	@go test -p 1 ./internal/... -race -covermode=atomic -coverprofile=coverage.out -v -tags=skip_coverage
	@go tool cover -html=coverage.out -o coverage.html

mocks: install-mockgen
	mockgen -source=./internal/repository/repository.go -destination=./internal/repository/mocks/repository_mock.go -package=mocks
	mockgen -source=./internal/repository/task.go -destination=./internal/repository/mocks/task_mock.go -package=mocks
	mockgen -source=./internal/clients/clients.go -destination=./internal/clients/mocks/client_mock.go -package=clientmocks