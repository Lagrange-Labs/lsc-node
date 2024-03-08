set -e

protoVer=1.12.0
protoImageName=bufbuild/buf:${protoVer}

echo "Generating gogo proto code"
cd proto
proto_dirs=$(find . -name '*.proto' -exec dirname {} \; | sort | uniq)
for dir in $proto_dirs; do
  for file in $(find "${dir}" -maxdepth 1 -name '*.proto'); do
    sudo rm -rf gen
    docker run --rm -v $PWD:/workspace --workdir /workspace ${protoImageName} generate --template buf.gen.yaml $file
    sudo chmod -R 775 gen
    cp -r gen/go/${dir}/* ../$(echo "${dir}" | sed 's/\(.*\)\/v\([2-9]\)/\1\/types\/v\2/; s/\(.*\)\/v1/\1\/types/')
    sudo rm -rf gen
  done
done
