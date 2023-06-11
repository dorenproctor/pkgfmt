set -e

go build -o cmd/bin/pkgfmt
cmd/bin/pkgfmt "$@"
