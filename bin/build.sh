#!/usr/bin/env bash
set -e

main () {
  git_root=$(git rev-parse --show-toplevel)
  rm -rf ./target
  mkdir ./target

  if command -v nix &> /dev/null
  then
    nix-shell --pure --keep GOROOT ./default.nix --run \
      "go build -o ./target"
  else
    echo "Nix not found. See ./bin/install.sh."
    echo "Trying to build without nix, assumes you have go and dependencies installed..."
    go build -o ./target
  fi

  cp "${git_root}"/image_processor/tesseract.ini ./target
  cp "${git_root}"/.env ./target
  cp "${git_root}"/image_processor/resources/*.png ./target

  echo "Build done in ./target/"
}

main
