#!/bin/bash

function checkTempFileError() {
  if [[ $1 != 0 ]]; then
    echo "Nie powiodło się stworzenie pliku tymczasowego"
    exit 1
  fi
}

function createTempFile() {
  local temp
  temp=$(mktemp)
  checkTempFileError $?
  echo "$temp"
}

function createTempDir() {
  local temp
  temp=$(mktemp -d)
  checkTempFileError $?
  echo "$temp"
}

if [[ $# != 3 ]]; then
  echo "Sposób uzytkowania: $0 <dojscie_do_main> <ścieżka/do/folderu/z/projektem> <ścieżka/do/folderu/z/testami>." >&2
  exit 1
fi

name=$1
java_dir=$2
dir=$3

if ! [[ -d "$dir" ]]; then
  echo "Podany folder z testami nie istnieje"
  exit 1
fi

dir=$(realpath "$dir")

total=0
correct=0

cd "$java_dir" || (
  echo "Podany folder projektu nie istnieje" &
  exit 1
)
tempFolder=$(createTempDir)
javac -d "$tempFolder" "$name.java"
cd "$tempFolder" || (
  echo "Podany folder projektu nie istnieje" &
  exit 1
)

out_file=$(createTempFile)

function run_test() {
  f="$1"
  ((total++))
  echo -e "\e[1mTest $f \e[0m"

  time java "$name" <"$f" >"$out_file"

  d_out=$(diff "${f%in}"out "$out_file")

  if [[ $d_out != "" ]]; then
    echo -e "\e[1;31m\tZły wynik testu\e[0m"
  else
    echo -e "\e[1;32m\tPoprawny wynik testu\e[0m"
    ((correct++))
  fi

  echo ""
}

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

traverse_folder "$dir"

echo -e "Poprawne \e[1m$correct\e[0m na \e[1m$total\e[0m testów"

if [[ $correct == "$total" ]]; then
  echo -e "\e[1;92mWszystko dobrze! \e[0m"
fi

rm "$out_file" -r "$tempFolder"
