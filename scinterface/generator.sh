#!/bin/sh

set -e

gen() {
    local package=$1
    local dir=$2

    abigen --bin bin/${package}.bin --abi bin/${package}.abi --pkg=${dir} --out=${dir}/${package}.go
}

gen committee committee
gen committeet committeet
gen arbitrum_sequencer_inbox arbinbox
