Proto files are copied from the [esphome GitHub repository](https://github.com/esphome/esphome)

* [api.proto](https://github.com/esphome/esphome/blob/dev/esphome/components/api/api.proto)
* [api_options.proto](https://github.com/esphome/esphome/blob/dev/esphome/components/api/api_options.proto)

### Steps to update the generated file
* Install [protoc](https://github.com/protocolbuffers/protobuf)
* `google/protobuf/descriptor.proto` should be on the import path
* Execute the following script
```bash
cd proto
./generate.sh
```