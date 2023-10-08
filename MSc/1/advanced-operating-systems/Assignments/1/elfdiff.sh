#!/usr/bin/env bash

diff <(readelf -a "$1") <(readelf -a "$2")
