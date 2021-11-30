VERSION=${VERSION-latest}
API_URL := $(shell terraform -chdir=infra/ output --raw load_balancer_ip)

install:
	go mod download
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.43.0
lint:
	golangci-lint run
test:
	go test ./pkg/
run:
	go run ./pkg/
go: lint test run
tfInit:
	terraform -chdir=infra/ init
tfCheck:
	terraform -chdir=infra/ fmt -check && echo "Terraform check passed" || (echo "Terraform check failed" && exit 1)
tfPlan:
	terraform -chdir=infra/ plan -var app_version="${VERSION}"
tfApply:
	terraform -chdir=infra/ apply -var app_version="${VERSION}"
tfDestroy:
	terraform -chdir=infra/ destroy -var app_version="${VERSION}"
tf: tfInit tfCheck tfPlan
tfOut: 
	@echo "http://${API_URL}/api/encode/string=teste"
	@echo "http://${API_URL}/api/decode/string=dGVzdGU="
testInfra:
	./test_infra.sh ${API_URL}