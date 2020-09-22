# Weeber gRPC Server for Internal Storage Engine

## Protocol Buffer Generator

using `libprotoc 3.12.1`

- Go:
  - Public storage
    >protoc --go_out=plugins=grpc:./ --go_opt=paths=source_relative protobuf/v1/PublicStorage/public_storage.proto

  - Private storage
    >protoc --go_out=plugins=grpc:./ --go_opt=paths=source_relative protobuf/v1/PrivateStorage/private_storage.proto