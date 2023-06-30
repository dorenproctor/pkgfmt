set -e

go build \
    -ldflags " \
    -X 'github.com/dorenproctor/pkgfmt/cmd.version=dev-version' \
    -X 'github.com/dorenproctor/pkgfmt/cmd.buildDate=$(date -u +"%Y-%m-%dT%H:%M:%SZ")' \
    " \
    -o bin/pkgfmt

bin/pkgfmt "$@"
