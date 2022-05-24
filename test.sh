#!/usr/bin/env bash
SCRIPT_PATH="$(
  cd "$(dirname "$0")"
  pwd -P
)"
cat "$SCRIPT_PATH"/test.json | "$SCRIPT_PATH"/flatten
