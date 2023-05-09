#!/usr/bin/env bash

# shellcheck disable=SC2128
CURRENT_PATH=$(dirname "$BASH_SOURCE")
ROOT_PATH=$(dirname "$CURRENT_PATH")
EXAMPLE_PROTO_PATH="$ROOT_PATH/example/proto"
THIRD_PROTO_PATH="$ROOT_PATH/proto"

protoc -I="$ROOT_PATH" -I="$THIRD_PROTO_PATH" \
  --go_out="." --go_opt=paths=source_relative \
  --yggdrasil-domain_out="." --yggdrasil-domain_opt=paths=source_relative \
  "$EXAMPLE_PROTO_PATH"/*.proto