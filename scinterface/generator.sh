#!/bin/sh

set -e

gen() {
    local package=$1

    ../scripts/abigen --bin bin/${package}.bin --abi bin/${package}.abi --pkg=${package} --out=${package}/${package}.go
}

if [ -z $1 ]
then
    echo "Usage: ./generator.sh [package]"
else
    gen $1
fi
