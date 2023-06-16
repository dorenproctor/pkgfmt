set -e

go build -o bin/pkgfmt
bin/pkgfmt "$@"
