#!/usr/bin/env bash
set -e

main() {
  # Check if nix is installed
  if command -v nix &> /dev/null
  then
    echo "Nix installed! Continuing..."
  else
    echo "Command nix nout found. Please install nix: https://nixos.org/download.html" \
    exit 1
  fi

  # Install nixGL for GUI
  echo "Installing nixGL for GUI"
  nix-channel --add https://github.com/guibou/nixGL/archive/main.tar.gz nixgl
  nix-channel --update
  nix-env -iA nixgl.auto.nixGLDefault

  nix-shell --pure --run "echo Installation successful!"
}

main