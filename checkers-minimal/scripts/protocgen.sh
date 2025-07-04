#!/usr/bin/env bash

set -e

echo "Generating gogo proto code"
cd proto
proto_dirs=$(find . -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  for file in $(find "${dir}" -maxdepth 1 -name '*.proto'); do
    # this regex checks if a proto file has its go_package set to github.com/alice/checkers/api/...
    # gogo proto files SHOULD ONLY be generated if this is false
    # we don't want gogo proto to run for proto files which are natively built for google.golang.org/protobuf
    echo "$file"
    if grep -q "option go_package" "$file" && grep -H -o -c 'option go_package.*github.com/alice/checkers/api' "$file" | grep -q ':0$'; then
      buf generate --template buf.gen.gogo.yaml $file
    fi
  done
done

echo "Generating pulsar proto code"
buf generate --template $(pwd)/buf.gen.pulsar.yaml

cd ..

cp -r github.com/alice/checkers/* ./
rm -rf api && mkdir api
mv alice/checkers/* ./api
rm -rf github.com alice
