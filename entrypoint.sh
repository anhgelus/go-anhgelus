#!/usr/bin/sh

if [ "$1" != "" ]; then
  for link in $1; do
    ./go-anhgelus create --path config.toml --url "$link"
  done
fi
