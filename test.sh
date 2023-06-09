set -e

go test ./...

for f in ./tests/input/*; do
    diff "$f/expected_output" "$f/generated_pkgsplit"
    echo f $f
done
