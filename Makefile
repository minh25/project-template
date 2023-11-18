PROTOC_LINUX_VERSION = 3.11.4
PROTOC_LINUX_ZIP = protoc-$(PROTOC_LINUX_VERSION)-linux-x86_64.zip

.PHONY: install-go-tools install-protoc gen-proto migrate-up migrate-down-1 run lint test get-next-err-code check-sql

install-go-tools:
	go install github.com/matryer/moq
	go install github.com/mgechev/revive

install-protoc:
	curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v$(PROTOC_LINUX_VERSION)/$(PROTOC_LINUX_ZIP)
	sudo unzip -o $(PROTOC_LINUX_ZIP) -d /usr/local bin/protoc
	sudo unzip -o $(PROTOC_LINUX_ZIP) -d /usr/local 'include/*'
	rm -f $(PROTOC_LINUX_ZIP)
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

gen-proto:
	rm -r pb
	mkdir -p pb
	sh gen-proto

migrate:
	echo \# make migrate name="$(name)"
	go run cmd/migrate/main.go create $(name)

migrate-up:
	go run cmd/migrate/main.go up

migrate-down-1:
	go run cmd/migrate/main.go down 1

run:
	go run cmd/server/main.go start

email-noti:
	go run cmd/cronjob/main.go email-notify

consume:
	go run cmd/server/main.go start-email-consumer

lint:
	$(foreach f,$(shell go fmt ./...),@echo "Forgot to format file: ${f}"; exit 1;)
	go vet ./...
	revive -config revive.toml -formatter friendly ./...

test:
	go test -p 1 -v ./...

# rpc-status more details: https://grpc.github.io/grpc/core/md_doc_statuscodes.html
# ex: get-next-err-code rpc-status=13
get-next-err-code:
	go run cmd/server/main.go next-error-code $(rpc-status)

check-sql:
	go run cmd/server/main.go check-sql

gen-mocks:
	find . -name \*_mocks.go -type f -delete
	go generate ./...

install-wkhtmltopdf:
	sudo apt install wkhtmltopdf

jira-test:
	rm -f testrun.*
	REPORTED_ISSUE_LINKS=${REPORTED_ISSUE_LINKS} JIRA_PWD=${PWD} \
		go test -count=1 -p 1 -covermode=count \
		./... && (cat testrun.tmp.json | jq -s "." > testrun.json)
