# To avtivate:
# $ nix-build
# $ nix-shell --pure
{ pkgs ? import <nixpkgs> {} }:
pkgs.mkShell {
  buildInputs = [
    # https://github.com/guibou/nixGL
    pkgs.gcc
    pkgs.git
    pkgs.go
    pkgs.gtk3-x11
    pkgs.leptonica
    pkgs.mesa
    pkgs.neovim
    pkgs.tesseract4
    pkgs.which
    pkgs.xorg.libXcursor
    pkgs.xorg.libXi
    pkgs.xorg.libXinerama
    pkgs.xorg.libXrandr
    pkgs.xorg.libXxf86vm
  ];
  shellHook = ''
    EDITOR=nvim
  '';
}
