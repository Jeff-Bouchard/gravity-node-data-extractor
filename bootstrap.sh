#!/bin/bash

maingo_entry="main.go"

grab_info () {
  # shellcheck disable=SC2034
  read -p "Enter extractor host: " -e host
  # shellcheck disable=SC2034
  read -p "Enter contact website: " -e website
  # shellcheck disable=SC2034
  read -p "Enter contact email: " -e email
}

main () {
  grab_info

  # shellcheck disable=SC2046

  IFS=$'\n'
  ln=$(cat -s template/swagger-meta.txt)

  # shellcheck disable=SC2059
  metadata=$(printf "$ln" "$host" "$website" "$email" "$website")
  maingo=$(cat -v $maingo_entry)
  rm "$maingo_entry"

  echo "$metadata" >> $maingo_entry
  echo "$maingo" >> $maingo_entry
}

# shellcheck disable=SC2068
main $@