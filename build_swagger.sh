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

# 遍历指定目录下的所有 .proto 文件
for proto_file in $(find "$PROTO_DIR" -name "*.proto"); do
    echo "编译 $proto_file"
    protoc -I . \
        -I ./google/protobuf \
        --swagger_out=logtostderr=true:. \
        "$proto_file"
done

echo "编译完成。"