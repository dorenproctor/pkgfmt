package fileutils

import (
	"fmt"
	"os/exec"
)

// diff two files using the diff command built into your shell
//
// only returns nil if no differences were found
func Diff(src, dst string) error {
	bytes, err := exec.Command("diff", src, dst).Output()
	if err != nil && err.Error() == "exit status 2" {
		return fmt.Errorf("'%s' or '%s' does not exist", src, dst)
	}
	if len(bytes) > 0 {
		return fmt.Errorf("%s", string(bytes))
	}
	return nil
}
