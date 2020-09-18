#!/usr/bin/env bash
protoc -I=. -I=../../../ --proto_path=${GOPATH}/src --micro_out=. --go_out=. ./*.proto
sed -i.bak 's/,omitempty//g' *.pb.go