#!/usr/bin/sh

for link in $1; do
  ./go-anhgelus create --path config.toml --url "$link"
done
