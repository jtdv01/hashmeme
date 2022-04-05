# To avtivate:
# $ nix-build
# $ nix-shell --pure
{ pkgs ? import <nixpkgs> {} }:
pkgs.mkShell {
  buildInputs = [
    pkgs.git
    pkgs.go
    pkgs.neovim
    pkgs.mesa
    pkgs.xorg.libX11
    pkgs.xorg.libX11.dev
    # pkgs.xorg.libXcursor
    # pkgs.xorg.libXrandr
    # pkgs.xorg.libXinerama
    pkgs.tesseract4
    pkgs.leptonica
    pkgs.which
  ];
  shellHook = ''
    EDITOR=nvim
  '';
}
