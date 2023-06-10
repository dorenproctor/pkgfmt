set -e

go build -o cmd/bin/pksplit
cmd/bin/pksplit "$@"