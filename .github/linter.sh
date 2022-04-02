git_root=$(git rev-parse --show-toplevel)
gofmt -w "${git_root}"/**/*.go