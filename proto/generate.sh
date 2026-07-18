#!/bin/bash
set -euo pipefail

# Prerequisites:
#   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
#   protoc (https://github.com/protocolbuffers/protobuf/releases)
#   google/protobuf/descriptor.proto must be on the import path
#   (bundled with protoc under include/)

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
OUT_DIR="${SCRIPT_DIR}/../pkg/api"
MODULE="github.com/mycontroller-org/esphome_api/pkg/api"
mkdir -p "${OUT_DIR}"

# Prefer PROTOC_INCLUDE if set; otherwise try common locations.
INCLUDE_FLAGS=()
if [[ -n "${PROTOC_INCLUDE:-}" ]]; then
  INCLUDE_FLAGS+=(-I"${PROTOC_INCLUDE}")
elif [[ -d /tmp/protoc-install/include ]]; then
  INCLUDE_FLAGS+=(-I/tmp/protoc-install/include)
elif [[ -d /usr/include ]]; then
  INCLUDE_FLAGS+=(-I/usr/include)
fi

cd "${SCRIPT_DIR}"
protoc \
  -I. \
  "${INCLUDE_FLAGS[@]}" \
  --go_out="${OUT_DIR}" --go_opt=paths=source_relative \
  --go_opt=Mapi.proto="${MODULE}" \
  --go_opt=Mapi_options.proto="${MODULE}" \
  api.proto api_options.proto

echo "Generated Go bindings in ${OUT_DIR}"
echo "Note: after regenerating, update pkg/api/helper.go type-id maps if message IDs changed."
