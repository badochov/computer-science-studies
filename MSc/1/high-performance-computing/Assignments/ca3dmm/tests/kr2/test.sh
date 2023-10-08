#!/bin/bash

set -e

PROG=$(realpath "$1")

cd "$(dirname "$0")"

mpirun  --oversubscribe -n "$2" "$PROG" 5 100 100 -s 28,31,87,15,12,44,56,29 -g 25 >res 2> /dev/null
trap 'rm res' EXIT

diff -Z out res
