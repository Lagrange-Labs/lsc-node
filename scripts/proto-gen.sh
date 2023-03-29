set -e

protoVer=1.12.0
protoImageName=bufbuild/buf:${protoVer}

echo "Generating gogo proto code"
cd proto
proto_dirs=$(find ./ -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  for file in $(find "${dir}" -maxdepth 1 -name '*.proto'); do
    sudo rm -rf gen
    docker run --rm -v ./:/workspace --workdir /workspace ${protoImageName} generate  --template buf.gen.yaml $file
    sudo chmod -R 775 gen
    cp -r gen/go/* ../$(echo "${dir}" | sed 's#^\./\([^/]*\)/.*$#\1#')/types/
    sudo rm -rf gen
  done
done
