#!/usr/bin/env bash
SCRIPT_PATH="$(
  cd "$(dirname "$0")"
  pwd -P
)"
go test -v "$SCRIPT_PATH"/...