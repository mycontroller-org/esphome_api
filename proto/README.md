Proto files are copied from the [esphome GitHub repository](https://github.com/esphome/esphome)
release [**2026.7.0**](https://github.com/esphome/esphome/releases/tag/2026.7.0):

* [api.proto](https://github.com/esphome/esphome/blob/2026.7.0/esphome/components/api/api.proto)
* [api_options.proto](https://github.com/esphome/esphome/blob/2026.7.0/esphome/components/api/api_options.proto)

### Steps to update the generated file
* Install [protoc](https://github.com/protocolbuffers/protobuf)
* Install `protoc-gen-go`: `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
* `google/protobuf/descriptor.proto` should be on the import path (set `PROTOC_INCLUDE` if needed)
* Refresh protos from an ESPHome release tag if desired, then:
```bash
ESPHOME_TAG=2026.7.0
curl -sL "https://raw.githubusercontent.com/esphome/esphome/${ESPHOME_TAG}/esphome/components/api/api.proto" \
  -o api.proto
curl -sL "https://raw.githubusercontent.com/esphome/esphome/${ESPHOME_TAG}/esphome/components/api/api_options.proto" \
  -o api_options.proto
./generate.sh
```
* Update `pkg/api/helper.go` type-id maps when message IDs change
* Document changes in `docs/API_PROTO_UPDATE_<from>_to_<to>.md` (ESPHome release versions) and `CHANGELOG.md`
