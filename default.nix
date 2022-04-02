# To avtivate:
# $ nix-build
# $ nix-shell --pure
{ pkgs ? import <nixpkgs> {} }:
pkgs.mkShell {
  buildInputs = [
    pkgs.bazel
    pkgs.gh
    pkgs.git
    pkgs.go
    pkgs.neovim
    pkgs.which
  ];
  shellHook = ''
    EDITOR=nvim
  '';
}
