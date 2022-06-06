
APP_RELATIVE_PATH=$(shell a=`basename $$PWD` && echo $$a)
API_PROTO_FILES=$(shell cd api && find . -name *.proto)
PB_FILES=$(shell cd api && find . -name *.pb.go)

.PHONY: init
init:
	# 自动生成proto 命令包
	go install github.com/china-xs/gin-tpl/cmd/proto@latest
	# 修改proto json tog 包
	go install github.com/favadi/protoc-go-inject-tag
	# swagger 依赖包
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	# proto&validate 请求参数校验
	go install github.com/envoyproxy/protoc-gen-validate@latest
	# 错误状态返回提示语，使用kratos error
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2
	# proto 生成gin 代码包
	go install github.com/china-xs/gin-tpl/cmd/protoc-gen-go-gin@latest
	# 安装 grpc 模块代码包
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	# wire 安装
	go install github.com/google/wire/cmd/wire

.PHONY: wire
# generate wire
wire:
	cd cmd/$(APP_RELATIVE_PATH) && wire

.PHONY: gorm
gorm:
	cd cmd/gorm && go run .

.PHONY: http
# generate http、swagger、grpc、gin、validate error
http:
	cd api && protoc --proto_path=. \
           --proto_path=../third_party \
           --go_out=paths=source_relative:. \
           --go-grpc_out=paths=source_relative:. \
           --go-gin_out=paths=source_relative:. \
           --validate_out=paths=source_relative,lang=go:. \
           --go-errors_out=paths=source_relative:. \
           --openapiv2_out . \
           --openapiv2_opt logtostderr=true \
           $(API_PROTO_FILES)


.PHONY: build
build:
	@go env -w GOPROXY=https://goproxy.cn
	@go mod tidy
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/${APP_RELATIVE_PATH}/


.PHONY: tag
tag:
	cd api && \
	 for name in $(PB_FILES); \
		do \
		protoc-go-inject-tag -input=$$name; \
		done

.PHONY: errors
errors:
	cd api && protoc --proto_path=. \
					  --proto_path=../third_party \
					  --go_out=paths=source_relative:. \
					  --go-errors_out=paths=source_relative:. \
					  ./errs/errors.proto

.PHONY: gen
gen:
	make http
	#make errors
	make tag