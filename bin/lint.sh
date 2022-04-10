#!/usr/bin/env bash
set -e

git_root=$(git rev-parse --show-toplevel)

main() {
    nix-shell --pure \
        --keep git_root \
        --run \
        """gofmt -w "${git_root}"/*/*.go;
        gofmt -w "${git_root}"/*.go"""
}

main
