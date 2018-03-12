go-graphite/protocol
======

This repository contains protocol definitions for go-graphite services

* zipper-* - protocols used for reading path
  * zipper_v3_grpc - gRPC service definitions for new zipper protocol. Status: **WIP**
  * zipper_v3_pb - new protobuf3-based protocol that's used as a base for gRPC communcation Status: **WIP**
  * zipper_v2_pb - current production protobuf3-based protocol. Status: **STABLE**
  * zipper_v1_pb - protobuf2 protocol, used in carbonzipper 0.6x or earlier. Status: **Deprecated** 