#!/bin/bash


generateProtoFiles(){
   protoc \
   --go_out=./packages \
   --go_opt=paths=import \
   --go-grpc_out=./packages \
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
    generateProtoFiles $strip_path
done

