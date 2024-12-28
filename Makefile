# 用于生成protobuf的Makefile

PROTOBUF_DIR := protobuf

# 定义grpc目录路径
GRPC_DIR := $(PROTOBUF_DIR)/grpc

# 定义buf generate命令
BUF_GENERATE_CMD := buf generate


# 生成go的protobuf代码
gen:
	@echo "Cleaning grpc directory..."
	@rm -rf $(GRPC_DIR)/*
	@echo "Generating Go code..."
	@cd $(PROTOBUF_DIR) && $(BUF_GENERATE_CMD)


# 定义push目标
push:
	@echo "Generating new tag..."
	@TAG_NAME="v$(date +%Y%m%d%H%M%S)" && git tag -a $$TAG_NAME -m "Auto tag for commit $(shell git rev-parse HEAD) on branch $(shell git rev-parse --abbrev-ref HEAD)"
	@echo "Pushing code and tag to remote repository..."
	@git push origin --tags


.PHONY: gen