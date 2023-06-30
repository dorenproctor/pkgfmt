set -e

go build \
    -ldflags " \
    -X 'github.com/dorenproctor/pkgfmt/cmd.version=dev-version' \
    -X 'github.com/dorenproctor/pkgfmt/cmd.buildDate=$(date)' \
    " \
    -o bin/pkgfmt

bin/pkgfmt "$@"
