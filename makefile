VERSION=${VERSION-latest}

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