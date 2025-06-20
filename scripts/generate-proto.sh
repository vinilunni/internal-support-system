#!/bin/bash

set -e

PROTO_DIR="proto"
OUT_DIR="."

echo "Generating protobuf files..."

protoc --proto_path=$PROTO_DIR \
       --go_out=$OUT_DIR \
       --go_opt=paths=source_relative \
       --go-grpc_out=$OUT_DIR \
       --go-grpc_opt=paths=source_relative \
       $PROTO_DIR/auth/auth.proto \
       $PROTO_DIR/user/user.proto \
       $PROTO_DIR/asset/asset.proto \
       $PROTO_DIR/assignment/assignment.proto \
       $PROTO_DIR/ticket/ticket.proto \
       $PROTO_DIR/notification/notification.proto

echo "Protobuf generation completed!" 