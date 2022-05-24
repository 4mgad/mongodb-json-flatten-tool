#!/usr/bin/env bash
SCRIPT_PATH="$(
  cd "$(dirname "$0")"
  pwd -P
)"
go build -o "$SCRIPT_PATH"/flatten
