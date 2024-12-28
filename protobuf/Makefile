# 获取当前提交中被修改的文件列表，并过滤掉非 protocol/proto 路径的文件
modified_files := $(shell git diff --name-only | grep '^protocol/proto/')

new_files := $(shell git ls-files --others --exclude-standard | grep '^proto/')

all_files := $(modified_files) $(new_files)

# 将剩下的文件路径进行分组并去掉路径的最后一个斜杠
grouped_dirs := $(sort $(patsubst %/,%,$(dir $(all_files))))

# 过滤掉重复的目录路径
unique_dirs := $(foreach dir,$(grouped_dirs),$(if $(findstring $(dir),$(DIRS)),,$(eval DIRS+=$(dir))$(dir)))


.PHONY: default

go:
	$(shell buf generate)

doc:
	$(foreach dir,$(unique_dirs), \
		$(call gen_doc, "$(dir)") \
	)

clean:
	$(foreach dir,$(unique_dirs), \
		$(call clean_go, "$(dir)") \
	)

tools:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest


get_last_segment = $(lastword $(subst /, ,$(patsubst "%",%,$(1))))

# 废弃 使用 buf generate 生成 pb 文件
define compile_go 

	$(eval package_name := $(call get_last_segment,$(1)))
	$(eval out_package_name := $(strip $(subst -,, $(package_name))))
	$(eval proto_file := $(wildcard proto/$(package_name)/*.proto))
	$(eval go_out := ${PWD}/grpc/$(out_package_name))
	$(eval ARCH := "`uname -s`")

	@echo "[INFO] input  package name: $(package_name)"
	@echo "[INFO] output package name: $(out_package_name)"

	@rm -rf $(go_out)
	@mkdir -p $(go_out)

	@echo "[INFO] go out: $(go_out)"
	@echo "[INFO] compiling"
	@protoc \
	-I proto \
	-I vendor \
	--go_out=$(go_out) \
	--go-grpc_opt=require_unimplemented_servers=false \
	--go-grpc_out=$(go_out) \
	--grpc-gateway_out=$(go_out) \
	--grpc-gateway_opt=logtostderr=true \
	--grpc-gateway_opt=paths=source_relative \
	$(proto_file)

	@echo "[INFO] move or delete useless directory..."
	@bash script/mvdefine.sh $(package_name) $(go_out)
	@echo "[INFO] compilation done"
	@echo "================================================"
endef

define clean_go
	$(eval package_name := $(call get_last_segment,$(1)))
	$(eval out_package_name := $(strip $(subst -,, $(package_name))))
	$(eval go_out := ${PWD}/grpc/$(out_package_name))
	@rm -rf $(go_out) 
	@echo "[INFO] clean done"
endef

define gen_doc
	$(eval package_name := $(call get_last_segment,$(1)))
	$(eval out_package_name := $(strip $(subst -,, $(package_name))))
	@echo "[INFO] input  package name: $(package_name)"
	@echo "[INFO] output package name: $(out_package_name)"
	@protoc -I proto \
	 -I vendor \
     --openapiv2_out=proto/$(package_name) \
     --openapiv2_opt logtostderr=true \
	 --openapiv2_opt json_names_for_fields=false  \
	$(proto_file)
	@bash script/openapiv2_doc.sh $(package_name) $(package_name)
	@echo "[INFO] doc done"
endef 
