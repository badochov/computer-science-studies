#!/bin/bash

set -e

cd -P -- "$(dirname -- "${BASH_SOURCE[0]}")"

make -C ../build clean all

cp ../build/converter .