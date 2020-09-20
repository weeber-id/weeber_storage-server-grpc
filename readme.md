# Weeber gRPC Server for Internal Storage Engine

## Protocol Buffer Generator

using `libprotoc 3.12.1`

- Go:
  >protoc --go_out=plugins=grpc:./ --go_opt=paths=source_relative storage/storage.proto`