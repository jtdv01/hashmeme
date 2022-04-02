# To avtivate:
# $ nix-build
# $ nix-shell --pure
{ pkgs ? import <nixpkgs> {} }:
pkgs.mkShell {
  buildInputs = [
    pkgs.go
    pkgs.neovim
    pkgs.which
  ];
  shellHook = ''
    EDITOR=nvim
  '';
}
