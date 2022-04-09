#!/usr/bin/env bash

main () {
  nix-shell --pure ./default.nix --run \
    "cd gui && \
    /nix/var/nix/profiles/per-user/${USER}/profile/bin/nixGL go run ./main.go"
}

main
