#!/bin/bash

set -e

PROG=$(realpath "$1")

cd "$(dirname "$0")"

mpirun  --oversubscribe -n "$2" "$PROG" 4 6 5 -s 3,3 -v >res
trap 'rm res' EXIT

diff -Z out res
