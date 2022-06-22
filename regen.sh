#!/usr/bin/env bash

if [[ -z ${GOBIN} ]]; then
	echo "GOBIN is not set, please update it if script doesn't work"
	GOBIN="${HOME}/go/bin"
fi
echo "Will use ${GOBIN} as path for go and vt proto generators"

for proto in carbonapi_v2_pb carbonapi_v3_pb; do
	echo "Generating files for ${proto}"
	protoc --go_out=paths=source_relative:. --plugin protoc-gen-go="${GOBIN}/protoc-gen-go" --go-vtproto_out=paths=source_relative:. --plugin protoc-gen-go-vtproto="${GOBIN}/protoc-gen-go-vtproto" --go-vtproto_opt=features=marshal+unmarshal+size+equal+pool "${proto}"/${proto}.proto
done
