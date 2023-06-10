package fileutils

import (
	"fmt"
	"pkgfmt/cmd/utils/cmdutils"
)

func Diff(src, dst string) error {
	o := cmdutils.Run("diff", src, dst)
	s := o.Stdout + o.Stderr
	if s == "" {
		return nil
	}
	return fmt.Errorf(s)
}

// func Diff(src, dst string) (string, error) {
// 	bytes, err := exec.Command("diff", src, dst).Output()
// 	if err != nil && err.Error() == "exit status 2" {
// 		return string(bytes), fmt.Errorf("'%s' or '%s' does not exist", src, dst)
// 	}
// 	return string(bytes), err
// }
