install:
	go mod download
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.43.0
lint:
	golangci-lint run
test:
	go test ./pkg/
run:
	go run ./pkg/