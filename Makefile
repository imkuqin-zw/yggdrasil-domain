# New Makefile for multi-architecture
.PHONY: all

gen-example-proto:
	@bash ./scripts/gen_example_proto.sh

build:
	@cd cmd/protoc-gen-yggdrasil-domain && go build

install:
	@cd cmd/protoc-gen-yggdrasil-domain && go install