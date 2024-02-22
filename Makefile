include _env

run:
	SERVICE_PORT=${SERVICE_PORT} SERVICE_LOGLEVEL=${SERVICE_LOGLEVEL} go run main.go

test:
	go test ./...

lint:
	golangci-lint run ./...
