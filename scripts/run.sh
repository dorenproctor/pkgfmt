set -e

go build -o cmd/bin/github.com/dorenproctor/pkgfmt
cmd/bin/github.com/dorenproctor/pkgfmt "$@"
