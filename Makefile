GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
API_PROTO_FILES=$(shell find api -name *.proto)

CODE_DIC = /root/go/src/fxkt.tech/bj21
INSTALL_DIC = /webroot/joker/v1
LOG_DIC = /data/weblog/joker/v1
BIN_FILE = joker_lms_v1

.PHONY: init
# init env
init:
	go get -u github.com/go-kratos/kratos/cmd/kratos/v2
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get -u github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2
	go get -u github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2
	go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	go get -u github.com/envoyproxy/protoc-gen-validate


.PHONY: proto
# proto => pb.go
proto:
	protoc --proto_path=. \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:. \
 	       --go-http_out=paths=source_relative:. \
 	       --go-grpc_out=paths=source_relative:. \
               --validate_out=paths=source_relative,lang=go:. \
               --openapiv2_out . \
	       $(API_PROTO_FILES)
	protoc --proto_path=. \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:. \
	       $(INTERNAL_PROTO_FILES)
	protoc --proto_path=. \
               --proto_path=./third_party \
               --go_out=paths=source_relative:. \
               --go-errors_out=paths=source_relative:. \
               $(API_PROTO_FILES)


.PHONY: ps
# list process
ps:
	ps aux | grep $(BIN_FILE) | grep -v grep


.PHONY: install
# install
install:
	mkdir -p $(INSTALL_DIC)
	mkdir -p $(LOG_DIC)
	cp -r conf $(INSTALL_DIC)
	make build;
	mv bin/lms bin/$(BIN_FILE)
	mv bin $(INSTALL_DIC)/bin
	cp Makefile $(INSTALL_DIC)


.PHONY: build
# build at code directory
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...


.PHONY: rebuild
# build at install directory
rebuild:
	cd $(CODE_DIC) && git pull && make build;
	rm -rf $(INSTALL_DIC)/bin
	mv $(CODE_DIC)/bin $(INSTALL_DIC)/bin
	mv bin/lms bin/$(BIN_FILE)

.PHONY: serve
# http service
serve:
	bin/$(BIN_FILE) -cmd serve -conf conf/ol.yaml > $(LOG_DIC)/lms_serve.log 2>&1 &

.PHONY: stat_status
stat_status:
	bin/$(BIN_FILE) -cmd stat_status -conf conf/ol.yaml > /dev/null 2>&1 &

.PHONY: stat_speed
stat_speed:
	bin/$(BIN_FILE) -cmd stat_speed -conf conf/ol.yaml > /dev/null 2>&1 &

# show help
help:
	@echo 'LocalMediaService(code:joker)'
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help