build:
	go build

run:
	APP_ENV=local go run main.go

lint:
	golangci-lint run -c .dev/.golangci.yml

lint-fix:
	golangci-lint run --fix

unit-test:
	go test ./...

code-coverage:
	go test ./... -coverprofile cover.out
	go tool cover -html=cover.out -o coverage.html
	echo `go tool cover -func cover.out | grep total`

all-tests: lint unit-test

generate-mocks:
	mockgen -destination=internal/mocks/mock_ethereum_service.go -source=internal/ethereum/handler.go -package=mocks
	mockgen -destination=internal/mocks/mock_ethereum_client.go -source=internal/ethereum/service.go -package=mocks


