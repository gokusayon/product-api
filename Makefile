.PHONY: run

check_install:
	which swagger || GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger: check_install
	swagger generate spec -o ./swagger.yaml --scan-models

run:
	go run main.go

build:
	go build main.go
