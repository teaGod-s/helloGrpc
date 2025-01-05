.PHONY: init
# 安装buf
init:
	go install github.com/bufbuild/buf/cmd/buf@v1.48.0

.PHONY: break
# proto破坏性变更检查
break:
	buf breaking --against ".git#subdir=./proto"

.PHONY: gen
# pb生成 swagger生成
gen:
	buf dep update && buf build && buf generate && cd docs && go generate && go fmt

.PHONY: build
# 构建二进制
build:
	go mod tidy
	go build -o example-gateway-server cmd/example-gateway-server/main.go
	go build -o example-grpc-server cmd/example-grpc-server/main.go

.PHONY: all
# 依次执行上面所有命令
all:
	make init;
	make break;
	make gen;
	make build;

.PHONY: help
# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)
.DEFAULT_GOAL := help
