#!/bin/sh
# instruction to make sure that the script will exit immediately  if any command return none zer0 status
set -e

./bin/air -c .air.toml

# take all the parameter pass to the script and run it
exec "$@"
