#!/usr/bin/env bash
set -e

main () {
    git_root=$(git rev-parse --show-toplevel)
    nix-shell --pure "${git_root}/default.nix" \
      --run \
      "go run ${git_root}/bin/bootstrap/create_topic.go"
}

main
