git_root=$(git rev-parse --show-toplevel)
gofmt -w "${git_root}"/*/*.go
gofmt -w "${git_root}"/*.go
