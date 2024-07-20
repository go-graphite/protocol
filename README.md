go-graphite/protocol
====================

This repository contains protocol definitions for go-graphite services

* zipper-* - protocols used for reading path
  * carbonapi_v3_pb - new protobuf3-based protocol that's used as a base for gRPC communcation Status: **STABLE**
  * carbonapi_v2_pb - current production protobuf3-based protocol. Status: **STABLE**

Historical proto releases (deprecated and removed after v1.0.0 release):
  * carbonapi_v3_grpc - gRPC service definitions for new zipper protocol. Was never used in any known system, might need a redisign if carbonapi or something else will implement gRPC support.
  * carbonapi_v1_pb - protobuf2 protocol, used in carbonzipper 0.6x or earlier. Carbonapi no longer supports that protocol. 


Generating Go from *.proto files
--------------------------------

We use [planetscale/vtprotobuf](https://github.com/planetscale/vtprotobuf) generator.

To install required generators please use:
```
$ go install github.com/planetscale/vtprotobuf/cmd/protoc-gen-go-vtproto@v0.6.0
```

Then generate Go files with:
```
$ go generate carbonapi_v2_pb/gen.go
$ go generate carbonapi_v3_pb/gen.go
```
