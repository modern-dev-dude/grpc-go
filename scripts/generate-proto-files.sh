#!/bin/bash
# protoc \
# --go_out=./packages/renderer \
# --go_opt=paths=source_relative \
# --go-grpc_out=./packages/renderer \
# --go-grpc_opt=paths=source_relative \
# ./proto/renderer.proto

generate(){
   echo "1: $1"
   echo "2: $2"

   protoc \
   --go_out=./packages/$2 \
   --go_opt=paths=import \
   --go-grpc_out=./packages/$2 \
   --go-grpc_opt=paths=import \
   ./proto/$1
}

proto_dir="$(pwd)/proto"
if [ ! -d "$proto_dir" ]; then
  echo "not a directory: $proto_dir "
  exit 1
fi

for file in "$proto_dir"/*; do
    strip_path=${file##*/}
    strip_ext=${strip_path%.*}
    generate $strip_path $strip_ext
done

