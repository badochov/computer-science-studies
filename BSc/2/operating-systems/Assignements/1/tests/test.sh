#!/bin/bash

if [[ $# != 1 ]]; then
    echo "Sposób uzytkowania: $0 <ścieżka/do/folderu/z/testami>" >&2
    exit 1
fi

tests=$(realpath "$1")

if ! [[ -d "$tests" ]]; then
  echo "Podany folder z testami nie istnieje"
  exit 1
fi


total=0
correct=0

function traverse_folder() {
  folder="$1"
  shopt -s nullglob
  for f in "$folder"/*.in; do
    run_test "$f"
  done

  shopt -s nullglob
  for d in "$folder"/*/; do
    echo "$d"
    traverse_folder "$(realpath "$d")"
  done
}

RED='\033[0;31m'
GREEN='\033[0;32m'
NOCOLOR='\033[0m'

function run_test() {
  input_file="$1"
  output_file=${input_file//.in/.out}

  ((total++))
  echo -e "\e[1mTest $f \e[0m"

  sh "$input_file" > "$temp_out"

  if cmp -s "$output_file" "$temp_out"; then
    echo -ne "${GREEN}stdout ok, ${NOCOLOR}"
    ((correct++))
  else
    echo -ne "${RED}stdout nieprawidlowe, ${NOCOLOR}"
    diff -d "$output_file" "$temp_out"
  fi
}

temp_out=$(mktemp)
temp_err=$(mktemp)
trap 'rm -f "$temp_out"' INT TERM HUP EXIT
trap 'rm -f "$temp_err"' INT TERM HUP EXIT

traverse_folder "$tests"

echo -e "Poprawne \e[1m$correct\e[0m na \e[1m$total\e[0m testów"

if [[ $correct == "$total" ]]; then
  echo -e "\e[1;92mWszystko dobrze! \e[0m"
fi
