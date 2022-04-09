#!/usr/bin/env bash

main () {
  # export GOROOT=`pwd`:${GOROOT}
  nix-shell --pure --keep GOROOT ./default.nix --run \
    "/nix/var/nix/profiles/per-user/${USER}/profile/bin/nixGL go run ./main.go"
}

main
