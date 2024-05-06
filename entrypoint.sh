#!/usr/bin/sh

if [ "$1" != "" ]; then
  for link in "$@"; do
    if [ "$link" != "entrypoint.sh" ]; then
      ./go-anhgelus config config.toml "$link"
    fi
  done
fi

./go-anhgelus run
