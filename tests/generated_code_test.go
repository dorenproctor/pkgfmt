package tests

import (
	"fmt"
	"os"
	"pkgsplit/cmd/pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratedCode(t *testing.T) {
	inputDir := "input/example"
	fileName := "x.go"
	inputFile := fmt.Sprintf("%s/%s", inputDir, fileName)
	p, err := pkg.New(inputFile)
	assert.NoError(t, err)

	err = os.RemoveAll(fmt.Sprintf("%s/generated_pkgsplit", inputDir))
	assert.NoError(t, err)

	assert.NoError(t, p.WriteOutput())
}
