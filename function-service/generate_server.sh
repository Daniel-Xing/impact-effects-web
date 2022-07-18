#!/bin/bash 

# generate the front end
python -m grpc_tools.protoc -I../protos --python_out=. --grpc_python_out=. ../protos/impactEffect.proto