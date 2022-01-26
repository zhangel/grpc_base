#!/bin/bash
#File Name    : protoc.sh
#Author       : zhangerlong
#Mail         : zhangerlong@qianxin.com
#Create Time  : 2022-01-25 08:58
#Description  : 
cd ..
GOOLGE_API=$GOPATH/src/googleapis
API=$GOPATH/src/googleapis
protoc -I$(pwd) -I$API --go_out=plugins=grpc:.  grpc_base/api/hello/message.proto  grpc_base/api/hello.proto
cd grpc_base
