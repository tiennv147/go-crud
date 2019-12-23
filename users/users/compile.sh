#!/usr/bin/env sh

# Generate Go code for this proto file

protoc -I users/ users/users.proto --go_out=plugins=grpc:users
