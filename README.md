go-graphite/protocol
====================

This repository contains protocol definitions for go-graphite services

* zipper-* - protocols used for reading path
  * carbonapi_v3_grpc - gRPC service definitions for new zipper protocol. Status: **WIP**
  * carbonapi_v3_pb - new protobuf3-based protocol that's used as a base for gRPC communcation Status: **STABLE**
  * carbonapi_v2_pb - current production protobuf3-based protocol. Status: **STABLE**
  * carbonapi_v1_pb - protobuf2 protocol, used in carbonzipper 0.6x or earlier. Status: **Deprecated** 


Generating Go from *.proto files
--------------------------------

We use [gogo/protobuf](https://github.com/gogo/protobuf). Install the binaries
needed with:
```
$ go get github.com/gogo/protobuf/proto
$ go get github.com/gogo/protobuf/protoc-gen-gogofast 
$ go get github.com/gogo/protobuf/protoc-gen-gogofaster
$ go get github.com/gogo/protobuf/protoc-gen-gogoslick 
$ go get github.com/gogo/protobuf/gogoproto
```
Then generate Go files with:
```
$ go generate ./...
```
