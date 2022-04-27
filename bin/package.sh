#!/usr/bin/env bash
set -e

git_root=$(git rev-parse --show-toplevel)

main() {
  local fyne_path="${git_root}/fyne"
  mkdir -p "${fyne_path}"
  rm -rf "${git_root}"/target

  nix-shell --pure \
    --keep fyne_path \
    --run "GOPATH=${fyne_path} go install fyne.io/fyne/v2/cmd/fyne@latest"

  echo "Packaging for linux..."
  nix-shell --pure \
    --keep fyne_path \
    --run "${fyne_path}/bin/fyne package --os linux --icon ${git_root}/image_processor/resources/hashmeme.png"

  # echo "Packaging for windows..."
  # nix-shell --pure \
  #   --keep fyne_path \
  #   --run "${fyne_path}/bin/fyne package --os windows --icon ${git_root}/image_processor/resources/hashmeme.png"

  echo "Packaging for android..."
  nix-shell --pure \
    --keep fyne_path \
    --run "${fyne_path}/bin/fyne package --os android --icon ${git_root}/image_processor/resources/hashmeme.png --appID com.jtdv01.hashmeme"
}

main