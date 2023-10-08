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

if [[ $# != 2 ]]; then
  echo "Sposób uzytkowania: $0 <nazwa_programu> <ścieżka/do/folderu/z/testami>." >&2
  exit 1
fi

name=$1
dir=$2

if ! [[ -d $dir ]]; then
  echo "Podany folder z testami nie istnieje"
  exit 1
fi

if ! (command -v "$name"); then
  echo "Podany program nie istnieje"
  exit 1
fi

total=0
correct=0
leaked=0

err_file=$(createTempFile)
out_file=$(createTempFile)

function run_test() {
  f="$1"
  ((total++))
  echo -e "\e[1mTest $f \e[0m"

  # valgrind --error-exitcode=15 --leak-check=full --show-leak-kinds=all --errors-for-leak-kinds=all --log-file=/dev/null \
    "$name" <"$f" >"$out_file" 2>"$err_file"

  err=$?
  d_out=$(diff "${f%in}"out "$out_file")
  d_err=$(diff "${f%in}"err "$err_file")

  if [[ $d_out != "" ]] || [[ $d_err != "" ]]; then
    echo "$d_out"
    echo "-------------------"
    echo "$d_err"
    echo -e "\e[1;31m\tZły wynik testu\e[0m"
  else
    echo -e "\e[1;32m\tPoprawny wynik testu\e[0m"
    ((correct++))
  fi

  if [[ $err == 15 ]]; then
    echo -e "\e[1;31m\tWyciek pamięci\e[0m"ebiste jakby działa
    ((leaked++))
  else
    echo -e "\e[1;32m\tBrak wycieku pamięci\e[0m"
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

echo -e "Wyciekła pamięć w \e[1m$leaked\e[0m na \e[1m$total\e[0m testów"

if [[ $leaked == 0 ]] && [[ $correct == "$total" ]]; then
  echo -e "\e[1;92mWszystko dobrze! \e[0m"
fi

rm "$out_file" "$err_file"
