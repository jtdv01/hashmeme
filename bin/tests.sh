#!/usr/bin/env bash
set -e

git_root=$(git rev-parse --show-toplevel)

main () {
  rm -rf ./target
  mkdir ./target

  if command -v nix &> /dev/null
  then
    nix-shell --pure "${git_root}"/default.nix --run \
      "go test image_processor/*.go"
  else
    echo "Nix not found. See ./bin/install.sh."
    echo "Trying to build without nix, assumes you have go and dependencies installed..."
    go test "${git_root}"/image_processor/*.go
  fi
}

main
