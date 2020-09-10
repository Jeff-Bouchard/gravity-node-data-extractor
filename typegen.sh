#!/bin/bash

typegen_waves () {
    swagger generate model --target=$PWD/swagger-types --skip-validation --spec=https://nodes.wavesplatform.com/api-docs/swagger.json
}

while [ -n "$1" ]; do

    case "$1" in
        --waves) typegen_waves ;;
    esac
    shift
done
