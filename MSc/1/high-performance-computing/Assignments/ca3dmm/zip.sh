#!/bin/bash

set -e

cd "$(dirname "$0")"

md2pdf -o report.pdf README.md

rm hb417666.zip || :

zip hb417666.zip densematgen.cpp densematgen.h main.cpp params.cpp params.hpp CMakeLists.txt types.hpp README.md param_test.cpp report.pdf
