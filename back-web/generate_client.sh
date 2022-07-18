#!/bin/bash 

# generate the back end
protoc --proto_path=../protos --go_out=. --go-grpc_out=. ../protos/impactEffect.proto