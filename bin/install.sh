#!/usr/bin/env bash

main() {
  # Install nixGL for GUI
  nix-channel --add https://github.com/guibou/nixGL/archive/main.tar.gz nixgl && nix-channel --update
  nix-env -iA nixgl.auto.nixGLDefault

  nix-shell --pure --run "echo Installation succesful!"
}

main