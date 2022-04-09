#!/usr/bin/env bash

main () {
  export NIX_USER_PATH=/nix/var/nix/profiles/per-user/"${USER}"/profile/bin
  nix-shell --pure \
    --keep GOROOT --keep PATH \
    ./default.nix --run \
    "${NIX_USER_PATH}/nixGL go run ./main.go"
}

main
