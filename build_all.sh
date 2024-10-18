#!/bin/bash

# 检查是否传递了目录参数
if [ -z "$1" ]; then
    echo "错误：未指定目录。请提供要编译的目录路径。"
    exit 1
fi

# 获取指定目录
PROTO_DIR=$1

# 检查目录是否存在
if [ ! -d "$PROTO_DIR" ]; then
    echo "错误：目录 $PROTO_DIR 不存在。"
    exit 1
fi

# 获取 GOPATH
GOPATH=$(go env GOPATH)

# 编译所有 .proto 文件
./build.sh "$PROTO_DIR"

# 编译 API
./build_api.sh "$PROTO_DIR"

# 编译 gRPC
./build_grpc.sh "$PROTO_DIR"

# 编译 Swagger
./build_swagger.sh "$PROTO_DIR"

echo "编译完成。"