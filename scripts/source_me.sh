alias r='./scripts/run.sh'
alias t='go test $(go list ./... | grep -v tests/input/)'
alias tt='t | grep FAIL'
alias ttt='OVERWRITE_TEST_EXPECTED_OUTPUT=true go test $(go list ./... | grep -v tests/input/)'
