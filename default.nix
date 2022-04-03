# To avtivate:
# $ nix-build
# $ nix-shell --pure
{ pkgs ? import <nixpkgs> {} }:
pkgs.mkShell {
  buildInputs = [
    pkgs.git
    pkgs.go
    pkgs.neovim
    pkgs.nodejs
    pkgs.tesseract4
    pkgs.leptonica
    pkgs.which
  ];
  shellHook = ''
    EDITOR=nvim
  '';
}
