export GO111MODULE=on

install:                  ## Install this program.
	go install -v ./...

exporters:                ## Install exporters.
	go install -v ./vendor/github.com/percona/mysqld_exporter

serve: install exporters  ## Start program as server and listen for incoming http requests.
	pmm-agent --config=testdata/.pmm-agent.yml serve

test: exporters           ## Run tests.
	go test -mod=vendor -v -race ./...

test-cover: exporters
	go install -v ./vendor/github.com/AlekSi/gocoverutil
	gocoverutil test -v -race ./...

send-cover: SHELL:=/bin/bash
send-cover:
	bash <(curl -s https://codecov.io/bash) -X fix

gen:                      ## Run `go generate`.
	go generate ./...

api:                      ## Generate api.
	go install -v ./vendor/github.com/golang/protobuf/protoc-gen-go
	protoc -Iapi api/*.proto --go_out=plugins=grpc:api

lint:	                  ## Run `golangci-lint`.
	golangci-lint run

format:	                  ## Run `goimports`.
	go install -v ./vendor/golang.org/x/tools/cmd/goimports
	goimports -local github.com/percona/pmm-agent -l -w $(shell find . -type f -name '*.go' -not -path "./vendor/*")

demo: install exporters   ## Demo.
	cp testdata/.pmm-agent.yml ~/.pmm-agent.yml
	pmm-agent list
	pmm-agent add mysql-1 --env DATA_SOURCE_NAME=root@/ -- mysqld_exporter --collect.all
	pmm-agent list
	pmm-agent stop mysql-1
	pmm-agent list
	pmm-agent start mysql-1
	pmm-agent list
	pmm-agent add mysql-2 --env DATA_SOURCE_NAME=root@/ -- mysqld_exporter --collect.all --web.listen-address=:9204
	pmm-agent list
	pmm-agent stop mysql-1 mysql-2
	pmm-agent list
	pmm-agent start mysql-1 mysql-2
	pmm-agent list
	pmm-agent stop
	pmm-agent list
	pmm-agent start
	pmm-agent list
	pmm-agent remove mysql-1
	pmm-agent list
	pmm-agent remove
	pmm-agent list

verify:                   ## Ensure that vendor/ is in sync with `go.*`.
	go mod verify
	go mod vendor
	git diff --exit-code

help: Makefile            ## Display this help message.
	@echo "Please use \`make <target>\` where <target> is one of:"
	@grep '^[a-zA-Z]' $(MAKEFILE_LIST) | \
	    sort | \
	    awk -F ':.*?## ' 'NF==2 {printf "  %-26s%s\n", $$1, $$2}'

.DEFAULT_GOAL := help
.PHONY: install exporters serve test test-cover send-cover gen api lint format demo verify help
